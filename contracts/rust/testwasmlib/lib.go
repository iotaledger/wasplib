// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package testwasmlib

import "github.com/iotaledger/wasplib/packages/vm/wasmlib"

func OnLoad() {
	exports := wasmlib.NewScExports()
	exports.AddFunc(FuncParamTypes, funcParamTypesThunk)
	exports.AddView(ViewBlockRecord, viewBlockRecordThunk)
	exports.AddView(ViewBlockRecords, viewBlockRecordsThunk)

	for i, key := range keyMap {
		idxMap[i] = key.KeyID()
	}
}

type ParamTypesContext struct {
	Params ImmutableParamTypesParams
	State  MutableTestWasmLibState
}

func funcParamTypesThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("testwasmlib.funcParamTypes")
	f := &ParamTypesContext{
		Params: ImmutableParamTypesParams{
			id: wasmlib.OBJ_ID_PARAMS,
		},
		State: MutableTestWasmLibState{
			id: wasmlib.OBJ_ID_STATE,
		},
	}
	funcParamTypes(ctx, f)
	ctx.Log("testwasmlib.funcParamTypes ok")
}

type BlockRecordContext struct {
	Params  ImmutableBlockRecordParams
	Results MutableBlockRecordResults
	State   ImmutableTestWasmLibState
}

func viewBlockRecordThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("testwasmlib.viewBlockRecord")
	f := &BlockRecordContext{
		Params: ImmutableBlockRecordParams{
			id: wasmlib.OBJ_ID_PARAMS,
		},
		Results: MutableBlockRecordResults{
			id: wasmlib.OBJ_ID_RESULTS,
		},
		State: ImmutableTestWasmLibState{
			id: wasmlib.OBJ_ID_STATE,
		},
	}
	ctx.Require(f.Params.BlockIndex().Exists(), "missing mandatory blockIndex")
	ctx.Require(f.Params.RecordIndex().Exists(), "missing mandatory recordIndex")
	viewBlockRecord(ctx, f)
	ctx.Log("testwasmlib.viewBlockRecord ok")
}

type BlockRecordsContext struct {
	Params  ImmutableBlockRecordsParams
	Results MutableBlockRecordsResults
	State   ImmutableTestWasmLibState
}

func viewBlockRecordsThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("testwasmlib.viewBlockRecords")
	f := &BlockRecordsContext{
		Params: ImmutableBlockRecordsParams{
			id: wasmlib.OBJ_ID_PARAMS,
		},
		Results: MutableBlockRecordsResults{
			id: wasmlib.OBJ_ID_RESULTS,
		},
		State: ImmutableTestWasmLibState{
			id: wasmlib.OBJ_ID_STATE,
		},
	}
	ctx.Require(f.Params.BlockIndex().Exists(), "missing mandatory blockIndex")
	viewBlockRecords(ctx, f)
	ctx.Log("testwasmlib.viewBlockRecords ok")
}
