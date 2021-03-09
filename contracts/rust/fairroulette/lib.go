// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package fairroulette

import "github.com/iotaledger/wasplib/packages/vm/wasmlib"

func OnLoad() {
	exports := wasmlib.NewScExports()
	exports.AddFunc(FuncLockBets, funcLockBetsThunk)
	exports.AddFunc(FuncPayWinners, funcPayWinnersThunk)
	exports.AddFunc(FuncPlaceBet, funcPlaceBetThunk)
	exports.AddFunc(FuncPlayPeriod, funcPlayPeriodThunk)
}

type FuncLockBetsParams struct {
}

func funcLockBetsThunk(ctx wasmlib.ScFuncContext) {
	// only SC itself can invoke this function
	ctx.Require(ctx.Caller() == ctx.ContractId().AsAgentId(), "no permission")

	params := &FuncLockBetsParams{
	}
	ctx.Log("fairroulette.funcLockBets")
	funcLockBets(ctx, params)
	ctx.Log("fairroulette.funcLockBets ok")
}

type FuncPayWinnersParams struct {
}

func funcPayWinnersThunk(ctx wasmlib.ScFuncContext) {
	// only SC itself can invoke this function
	ctx.Require(ctx.Caller() == ctx.ContractId().AsAgentId(), "no permission")

	params := &FuncPayWinnersParams{
	}
	ctx.Log("fairroulette.funcPayWinners")
	funcPayWinners(ctx, params)
	ctx.Log("fairroulette.funcPayWinners ok")
}

type FuncPlaceBetParams struct {
	Number wasmlib.ScImmutableInt64 // the number a better bets on
}

func funcPlaceBetThunk(ctx wasmlib.ScFuncContext) {
	p := ctx.Params()
	params := &FuncPlaceBetParams{
		Number: p.GetInt64(ParamNumber),
	}
	ctx.Require(params.Number.Exists(), "missing mandatory number")
	ctx.Log("fairroulette.funcPlaceBet")
	funcPlaceBet(ctx, params)
	ctx.Log("fairroulette.funcPlaceBet ok")
}

type FuncPlayPeriodParams struct {
	PlayPeriod wasmlib.ScImmutableInt64 // number of minutes in one playing round
}

func funcPlayPeriodThunk(ctx wasmlib.ScFuncContext) {
	// only SC creator can update the play period
	ctx.Require(ctx.Caller() == ctx.ContractCreator(), "no permission")

	p := ctx.Params()
	params := &FuncPlayPeriodParams{
		PlayPeriod: p.GetInt64(ParamPlayPeriod),
	}
	ctx.Require(params.PlayPeriod.Exists(), "missing mandatory playPeriod")
	ctx.Log("fairroulette.funcPlayPeriod")
	funcPlayPeriod(ctx, params)
	ctx.Log("fairroulette.funcPlayPeriod ok")
}
