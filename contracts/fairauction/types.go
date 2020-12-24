// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package fairauction

import "github.com/iotaledger/wasplib/client"

type AuctionInfo struct {
	auctionOwner  *client.ScAgent // issuer of start_auction transaction
	color         *client.ScColor // color of tokens for sale
	deposit       int64           // deposit by auction owner to cover the SC fees
	description   string          // auction description
	duration      int64           // auction duration in minutes
	highestBid    int64           // the current highest bid amount
	highestBidder *client.ScAgent // the current highest bidder
	minimumBid    int64           // minimum bid amount
	numTokens     int64           // number of tokens for sale
	ownerMargin   int64           // auction owner's margin in promilles
	whenStarted   int64           // timestamp when auction started
}

type BidInfo struct {
	amount    int64 // cumulative amount of bids from same bidder
	index     int64 // index of bidder in bidder list
	timestamp int64 // timestamp of most recent bid
}

func encodeAuctionInfo(o *AuctionInfo) []byte {
	return client.NewBytesEncoder().
		Agent(o.auctionOwner).
		Color(o.color).
		Int(o.deposit).
		String(o.description).
		Int(o.duration).
		Int(o.highestBid).
		Agent(o.highestBidder).
		Int(o.minimumBid).
		Int(o.numTokens).
		Int(o.ownerMargin).
		Int(o.whenStarted).
		Data()
}

func decodeAuctionInfo(bytes []byte) *AuctionInfo {
	decode := client.NewBytesDecoder(bytes)
	data := &AuctionInfo{}
	data.auctionOwner = decode.Agent()
	data.color = decode.Color()
	data.deposit = decode.Int()
	data.description = decode.String()
	data.duration = decode.Int()
	data.highestBid = decode.Int()
	data.highestBidder = decode.Agent()
	data.minimumBid = decode.Int()
	data.numTokens = decode.Int()
	data.ownerMargin = decode.Int()
	data.whenStarted = decode.Int()
	return data
}

func encodeBidInfo(o *BidInfo) []byte {
	return client.NewBytesEncoder().
		Int(o.amount).
		Int(o.index).
		Int(o.timestamp).
		Data()
}

func decodeBidInfo(bytes []byte) *BidInfo {
	decode := client.NewBytesDecoder(bytes)
	data := &BidInfo{}
	data.amount = decode.Int()
	data.index = decode.Int()
	data.timestamp = decode.Int()
	return data
}
