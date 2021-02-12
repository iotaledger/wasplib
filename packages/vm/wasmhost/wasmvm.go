// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package wasmhost

import (
	"encoding/binary"
	"fmt"
)

type WasmVM interface {
	LinkHost(impl WasmVM, host *WasmHost) error
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

func (vm *WasmVmBase) LinkHost(impl WasmVM, host *WasmHost) error {
	vm.impl = impl
	vm.host = host
	host.vm = impl
	return nil
}

func (vm *WasmVmBase) HostFdWrite(fd int32, iovs int32, size int32, written int32) int32 {
	vm.host.TraceAll("HostFdWrite(...)")
	// very basic implementation that expects fd to be stdout and iovs to be only one element
	ptr := vm.impl.UnsafeMemory()
	txt := binary.LittleEndian.Uint32(ptr[iovs : iovs+4])
	siz := binary.LittleEndian.Uint32(ptr[iovs+4 : iovs+8])
	fmt.Print(string(ptr[txt : txt+siz]))
	binary.LittleEndian.PutUint32(ptr[written:written+4], siz)
	return int32(siz)
}

func (vm *WasmVmBase) HostGetBytes(objId int32, keyId int32, typeId int32, stringRef int32, size int32) int32 {
	host := vm.host
	host.TraceAll("HostGetBytes(o%d,k%d,t%d,r%d,s%d)", objId, keyId, typeId, stringRef, size)

	// negative size means only check for existence
	if size < 0 {
		if host.Exists(objId, keyId, typeId) {
			return 0
		}
		// missing key is indicated by -1
		return -1
	}

	bytes := host.GetBytes(objId, keyId, typeId)
	if bytes == nil {
		return -1
	}
	return vm.vmSetBytes(stringRef, size, bytes)
}

func (vm *WasmVmBase) HostGetKeyId(keyRef int32, size int32) int32 {
	host := vm.host
	host.TraceAll("HostGetKeyId(r%d,s%d)", keyRef, size)
	// non-negative size means original key was a string
	if size >= 0 {
		bytes := vm.vmGetBytes(keyRef, size)
		return host.GetKeyIdFromString(string(bytes))
	}

	// negative size means original key was a byte slice
	bytes := vm.vmGetBytes(keyRef, -size-1)
	return host.GetKeyIdFromBytes(bytes)
}

func (vm *WasmVmBase) HostGetObjectId(objId int32, keyId int32, typeId int32) int32 {
	host := vm.host
	host.TraceAll("HostGetObjectId(o%d,k%d,t%d)", objId, keyId, typeId)
	return host.GetObjectId(objId, keyId, typeId)
}

func (vm *WasmVmBase) HostSetBytes(objId int32, keyId int32, typeId int32, stringRef int32, size int32) {
	host := vm.host
	host.TraceAll("HostSetBytes(o%d,k%d,t%d,r%d,s%d)", objId, keyId, typeId, stringRef, size)
	bytes := vm.vmGetBytes(stringRef, size)
	host.SetBytes(objId, keyId, typeId, bytes)
}

func (vm *WasmVmBase) PreCall() []byte {
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

func (vm *WasmVmBase) PostCall(frame []byte) {
	ptr := vm.impl.UnsafeMemory()
	copy(ptr, frame)
}

func (vm *WasmVmBase) SaveMemory() {
	// find initialized data range in memory
	ptr := vm.impl.UnsafeMemory()
	if ptr == nil {
		// this vm implementation does not communicate via mem pool
		return
	}
	firstNonZero := -1
	lastNonZero := 0
	for i, b := range ptr {
		if b != 0 {
			if firstNonZero < 0 {
				firstNonZero = i
			}
			lastNonZero = i
		}
	}

	// save copy of initialized data range
	vm.memoryNonZero = len(ptr)
	if firstNonZero >= 0 {
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

func (vm *WasmVmBase) vmSetBytes(offset int32, size int32, bytes []byte) int32 {
	if size != 0 {
		ptr := vm.impl.UnsafeMemory()
		copy(ptr[offset:offset+size], bytes)
	}
	return int32(len(bytes))
}
