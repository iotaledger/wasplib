// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package dividend

import "github.com/iotaledger/wasplib/packages/vm/wasmlib"

type DivideCall struct {
	Func *wasmlib.ScFunc
}

func NewDivideCall(ctx wasmlib.ScFuncContext) *DivideCall {
	return &DivideCall{Func: wasmlib.NewScFunc(HScName, HFuncDivide)}
}

type InitCall struct {
	Func *wasmlib.ScFunc
	Params MutableInitParams
}

func NewInitCall(ctx wasmlib.ScFuncContext) *InitCall {
	f := &InitCall{Func: wasmlib.NewScFunc(HScName, HFuncInit)}
	f.Func.SetPtrs(&f.Params.id, nil)
	return f
}

type MemberCall struct {
	Func *wasmlib.ScFunc
	Params MutableMemberParams
}

func NewMemberCall(ctx wasmlib.ScFuncContext) *MemberCall {
	f := &MemberCall{Func: wasmlib.NewScFunc(HScName, HFuncMember)}
	f.Func.SetPtrs(&f.Params.id, nil)
	return f
}

type SetOwnerCall struct {
	Func *wasmlib.ScFunc
	Params MutableSetOwnerParams
}

func NewSetOwnerCall(ctx wasmlib.ScFuncContext) *SetOwnerCall {
	f := &SetOwnerCall{Func: wasmlib.NewScFunc(HScName, HFuncSetOwner)}
	f.Func.SetPtrs(&f.Params.id, nil)
	return f
}

type GetFactorCall struct {
	Func *wasmlib.ScView
	Params MutableGetFactorParams
	Results ImmutableGetFactorResults
}

func NewGetFactorCall(ctx wasmlib.ScFuncContext) *GetFactorCall {
	f := &GetFactorCall{Func: wasmlib.NewScView(HScName, HViewGetFactor)}
	f.Func.SetPtrs(&f.Params.id, &f.Results.id)
	return f
}

func NewGetFactorCallFromView(ctx wasmlib.ScViewContext) *GetFactorCall {
	return NewGetFactorCall(wasmlib.ScFuncContext{})
}
