// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use std::convert::TryInto;

use crate::client::ScContext;

#[derive(Eq, PartialEq)]
pub struct ScAddress {
    address: [u8; 33],
}

impl ScAddress {
    pub fn from_bytes(bytes: &[u8]) -> ScAddress {
        ScAddress { address: bytes.try_into().expect("address should be 33 bytes") }
    }

    pub fn to_bytes(&self) -> &[u8] {
        &self.address
    }

    pub fn to_string(&self) -> String {
        base58_encode(&self.address)
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

#[derive(Eq, PartialEq)]
pub struct ScAgent {
    id: [u8; 37],
}

impl ScAgent {
    pub fn from_bytes(bytes: &[u8]) -> ScAgent {
        ScAgent { id: bytes.try_into().expect("agent id should be 37 bytes") }
    }

    pub fn to_bytes(&self) -> &[u8] {
        &self.id
    }

    pub fn to_string(&self) -> String {
        base58_encode(&self.id)
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

#[derive(Eq, PartialEq)]
pub struct ScColor {
    color: [u8; 32],
}

impl ScColor {
    pub const IOTA: ScColor = ScColor { color: [0x00; 32] };
    pub const MINT: ScColor = ScColor { color: [0xff; 32] };

    pub fn from_bytes(bytes: &[u8]) -> ScColor {
        ScColor { color: bytes.try_into().expect("color should be 32 bytes") }
    }

    pub fn to_bytes(&self) -> &[u8] {
        &self.color
    }

    pub fn to_string(&self) -> String {
        base58_encode(&self.color)
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

#[derive(Eq, PartialEq)]
pub struct ScRequestId {
    id: [u8; 34],
}

impl ScRequestId {
    pub fn from_bytes(bytes: &[u8]) -> ScRequestId {
        ScRequestId { id: bytes.try_into().expect("request id should be 34 bytes") }
    }

    pub fn to_bytes(&self) -> &[u8] {
        &self.id
    }

    pub fn to_string(&self) -> String {
        base58_encode(&self.id)
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

#[derive(Eq, PartialEq)]
pub struct ScTxHash {
    hash: [u8; 32],
}

impl ScTxHash {
    pub fn from_bytes(bytes: &[u8]) -> ScTxHash {
        ScTxHash { hash: bytes.try_into().expect("tx hash should be 32 bytes") }
    }

    pub fn to_bytes(&self) -> &[u8] {
        &self.hash
    }

    pub fn to_string(&self) -> String {
        base58_encode(&self.hash)
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

fn base58_encode(bytes: &[u8]) -> String {
    ScContext::new().utility().base58_encode(bytes)
}
