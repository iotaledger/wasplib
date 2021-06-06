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
		idxMap[i] = wasmlib.GetKeyIdFromString(key)
	}
}

type FuncDonateParams struct {
	Feedback wasmlib.ScImmutableString // feedback for the person you donate to
}

type FuncDonateContext struct {
	Params FuncDonateParams
	State  DonateWithFeedbackFuncState
}

func funcDonateThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("donatewithfeedback.funcDonate")
	p := ctx.Params().MapId()
	f := &FuncDonateContext{
		Params: FuncDonateParams{
			Feedback: wasmlib.NewScImmutableString(p, idxMap[IdxParamFeedback]),
		},
		State: DonateWithFeedbackFuncState{
			stateId: wasmlib.GetObjectId(1, wasmlib.KeyState, wasmlib.TYPE_MAP),
		},
	}
	funcDonate(ctx, f)
	ctx.Log("donatewithfeedback.funcDonate ok")
}

type FuncWithdrawParams struct {
	Amount wasmlib.ScImmutableInt64 // amount to withdraw
}

type FuncWithdrawContext struct {
	Params FuncWithdrawParams
	State  DonateWithFeedbackFuncState
}

func funcWithdrawThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("donatewithfeedback.funcWithdraw")
	// only SC creator can withdraw donated funds
	ctx.Require(ctx.Caller() == ctx.ContractCreator(), "no permission")

	p := ctx.Params().MapId()
	f := &FuncWithdrawContext{
		Params: FuncWithdrawParams{
			Amount: wasmlib.NewScImmutableInt64(p, idxMap[IdxParamAmount]),
		},
		State: DonateWithFeedbackFuncState{
			stateId: wasmlib.GetObjectId(1, wasmlib.KeyState, wasmlib.TYPE_MAP),
		},
	}
	funcWithdraw(ctx, f)
	ctx.Log("donatewithfeedback.funcWithdraw ok")
}

type ViewDonationParams struct {
	Nr wasmlib.ScImmutableInt64
}

type ViewDonationResults struct {
	Amount    wasmlib.ScMutableInt64   // amount donated
	Donator   wasmlib.ScMutableAgentId // who donated
	Error     wasmlib.ScMutableString  // error to be reported to donator if anything goes wrong
	Feedback  wasmlib.ScMutableString  // the feedback for the person donated to
	Timestamp wasmlib.ScMutableInt64   // when the donation took place
}

type ViewDonationContext struct {
	Params  ViewDonationParams
	Results ViewDonationResults
	State   DonateWithFeedbackViewState
}

func viewDonationThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("donatewithfeedback.viewDonation")
	p := ctx.Params().MapId()
	r := ctx.Results().MapId()
	f := &ViewDonationContext{
		Params: ViewDonationParams{
			Nr: wasmlib.NewScImmutableInt64(p, idxMap[IdxParamNr]),
		},
		Results: ViewDonationResults{
			Amount:    wasmlib.NewScMutableInt64(r, idxMap[IdxResultAmount]),
			Donator:   wasmlib.NewScMutableAgentId(r, idxMap[IdxResultDonator]),
			Error:     wasmlib.NewScMutableString(r, idxMap[IdxResultError]),
			Feedback:  wasmlib.NewScMutableString(r, idxMap[IdxResultFeedback]),
			Timestamp: wasmlib.NewScMutableInt64(r, idxMap[IdxResultTimestamp]),
		},
		State: DonateWithFeedbackViewState{
			stateId: wasmlib.GetObjectId(1, wasmlib.KeyState, wasmlib.TYPE_MAP),
		},
	}
	ctx.Require(f.Params.Nr.Exists(), "missing mandatory nr")
	viewDonation(ctx, f)
	ctx.Log("donatewithfeedback.viewDonation ok")
}

type ViewDonationInfoResults struct {
	Count         wasmlib.ScMutableInt64
	MaxDonation   wasmlib.ScMutableInt64
	TotalDonation wasmlib.ScMutableInt64
}

type ViewDonationInfoContext struct {
	Results ViewDonationInfoResults
	State   DonateWithFeedbackViewState
}

func viewDonationInfoThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("donatewithfeedback.viewDonationInfo")
	r := ctx.Results().MapId()
	f := &ViewDonationInfoContext{
		Results: ViewDonationInfoResults{
			Count:         wasmlib.NewScMutableInt64(r, idxMap[IdxResultCount]),
			MaxDonation:   wasmlib.NewScMutableInt64(r, idxMap[IdxResultMaxDonation]),
			TotalDonation: wasmlib.NewScMutableInt64(r, idxMap[IdxResultTotalDonation]),
		},
		State: DonateWithFeedbackViewState{
			stateId: wasmlib.GetObjectId(1, wasmlib.KeyState, wasmlib.TYPE_MAP),
		},
	}
	viewDonationInfo(ctx, f)
	ctx.Log("donatewithfeedback.viewDonationInfo ok")
}
