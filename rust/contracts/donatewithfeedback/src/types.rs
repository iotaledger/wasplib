// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use wasplib::client::*;

pub struct DonationInfo {
    pub amount: i64,
    pub donator: ScAgent,
    pub error: String,
    pub feedback: String,
    pub timestamp: i64,
}

pub fn encode_donation_info(o: &DonationInfo) -> Vec<u8> {
    let mut e = BytesEncoder::new();
    e.int(o.amount);
    e.agent(&o.donator);
    e.string(&o.error);
    e.string(&o.feedback);
    e.int(o.timestamp);
    return e.data();
}

pub fn decode_donation_info(bytes: &[u8]) -> DonationInfo {
    let mut d = BytesDecoder::new(bytes);
    DonationInfo {
        amount: d.int(),
        donator: d.agent(),
        error: d.string(),
        feedback: d.string(),
        timestamp: d.int(),
    }
}
