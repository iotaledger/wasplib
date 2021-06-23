// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package fairauction

import "github.com/iotaledger/wasplib/packages/vm/wasmlib"

type FinalizeAuctionCall struct {
	Func   wasmlib.ScFunc
	Params MutableFinalizeAuctionParams
}

func NewFinalizeAuctionCall(ctx wasmlib.ScFuncContext) *FinalizeAuctionCall {
	f := &FinalizeAuctionCall{}
	f.Func.Init(HScName, HFuncFinalizeAuction, &f.Params.id, nil)
	return f
}

type PlaceBidCall struct {
	Func   wasmlib.ScFunc
	Params MutablePlaceBidParams
}

func NewPlaceBidCall(ctx wasmlib.ScFuncContext) *PlaceBidCall {
	f := &PlaceBidCall{}
	f.Func.Init(HScName, HFuncPlaceBid, &f.Params.id, nil)
	return f
}

type SetOwnerMarginCall struct {
	Func   wasmlib.ScFunc
	Params MutableSetOwnerMarginParams
}

func NewSetOwnerMarginCall(ctx wasmlib.ScFuncContext) *SetOwnerMarginCall {
	f := &SetOwnerMarginCall{}
	f.Func.Init(HScName, HFuncSetOwnerMargin, &f.Params.id, nil)
	return f
}

type StartAuctionCall struct {
	Func   wasmlib.ScFunc
	Params MutableStartAuctionParams
}

func NewStartAuctionCall(ctx wasmlib.ScFuncContext) *StartAuctionCall {
	f := &StartAuctionCall{}
	f.Func.Init(HScName, HFuncStartAuction, &f.Params.id, nil)
	return f
}

type GetInfoCall struct {
	Func    wasmlib.ScView
	Params  MutableGetInfoParams
	Results ImmutableGetInfoResults
}

func NewGetInfoCall(ctx wasmlib.ScFuncContext) *GetInfoCall {
	f := &GetInfoCall{}
	f.Func.Init(HScName, HViewGetInfo, &f.Params.id, &f.Results.id)
	return f
}

func NewGetInfoCallFromView(ctx wasmlib.ScViewContext) *GetInfoCall {
	f := &GetInfoCall{}
	f.Func.Init(HScName, HViewGetInfo, &f.Params.id, &f.Results.id)
	return f
}
