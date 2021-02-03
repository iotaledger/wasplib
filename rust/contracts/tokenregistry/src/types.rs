// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use wasplib::client::*;

pub struct TokenInfo {
    //@formatter:off
    pub created:      i64,       // creation timestamp
    pub description:  String,    // description what minted token represents
    pub minted_by:    ScAgentId, // original minter
    pub owner:        ScAgentId, // current owner
    pub supply:       i64,       // amount of tokens originally minted
    pub updated:      i64,       // last update timestamp
    pub user_defined: String,    // any user defined text
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
