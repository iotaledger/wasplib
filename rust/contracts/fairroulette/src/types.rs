// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use wasplib::client::*;

pub struct BetInfo {
    pub amount: i64,
    pub better: ScAgent,
    pub color: i64,
}

pub fn encode_bet_info(o: &BetInfo) -> Vec<u8> {
    let mut e = BytesEncoder::new();
    e.int(o.amount);
    e.agent(&o.better);
    e.int(o.color);
    return e.data();
}

pub fn decode_bet_info(bytes: &[u8]) -> BetInfo {
    let mut d = BytesDecoder::new(bytes);
    BetInfo {
        amount: d.int(),
        better: d.agent(),
        color: d.int(),
    }
}
