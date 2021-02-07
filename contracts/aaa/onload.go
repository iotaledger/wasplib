// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package aaa

import "github.com/iotaledger/wasp/packages/vm/wasmlib"

func OnLoad() {
	exports := wasmlib.NewScExports()
	exports.AddCall(FuncDonate, funcDonateThunk)
	exports.AddCall(FuncWithdraw, funcWithdrawThunk)
	exports.AddView(ViewDonations, viewDonationsThunk)
}

type FuncDonateParams struct {
	Feedback wasmlib.ScImmutableString
}

func funcDonateThunk(ctx *wasmlib.ScCallContext) {
	p := ctx.Params()
	params := &FuncDonateParams{
		Feedback: p.GetString(ParamFeedback),
	}
	ctx.Require(params.Feedback.Exists(), "missing mandatory feedback")
	funcDonate(ctx, params)
}

type FuncWithdrawParams struct {
	Amount wasmlib.ScImmutableInt
}

func funcWithdrawThunk(ctx *wasmlib.ScCallContext) {
	p := ctx.Params()
	params := &FuncWithdrawParams{
		Amount: p.GetInt(ParamAmount),
	}
	ctx.Require(params.Amount.Exists(), "missing mandatory amount")
	funcWithdraw(ctx, params)
}

type ViewDonationsParams struct {
}

func viewDonationsThunk(ctx *wasmlib.ScViewContext) {
	params := &ViewDonationsParams{
	}
	viewDonations(ctx, params)
}
