package wasmhost

import (
	"errors"
	"github.com/bytecodealliance/wasmtime-go"
	"strings"
)

type WasmTimeVM struct {
	instance *wasmtime.Instance
	linker   *wasmtime.Linker
	memory   *wasmtime.Memory
	module   *wasmtime.Module
	store    *wasmtime.Store
}

func NewWasmTimeVM() *WasmTimeVM {
	host := &WasmTimeVM{}
	host.store = wasmtime.NewStore(wasmtime.NewEngine())
	host.linker = wasmtime.NewLinker(host.store)
	return host
}

func (vm *WasmTimeVM) LinkHost(host *WasmHost) error {
	var externals = map[string]interface{}{
		"wasplib.hostGetBytes": func(objId int32, keyId int32, stringRef int32, size int32) int32 {
			if objId >= 0 {
				return host.vmSetBytes(stringRef, size, host.GetBytes(objId, keyId))
			}
			return host.vmSetBytes(stringRef, size, []byte(host.GetString(-objId, keyId)))
		},
		"wasplib.hostGetInt": func(objId int32, keyId int32) int64 {
			return host.GetInt(objId, keyId)
		},
		"wasplib.hostGetIntRef": func(objId int32, keyId int32, intRef int32) {
			host.vmSetInt(intRef, host.GetInt(objId, keyId))
		},
		"wasplib.hostGetKeyId": func(keyRef int32, size int32) int32 {
			return host.GetKeyId(string(host.vmGetBytes(keyRef, size)))
		},
		"wasplib.hostGetObjectId": func(objId int32, keyId int32, typeId int32) int32 {
			return host.GetObjectId(objId, keyId, typeId)
		},
		"wasplib.hostSetBytes": func(objId int32, keyId int32, stringRef int32, size int32) {
			if objId >= 0 {
				host.SetBytes(objId, keyId, host.vmGetBytes(stringRef, size))
				return
			}
			host.SetString(-objId, keyId, string(host.vmGetBytes(stringRef, size)))
		},
		"wasplib.hostSetInt": func(objId int32, keyId int32, value int64) {
			host.SetInt(objId, keyId, value)
		},
		"wasplib.hostSetIntRef": func(objId int32, keyId int32, intRef int32) {
			host.SetInt(objId, keyId, host.vmGetInt(intRef))
		},
		//TODO: go implementation uses this one to write panic message
		"wasi_unstable.fd_write": func(fd int32, iovs int32, size int32, written int32) int32 {
			return host.fdWrite(fd, iovs, size, written)
		},
	}

	for name, function := range externals {
		names := strings.Split(name, ".")
		err := vm.linker.DefineFunc(names[0], names[1], function)
		if err != nil {
			return err
		}
	}
	return nil
}

func (vm *WasmTimeVM) LoadWasm(wasmData []byte) error {
	var err error
	vm.module, err = wasmtime.NewModule(vm.store.Engine, wasmData)
	if err != nil {
		return err
	}
	vm.instance, err = vm.linker.Instantiate(vm.module)
	if err != nil {
		return err
	}
	memory := vm.instance.GetExport("memory")
	if memory == nil {
		return errors.New("no memory export")
	}
	vm.memory = memory.Memory()
	if vm.memory == nil {
		return errors.New("not a memory type")
	}
	return nil
}

func (vm *WasmTimeVM) RunFunction(functionName string) error {
	function := vm.instance.GetExport(functionName).Func()
	_, err := function.Call()
	return err
}

func (vm *WasmTimeVM) UnsafeMemory() []byte {
	return vm.memory.UnsafeData()
}
