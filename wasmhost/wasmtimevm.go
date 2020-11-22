// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package wasmhost

import (
	"errors"
	"github.com/bytecodealliance/wasmtime-go"
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
	err := vm.linker.DefineFunc("wasplib", "hostGetBytes",
		func(objId int32, keyId int32, stringRef int32, size int32) int32 {
			vm.traceHost(host, "hostGetBytes")
			return host.GetBytes(objId, keyId, stringRef, size)
		})
	if err != nil {
		return err
	}
	err = vm.linker.DefineFunc("wasplib", "hostGetInt",
		func(objId int32, keyId int32) int64 {
			vm.traceHost(host, "hostGetInt")
			return host.GetInt(objId, keyId)
		})
	if err != nil {
		return err
	}
	err = vm.linker.DefineFunc("wasplib", "hostGetIntRef",
		func(objId int32, keyId int32, intRef int32) {
			vm.traceHost(host, "hostGetIntRef")
			host.vmSetInt(intRef, host.GetInt(objId, keyId))
		})
	if err != nil {
		return err
	}
	err = vm.linker.DefineFunc("wasplib", "hostGetKeyId",
		func(keyRef int32, size int32) int32 {
			vm.traceHost(host, "hostGetKeyId")
			return host.GetKeyId(keyRef, size)
		})
	if err != nil {
		return err
	}
	err = vm.linker.DefineFunc("wasplib", "hostGetObjectId",
		func(objId int32, keyId int32, typeId int32) int32 {
			vm.traceHost(host, "hostGetObjectId")
			return host.GetObjectId(objId, keyId, typeId)
		})
	if err != nil {
		return err
	}
	err = vm.linker.DefineFunc("wasplib", "hostSetBytes",
		func(objId int32, keyId int32, stringRef int32, size int32) {
			vm.traceHost(host, "hostSetBytes")
			host.SetBytes(objId, keyId, stringRef, size)
		})
	if err != nil {
		return err
	}
	err = vm.linker.DefineFunc("wasplib", "hostSetInt",
		func(objId int32, keyId int32, value int64) {
			vm.traceHost(host, "hostSetInt")
			host.SetInt(objId, keyId, value)
		})
	if err != nil {
		return err
	}
	err = vm.linker.DefineFunc("wasplib", "hostSetIntRef",
		func(objId int32, keyId int32, intRef int32) {
			vm.traceHost(host, "hostSetIntRef")
			host.SetInt(objId, keyId, host.vmGetInt(intRef))
		})
	if err != nil {
		return err
	}
	// go implementation uses this one to write panic message
	err = vm.linker.DefineFunc("wasi_unstable", "fd_write",
		func(fd int32, iovs int32, size int32, written int32) int32 {
			return host.fdWrite(fd, iovs, size, written)
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
	function := export.Func()
	_, err := function.Call()
	return err
}

func (vm *WasmTimeVM) RunScFunction(index int32) error {
	export := vm.instance.GetExport("sc_call_entrypoint")
	if export == nil {
		return errors.New("unknown export function: 'sc_call_entrypoint'")
	}
	function := export.Func()
	_, err := function.Call(index)
	return err
}

func (vm *WasmTimeVM) traceHost(host *WasmHost, text string) {
	host.logger.Log(KeyTraceHost, text)
}

func (vm *WasmTimeVM) UnsafeMemory() []byte {
	return vm.memory.UnsafeData()
}
