// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package fairauction

import "github.com/iotaledger/wasp/packages/vm/wasmlib"

const DurationDefault = 60
const DurationMin = 1
const DurationMax = 120
const MaxDescriptionLength = 150
const OwnerMarginDefault = 50
const OwnerMarginMin = 5
const OwnerMarginMax = 100

func funcFinalizeAuction(ctx *wasmlib.ScFuncContext, params *FuncFinalizeAuctionParams) {
	color := params.Color.Value()
	state := ctx.State()
	auctions := state.GetMap(VarAuctions)
	currentAuction := auctions.GetMap(color)
	auctionInfo := currentAuction.GetBytes(VarInfo)
	if !auctionInfo.Exists() {
		ctx.Panic("Missing auction info")
	}
	auction := NewAuctionFromBytes(auctionInfo.Value())
	if auction.HighestBid < 0 {
		ctx.Log("No one bid on " + color.String())
		ownerFee := auction.MinimumBid * auction.OwnerMargin / 1000
		if ownerFee == 0 {
			ownerFee = 1
		}
		// finalizeAuction request token was probably not confirmed yet
		transfer(ctx, ctx.ContractCreator(), wasmlib.IOTA, ownerFee-1)
		transfer(ctx, auction.Creator, auction.Color, auction.NumTokens)
		transfer(ctx, auction.Creator, wasmlib.IOTA, auction.Deposit-ownerFee)
		return
	}

	ownerFee := auction.HighestBid * auction.OwnerMargin / 1000
	if ownerFee == 0 {
		ownerFee = 1
	}

	// return staked bids to losers
	bidders := currentAuction.GetMap(VarBidders)
	bidderList := currentAuction.GetAgentIdArray(VarBidderList)
	size := bidderList.Length()
	for i := int32(0); i < size; i++ {
		bidder := bidderList.GetAgentId(i).Value()
		if !bidder.Equals(auction.HighestBidder) {
			loser := bidders.GetBytes(bidder)
			bid := NewBidFromBytes(loser.Value())
			transfer(ctx, bidder, wasmlib.IOTA, bid.Amount)
		}
	}

	// finalizeAuction request token was probably not confirmed yet
	transfer(ctx, ctx.ContractCreator(), wasmlib.IOTA, ownerFee-1)
	transfer(ctx, auction.HighestBidder, auction.Color, auction.NumTokens)
	transfer(ctx, auction.Creator, wasmlib.IOTA, auction.Deposit+auction.HighestBid-ownerFee)
}

func funcPlaceBid(ctx *wasmlib.ScFuncContext, params *FuncPlaceBidParams) {
	bidAmount := ctx.Incoming().Balance(wasmlib.IOTA)
	if bidAmount == 0 {
		ctx.Panic("Missing bid amount")
	}

	color := params.Color.Value()
	state := ctx.State()
	auctions := state.GetMap(VarAuctions)
	currentAuction := auctions.GetMap(color)
	auctionInfo := currentAuction.GetBytes(VarInfo)
	if !auctionInfo.Exists() {
		ctx.Panic("Missing auction info")
	}

	auction := NewAuctionFromBytes(auctionInfo.Value())
	bidders := currentAuction.GetMap(VarBidders)
	bidderList := currentAuction.GetAgentIdArray(VarBidderList)
	caller := ctx.Caller()
	bidder := bidders.GetBytes(caller)
	if bidder.Exists() {
		ctx.Log("Upped bid from: " + caller.String())
		bid := NewBidFromBytes(bidder.Value())
		bidAmount += bid.Amount
		bid.Amount = bidAmount
		bid.Timestamp = ctx.Timestamp()
		bidder.SetValue(bid.Bytes())
	} else {
		if bidAmount < auction.MinimumBid {
			ctx.Panic("Insufficient bid amount")
		}
		ctx.Log("New bid from: " + caller.String())
		index := bidderList.Length()
		bidderList.GetAgentId(index).SetValue(caller)
		bid := &Bid{
			Index:     int64(index),
			Amount:    bidAmount,
			Timestamp: ctx.Timestamp(),
		}
		bidder.SetValue(bid.Bytes())
	}
	if bidAmount > auction.HighestBid {
		ctx.Log("New highest bidder")
		auction.HighestBid = bidAmount
		auction.HighestBidder = caller
		auctionInfo.SetValue(auction.Bytes())
	}
}

func funcSetOwnerMargin(ctx *wasmlib.ScFuncContext, params *FuncSetOwnerMarginParams) {
	ownerMargin := params.OwnerMargin.Value()
	if ownerMargin < OwnerMarginMin {
		ownerMargin = OwnerMarginMin
	}
	if ownerMargin > OwnerMarginMax {
		ownerMargin = OwnerMarginMax
	}
	ctx.State().GetInt(VarOwnerMargin).SetValue(ownerMargin)
	ctx.Log("Updated owner margin")
}

func funcStartAuction(ctx *wasmlib.ScFuncContext, params *FuncStartAuctionParams) {
	color := params.Color.Value()
	if color.Equals(wasmlib.IOTA) || color.Equals(wasmlib.MINT) {
		ctx.Panic("Reserved auction token color")
	}
	numTokens := ctx.Incoming().Balance(color)
	if numTokens == 0 {
		ctx.Panic("Missing auction tokens")
	}

	minimumBid := params.MinimumBid.Value()

	// duration in minutes
	duration := params.Duration.Value()
	if duration == 0 {
		duration = DurationDefault
	}
	if duration < DurationMin {
		duration = DurationMin
	}
	if duration > DurationMax {
		duration = DurationMax
	}

	description := params.Description.Value()
	if description == "" {
		description = "N/A"
	}
	if len(description) > MaxDescriptionLength {
		description = description[:MaxDescriptionLength] + "[...]"
	}

	state := ctx.State()
	ownerMargin := state.GetInt(VarOwnerMargin).Value()
	if ownerMargin == 0 {
		ownerMargin = OwnerMarginDefault
	}

	// need at least 1 iota to run SC
	margin := minimumBid * ownerMargin / 1000
	if margin == 0 {
		margin = 1
	}
	deposit := ctx.Incoming().Balance(wasmlib.IOTA)
	if deposit < margin {
		ctx.Panic("Insufficient deposit")
	}

	auctions := state.GetMap(VarAuctions)
	currentAuction := auctions.GetMap(color)
	auctionInfo := currentAuction.GetBytes(VarInfo)
	if auctionInfo.Exists() {
		ctx.Panic("Auction for this token color already exists")
	}

	auction := &Auction{
		Creator:       ctx.Caller(),
		Color:         color,
		Deposit:       deposit,
		Description:   description,
		Duration:      duration,
		HighestBid:    -1,
		HighestBidder: &wasmlib.ScAgentId{},
		MinimumBid:    minimumBid,
		NumTokens:     numTokens,
		OwnerMargin:   ownerMargin,
		WhenStarted:   ctx.Timestamp(),
	}
	auctionInfo.SetValue(auction.Bytes())

	finalizeParams := wasmlib.NewScMutableMap()
	finalizeParams.GetColor(VarColor).SetValue(auction.Color)
	ctx.Post(&wasmlib.PostRequestParams{
		ContractId: ctx.ContractId(),
		Function:   HFuncFinalizeAuction,
		Params:     finalizeParams,
		Transfer:   nil,
		Delay:      duration * 60,
	})
	ctx.Log("New auction started")
}

func viewGetInfo(ctx *wasmlib.ScViewContext, params *ViewGetInfoParams) {
	color := params.Color.Value()
	state := ctx.State()
	auctions := state.GetMap(VarAuctions)
	currentAuction := auctions.GetMap(color)
	auctionInfo := currentAuction.GetBytes(VarInfo)
	if !auctionInfo.Exists() {
		ctx.Panic("Missing auction info")
	}

	auction := NewAuctionFromBytes(auctionInfo.Value())
	results := ctx.Results()
	results.GetColor(VarColor).SetValue(auction.Color)
	results.GetAgentId(VarCreator).SetValue(auction.Creator)
	results.GetInt(VarDeposit).SetValue(auction.Deposit)
	results.GetString(VarDescription).SetValue(auction.Description)
	results.GetInt(VarDuration).SetValue(auction.Duration)
	results.GetInt(VarHighestBid).SetValue(auction.HighestBid)
	results.GetAgentId(VarHighestBidder).SetValue(auction.HighestBidder)
	results.GetInt(VarMinimumBid).SetValue(auction.MinimumBid)
	results.GetInt(VarNumTokens).SetValue(auction.NumTokens)
	results.GetInt(VarOwnerMargin).SetValue(auction.OwnerMargin)
	results.GetInt(VarWhenStarted).SetValue(auction.WhenStarted)

	bidderList := currentAuction.GetAgentIdArray(VarBidderList)
	results.GetInt(VarBidders).SetValue(int64(bidderList.Length()))
}

func transfer(ctx *wasmlib.ScFuncContext, agent *wasmlib.ScAgentId, color *wasmlib.ScColor, amount int64) {
	if agent.IsAddress() {
		// send back to original Tangle address
		ctx.TransferToAddress(agent.Address(), wasmlib.NewScTransfer(color, amount))
		return
	}

	// TODO not an address, deposit into account on chain
	ctx.TransferToAddress(agent.Address(), wasmlib.NewScTransfer(color, amount))
}
