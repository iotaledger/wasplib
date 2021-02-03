// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package testcore

import "github.com/iotaledger/wasplib/client"

const ParamIntParamName = client.Key("intParamName")
const ParamIntParamValue = client.Key("intParamValue")
const ParamHnameContract = client.Key("hnameContract")
const ParamHnameEp = client.Key("hnameEP")

const ParamAddress = client.Key("address")
const ParamChainOwnerId = client.Key("chainOwnerID")
const ParamContractId = client.Key("contractID")
const ParamChainId = client.Key("chainid")
const ParamCaller = client.Key("caller")
const ParamAgentId = client.Key("agentID")
const ParamCreator = client.Key("contractCreator")

const ParamInt64 = client.Key("int64")
const ParamInt64Zero = client.Key("int64-0")
const ParamHash = client.Key("Hash")
const ParamHname = client.Key("Hname")
const ParamHnameZero = client.Key("Hname-0")
const ParamString = client.Key("string")
const ParamStringZero = client.Key("string-0")

const VarCounter = client.Key("counter")
const VarContractNameDeployed = client.Key("exampleDeployTR")

const MsgFullPanic string = "========== panic FULL ENTRY POINT ========="
const MsgViewPanic string = "========== panic VIEW ========="
const MsgPanicUnauthorized string = "============== panic due to unauthorized call"

func OnLoad() {
	exports := client.NewScExports()
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

func onInit(ctx *client.ScCallContext) {
	ctx.Log("testcore.on_init.wasm.begin")
}

func doNothing(ctx *client.ScCallContext) {
	ctx.Log("testcore.do_nothing.begin")
}

func setInt(ctx *client.ScCallContext) {
	ctx.Log("testcore.set_int.begin")
	paramName := ctx.Params().GetString(ParamIntParamName)
	ctx.Require(paramName.Exists(), "param 'name' not found")

	paramValue := ctx.Params().GetInt(ParamIntParamValue)
	ctx.Require(paramValue.Exists(), "param 'value' not found")

	ctx.State().GetInt(client.Key(paramName.Value())).SetValue(paramValue.Value())
}

func getInt(ctx *client.ScViewContext) {
	ctx.Log("testcore.get_int.begin")
	paramName := ctx.Params().GetString(ParamIntParamName)
	ctx.Require(paramName.Exists(), "param 'name' not found")

	paramValue := ctx.State().GetInt(client.Key(paramName.Value()))
	ctx.Require(paramValue.Exists(), "param 'value' not found")

	ctx.Results().GetInt(client.Key(paramName.Value())).SetValue(paramValue.Value())
}

func callOnChain(ctx *client.ScCallContext) {
	paramValue := ctx.Params().GetInt(ParamIntParamValue)
	ctx.Require(paramValue.Exists(), "param 'value' not found")
	paramIn := paramValue.Value()

	targetContract := ctx.ContractId().Hname()
	paramHnameContract := ctx.Params().GetHname(ParamHnameContract)
	if paramHnameContract.Exists() {
		targetContract = paramHnameContract.Value()
	}

	targetEp := client.NewScHname("callOnChain")
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

	par := client.NewScMutableMap()
	par.GetInt(ParamIntParamValue).SetValue(paramIn)
	ret := ctx.Call(targetContract, targetEp, par, nil)

	retVal := ret.GetInt(ParamIntParamValue)

	ctx.Results().GetInt(ParamIntParamValue).SetValue(retVal.Value())
}

func getCounter(ctx *client.ScViewContext) {
	ctx.Log("testcore.get_counter.begin")
	counter := ctx.State().GetInt(VarCounter)
	ctx.Results().GetInt(VarCounter).SetValue(counter.Value())
}

func runRecursion(ctx *client.ScCallContext) {
	paramValue := ctx.Params().GetInt(ParamIntParamValue)
	ctx.Require(paramValue.Exists(), "param no found")
	depth := paramValue.Value()
	if depth <= 0 {
		return
	}
	par := client.NewScMutableMap()
	par.GetInt(ParamIntParamValue).SetValue(depth - 1)
	par.GetHname(ParamHnameEp).SetValue(client.NewScHname("runRecursion"))
	ctx.Call(ctx.ContractId().Hname(), client.NewScHname("callOnChain"), par, nil)
	// TODO how would I return result of the call ???
	ctx.Results().GetInt(ParamIntParamValue).SetValue(depth - 1)
}

func fibonacci(ctx *client.ScViewContext) {
	nParam := ctx.Params().GetInt(ParamIntParamValue)
	ctx.Require(nParam.Exists(), "param 'value' not found")

	n := nParam.Value()
	if n == 0 || n == 1 {
		ctx.Results().GetInt(ParamIntParamValue).SetValue(n)
		return
	}
	params1 := client.NewScMutableMap()
	params1.GetInt(ParamIntParamValue).SetValue(n - 1)
	results1 := ctx.Call(ctx.ContractId().Hname(), client.NewScHname("fibonacci"), params1)
	n1 := results1.GetInt(ParamIntParamValue).Value()

	params2 := client.NewScMutableMap()
	params2.GetInt(ParamIntParamValue).SetValue(n - 2)
	results2 := ctx.Call(ctx.ContractId().Hname(), client.NewScHname("fibonacci"), params2)
	n2 := results2.GetInt(ParamIntParamValue).Value()

	ctx.Results().GetInt(ParamIntParamValue).SetValue(n1 + n2)
}

func testPanicFullEp(ctx *client.ScCallContext) {
	ctx.Panic(MsgFullPanic)
}

func testPanicViewEp(ctx *client.ScViewContext) {
	ctx.Panic(MsgViewPanic)
}

func testCallPanicFullEp(ctx *client.ScCallContext) {
	ctx.Call(ctx.ContractId().Hname(), client.NewScHname("testPanicFullEP"), nil, nil)
}

func testCallPanicViewFromFull(ctx *client.ScCallContext) {
	ctx.Call(ctx.ContractId().Hname(), client.NewScHname("testPanicViewEP"), nil, nil)
}

func testCallPanicViewFromView(ctx *client.ScViewContext) {
	ctx.Call(ctx.ContractId().Hname(), client.NewScHname("testPanicViewEP"), nil)
}

func testJustView(ctx *client.ScViewContext) {
	ctx.Log("calling empty view entry point")
}

func sendToAddress(ctx *client.ScCallContext) {
	ctx.Log("sendToAddress")
	ctx.Require(ctx.Caller().Equals(ctx.ContractCreator()), MsgPanicUnauthorized)

	targetAddr := ctx.Params().GetAddress(ParamAddress)
	ctx.Require(targetAddr.Exists(), "parameter 'address' not found")

	myBalances := ctx.Balances()
	ctx.TransferToAddress(targetAddr.Value(), myBalances)
}

func testChainOwnerIdView(ctx *client.ScViewContext) {
	ctx.Results().GetAgentId(ParamChainOwnerId).SetValue(ctx.ChainOwnerId())
}

func testChainOwnerIdFull(ctx *client.ScCallContext) {
	ctx.Results().GetAgentId(ParamChainOwnerId).SetValue(ctx.ChainOwnerId())
}

func testContractIdView(ctx *client.ScViewContext) {
	ctx.Results().GetContractId(ParamContractId).SetValue(ctx.ContractId())
}

func testContractIdFull(ctx *client.ScCallContext) {
	ctx.Results().GetContractId(ParamContractId).SetValue(ctx.ContractId())
}

func testSandboxCall(ctx *client.ScViewContext) {
	ret := ctx.Call(client.CoreRoot, client.CoreRootViewGetChainInfo, nil)
	desc := ret.GetString(client.Key("d")).Value()
	ctx.Results().GetString(client.Key("sandboxCall")).SetValue(desc)
}

func passTypesFull(ctx *client.ScCallContext) {
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
	ctx.Require(ctx.Params().GetHname(ParamHname).Value().Equals(client.NewScHname("Hname")), "Hname wrong")

	ctx.Require(ctx.Params().GetHname(ParamHnameZero).Exists(), "!Hname-0.exist")
	ctx.Require(ctx.Params().GetHname(ParamHnameZero).Value().Equals(client.ScHname(0)), "Hname-0 wrong")
}

func passTypesView(ctx *client.ScViewContext) {
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
	ctx.Require(ctx.Params().GetHname(ParamHname).Value().Equals(client.NewScHname("Hname")), "Hname wrong")

	ctx.Require(ctx.Params().GetHname(ParamHnameZero).Exists(), "!Hname-0.exist")
	ctx.Require(ctx.Params().GetHname(ParamHnameZero).Value().Equals(client.ScHname(0)), "Hname-0 wrong")
}

func checkCtxFromFull(ctx *client.ScCallContext) {
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

func checkCtxFromView(ctx *client.ScViewContext) {
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

func testEventLogGenericData(ctx *client.ScCallContext) {
	counter := ctx.Params().GetInt(VarCounter)
	ctx.Require(counter.Exists(), "!counter.exist")
	event := "[GenericData] Counter Number: " + counter.String()
	ctx.Event(event)
}

func testEventLogEventData(ctx *client.ScCallContext) {
	ctx.Event("[Event] - Testing Event...")
}

func testEventLogDeploy(ctx *client.ScCallContext) {
	//Deploy the same contract with another name
	programHash := ctx.Utility().HashBlake2b([]byte("test_sandbox"))
	ctx.Deploy(programHash, string(VarContractNameDeployed),
		"test contract deploy log", nil)
}

func withdrawToChain(ctx *client.ScCallContext) {
	//Deploy the same contract with another name
	targetChain := ctx.Params().GetChainId(ParamChainId)
	ctx.Require(targetChain.Exists(), "chainID not provided")

	targetContractId := client.NewScContractId(targetChain.Value(), client.CoreAccounts)
	ctx.Post(&client.PostRequestParams{
		ContractId: targetContractId,
		Function:   client.CoreAccountsViewWithdrawToChain,
		Params:     nil,
		Transfer:   client.NewScTransfer(client.IOTA, 2),
		Delay:      0,
	})
	ctx.Log("====  success ====")
	// TODO how to check if post was successful
}
