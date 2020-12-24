// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use wasplib::client::*;

pub struct Member {
    //@formatter:off
    pub address: ScAddress,
    pub factor:  i64,
    //@formatter:on
}

pub fn encode_member(o: &Member) -> Vec<u8> {
    let mut encode = BytesEncoder::new();
    encode.address(&o.address);
    encode.int(o.factor);
    return encode.data();
}

pub fn decode_member(bytes: &[u8]) -> Member {
    let mut decode = BytesDecoder::new(bytes);
    Member {
        address: decode.address(),
        factor: decode.int(),
    }
}
