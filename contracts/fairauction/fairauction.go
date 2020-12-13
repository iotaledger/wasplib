// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package fairauction

import (
	"github.com/iotaledger/wasplib/client"
)

const (
	keyAuctions    = client.Key("auctions")
	keyColor       = client.Key("color")
	keyDescription = client.Key("description")
	keyDuration    = client.Key("duration")
	keyMinimum     = client.Key("minimum")
	keyOwnerMargin = client.Key("owner_margin")
)

const (
	durationDefault      = 60
	durationMin          = 1
	durationMax          = 120
	maxDescriptionLength = 150
	ownerMarginDefault   = 50
	ownerMarginMin       = 5
	ownerMarginMax       = 100
)

type AuctionInfo struct {
	// color of tokens for sale
	color *client.ScColor
	// number of tokens for sale
	numTokens int64
	// minimum bid. Set by the auction initiator
	minimumBid int64
	// any text, like "AuctionOwner of the token have a right to call me for a date". Set by auction initiator
	description string
	// timestamp when auction started
	whenStarted int64
	// duration of the auctions in minutes. Should be >= MinAuctionDurationMinutes
	duration int64
	// address which issued StartAuction transaction
	auctionOwner *client.ScAgent
	// deposit by the auction owner. Iotas sent by the auction owner together with the tokens for sale in the same
	// transaction.
	deposit int64
	// AuctionOwner's margin in promilles, taken at the moment of creation of smart contract
	ownerMargin int64
	// list of bids to the auction
	bids []*BidInfo
}

type BidInfo struct {
	// originator of the bid
	bidder *client.ScAgent
	// the amount is a cumulative sum of all bids from the same bidder
	amount int64
	// most recent bid update time
	when int64
}

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

	if color.Equals(client.IOTA) || color.Equals(client.MINT) {
		refund(sc, deposit/2, "Reserved token color...")
		return
	}

	numTokens := sc.Incoming().Balance(color)
	if numTokens == 0 {
		refund(sc, deposit/2, "Auction tokens missing from request...")
		return
	}

	minimumBid := params.GetInt(keyMinimum).Value()
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
	if len(description) == 0 {
		description = "N/A"
	}
	if len(description) > maxDescriptionLength {
		description = description[:maxDescriptionLength] + "[...]"
	}

	auctions := state.GetMap(keyAuctions)
	currentAuction := auctions.GetBytes(color)
	if len(currentAuction.Value()) != 0 {
		refund(sc, deposit/2, "Auction for this token already exists...")
		return
	}

	auction := &AuctionInfo{
		color:        color,
		numTokens:    numTokens,
		minimumBid:   minimumBid,
		description:  description,
		whenStarted:  sc.Timestamp(),
		duration:     duration,
		auctionOwner: sc.Caller(),
		deposit:      deposit,
		ownerMargin:  ownerMargin,
	}
	bytes := encodeAuctionInfo(auction)
	currentAuction.SetValue(bytes)
	finalizeRequest := sc.Post("finalize_auction")
	finalizeParams := finalizeRequest.Params()
	finalizeParams.GetColor(keyColor).SetValue(auction.color)
	finalizeRequest.Post(auction.duration * 60)
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
		sc.Log("INTERNAL INCONSISTENCY: missing color")
		return
	}
	color := colorParam.Value()

	state := sc.State()
	auctions := state.GetMap(keyAuctions)
	currentAuction := auctions.GetBytes(color)
	bytes := currentAuction.Value()
	if len(bytes) == 0 {
		sc.Log("INTERNAL INCONSISTENCY missing auction info")
		return
	}
	auction := decodeAuctionInfo(bytes)
	if len(auction.bids) == 0 {
		sc.Log("No one bid on " + color.String())
		ownerFee := auction.minimumBid * auction.ownerMargin / 1000
		if ownerFee == 0 {
			ownerFee = 1
		}
		// finalizeAuction request token was probably not confirmed yet
		sc.Transfer(sc.Contract().Owner(), client.IOTA, ownerFee-1)
		sc.Transfer(auction.auctionOwner, auction.color, auction.numTokens)
		sc.Transfer(auction.auctionOwner, client.IOTA, auction.deposit-ownerFee)
		return
	}

	winner := &BidInfo{}
	for _, bidder := range auction.bids {
		if bidder.amount >= winner.amount {
			if bidder.amount > winner.amount || bidder.when < winner.when {
				winner = bidder
			}
		}
	}
	ownerFee := winner.amount * auction.ownerMargin / 1000
	if ownerFee == 0 {
		ownerFee = 1
	}

	// return staked bids to losers
	for _, bidder := range auction.bids {
		if bidder != winner {
			sc.Transfer(bidder.bidder, client.IOTA, bidder.amount)
		}
	}

	// finalizeAuction request token was probably not confirmed yet
	sc.Transfer(sc.Contract().Owner(), client.IOTA, ownerFee-1)
	sc.Transfer(winner.bidder, auction.color, auction.numTokens)
	sc.Transfer(auction.auctionOwner, client.IOTA, auction.deposit+winner.amount-ownerFee)
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
	currentAuction := auctions.GetBytes(color)
	bytes := currentAuction.Value()
	if len(bytes) == 0 {
		refund(sc, bidAmount, "Missing auction")
		return
	}

	caller := sc.Caller()
	auction := decodeAuctionInfo(bytes)
	var bid *BidInfo
	for _, bidder := range auction.bids {
		if bidder.bidder.Equals(caller) {
			bid = bidder
			break
		}
	}
	if bid == nil {
		sc.Log("New bid from: " + caller.String())
		bid = &BidInfo{bidder: caller}
		auction.bids = append(auction.bids, bid)
	}
	bid.amount += bidAmount
	bid.when = sc.Timestamp()

	bytes = encodeAuctionInfo(auction)
	currentAuction.SetValue(bytes)
	sc.Log("Updated auction with bid...")
}

func setOwnerMargin(sc *client.ScCallContext) {
	// can only be sent by SC owner
	if !sc.From(sc.Contract().Owner()) {
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

func decodeAuctionInfo(bytes []byte) *AuctionInfo {
	decoder := client.NewBytesDecoder(bytes)
	auction := &AuctionInfo{
		color:        decoder.Color(),
		numTokens:    decoder.Int(),
		minimumBid:   decoder.Int(),
		description:  decoder.String(),
		whenStarted:  decoder.Int(),
		duration:     decoder.Int(),
		auctionOwner: decoder.Agent(),
		deposit:      decoder.Int(),
		ownerMargin:  decoder.Int(),
	}
	bids := int(decoder.Int())
	for i := 0; i < bids; i++ {
		bytes = decoder.Bytes()
		bid := decodeBidInfo(bytes)
		auction.bids = append(auction.bids, bid)
	}
	return auction
}

func decodeBidInfo(bytes []byte) *BidInfo {
	decoder := client.NewBytesDecoder(bytes)
	return &BidInfo{
		bidder: decoder.Agent(),
		amount: decoder.Int(),
		when:   decoder.Int(),
	}
}

func encodeAuctionInfo(auction *AuctionInfo) []byte {
	encoder := client.NewBytesEncoder().
		Color(auction.color).
		Int(auction.numTokens).
		Int(auction.minimumBid).
		String(auction.description).
		Int(auction.whenStarted).
		Int(auction.duration).
		Agent(auction.auctionOwner).
		Int(auction.deposit).
		Int(auction.ownerMargin).
		Int(int64(len(auction.bids)))
	for _, bid := range auction.bids {
		bytes := encodeBidInfo(bid)
		encoder.Bytes(bytes)
	}
	return encoder.Data()
}

func encodeBidInfo(bid *BidInfo) []byte {
	return client.NewBytesEncoder().
		Agent(bid.bidder).
		Int(bid.amount).
		Int(bid.when).
		Data()
}

func refund(sc *client.ScCallContext, amount int64, reason string) {
	sc.Log(reason)
	caller := sc.Caller()
	if amount != 0 {
		sc.Transfer(caller, client.IOTA, amount)
	}
	deposit := sc.Incoming().Balance(client.IOTA)
	if deposit-amount != 0 {
		sc.Transfer(sc.Contract().Owner(), client.IOTA, deposit-amount)
	}

	//TODO
	//// refund all other token colors, don't keep tokens that were to be auctioned
	//colors := request.Colors()
	//items := colors.Length()
	//for i := int32(0); i < items; i++ {
	//	color := colors.GetColor(i).Value()
	//	if !color.Equals(client.IOTA) {
	//		sc.Transfer(caller, color, request.Balance(color))
	//	}
	//}
}
