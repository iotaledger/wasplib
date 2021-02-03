// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use wasplib::client::*;

pub struct AuctionInfo {
    //@formatter:off
    pub color:          ScColor,   // color of tokens for sale
    pub creator:        ScAgentId, // issuer of start_auction transaction
    pub deposit:        i64,       // deposit by auction owner to cover the SC fees
    pub description:    String,    // auction description
    pub duration:       i64,       // auction duration in minutes
    pub highest_bid:    i64,       // the current highest bid amount
    pub highest_bidder: ScAgentId, // the current highest bidder
    pub minimum_bid:    i64,       // minimum bid amount
    pub num_tokens:     i64,       // number of tokens for sale
    pub owner_margin:   i64,       // auction owner's margin in promilles
    pub when_started:   i64,       // timestamp when auction started
    //@formatter:on
}

pub struct BidInfo {
    //@formatter:off
    pub amount:    i64, // cumulative amount of bids from same bidder
    pub index:     i64, // index of bidder in bidder list
    pub timestamp: i64, // timestamp of most recent bid
    //@formatter:on
}

pub fn encode_auction_info(o: &AuctionInfo) -> Vec<u8> {
    let mut encode = BytesEncoder::new();
    encode.color(&o.color);
    encode.agent(&o.creator);
    encode.int(o.deposit);
    encode.string(&o.description);
    encode.int(o.duration);
    encode.int(o.highest_bid);
    encode.agent(&o.highest_bidder);
    encode.int(o.minimum_bid);
    encode.int(o.num_tokens);
    encode.int(o.owner_margin);
    encode.int(o.when_started);
    return encode.data();
}

pub fn decode_auction_info(bytes: &[u8]) -> AuctionInfo {
    let mut decode = BytesDecoder::new(bytes);
    AuctionInfo {
        color: decode.color(),
        creator: decode.agent(),
        deposit: decode.int(),
        description: decode.string(),
        duration: decode.int(),
        highest_bid: decode.int(),
        highest_bidder: decode.agent(),
        minimum_bid: decode.int(),
        num_tokens: decode.int(),
        owner_margin: decode.int(),
        when_started: decode.int(),
    }
}

pub fn encode_bid_info(o: &BidInfo) -> Vec<u8> {
    let mut encode = BytesEncoder::new();
    encode.int(o.amount);
    encode.int(o.index);
    encode.int(o.timestamp);
    return encode.data();
}

pub fn decode_bid_info(bytes: &[u8]) -> BidInfo {
    let mut decode = BytesDecoder::new(bytes);
    BidInfo {
        amount: decode.int(),
        index: decode.int(),
        timestamp: decode.int(),
    }
}
