// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package testcore

import "github.com/iotaledger/wasplib/packages/vm/wasmlib"

func OnLoad() {
	exports := wasmlib.NewScExports()
	exports.AddFunc(FuncCallOnChain, funcCallOnChainThunk)
	exports.AddFunc(FuncCheckContextFromFullEP, funcCheckContextFromFullEPThunk)
	exports.AddFunc(FuncDoNothing, funcDoNothingThunk)
	exports.AddFunc(FuncGetMintedSupply, funcGetMintedSupplyThunk)
	exports.AddFunc(FuncIncCounter, funcIncCounterThunk)
	exports.AddFunc(FuncInit, funcInitThunk)
	exports.AddFunc(FuncPassTypesFull, funcPassTypesFullThunk)
	exports.AddFunc(FuncRunRecursion, funcRunRecursionThunk)
	exports.AddFunc(FuncSendToAddress, funcSendToAddressThunk)
	exports.AddFunc(FuncSetInt, funcSetIntThunk)
	exports.AddFunc(FuncTestCallPanicFullEP, funcTestCallPanicFullEPThunk)
	exports.AddFunc(FuncTestCallPanicViewEPFromFull, funcTestCallPanicViewEPFromFullThunk)
	exports.AddFunc(FuncTestChainOwnerIDFull, funcTestChainOwnerIDFullThunk)
	exports.AddFunc(FuncTestEventLogDeploy, funcTestEventLogDeployThunk)
	exports.AddFunc(FuncTestEventLogEventData, funcTestEventLogEventDataThunk)
	exports.AddFunc(FuncTestEventLogGenericData, funcTestEventLogGenericDataThunk)
	exports.AddFunc(FuncTestPanicFullEP, funcTestPanicFullEPThunk)
	exports.AddFunc(FuncWithdrawToChain, funcWithdrawToChainThunk)
	exports.AddView(ViewCheckContextFromViewEP, viewCheckContextFromViewEPThunk)
	exports.AddView(ViewFibonacci, viewFibonacciThunk)
	exports.AddView(ViewGetCounter, viewGetCounterThunk)
	exports.AddView(ViewGetInt, viewGetIntThunk)
	exports.AddView(ViewJustView, viewJustViewThunk)
	exports.AddView(ViewPassTypesView, viewPassTypesViewThunk)
	exports.AddView(ViewTestCallPanicViewEPFromView, viewTestCallPanicViewEPFromViewThunk)
	exports.AddView(ViewTestChainOwnerIDView, viewTestChainOwnerIDViewThunk)
	exports.AddView(ViewTestPanicViewEP, viewTestPanicViewEPThunk)
	exports.AddView(ViewTestSandboxCall, viewTestSandboxCallThunk)

	for i, key := range keyMap {
		idxMap[i] = key.KeyID()
	}
}

type CallOnChainContext struct {
	Params  ImmutableCallOnChainParams
	Results MutableCallOnChainResults
	State   MutableTestCoreState
}

func funcCallOnChainThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("testcore.funcCallOnChain")
	f := &CallOnChainContext{
		Params: ImmutableCallOnChainParams{
			id: wasmlib.OBJ_ID_PARAMS,
		},
		Results: MutableCallOnChainResults{
			id: wasmlib.OBJ_ID_RESULTS,
		},
		State: MutableTestCoreState{
			id: wasmlib.OBJ_ID_STATE,
		},
	}
	ctx.Require(f.Params.IntValue().Exists(), "missing mandatory intValue")
	funcCallOnChain(ctx, f)
	ctx.Log("testcore.funcCallOnChain ok")
}

type CheckContextFromFullEPContext struct {
	Params ImmutableCheckContextFromFullEPParams
	State  MutableTestCoreState
}

func funcCheckContextFromFullEPThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("testcore.funcCheckContextFromFullEP")
	f := &CheckContextFromFullEPContext{
		Params: ImmutableCheckContextFromFullEPParams{
			id: wasmlib.OBJ_ID_PARAMS,
		},
		State: MutableTestCoreState{
			id: wasmlib.OBJ_ID_STATE,
		},
	}
	ctx.Require(f.Params.AgentID().Exists(), "missing mandatory agentID")
	ctx.Require(f.Params.Caller().Exists(), "missing mandatory caller")
	ctx.Require(f.Params.ChainID().Exists(), "missing mandatory chainID")
	ctx.Require(f.Params.ChainOwnerID().Exists(), "missing mandatory chainOwnerID")
	ctx.Require(f.Params.ContractCreator().Exists(), "missing mandatory contractCreator")
	funcCheckContextFromFullEP(ctx, f)
	ctx.Log("testcore.funcCheckContextFromFullEP ok")
}

type DoNothingContext struct {
	State MutableTestCoreState
}

func funcDoNothingThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("testcore.funcDoNothing")
	f := &DoNothingContext{
		State: MutableTestCoreState{
			id: wasmlib.OBJ_ID_STATE,
		},
	}
	funcDoNothing(ctx, f)
	ctx.Log("testcore.funcDoNothing ok")
}

type GetMintedSupplyContext struct {
	Results MutableGetMintedSupplyResults
	State   MutableTestCoreState
}

func funcGetMintedSupplyThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("testcore.funcGetMintedSupply")
	f := &GetMintedSupplyContext{
		Results: MutableGetMintedSupplyResults{
			id: wasmlib.OBJ_ID_RESULTS,
		},
		State: MutableTestCoreState{
			id: wasmlib.OBJ_ID_STATE,
		},
	}
	funcGetMintedSupply(ctx, f)
	ctx.Log("testcore.funcGetMintedSupply ok")
}

type IncCounterContext struct {
	State MutableTestCoreState
}

func funcIncCounterThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("testcore.funcIncCounter")
	f := &IncCounterContext{
		State: MutableTestCoreState{
			id: wasmlib.OBJ_ID_STATE,
		},
	}
	funcIncCounter(ctx, f)
	ctx.Log("testcore.funcIncCounter ok")
}

type InitContext struct {
	State MutableTestCoreState
}

func funcInitThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("testcore.funcInit")
	f := &InitContext{
		State: MutableTestCoreState{
			id: wasmlib.OBJ_ID_STATE,
		},
	}
	funcInit(ctx, f)
	ctx.Log("testcore.funcInit ok")
}

type PassTypesFullContext struct {
	Params ImmutablePassTypesFullParams
	State  MutableTestCoreState
}

func funcPassTypesFullThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("testcore.funcPassTypesFull")
	f := &PassTypesFullContext{
		Params: ImmutablePassTypesFullParams{
			id: wasmlib.OBJ_ID_PARAMS,
		},
		State: MutableTestCoreState{
			id: wasmlib.OBJ_ID_STATE,
		},
	}
	ctx.Require(f.Params.Hash().Exists(), "missing mandatory hash")
	ctx.Require(f.Params.Hname().Exists(), "missing mandatory hname")
	ctx.Require(f.Params.HnameZero().Exists(), "missing mandatory hnameZero")
	ctx.Require(f.Params.Int64().Exists(), "missing mandatory int64")
	ctx.Require(f.Params.Int64Zero().Exists(), "missing mandatory int64Zero")
	ctx.Require(f.Params.String().Exists(), "missing mandatory string")
	ctx.Require(f.Params.StringZero().Exists(), "missing mandatory stringZero")
	funcPassTypesFull(ctx, f)
	ctx.Log("testcore.funcPassTypesFull ok")
}

type RunRecursionContext struct {
	Params  ImmutableRunRecursionParams
	Results MutableRunRecursionResults
	State   MutableTestCoreState
}

func funcRunRecursionThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("testcore.funcRunRecursion")
	f := &RunRecursionContext{
		Params: ImmutableRunRecursionParams{
			id: wasmlib.OBJ_ID_PARAMS,
		},
		Results: MutableRunRecursionResults{
			id: wasmlib.OBJ_ID_RESULTS,
		},
		State: MutableTestCoreState{
			id: wasmlib.OBJ_ID_STATE,
		},
	}
	ctx.Require(f.Params.IntValue().Exists(), "missing mandatory intValue")
	funcRunRecursion(ctx, f)
	ctx.Log("testcore.funcRunRecursion ok")
}

type SendToAddressContext struct {
	Params ImmutableSendToAddressParams
	State  MutableTestCoreState
}

func funcSendToAddressThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("testcore.funcSendToAddress")
	ctx.Require(ctx.Caller() == ctx.ContractCreator(), "no permission")

	f := &SendToAddressContext{
		Params: ImmutableSendToAddressParams{
			id: wasmlib.OBJ_ID_PARAMS,
		},
		State: MutableTestCoreState{
			id: wasmlib.OBJ_ID_STATE,
		},
	}
	ctx.Require(f.Params.Address().Exists(), "missing mandatory address")
	funcSendToAddress(ctx, f)
	ctx.Log("testcore.funcSendToAddress ok")
}

type SetIntContext struct {
	Params ImmutableSetIntParams
	State  MutableTestCoreState
}

func funcSetIntThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("testcore.funcSetInt")
	f := &SetIntContext{
		Params: ImmutableSetIntParams{
			id: wasmlib.OBJ_ID_PARAMS,
		},
		State: MutableTestCoreState{
			id: wasmlib.OBJ_ID_STATE,
		},
	}
	ctx.Require(f.Params.IntValue().Exists(), "missing mandatory intValue")
	ctx.Require(f.Params.Name().Exists(), "missing mandatory name")
	funcSetInt(ctx, f)
	ctx.Log("testcore.funcSetInt ok")
}

type TestCallPanicFullEPContext struct {
	State MutableTestCoreState
}

func funcTestCallPanicFullEPThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("testcore.funcTestCallPanicFullEP")
	f := &TestCallPanicFullEPContext{
		State: MutableTestCoreState{
			id: wasmlib.OBJ_ID_STATE,
		},
	}
	funcTestCallPanicFullEP(ctx, f)
	ctx.Log("testcore.funcTestCallPanicFullEP ok")
}

type TestCallPanicViewEPFromFullContext struct {
	State MutableTestCoreState
}

func funcTestCallPanicViewEPFromFullThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("testcore.funcTestCallPanicViewEPFromFull")
	f := &TestCallPanicViewEPFromFullContext{
		State: MutableTestCoreState{
			id: wasmlib.OBJ_ID_STATE,
		},
	}
	funcTestCallPanicViewEPFromFull(ctx, f)
	ctx.Log("testcore.funcTestCallPanicViewEPFromFull ok")
}

type TestChainOwnerIDFullContext struct {
	Results MutableTestChainOwnerIDFullResults
	State   MutableTestCoreState
}

func funcTestChainOwnerIDFullThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("testcore.funcTestChainOwnerIDFull")
	f := &TestChainOwnerIDFullContext{
		Results: MutableTestChainOwnerIDFullResults{
			id: wasmlib.OBJ_ID_RESULTS,
		},
		State: MutableTestCoreState{
			id: wasmlib.OBJ_ID_STATE,
		},
	}
	funcTestChainOwnerIDFull(ctx, f)
	ctx.Log("testcore.funcTestChainOwnerIDFull ok")
}

type TestEventLogDeployContext struct {
	State MutableTestCoreState
}

func funcTestEventLogDeployThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("testcore.funcTestEventLogDeploy")
	f := &TestEventLogDeployContext{
		State: MutableTestCoreState{
			id: wasmlib.OBJ_ID_STATE,
		},
	}
	funcTestEventLogDeploy(ctx, f)
	ctx.Log("testcore.funcTestEventLogDeploy ok")
}

type TestEventLogEventDataContext struct {
	State MutableTestCoreState
}

func funcTestEventLogEventDataThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("testcore.funcTestEventLogEventData")
	f := &TestEventLogEventDataContext{
		State: MutableTestCoreState{
			id: wasmlib.OBJ_ID_STATE,
		},
	}
	funcTestEventLogEventData(ctx, f)
	ctx.Log("testcore.funcTestEventLogEventData ok")
}

type TestEventLogGenericDataContext struct {
	Params ImmutableTestEventLogGenericDataParams
	State  MutableTestCoreState
}

func funcTestEventLogGenericDataThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("testcore.funcTestEventLogGenericData")
	f := &TestEventLogGenericDataContext{
		Params: ImmutableTestEventLogGenericDataParams{
			id: wasmlib.OBJ_ID_PARAMS,
		},
		State: MutableTestCoreState{
			id: wasmlib.OBJ_ID_STATE,
		},
	}
	ctx.Require(f.Params.Counter().Exists(), "missing mandatory counter")
	funcTestEventLogGenericData(ctx, f)
	ctx.Log("testcore.funcTestEventLogGenericData ok")
}

type TestPanicFullEPContext struct {
	State MutableTestCoreState
}

func funcTestPanicFullEPThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("testcore.funcTestPanicFullEP")
	f := &TestPanicFullEPContext{
		State: MutableTestCoreState{
			id: wasmlib.OBJ_ID_STATE,
		},
	}
	funcTestPanicFullEP(ctx, f)
	ctx.Log("testcore.funcTestPanicFullEP ok")
}

type WithdrawToChainContext struct {
	Params ImmutableWithdrawToChainParams
	State  MutableTestCoreState
}

func funcWithdrawToChainThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("testcore.funcWithdrawToChain")
	f := &WithdrawToChainContext{
		Params: ImmutableWithdrawToChainParams{
			id: wasmlib.OBJ_ID_PARAMS,
		},
		State: MutableTestCoreState{
			id: wasmlib.OBJ_ID_STATE,
		},
	}
	ctx.Require(f.Params.ChainID().Exists(), "missing mandatory chainID")
	funcWithdrawToChain(ctx, f)
	ctx.Log("testcore.funcWithdrawToChain ok")
}

type CheckContextFromViewEPContext struct {
	Params ImmutableCheckContextFromViewEPParams
	State  ImmutableTestCoreState
}

func viewCheckContextFromViewEPThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("testcore.viewCheckContextFromViewEP")
	f := &CheckContextFromViewEPContext{
		Params: ImmutableCheckContextFromViewEPParams{
			id: wasmlib.OBJ_ID_PARAMS,
		},
		State: ImmutableTestCoreState{
			id: wasmlib.OBJ_ID_STATE,
		},
	}
	ctx.Require(f.Params.AgentID().Exists(), "missing mandatory agentID")
	ctx.Require(f.Params.ChainID().Exists(), "missing mandatory chainID")
	ctx.Require(f.Params.ChainOwnerID().Exists(), "missing mandatory chainOwnerID")
	ctx.Require(f.Params.ContractCreator().Exists(), "missing mandatory contractCreator")
	viewCheckContextFromViewEP(ctx, f)
	ctx.Log("testcore.viewCheckContextFromViewEP ok")
}

type FibonacciContext struct {
	Params  ImmutableFibonacciParams
	Results MutableFibonacciResults
	State   ImmutableTestCoreState
}

func viewFibonacciThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("testcore.viewFibonacci")
	f := &FibonacciContext{
		Params: ImmutableFibonacciParams{
			id: wasmlib.OBJ_ID_PARAMS,
		},
		Results: MutableFibonacciResults{
			id: wasmlib.OBJ_ID_RESULTS,
		},
		State: ImmutableTestCoreState{
			id: wasmlib.OBJ_ID_STATE,
		},
	}
	ctx.Require(f.Params.IntValue().Exists(), "missing mandatory intValue")
	viewFibonacci(ctx, f)
	ctx.Log("testcore.viewFibonacci ok")
}

type GetCounterContext struct {
	Results MutableGetCounterResults
	State   ImmutableTestCoreState
}

func viewGetCounterThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("testcore.viewGetCounter")
	f := &GetCounterContext{
		Results: MutableGetCounterResults{
			id: wasmlib.OBJ_ID_RESULTS,
		},
		State: ImmutableTestCoreState{
			id: wasmlib.OBJ_ID_STATE,
		},
	}
	viewGetCounter(ctx, f)
	ctx.Log("testcore.viewGetCounter ok")
}

type GetIntContext struct {
	Params  ImmutableGetIntParams
	Results MutableGetIntResults
	State   ImmutableTestCoreState
}

func viewGetIntThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("testcore.viewGetInt")
	f := &GetIntContext{
		Params: ImmutableGetIntParams{
			id: wasmlib.OBJ_ID_PARAMS,
		},
		Results: MutableGetIntResults{
			id: wasmlib.OBJ_ID_RESULTS,
		},
		State: ImmutableTestCoreState{
			id: wasmlib.OBJ_ID_STATE,
		},
	}
	ctx.Require(f.Params.Name().Exists(), "missing mandatory name")
	viewGetInt(ctx, f)
	ctx.Log("testcore.viewGetInt ok")
}

type JustViewContext struct {
	State ImmutableTestCoreState
}

func viewJustViewThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("testcore.viewJustView")
	f := &JustViewContext{
		State: ImmutableTestCoreState{
			id: wasmlib.OBJ_ID_STATE,
		},
	}
	viewJustView(ctx, f)
	ctx.Log("testcore.viewJustView ok")
}

type PassTypesViewContext struct {
	Params ImmutablePassTypesViewParams
	State  ImmutableTestCoreState
}

func viewPassTypesViewThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("testcore.viewPassTypesView")
	f := &PassTypesViewContext{
		Params: ImmutablePassTypesViewParams{
			id: wasmlib.OBJ_ID_PARAMS,
		},
		State: ImmutableTestCoreState{
			id: wasmlib.OBJ_ID_STATE,
		},
	}
	ctx.Require(f.Params.Hash().Exists(), "missing mandatory hash")
	ctx.Require(f.Params.Hname().Exists(), "missing mandatory hname")
	ctx.Require(f.Params.HnameZero().Exists(), "missing mandatory hnameZero")
	ctx.Require(f.Params.Int64().Exists(), "missing mandatory int64")
	ctx.Require(f.Params.Int64Zero().Exists(), "missing mandatory int64Zero")
	ctx.Require(f.Params.String().Exists(), "missing mandatory string")
	ctx.Require(f.Params.StringZero().Exists(), "missing mandatory stringZero")
	viewPassTypesView(ctx, f)
	ctx.Log("testcore.viewPassTypesView ok")
}

type TestCallPanicViewEPFromViewContext struct {
	State ImmutableTestCoreState
}

func viewTestCallPanicViewEPFromViewThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("testcore.viewTestCallPanicViewEPFromView")
	f := &TestCallPanicViewEPFromViewContext{
		State: ImmutableTestCoreState{
			id: wasmlib.OBJ_ID_STATE,
		},
	}
	viewTestCallPanicViewEPFromView(ctx, f)
	ctx.Log("testcore.viewTestCallPanicViewEPFromView ok")
}

type TestChainOwnerIDViewContext struct {
	Results MutableTestChainOwnerIDViewResults
	State   ImmutableTestCoreState
}

func viewTestChainOwnerIDViewThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("testcore.viewTestChainOwnerIDView")
	f := &TestChainOwnerIDViewContext{
		Results: MutableTestChainOwnerIDViewResults{
			id: wasmlib.OBJ_ID_RESULTS,
		},
		State: ImmutableTestCoreState{
			id: wasmlib.OBJ_ID_STATE,
		},
	}
	viewTestChainOwnerIDView(ctx, f)
	ctx.Log("testcore.viewTestChainOwnerIDView ok")
}

type TestPanicViewEPContext struct {
	State ImmutableTestCoreState
}

func viewTestPanicViewEPThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("testcore.viewTestPanicViewEP")
	f := &TestPanicViewEPContext{
		State: ImmutableTestCoreState{
			id: wasmlib.OBJ_ID_STATE,
		},
	}
	viewTestPanicViewEP(ctx, f)
	ctx.Log("testcore.viewTestPanicViewEP ok")
}

type TestSandboxCallContext struct {
	Results MutableTestSandboxCallResults
	State   ImmutableTestCoreState
}

func viewTestSandboxCallThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("testcore.viewTestSandboxCall")
	f := &TestSandboxCallContext{
		Results: MutableTestSandboxCallResults{
			id: wasmlib.OBJ_ID_RESULTS,
		},
		State: ImmutableTestCoreState{
			id: wasmlib.OBJ_ID_STATE,
		},
	}
	viewTestSandboxCall(ctx, f)
	ctx.Log("testcore.viewTestSandboxCall ok")
}
