// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package common

import (
	"fmt"

	"github.com/iotaledger/wart/wasm/consts/value"
	"github.com/iotaledger/wart/wasm/executors"
	"github.com/iotaledger/wart/wasm/sections"
	"github.com/iotaledger/wasp/packages/vm/wasmhost"
)

type WartVM struct {
	wasmhost.WasmVmBase
	runner *executors.WasmRunner
}

var _ wasmhost.WasmVM = &WartVM{}

func NewWartVM() *WartVM {
	vm := &WartVM{}
	vm.runner = executors.NewWasmRunner()
	return vm
}

func (vm *WartVM) Interrupt() {
	vm.runner.Interrupt()
}

func (vm *WartVM) LinkHost(impl wasmhost.WasmVM, host *wasmhost.WasmHost) error {
	_ = vm.WasmVmBase.LinkHost(impl, host)

	m := executors.DefineModule("WasmLib")
	lnk := executors.NewWasmLinker(m)
	_ = lnk.DefineFunction("hostGetBytes",
		[]value.DataType{value.I32, value.I32, value.I32, value.I32, value.I32},
		[]value.DataType{value.I32},
		func(ctx *sections.HostContext) error {
			args := ctx.Frame[ctx.SP:]
			objId := args[0].I32
			keyId := args[1].I32
			typeId := args[2].I32
			stringRef := args[3].I32
			size := args[4].I32
			args[0].I32 = vm.HostGetBytes(objId, keyId, typeId, stringRef, size)
			return nil
		})
	_ = lnk.DefineFunction("hostGetKeyId",
		[]value.DataType{value.I32, value.I32},
		[]value.DataType{value.I32},
		func(ctx *sections.HostContext) error {
			args := ctx.Frame[ctx.SP:]
			keyRef := args[0].I32
			size := args[1].I32
			args[0].I32 = vm.HostGetKeyId(keyRef, size)
			return nil
		})
	_ = lnk.DefineFunction("hostGetObjectId",
		[]value.DataType{value.I32, value.I32, value.I32},
		[]value.DataType{value.I32},
		func(ctx *sections.HostContext) error {
			args := ctx.Frame[ctx.SP:]
			objId := args[0].I32
			keyId := args[1].I32
			typeId := args[2].I32
			args[0].I32 = vm.HostGetObjectId(objId, keyId, typeId)
			return nil
		})
	_ = lnk.DefineFunction("hostSetBytes",
		[]value.DataType{value.I32, value.I32, value.I32, value.I32, value.I32},
		nil,
		func(ctx *sections.HostContext) error {
			args := ctx.Frame[ctx.SP:]
			objId := args[0].I32
			keyId := args[1].I32
			typeId := args[2].I32
			stringRef := args[3].I32
			size := args[4].I32
			vm.HostSetBytes(objId, keyId, typeId, stringRef, size)
			return nil
		})
	// go implementation uses this one to write panic message
	m = executors.DefineModule("wasi_snapshot_preview1")
	lnk = executors.NewWasmLinker(m)
	_ = lnk.DefineFunction("fd_write",
		[]value.DataType{value.I32, value.I32, value.I32, value.I32},
		[]value.DataType{value.I32},
		func(ctx *sections.HostContext) error {
			args := ctx.Frame[ctx.SP:]
			fd := args[0].I32
			iovs := args[1].I32
			size := args[2].I32
			written := args[3].I32
			args[0].I32 = vm.HostFdWrite(fd, iovs, size, written)
			return nil
		})
	return nil
}

func (vm *WartVM) LoadWasm(wasmData []byte) error {
	return vm.runner.Load(wasmData)
}

func (vm *WartVM) RunFunction(functionName string, args ...interface{}) error {
	if len(args) != 0 {
		panic("RunFunction.args not implemented for Wart")
	}
	return vm.Run(func() error {
		err := vm.runner.RunExport(functionName, nil)
		fmt.Printf("%s gas used: %d\n", functionName, vm.runner.GasUsed())
		return err
	})
}

func (vm *WartVM) RunScFunction(index int32) error {
	params := make([]sections.Variable, 1)
	params[0].I32 = index
	frame := vm.PreCall()
	err := vm.Run(func() error {
		err := vm.runner.RunExport("on_call", params)
		fmt.Printf("on_call(%d) gas used: %d\n", index, vm.runner.GasUsed())
		return err
	})
	vm.PostCall(frame)
	return err
}

func (vm *WartVM) UnsafeMemory() []byte {
	return vm.runner.Memory()
}
