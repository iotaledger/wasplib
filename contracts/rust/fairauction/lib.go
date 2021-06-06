// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package fairauction

import "github.com/iotaledger/wasplib/packages/vm/wasmlib"

func OnLoad() {
	exports := wasmlib.NewScExports()
	exports.AddFunc(FuncFinalizeAuction, funcFinalizeAuctionThunk)
	exports.AddFunc(FuncPlaceBid, funcPlaceBidThunk)
	exports.AddFunc(FuncSetOwnerMargin, funcSetOwnerMarginThunk)
	exports.AddFunc(FuncStartAuction, funcStartAuctionThunk)
	exports.AddView(ViewGetInfo, viewGetInfoThunk)

	for i, key := range keyMap {
		idxMap[i] = wasmlib.GetKeyIdFromString(key)
	}
}

type FuncFinalizeAuctionParams struct {
	Color wasmlib.ScImmutableColor // color identifies the auction
}

type FuncFinalizeAuctionContext struct {
	Params FuncFinalizeAuctionParams
	State  FairAuctionFuncState
}

func funcFinalizeAuctionThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("fairauction.funcFinalizeAuction")
	// only SC itself can invoke this function
	ctx.Require(ctx.Caller() == ctx.AccountId(), "no permission")

	p := ctx.Params().MapId()
	f := &FuncFinalizeAuctionContext{
		Params: FuncFinalizeAuctionParams{
			Color: wasmlib.NewScImmutableColor(p, idxMap[IdxParamColor]),
		},
		State: FairAuctionFuncState{
			stateId: wasmlib.GetObjectId(1, wasmlib.KeyState, wasmlib.TYPE_MAP),
		},
	}
	ctx.Require(f.Params.Color.Exists(), "missing mandatory color")
	funcFinalizeAuction(ctx, f)
	ctx.Log("fairauction.funcFinalizeAuction ok")
}

type FuncPlaceBidParams struct {
	Color wasmlib.ScImmutableColor // color identifies the auction
}

type FuncPlaceBidContext struct {
	Params FuncPlaceBidParams
	State  FairAuctionFuncState
}

func funcPlaceBidThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("fairauction.funcPlaceBid")
	p := ctx.Params().MapId()
	f := &FuncPlaceBidContext{
		Params: FuncPlaceBidParams{
			Color: wasmlib.NewScImmutableColor(p, idxMap[IdxParamColor]),
		},
		State: FairAuctionFuncState{
			stateId: wasmlib.GetObjectId(1, wasmlib.KeyState, wasmlib.TYPE_MAP),
		},
	}
	ctx.Require(f.Params.Color.Exists(), "missing mandatory color")
	funcPlaceBid(ctx, f)
	ctx.Log("fairauction.funcPlaceBid ok")
}

type FuncSetOwnerMarginParams struct {
	OwnerMargin wasmlib.ScImmutableInt64 // new SC owner margin in promilles
}

type FuncSetOwnerMarginContext struct {
	Params FuncSetOwnerMarginParams
	State  FairAuctionFuncState
}

func funcSetOwnerMarginThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("fairauction.funcSetOwnerMargin")
	// only SC creator can set owner margin
	ctx.Require(ctx.Caller() == ctx.ContractCreator(), "no permission")

	p := ctx.Params().MapId()
	f := &FuncSetOwnerMarginContext{
		Params: FuncSetOwnerMarginParams{
			OwnerMargin: wasmlib.NewScImmutableInt64(p, idxMap[IdxParamOwnerMargin]),
		},
		State: FairAuctionFuncState{
			stateId: wasmlib.GetObjectId(1, wasmlib.KeyState, wasmlib.TYPE_MAP),
		},
	}
	ctx.Require(f.Params.OwnerMargin.Exists(), "missing mandatory ownerMargin")
	funcSetOwnerMargin(ctx, f)
	ctx.Log("fairauction.funcSetOwnerMargin ok")
}

type FuncStartAuctionParams struct {
	Color       wasmlib.ScImmutableColor  // color of the tokens being auctioned
	Description wasmlib.ScImmutableString // description of the tokens being auctioned
	Duration    wasmlib.ScImmutableInt64  // duration of auction in minutes
	MinimumBid  wasmlib.ScImmutableInt64  // minimum required amount for any bid
}

type FuncStartAuctionContext struct {
	Params FuncStartAuctionParams
	State  FairAuctionFuncState
}

func funcStartAuctionThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("fairauction.funcStartAuction")
	p := ctx.Params().MapId()
	f := &FuncStartAuctionContext{
		Params: FuncStartAuctionParams{
			Color:       wasmlib.NewScImmutableColor(p, idxMap[IdxParamColor]),
			Description: wasmlib.NewScImmutableString(p, idxMap[IdxParamDescription]),
			Duration:    wasmlib.NewScImmutableInt64(p, idxMap[IdxParamDuration]),
			MinimumBid:  wasmlib.NewScImmutableInt64(p, idxMap[IdxParamMinimumBid]),
		},
		State: FairAuctionFuncState{
			stateId: wasmlib.GetObjectId(1, wasmlib.KeyState, wasmlib.TYPE_MAP),
		},
	}
	ctx.Require(f.Params.Color.Exists(), "missing mandatory color")
	ctx.Require(f.Params.MinimumBid.Exists(), "missing mandatory minimumBid")
	funcStartAuction(ctx, f)
	ctx.Log("fairauction.funcStartAuction ok")
}

type ViewGetInfoParams struct {
	Color wasmlib.ScImmutableColor // color identifies the auction
}

type ViewGetInfoResults struct {
	Bidders       wasmlib.ScMutableInt64   // nr of bidders
	Color         wasmlib.ScMutableColor   // color of tokens for sale
	Creator       wasmlib.ScMutableAgentId // issuer of start_auction transaction
	Deposit       wasmlib.ScMutableInt64   // deposit by auction owner to cover the SC fees
	Description   wasmlib.ScMutableString  // auction description
	Duration      wasmlib.ScMutableInt64   // auction duration in minutes
	HighestBid    wasmlib.ScMutableInt64   // the current highest bid amount
	HighestBidder wasmlib.ScMutableAgentId // the current highest bidder
	MinimumBid    wasmlib.ScMutableInt64   // minimum bid amount
	NumTokens     wasmlib.ScMutableInt64   // number of tokens for sale
	OwnerMargin   wasmlib.ScMutableInt64   // auction owner's margin in promilles
	WhenStarted   wasmlib.ScMutableInt64   // timestamp when auction started
}

type ViewGetInfoContext struct {
	Params  ViewGetInfoParams
	Results ViewGetInfoResults
	State   FairAuctionViewState
}

func viewGetInfoThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("fairauction.viewGetInfo")
	p := ctx.Params().MapId()
	r := ctx.Results().MapId()
	f := &ViewGetInfoContext{
		Params: ViewGetInfoParams{
			Color: wasmlib.NewScImmutableColor(p, idxMap[IdxParamColor]),
		},
		Results: ViewGetInfoResults{
			Bidders:       wasmlib.NewScMutableInt64(r, idxMap[IdxResultBidders]),
			Color:         wasmlib.NewScMutableColor(r, idxMap[IdxResultColor]),
			Creator:       wasmlib.NewScMutableAgentId(r, idxMap[IdxResultCreator]),
			Deposit:       wasmlib.NewScMutableInt64(r, idxMap[IdxResultDeposit]),
			Description:   wasmlib.NewScMutableString(r, idxMap[IdxResultDescription]),
			Duration:      wasmlib.NewScMutableInt64(r, idxMap[IdxResultDuration]),
			HighestBid:    wasmlib.NewScMutableInt64(r, idxMap[IdxResultHighestBid]),
			HighestBidder: wasmlib.NewScMutableAgentId(r, idxMap[IdxResultHighestBidder]),
			MinimumBid:    wasmlib.NewScMutableInt64(r, idxMap[IdxResultMinimumBid]),
			NumTokens:     wasmlib.NewScMutableInt64(r, idxMap[IdxResultNumTokens]),
			OwnerMargin:   wasmlib.NewScMutableInt64(r, idxMap[IdxResultOwnerMargin]),
			WhenStarted:   wasmlib.NewScMutableInt64(r, idxMap[IdxResultWhenStarted]),
		},
		State: FairAuctionViewState{
			stateId: wasmlib.GetObjectId(1, wasmlib.KeyState, wasmlib.TYPE_MAP),
		},
	}
	ctx.Require(f.Params.Color.Exists(), "missing mandatory color")
	viewGetInfo(ctx, f)
	ctx.Log("fairauction.viewGetInfo ok")
}
