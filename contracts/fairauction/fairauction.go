// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package fairauction

import (
	"github.com/iotaledger/wasplib/client"
)

const KeyAuctions = client.Key("auctions")
const KeyBidders = client.Key("bidders")
const KeyBidderList = client.Key("bidder_list")
const KeyColor = client.Key("color")
const KeyDescription = client.Key("description")
const KeyDuration = client.Key("duration")
const KeyInfo = client.Key("info")
const KeyMinimumBid = client.Key("minimum")
const KeyOwnerMargin = client.Key("owner_margin")

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
}

func startAuction(sc *client.ScCallContext) {
	deposit := sc.Incoming().Balance(client.IOTA)
	if deposit < 1 {
		sc.Log("Empty deposit...")
		return
	}

	state := sc.State()
	ownerMargin := state.GetInt(KeyOwnerMargin).Value()
	if ownerMargin == 0 {
		ownerMargin = OwnerMarginDefault
	}

	params := sc.Params()
	colorParam := params.GetColor(KeyColor)
	if !colorParam.Exists() {
		refund(sc, deposit / 2, "Missing token color...")
		return
	}
	color := colorParam.Value()

	if color == client.IOTA || color == client.MINT {
		refund(sc, deposit / 2, "Reserved token color...")
		return
	}

	numTokens := sc.Incoming().Balance(color)
	if numTokens == 0 {
		refund(sc, deposit / 2, "Auction tokens missing from request...")
		return
	}

	minimumBid := params.GetInt(KeyMinimumBid).Value()
	if minimumBid == 0 {
		refund(sc, deposit / 2, "Missing minimum bid...")
		return
	}

	// need at least 1 iota to run SC
	margin := minimumBid * ownerMargin / 1000
	if margin == 0 {
		margin = 1
	}
	if deposit < margin {
		refund(sc, deposit / 2, "Insufficient deposit...")
		return
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

	auctions := state.GetMap(KeyAuctions)
	currentAuction := auctions.GetMap(color)
	currentInfo := currentAuction.GetBytes(KeyInfo)
	if currentInfo.Exists() {
		refund(sc, deposit / 2, "Auction for this token already exists...")
		return
	}

	auction := &AuctionInfo {
		auctionOwner: sc.Caller(),
		color: color,
		deposit: deposit,
		description: description,
		duration: duration,
		highestBid: -1,
		highestBidder: &client.ScAgent{},
		minimumBid:minimumBid,
		numTokens: numTokens,
		ownerMargin: ownerMargin,
		whenStarted: sc.Timestamp(),
	}
	currentInfo.SetValue(encodeAuctionInfo(auction))

	finalizeRequest := sc.Post("finalize_auction")
	finalizeParams := finalizeRequest.Params()
	finalizeParams.GetColor(KeyColor).SetValue(auction.color)
	finalizeRequest.Post(duration * 60)
	sc.Log("New auction started...")
}

func finalizeAuction(sc *client.ScCallContext) {
	// can only be sent by SC itself
	if !sc.From(sc.Contract().Id()) {
		sc.Log("Cancel spoofed request")
		return
	}

	colorParam := sc.Params().GetColor(KeyColor)
	if !colorParam.Exists() {
		sc.Log("INTERNAL INCONSISTENCY: missing color")
		return
	}
	color := colorParam.Value()

	state := sc.State()
	auctions := state.GetMap(KeyAuctions)
	currentAuction := auctions.GetMap(color)
	currentInfo := currentAuction.GetBytes(KeyInfo)
	if !currentInfo.Exists() {
		sc.Log("INTERNAL INCONSISTENCY missing auction info")
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
		sc.Transfer(sc.Contract().Owner(), client.IOTA, ownerFee - 1)
		sc.Transfer(auction.auctionOwner, auction.color, auction.numTokens)
		sc.Transfer(auction.auctionOwner, client.IOTA, auction.deposit - ownerFee)
		return
	}

	ownerFee := auction.highestBid * auction.ownerMargin / 1000
	if ownerFee == 0 {
		ownerFee = 1
	}

	// return staked bids to losers
	bidders := currentAuction.GetMap(KeyBidders)
	bidderList := currentAuction.GetAgentArray(KeyBidderList)
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
	sc.Transfer(sc.Contract().Owner(), client.IOTA, ownerFee - 1)
	sc.Transfer(auction.highestBidder, auction.color, auction.numTokens)
	sc.Transfer(auction.auctionOwner, client.IOTA, auction.deposit + auction.highestBid - ownerFee)
}

func placeBid(sc *client.ScCallContext) {
	bidAmount := sc.Incoming().Balance(client.IOTA)
	if bidAmount == 0 {
		sc.Log("Insufficient bid amount")
		return
	}

	colorParam := sc.Params().GetColor(KeyColor)
	if !colorParam.Exists() {
		refund(sc, bidAmount, "Missing token color")
		return
	}
	color := colorParam.Value()

	state := sc.State()
	auctions := state.GetMap(KeyAuctions)
	currentAuction := auctions.GetMap(color)
	currentInfo := currentAuction.GetBytes(KeyInfo)
	if !currentInfo.Exists() {
		refund(sc, bidAmount, "Missing auction")
		return
	}

	bytes := currentInfo.Value()
	auction := decodeAuctionInfo(bytes)
	bidders := currentAuction.GetMap(KeyBidders)
	bidderList := currentAuction.GetAgentArray(KeyBidderList)
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
		bid := &BidInfo { index: int64(index), amount: bidAmount, timestamp: sc.Timestamp() }
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
	// can only be sent by SC owner
	if !sc.From(sc.Contract().Owner()) {
		sc.Log("Cancel spoofed request")
		return
	}

	ownerMargin := sc.Params().GetInt(KeyOwnerMargin).Value()
	if ownerMargin < OwnerMarginMin {
		ownerMargin = OwnerMarginMin
	}
	if ownerMargin > OwnerMarginMax {
		ownerMargin = OwnerMarginMax
	}
	sc.State().GetInt(KeyOwnerMargin).SetValue(ownerMargin)
	sc.Log("Updated owner margin...")
}

func refund(sc *client.ScCallContext, amount int64, reason string) {
	sc.Log(reason)
	caller := sc.Caller()
	if amount != 0 {
		sc.Transfer(caller, client.IOTA, amount)
	}
	deposit := sc.Incoming().Balance(client.IOTA)
	if deposit - amount != 0 {
		sc.Transfer(sc.Contract().Owner(), client.IOTA, deposit - amount)
	}

	//TODO iterate over sc.Incoming() balances
	// // refund all other token colors, don't keep tokens that were to be auctioned
	// colors := request.colors()
	// items := colors.Length()
	// for i in 0..items {
	//     color := colors.GetColor(i).Value()
	//     if color != client.IOTA {
	//         sc.Transfer(caller, &color, sc.Incoming().balance(color))
	//     }
	// }
}
