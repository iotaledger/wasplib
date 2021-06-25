// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package testwasmlib

import "github.com/iotaledger/wasplib/packages/vm/wasmlib"

type ParamTypesCall struct {
	Func *wasmlib.ScFunc
	Params MutableParamTypesParams
}

func NewParamTypesCall(ctx wasmlib.ScFuncContext) *ParamTypesCall {
	f := &ParamTypesCall{Func: wasmlib.NewScFunc(HScName, HFuncParamTypes)}
	f.Func.SetPtrs(&f.Params.id, nil)
	return f
}

type BlockRecordCall struct {
	Func *wasmlib.ScView
	Params MutableBlockRecordParams
	Results ImmutableBlockRecordResults
}

func NewBlockRecordCall(ctx wasmlib.ScFuncContext) *BlockRecordCall {
	f := &BlockRecordCall{Func: wasmlib.NewScView(HScName, HViewBlockRecord)}
	f.Func.SetPtrs(&f.Params.id, &f.Results.id)
	return f
}

func NewBlockRecordCallFromView(ctx wasmlib.ScViewContext) *BlockRecordCall {
	return NewBlockRecordCall(wasmlib.ScFuncContext{})
}

type BlockRecordsCall struct {
	Func *wasmlib.ScView
	Params MutableBlockRecordsParams
	Results ImmutableBlockRecordsResults
}

func NewBlockRecordsCall(ctx wasmlib.ScFuncContext) *BlockRecordsCall {
	f := &BlockRecordsCall{Func: wasmlib.NewScView(HScName, HViewBlockRecords)}
	f.Func.SetPtrs(&f.Params.id, &f.Results.id)
	return f
}

func NewBlockRecordsCallFromView(ctx wasmlib.ScViewContext) *BlockRecordsCall {
	return NewBlockRecordsCall(wasmlib.ScFuncContext{})
}
