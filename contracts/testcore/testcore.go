// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package testcore

import "github.com/iotaledger/wasp/packages/vm/wasmlib"

const ParamIntParamName = wasmlib.Key("intParamName")
const ParamIntParamValue = wasmlib.Key("intParamValue")
const ParamHnameContract = wasmlib.Key("hnameContract")
const ParamHnameEp = wasmlib.Key("hnameEP")

const ParamAddress = wasmlib.Key("address")
const ParamChainOwnerId = wasmlib.Key("chainOwnerID")
const ParamContractId = wasmlib.Key("contractID")
const ParamChainId = wasmlib.Key("chainid")
const ParamCaller = wasmlib.Key("caller")
const ParamAgentId = wasmlib.Key("agentID")
const ParamCreator = wasmlib.Key("contractCreator")

const ParamInt64 = wasmlib.Key("int64")
const ParamInt64Zero = wasmlib.Key("int64-0")
const ParamHash = wasmlib.Key("Hash")
const ParamHname = wasmlib.Key("Hname")
const ParamHnameZero = wasmlib.Key("Hname-0")
const ParamString = wasmlib.Key("string")
const ParamStringZero = wasmlib.Key("string-0")

const VarCounter = wasmlib.Key("counter")
const VarContractNameDeployed = wasmlib.Key("exampleDeployTR")

const MsgFullPanic string = "========== panic FULL ENTRY POINT ========="
const MsgViewPanic string = "========== panic VIEW ========="
const MsgPanicUnauthorized string = "============== panic due to unauthorized call"

func OnLoad() {
	exports := wasmlib.NewScExports()
	exports.AddCall("init", onInit)
	exports.AddCall("doNothing", doNothing)
	exports.AddCall("callOnChain", callOnChain)
	exports.AddCall("setInt", setInt)
	exports.AddView("getInt", getInt)
	exports.AddView("fibonacci", fibonacci)
	exports.AddView("getCounter", getCounter)
	exports.AddCall("runRecursion", runRecursion)

	exports.AddCall("testPanicFullEP", testPanicFullEp)
	exports.AddView("testPanicViewEP", testPanicViewEp)
	exports.AddCall("testCallPanicFullEP", testCallPanicFullEp)
	exports.AddCall("testCallPanicViewEPFromFull", testCallPanicViewFromFull)
	exports.AddView("testCallPanicViewEPFromView", testCallPanicViewFromView)

	exports.AddView("testChainOwnerIDView", testChainOwnerIdView)
	exports.AddCall("testChainOwnerIDFull", testChainOwnerIdFull)
	exports.AddView("testContractIDView", testContractIdView)
	exports.AddCall("testContractIDFull", testContractIdFull)
	exports.AddView("testSandboxCall", testSandboxCall)

	exports.AddCall("passTypesFull", passTypesFull)
	exports.AddView("passTypesView", passTypesView)
	exports.AddCall("checkContextFromFullEP", checkCtxFromFull)
	exports.AddView("checkContextFromViewEP", checkCtxFromView)

	exports.AddCall("sendToAddress", sendToAddress)
	exports.AddView("justView", testJustView)

	exports.AddCall("testEventLogGenericData", testEventLogGenericData)
	exports.AddCall("testEventLogEventData", testEventLogEventData)
	exports.AddCall("testEventLogDeploy", testEventLogDeploy)

	exports.AddCall("withdrawToChain", withdrawToChain)
}

func onInit(ctx *wasmlib.ScCallContext) {
	ctx.Log("testcore.on_init.wasm.begin")
}

func doNothing(ctx *wasmlib.ScCallContext) {
	ctx.Log("testcore.do_nothing.begin")
}

func setInt(ctx *wasmlib.ScCallContext) {
	ctx.Log("testcore.set_int.begin")
	paramName := ctx.Params().GetString(ParamIntParamName)
	ctx.Require(paramName.Exists(), "param 'name' not found")

	paramValue := ctx.Params().GetInt(ParamIntParamValue)
	ctx.Require(paramValue.Exists(), "param 'value' not found")

	ctx.State().GetInt(wasmlib.Key(paramName.Value())).SetValue(paramValue.Value())
}

func getInt(ctx *wasmlib.ScViewContext) {
	ctx.Log("testcore.get_int.begin")
	paramName := ctx.Params().GetString(ParamIntParamName)
	ctx.Require(paramName.Exists(), "param 'name' not found")

	paramValue := ctx.State().GetInt(wasmlib.Key(paramName.Value()))
	ctx.Require(paramValue.Exists(), "param 'value' not found")

	ctx.Results().GetInt(wasmlib.Key(paramName.Value())).SetValue(paramValue.Value())
}

func callOnChain(ctx *wasmlib.ScCallContext) {
	paramValue := ctx.Params().GetInt(ParamIntParamValue)
	ctx.Require(paramValue.Exists(), "param 'value' not found")
	paramIn := paramValue.Value()

	targetContract := ctx.ContractId().Hname()
	paramHnameContract := ctx.Params().GetHname(ParamHnameContract)
	if paramHnameContract.Exists() {
		targetContract = paramHnameContract.Value()
	}

	targetEp := wasmlib.NewScHname("callOnChain")
	paramHnameEp := ctx.Params().GetHname(ParamHnameEp)
	if paramHnameEp.Exists() {
		targetEp = paramHnameEp.Value()
	}

	varCounter := ctx.State().GetInt(VarCounter)
	counter := int64(0)
	if varCounter.Exists() {
		counter = varCounter.Value()
	}
	varCounter.SetValue(counter + 1)

	msg := "call depth = " + ctx.Utility().String(paramIn) +
		" hnameContract = " + targetContract.String() +
		" hnameEP = " + targetEp.String() +
		" counter = " + ctx.Utility().String(counter)
	ctx.Log(msg)

	par := wasmlib.NewScMutableMap()
	par.GetInt(ParamIntParamValue).SetValue(paramIn)
	ret := ctx.Call(targetContract, targetEp, par, nil)

	retVal := ret.GetInt(ParamIntParamValue)

	ctx.Results().GetInt(ParamIntParamValue).SetValue(retVal.Value())
}

func getCounter(ctx *wasmlib.ScViewContext) {
	ctx.Log("testcore.get_counter.begin")
	counter := ctx.State().GetInt(VarCounter)
	ctx.Results().GetInt(VarCounter).SetValue(counter.Value())
}

func runRecursion(ctx *wasmlib.ScCallContext) {
	paramValue := ctx.Params().GetInt(ParamIntParamValue)
	ctx.Require(paramValue.Exists(), "param no found")
	depth := paramValue.Value()
	if depth <= 0 {
		return
	}
	par := wasmlib.NewScMutableMap()
	par.GetInt(ParamIntParamValue).SetValue(depth - 1)
	par.GetHname(ParamHnameEp).SetValue(wasmlib.NewScHname("runRecursion"))
	ctx.Call(ctx.ContractId().Hname(), wasmlib.NewScHname("callOnChain"), par, nil)
	// TODO how would I return result of the call ???
	ctx.Results().GetInt(ParamIntParamValue).SetValue(depth - 1)
}

func fibonacci(ctx *wasmlib.ScViewContext) {
	nParam := ctx.Params().GetInt(ParamIntParamValue)
	ctx.Require(nParam.Exists(), "param 'value' not found")

	n := nParam.Value()
	if n == 0 || n == 1 {
		ctx.Results().GetInt(ParamIntParamValue).SetValue(n)
		return
	}
	params1 := wasmlib.NewScMutableMap()
	params1.GetInt(ParamIntParamValue).SetValue(n - 1)
	results1 := ctx.Call(ctx.ContractId().Hname(), wasmlib.NewScHname("fibonacci"), params1)
	n1 := results1.GetInt(ParamIntParamValue).Value()

	params2 := wasmlib.NewScMutableMap()
	params2.GetInt(ParamIntParamValue).SetValue(n - 2)
	results2 := ctx.Call(ctx.ContractId().Hname(), wasmlib.NewScHname("fibonacci"), params2)
	n2 := results2.GetInt(ParamIntParamValue).Value()

	ctx.Results().GetInt(ParamIntParamValue).SetValue(n1 + n2)
}

func testPanicFullEp(ctx *wasmlib.ScCallContext) {
	ctx.Panic(MsgFullPanic)
}

func testPanicViewEp(ctx *wasmlib.ScViewContext) {
	ctx.Panic(MsgViewPanic)
}

func testCallPanicFullEp(ctx *wasmlib.ScCallContext) {
	ctx.Call(ctx.ContractId().Hname(), wasmlib.NewScHname("testPanicFullEP"), nil, nil)
}

func testCallPanicViewFromFull(ctx *wasmlib.ScCallContext) {
	ctx.Call(ctx.ContractId().Hname(), wasmlib.NewScHname("testPanicViewEP"), nil, nil)
}

func testCallPanicViewFromView(ctx *wasmlib.ScViewContext) {
	ctx.Call(ctx.ContractId().Hname(), wasmlib.NewScHname("testPanicViewEP"), nil)
}

func testJustView(ctx *wasmlib.ScViewContext) {
	ctx.Log("calling empty view entry point")
}

func sendToAddress(ctx *wasmlib.ScCallContext) {
	ctx.Log("sendToAddress")
	ctx.Require(ctx.Caller().Equals(ctx.ContractCreator()), MsgPanicUnauthorized)

	targetAddr := ctx.Params().GetAddress(ParamAddress)
	ctx.Require(targetAddr.Exists(), "parameter 'address' not found")

	myBalances := ctx.Balances()
	ctx.TransferToAddress(targetAddr.Value(), myBalances)
}

func testChainOwnerIdView(ctx *wasmlib.ScViewContext) {
	ctx.Results().GetAgentId(ParamChainOwnerId).SetValue(ctx.ChainOwnerId())
}

func testChainOwnerIdFull(ctx *wasmlib.ScCallContext) {
	ctx.Results().GetAgentId(ParamChainOwnerId).SetValue(ctx.ChainOwnerId())
}

func testContractIdView(ctx *wasmlib.ScViewContext) {
	ctx.Results().GetContractId(ParamContractId).SetValue(ctx.ContractId())
}

func testContractIdFull(ctx *wasmlib.ScCallContext) {
	ctx.Results().GetContractId(ParamContractId).SetValue(ctx.ContractId())
}

func testSandboxCall(ctx *wasmlib.ScViewContext) {
	ret := ctx.Call(wasmlib.CoreRoot, wasmlib.CoreRootViewGetChainInfo, nil)
	desc := ret.GetString(wasmlib.Key("d")).Value()
	ctx.Results().GetString(wasmlib.Key("sandboxCall")).SetValue(desc)
}

func passTypesFull(ctx *wasmlib.ScCallContext) {
	ctx.Require(ctx.Params().GetInt(ParamInt64).Exists(), "!int64.exist")
	ctx.Require(ctx.Params().GetInt(ParamInt64).Value() == 42, "int64 wrong")

	ctx.Require(ctx.Params().GetInt(ParamInt64Zero).Exists(), "!int64-0.exist")
	ctx.Require(ctx.Params().GetInt(ParamInt64Zero).Value() == 0, "int64-0 wrong")

	ctx.Require(ctx.Params().GetString(ParamString).Exists(), "!string.exist")
	ctx.Require(ctx.Params().GetString(ParamString).Value() == "string", "string wrong")

	ctx.Require(ctx.Params().GetString(ParamStringZero).Exists(), "!string-0.exist")
	ctx.Require(ctx.Params().GetString(ParamStringZero).Value() == "", "string-0 wrong")

	ctx.Require(ctx.Params().GetHash(ParamHash).Exists(), "!Hash.exist")

	hash := ctx.Utility().HashBlake2b([]byte("Hash"))
	ctx.Require(ctx.Params().GetHash(ParamHash).Value().Equals(hash), "Hash wrong")

	ctx.Require(ctx.Params().GetHname(ParamHname).Exists(), "!Hname.exist")
	ctx.Require(ctx.Params().GetHname(ParamHname).Value().Equals(wasmlib.NewScHname("Hname")), "Hname wrong")

	ctx.Require(ctx.Params().GetHname(ParamHnameZero).Exists(), "!Hname-0.exist")
	ctx.Require(ctx.Params().GetHname(ParamHnameZero).Value().Equals(wasmlib.ScHname(0)), "Hname-0 wrong")
}

func passTypesView(ctx *wasmlib.ScViewContext) {
	ctx.Require(ctx.Params().GetInt(ParamInt64).Exists(), "!int64.exist")
	ctx.Require(ctx.Params().GetInt(ParamInt64).Value() == 42, "int64 wrong")

	ctx.Require(ctx.Params().GetInt(ParamInt64Zero).Exists(), "!int64-0.exist")
	ctx.Require(ctx.Params().GetInt(ParamInt64Zero).Value() == 0, "int64-0 wrong")

	ctx.Require(ctx.Params().GetString(ParamString).Exists(), "!string.exist")
	ctx.Require(ctx.Params().GetString(ParamString).Value() == "string", "string wrong")

	ctx.Require(ctx.Params().GetString(ParamStringZero).Exists(), "!string-0.exist")
	ctx.Require(ctx.Params().GetString(ParamStringZero).Value() == "", "string-0 wrong")

	ctx.Require(ctx.Params().GetHash(ParamHash).Exists(), "!Hash.exist")

	hash := ctx.Utility().HashBlake2b([]byte("Hash"))
	ctx.Require(ctx.Params().GetHash(ParamHash).Value().Equals(hash), "Hash wrong")

	ctx.Require(ctx.Params().GetHname(ParamHname).Exists(), "!Hname.exist")
	ctx.Require(ctx.Params().GetHname(ParamHname).Value().Equals(wasmlib.NewScHname("Hname")), "Hname wrong")

	ctx.Require(ctx.Params().GetHname(ParamHnameZero).Exists(), "!Hname-0.exist")
	ctx.Require(ctx.Params().GetHname(ParamHnameZero).Value().Equals(wasmlib.ScHname(0)), "Hname-0 wrong")
}

func checkCtxFromFull(ctx *wasmlib.ScCallContext) {
	par := ctx.Params()

	chainId := par.GetChainId(ParamChainId)
	ctx.Require(chainId.Exists() && chainId.Value().Equals(ctx.ContractId().ChainId()), "fail: chainID")

	chainOwnerId := par.GetAgentId(ParamChainOwnerId)
	ctx.Require(chainOwnerId.Exists() && chainOwnerId.Value().Equals(ctx.ChainOwnerId()), "fail: chainOwnerID")

	caller := par.GetAgentId(ParamCaller)
	ctx.Require(caller.Exists() && caller.Value().Equals(ctx.Caller()), "fail: caller")

	contractId := par.GetContractId(ParamContractId)
	ctx.Require(contractId.Exists() && contractId.Value().Equals(ctx.ContractId()), "fail: contractID")

	agentId := par.GetAgentId(ParamAgentId)
	asAgentId := ctx.ContractId().AsAgentId()
	ctx.Require(agentId.Exists() && agentId.Value().Equals(asAgentId), "fail: agentID")

	creator := par.GetAgentId(ParamCreator)
	ctx.Require(creator.Exists() && creator.Value().Equals(ctx.ContractCreator()), "fail: contractCreator")
}

func checkCtxFromView(ctx *wasmlib.ScViewContext) {
	par := ctx.Params()

	chainId := par.GetChainId(ParamChainId)
	ctx.Require(chainId.Exists() && chainId.Value().Equals(ctx.ContractId().ChainId()), "fail: chainID")

	chainOwnerId := par.GetAgentId(ParamChainOwnerId)
	ctx.Require(chainOwnerId.Exists() && chainOwnerId.Value().Equals(ctx.ChainOwnerId()), "fail: chainOwnerID")

	contractId := par.GetContractId(ParamContractId)
	ctx.Require(contractId.Exists() && contractId.Value().Equals(ctx.ContractId()), "fail: contractID")

	agentId := par.GetAgentId(ParamAgentId)
	asAgentId := ctx.ContractId().AsAgentId()
	ctx.Require(agentId.Exists() && agentId.Value().Equals(asAgentId), "fail: agentID")

	creator := par.GetAgentId(ParamCreator)
	ctx.Require(creator.Exists() && creator.Value().Equals(ctx.ContractCreator()), "fail: contractCreator")
}

func testEventLogGenericData(ctx *wasmlib.ScCallContext) {
	counter := ctx.Params().GetInt(VarCounter)
	ctx.Require(counter.Exists(), "!counter.exist")
	event := "[GenericData] Counter Number: " + counter.String()
	ctx.Event(event)
}

func testEventLogEventData(ctx *wasmlib.ScCallContext) {
	ctx.Event("[Event] - Testing Event...")
}

func testEventLogDeploy(ctx *wasmlib.ScCallContext) {
	//Deploy the same contract with another name
	programHash := ctx.Utility().HashBlake2b([]byte("test_sandbox"))
	ctx.Deploy(programHash, string(VarContractNameDeployed),
		"test contract deploy log", nil)
}

func withdrawToChain(ctx *wasmlib.ScCallContext) {
	//Deploy the same contract with another name
	targetChain := ctx.Params().GetChainId(ParamChainId)
	ctx.Require(targetChain.Exists(), "chainID not provided")

	targetContractId := wasmlib.NewScContractId(targetChain.Value(), wasmlib.CoreAccounts)
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
