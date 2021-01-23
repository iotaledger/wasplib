// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package testcore

import "github.com/iotaledger/wasplib/client"

const ParamIntParamName = client.Key("intParamName")
const ParamIntParamValue = client.Key("intParamValue")

// const PARAM_HNAME: &str = "hname";
// const PARAM_CALL_OPTION: &str = "callOption";
const ParamAddress = client.Key("address")
const ParamChainOwner = client.Key("chainOwner")
const ParamContractId = client.Key("contractID")

const MsgFullPanic string = "========== panic FULL ENTRY POINT ========="
const MsgViewPanic string = "========== panic VIEW ========="
const MsgPanicUnauthorized string = "============== panic due to unauthorized call"

const SelfName = client.Key("test_sandbox") // temporary, until hname in the call will become available

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

	exports.AddCall("sendToAddress", sendToAddress)
}

func onInit(ctx *client.ScCallContext) {
	ctx.Log("testcore.OnInit.Wasm.Begin")
}

func doNothing(ctx *client.ScCallContext) {
	ctx.Log("testcore.DoNothing.Begin")
}

func setInt(ctx *client.ScCallContext) {
	ctx.Log("testcore.SetInt.Begin")
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
	ctx.Log("testcore.GetInt.Begin")
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
	ctx.Log("testcore.CallOnChain.Begin")
	//
	// let param_call_option = ctx.params().get_string(PARAM_CALL_OPTION);
	// let param_value = ctx.params().get_int(PARAM_INT_PARAM_VALUE);
	// if !param_value.exists(){
	//     ctx.panic("param value not found")
	// }
	// TODO cannot get hname type
}

func fibonacci(ctx *client.ScViewContext) {
	nParam := ctx.Params().GetInt(ParamIntParamValue)
	if !nParam.Exists() {
		ctx.Panic("param value not found")
	}
	n := nParam.Value()
	ctx.Log("fibonacci: " + nParam.String())
	if n == 0 || n == 1 {
		ctx.Log("return 1")
		ctx.Results().GetInt(ParamIntParamValue).SetValue(n)
		return
	}
	ctx.Log("before call 1")
	params1 := client.NewScMutableMap()
	params1.GetInt(ParamIntParamValue).SetValue(n - 1)
	results1 := ctx.Call(0, client.NewHname("fibonacci"), params1)
	n1 := results1.GetInt(ParamIntParamValue)
	ctx.Log("    fibonacci-1: " + n1.String())

	params2 := client.NewScMutableMap()
	params2.GetInt(ParamIntParamValue).SetValue(n - 2)
	results2 := ctx.Call(0, client.NewHname("fibonacci"), params2)
	n2 := results2.GetInt(ParamIntParamValue)
	ctx.Log("    fibonacci-2: " + n2.String())

	ctx.Results().GetInt(ParamIntParamValue).SetValue(n1.Value() + n2.Value())
}

func testPanicFullEp(ctx *client.ScCallContext) {
	ctx.Panic(MsgFullPanic)
}

func testPanicViewEp(ctx *client.ScViewContext) {
	ctx.Panic(MsgViewPanic)
}

func testCallPanicFullEp(ctx *client.ScCallContext) {
	ctx.Call(0, client.NewHname("testPanicFullEP"), nil, nil)
}

// FIXME no need for 'view method special'
func testCallPanicViewFromFull(ctx *client.ScCallContext) {
	ctx.Call(0, client.NewHname("testPanicViewEP"), nil, nil)
}

// FIXME no need for 'view method special'
func testCallPanicViewFromView(ctx *client.ScViewContext) {
	ctx.Call(0, client.NewHname("testPanicViewEP"), nil)
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
	//myColors := myBalances.Colors()
	//length := myColors.Length()
	//for i := int32(0); i < length; i++ {
	//	color := myColors.GetColor(i)
    //	ctx.Log("Color: " + color.String())
	//}
	ctx.TransferToAddress(targetAddr.Value(), myBalances)
}

func testChainOwnerIdView(ctx *client.ScViewContext) {
	ctx.Results().GetAgent(ParamChainOwner).SetValue(ctx.ChainOwner())
}

func testChainOwnerIdFull(ctx *client.ScCallContext) {
	ctx.Results().GetAgent(ParamChainOwner).SetValue(ctx.ChainOwner())
}

func testContractIdView(_ctx *client.ScViewContext) {
	// TODO there's no way to return contact ID
	// ctx.results().(PARAM_CONTRACT_ID).set_value(ctx.chain_owner().value)
}

func testContractIdFull(_ctx *client.ScCallContext) {

}

func testSandboxCall(ctx *client.ScViewContext) {
	ret := ctx.Call(client.CoreRoot, client.ViewGetChainInfo, nil)
	desc := ret.GetString(client.Key("d")).Value()
	ctx.Results().GetString(client.Key("sandboxCall")).SetValue(desc)
}
