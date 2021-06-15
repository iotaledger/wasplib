// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package testwasmlib

import "github.com/iotaledger/wasplib/packages/vm/wasmlib"

type TestWasmLibFunc struct {
	sc wasmlib.ScContractFunc
}

func NewTestWasmLibFunc(ctx wasmlib.ScFuncContext) *TestWasmLibFunc {
	return &TestWasmLibFunc{sc: wasmlib.NewScContractFunc(ctx, HScName)}
}

func (f *TestWasmLibFunc) Delay(seconds int32) *TestWasmLibFunc {
	f.sc.Delay(seconds)
	return f
}

func (f *TestWasmLibFunc) OfContract(contract wasmlib.ScHname) *TestWasmLibFunc {
	f.sc.OfContract(contract)
	return f
}

func (f *TestWasmLibFunc) Post() *TestWasmLibFunc {
	f.sc.Post()
	return f
}

func (f *TestWasmLibFunc) PostToChain(chainId wasmlib.ScChainId) *TestWasmLibFunc {
	f.sc.PostToChain(chainId)
	return f
}

func (f *TestWasmLibFunc) ParamTypes(params MutableFuncParamTypesParams, transfer wasmlib.ScTransfers) {
	f.sc.Run(HFuncParamTypes, params.id, &transfer)
}

type TestWasmLibView struct {
	sc wasmlib.ScContractView
}

func NewTestWasmLibView(ctx wasmlib.ScViewContext) *TestWasmLibView {
	return &TestWasmLibView{sc: wasmlib.NewScContractView(ctx, HScName)}
}

func (v *TestWasmLibView) OfContract(contract wasmlib.ScHname) *TestWasmLibView {
	v.sc.OfContract(contract)
	return v
}
