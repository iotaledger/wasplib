// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use std::convert::TryInto;

use super::context::ROOT_CALL_CONTEXT;

#[derive(Eq, PartialEq)]
pub struct ScAddress {
    address: [u8; 33],
}

impl ScAddress {
    pub fn as_agent(&self) -> ScAgent {
        let mut agent = ScAgent { id: [0; 37] };
        agent.id[0..33].copy_from_slice(&self.address[0..33]);
        agent
    }

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

fn base58_encode(bytes: &[u8]) -> String {
    ROOT_CALL_CONTEXT.utility().base58_encode(bytes)
}
