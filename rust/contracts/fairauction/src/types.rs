// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use wasplib::client::*;

pub struct AuctionInfo {
    pub auction_owner: ScAgent, // issuer of start_auction transaction
    pub color: ScColor, // color of tokens for sale
    pub deposit: i64, // deposit by auction owner to cover the SC fees
    pub description: String, // auction description
    pub duration: i64, // auction duration in minutes
    pub highest_bid: i64, // the current highest bid amount
    pub highest_bidder: ScAgent, // the current highest bidder
    pub minimum_bid: i64, // minimum bid amount
    pub num_tokens: i64, // number of tokens for sale
    pub owner_margin: i64, // auction owner's margin in promilles
    pub when_started: i64, // timestamp when auction started
}

pub struct BidInfo {
    pub amount: i64, // cumulative amount of bids from same bidder
    pub index: i64, // index of bidder in bidder list
    pub timestamp: i64, // timestamp of most recent bid
}

pub fn encode_auction_info(o: &AuctionInfo) -> Vec<u8> {
    let mut e = BytesEncoder::new();
    e.agent(&o.auction_owner);
    e.color(&o.color);
    e.int(o.deposit);
    e.string(&o.description);
    e.int(o.duration);
    e.int(o.highest_bid);
    e.agent(&o.highest_bidder);
    e.int(o.minimum_bid);
    e.int(o.num_tokens);
    e.int(o.owner_margin);
    e.int(o.when_started);
    return e.data();
}

pub fn decode_auction_info(bytes: &[u8]) -> AuctionInfo {
    let mut d = BytesDecoder::new(bytes);
    AuctionInfo {
        auction_owner: d.agent(),
        color: d.color(),
        deposit: d.int(),
        description: d.string(),
        duration: d.int(),
        highest_bid: d.int(),
        highest_bidder: d.agent(),
        minimum_bid: d.int(),
        num_tokens: d.int(),
        owner_margin: d.int(),
        when_started: d.int(),
    }
}

pub fn encode_bid_info(o: &BidInfo) -> Vec<u8> {
    let mut e = BytesEncoder::new();
    e.int(o.amount);
    e.int(o.index);
    e.int(o.timestamp);
    return e.data();
}

pub fn decode_bid_info(bytes: &[u8]) -> BidInfo {
    let mut d = BytesDecoder::new(bytes);
    BidInfo {
        amount: d.int(),
        index: d.int(),
        timestamp: d.int(),
    }
}
