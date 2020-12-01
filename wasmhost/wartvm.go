// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package wasmhost

import (
	"github.com/iotaledger/wart/wasm/consts/value"
	"github.com/iotaledger/wart/wasm/executors"
	"github.com/iotaledger/wart/wasm/sections"
)

type WartVM struct {
	WasmVmBase
	runner *executors.WasmRunner
}

func NewWartVM() *WartVM {
	host := &WartVM{}
	host.impl = host
	host.runner = executors.NewWasmRunner()
	return host
}

func (vm *WartVM) LinkHost(host *WasmHost) error {
	vm.host = host
	m := executors.DefineModule("wasplib")
	lnk := executors.NewWasmLinker(m)
	_ = lnk.DefineFunction("hostGetBytes",
		[]value.DataType{value.I32, value.I32, value.I32, value.I32},
		[]value.DataType{value.I32},
		func(ctx *sections.HostContext) error {
			objId := ctx.Frame[ctx.SP].I32
			keyId := ctx.Frame[ctx.SP+1].I32
			stringRef := ctx.Frame[ctx.SP+2].I32
			size := ctx.Frame[ctx.SP+3].I32
			ctx.Frame[ctx.SP].I32 = vm.hostGetBytes(objId, keyId, stringRef, size)
			return nil
		})
	_ = lnk.DefineFunction("hostGetInt",
		[]value.DataType{value.I32, value.I32},
		[]value.DataType{value.I64},
		func(ctx *sections.HostContext) error {
			objId := ctx.Frame[ctx.SP].I32
			keyId := ctx.Frame[ctx.SP+1].I32
			ctx.Frame[ctx.SP].I64 = vm.hostGetInt(objId, keyId)
			return nil
		})
	_ = lnk.DefineFunction("hostGetIntRef",
		[]value.DataType{value.I32, value.I32, value.I32},
		nil,
		func(ctx *sections.HostContext) error {
			objId := ctx.Frame[ctx.SP].I32
			keyId := ctx.Frame[ctx.SP+1].I32
			intRef := ctx.Frame[ctx.SP+2].I32
			vm.hostGetIntRef(objId, keyId, intRef)
			return nil
		})
	_ = lnk.DefineFunction("hostGetKeyId",
		[]value.DataType{value.I32, value.I32},
		[]value.DataType{value.I32},
		func(ctx *sections.HostContext) error {
			keyRef := ctx.Frame[ctx.SP].I32
			size := ctx.Frame[ctx.SP+1].I32
			ctx.Frame[ctx.SP].I32 = vm.hostGetKeyId(keyRef, size)
			return nil
		})
	_ = lnk.DefineFunction("hostGetObjectId",
		[]value.DataType{value.I32, value.I32, value.I32},
		[]value.DataType{value.I32},
		func(ctx *sections.HostContext) error {
			objId := ctx.Frame[ctx.SP].I32
			keyId := ctx.Frame[ctx.SP+1].I32
			typeId := ctx.Frame[ctx.SP+2].I32
			ctx.Frame[ctx.SP].I32 = vm.hostGetObjectId(objId, keyId, typeId)
			return nil
		})
	_ = lnk.DefineFunction("hostSetBytes",
		[]value.DataType{value.I32, value.I32, value.I32, value.I32},
		nil,
		func(ctx *sections.HostContext) error {
			objId := ctx.Frame[ctx.SP].I32
			keyId := ctx.Frame[ctx.SP+1].I32
			stringRef := ctx.Frame[ctx.SP+2].I32
			size := ctx.Frame[ctx.SP+3].I32
			vm.hostSetBytes(objId, keyId, stringRef, size)
			return nil
		})
	_ = lnk.DefineFunction("hostSetInt",
		[]value.DataType{value.I32, value.I32, value.I64},
		nil,
		func(ctx *sections.HostContext) error {
			objId := ctx.Frame[ctx.SP].I32
			keyId := ctx.Frame[ctx.SP+1].I32
			value := ctx.Frame[ctx.SP+2].I64
			vm.hostSetInt(objId, keyId, value)
			return nil
		})
	_ = lnk.DefineFunction("hostSetIntRef",
		[]value.DataType{value.I32, value.I32, value.I32},
		nil,
		func(ctx *sections.HostContext) error {
			objId := ctx.Frame[ctx.SP].I32
			keyId := ctx.Frame[ctx.SP+1].I32
			intRef := ctx.Frame[ctx.SP+2].I32
			vm.hostSetIntRef(objId, keyId, intRef)
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
			ctx.Frame[ctx.SP].I32 = vm.hostFdWrite(fd, iovs, size, written)
			return nil
		})
	return nil
}

func (vm *WartVM) LoadWasm(wasmData []byte) error {
	return vm.runner.Load(wasmData)
}

func (vm *WartVM) RunFunction(functionName string) error {
	err := vm.runner.RunExport(functionName, nil)
	return err
}

func (vm *WartVM) RunScFunction(index int32) error {
	params := make([]sections.Variable, 1)
	params[0].I32 = index
	frame := vm.preCall()
	err := vm.runner.RunExport("sc_call_entrypoint", params)
	vm.postCall(frame)
	return err
}

func (vm *WartVM) UnsafeMemory() []byte {
	return vm.runner.Memory()
}
