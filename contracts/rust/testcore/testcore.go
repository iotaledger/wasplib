// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package testcore

import (
	"github.com/iotaledger/wasplib/packages/vm/wasmlib"
)

const ContractNameDeployed = "exampleDeployTR"
const MsgFullPanic = "========== panic FULL ENTRY POINT ========="
const MsgViewPanic = "========== panic VIEW ========="

func funcCallOnChain(ctx wasmlib.ScFuncContext, f *FuncCallOnChainContext) {
	paramInt := f.Params.IntValue().Value()

	targetContract := ctx.Contract()
	if f.Params.HnameContract().Exists() {
		targetContract = f.Params.HnameContract().Value()
	}

	targetEp := HFuncCallOnChain
	if f.Params.HnameEP().Exists() {
		targetEp = f.Params.HnameEP().Value()
	}

	varCounter := f.State.Counter()

	ctx.Log("call depth = " + f.Params.IntValue().String() +
		", hnameContract = " + targetContract.String() +
		", hnameEP = " + targetEp.String() +
		", counter = " + varCounter.String())

	varCounter.SetValue(varCounter.Value() + 1)

	params := wasmlib.NewScMutableMap()
	params.GetInt64(ParamIntValue).SetValue(paramInt)
	ret := ctx.Call(targetContract, targetEp, params, nil)
	retVal := ret.GetInt64(ParamIntValue)
	f.Results.IntValue().SetValue(retVal.Value())
}

func funcCheckContextFromFullEP(ctx wasmlib.ScFuncContext, f *FuncCheckContextFromFullEPContext) {
	ctx.Require(f.Params.AgentId().Value() == ctx.AccountId(), "fail: agentID")
	ctx.Require(f.Params.Caller().Value() == ctx.Caller(), "fail: caller")
	ctx.Require(f.Params.ChainId().Value() == ctx.ChainId(), "fail: chainID")
	ctx.Require(f.Params.ChainOwnerId().Value() == ctx.ChainOwnerId(), "fail: chainOwnerID")
	ctx.Require(f.Params.ContractCreator().Value() == ctx.ContractCreator(), "fail: contractCreator")
}

func funcDoNothing(ctx wasmlib.ScFuncContext, f *FuncDoNothingContext) {
	ctx.Log("doing nothing...")
}

func funcGetMintedSupply(ctx wasmlib.ScFuncContext, f *FuncGetMintedSupplyContext) {
	minted := ctx.Minted()
	mintedColors := minted.Colors()
	ctx.Require(mintedColors.Length() == 1, "test only supports one minted color")
	color := mintedColors.GetColor(0).Value()
	amount := minted.Balance(color)
	f.Results.MintedColor().SetValue(color)
	f.Results.MintedSupply().SetValue(amount)
}

func funcIncCounter(ctx wasmlib.ScFuncContext, f *FuncIncCounterContext) {
	counter := f.State.Counter()
	counter.SetValue(counter.Value() + 1)
}

func funcInit(ctx wasmlib.ScFuncContext, f *FuncInitContext) {
	ctx.Log("doing nothing...")
}

func funcPassTypesFull(ctx wasmlib.ScFuncContext, f *FuncPassTypesFullContext) {
	hash := ctx.Utility().HashBlake2b([]byte(ParamHash))
	ctx.Require(f.Params.Hash().Value() == hash, "Hash wrong")
	ctx.Require(f.Params.Int64().Value() == 42, "int64 wrong")
	ctx.Require(f.Params.Int64Zero().Value() == 0, "int64-0 wrong")
	ctx.Require(f.Params.String().Value() == string(ParamString), "string wrong")
	ctx.Require(f.Params.StringZero().Value() == "", "string-0 wrong")
	ctx.Require(f.Params.Hname().Value() == wasmlib.NewScHname(string(ParamHname)), "Hname wrong")
	ctx.Require(f.Params.HnameZero().Value() == wasmlib.ScHname(0), "Hname-0 wrong")
}

func funcRunRecursion(ctx wasmlib.ScFuncContext, f *FuncRunRecursionContext) {
	depth := f.Params.IntValue().Value()
	if depth <= 0 {
		return
	}

	parms := wasmlib.NewScMutableMap()
	parms.GetInt64(ParamIntValue).SetValue(depth - 1)
	parms.GetHname(ParamHnameEP).SetValue(HFuncRunRecursion)
	ctx.CallSelf(HFuncCallOnChain, parms, nil)
	// TODO how would I return result of the call ???
	f.Results.IntValue().SetValue(depth - 1)
}

func funcSendToAddress(ctx wasmlib.ScFuncContext, f *FuncSendToAddressContext) {
	balances := wasmlib.NewScTransfersFromBalances(ctx.Balances())
	ctx.TransferToAddress(f.Params.Address().Value(), balances)
}

func funcSetInt(ctx wasmlib.ScFuncContext, f *FuncSetIntContext) {
	ctx.State().GetInt64(wasmlib.Key(f.Params.Name().Value())).SetValue(f.Params.IntValue().Value())
}

func funcTestCallPanicFullEP(ctx wasmlib.ScFuncContext, f *FuncTestCallPanicFullEPContext) {
	ctx.CallSelf(HFuncTestPanicFullEP, nil, nil)
}

func funcTestCallPanicViewEPFromFull(ctx wasmlib.ScFuncContext, f *FuncTestCallPanicViewEPFromFullContext) {
	ctx.CallSelf(HViewTestPanicViewEP, nil, nil)
}

func funcTestChainOwnerIDFull(ctx wasmlib.ScFuncContext, f *FuncTestChainOwnerIDFullContext) {
	f.Results.ChainOwnerId().SetValue(ctx.ChainOwnerId())
}

func funcTestEventLogDeploy(ctx wasmlib.ScFuncContext, f *FuncTestEventLogDeployContext) {
	// deploy the same contract with another name
	programHash := ctx.Utility().HashBlake2b([]byte("testcore"))
	ctx.Deploy(programHash, ContractNameDeployed, "test contract deploy log", nil)
}

func funcTestEventLogEventData(ctx wasmlib.ScFuncContext, f *FuncTestEventLogEventDataContext) {
	ctx.Event("[Event] - Testing Event...")
}

func funcTestEventLogGenericData(ctx wasmlib.ScFuncContext, f *FuncTestEventLogGenericDataContext) {
	event := "[GenericData] Counter Number: " + f.Params.Counter().String()
	ctx.Event(event)
}

func funcTestPanicFullEP(ctx wasmlib.ScFuncContext, f *FuncTestPanicFullEPContext) {
	ctx.Panic(MsgFullPanic)
}

func funcWithdrawToChain(ctx wasmlib.ScFuncContext, f *FuncWithdrawToChainContext) {
	transfers := wasmlib.NewScTransferIotas(1)
	ctx.Post(f.Params.ChainId().Value(), wasmlib.CoreAccounts, wasmlib.CoreAccountsFuncWithdraw, nil, transfers, 0)
}

func viewCheckContextFromViewEP(ctx wasmlib.ScViewContext, f *ViewCheckContextFromViewEPContext) {
	ctx.Require(f.Params.AgentId().Value() == ctx.AccountId(), "fail: agentID")
	ctx.Require(f.Params.ChainId().Value() == ctx.ChainId(), "fail: chainID")
	ctx.Require(f.Params.ChainOwnerId().Value() == ctx.ChainOwnerId(), "fail: chainOwnerID")
	ctx.Require(f.Params.ContractCreator().Value() == ctx.ContractCreator(), "fail: contractCreator")
}

func viewFibonacci(ctx wasmlib.ScViewContext, f *ViewFibonacciContext) {
	n := f.Params.IntValue().Value()
	if n == 0 || n == 1 {
		f.Results.IntValue().SetValue(n)
		return
	}
	parms1 := wasmlib.NewScMutableMap()
	parms1.GetInt64(ParamIntValue).SetValue(n - 1)
	results1 := ctx.CallSelf(HViewFibonacci, parms1)
	n1 := results1.GetInt64(ParamIntValue).Value()

	parms2 := wasmlib.NewScMutableMap()
	parms2.GetInt64(ParamIntValue).SetValue(n - 2)
	results2 := ctx.CallSelf(HViewFibonacci, parms2)
	n2 := results2.GetInt64(ParamIntValue).Value()

	f.Results.IntValue().SetValue(n1 + n2)
}

func viewGetCounter(ctx wasmlib.ScViewContext, f *ViewGetCounterContext) {
	f.Results.Counter().SetValue(f.State.Counter().Value())
}

func viewGetInt(ctx wasmlib.ScViewContext, f *ViewGetIntContext) {
	name := f.Params.Name().Value()
	value := ctx.State().GetInt64(wasmlib.Key(name))
	ctx.Require(value.Exists(), "param 'value' not found")
	ctx.Results().GetInt64(wasmlib.Key(name)).SetValue(value.Value())
}

func viewJustView(ctx wasmlib.ScViewContext, f *ViewJustViewContext) {
	ctx.Log("doing nothing...")
}

func viewPassTypesView(ctx wasmlib.ScViewContext, f *ViewPassTypesViewContext) {
	hash := ctx.Utility().HashBlake2b([]byte(ParamHash))
	ctx.Require(f.Params.Hash().Value() == hash, "Hash wrong")
	ctx.Require(f.Params.Int64().Value() == 42, "int64 wrong")
	ctx.Require(f.Params.Int64Zero().Value() == 0, "int64-0 wrong")
	ctx.Require(f.Params.String().Value() == string(ParamString), "string wrong")
	ctx.Require(f.Params.StringZero().Value() == "", "string-0 wrong")
	ctx.Require(f.Params.Hname().Value() == wasmlib.NewScHname(string(ParamHname)), "Hname wrong")
	ctx.Require(f.Params.HnameZero().Value() == wasmlib.ScHname(0), "Hname-0 wrong")
}

func viewTestCallPanicViewEPFromView(ctx wasmlib.ScViewContext, f *ViewTestCallPanicViewEPFromViewContext) {
	ctx.CallSelf(HViewTestPanicViewEP, nil)
}

func viewTestChainOwnerIDView(ctx wasmlib.ScViewContext, f *ViewTestChainOwnerIDViewContext) {
	f.Results.ChainOwnerId().SetValue(ctx.ChainOwnerId())
}

func viewTestPanicViewEP(ctx wasmlib.ScViewContext, f *ViewTestPanicViewEPContext) {
	ctx.Panic(MsgViewPanic)
}

func viewTestSandboxCall(ctx wasmlib.ScViewContext, f *ViewTestSandboxCallContext) {
	ret := ctx.Call(wasmlib.CoreRoot, wasmlib.CoreRootViewGetChainInfo, nil)
	desc := ret.GetString(wasmlib.Key("d")).Value()
	f.Results.SandboxCall().SetValue(desc)
}
