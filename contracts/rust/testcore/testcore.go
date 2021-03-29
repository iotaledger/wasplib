// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package testcore

import (
	"github.com/iotaledger/wasplib/packages/vm/wasmlib"
)

const ContractNameDeployed = "exampleDeployTR"
const MsgFullPanic = "========== panic FULL ENTRY POINT ========="
const MsgViewPanic = "========== panic VIEW ========="

func funcCallOnChain(ctx wasmlib.ScFuncContext, params *FuncCallOnChainParams) {
	paramInt := params.IntValue.Value()

	targetContract := ctx.Contract()
	if params.HnameContract.Exists() {
		targetContract = params.HnameContract.Value()
	}

	targetEp := HFuncCallOnChain
	if params.HnameEP.Exists() {
		targetEp = params.HnameEP.Value()
	}

	varCounter := ctx.State().GetInt64(VarCounter)
	counter := varCounter.Value()
	varCounter.SetValue(counter + 1)

	ctx.Log("call depth = " + params.IntValue.String() +
		" hnameContract = " + targetContract.String() +
		" hnameEP = " + targetEp.String() +
		" counter = " + ctx.Utility().String(counter))

	parms := wasmlib.NewScMutableMap()
	parms.GetInt64(ParamIntValue).SetValue(paramInt)
	ret := ctx.Call(targetContract, targetEp, parms, nil)

	retVal := ret.GetInt64(ParamIntValue)
	ctx.Results().GetInt64(ParamIntValue).SetValue(retVal.Value())
}

func funcCheckContextFromFullEP(ctx wasmlib.ScFuncContext, params *FuncCheckContextFromFullEPParams) {
	ctx.Require(params.AgentId.Value() == ctx.AccountId(), "fail: agentID")
	ctx.Require(params.Caller.Value() == ctx.Caller(), "fail: caller")
	ctx.Require(params.ChainId.Value() == ctx.ChainId(), "fail: chainID")
	ctx.Require(params.ChainOwnerId.Value() == ctx.ChainOwnerId(), "fail: chainOwnerID")
	ctx.Require(params.ContractCreator.Value() == ctx.ContractCreator(), "fail: contractCreator")
}

func funcDoNothing(ctx wasmlib.ScFuncContext, params *FuncDoNothingParams) {
	ctx.Log("doing nothing...")
}

func funcGetMintedSupply(ctx wasmlib.ScFuncContext, params *FuncGetMintedSupplyParams) {
	minted := ctx.Minted()
	mintedColors := minted.Colors()
	ctx.Require(mintedColors.Length() == 1, "test only supports one minted color")
	color := mintedColors.GetColor(0).Value()
	amount := minted.Balance(color)
	ctx.Results().GetInt64(VarMintedSupply).SetValue(amount)
	ctx.Results().GetColor(VarMintedColor).SetValue(color)
}

func funcIncCounter(ctx wasmlib.ScFuncContext, params *FuncIncCounterParams) {
	ctx.State().GetInt64(VarCounter).SetValue(ctx.State().GetInt64(VarCounter).Value() + 1)
}

func funcInit(ctx wasmlib.ScFuncContext, params *FuncInitParams) {
	ctx.Log("doing nothing...")
}

func funcPassTypesFull(ctx wasmlib.ScFuncContext, params *FuncPassTypesFullParams) {
	hash := ctx.Utility().HashBlake2b([]byte(ParamHash))
	ctx.Require(params.Hash.Value() == hash, "Hash wrong")
	ctx.Require(params.Int64.Value() == 42, "int64 wrong")
	ctx.Require(params.Int64Zero.Value() == 0, "int64-0 wrong")
	ctx.Require(params.String.Value() == string(ParamString), "string wrong")
	ctx.Require(params.StringZero.Value() == "", "string-0 wrong")
	ctx.Require(params.Hname.Value() == wasmlib.NewScHname(string(ParamHname)), "Hname wrong")
	ctx.Require(params.HnameZero.Value() == wasmlib.ScHname(0), "Hname-0 wrong")
}

func funcRunRecursion(ctx wasmlib.ScFuncContext, params *FuncRunRecursionParams) {
	depth := params.IntValue.Value()
	if depth <= 0 {
		return
	}

	parms := wasmlib.NewScMutableMap()
	parms.GetInt64(ParamIntValue).SetValue(depth - 1)
	parms.GetHname(ParamHnameEP).SetValue(HFuncRunRecursion)
	ctx.CallSelf(HFuncCallOnChain, parms, nil)
	// TODO how would I return result of the call ???
	ctx.Results().GetInt64(ParamIntValue).SetValue(depth - 1)
}

func funcSendToAddress(ctx wasmlib.ScFuncContext, params *FuncSendToAddressParams) {
	balances := wasmlib.NewScTransfersFromBalances(ctx.Balances())
	ctx.TransferToAddress(params.Address.Value(), balances)
}

func funcSetInt(ctx wasmlib.ScFuncContext, params *FuncSetIntParams) {
	ctx.State().GetInt64(wasmlib.Key(params.Name.Value())).SetValue(params.IntValue.Value())
}

func funcTestCallPanicFullEP(ctx wasmlib.ScFuncContext, params *FuncTestCallPanicFullEPParams) {
	ctx.CallSelf(HFuncTestPanicFullEP, nil, nil)
}

func funcTestCallPanicViewEPFromFull(ctx wasmlib.ScFuncContext, params *FuncTestCallPanicViewEPFromFullParams) {
	ctx.CallSelf(HViewTestPanicViewEP, nil, nil)
}

func funcTestChainOwnerIDFull(ctx wasmlib.ScFuncContext, params *FuncTestChainOwnerIDFullParams) {
	ctx.Results().GetAgentId(ParamChainOwnerId).SetValue(ctx.ChainOwnerId())
}

func funcTestEventLogDeploy(ctx wasmlib.ScFuncContext, params *FuncTestEventLogDeployParams) {
	//Deploy the same contract with another name
	programHash := ctx.Utility().HashBlake2b([]byte("test_sandbox"))
	ctx.Deploy(programHash, ContractNameDeployed, "test contract deploy log", nil)
}

func funcTestEventLogEventData(ctx wasmlib.ScFuncContext, params *FuncTestEventLogEventDataParams) {
	ctx.Event("[Event] - Testing Event...")
}

func funcTestEventLogGenericData(ctx wasmlib.ScFuncContext, params *FuncTestEventLogGenericDataParams) {
	event := "[GenericData] Counter Number: " + params.Counter.String()
	ctx.Event(event)
}

func funcTestPanicFullEP(ctx wasmlib.ScFuncContext, params *FuncTestPanicFullEPParams) {
	ctx.Panic(MsgFullPanic)
}

func funcWithdrawToChain(ctx wasmlib.ScFuncContext, params *FuncWithdrawToChainParams) {
	transfers := wasmlib.NewScTransferIotas(1)
	ctx.Post(params.ChainId.Value(), wasmlib.CoreAccounts, wasmlib.CoreAccountsFuncWithdraw, nil, transfers, 0)
}

func viewCheckContextFromViewEP(ctx wasmlib.ScViewContext, params *ViewCheckContextFromViewEPParams) {
	ctx.Require(params.AgentId.Value() == ctx.AccountId(), "fail: agentID")
	ctx.Require(params.ChainId.Value() == ctx.ChainId(), "fail: chainID")
	ctx.Require(params.ChainOwnerId.Value() == ctx.ChainOwnerId(), "fail: chainOwnerID")
	ctx.Require(params.ContractCreator.Value() == ctx.ContractCreator(), "fail: contractCreator")
}

func viewFibonacci(ctx wasmlib.ScViewContext, params *ViewFibonacciParams) {
	n := params.IntValue.Value()
	if n == 0 || n == 1 {
		ctx.Results().GetInt64(ParamIntValue).SetValue(n)
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

	ctx.Results().GetInt64(ParamIntValue).SetValue(n1 + n2)
}

func viewGetCounter(ctx wasmlib.ScViewContext, params *ViewGetCounterParams) {
	counter := ctx.State().GetInt64(VarCounter)
	ctx.Results().GetInt64(VarCounter).SetValue(counter.Value())
}

func viewGetInt(ctx wasmlib.ScViewContext, params *ViewGetIntParams) {
	name := params.Name.Value()
	value := ctx.State().GetInt64(wasmlib.Key(name))
	ctx.Require(value.Exists(), "param 'value' not found")
	ctx.Results().GetInt64(wasmlib.Key(name)).SetValue(value.Value())
}

func viewJustView(ctx wasmlib.ScViewContext, params *ViewJustViewParams) {
	ctx.Log("doing nothing...")
}

func viewPassTypesView(ctx wasmlib.ScViewContext, params *ViewPassTypesViewParams) {
	hash := ctx.Utility().HashBlake2b([]byte(ParamHash))
	ctx.Require(params.Hash.Value() == hash, "Hash wrong")
	ctx.Require(params.Int64.Value() == 42, "int64 wrong")
	ctx.Require(params.Int64Zero.Value() == 0, "int64-0 wrong")
	ctx.Require(params.String.Value() == string(ParamString), "string wrong")
	ctx.Require(params.StringZero.Value() == "", "string-0 wrong")
	ctx.Require(params.Hname.Value() == wasmlib.NewScHname(string(ParamHname)), "Hname wrong")
	ctx.Require(params.HnameZero.Value() == wasmlib.ScHname(0), "Hname-0 wrong")
}

func viewTestCallPanicViewEPFromView(ctx wasmlib.ScViewContext, params *ViewTestCallPanicViewEPFromViewParams) {
	ctx.CallSelf(HViewTestPanicViewEP, nil)
}

func viewTestChainOwnerIDView(ctx wasmlib.ScViewContext, params *ViewTestChainOwnerIDViewParams) {
	ctx.Results().GetAgentId(ParamChainOwnerId).SetValue(ctx.ChainOwnerId())
}

func viewTestPanicViewEP(ctx wasmlib.ScViewContext, params *ViewTestPanicViewEPParams) {
	ctx.Panic(MsgViewPanic)
}

func viewTestSandboxCall(ctx wasmlib.ScViewContext, params *ViewTestSandboxCallParams) {
	ret := ctx.Call(wasmlib.CoreRoot, wasmlib.CoreRootViewGetChainInfo, nil)
	desc := ret.GetString(wasmlib.Key("d")).Value()
	ctx.Results().GetString(wasmlib.Key("sandboxCall")).SetValue(desc)
}
