// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package fairauction

import "github.com/iotaledger/wasplib/client"

const keyAuctions = client.Key("auctions")
const keyBidders = client.Key("bidders")
const keyBidderList = client.Key("bidder_list")
const keyColor = client.Key("color")
const keyDescription = client.Key("description")
const keyDuration = client.Key("duration")
const keyInfo = client.Key("info")
const keyMinimumBid = client.Key("minimum")
const keyOwnerMargin = client.Key("owner_margin")

const durationDefault = 60
const durationMin = 1
const durationMax = 120
const maxDescriptionLength = 150
const ownerMarginDefault = 50
const ownerMarginMin = 5
const ownerMarginMax = 100

func OnLoad() {
	exports := client.NewScExports()
	exports.AddCall("start_auction", startAuction)
	exports.AddCall("finalize_auction", finalizeAuction)
	exports.AddCall("place_bid", placeBid)
	exports.AddCall("set_owner_margin", setOwnerMargin)
}

func startAuction(sc *client.ScCallContext) {
	deposit := sc.Incoming().Balance(client.IOTA)
	if deposit < 1 {
		sc.Log("Empty deposit...")
		return
	}

	state := sc.State()
	ownerMargin := state.GetInt(keyOwnerMargin).Value()
	if ownerMargin == 0 {
		ownerMargin = ownerMarginDefault
	}

	params := sc.Params()
	colorParam := params.GetColor(keyColor)
	if !colorParam.Exists() {
		refund(sc, deposit/2, "Missing token color...")
		return
	}
	color := colorParam.Value()

	if color == client.IOTA || color == client.MINT {
		refund(sc, deposit/2, "Reserved token color...")
		return
	}

	numTokens := sc.Incoming().Balance(color)
	if numTokens == 0 {
		refund(sc, deposit/2, "Auction tokens missing from request...")
		return
	}

	minimumBid := params.GetInt(keyMinimumBid).Value()
	if minimumBid == 0 {
		refund(sc, deposit/2, "Missing minimum bid...")
		return
	}

	// need at least 1 iota to run SC
	margin := minimumBid * ownerMargin / 1000
	if margin == 0 {
		margin = 1
	}
	if deposit < margin {
		refund(sc, deposit/2, "Insufficient deposit...")
		return
	}

	// duration in minutes
	duration := params.GetInt(keyDuration).Value()
	if duration == 0 {
		duration = durationDefault
	}
	if duration < durationMin {
		duration = durationMin
	}
	if duration > durationMax {
		duration = durationMax
	}

	description := params.GetString(keyDescription).Value()
	if description == "" {
		description = "N/A"
	}
	if len(description) > maxDescriptionLength {
		description = description[:maxDescriptionLength] + "[...]"
	}

	auctions := state.GetMap(keyAuctions)
	currentAuction := auctions.GetMap(color)
	currentInfo := currentAuction.GetBytes(keyInfo)
	if currentInfo.Exists() {
		refund(sc, deposit/2, "Auction for this token already exists...")
		return
	}

	auction := &AuctionInfo{
		auctionOwner:  sc.Caller(),
		color:         color,
		deposit:       deposit,
		description:   description,
		duration:      duration,
		highestBid:    -1,
		highestBidder: &client.ScAgent{},
		minimumBid:    minimumBid,
		numTokens:     numTokens,
		ownerMargin:   ownerMargin,
		whenStarted:   sc.Timestamp(),
	}
	currentInfo.SetValue(encodeAuctionInfo(auction))

	finalizeRequest := sc.Post("finalize_auction")
	finalizeParams := finalizeRequest.Params()
	finalizeParams.GetColor(keyColor).SetValue(auction.color)
	finalizeRequest.Post(duration * 60)
	sc.Log("New auction started...")
}

func finalizeAuction(sc *client.ScCallContext) {
	// can only be sent by SC itself
	if !sc.From(sc.Contract().Id()) {
		sc.Log("Cancel spoofed request")
		return
	}

	colorParam := sc.Params().GetColor(keyColor)
	if !colorParam.Exists() {
		sc.Log("Internal inconsistency: missing color")
		return
	}
	color := colorParam.Value()

	state := sc.State()
	auctions := state.GetMap(keyAuctions)
	currentAuction := auctions.GetMap(color)
	currentInfo := currentAuction.GetBytes(keyInfo)
	if !currentInfo.Exists() {
		sc.Log("Internal inconsistency: missing auction info")
		return
	}
	auction := decodeAuctionInfo(currentInfo.Value())
	if auction.highestBid < 0 {
		sc.Log("No one bid on " + color.String())
		ownerFee := auction.minimumBid * auction.ownerMargin / 1000
		if ownerFee == 0 {
			ownerFee = 1
		}
		// finalizeAuction request token was probably not confirmed yet
		sc.Transfer(sc.Contract().Creator(), client.IOTA, ownerFee-1)
		sc.Transfer(auction.auctionOwner, auction.color, auction.numTokens)
		sc.Transfer(auction.auctionOwner, client.IOTA, auction.deposit-ownerFee)
		return
	}

	ownerFee := auction.highestBid * auction.ownerMargin / 1000
	if ownerFee == 0 {
		ownerFee = 1
	}

	// return staked bids to losers
	bidders := currentAuction.GetMap(keyBidders)
	bidderList := currentAuction.GetAgentArray(keyBidderList)
	size := bidderList.Length()
	for i := int32(0); i < size; i++ {
		bidder := bidderList.GetAgent(i).Value()
		if !bidder.Equals(auction.highestBidder) {
			loser := bidders.GetBytes(bidder)
			bid := decodeBidInfo(loser.Value())
			sc.Transfer(bidder, client.IOTA, bid.amount)
		}
	}

	// finalizeAuction request token was probably not confirmed yet
	sc.Transfer(sc.Contract().Creator(), client.IOTA, ownerFee-1)
	sc.Transfer(auction.highestBidder, auction.color, auction.numTokens)
	sc.Transfer(auction.auctionOwner, client.IOTA, auction.deposit+auction.highestBid-ownerFee)
}

func placeBid(sc *client.ScCallContext) {
	bidAmount := sc.Incoming().Balance(client.IOTA)
	if bidAmount == 0 {
		sc.Log("Insufficient bid amount")
		return
	}

	colorParam := sc.Params().GetColor(keyColor)
	if !colorParam.Exists() {
		refund(sc, bidAmount, "Missing token color")
		return
	}
	color := colorParam.Value()

	state := sc.State()
	auctions := state.GetMap(keyAuctions)
	currentAuction := auctions.GetMap(color)
	currentInfo := currentAuction.GetBytes(keyInfo)
	if !currentInfo.Exists() {
		refund(sc, bidAmount, "Missing auction")
		return
	}

	auction := decodeAuctionInfo(currentInfo.Value())
	bidders := currentAuction.GetMap(keyBidders)
	bidderList := currentAuction.GetAgentArray(keyBidderList)
	caller := sc.Caller()
	bidder := bidders.GetBytes(caller)
	if bidder.Exists() {
		sc.Log("Upped bid from: " + caller.String())
		bid := decodeBidInfo(bidder.Value())
		bidAmount += bid.amount
		bid.amount = bidAmount
		bid.timestamp = sc.Timestamp()
		bidder.SetValue(encodeBidInfo(bid))
	} else {
		sc.Log("New bid from: " + caller.String())
		index := bidderList.Length()
		bidderList.GetAgent(index).SetValue(caller)
		bid := &BidInfo{
			index:     int64(index),
			amount:    bidAmount,
			timestamp: sc.Timestamp(),
		}
		bidder.SetValue(encodeBidInfo(bid))
	}
	if bidAmount > auction.highestBid {
		sc.Log("New highest bidder...")
		auction.highestBid = bidAmount
		auction.highestBidder = caller
		currentInfo.SetValue(encodeAuctionInfo(auction))
	}
}

func setOwnerMargin(sc *client.ScCallContext) {
	// can only be sent by SC creator
	if !sc.From(sc.Contract().Creator()) {
		sc.Log("Cancel spoofed request")
		return
	}

	ownerMargin := sc.Params().GetInt(keyOwnerMargin).Value()
	if ownerMargin < ownerMarginMin {
		ownerMargin = ownerMarginMin
	}
	if ownerMargin > ownerMarginMax {
		ownerMargin = ownerMarginMax
	}
	sc.State().GetInt(keyOwnerMargin).SetValue(ownerMargin)
	sc.Log("Updated owner margin...")
}

func refund(sc *client.ScCallContext, amount int64, reason string) {
	sc.Log(reason)
	caller := sc.Caller()
	if amount != 0 {
		sc.Transfer(caller, client.IOTA, amount)
	}
	incoming := sc.Incoming()
	deposit := incoming.Balance(client.IOTA)
	if deposit-amount != 0 {
		sc.Transfer(sc.Contract().Creator(), client.IOTA, deposit-amount)
	}

	// refund all other token colors, don't keep tokens that were to be auctioned
	colors := incoming.Colors()
	size := colors.Length()
	for i := int32(0); i < size; i++ {
		color := colors.GetColor(i).Value()
		if color != client.IOTA {
			sc.Transfer(caller, color, incoming.Balance(color))
		}
	}
}
