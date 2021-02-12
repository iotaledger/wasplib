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
	vm := &WasmTimeVM{}
	vm.store = wasmtime.NewStore(wasmtime.NewEngine())
	vm.linker = wasmtime.NewLinker(vm.store)
	return vm
}

func (vm *WasmTimeVM) LinkHost(impl WasmVM, host *WasmHost) error {
	vm.WasmVmBase.LinkHost(impl, host)
	err := vm.linker.DefineFunc("wasplib", "hostGetBytes",
		func(objId int32, keyId int32, typeId int32, stringRef int32, size int32) int32 {
			return vm.HostGetBytes(objId, keyId, typeId, stringRef, size)
		})
	if err != nil {
		return err
	}
	err = vm.linker.DefineFunc("wasplib", "hostGetKeyId",
		func(keyRef int32, size int32) int32 {
			return vm.HostGetKeyId(keyRef, size)
		})
	if err != nil {
		return err
	}
	err = vm.linker.DefineFunc("wasplib", "hostGetObjectId",
		func(objId int32, keyId int32, typeId int32) int32 {
			return vm.HostGetObjectId(objId, keyId, typeId)
		})
	if err != nil {
		return err
	}
	err = vm.linker.DefineFunc("wasplib", "hostSetBytes",
		func(objId int32, keyId int32, typeId int32, stringRef int32, size int32) {
			vm.HostSetBytes(objId, keyId, typeId, stringRef, size)
		})
	if err != nil {
		return err
	}
	// go implementation uses this one to write panic message
	err = vm.linker.DefineFunc("wasi_unstable", "fd_write",
		func(fd int32, iovs int32, size int32, written int32) int32 {
			return vm.HostFdWrite(fd, iovs, size, written)
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
	export := vm.instance.GetExport("on_call_entrypoint")
	if export == nil {
		return errors.New("unknown export function: 'on_call_entrypoint'")
	}
	frame := vm.PreCall()
	_, err := export.Func().Call(index)
	vm.PostCall(frame)
	return err
}

func (vm *WasmTimeVM) UnsafeMemory() []byte {
	return vm.memory.UnsafeData()
}
