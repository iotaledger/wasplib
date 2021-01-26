// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package testcore

import "github.com/iotaledger/wasplib/client"

const ParamIntParamName = client.Key("intParamName")
const ParamIntParamValue = client.Key("intParamValue")
const ParamHnameContract = client.Key("hnameContract")
const ParamHnameEp = client.Key("hnameEP")

const ParamAddress = client.Key("address")
const ParamChainOwner = client.Key("chainOwner")
const ParamContractId = client.Key("contractID")

const ParamInt64 = client.Key("int64")
const ParamInt64Zero = client.Key("int64-0")
const ParamHash = client.Key("Hash")
const ParamHname = client.Key("Hname")
const ParamHnameZero = client.Key("Hname-0")
const ParamString = client.Key("string")
const ParamStringZero = client.Key("string-0")

const VarCounter = client.Key("counter")

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

	exports.AddCall("sendToAddress", sendToAddress)
	exports.AddView("justView", testJustView)
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

	targetContract := client.Hname(0) // Todo strange a bit: client.Hname(0) is a constant

	paramHnameContract := ctx.Params().GetHname(ParamHnameContract)
	if paramHnameContract.Exists() {
		targetContract = paramHnameContract.Value()
	}

	targetEp := client.NewHname("callOnChain")
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
	par.GetHname(ParamHnameEp).SetValue(client.NewHname("runRecursion"))
	ctx.Call(client.Hname(0), client.NewHname("callOnChain"), par, nil)
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
	results1 := ctx.Call(client.Hname(0), client.NewHname("fibonacci"), params1)
	n1 := results1.GetInt(ParamIntParamValue).Value()

	params2 := client.NewScMutableMap()
	params2.GetInt(ParamIntParamValue).SetValue(n - 2)
	results2 := ctx.Call(client.Hname(0), client.NewHname("fibonacci"), params2)
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
	ctx.Call(client.Hname(0), client.NewHname("testPanicFullEP"), nil, nil)
}

func testCallPanicViewFromFull(ctx *client.ScCallContext) {
	ctx.Call(client.Hname(0), client.NewHname("testPanicViewEP"), nil, nil)
}

func testCallPanicViewFromView(ctx *client.ScViewContext) {
	ctx.Call(client.Hname(0), client.NewHname("testPanicViewEP"), nil)
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
	ctx.Results().GetAgent(ParamChainOwner).SetValue(ctx.ChainOwner())
}

func testChainOwnerIdFull(ctx *client.ScCallContext) {
	ctx.Results().GetAgent(ParamChainOwner).SetValue(ctx.ChainOwner())
}

func testContractIdView(ctx *client.ScViewContext) {
	//TODO discussion about using ChainID vs ContractID because one of those seems redundant
	ctx.Results().GetContractId(ParamContractId).SetValue(ctx.ContractId())
}

func testContractIdFull(ctx *client.ScCallContext) {
	ctx.Results().GetContractId(ParamContractId).SetValue(ctx.ContractId())
}

func testSandboxCall(ctx *client.ScViewContext) {
	ret := ctx.Call(client.CoreRoot, client.ViewGetChainInfo, nil)
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

	hash := ctx.Utility().Hash([]byte("Hash"))
	ctx.Require(ctx.Params().GetHash(ParamHash).Value().Equals(hash), "Hash wrong")

	ctx.Require(ctx.Params().GetHname(ParamHname).Exists(), "!Hname.exist")
	ctx.Require(ctx.Params().GetHname(ParamHname).Value().Equals(client.NewHname("Hname")), "Hname wrong")

	ctx.Require(ctx.Params().GetHname(ParamHnameZero).Exists(), "!Hname-0.exist")
	ctx.Require(ctx.Params().GetHname(ParamHnameZero).Value().Equals(client.Hname(0)), "Hname-0 wrong")
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

	hash := ctx.Utility().Hash([]byte("Hash"))
	ctx.Require(ctx.Params().GetHash(ParamHash).Value().Equals(hash), "Hash wrong")

	ctx.Require(ctx.Params().GetHname(ParamHname).Exists(), "!Hname.exist")
	ctx.Require(ctx.Params().GetHname(ParamHname).Value().Equals(client.NewHname("Hname")), "Hname wrong")

	ctx.Require(ctx.Params().GetHname(ParamHnameZero).Exists(), "!Hname-0.exist")
	ctx.Require(ctx.Params().GetHname(ParamHnameZero).Value().Equals(client.Hname(0)), "Hname-0 wrong")
}
