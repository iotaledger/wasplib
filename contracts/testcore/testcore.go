// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package testcore

import (
	"github.com/iotaledger/wasplib/client"
)

const ParamIntParamName = client.Key("intParamName")
const ParamIntParamValue = client.Key("intParamValue")
const ParamInt64 = client.Key("int64")
const ParamInt64_0 = client.Key("int64-0")
const Paramhname = client.Key("hname")
const ParamHname = client.Key("Hname")
const ParamHname0 = client.Key("Hname-0")
const ParamCallOption = client.Key("callOption")
const ParamAddress = client.Key("address")
const ParamChainOwner = client.Key("chainOwner")
const ParamContractId = client.Key("contractID")

const VarCounter = client.Key("counter")

const MsgFullPanic string = "========== panic FULL ENTRY POINT ========="
const MsgViewPanic string = "========== panic VIEW ========="
const MsgPanicUnauthorized string = "============== panic due to unauthorized call"

const CallOptionForward = client.Key("forward")

func OnLoad() {
	exports := client.NewScExports()
	exports.AddCall("init", onInit)
	exports.AddCall("doNothing", doNothing)
	exports.AddCall("callOnChain", callOnChain)
	exports.AddCall("setInt", setInt)
	exports.AddView("getInt", getInt)
	exports.AddView("fibonacci", fibonacci)

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
	if !paramName.Exists() {
		ctx.Panic("param name not found")
	}
	paramValue := ctx.Params().GetInt(ParamIntParamValue)
	if !paramValue.Exists() {
		ctx.Panic("param value not found")
	}
	ctx.State().GetInt(client.Key(paramName.Value())).SetValue(paramValue.Value())
}

func getInt(ctx *client.ScViewContext) {
	ctx.Log("testcore.get_int.begin")
	paramName := ctx.Params().GetString(ParamIntParamName)
	if !paramName.Exists() {
		ctx.Panic("param name not found")
	}
	paramValue := ctx.State().GetInt(client.Key(paramName.Value()))
	if !paramValue.Exists() {
		ctx.Panic("param value is not in state")
	}
	ctx.Results().GetInt(client.Key(paramName.Value())).SetValue(paramValue.Value())
}

func callOnChain(ctx *client.ScCallContext) {
	paramCallOption := ctx.Params().GetString(ParamCallOption)
	if !paramCallOption.Exists() {
		ctx.Panic("'callOption' not specified")
	}
	callOption := paramCallOption.Value()

	paramValue := ctx.Params().GetInt(ParamIntParamValue)
	if !paramValue.Exists() {
		ctx.Panic("param value not found")
	}
	callDepth := paramValue.Value()

	target := client.Hname(0)
	paramHname := ctx.Params().GetHname(Paramhname)
	if paramHname.Exists() {
		target = paramHname.Value()
	}

	varCounter := ctx.State().GetInt(VarCounter)
	counter := int64(0)
	if varCounter.Exists() {
		counter = varCounter.Value()
	}

	// TODO ctx.contract_id() ContactID is not an AgentID type.
	//  should be

	//ctx.Log(fmt.Sprintf("call depth = %d option = '%s' hname = %s counter = %d",
	//	callDepth, callOption, target.String(), counter))

	if callDepth <= 0 {
		ctx.Results().GetInt(VarCounter).SetValue(varCounter.Value())
		return
	}

    varCounter.SetValue(counter + 1)
    callDepth = callDepth - 1
	if callOption == string(CallOptionForward) {
		par := client.NewScMutableMap()
		par.GetString(ParamCallOption).SetValue(callOption)
		par.GetInt(ParamIntParamValue).SetValue(callDepth)
		ret := ctx.Call(target, client.NewHname("callOnChain"), par, nil)
		ctx.Results().GetInt(VarCounter).SetValue(ret.GetInt(VarCounter).Value())
	} else {
		ctx.Panic("unknown call option")
	}
}

func fibonacci(ctx *client.ScViewContext) {
	nParam := ctx.Params().GetInt(ParamIntParamValue)
	if !nParam.Exists() {
		ctx.Panic("param value not found")
	}
	n := nParam.Value()
	// ctx.log(&("fibonacci: ".to_string() + &n.to_string()));
	if n == 0 || n == 1 {
		ctx.Results().GetInt(ParamIntParamValue).SetValue(n)
		return
	}
	params1 := client.NewScMutableMap()
	params1.GetInt(ParamIntParamValue).SetValue(n - 1)
	results1 := ctx.Call(client.Hname(0), client.NewHname("fibonacci"), params1)
	n1 := results1.GetInt(ParamIntParamValue).Value()
	// ctx.log(&("    fibonacci-1: ".to_string() + &n1.to_string()));

	params2 := client.NewScMutableMap()
	params2.GetInt(ParamIntParamValue).SetValue(n - 2)
	results2 := ctx.Call(client.Hname(0), client.NewHname("fibonacci"), params2)
	n2 := results2.GetInt(ParamIntParamValue).Value()
	// ctx.log(&("    fibonacci-2: ".to_string() + &n2.to_string()));

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

// FIXME no need for 'view method special'
func testCallPanicViewFromFull(ctx *client.ScCallContext) {
	ctx.Call(client.Hname(0), client.NewHname("testPanicViewEP"), nil, nil)
}

// FIXME no need for 'view method special'
func testCallPanicViewFromView(ctx *client.ScViewContext) {
	ctx.Call(client.Hname(0), client.NewHname("testPanicViewEP"), nil)
}

func sendToAddress(ctx *client.ScCallContext) {
	ctx.Log("sendToAddress")
	if !ctx.Caller().Equals(ctx.ContractCreator()) {
		ctx.Panic(MsgPanicUnauthorized)
	}
	targetAddr := ctx.Params().GetAddress(ParamAddress)
	if !targetAddr.Exists() {
		ctx.Panic("parameter 'address' not provided")
	}
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
	ctx.Results().GetAgent(ParamContractId).SetValue(ctx.ContractId());
	// alternatively do not use agent but bytes instead for now:
	// ctx.Results().GetBytes(PARAM_CONTRACT_ID).SetValue(ctx.ContractId().Bytes());
}

func testContractIdFull(ctx *client.ScCallContext) {
	ctx.Results().GetAgent(ParamContractId).SetValue(ctx.ContractId());
	// alternatively do not use agent but bytes instead for now:
	// ctx.Results().GetBytes(PARAM_CONTRACT_ID).SetValue(ctx.ContractId().Bytes());
}

func testSandboxCall(ctx *client.ScViewContext) {
	ret := ctx.Call(client.CoreRoot, client.ViewGetChainInfo, nil)
	desc := ret.GetString(client.Key("d")).Value()
	ctx.Results().GetString(client.Key("sandboxCall")).SetValue(desc)
}

func passTypesFull(ctx *client.ScCallContext) {
	if !ctx.Params().GetInt(ParamInt64).Exists() {
		ctx.Panic("!int64.exist")
	}
	if ctx.Params().GetInt(ParamInt64).Value() != 42 {
		ctx.Panic("int64 wrong")
	}
	if !ctx.Params().GetInt(ParamInt64_0).Exists() {
		ctx.Panic("!int64-0.exist")
	}
	if ctx.Params().GetInt(ParamInt64_0).Value() != 0 {
		ctx.Panic("int64-0 wrong")
	}
	if !ctx.Params().GetHash(client.Key("Hash")).Exists() {
		ctx.Panic("!Hash.exist")
	}
	hash := ctx.Utility().Hash([]byte("Hash"))
	if !ctx.Params().GetHash(client.Key("Hash")).Value().Equals(hash) {
		ctx.Panic("Hash wrong")
	}
	if !ctx.Params().GetHname(ParamHname).Exists() {
		ctx.Panic("!Hname. exist")
	}
	if !ctx.Params().GetHname(ParamHname).Value().Equals(client.NewHname(string(ParamHname))) {
		ctx.Panic("Hname wrong")
	}
	if !ctx.Params().GetHname(ParamHname0).Exists() {
		ctx.Panic("!Hname-0.exist")
	}
	if !ctx.Params().GetHname(ParamHname0).Value().Equals(client.Hname(0)) {
		ctx.Panic("Hname-0 wrong")
	}
}

func passTypesView(ctx *client.ScViewContext) {
	if !ctx.Params().GetInt(ParamInt64).Exists() {
		ctx.Panic("!int64. exist")
	}
	if ctx.Params().GetInt(ParamInt64).Value() != 42 {
		ctx.Panic("int64 wrong")
	}
	if !ctx.Params().GetInt(ParamInt64_0).Exists() {
		ctx.Panic("!int64-0. exist")
	}
	if ctx.Params().GetInt(ParamInt64_0).Value() != 0 {
		ctx.Panic("int64-0 wrong")
	}
	if !ctx.Params().GetHash(client.Key("Hash")).Exists() {
		ctx.Panic("!Hash.exist")
	}
	hash := ctx.Utility().Hash([]byte("Hash"))
	if !ctx.Params().GetHash(client.Key("Hash")).Value().Equals(hash) {
		ctx.Panic("Hash wrong")
	}
	if !ctx.Params().GetHname(ParamHname).Exists() {
		ctx.Panic("!Hname. exist")
	}
	if !ctx.Params().GetHname(ParamHname).Value().Equals(client.NewHname(string(ParamHname))) {
		ctx.Panic("Hname wrong")
	}
	if !ctx.Params().GetHname(ParamHname0).Exists() {
		ctx.Panic("!Hname-0.exist")
	}
	if !ctx.Params().GetHname(ParamHname0).Value().Equals(client.Hname(0)) {
		ctx.Panic("Hname-0 wrong")
	}
}
