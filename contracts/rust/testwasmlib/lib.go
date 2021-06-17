// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package testwasmlib

import "github.com/iotaledger/wasplib/packages/vm/wasmlib"

func OnLoad() {
	exports := wasmlib.NewScExports()
	exports.AddFunc(FuncParamTypes, funcParamTypesThunk)
	exports.AddView(ViewBlockRecord, viewBlockRecordThunk)
	exports.AddView(ViewBlockRecords, viewBlockRecordsThunk)

	for i, key := range keyMap {
		idxMap[i] = key.KeyId()
	}
}

type FuncParamTypesContext struct {
	Params ImmutableFuncParamTypesParams
	State  MutableTestWasmLibState
}

func funcParamTypesThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("testwasmlib.funcParamTypes")
	f := &FuncParamTypesContext{
		Params: ImmutableFuncParamTypesParams{
			id: wasmlib.GetObjectId(1, wasmlib.KeyParams, wasmlib.TYPE_MAP),
		},
		State: MutableTestWasmLibState{
			id: wasmlib.GetObjectId(1, wasmlib.KeyState, wasmlib.TYPE_MAP),
		},
	}
	funcParamTypes(ctx, f)
	ctx.Log("testwasmlib.funcParamTypes ok")
}

type ViewBlockRecordContext struct {
	Params  ImmutableViewBlockRecordParams
	Results MutableViewBlockRecordResults
	State   ImmutableTestWasmLibState
}

func viewBlockRecordThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("testwasmlib.viewBlockRecord")
	f := &ViewBlockRecordContext{
		Params: ImmutableViewBlockRecordParams{
			id: wasmlib.GetObjectId(1, wasmlib.KeyParams, wasmlib.TYPE_MAP),
		},
		Results: MutableViewBlockRecordResults{
			id: wasmlib.GetObjectId(1, wasmlib.KeyResults, wasmlib.TYPE_MAP),
		},
		State: ImmutableTestWasmLibState{
			id: wasmlib.GetObjectId(1, wasmlib.KeyState, wasmlib.TYPE_MAP),
		},
	}
	ctx.Require(f.Params.BlockIndex().Exists(), "missing mandatory blockIndex")
	ctx.Require(f.Params.RecordIndex().Exists(), "missing mandatory recordIndex")
	viewBlockRecord(ctx, f)
	ctx.Log("testwasmlib.viewBlockRecord ok")
}

type ViewBlockRecordsContext struct {
	Params  ImmutableViewBlockRecordsParams
	Results MutableViewBlockRecordsResults
	State   ImmutableTestWasmLibState
}

func viewBlockRecordsThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("testwasmlib.viewBlockRecords")
	f := &ViewBlockRecordsContext{
		Params: ImmutableViewBlockRecordsParams{
			id: wasmlib.GetObjectId(1, wasmlib.KeyParams, wasmlib.TYPE_MAP),
		},
		Results: MutableViewBlockRecordsResults{
			id: wasmlib.GetObjectId(1, wasmlib.KeyResults, wasmlib.TYPE_MAP),
		},
		State: ImmutableTestWasmLibState{
			id: wasmlib.GetObjectId(1, wasmlib.KeyState, wasmlib.TYPE_MAP),
		},
	}
	ctx.Require(f.Params.BlockIndex().Exists(), "missing mandatory blockIndex")
	viewBlockRecords(ctx, f)
	ctx.Log("testwasmlib.viewBlockRecords ok")
}
