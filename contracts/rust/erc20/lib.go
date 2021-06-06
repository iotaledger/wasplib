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

	for i, key := range keyMap {
		idxMap[i] = wasmlib.GetKeyIdFromString(key)
	}
}

type FuncApproveParams struct {
	Amount     wasmlib.ScImmutableInt64   // allowance value for delegated account
	Delegation wasmlib.ScImmutableAgentId // delegated account
}

type FuncApproveContext struct {
	Params FuncApproveParams
	State  Erc20FuncState
}

func funcApproveThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("erc20.funcApprove")
	p := ctx.Params().MapId()
	f := &FuncApproveContext{
		Params: FuncApproveParams{
			Amount:     wasmlib.NewScImmutableInt64(p, idxMap[IdxParamAmount]),
			Delegation: wasmlib.NewScImmutableAgentId(p, idxMap[IdxParamDelegation]),
		},
		State: Erc20FuncState{
			stateId: wasmlib.GetObjectId(1, wasmlib.KeyState, wasmlib.TYPE_MAP),
		},
	}
	ctx.Require(f.Params.Amount.Exists(), "missing mandatory amount")
	ctx.Require(f.Params.Delegation.Exists(), "missing mandatory delegation")
	funcApprove(ctx, f)
	ctx.Log("erc20.funcApprove ok")
}

type FuncInitParams struct {
	Creator wasmlib.ScImmutableAgentId // creator/owner of the initial supply
	Supply  wasmlib.ScImmutableInt64   // initial token supply
}

type FuncInitContext struct {
	Params FuncInitParams
	State  Erc20FuncState
}

func funcInitThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("erc20.funcInit")
	p := ctx.Params().MapId()
	f := &FuncInitContext{
		Params: FuncInitParams{
			Creator: wasmlib.NewScImmutableAgentId(p, idxMap[IdxParamCreator]),
			Supply:  wasmlib.NewScImmutableInt64(p, idxMap[IdxParamSupply]),
		},
		State: Erc20FuncState{
			stateId: wasmlib.GetObjectId(1, wasmlib.KeyState, wasmlib.TYPE_MAP),
		},
	}
	ctx.Require(f.Params.Creator.Exists(), "missing mandatory creator")
	ctx.Require(f.Params.Supply.Exists(), "missing mandatory supply")
	funcInit(ctx, f)
	ctx.Log("erc20.funcInit ok")
}

type FuncTransferParams struct {
	Account wasmlib.ScImmutableAgentId // target account
	Amount  wasmlib.ScImmutableInt64   // amount of tokens to transfer
}

type FuncTransferContext struct {
	Params FuncTransferParams
	State  Erc20FuncState
}

func funcTransferThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("erc20.funcTransfer")
	p := ctx.Params().MapId()
	f := &FuncTransferContext{
		Params: FuncTransferParams{
			Account: wasmlib.NewScImmutableAgentId(p, idxMap[IdxParamAccount]),
			Amount:  wasmlib.NewScImmutableInt64(p, idxMap[IdxParamAmount]),
		},
		State: Erc20FuncState{
			stateId: wasmlib.GetObjectId(1, wasmlib.KeyState, wasmlib.TYPE_MAP),
		},
	}
	ctx.Require(f.Params.Account.Exists(), "missing mandatory account")
	ctx.Require(f.Params.Amount.Exists(), "missing mandatory amount")
	funcTransfer(ctx, f)
	ctx.Log("erc20.funcTransfer ok")
}

type FuncTransferFromParams struct {
	Account   wasmlib.ScImmutableAgentId // sender account
	Amount    wasmlib.ScImmutableInt64   // amount of tokens to transfer
	Recipient wasmlib.ScImmutableAgentId // recipient account
}

type FuncTransferFromContext struct {
	Params FuncTransferFromParams
	State  Erc20FuncState
}

func funcTransferFromThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("erc20.funcTransferFrom")
	p := ctx.Params().MapId()
	f := &FuncTransferFromContext{
		Params: FuncTransferFromParams{
			Account:   wasmlib.NewScImmutableAgentId(p, idxMap[IdxParamAccount]),
			Amount:    wasmlib.NewScImmutableInt64(p, idxMap[IdxParamAmount]),
			Recipient: wasmlib.NewScImmutableAgentId(p, idxMap[IdxParamRecipient]),
		},
		State: Erc20FuncState{
			stateId: wasmlib.GetObjectId(1, wasmlib.KeyState, wasmlib.TYPE_MAP),
		},
	}
	ctx.Require(f.Params.Account.Exists(), "missing mandatory account")
	ctx.Require(f.Params.Amount.Exists(), "missing mandatory amount")
	ctx.Require(f.Params.Recipient.Exists(), "missing mandatory recipient")
	funcTransferFrom(ctx, f)
	ctx.Log("erc20.funcTransferFrom ok")
}

type ViewAllowanceParams struct {
	Account    wasmlib.ScImmutableAgentId // sender account
	Delegation wasmlib.ScImmutableAgentId // delegated account
}

type ViewAllowanceResults struct {
	Amount wasmlib.ScMutableInt64
}

type ViewAllowanceContext struct {
	Params  ViewAllowanceParams
	Results ViewAllowanceResults
	State   Erc20ViewState
}

func viewAllowanceThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("erc20.viewAllowance")
	p := ctx.Params().MapId()
	r := ctx.Results().MapId()
	f := &ViewAllowanceContext{
		Params: ViewAllowanceParams{
			Account:    wasmlib.NewScImmutableAgentId(p, idxMap[IdxParamAccount]),
			Delegation: wasmlib.NewScImmutableAgentId(p, idxMap[IdxParamDelegation]),
		},
		Results: ViewAllowanceResults{
			Amount: wasmlib.NewScMutableInt64(r, idxMap[IdxResultAmount]),
		},
		State: Erc20ViewState{
			stateId: wasmlib.GetObjectId(1, wasmlib.KeyState, wasmlib.TYPE_MAP),
		},
	}
	ctx.Require(f.Params.Account.Exists(), "missing mandatory account")
	ctx.Require(f.Params.Delegation.Exists(), "missing mandatory delegation")
	viewAllowance(ctx, f)
	ctx.Log("erc20.viewAllowance ok")
}

type ViewBalanceOfParams struct {
	Account wasmlib.ScImmutableAgentId // sender account
}

type ViewBalanceOfResults struct {
	Amount wasmlib.ScMutableInt64
}

type ViewBalanceOfContext struct {
	Params  ViewBalanceOfParams
	Results ViewBalanceOfResults
	State   Erc20ViewState
}

func viewBalanceOfThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("erc20.viewBalanceOf")
	p := ctx.Params().MapId()
	r := ctx.Results().MapId()
	f := &ViewBalanceOfContext{
		Params: ViewBalanceOfParams{
			Account: wasmlib.NewScImmutableAgentId(p, idxMap[IdxParamAccount]),
		},
		Results: ViewBalanceOfResults{
			Amount: wasmlib.NewScMutableInt64(r, idxMap[IdxResultAmount]),
		},
		State: Erc20ViewState{
			stateId: wasmlib.GetObjectId(1, wasmlib.KeyState, wasmlib.TYPE_MAP),
		},
	}
	ctx.Require(f.Params.Account.Exists(), "missing mandatory account")
	viewBalanceOf(ctx, f)
	ctx.Log("erc20.viewBalanceOf ok")
}

type ViewTotalSupplyResults struct {
	Supply wasmlib.ScMutableInt64
}

type ViewTotalSupplyContext struct {
	Results ViewTotalSupplyResults
	State   Erc20ViewState
}

func viewTotalSupplyThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("erc20.viewTotalSupply")
	r := ctx.Results().MapId()
	f := &ViewTotalSupplyContext{
		Results: ViewTotalSupplyResults{
			Supply: wasmlib.NewScMutableInt64(r, idxMap[IdxResultSupply]),
		},
		State: Erc20ViewState{
			stateId: wasmlib.GetObjectId(1, wasmlib.KeyState, wasmlib.TYPE_MAP),
		},
	}
	viewTotalSupply(ctx, f)
	ctx.Log("erc20.viewTotalSupply ok")
}
