// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use wasplib::client::*;

pub struct TokenInfo {
    pub supply: i64,
    pub minted_by: ScAgent,
    pub owner: ScAgent,
    pub created: i64,
    pub updated: i64,
    pub description: String,
    pub user_defined: String,
}

pub fn encode_token_info(o: &TokenInfo) -> Vec<u8> {
    let mut e = BytesEncoder::new();
    e.int(o.supply);
    e.agent(&o.minted_by);
    e.agent(&o.owner);
    e.int(o.created);
    e.int(o.updated);
    e.string(&o.description);
    e.string(&o.user_defined);
    return e.data();
}

pub fn decode_token_info(bytes: &[u8]) -> TokenInfo {
    let mut d = BytesDecoder::new(bytes);
    TokenInfo {
        supply: d.int(),
        minted_by: d.agent(),
        owner: d.agent(),
        created: d.int(),
        updated: d.int(),
        description: d.string(),
        user_defined: d.string(),
    }
}
