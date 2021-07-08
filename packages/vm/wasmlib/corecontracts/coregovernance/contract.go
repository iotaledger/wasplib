// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package coregovernance

import "github.com/iotaledger/wasplib/packages/vm/wasmlib"

type AddAllowedStateControllerAddressCall struct {
	Func   *wasmlib.ScFunc
	Params MutableAddAllowedStateControllerAddressParams
}

type RemoveAllowedStateControllerAddressCall struct {
	Func   *wasmlib.ScFunc
	Params MutableRemoveAllowedStateControllerAddressParams
}

type RotateStateControllerCall struct {
	Func   *wasmlib.ScFunc
	Params MutableRotateStateControllerParams
}

type GetAllowedStateControllerAddressesCall struct {
	Func    *wasmlib.ScView
	Results ImmutableGetAllowedStateControllerAddressesResults
}

type Funcs struct{}

var ScFuncs Funcs

func (sc Funcs) AddAllowedStateControllerAddress(ctx wasmlib.ScFuncCallContext) *AddAllowedStateControllerAddressCall {
	f := &AddAllowedStateControllerAddressCall{Func: wasmlib.NewScFunc(HScName, HFuncAddAllowedStateControllerAddress)}
	f.Func.SetPtrs(&f.Params.id, nil)
	return f
}

func (sc Funcs) RemoveAllowedStateControllerAddress(ctx wasmlib.ScFuncCallContext) *RemoveAllowedStateControllerAddressCall {
	f := &RemoveAllowedStateControllerAddressCall{Func: wasmlib.NewScFunc(HScName, HFuncRemoveAllowedStateControllerAddress)}
	f.Func.SetPtrs(&f.Params.id, nil)
	return f
}

func (sc Funcs) RotateStateController(ctx wasmlib.ScFuncCallContext) *RotateStateControllerCall {
	f := &RotateStateControllerCall{Func: wasmlib.NewScFunc(HScName, HFuncRotateStateController)}
	f.Func.SetPtrs(&f.Params.id, nil)
	return f
}

func (sc Funcs) GetAllowedStateControllerAddresses(ctx wasmlib.ScViewCallContext) *GetAllowedStateControllerAddressesCall {
	f := &GetAllowedStateControllerAddressesCall{Func: wasmlib.NewScView(HScName, HViewGetAllowedStateControllerAddresses)}
	f.Func.SetPtrs(nil, &f.Results.id)
	return f
}
