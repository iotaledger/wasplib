package wasmhost

import (
	"encoding/binary"
	"fmt"
	"github.com/mr-tron/base58"
)

type WasmVM interface {
	LinkHost(vm *WasmHost) error
	LoadWasm(wasmData []byte) error
	RunFunction(functionName string) error
	RunScFunction(index int32) error
	UnsafeMemory() []byte
	SaveMemory()
}

type WasmVmBase struct {
	impl          WasmVM
	host          *WasmHost
	memoryCopy    []byte
	memoryDirty   bool
	memoryNonZero int
}

func (vm *WasmVmBase) hostFdWrite(fd int32, iovs int32, size int32, written int32) int32 {
	host := vm.host
	host.TraceHost("hostFdWrite(...)")
	// very basic implementation that expects fd to be stdout and iovs to be only one element
	ptr := vm.impl.UnsafeMemory()
	txt := binary.LittleEndian.Uint32(ptr[iovs : iovs+4])
	siz := binary.LittleEndian.Uint32(ptr[iovs+4 : iovs+8])
	fmt.Print(string(ptr[txt : txt+siz]))
	binary.LittleEndian.PutUint32(ptr[written:written+4], siz)
	return int32(siz)
}

func (vm *WasmVmBase) hostGetBytes(objId int32, keyId int32, stringRef int32, size int32) int32 {
	host := vm.host
	host.TraceHost("hostGetBytes(o%d,k%d,r%d,s%d)", objId, keyId, stringRef, size)
	if objId < 0 {
		// negative objId means get string
		value := host.getString(-objId, keyId)
		// missing key is indicated by -1
		if value == nil {
			return -1
		}
		return vm.vmSetBytes(stringRef, size, []byte(*value))
	}

	bytes := host.GetBytes(objId, keyId)
	if bytes == nil {
		return -1
	}
	return vm.vmSetBytes(stringRef, size, bytes)
}

func (vm *WasmVmBase) hostGetInt(objId int32, keyId int32) int64 {
	host := vm.host
	host.TraceHost("hostGetInt(o%d,k%d)", objId, keyId)
	return host.GetInt(objId, keyId)
}

func (vm *WasmVmBase) hostGetIntRef(objId int32, keyId int32, intRef int32) {
	host := vm.host
	host.TraceHost("hostGetIntRef(o%d,k%d,r%d)", objId, keyId, intRef)
	vm.vmSetInt(intRef, host.GetInt(objId, keyId))
}

func (vm *WasmVmBase) hostGetKeyId(keyRef int32, size int32) int32 {
	host := vm.host
	host.TraceHost("hostGetKeyId(r%d,s%d)", keyRef, size)
	// non-negative size means original key was a string
	if size >= 0 {
		bytes := vm.vmGetBytes(keyRef, size)
		return host.GetKeyId(string(bytes))
	}

	// negative size means original key was a byte slice
	bytes := vm.vmGetBytes(keyRef, -size-1)
	if !host.useBase58Keys {
		// use byte slice key as is
		return host.GetKey(bytes)
	}

	// transform byte slice key into base58 string
	// now all keys are byte slices from strings
	return host.GetKeyId(base58.Encode(bytes))
}

func (vm *WasmVmBase) hostGetObjectId(objId int32, keyId int32, typeId int32) int32 {
	host := vm.host
	host.TraceHost("hostGetObjectId(o%d,k%d,t%d)", objId, keyId, typeId)
	return host.GetObjectId(objId, keyId, typeId)
}

func (vm *WasmVmBase) hostSetBytes(objId int32, keyId int32, stringRef int32, size int32) {
	host := vm.host
	host.TraceHost("hostSetBytes(o%d,k%d,r%d,s%d)", objId, keyId, stringRef, size)
	bytes := vm.vmGetBytes(stringRef, size)
	if objId < 0 {
		host.SetString(-objId, keyId, string(bytes))
		return
	}
	host.SetBytes(objId, keyId, bytes)
}

func (vm *WasmVmBase) hostSetInt(objId int32, keyId int32, value int64) {
	host := vm.host
	host.TraceHost("hostSetInt(o%d,k%d,v%d)", objId, keyId, value)
	host.SetInt(objId, keyId, value)
}

func (vm *WasmVmBase) hostSetIntRef(objId int32, keyId int32, intRef int32) {
	host := vm.host
	host.TraceHost("hostSetIntRef(o%d,k%d,r%d)", objId, keyId, intRef)
	host.SetInt(objId, keyId, vm.vmGetInt(intRef))
}

func (vm *WasmVmBase) preCall() []byte {
	ptr := vm.impl.UnsafeMemory()
	frame := make([]byte, len(ptr))
	copy(frame, ptr)
	if vm.memoryDirty {
		// clear memory and restore initialized data range
		copy(ptr, make([]byte, len(ptr)))
		copy(ptr[vm.memoryNonZero:], vm.memoryCopy)
	}
	vm.memoryDirty = true
	return frame
}

func (vm *WasmVmBase) postCall(frame []byte) {
	ptr := vm.impl.UnsafeMemory()
	copy(ptr, frame)
}

func (vm *WasmVmBase) SaveMemory() {
	// find initialized data range in memory
	ptr := vm.impl.UnsafeMemory()
	firstNonZero := 0
	lastNonZero := 0
	for i, b := range ptr {
		if b != 0 {
			if firstNonZero == 0 {
				firstNonZero = i
			}
			lastNonZero = i
		}
	}

	// save copy of initialized data range
	vm.memoryNonZero = len(ptr)
	if ptr[firstNonZero] != 0 {
		vm.memoryNonZero = firstNonZero
		size := lastNonZero + 1 - firstNonZero
		vm.memoryCopy = make([]byte, size)
		copy(vm.memoryCopy, ptr[vm.memoryNonZero:])
	}
}

func (vm *WasmVmBase) vmGetBytes(offset int32, size int32) []byte {
	ptr := vm.impl.UnsafeMemory()
	bytes := make([]byte, size)
	copy(bytes, ptr[offset:offset+size])
	return bytes
}

func (vm *WasmVmBase) vmGetInt(offset int32) int64 {
	ptr := vm.impl.UnsafeMemory()
	return int64(binary.LittleEndian.Uint64(ptr[offset : offset+8]))
}

func (vm *WasmVmBase) vmSetBytes(offset int32, size int32, bytes []byte) int32 {
	if size != 0 {
		ptr := vm.impl.UnsafeMemory()
		copy(ptr[offset:offset+size], bytes)
	}
	return int32(len(bytes))
}

func (vm *WasmVmBase) vmSetInt(offset int32, value int64) {
	ptr := vm.impl.UnsafeMemory()
	binary.LittleEndian.PutUint64(ptr[offset:offset+8], uint64(value))
}
