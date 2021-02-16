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
	exports.AddView(ViewDonations, viewDonationsThunk)
}

type FuncDonateParams struct {
	Feedback wasmlib.ScImmutableString // feedback for the person you donate to
}

func funcDonateThunk(ctx wasmlib.ScFuncContext) {
	p := ctx.Params()
	params := &FuncDonateParams{
		Feedback: p.GetString(ParamFeedback),
	}
	funcDonate(ctx, params)
}

type FuncWithdrawParams struct {
	Amount wasmlib.ScImmutableInt // amount to withdraw
}

func funcWithdrawThunk(ctx wasmlib.ScFuncContext) {
	// only SC creator can withdraw donated funds
	ctx.Require(ctx.Caller() == ctx.ContractCreator(), "no permission")

	p := ctx.Params()
	params := &FuncWithdrawParams{
		Amount: p.GetInt(ParamAmount),
	}
	funcWithdraw(ctx, params)
}

type ViewDonationsParams struct {
}

func viewDonationsThunk(ctx wasmlib.ScViewContext) {
	params := &ViewDonationsParams{
	}
	viewDonations(ctx, params)
}
