// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package wasmlocalhost

import (
	"github.com/iotaledger/wart/wasm/consts/value"
	"github.com/iotaledger/wart/wasm/executors"
	"github.com/iotaledger/wart/wasm/sections"
	"github.com/iotaledger/wasp/packages/vm/wasmhost"
)

type WartVM struct {
	wasmhost.WasmVmBase
	runner *executors.WasmRunner
}

func NewWartVM() *WartVM {
	vm := &WartVM{}
	vm.runner = executors.NewWasmRunner()
	return vm
}

func (vm *WartVM) LinkHost(impl wasmhost.WasmVM, host *wasmhost.WasmHost) error {
	vm.WasmVmBase.LinkHost(impl, host)
	m := executors.DefineModule("wasplib")
	lnk := executors.NewWasmLinker(m)
	_ = lnk.DefineFunction("hostGetBytes",
		[]value.DataType{value.I32, value.I32, value.I32, value.I32, value.I32},
		[]value.DataType{value.I32},
		func(ctx *sections.HostContext) error {
			objId := ctx.Frame[ctx.SP].I32
			keyId := ctx.Frame[ctx.SP+1].I32
			typeId := ctx.Frame[ctx.SP+2].I32
			stringRef := ctx.Frame[ctx.SP+3].I32
			size := ctx.Frame[ctx.SP+4].I32
			ctx.Frame[ctx.SP].I32 = vm.HostGetBytes(objId, keyId, typeId, stringRef, size)
			return nil
		})
	_ = lnk.DefineFunction("hostGetKeyId",
		[]value.DataType{value.I32, value.I32},
		[]value.DataType{value.I32},
		func(ctx *sections.HostContext) error {
			keyRef := ctx.Frame[ctx.SP].I32
			size := ctx.Frame[ctx.SP+1].I32
			ctx.Frame[ctx.SP].I32 = vm.HostGetKeyId(keyRef, size)
			return nil
		})
	_ = lnk.DefineFunction("hostGetObjectId",
		[]value.DataType{value.I32, value.I32, value.I32},
		[]value.DataType{value.I32},
		func(ctx *sections.HostContext) error {
			objId := ctx.Frame[ctx.SP].I32
			keyId := ctx.Frame[ctx.SP+1].I32
			typeId := ctx.Frame[ctx.SP+2].I32
			ctx.Frame[ctx.SP].I32 = vm.HostGetObjectId(objId, keyId, typeId)
			return nil
		})
	_ = lnk.DefineFunction("hostSetBytes",
		[]value.DataType{value.I32, value.I32, value.I32, value.I32},
		nil,
		func(ctx *sections.HostContext) error {
			objId := ctx.Frame[ctx.SP].I32
			keyId := ctx.Frame[ctx.SP+1].I32
			typeId := ctx.Frame[ctx.SP+2].I32
			stringRef := ctx.Frame[ctx.SP+3].I32
			size := ctx.Frame[ctx.SP+4].I32
			vm.HostSetBytes(objId, keyId, typeId, stringRef, size)
			return nil
		})
	// go implementation uses this one to write panic message
	m = executors.DefineModule("wasi_unstable")
	lnk = executors.NewWasmLinker(m)
	_ = lnk.DefineFunction("fd_write",
		[]value.DataType{value.I32, value.I32, value.I32, value.I32},
		[]value.DataType{value.I32},
		func(ctx *sections.HostContext) error {
			fd := ctx.Frame[ctx.SP].I32
			iovs := ctx.Frame[ctx.SP+1].I32
			size := ctx.Frame[ctx.SP+2].I32
			written := ctx.Frame[ctx.SP+3].I32
			ctx.Frame[ctx.SP].I32 = vm.HostFdWrite(fd, iovs, size, written)
			return nil
		})
	return nil
}

func (vm *WartVM) LoadWasm(wasmData []byte) error {
	return vm.runner.Load(wasmData)
}

func (vm *WartVM) RunFunction(functionName string) error {
	return vm.runner.RunExport(functionName, nil)
}

func (vm *WartVM) RunScFunction(index int32) error {
	params := make([]sections.Variable, 1)
	params[0].I32 = index
	frame := vm.PreCall()
	err := vm.runner.RunExport("on_call_entrypoint", params)
	vm.PostCall(frame)
	return err
}

func (vm *WartVM) UnsafeMemory() []byte {
	return vm.runner.Memory()
}
