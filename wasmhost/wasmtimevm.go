// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package wasmhost

import (
	"errors"
	"github.com/bytecodealliance/wasmtime-go"
)

type WasmTimeVM struct {
	WasmVmBase
	instance *wasmtime.Instance
	linker   *wasmtime.Linker
	memory   *wasmtime.Memory
	module   *wasmtime.Module
	store    *wasmtime.Store
}

func NewWasmTimeVM() *WasmTimeVM {
	host := &WasmTimeVM{}
	host.impl = host
	host.store = wasmtime.NewStore(wasmtime.NewEngine())
	host.linker = wasmtime.NewLinker(host.store)
	return host
}

func (vm *WasmTimeVM) LinkHost(host *WasmHost) error {
	vm.host = host
	err := vm.linker.DefineFunc("wasplib", "hostGetBytes",
		func(objId int32, keyId int32, stringRef int32, size int32) int32 {
			return vm.hostGetBytes(objId, keyId, stringRef, size)
		})
	if err != nil {
		return err
	}
	err = vm.linker.DefineFunc("wasplib", "hostGetInt",
		func(objId int32, keyId int32) int64 {
			return vm.hostGetInt(objId, keyId)
		})
	if err != nil {
		return err
	}
	err = vm.linker.DefineFunc("wasplib", "hostGetIntRef",
		func(objId int32, keyId int32, intRef int32) {
			vm.hostGetIntRef(objId, keyId, intRef)
		})
	if err != nil {
		return err
	}
	err = vm.linker.DefineFunc("wasplib", "hostGetKeyId",
		func(keyRef int32, size int32) int32 {
			return vm.hostGetKeyId(keyRef, size)
		})
	if err != nil {
		return err
	}
	err = vm.linker.DefineFunc("wasplib", "hostGetObjectId",
		func(objId int32, keyId int32, typeId int32) int32 {
			return vm.hostGetObjectId(objId, keyId, typeId)
		})
	if err != nil {
		return err
	}
	err = vm.linker.DefineFunc("wasplib", "hostSetBytes",
		func(objId int32, keyId int32, stringRef int32, size int32) {
			vm.hostSetBytes(objId, keyId, stringRef, size)
		})
	if err != nil {
		return err
	}
	err = vm.linker.DefineFunc("wasplib", "hostSetInt",
		func(objId int32, keyId int32, value int64) {
			vm.hostSetInt(objId, keyId, value)
		})
	if err != nil {
		return err
	}
	err = vm.linker.DefineFunc("wasplib", "hostSetIntRef",
		func(objId int32, keyId int32, intRef int32) {
			vm.hostSetIntRef(objId, keyId, intRef)
		})
	if err != nil {
		return err
	}
	// go implementation uses this one to write panic message
	err = vm.linker.DefineFunc("wasi_unstable", "fd_write",
		func(fd int32, iovs int32, size int32, written int32) int32 {
			return vm.hostFdWrite(fd, iovs, size, written)
		})
	if err != nil {
		return err
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
	export := vm.instance.GetExport(functionName)
	if export == nil {
		return errors.New("unknown export function: '" + functionName + "'")
	}
	_, err := export.Func().Call()
	return err
}

func (vm *WasmTimeVM) RunScFunction(index int32) error {
	export := vm.instance.GetExport("sc_call_entrypoint")
	if export == nil {
		return errors.New("unknown export function: 'sc_call_entrypoint'")
	}
	frame := vm.preCall()
	_, err := export.Func().Call(index)
	vm.postCall(frame)
	return err
}

func (vm *WasmTimeVM) UnsafeMemory() []byte {
	return vm.memory.UnsafeData()
}
