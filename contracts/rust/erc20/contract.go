// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package erc20

import "github.com/iotaledger/wasplib/packages/vm/wasmlib"

type ApproveCall struct {
	Func *wasmlib.ScFunc
	Params MutableApproveParams
}

func NewApproveCall(ctx wasmlib.ScFuncContext) *ApproveCall {
	f := &ApproveCall{Func: wasmlib.NewScFunc(HScName, HFuncApprove)}
	f.Func.SetPtrs(&f.Params.id, nil)
	return f
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

type TransferCall struct {
	Func *wasmlib.ScFunc
	Params MutableTransferParams
}

func NewTransferCall(ctx wasmlib.ScFuncContext) *TransferCall {
	f := &TransferCall{Func: wasmlib.NewScFunc(HScName, HFuncTransfer)}
	f.Func.SetPtrs(&f.Params.id, nil)
	return f
}

type TransferFromCall struct {
	Func *wasmlib.ScFunc
	Params MutableTransferFromParams
}

func NewTransferFromCall(ctx wasmlib.ScFuncContext) *TransferFromCall {
	f := &TransferFromCall{Func: wasmlib.NewScFunc(HScName, HFuncTransferFrom)}
	f.Func.SetPtrs(&f.Params.id, nil)
	return f
}

type AllowanceCall struct {
	Func *wasmlib.ScView
	Params MutableAllowanceParams
	Results ImmutableAllowanceResults
}

func NewAllowanceCall(ctx wasmlib.ScFuncContext) *AllowanceCall {
	f := &AllowanceCall{Func: wasmlib.NewScView(HScName, HViewAllowance)}
	f.Func.SetPtrs(&f.Params.id, &f.Results.id)
	return f
}

func NewAllowanceCallFromView(ctx wasmlib.ScViewContext) *AllowanceCall {
	return NewAllowanceCall(wasmlib.ScFuncContext{})
}

type BalanceOfCall struct {
	Func *wasmlib.ScView
	Params MutableBalanceOfParams
	Results ImmutableBalanceOfResults
}

func NewBalanceOfCall(ctx wasmlib.ScFuncContext) *BalanceOfCall {
	f := &BalanceOfCall{Func: wasmlib.NewScView(HScName, HViewBalanceOf)}
	f.Func.SetPtrs(&f.Params.id, &f.Results.id)
	return f
}

func NewBalanceOfCallFromView(ctx wasmlib.ScViewContext) *BalanceOfCall {
	return NewBalanceOfCall(wasmlib.ScFuncContext{})
}

type TotalSupplyCall struct {
	Func *wasmlib.ScView
	Results ImmutableTotalSupplyResults
}

func NewTotalSupplyCall(ctx wasmlib.ScFuncContext) *TotalSupplyCall {
	f := &TotalSupplyCall{Func: wasmlib.NewScView(HScName, HViewTotalSupply)}
	f.Func.SetPtrs(nil, &f.Results.id)
	return f
}

func NewTotalSupplyCallFromView(ctx wasmlib.ScViewContext) *TotalSupplyCall {
	return NewTotalSupplyCall(wasmlib.ScFuncContext{})
}
