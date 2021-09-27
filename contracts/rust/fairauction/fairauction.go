// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package fairauction

import (
	"github.com/iotaledger/wasp/packages/vm/wasmlib"
)

const (
	DurationDefault      = 60
	DurationMin          = 1
	DurationMax          = 120
	MaxDescriptionLength = 150
	OwnerMarginDefault   = 50
	OwnerMarginMin       = 5
	OwnerMarginMax       = 100
)

func funcFinalizeAuction(ctx wasmlib.ScFuncContext, f *FinalizeAuctionContext) {
	color := f.Params.Color().Value()
	currentAuction := f.State.Auctions().GetAuction(color)
	ctx.Require(currentAuction.Exists(), "Missing auction info")
	auction := currentAuction.Value()
	if auction.HighestBid < 0 {
		ctx.Log("No one bid on " + color.String())
		ownerFee := auction.MinimumBid * auction.OwnerMargin / 1000
		if ownerFee == 0 {
			ownerFee = 1
		}
		// finalizeAuction request token was probably not confirmed yet
		transferTokens(ctx, ctx.ContractCreator(), wasmlib.IOTA, ownerFee-1)
		transferTokens(ctx, auction.Creator, auction.Color, auction.NumTokens)
		transferTokens(ctx, auction.Creator, wasmlib.IOTA, auction.Deposit-ownerFee)
		return
	}

	ownerFee := auction.HighestBid * auction.OwnerMargin / 1000
	if ownerFee == 0 {
		ownerFee = 1
	}

	// return staked bids to losers
	bids := f.State.Bids().GetBids(color)
	bidderList := f.State.BidderList().GetBidderList(color)
	size := bidderList.Length()
	for i := int32(0); i < size; i++ {
		loser := bidderList.GetAgentID(i).Value()
		if loser != auction.HighestBidder {
			bid := bids.GetBid(loser).Value()
			transferTokens(ctx, loser, wasmlib.IOTA, bid.Amount)
		}
	}

	// finalizeAuction request token was probably not confirmed yet
	transferTokens(ctx, ctx.ContractCreator(), wasmlib.IOTA, ownerFee-1)
	transferTokens(ctx, auction.HighestBidder, auction.Color, auction.NumTokens)
	transferTokens(ctx, auction.Creator, wasmlib.IOTA, auction.Deposit+auction.HighestBid-ownerFee)
}

func funcPlaceBid(ctx wasmlib.ScFuncContext, f *PlaceBidContext) {
	bidAmount := ctx.Incoming().Balance(wasmlib.IOTA)
	ctx.Require(bidAmount > 0, "Missing bid amount")

	color := f.Params.Color().Value()
	currentAuction := f.State.Auctions().GetAuction(color)
	ctx.Require(currentAuction.Exists(), "Missing auction info")

	auction := currentAuction.Value()
	bids := f.State.Bids().GetBids(color)
	bidderList := f.State.BidderList().GetBidderList(color)
	caller := ctx.Caller()
	currentBid := bids.GetBid(caller)
	if currentBid.Exists() {
		ctx.Log("Upped bid from: " + caller.String())
		bid := currentBid.Value()
		bidAmount += bid.Amount
		bid.Amount = bidAmount
		bid.Timestamp = ctx.Timestamp()
		currentBid.SetValue(bid)
	} else {
		ctx.Require(bidAmount >= auction.MinimumBid, "Insufficient bid amount")
		ctx.Log("New bid from: " + caller.String())
		index := bidderList.Length()
		bidderList.GetAgentID(index).SetValue(caller)
		bid := &Bid{
			Index:     index,
			Amount:    bidAmount,
			Timestamp: ctx.Timestamp(),
		}
		currentBid.SetValue(bid)
	}
	if bidAmount > auction.HighestBid {
		ctx.Log("New highest bidder")
		auction.HighestBid = bidAmount
		auction.HighestBidder = caller
		currentAuction.SetValue(auction)
	}
}

func funcSetOwnerMargin(ctx wasmlib.ScFuncContext, f *SetOwnerMarginContext) {
	ownerMargin := f.Params.OwnerMargin().Value()
	if ownerMargin < OwnerMarginMin {
		ownerMargin = OwnerMarginMin
	}
	if ownerMargin > OwnerMarginMax {
		ownerMargin = OwnerMarginMax
	}
	f.State.OwnerMargin().SetValue(ownerMargin)
}

func funcStartAuction(ctx wasmlib.ScFuncContext, f *StartAuctionContext) {
	color := f.Params.Color().Value()
	if color == wasmlib.IOTA || color == wasmlib.MINT {
		ctx.Panic("Reserved auction token color")
	}
	numTokens := ctx.Incoming().Balance(color)
	if numTokens == 0 {
		ctx.Panic("Missing auction tokens")
	}

	minimumBid := f.Params.MinimumBid().Value()

	// duration in minutes
	duration := f.Params.Duration().Value()
	if duration == 0 {
		duration = DurationDefault
	}
	if duration < DurationMin {
		duration = DurationMin
	}
	if duration > DurationMax {
		duration = DurationMax
	}

	description := f.Params.Description().Value()
	if description == "" {
		description = "N/A"
	}
	if len(description) > MaxDescriptionLength {
		ss := description[:MaxDescriptionLength]
		description = ss + "[...]"
	}

	ownerMargin := f.State.OwnerMargin().Value()
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

	currentAuction := f.State.Auctions().GetAuction(color)
	if currentAuction.Exists() {
		ctx.Panic("Auction for this token color already exists")
	}

	auction := &Auction{
		Creator:       ctx.Caller(),
		Color:         color,
		Deposit:       deposit,
		Description:   description,
		Duration:      duration,
		HighestBid:    -1,
		HighestBidder: wasmlib.ScAgentID{},
		MinimumBid:    minimumBid,
		NumTokens:     numTokens,
		OwnerMargin:   ownerMargin,
		WhenStarted:   ctx.Timestamp(),
	}
	currentAuction.SetValue(auction)

	fa := ScFuncs.FinalizeAuction(ctx)
	fa.Params.Color().SetValue(auction.Color)
	fa.Func.Delay(duration * 60).TransferIotas(1).Post()
}

func viewGetInfo(ctx wasmlib.ScViewContext, f *GetInfoContext) {
	color := f.Params.Color().Value()
	currentAuction := f.State.Auctions().GetAuction(color)
	if !currentAuction.Exists() {
		ctx.Panic("Missing auction info")
	}

	auction := currentAuction.Value()
	f.Results.Color().SetValue(auction.Color)
	f.Results.Creator().SetValue(auction.Creator)
	f.Results.Deposit().SetValue(auction.Deposit)
	f.Results.Description().SetValue(auction.Description)
	f.Results.Duration().SetValue(auction.Duration)
	f.Results.HighestBid().SetValue(auction.HighestBid)
	f.Results.HighestBidder().SetValue(auction.HighestBidder)
	f.Results.MinimumBid().SetValue(auction.MinimumBid)
	f.Results.NumTokens().SetValue(auction.NumTokens)
	f.Results.OwnerMargin().SetValue(auction.OwnerMargin)
	f.Results.WhenStarted().SetValue(auction.WhenStarted)

	bidderList := f.State.BidderList().GetBidderList(color)
	f.Results.Bidders().SetValue(bidderList.Length())
}

func transferTokens(ctx wasmlib.ScFuncContext, agent wasmlib.ScAgentID, color wasmlib.ScColor, amount int64) {
	if agent.IsAddress() {
		// send back to original Tangle address
		ctx.TransferToAddress(agent.Address(), wasmlib.NewScTransfer(color, amount))
		return
	}

	// TODO not an address, deposit into account on chain
	ctx.TransferToAddress(agent.Address(), wasmlib.NewScTransfer(color, amount))
}
