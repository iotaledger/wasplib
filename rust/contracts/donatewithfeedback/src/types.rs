// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use wasplib::client::*;

pub struct DonationInfo {
    //@formatter:off
    pub amount:    i64,
    pub donator:   ScAgentId,
    pub error:     String,
    pub feedback:  String,
    pub timestamp: i64,
    //@formatter:on
}

pub fn encode_donation_info(o: &DonationInfo) -> Vec<u8> {
    let mut encode = BytesEncoder::new();
    encode.int(o.amount);
    encode.agent(&o.donator);
    encode.string(&o.error);
    encode.string(&o.feedback);
    encode.int(o.timestamp);
    return encode.data();
}

pub fn decode_donation_info(bytes: &[u8]) -> DonationInfo {
    let mut decode = BytesDecoder::new(bytes);
    DonationInfo {
        amount: decode.int(),
        donator: decode.agent(),
        error: decode.string(),
        feedback: decode.string(),
        timestamp: decode.int(),
    }
}
