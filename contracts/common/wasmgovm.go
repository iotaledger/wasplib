// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package common

import (
	"errors"
	"strings"

	"github.com/iotaledger/wasp/packages/vm/wasmhost"
	"github.com/iotaledger/wasplib/packages/vm/wasmlib"
)

type WasmGoVM struct {
	wasmhost.WasmVMBase
	contract string
	onLoad   map[string]func()
}

var _ wasmhost.WasmVM = &WasmGoVM{}

func NewWasmGoVM(onLoad map[string]func()) *WasmGoVM {
	vm := &WasmGoVM{}
	vm.onLoad = onLoad
	return vm
}

func (vm *WasmGoVM) Interrupt() {
	// disabled for now
	// panic("implement me")
}

func (vm *WasmGoVM) LinkHost(impl wasmhost.WasmVM, host *wasmhost.WasmHost) error {
	_ = vm.WasmVMBase.LinkHost(impl, host)
	wasmlib.ConnectHost(host)
	return nil
}

func (vm *WasmGoVM) LoadWasm(wasmData []byte) error {
	contract := string(wasmData)
	if !strings.HasPrefix(contract, "go:") {
		return errors.New("WasmGoVM: not a Go contract: " + contract)
	}
	vm.contract = contract[3:]
	onLoad, ok := vm.onLoad[vm.contract]
	if !ok {
		return errors.New("WasmGoVM: unknown contract: " + vm.contract)
	}
	onLoad()
	return nil
}

func (vm *WasmGoVM) RunFunction(functionName string, args ...interface{}) error {
	// already ran on_load in LoadWasm, other functions are not supported
	if functionName != "on_load" {
		return errors.New("WasmGoVM: cannot run function: " + functionName)
	}
	return nil
}

func (vm *WasmGoVM) RunScFunction(index int32) error {
	return vm.Run(func() error {
		wasmlib.OnCall(index)
		return nil
	})
}

func (vm *WasmGoVM) UnsafeMemory() []byte {
	// no need to communicate through Wasm mem pool
	return nil
}
