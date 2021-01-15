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

func startAuction(sc *client.ScCallContext) {
	params := sc.Params()
	colorParam := params.GetColor(KeyColor)
	if !colorParam.Exists() {
		sc.Panic("Missing auction token color")
	}
	color := colorParam.Value()
	if color.Equals(client.IOTA) || color.Equals(client.MINT) {
		sc.Panic("Reserved auction token color")
	}
	numTokens := sc.Incoming().Balance(color)
	if numTokens == 0 {
		sc.Panic("Missing auction tokens")
	}

	minimumBid := params.GetInt(KeyMinimumBid).Value()
	if minimumBid == 0 {
		sc.Panic("Missing minimum bid")
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

	state := sc.State()
	ownerMargin := state.GetInt(KeyOwnerMargin).Value()
	if ownerMargin == 0 {
		ownerMargin = OwnerMarginDefault
	}

	// need at least 1 iota to run SC
	margin := minimumBid * ownerMargin / 1000
	if margin == 0 {
		margin = 1
	}
	deposit := sc.Incoming().Balance(client.IOTA)
	if deposit < margin {
		sc.Panic("Insufficient deposit")
	}

	auctions := state.GetMap(KeyAuctions)
	currentAuction := auctions.GetMap(color)
	auctionInfo := currentAuction.GetBytes(KeyInfo)
	if auctionInfo.Exists() {
		sc.Panic("Auction for this token color already exists")
	}

	auction := &AuctionInfo {
		Creator: sc.Caller(),
		Color: color,
		Deposit: deposit,
		Description: description,
		Duration: duration,
		HighestBid: -1,
		HighestBidder: &client.ScAgent{},
		MinimumBid: minimumBid,
		NumTokens: numTokens,
		OwnerMargin: ownerMargin,
		WhenStarted: sc.Timestamp(),
	}
	auctionInfo.SetValue(EncodeAuctionInfo(auction))

	finalizeRequest := sc.Post("finalize_auction")
	finalizeParams := finalizeRequest.Params()
	finalizeParams.GetColor(KeyColor).SetValue(auction.Color)
	finalizeRequest.Post(duration * 60)
	sc.Log("New auction started")
}

func finalizeAuction(sc *client.ScCallContext) {
	// can only be sent by SC itself
	if !sc.From(sc.ContractId()) {
		sc.Panic("Cancel spoofed request")
	}

	colorParam := sc.Params().GetColor(KeyColor)
	if !colorParam.Exists() {
		sc.Panic("Missing token color")
	}
	color := colorParam.Value()

	state := sc.State()
	auctions := state.GetMap(KeyAuctions)
	currentAuction := auctions.GetMap(color)
	auctionInfo := currentAuction.GetBytes(KeyInfo)
	if !auctionInfo.Exists() {
		sc.Panic("Missing auction info")
	}
	auction := DecodeAuctionInfo(auctionInfo.Value())
	if auction.HighestBid < 0 {
		sc.Log("No one bid on " + color.String())
		ownerFee := auction.MinimumBid * auction.OwnerMargin / 1000
		if ownerFee == 0 {
			ownerFee = 1
		}
		// finalizeAuction request token was probably not confirmed yet
		transfer(sc, sc.ContractCreator(), client.IOTA, ownerFee - 1)
		transfer(sc, auction.Creator, auction.Color, auction.NumTokens)
		transfer(sc, auction.Creator, client.IOTA, auction.Deposit - ownerFee)
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
			transfer(sc, bidder, client.IOTA, bid.Amount)
		}
	}

	// finalizeAuction request token was probably not confirmed yet
	transfer(sc, sc.ContractCreator(), client.IOTA, ownerFee-1)
	transfer(sc, auction.HighestBidder, auction.Color, auction.NumTokens)
	transfer(sc, auction.Creator, client.IOTA, auction.Deposit+auction.HighestBid-ownerFee)
}

func placeBid(sc *client.ScCallContext) {
	bidAmount := sc.Incoming().Balance(client.IOTA)
	if bidAmount == 0 {
		sc.Panic("Missing bid amount")
	}

	colorParam := sc.Params().GetColor(KeyColor)
	if !colorParam.Exists() {
		sc.Panic("Missing token color")
	}
	color := colorParam.Value()

	state := sc.State()
	auctions := state.GetMap(KeyAuctions)
	currentAuction := auctions.GetMap(color)
	auctionInfo := currentAuction.GetBytes(KeyInfo)
	if !auctionInfo.Exists() {
		sc.Panic("Missing auction info")
	}

	auction := DecodeAuctionInfo(auctionInfo.Value())
	bidders := currentAuction.GetMap(KeyBidders)
	bidderList := currentAuction.GetAgentArray(KeyBidderList)
	caller := sc.Caller()
	bidder := bidders.GetBytes(caller)
	if bidder.Exists() {
		sc.Log("Upped bid from: " + caller.String())
		bid := DecodeBidInfo(bidder.Value())
		bidAmount += bid.Amount
		bid.Amount = bidAmount
		bid.Timestamp = sc.Timestamp()
		bidder.SetValue(EncodeBidInfo(bid))
	} else {
		if bidAmount < auction.MinimumBid {
			sc.Panic("Insufficient bid amount")
		}
		sc.Log("New bid from: " + caller.String())
		index := bidderList.Length()
		bidderList.GetAgent(index).SetValue(caller)
		bid := &BidInfo {
			Index: int64(index),
			Amount: bidAmount,
			Timestamp: sc.Timestamp(),
		}
		bidder.SetValue(EncodeBidInfo(bid))
	}
	if bidAmount > auction.HighestBid {
		sc.Log("New highest bidder")
		auction.HighestBid = bidAmount
		auction.HighestBidder = caller
		auctionInfo.SetValue(EncodeAuctionInfo(auction))
	}
}

func setOwnerMargin(sc *client.ScCallContext) {
	// can only be sent by SC creator
	if !sc.From(sc.ContractCreator()) {
		sc.Panic("Cancel spoofed request")
	}

	ownerMargin := sc.Params().GetInt(KeyOwnerMargin).Value()
	if ownerMargin < OwnerMarginMin {
		ownerMargin = OwnerMarginMin
	}
	if ownerMargin > OwnerMarginMax {
		ownerMargin = OwnerMarginMax
	}
	sc.State().GetInt(KeyOwnerMargin).SetValue(ownerMargin)
	sc.Log("Updated owner margin")
}

func getInfo(sc *client.ScViewContext) {
	colorParam := sc.Params().GetColor(KeyColor)
	if !colorParam.Exists() {
		sc.Panic("Missing token color")
	}
	color := colorParam.Value()

	state := sc.State()
	auctions := state.GetMap(KeyAuctions)
	currentAuction := auctions.GetMap(color)
	auctionInfo := currentAuction.GetBytes(KeyInfo)
	if !auctionInfo.Exists() {
		sc.Panic("Missing auction info")
	}

	auction := DecodeAuctionInfo(auctionInfo.Value())
	results := sc.Results()
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

func transfer(sc *client.ScCallContext, agent *client.ScAgent, color *client.ScColor, amount int64) {
	if ! agent.IsAddress() {
		// not an address, deposit into account on chain
		sc.Transfer(agent, color, amount)
		return
	}

	// send to original Tangle address
	sc.TransferToAddress(agent.Address()).Transfer(color, amount).Send()
}
