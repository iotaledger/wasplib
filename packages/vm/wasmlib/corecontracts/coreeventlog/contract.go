// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package coreeventlog

import "github.com/iotaledger/wasplib/packages/vm/wasmlib"

type GetNumRecordsCall struct {
	Func wasmlib.ScView
	Params MutableGetNumRecordsParams
	Results ImmutableGetNumRecordsResults
}

func NewGetNumRecordsCall(ctx wasmlib.ScFuncContext) *GetNumRecordsCall {
	f := &GetNumRecordsCall{}
	f.Func.Init(HScName, HViewGetNumRecords, &f.Params.id, &f.Results.id)
	return f
}

func NewGetNumRecordsCallFromView(ctx wasmlib.ScViewContext) *GetNumRecordsCall {
	f := &GetNumRecordsCall{}
	f.Func.Init(HScName, HViewGetNumRecords, &f.Params.id, &f.Results.id)
	return f
}

type GetRecordsCall struct {
	Func wasmlib.ScView
	Params MutableGetRecordsParams
}

func NewGetRecordsCall(ctx wasmlib.ScFuncContext) *GetRecordsCall {
	f := &GetRecordsCall{}
	f.Func.Init(HScName, HViewGetRecords, &f.Params.id, nil)
	return f
}

func NewGetRecordsCallFromView(ctx wasmlib.ScViewContext) *GetRecordsCall {
	f := &GetRecordsCall{}
	f.Func.Init(HScName, HViewGetRecords, &f.Params.id, nil)
	return f
}
