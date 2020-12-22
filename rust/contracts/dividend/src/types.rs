// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use wasplib::client::*;

pub struct Member {
    pub address: ScAddress,
    pub factor: i64,
}

pub fn encode_member(o: &Member) -> Vec<u8> {
    let mut e = BytesEncoder::new();
    e.address(&o.address);
    e.int(o.factor);
    return e.data();
}

pub fn decode_member(bytes: &[u8]) -> Member {
    let mut d = BytesDecoder::new(bytes);
    Member {
        address: d.address(),
        factor: d.int(),
    }
}
