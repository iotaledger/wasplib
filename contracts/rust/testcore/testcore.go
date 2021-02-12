// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package testcore

import "github.com/iotaledger/wasp/packages/vm/wasmlib"

const ContractNameDeployed = "exampleDeployTR"
const MsgFullPanic = "========== panic FULL ENTRY POINT ========="
const MsgViewPanic = "========== panic VIEW ========="

func funcCallOnChain(ctx *wasmlib.ScFuncContext, params *FuncCallOnChainParams) {
	ctx.Log("calling callOnChain")
	paramIn := params.IntValue.Value()

	targetContract := ctx.ContractId().Hname()
	if params.HnameContract.Exists() {
		targetContract = params.HnameContract.Value()
	}

	targetEp := HFuncCallOnChain
	if params.HnameEP.Exists() {
		targetEp = params.HnameEP.Value()
	}

	varCounter := ctx.State().GetInt(VarCounter)
	counter := varCounter.Value()
	varCounter.SetValue(counter + 1)

	msg := "call depth = " + params.IntValue.String() +
		" hnameContract = " + targetContract.String() +
		" hnameEP = " + targetEp.String() +
		" counter = " + ctx.Utility().String(counter)
	ctx.Log(msg)

	par := wasmlib.NewScMutableMap()
	par.GetInt(ParamIntValue).SetValue(paramIn)
	ret := ctx.Call(targetContract, targetEp, par, nil)

	retVal := ret.GetInt(ParamIntValue)

	ctx.Results().GetInt(ParamIntValue).SetValue(retVal.Value())
}

func funcCheckContextFromFullEP(ctx *wasmlib.ScFuncContext, params *FuncCheckContextFromFullEPParams) {
	ctx.Log("calling checkContextFromFullEP")

	ctx.Require(params.ChainId.Value().Equals(ctx.ContractId().ChainId()), "fail: chainID")
	ctx.Require(params.ChainOwnerId.Value().Equals(ctx.ChainOwnerId()), "fail: chainOwnerID")
	ctx.Require(params.Caller.Value().Equals(ctx.Caller()), "fail: caller")
	ctx.Require(params.ContractId.Value().Equals(ctx.ContractId()), "fail: contractID")
	ctx.Require(params.AgentId.Value().Equals(ctx.ContractId().AsAgentId()), "fail: agentID")
	ctx.Require(params.ContractCreator.Value().Equals(ctx.ContractCreator()), "fail: contractCreator")
}

func funcDoNothing(ctx *wasmlib.ScFuncContext, params *FuncDoNothingParams) {
	ctx.Log("calling doNothing")
}

func funcInit(ctx *wasmlib.ScFuncContext, params *FuncInitParams) {
	ctx.Log("calling init")
}

func funcPassTypesFull(ctx *wasmlib.ScFuncContext, params *FuncPassTypesFullParams) {
	ctx.Log("calling passTypesFull")

	ctx.Require(params.Int64.Value() == 42, "int64 wrong")
	ctx.Require(params.Int64Zero.Value() == 0, "int64-0 wrong")
	ctx.Require(params.String.Value() == string(ParamString), "string wrong")
	ctx.Require(params.StringZero.Value() == "", "string-0 wrong")

	hash := ctx.Utility().HashBlake2b([]byte(ParamHash))
	ctx.Require(params.Hash.Value().Equals(hash), "Hash wrong")

	ctx.Require(params.Hname.Value().Equals(wasmlib.NewScHname(string(ParamHname))), "Hname wrong")
	ctx.Require(params.HnameZero.Value().Equals(wasmlib.ScHname(0)), "Hname-0 wrong")
}

func funcRunRecursion(ctx *wasmlib.ScFuncContext, params *FuncRunRecursionParams) {
	ctx.Log("calling runRecursion")
	depth := params.IntValue.Value()
	if depth <= 0 {
		return
	}
	par := wasmlib.NewScMutableMap()
	par.GetInt(ParamIntValue).SetValue(depth - 1)
	par.GetHname(VarHnameEP).SetValue(HFuncRunRecursion)
	ctx.CallSelf(HFuncCallOnChain, par, nil)
	// TODO how would I return result of the call ???
	ctx.Results().GetInt(ParamIntValue).SetValue(depth - 1)
}

func funcSendToAddress(ctx *wasmlib.ScFuncContext, params *FuncSendToAddressParams) {
	ctx.Log("calling sendToAddress")
	ctx.TransferToAddress(params.Address.Value(), ctx.Balances())
}

func funcSetInt(ctx *wasmlib.ScFuncContext, params *FuncSetIntParams) {
	ctx.Log("calling setInt")
	ctx.State().GetInt(wasmlib.Key(params.Name.Value())).SetValue(params.IntValue.Value())
}

func funcTestCallPanicFullEP(ctx *wasmlib.ScFuncContext, params *FuncTestCallPanicFullEPParams) {
	ctx.Log("calling testCallPanicFullEP")
	ctx.CallSelf(HFuncTestPanicFullEP, nil, nil)
}

func funcTestCallPanicViewEPFromFull(ctx *wasmlib.ScFuncContext, params *FuncTestCallPanicViewEPFromFullParams) {
	ctx.Log("calling testCallPanicViewEPFromFull")
	ctx.CallSelf(HViewTestPanicViewEP, nil, nil)
}

func funcTestChainOwnerIDFull(ctx *wasmlib.ScFuncContext, params *FuncTestChainOwnerIDFullParams) {
	ctx.Log("calling testChainOwnerIDFull")
	ctx.Results().GetAgentId(ParamChainOwnerId).SetValue(ctx.ChainOwnerId())
}

func funcTestContractIDFull(ctx *wasmlib.ScFuncContext, params *FuncTestContractIDFullParams) {
	ctx.Log("calling testContractIDFull")
	ctx.Results().GetContractId(ParamContractId).SetValue(ctx.ContractId())
}

func funcTestEventLogDeploy(ctx *wasmlib.ScFuncContext, params *FuncTestEventLogDeployParams) {
	ctx.Log("calling testEventLogDeploy")
	//Deploy the same contract with another name
	programHash := ctx.Utility().HashBlake2b([]byte("test_sandbox"))
	ctx.Deploy(programHash, string(ContractNameDeployed),
		"test contract deploy log", nil)
}

func funcTestEventLogEventData(ctx *wasmlib.ScFuncContext, params *FuncTestEventLogEventDataParams) {
	ctx.Log("calling testEventLogEventData")
	ctx.Event("[Event] - Testing Event...")
}

func funcTestEventLogGenericData(ctx *wasmlib.ScFuncContext, params *FuncTestEventLogGenericDataParams) {
	ctx.Log("calling testEventLogGenericData")
	event := "[GenericData] Counter Number: " + params.Counter.String()
	ctx.Event(event)
}

func funcTestPanicFullEP(ctx *wasmlib.ScFuncContext, params *FuncTestPanicFullEPParams) {
	ctx.Log("calling testPanicFullEP")
	ctx.Panic(MsgFullPanic)
}

func funcWithdrawToChain(ctx *wasmlib.ScFuncContext, params *FuncWithdrawToChainParams) {
	ctx.Log("calling withdrawToChain")
	//Deploy the same contract with another name
	targetContractId := wasmlib.NewScContractId(params.ChainId.Value(), wasmlib.CoreAccounts)
	ctx.Post(&wasmlib.PostRequestParams{
		ContractId: targetContractId,
		Function:   wasmlib.CoreAccountsFuncWithdrawToChain,
		Params:     nil,
		Transfer:   wasmlib.NewScTransfer(wasmlib.IOTA, 2),
		Delay:      0,
	})
	ctx.Log("====  success ====")
	// TODO how to check if post was successful
}

func viewCheckContextFromViewEP(ctx *wasmlib.ScViewContext, params *ViewCheckContextFromViewEPParams) {
	ctx.Log("calling checkContextFromViewEP")

	ctx.Require(params.ChainId.Value().Equals(ctx.ContractId().ChainId()), "fail: chainID")
	ctx.Require(params.ChainOwnerId.Value().Equals(ctx.ChainOwnerId()), "fail: chainOwnerID")
	ctx.Require(params.ContractId.Value().Equals(ctx.ContractId()), "fail: contractID")
	ctx.Require(params.AgentId.Value().Equals(ctx.ContractId().AsAgentId()), "fail: agentID")
	ctx.Require(params.ContractCreator.Value().Equals(ctx.ContractCreator()), "fail: contractCreator")
}

func viewFibonacci(ctx *wasmlib.ScViewContext, params *ViewFibonacciParams) {
	ctx.Log("calling fibonacci")
	n := params.IntValue.Value()
	if n == 0 || n == 1 {
		ctx.Results().GetInt(ParamIntValue).SetValue(n)
		return
	}
	params1 := wasmlib.NewScMutableMap()
	params1.GetInt(ParamIntValue).SetValue(n - 1)
	results1 := ctx.CallSelf(HViewFibonacci, params1)
	n1 := results1.GetInt(ParamIntValue).Value()

	params2 := wasmlib.NewScMutableMap()
	params2.GetInt(ParamIntValue).SetValue(n - 2)
	results2 := ctx.CallSelf(HViewFibonacci, params2)
	n2 := results2.GetInt(ParamIntValue).Value()

	ctx.Results().GetInt(ParamIntValue).SetValue(n1 + n2)
}

func viewGetCounter(ctx *wasmlib.ScViewContext, params *ViewGetCounterParams) {
	ctx.Log("calling getCounter")
	counter := ctx.State().GetInt(VarCounter)
	ctx.Results().GetInt(VarCounter).SetValue(counter.Value())
}

func viewGetInt(ctx *wasmlib.ScViewContext, params *ViewGetIntParams) {
	ctx.Log("calling getInt")
	name := params.Name.Value()
	value := ctx.State().GetInt(wasmlib.Key(name))
	ctx.Require(value.Exists(), "param 'value' not found")
	ctx.Results().GetInt(wasmlib.Key(name)).SetValue(value.Value())
}

func viewJustView(ctx *wasmlib.ScViewContext, params *ViewJustViewParams) {
	ctx.Log("calling justView")
}

func viewPassTypesView(ctx *wasmlib.ScViewContext, params *ViewPassTypesViewParams) {
	ctx.Log("calling passTypesView")

	ctx.Require(params.Int64.Value() == 42, "int64 wrong")
	ctx.Require(params.Int64Zero.Value() == 0, "int64-0 wrong")
	ctx.Require(params.String.Value() == string(ParamString), "string wrong")
	ctx.Require(params.StringZero.Value() == "", "string-0 wrong")

	hash := ctx.Utility().HashBlake2b([]byte(ParamHash))
	ctx.Require(params.Hash.Value().Equals(hash), "Hash wrong")

	ctx.Require(params.Hname.Value().Equals(wasmlib.NewScHname("Hname")), "Hname wrong")
	ctx.Require(params.HnameZero.Value().Equals(wasmlib.ScHname(0)), "Hname-0 wrong")
}

func viewTestCallPanicViewEPFromView(ctx *wasmlib.ScViewContext, params *ViewTestCallPanicViewEPFromViewParams) {
	ctx.Log("calling testCallPanicViewEPFromView")
	ctx.CallSelf(HViewTestPanicViewEP, nil)
}

func viewTestChainOwnerIDView(ctx *wasmlib.ScViewContext, params *ViewTestChainOwnerIDViewParams) {
	ctx.Log("calling testChainOwnerIDView")
	ctx.Results().GetAgentId(ParamChainOwnerId).SetValue(ctx.ChainOwnerId())
}

func viewTestContractIDView(ctx *wasmlib.ScViewContext, params *ViewTestContractIDViewParams) {
	ctx.Log("calling testContractIDView")
	ctx.Results().GetContractId(ParamContractId).SetValue(ctx.ContractId())
}

func viewTestPanicViewEP(ctx *wasmlib.ScViewContext, params *ViewTestPanicViewEPParams) {
	ctx.Log("calling testPanicViewEP")
	ctx.Panic(MsgViewPanic)
}

func viewTestSandboxCall(ctx *wasmlib.ScViewContext, params *ViewTestSandboxCallParams) {
	ctx.Log("calling testSandboxCall")
	ret := ctx.Call(wasmlib.CoreRoot, wasmlib.CoreRootViewGetChainInfo, nil)
	desc := ret.GetString(wasmlib.Key("d")).Value()
	ctx.Results().GetString(wasmlib.Key("sandboxCall")).SetValue(desc)
}
