// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use std::convert::TryInto;

use crate::client::host::get_key_id_from_bytes;

use super::context::*;
use super::keys::*;

pub struct ScAddress {
    address: [u8; 33],
}

impl ScAddress {
    pub const NULL: ScAddress = ScAddress { address: [0x00; 33] };

    pub fn as_agent(&self) -> ScAgent {
        let mut agent = ScAgent { agent: [0; 37] };
        agent.agent[..33].copy_from_slice(&self.address[..33]);
        agent
    }

    pub fn equals(&self, other: &ScAddress) -> bool {
        self.address == other.address
    }

    pub fn from_bytes(bytes: &[u8]) -> ScAddress {
        ScAddress { address: bytes.try_into().expect("address id should be 33 bytes") }
    }

    pub fn to_bytes(&self) -> &[u8] {
        &self.address
    }

    pub fn to_string(&self) -> String {
        base58_encode(&self.address)
    }
}

impl MapKey for ScAddress {
    fn get_id(&self) -> i32 {
        get_key_id_from_bytes(self.to_bytes())
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScAgent {
    agent: [u8; 37],
}

impl ScAgent {
    pub const NULL: ScAgent = ScAgent { agent: [0x00; 37] };

    pub fn address(&self) -> ScAddress {
        let mut address = ScAddress { address: [0; 33] };
        address.address[..33].copy_from_slice(&self.agent[..33]);
        address
    }

    pub fn equals(&self, other: &ScAgent) -> bool {
        self.agent == other.agent
    }

    pub fn from_bytes(bytes: &[u8]) -> ScAgent {
        ScAgent { agent: bytes.try_into().expect("agent id should be 37 bytes") }
    }

    pub fn is_address(&self) -> bool {
        self.address().as_agent().equals(self)
    }

    pub fn to_bytes(&self) -> &[u8] {
        &self.agent
    }

    pub fn to_string(&self) -> String {
        base58_encode(&self.agent)
    }
}

impl MapKey for ScAgent {
    fn get_id(&self) -> i32 {
        get_key_id_from_bytes(self.to_bytes())
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScColor {
    color: [u8; 32],
}

impl ScColor {
    pub const IOTA: ScColor = ScColor { color: [0x00; 32] };
    pub const MINT: ScColor = ScColor { color: [0xff; 32] };

    pub fn equals(&self, other: &ScColor) -> bool {
        self.color == other.color
    }

    pub fn from_bytes(bytes: &[u8]) -> ScColor {
        ScColor { color: bytes.try_into().expect("color id should be 32 bytes") }
    }

    pub fn to_bytes(&self) -> &[u8] {
        &self.color
    }

    pub fn to_string(&self) -> String {
        base58_encode(&self.color)
    }
}

impl MapKey for ScColor {
    fn get_id(&self) -> i32 {
        get_key_id_from_bytes(self.to_bytes())
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScHash {
    hash: [u8; 32],
}

impl ScHash {
    pub const NULL: ScHash = ScHash { hash: [0x00; 32] };

    pub fn equals(&self, other: &ScHash) -> bool {
        self.hash == other.hash
    }

    pub fn from_bytes(bytes: &[u8]) -> ScHash {
        ScHash { hash: bytes.try_into().expect("hash should be 32 bytes") }
    }

    pub fn to_bytes(&self) -> &[u8] {
        &self.hash
    }

    pub fn to_string(&self) -> String {
        base58_encode(&self.hash)
    }
}

impl MapKey for ScHash {
    fn get_id(&self) -> i32 {
        get_key_id_from_bytes(self.to_bytes())
    }
}
