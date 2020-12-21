// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package fairauction

import "github.com/iotaledger/wasplib/client"

type AuctionInfo struct {
    auctionOwner *client.ScAgent
    color *client.ScColor
    deposit int64
    description string
    duration int64
    highestBid int64
    highestBidder *client.ScAgent
    minimumBid int64
    numTokens int64
    ownerMargin int64
    whenStarted int64
}

type BidInfo struct {
    amount int64
    index int64
    timestamp int64
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
