package main

import (
	"github.com/iotaledger/wasplib/client"
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
	color string
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
	auctionOwner string
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
	address string
	// the amount is a cumulative sum of all bids from the same bidder
	amount int64
	// most recent bid update time
	when int64
}

func main() {
}

//export onLoad
func onLoadFairAuction() {
	exports := client.NewScExports()
	exports.Add("startAuction")
	exports.Add("finalizeAuction")
	exports.Add("placeBid")
	exports.AddProtected("setOwnerMargin")
}

//export startAuction
func startAuction() {
	sc := client.NewScContext()
	request := sc.Request()
	deposit := request.Balance("iota")
	if deposit < 1 {
		sc.Log("Empty deposit...")
		return
	}

	state := sc.State()
	ownerMargin := state.GetInt("ownerMargin").Value()
	if ownerMargin == 0 {
		ownerMargin = ownerMarginDefault
	}

	params := request.Params()
	color := params.GetString("color").Value()
	if color == "" {
		refund(deposit/2, "Missing token color...")
		return
	}
	//TODO check if valid base58 string
	if color == "InvalidBase58" {
		refund(deposit/2, "Invalid token color...")
		return
	}
	//TODO determine appropriate color hashes
	if color == "iota" || color == "new" {
		refund(deposit/2, "Reserved token color...")
		return
	}

	numTokens := request.Balance(color)
	if numTokens == 0 {
		refund(deposit/2, "Auction tokens missing from request...")
		return
	}

	minimumBid := params.GetInt("minimum").Value()
	if minimumBid == 0 {
		refund(deposit/2, "Missing minimum bid...")
		return
	}

	// need at least 1 iota to run SC
	margin := minimumBid * ownerMargin / 1000
	if margin == 0 {
		margin = 1
	}
	if deposit < margin {
		refund(deposit/2, "Insufficient deposit...")
		return
	}

	// duration in minutes
	duration := params.GetInt("duration").Value()
	if duration == 0 {
		duration = durationDefault
	}
	if duration < durationMin {
		duration = durationMin
	}
	if duration > durationMax {
		duration = durationMax
	}

	description := params.GetString("dscr").Value()
	if description == "" {
		description = "N/A"
	}
	if len(description) > maxDescriptionLength {
		description = description[:maxDescriptionLength] + "[...]"
	}

	auctions := state.GetMap("auctions")
	currentAuction := auctions.GetBytes(color)
	if len(currentAuction.Value()) != 0 {
		refund(deposit/2, "Auction for this token already exists...")
		return
	}

	auction := &AuctionInfo{
		color:        color,
		numTokens:    numTokens,
		minimumBid:   minimumBid,
		description:  description,
		whenStarted:  request.Timestamp(),
		duration:     duration,
		auctionOwner: request.Address(),
		deposit:      deposit,
		ownerMargin:  ownerMargin,
	}
	bytes := encodeAuctionInfo(auction)
	currentAuction.SetValue(bytes)

	finalizeParams := sc.Event("", "finalizeAuction", auction.duration*60)
	finalizeParams.GetString("color").SetValue(auction.color)
	sc.Log("New auction started...")
}

//export finalizeAuction
func finalizeAuction() {
	// can only be sent by SC itself
	sc := client.NewScContext()
	request := sc.Request()
	if request.Address() != sc.Contract().Address() {
		sc.Log("Cancel spoofed request")
		return
	}

	color := request.Params().GetString("color").Value()
	if color == "" {
		sc.Log("INTERNAL INCONSISTENCY: missing color")
		return
	}

	state := sc.State()
	auctions := state.GetMap("auctions")
	currentAuction := auctions.GetBytes(color)
	bytes := currentAuction.Value()
	if len(bytes) == 0 {
		sc.Log("INTERNAL INCONSISTENCY missing auction info")
		return
	}
	auction := decodeAuctionInfo(bytes)
	if len(auction.bids) == 0 {
		sc.Log("No one bid on " + color)
		ownerFee := auction.minimumBid * auction.ownerMargin / 1000
		if ownerFee == 0 {
			ownerFee = 1
		}
		// finalizeAuction request token was probably not confirmed yet
		sc.Transfer(sc.Contract().Owner(), "iota", ownerFee - 1)
		sc.Transfer(auction.auctionOwner, "iota", auction.deposit-ownerFee)
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
			sc.Transfer(bidder.address, "iota", bidder.amount)
		}
	}

	// finalizeAuction request token was probably not confirmed yet
	sc.Transfer(sc.Contract().Owner(), "iota", ownerFee - 1)
	sc.Transfer(winner.address, auction.color, auction.numTokens)
	sc.Transfer(auction.auctionOwner, "iota", auction.deposit+winner.amount-ownerFee)
}

//export placeBid
func placeBid() {
	sc := client.NewScContext()
	request := sc.Request()
	bidAmount := request.Balance("iota")
	if bidAmount == 0 {
		sc.Log("Insufficient bid amount")
		return
	}

	color := request.Params().GetString("color").Value()
	if color == "" {
		refund(bidAmount, "Missing token color")
		return
	}

	state := sc.State()
	auctions := state.GetMap("auctions")
	currentAuction := auctions.GetBytes(color)
	bytes := currentAuction.Value()
	if len(bytes) == 0 {
		refund(bidAmount, "Missing auction")
		return
	}

	sender := request.Address()
	auction := decodeAuctionInfo(bytes)
	var bid *BidInfo
	for _, bidder := range auction.bids {
		if bidder.address == sender {
			bid = bidder
			break
		}
	}
	if bid == nil {
		sc.Log("New bid from: " + sender)
		bid = &BidInfo{address: sender}
		auction.bids = append(auction.bids, bid)
	}
	bid.amount += bidAmount
	bid.when = request.Timestamp()

	bytes = encodeAuctionInfo(auction)
	currentAuction.SetValue(bytes)
	sc.Log("Updated auction with bid...")
}

//export setOwnerMargin
func setOwnerMargin() {
	// can only be sent by SC owner
	sc := client.NewScContext()
	if sc.Request().Address() != sc.Contract().Owner() {
		sc.Log("Cancel spoofed request")
		return
	}

	ownerMargin := sc.Request().Params().GetInt("ownerMargin").Value()
	if ownerMargin < ownerMarginMin {
		ownerMargin = ownerMarginMin
	}
	if ownerMargin > ownerMarginMax {
		ownerMargin = ownerMarginMax
	}
	sc.State().GetInt("ownerMargin").SetValue(ownerMargin)
	sc.Log("Updated owner margin...")
}

func decodeAuctionInfo(bytes []byte) *AuctionInfo {
	decoder := client.NewBytesDecoder(bytes)
	auction := &AuctionInfo{
		color:        decoder.String(),
		numTokens:    decoder.Int(),
		minimumBid:   decoder.Int(),
		description:  decoder.String(),
		whenStarted:  decoder.Int(),
		duration:     decoder.Int(),
		auctionOwner: decoder.String(),
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
		address: decoder.String(),
		amount:  decoder.Int(),
		when:    decoder.Int(),
	}
}

func encodeAuctionInfo(auction *AuctionInfo) []byte {
	encoder := client.NewBytesEncoder().
		String(auction.color).
		Int(auction.numTokens).
		Int(auction.minimumBid).
		String(auction.description).
		Int(auction.whenStarted).
		Int(auction.duration).
		String(auction.auctionOwner).
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
		String(bid.address).
		Int(bid.amount).
		Int(bid.when).
		Data()
}

func refund(amount int64, reason string) {
	sc := client.NewScContext()
	sc.Log(reason)
	request := sc.Request()
	sender := request.Address()
	if amount != 0 {
		sc.Transfer(sender, "iota", amount)
	}
	deposit := request.Balance("iota")
	if deposit-amount != 0 {
		sc.Transfer(sc.Contract().Owner(), "iota", deposit-amount)
	}

	// refund all other token colors, don't keep tokens that were to be auctioned
	colors := request.Colors()
	items := colors.Length()
	for i := int32(0); i < items; i++ {
		color := colors.GetString(i).Value()
		if color != "iota" {
			sc.Transfer(sender, color, request.Balance(color))
		}
	}
}
