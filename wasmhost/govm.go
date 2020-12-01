// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package wasmhost

import "github.com/iotaledger/wasplib/client"

type GoVM struct {
	WasmVmBase
}

func NewGoVM() *GoVM {
	host := &GoVM{}
	host.impl = host
	return host
}

func (vm *GoVM) LinkHost(host *WasmHost) error {
	vm.host = host
	client.ConnectHost(host)
	return nil
}

func (vm *GoVM) LoadWasm(wasmData []byte) error {
	panic("GoVM.LoadWasm")
}

func (vm *GoVM) RunFunction(functionName string) error {
	panic("GoVM.RunFunction")
}

func (vm *GoVM) RunScFunction(index int32) error {
	//TODO how to clear global var state?
	client.ScCallEntrypoint(index)
	return nil
}

func (vm *GoVM) UnsafeMemory() []byte {
	panic("GoVM.UnsafeMemory")
}
