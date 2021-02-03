// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use std::convert::TryInto;

use crate::client::host::get_key_id_from_bytes;

use super::context::*;
use super::keys::*;

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScAddress {
    id: [u8; 33],
}

impl ScAddress {
    pub fn from_bytes(bytes: &[u8]) -> ScAddress {
        ScAddress { id: bytes.try_into().expect("invalid address id length") }
    }

    pub fn as_agent_id(&self) -> ScAgentId {
        let mut a = ScAgentId { id: [0; 37] };
        a.id[..33].copy_from_slice(&self.id[..33]);
        a
    }

    pub fn equals(&self, other: &ScAddress) -> bool {
        self.id == other.id
    }

    pub fn to_bytes(&self) -> &[u8] {
        &self.id
    }

    pub fn to_string(&self) -> String {
        base58_encode(&self.id)
    }
}

impl MapKey for ScAddress {
    fn get_id(&self) -> Key32 {
        get_key_id_from_bytes(self.to_bytes())
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

#[derive(PartialEq, Clone)]
pub struct ScAgentId {
    id: [u8; 37],
}

impl ScAgentId {
    pub fn from_bytes(bytes: &[u8]) -> ScAgentId {
        ScAgentId { id: bytes.try_into().expect("invalid agent id lengths") }
    }

    pub fn address(&self) -> ScAddress {
        let mut a = ScAddress { id: [0; 33] };
        a.id[..33].copy_from_slice(&self.id[..33]);
        a
    }

    pub fn equals(&self, other: &ScAgentId) -> bool {
        self.id == other.id
    }

    pub fn is_address(&self) -> bool {
        self.address().as_agent_id().equals(self)
    }

    pub fn to_bytes(&self) -> &[u8] {
        &self.id
    }

    pub fn to_string(&self) -> String {
        base58_encode(&self.id)
    }
}

impl MapKey for ScAgentId {
    fn get_id(&self) -> Key32 {
        get_key_id_from_bytes(self.to_bytes())
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

#[derive(PartialEq, Clone)]
pub struct ScChainId {
    id: [u8; 33],
}

impl ScChainId {
    pub fn from_bytes(bytes: &[u8]) -> ScChainId {
        ScChainId { id: bytes.try_into().expect("invalid chain id length") }
    }

    pub fn equals(&self, other: &ScChainId) -> bool {
        self.id == other.id
    }

    pub fn to_bytes(&self) -> &[u8] {
        &self.id
    }

    pub fn to_string(&self) -> String {
        base58_encode(&self.id)
    }
}

impl MapKey for ScChainId {
    fn get_id(&self) -> Key32 {
        get_key_id_from_bytes(self.to_bytes())
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

#[derive(PartialEq, Clone)]
pub struct ScContractId {
    id: [u8; 37],
}

impl ScContractId {
    pub fn new(chain_id: &ScChainId, hname: &ScHname) -> ScContractId {
        let mut c = ScContractId { id: [0; 37] };
        c.id[..33].copy_from_slice(&chain_id.to_bytes());
        c.id[33..].copy_from_slice(&hname.to_bytes());
        c
    }

    pub fn from_bytes(bytes: &[u8]) -> ScContractId {
        ScContractId { id: bytes.try_into().expect("invalid contract id length") }
    }

    pub fn as_agent_id(&self) -> ScAgentId {
        let mut a = ScAgentId { id: [0x00; 37] };
        a.id[..].copy_from_slice(&self.id[..]);
        a
    }

    pub fn chain_id(&self) -> ScChainId {
        let mut c = ScChainId { id: [0; 33] };
        c.id[..33].copy_from_slice(&self.id[..33]);
        c
    }

    pub fn equals(&self, other: &ScContractId) -> bool {
        self.id == other.id
    }

    pub fn hname(&self) -> ScHname {
        ScHname::from_bytes(&self.id[33..])
    }

    pub fn to_bytes(&self) -> &[u8] {
        &self.id
    }

    pub fn to_string(&self) -> String {
        base58_encode(&self.id)
    }
}

impl MapKey for ScContractId {
    fn get_id(&self) -> Key32 {
        get_key_id_from_bytes(self.to_bytes())
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

#[derive(PartialEq, Clone)]
pub struct ScColor {
    id: [u8; 32],
}

impl ScColor {
    pub const IOTA: ScColor = ScColor { id: [0x00; 32] };
    pub const MINT: ScColor = ScColor { id: [0xff; 32] };

    pub fn from_bytes(bytes: &[u8]) -> ScColor {
        ScColor { id: bytes.try_into().expect("invalid color id length") }
    }

    pub fn equals(&self, other: &ScColor) -> bool {
        self.id == other.id
    }

    pub fn to_bytes(&self) -> &[u8] {
        &self.id
    }

    pub fn to_string(&self) -> String {
        base58_encode(&self.id)
    }
}

impl MapKey for ScColor {
    fn get_id(&self) -> Key32 {
        get_key_id_from_bytes(self.to_bytes())
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

#[derive(PartialEq, Clone)]
pub struct ScHash {
    id: [u8; 32],
}

impl ScHash {
    pub fn from_bytes(bytes: &[u8]) -> ScHash {
        ScHash { id: bytes.try_into().expect("invalid hash id length") }
    }

    pub fn equals(&self, other: &ScHash) -> bool {
        self.id == other.id
    }

    pub fn to_bytes(&self) -> &[u8] {
        &self.id
    }

    pub fn to_string(&self) -> String {
        base58_encode(&self.id)
    }
}

impl MapKey for ScHash {
    fn get_id(&self) -> Key32 {
        get_key_id_from_bytes(self.to_bytes())
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

#[derive(Clone, Copy)]
pub struct ScHname(pub u32);

impl ScHname {
    pub fn new(name: &str) -> ScHname {
        ScCallContext {}.utility().hname(name)
    }

    pub fn from_bytes(bytes: &[u8]) -> ScHname {
        if bytes.len() != 4 { panic!("Hname should be 4 bytes"); }
        let val = bytes[3] as u32;
        let val = (val << 8) | (bytes[2] as u32);
        let val = (val << 8) | (bytes[1] as u32);
        let val = (val << 8) | (bytes[0] as u32);
        ScHname(val)
    }

    pub fn equals(&self, other: ScHname) -> bool {
        self.0 == other.0
    }

    pub fn to_bytes(&self) -> Vec<u8> {
        let val = self.0;
        let mut bytes: Vec<u8> = Vec::new();
        bytes.push((val >> 0) as u8);
        bytes.push((val >> 8) as u8);
        bytes.push((val >> 16) as u8);
        bytes.push((val >> 24) as u8);
        bytes
    }

    pub fn to_string(&self) -> String {
        self.0.to_string()
    }
}

impl MapKey for ScHname {
    fn get_id(&self) -> Key32 {
        get_key_id_from_bytes(&self.0.to_ne_bytes())
    }
}
