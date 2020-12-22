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
	d := client.NewBytesDecoder(bytes)
	data := &AuctionInfo{}
	data.auctionOwner = d.Agent()
	data.color = d.Color()
	data.deposit = d.Int()
	data.description = d.String()
	data.duration = d.Int()
	data.highestBid = d.Int()
	data.highestBidder = d.Agent()
	data.minimumBid = d.Int()
	data.numTokens = d.Int()
	data.ownerMargin = d.Int()
	data.whenStarted = d.Int()
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
	d := client.NewBytesDecoder(bytes)
	data := &BidInfo{}
	data.amount = d.Int()
	data.index = d.Int()
	data.timestamp = d.Int()
	return data
}
