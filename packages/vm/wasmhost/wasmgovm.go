// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package wasmhost

import (
	"errors"
	"github.com/iotaledger/wasp/packages/vm/wasmlib"
	"strings"
)

type WasmGoVM struct {
	WasmVmBase
	contract string
	onLoad   map[string]func()
}

func NewWasmGoVM(onLoad map[string]func()) *WasmGoVM {
	vm := &WasmGoVM{}
	vm.onLoad = onLoad
	return vm
}

func (vm *WasmGoVM) LinkHost(impl WasmVM, host *WasmHost) error {
	vm.WasmVmBase.LinkHost(impl, host)
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

func (vm *WasmGoVM) RunFunction(functionName string) error {
	// already ran on_load in LoadWasm, other functions are not supported
	if functionName != "on_load" {
		return errors.New("WasmGoVM: cannot run function: " + functionName)
	}
	return nil
}

func (vm *WasmGoVM) RunScFunction(index int32) error {
	wasmlib.ScCallEntrypoint(index)
	return nil
}

func (vm *WasmGoVM) UnsafeMemory() []byte {
	// no need to communicate through Wasm mem pool
	return nil
}
