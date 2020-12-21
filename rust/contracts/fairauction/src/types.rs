// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use wasplib::client::*;

pub struct AuctionInfo {
    pub auction_owner: ScAgent,
    pub color: ScColor,
    pub deposit: i64,
    pub description: String,
    pub duration: i64,
    pub highest_bid: i64,
    pub highest_bidder: ScAgent,
    pub minimum_bid: i64,
    pub num_tokens: i64,
    pub owner_margin: i64,
    pub when_started: i64,
}

pub struct BidInfo {
    pub amount: i64,
    pub index: i64,
    pub timestamp: i64,
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
