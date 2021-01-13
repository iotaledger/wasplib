// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package fairauction

import "github.com/iotaledger/wasplib/client"

type AuctionInfo struct {
	Color         *client.ScColor // color of tokens for sale
	Creator       *client.ScAgent // issuer of start_auction transaction
	Deposit       int64           // deposit by auction owner to cover the SC fees
	Description   string          // auction description
	Duration      int64           // auction duration in minutes
	HighestBid    int64           // the current highest bid amount
	HighestBidder *client.ScAgent // the current highest bidder
	MinimumBid    int64           // minimum bid amount
	NumTokens     int64           // number of tokens for sale
	OwnerMargin   int64           // auction owner's margin in promilles
	WhenStarted   int64           // timestamp when auction started
}

type BidInfo struct {
	Amount    int64 // cumulative amount of bids from same bidder
	Index     int64 // index of bidder in bidder list
	Timestamp int64 // timestamp of most recent bid
}

func EncodeAuctionInfo(o *AuctionInfo) []byte {
	return client.NewBytesEncoder().
		Color(o.Color).
		Agent(o.Creator).
		Int(o.Deposit).
		String(o.Description).
		Int(o.Duration).
		Int(o.HighestBid).
		Agent(o.HighestBidder).
		Int(o.MinimumBid).
		Int(o.NumTokens).
		Int(o.OwnerMargin).
		Int(o.WhenStarted).
		Data()
}

func DecodeAuctionInfo(bytes []byte) *AuctionInfo {
	decode := client.NewBytesDecoder(bytes)
	data := &AuctionInfo{}
	data.Color = decode.Color()
	data.Creator = decode.Agent()
	data.Deposit = decode.Int()
	data.Description = decode.String()
	data.Duration = decode.Int()
	data.HighestBid = decode.Int()
	data.HighestBidder = decode.Agent()
	data.MinimumBid = decode.Int()
	data.NumTokens = decode.Int()
	data.OwnerMargin = decode.Int()
	data.WhenStarted = decode.Int()
	return data
}

func EncodeBidInfo(o *BidInfo) []byte {
	return client.NewBytesEncoder().
		Int(o.Amount).
		Int(o.Index).
		Int(o.Timestamp).
		Data()
}

func DecodeBidInfo(bytes []byte) *BidInfo {
	decode := client.NewBytesDecoder(bytes)
	data := &BidInfo{}
	data.Amount = decode.Int()
	data.Index = decode.Int()
	data.Timestamp = decode.Int()
	return data
}
