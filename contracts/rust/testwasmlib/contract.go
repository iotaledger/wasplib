// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package testwasmlib

import "github.com/iotaledger/wasplib/packages/vm/wasmlib"

type ParamTypesCall struct {
	Func   *wasmlib.ScFunc
	Params MutableParamTypesParams
}

type BlockRecordCall struct {
	Func    *wasmlib.ScView
	Params  MutableBlockRecordParams
	Results ImmutableBlockRecordResults
}

type BlockRecordsCall struct {
	Func    *wasmlib.ScView
	Params  MutableBlockRecordsParams
	Results ImmutableBlockRecordsResults
}

type Funcs struct{}

var ScFuncs Funcs

func (sc Funcs) ParamTypes(ctx wasmlib.ScFuncCallContext) *ParamTypesCall {
	f := &ParamTypesCall{Func: wasmlib.NewScFunc(HScName, HFuncParamTypes)}
	f.Func.SetPtrs(&f.Params.id, nil)
	return f
}

func (sc Funcs) BlockRecord(ctx wasmlib.ScViewCallContext) *BlockRecordCall {
	f := &BlockRecordCall{Func: wasmlib.NewScView(HScName, HViewBlockRecord)}
	f.Func.SetPtrs(&f.Params.id, &f.Results.id)
	return f
}

func (sc Funcs) BlockRecords(ctx wasmlib.ScViewCallContext) *BlockRecordsCall {
	f := &BlockRecordsCall{Func: wasmlib.NewScView(HScName, HViewBlockRecords)}
	f.Func.SetPtrs(&f.Params.id, &f.Results.id)
	return f
}
