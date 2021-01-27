// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package fairauction

import "github.com/iotaledger/wasplib/client"

const KeyAuctions = client.Key("auctions")
const KeyBidders = client.Key("bidders")
const KeyBidderList = client.Key("bidder_list")
const KeyColor = client.Key("color")
const KeyCreator = client.Key("creator")
const KeyDeposit = client.Key("deposit")
const KeyDescription = client.Key("description")
const KeyDuration = client.Key("duration")
const KeyHighestBid = client.Key("highest_bid")
const KeyHighestBidder = client.Key("highest_bidder")
const KeyInfo = client.Key("info")
const KeyMinimumBid = client.Key("minimum")
const KeyNumTokens = client.Key("num_tokens")
const KeyOwnerMargin = client.Key("owner_margin")
const KeyWhenStarted = client.Key("when_started")

const DurationDefault = 60
const DurationMin = 1
const DurationMax = 120
const MaxDescriptionLength = 150
const OwnerMarginDefault = 50
const OwnerMarginMin = 5
const OwnerMarginMax = 100

func OnLoad() {
	exports := client.NewScExports()
	exports.AddCall("start_auction", startAuction)
	exports.AddCall("finalize_auction", finalizeAuction)
	exports.AddCall("place_bid", placeBid)
	exports.AddCall("set_owner_margin", setOwnerMargin)
	exports.AddView("get_info", getInfo)
}

func startAuction(ctx *client.ScCallContext) {
	params := ctx.Params()
	colorParam := params.GetColor(KeyColor)
	if !colorParam.Exists() {
		ctx.Panic("Missing auction token color")
	}
	color := colorParam.Value()
	if color.Equals(client.IOTA) || color.Equals(client.MINT) {
		ctx.Panic("Reserved auction token color")
	}
	numTokens := ctx.Incoming().Balance(color)
	if numTokens == 0 {
		ctx.Panic("Missing auction tokens")
	}

	minimumBid := params.GetInt(KeyMinimumBid).Value()
	if minimumBid == 0 {
		ctx.Panic("Missing minimum bid")
	}

	// duration in minutes
	duration := params.GetInt(KeyDuration).Value()
	if duration == 0 {
		duration = DurationDefault
	}
	if duration < DurationMin {
		duration = DurationMin
	}
	if duration > DurationMax {
		duration = DurationMax
	}

	description := params.GetString(KeyDescription).Value()
	if description == "" {
		description = "N/A"
	}
	if len(description) > MaxDescriptionLength {
		description = description[:MaxDescriptionLength] + "[...]"
	}

	state := ctx.State()
	ownerMargin := state.GetInt(KeyOwnerMargin).Value()
	if ownerMargin == 0 {
		ownerMargin = OwnerMarginDefault
	}

	// need at least 1 iota to run SC
	margin := minimumBid * ownerMargin / 1000
	if margin == 0 {
		margin = 1
	}
	deposit := ctx.Incoming().Balance(client.IOTA)
	if deposit < margin {
		ctx.Panic("Insufficient deposit")
	}

	auctions := state.GetMap(KeyAuctions)
	currentAuction := auctions.GetMap(color)
	auctionInfo := currentAuction.GetBytes(KeyInfo)
	if auctionInfo.Exists() {
		ctx.Panic("Auction for this token color already exists")
	}

	auction := &AuctionInfo{
		Creator:       ctx.Caller(),
		Color:         color,
		Deposit:       deposit,
		Description:   description,
		Duration:      duration,
		HighestBid:    -1,
		HighestBidder: &client.ScAgent{},
		MinimumBid:    minimumBid,
		NumTokens:     numTokens,
		OwnerMargin:   ownerMargin,
		WhenStarted:   ctx.Timestamp(),
	}
	auctionInfo.SetValue(EncodeAuctionInfo(auction))

	finalizeParams := client.NewScMutableMap()
	finalizeParams.GetColor(KeyColor).SetValue(auction.Color)
    ctx.Post(&client.PostRequestParams {
        Contract: ctx.ContractId(),
        Function: client.NewHname("finalize_auction"),
        Params: finalizeParams,
        Transfer: nil,
        Delay: duration * 60,
    })
	ctx.Log("New auction started")
}

func finalizeAuction(ctx *client.ScCallContext) {
	// can only be sent by SC itself
	if !ctx.From(ctx.ContractId().AsAgent()) {
		ctx.Panic("Cancel spoofed request")
	}

	colorParam := ctx.Params().GetColor(KeyColor)
	if !colorParam.Exists() {
		ctx.Panic("Missing token color")
	}
	color := colorParam.Value()

	state := ctx.State()
	auctions := state.GetMap(KeyAuctions)
	currentAuction := auctions.GetMap(color)
	auctionInfo := currentAuction.GetBytes(KeyInfo)
	if !auctionInfo.Exists() {
		ctx.Panic("Missing auction info")
	}
	auction := DecodeAuctionInfo(auctionInfo.Value())
	if auction.HighestBid < 0 {
		ctx.Log("No one bid on " + color.String())
		ownerFee := auction.MinimumBid * auction.OwnerMargin / 1000
		if ownerFee == 0 {
			ownerFee = 1
		}
		// finalizeAuction request token was probably not confirmed yet
		transfer(ctx, ctx.ContractCreator(), client.IOTA, ownerFee-1)
		transfer(ctx, auction.Creator, auction.Color, auction.NumTokens)
		transfer(ctx, auction.Creator, client.IOTA, auction.Deposit-ownerFee)
		return
	}

	ownerFee := auction.HighestBid * auction.OwnerMargin / 1000
	if ownerFee == 0 {
		ownerFee = 1
	}

	// return staked bids to losers
	bidders := currentAuction.GetMap(KeyBidders)
	bidderList := currentAuction.GetAgentArray(KeyBidderList)
	size := bidderList.Length()
	for i := int32(0); i < size; i++ {
		bidder := bidderList.GetAgent(i).Value()
		if !bidder.Equals(auction.HighestBidder) {
			loser := bidders.GetBytes(bidder)
			bid := DecodeBidInfo(loser.Value())
			transfer(ctx, bidder, client.IOTA, bid.Amount)
		}
	}

	// finalizeAuction request token was probably not confirmed yet
	transfer(ctx, ctx.ContractCreator(), client.IOTA, ownerFee-1)
	transfer(ctx, auction.HighestBidder, auction.Color, auction.NumTokens)
	transfer(ctx, auction.Creator, client.IOTA, auction.Deposit+auction.HighestBid-ownerFee)
}

func placeBid(ctx *client.ScCallContext) {
	bidAmount := ctx.Incoming().Balance(client.IOTA)
	if bidAmount == 0 {
		ctx.Panic("Missing bid amount")
	}

	colorParam := ctx.Params().GetColor(KeyColor)
	if !colorParam.Exists() {
		ctx.Panic("Missing token color")
	}
	color := colorParam.Value()

	state := ctx.State()
	auctions := state.GetMap(KeyAuctions)
	currentAuction := auctions.GetMap(color)
	auctionInfo := currentAuction.GetBytes(KeyInfo)
	if !auctionInfo.Exists() {
		ctx.Panic("Missing auction info")
	}

	auction := DecodeAuctionInfo(auctionInfo.Value())
	bidders := currentAuction.GetMap(KeyBidders)
	bidderList := currentAuction.GetAgentArray(KeyBidderList)
	caller := ctx.Caller()
	bidder := bidders.GetBytes(caller)
	if bidder.Exists() {
		ctx.Log("Upped bid from: " + caller.String())
		bid := DecodeBidInfo(bidder.Value())
		bidAmount += bid.Amount
		bid.Amount = bidAmount
		bid.Timestamp = ctx.Timestamp()
		bidder.SetValue(EncodeBidInfo(bid))
	} else {
		if bidAmount < auction.MinimumBid {
			ctx.Panic("Insufficient bid amount")
		}
		ctx.Log("New bid from: " + caller.String())
		index := bidderList.Length()
		bidderList.GetAgent(index).SetValue(caller)
		bid := &BidInfo{
			Index:     int64(index),
			Amount:    bidAmount,
			Timestamp: ctx.Timestamp(),
		}
		bidder.SetValue(EncodeBidInfo(bid))
	}
	if bidAmount > auction.HighestBid {
		ctx.Log("New highest bidder")
		auction.HighestBid = bidAmount
		auction.HighestBidder = caller
		auctionInfo.SetValue(EncodeAuctionInfo(auction))
	}
}

func setOwnerMargin(ctx *client.ScCallContext) {
	// can only be sent by SC creator
	if !ctx.From(ctx.ContractCreator()) {
		ctx.Panic("Cancel spoofed request")
	}

	ownerMargin := ctx.Params().GetInt(KeyOwnerMargin).Value()
	if ownerMargin < OwnerMarginMin {
		ownerMargin = OwnerMarginMin
	}
	if ownerMargin > OwnerMarginMax {
		ownerMargin = OwnerMarginMax
	}
	ctx.State().GetInt(KeyOwnerMargin).SetValue(ownerMargin)
	ctx.Log("Updated owner margin")
}

func getInfo(ctx *client.ScViewContext) {
	colorParam := ctx.Params().GetColor(KeyColor)
	if !colorParam.Exists() {
		ctx.Panic("Missing token color")
	}
	color := colorParam.Value()

	state := ctx.State()
	auctions := state.GetMap(KeyAuctions)
	currentAuction := auctions.GetMap(color)
	auctionInfo := currentAuction.GetBytes(KeyInfo)
	if !auctionInfo.Exists() {
		ctx.Panic("Missing auction info")
	}

	auction := DecodeAuctionInfo(auctionInfo.Value())
	results := ctx.Results()
	results.GetColor(KeyColor).SetValue(auction.Color)
	results.GetAgent(KeyCreator).SetValue(auction.Creator)
	results.GetInt(KeyDeposit).SetValue(auction.Deposit)
	results.GetString(KeyDescription).SetValue(auction.Description)
	results.GetInt(KeyDuration).SetValue(auction.Duration)
	results.GetInt(KeyHighestBid).SetValue(auction.HighestBid)
	results.GetAgent(KeyHighestBidder).SetValue(auction.HighestBidder)
	results.GetInt(KeyMinimumBid).SetValue(auction.MinimumBid)
	results.GetInt(KeyNumTokens).SetValue(auction.NumTokens)
	results.GetInt(KeyOwnerMargin).SetValue(auction.OwnerMargin)
	results.GetInt(KeyWhenStarted).SetValue(auction.WhenStarted)

	bidderList := currentAuction.GetAgentArray(KeyBidderList)
	results.GetInt(KeyBidders).SetValue(int64(bidderList.Length()))
}

func transfer(ctx *client.ScCallContext, agent *client.ScAgent, color *client.ScColor, amount int64) {
	if agent.IsAddress() {
		// send back to original Tangle address
		ctx.TransferToAddress(agent.Address(), client.NewScTransfer(color, amount))
		return
	}

	// TODO not an address, deposit into account on chain
	ctx.TransferToAddress(agent.Address(), client.NewScTransfer(color, amount))
}
