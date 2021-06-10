// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package donatewithfeedback

import "github.com/iotaledger/wasplib/packages/vm/wasmlib"

func OnLoad() {
	exports := wasmlib.NewScExports()
	exports.AddFunc(FuncDonate, funcDonateThunk)
	exports.AddFunc(FuncWithdraw, funcWithdrawThunk)
	exports.AddView(ViewDonation, viewDonationThunk)
	exports.AddView(ViewDonationInfo, viewDonationInfoThunk)

	for i, key := range keyMap {
		idxMap[i] = key.KeyId()
	}
}

type FuncDonateContext struct {
	Params ImmutableFuncDonateParams
	State  MutableDonateWithFeedbackState
}

func funcDonateThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("donatewithfeedback.funcDonate")
	f := &FuncDonateContext{
		Params: ImmutableFuncDonateParams{
			id: wasmlib.GetObjectId(1, wasmlib.KeyParams, wasmlib.TYPE_MAP),
		},
		State: MutableDonateWithFeedbackState{
			id: wasmlib.GetObjectId(1, wasmlib.KeyState, wasmlib.TYPE_MAP),
		},
	}
	funcDonate(ctx, f)
	ctx.Log("donatewithfeedback.funcDonate ok")
}

type FuncWithdrawContext struct {
	Params ImmutableFuncWithdrawParams
	State  MutableDonateWithFeedbackState
}

func funcWithdrawThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("donatewithfeedback.funcWithdraw")
	// only SC creator can withdraw donated funds
	ctx.Require(ctx.Caller() == ctx.ContractCreator(), "no permission")

	f := &FuncWithdrawContext{
		Params: ImmutableFuncWithdrawParams{
			id: wasmlib.GetObjectId(1, wasmlib.KeyParams, wasmlib.TYPE_MAP),
		},
		State: MutableDonateWithFeedbackState{
			id: wasmlib.GetObjectId(1, wasmlib.KeyState, wasmlib.TYPE_MAP),
		},
	}
	funcWithdraw(ctx, f)
	ctx.Log("donatewithfeedback.funcWithdraw ok")
}

type ViewDonationContext struct {
	Params  ImmutableViewDonationParams
	Results MutableViewDonationResults
	State   ImmutableDonateWithFeedbackState
}

func viewDonationThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("donatewithfeedback.viewDonation")
	f := &ViewDonationContext{
		Params: ImmutableViewDonationParams{
			id: wasmlib.GetObjectId(1, wasmlib.KeyParams, wasmlib.TYPE_MAP),
		},
		Results: MutableViewDonationResults{
			id: wasmlib.GetObjectId(1, wasmlib.KeyResults, wasmlib.TYPE_MAP),
		},
		State: ImmutableDonateWithFeedbackState{
			id: wasmlib.GetObjectId(1, wasmlib.KeyState, wasmlib.TYPE_MAP),
		},
	}
	ctx.Require(f.Params.Nr().Exists(), "missing mandatory nr")
	viewDonation(ctx, f)
	ctx.Log("donatewithfeedback.viewDonation ok")
}

type ViewDonationInfoContext struct {
	Results MutableViewDonationInfoResults
	State   ImmutableDonateWithFeedbackState
}

func viewDonationInfoThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("donatewithfeedback.viewDonationInfo")
	f := &ViewDonationInfoContext{
		Results: MutableViewDonationInfoResults{
			id: wasmlib.GetObjectId(1, wasmlib.KeyResults, wasmlib.TYPE_MAP),
		},
		State: ImmutableDonateWithFeedbackState{
			id: wasmlib.GetObjectId(1, wasmlib.KeyState, wasmlib.TYPE_MAP),
		},
	}
	viewDonationInfo(ctx, f)
	ctx.Log("donatewithfeedback.viewDonationInfo ok")
}
