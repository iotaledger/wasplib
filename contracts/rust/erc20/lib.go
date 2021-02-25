// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package erc20

import "github.com/iotaledger/wasplib/packages/vm/wasmlib"

func OnLoad() {
	exports := wasmlib.NewScExports()
	exports.AddFunc(FuncApprove, funcApproveThunk)
	exports.AddFunc(FuncInit, funcInitThunk)
	exports.AddFunc(FuncTransfer, funcTransferThunk)
	exports.AddFunc(FuncTransferFrom, funcTransferFromThunk)
	exports.AddView(ViewAllowance, viewAllowanceThunk)
	exports.AddView(ViewBalanceOf, viewBalanceOfThunk)
	exports.AddView(ViewTotalSupply, viewTotalSupplyThunk)
}

type FuncApproveParams struct {
	Amount     wasmlib.ScImmutableInt64             // allowance value for delegated account
	Delegation wasmlib.ScImmutableAgentId           // delegated account
}

func funcApproveThunk(ctx wasmlib.ScFuncContext) {
	p := ctx.Params()
	params := &FuncApproveParams{
		Amount:     p.GetInt64(ParamAmount),
		Delegation: p.GetAgentId(ParamDelegation),
	}
	ctx.Require(params.Amount.Exists(), "missing mandatory amount")
	ctx.Require(params.Delegation.Exists(), "missing mandatory delegation")
	funcApprove(ctx, params)
}

type FuncInitParams struct {
	Creator wasmlib.ScImmutableAgentId           // creator/owner of the initial supply
	Supply  wasmlib.ScImmutableInt64             // initial token supply
}

func funcInitThunk(ctx wasmlib.ScFuncContext) {
	p := ctx.Params()
	params := &FuncInitParams{
		Creator: p.GetAgentId(ParamCreator),
		Supply:  p.GetInt64(ParamSupply),
	}
	ctx.Require(params.Creator.Exists(), "missing mandatory creator")
	ctx.Require(params.Supply.Exists(), "missing mandatory supply")
	funcInit(ctx, params)
}

type FuncTransferParams struct {
	Account wasmlib.ScImmutableAgentId           // target account
	Amount  wasmlib.ScImmutableInt64             // amount of tokens to transfer
}

func funcTransferThunk(ctx wasmlib.ScFuncContext) {
	p := ctx.Params()
	params := &FuncTransferParams{
		Account: p.GetAgentId(ParamAccount),
		Amount:  p.GetInt64(ParamAmount),
	}
	ctx.Require(params.Account.Exists(), "missing mandatory account")
	ctx.Require(params.Amount.Exists(), "missing mandatory amount")
	funcTransfer(ctx, params)
}

type FuncTransferFromParams struct {
	Account   wasmlib.ScImmutableAgentId           // sender account
	Amount    wasmlib.ScImmutableInt64             // amount of tokens to transfer
	Recipient wasmlib.ScImmutableAgentId           // recipient account
}

func funcTransferFromThunk(ctx wasmlib.ScFuncContext) {
	p := ctx.Params()
	params := &FuncTransferFromParams{
		Account:   p.GetAgentId(ParamAccount),
		Amount:    p.GetInt64(ParamAmount),
		Recipient: p.GetAgentId(ParamRecipient),
	}
	ctx.Require(params.Account.Exists(), "missing mandatory account")
	ctx.Require(params.Amount.Exists(), "missing mandatory amount")
	ctx.Require(params.Recipient.Exists(), "missing mandatory recipient")
	funcTransferFrom(ctx, params)
}

type ViewAllowanceParams struct {
	Account    wasmlib.ScImmutableAgentId           // sender account
	Delegation wasmlib.ScImmutableAgentId           // delegated account
}

func viewAllowanceThunk(ctx wasmlib.ScViewContext) {
	p := ctx.Params()
	params := &ViewAllowanceParams{
		Account:    p.GetAgentId(ParamAccount),
		Delegation: p.GetAgentId(ParamDelegation),
	}
	ctx.Require(params.Account.Exists(), "missing mandatory account")
	ctx.Require(params.Delegation.Exists(), "missing mandatory delegation")
	viewAllowance(ctx, params)
}

type ViewBalanceOfParams struct {
	Account wasmlib.ScImmutableAgentId           // sender account
}

func viewBalanceOfThunk(ctx wasmlib.ScViewContext) {
	p := ctx.Params()
	params := &ViewBalanceOfParams{
		Account: p.GetAgentId(ParamAccount),
	}
	ctx.Require(params.Account.Exists(), "missing mandatory account")
	viewBalanceOf(ctx, params)
}

type ViewTotalSupplyParams struct {
}

func viewTotalSupplyThunk(ctx wasmlib.ScViewContext) {
	params := &ViewTotalSupplyParams{
	}
	viewTotalSupply(ctx, params)
}
