// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package coreeventlog

import "github.com/iotaledger/wasplib/packages/vm/wasmlib"

type GetNumRecordsCall struct {
	Func    *wasmlib.ScView
	Params  MutableGetNumRecordsParams
	Results ImmutableGetNumRecordsResults
}

type GetRecordsCall struct {
	Func    *wasmlib.ScView
	Params  MutableGetRecordsParams
	Results ImmutableGetRecordsResults
}

type coreeventlogFuncs struct{}

var ScFuncs coreeventlogFuncs

func (sc coreeventlogFuncs) GetNumRecords(ctx wasmlib.ScViewCallContext) *GetNumRecordsCall {
	f := &GetNumRecordsCall{Func: wasmlib.NewScView(HScName, HViewGetNumRecords)}
	f.Func.SetPtrs(&f.Params.id, &f.Results.id)
	return f
}

func (sc coreeventlogFuncs) GetRecords(ctx wasmlib.ScViewCallContext) *GetRecordsCall {
	f := &GetRecordsCall{Func: wasmlib.NewScView(HScName, HViewGetRecords)}
	f.Func.SetPtrs(&f.Params.id, &f.Results.id)
	return f
}
