// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package common

// import (
// 	"github.com/iotaledger/wasp/packages/vm/wasmhost"
// 	"github.com/wasmerio/wasmer-go/wasmer"
// )
//
// type WasmerVM struct {
// 	wasmhost.WasmVMBase
// 	instance *wasmer.Instance
// 	linker   *wasmer.ImportObject
// 	memory   *wasmer.Memory
// 	module   *wasmer.Module
// 	store    *wasmer.Store
// }
//
// var _ wasmhost.WasmVM = &WasmerVM{}
//
// func NewWasmerVM() *WasmerVM {
// 	vm := &WasmerVM{}
// 	vm.store = wasmer.NewStore(wasmer.NewEngine())
// 	vm.linker = wasmer.NewImportObject()
// 	return vm
// }
//
// func (vm *WasmerVM) Interrupt() {
// 	panic("implement me")
// }
//
// func (vm *WasmerVM) LinkHost(impl wasmhost.WasmVM, host *wasmhost.WasmHost) error {
// 	_ = vm.WasmVMBase.LinkHost(impl, host)
//
// 	typeVoid := wasmer.NewValueTypes()
// 	typeInt32 := wasmer.NewValueTypes(wasmer.I32)
// 	type2Int32 := wasmer.NewValueTypes(wasmer.I32, wasmer.I32)
// 	type3Int32 := wasmer.NewValueTypes(wasmer.I32, wasmer.I32, wasmer.I32)
// 	type4Int32 := wasmer.NewValueTypes(wasmer.I32, wasmer.I32, wasmer.I32, wasmer.I32)
// 	type5Int32 := wasmer.NewValueTypes(wasmer.I32, wasmer.I32, wasmer.I32, wasmer.I32, wasmer.I32)
//
// 	hostGetBytes := func(args []wasmer.Value) ([]wasmer.Value, error) { return vm.exportHostGetBytes(args) }
// 	hostGetKeyID := func(args []wasmer.Value) ([]wasmer.Value, error) { return vm.exportHostGetKeyID(args) }
// 	hostGetObjectID := func(args []wasmer.Value) ([]wasmer.Value, error) { return vm.exportHostGetObjectID(args) }
// 	hostSetBytes := func(args []wasmer.Value) ([]wasmer.Value, error) { return vm.exportHostSetBytes(args) }
//
// 	funcs := map[string]wasmer.IntoExtern{
// 		"hostGetBytes":    wasmer.NewFunction(vm.store, wasmer.NewFunctionType(type5Int32, typeInt32), hostGetBytes).IntoExtern(),
// 		"hostGetKeyID":    wasmer.NewFunction(vm.store, wasmer.NewFunctionType(type2Int32, typeInt32), hostGetKeyID).IntoExtern(),
// 		"hostGetObjectID": wasmer.NewFunction(vm.store, wasmer.NewFunctionType(type3Int32, typeInt32), hostGetObjectID).IntoExtern(),
// 		"hostSetBytes":    wasmer.NewFunction(vm.store, wasmer.NewFunctionType(type5Int32, typeVoid), hostSetBytes).IntoExtern(),
// 	}
// 	vm.linker.Register("WasmLib", funcs)
//
// 	// TinyGo Wasm implementation uses this one to write panic message to console
// 	fdWrite := func(args []wasmer.Value) ([]wasmer.Value, error) { return vm.exportFdWrite(args) }
// 	funcs = map[string]wasmer.IntoExtern{
// 		"fd_write": wasmer.NewFunction(vm.store, wasmer.NewFunctionType(type4Int32, typeInt32), fdWrite).IntoExtern(),
// 	}
// 	vm.linker.Register("wasi_unstable", funcs)
// 	return nil
// }
//
// func (vm *WasmerVM) LoadWasm(wasmData []byte) error {
// 	var err error
// 	vm.module, err = wasmer.NewModule(vm.store, wasmData)
// 	if err != nil {
// 		return err
// 	}
// 	vm.instance, err = wasmer.NewInstance(vm.module, vm.linker)
// 	if err != nil {
// 		return err
// 	}
// 	vm.memory, err = vm.instance.Exports.GetMemory("memory")
// 	return err
// }
//
// func (vm *WasmerVM) RunFunction(functionName string, args ...interface{}) error {
// 	export, err := vm.instance.Exports.GetFunction(functionName)
// 	if err != nil {
// 		return err
// 	}
// 	return vm.Run(func() error {
// 		_, err = export(args...)
// 		return err
// 	})
// }
//
// func (vm *WasmerVM) RunScFunction(index int32) error {
// 	export, err := vm.instance.Exports.GetFunction("on_call")
// 	if err != nil {
// 		return err
// 	}
// 	frame := vm.PreCall()
// 	err = vm.Run(func() error {
// 		_, err = export(index)
// 		return err
// 	})
// 	vm.PostCall(frame)
// 	return err
// }
//
// func (vm *WasmerVM) UnsafeMemory() []byte {
// 	return vm.memory.Data()
// }
//
// func (vm *WasmerVM) exportFdWrite(args []wasmer.Value) ([]wasmer.Value, error) {
// 	fd := args[0].I32()
// 	iovs := args[1].I32()
// 	size := args[2].I32()
// 	written := args[3].I32()
// 	ret := vm.HostFdWrite(fd, iovs, size, written)
// 	return []wasmer.Value{wasmer.NewI32(ret)}, nil
// }
//
// func (vm *WasmerVM) exportHostGetBytes(args []wasmer.Value) ([]wasmer.Value, error) {
// 	objID := args[0].I32()
// 	keyID := args[1].I32()
// 	typeID := args[2].I32()
// 	stringRef := args[3].I32()
// 	size := args[4].I32()
// 	ret := vm.HostGetBytes(objID, keyID, typeID, stringRef, size)
// 	return []wasmer.Value{wasmer.NewI32(ret)}, nil
// }
//
// func (vm *WasmerVM) exportHostGetKeyID(args []wasmer.Value) ([]wasmer.Value, error) {
// 	keyRef := args[0].I32()
// 	size := args[1].I32()
// 	ret := vm.HostGetKeyID(keyRef, size)
// 	return []wasmer.Value{wasmer.NewI32(ret)}, nil
// }
//
// func (vm *WasmerVM) exportHostGetObjectID(args []wasmer.Value) ([]wasmer.Value, error) {
// 	objID := args[0].I32()
// 	keyID := args[1].I32()
// 	typeID := args[2].I32()
// 	ret := vm.HostGetObjectID(objID, keyID, typeID)
// 	return []wasmer.Value{wasmer.NewI32(ret)}, nil
// }
//
// func (vm *WasmerVM) exportHostSetBytes(args []wasmer.Value) ([]wasmer.Value, error) {
// 	objID := args[0].I32()
// 	keyID := args[1].I32()
// 	typeID := args[2].I32()
// 	stringRef := args[3].I32()
// 	size := args[4].I32()
// 	vm.HostSetBytes(objID, keyID, typeID, stringRef, size)
// 	return []wasmer.Value{}, nil
// }
