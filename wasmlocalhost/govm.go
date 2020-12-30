// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package wasmlocalhost

import (
	"errors"
	"github.com/iotaledger/wasp/packages/vm/wasmhost"
	"github.com/iotaledger/wasplib/client"
	"strings"
)

type GoVM struct {
	wasmhost.WasmVmBase
	contract string
	onLoad   map[string]func()
}

func NewGoVM(onLoad map[string]func()) *GoVM {
	vm := &GoVM{}
	vm.onLoad = onLoad
	return vm
}

func (vm *GoVM) LinkHost(impl wasmhost.WasmVM, host *wasmhost.WasmHost) error {
	vm.WasmVmBase.LinkHost(impl, host)
	client.ConnectHost(host)
	return nil
}

func (vm *GoVM) LoadWasm(wasmData []byte) error {
	contract := string(wasmData)
	if !strings.HasPrefix(contract, "go:") {
		return errors.New("GoVM: not a Go contract: " + contract)
	}
	vm.contract = contract[3:]
	onLoad, ok := vm.onLoad[vm.contract]
	if !ok {
		return errors.New("Unknown contract: " + vm.contract)
	}
	onLoad()
	return nil
}

func (vm *GoVM) RunFunction(functionName string) error {
	// already ran on_load in LoadWasm, other functions are not supported
	if functionName != "on_load" {
		return errors.New("GoVM: cannot run function: " + functionName)
	}
	return nil
}

func (vm *GoVM) RunScFunction(index int32) error {
	client.ScCallEntrypoint(index)
	return nil
}

func (vm *GoVM) UnsafeMemory() []byte {
	// no need to communicate through Wasm mem pool
	return nil
}
