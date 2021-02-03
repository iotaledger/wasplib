// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use wasplib::client::*;

pub struct BetInfo {
    //@formatter:off
    pub amount: i64,
    pub better: ScAgentId,
    pub color:  i64,
    //@formatter:on
}

pub fn encode_bet_info(o: &BetInfo) -> Vec<u8> {
    let mut encode = BytesEncoder::new();
    encode.int(o.amount);
    encode.agent(&o.better);
    encode.int(o.color);
    return encode.data();
}

pub fn decode_bet_info(bytes: &[u8]) -> BetInfo {
    let mut decode = BytesDecoder::new(bytes);
    BetInfo {
        amount: decode.int(),
        better: decode.agent(),
        color: decode.int(),
    }
}
