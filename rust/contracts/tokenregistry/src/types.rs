// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use wasplib::client::*;

pub struct TokenInfo {
    //@formatter:off
    pub created:      i64,
    pub description:  String,
    pub minted_by:    ScAgentId,
    pub owner:        ScAgentId,
    pub supply:       i64,
    pub updated:      i64,
    pub user_defined: String,
    //@formatter:on
}

pub fn encode_token_info(o: &TokenInfo) -> Vec<u8> {
    let mut encode = BytesEncoder::new();
    encode.int(o.created);
    encode.string(&o.description);
    encode.agent(&o.minted_by);
    encode.agent(&o.owner);
    encode.int(o.supply);
    encode.int(o.updated);
    encode.string(&o.user_defined);
    return encode.data();
}

pub fn decode_token_info(bytes: &[u8]) -> TokenInfo {
    let mut decode = BytesDecoder::new(bytes);
    TokenInfo {
        created: decode.int(),
        description: decode.string(),
        minted_by: decode.agent(),
        owner: decode.agent(),
        supply: decode.int(),
        updated: decode.int(),
        user_defined: decode.string(),
    }
}
