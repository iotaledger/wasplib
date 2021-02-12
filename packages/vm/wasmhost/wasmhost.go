// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package wasmhost

import (
	"errors"
	"github.com/iotaledger/hive.go/logger"
	"github.com/iotaledger/wasp/packages/coretypes"
)

type WasmHost struct {
	KvStoreHost
	vm          WasmVM
	codeToFunc  map[uint32]string
	funcToCode  map[string]uint32
	funcToIndex map[string]int32
}

func (host *WasmHost) InitVM(vm WasmVM, useBase58Keys bool) error {
	host.useBase58Keys = useBase58Keys
	return vm.LinkHost(vm, host)
}

func (host *WasmHost) Init(null HostObject, root HostObject, log *logger.Logger) {
	host.KvStoreHost.Init(null, root, log)
	host.codeToFunc = make(map[uint32]string)
	host.funcToCode = make(map[string]uint32)
	host.funcToIndex = make(map[string]int32)
}

func (host *WasmHost) FunctionFromCode(code uint32) string {
	return host.codeToFunc[code]
}

func (host *WasmHost) IsView(function string) bool {
	return (host.funcToIndex[function] & 0x8000) != 0
}

func (host *WasmHost) LoadWasm(wasmData []byte) error {
	err := host.vm.LoadWasm(wasmData)
	if err != nil {
		return err
	}
	err = host.RunFunction("on_load")
	if err != nil {
		return err
	}
	host.vm.SaveMemory()
	return nil
}

func (host *WasmHost) RunFunction(functionName string) (err error) {
	return host.vm.RunFunction(functionName)
}

func (host *WasmHost) RunScFunction(functionName string) (err error) {
	index, ok := host.funcToIndex[functionName]
	if !ok {
		return errors.New("unknown SC function name: " + functionName)
	}
	return host.vm.RunScFunction(index)
}

func (host *WasmHost) SetExport(index int32, functionName string) {
	if index < 0 {
		// double check that predefined keys are in sync
		if index == KeyZzzzzzz {
			return
		}
		panic("SetExport: predefined key value mismatch")
	}
	_, ok := host.funcToCode[functionName]
	if ok {
		panic("SetExport: duplicate function name")
	}
	hn := coretypes.Hn(functionName)
	hashedName := uint32(hn)
	_, ok = host.codeToFunc[hashedName]
	if ok {
		panic("SetExport: duplicate hashed name")
	}
	host.codeToFunc[hashedName] = functionName
	host.funcToCode[functionName] = hashedName
	host.funcToIndex[functionName] = index
}
