// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use std::convert::TryInto;

use crate::context::*;
use crate::host::*;
use crate::keys::*;

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// container object for 33-byte Tangle address ids
#[derive(PartialEq, Clone)]
pub struct ScAddress {
    id: [u8; 33],
}

impl ScAddress {
    // construct from byte array
    pub fn from_bytes(bytes: &[u8]) -> ScAddress {
        ScAddress { id: bytes.try_into().expect("invalid address id length") }
    }

    // returns agent id representation of this Tangle address
    pub fn as_agent_id(&self) -> ScAgentId {
        let mut a = ScAgentId { id: [0; 37] };
        a.id[..33].copy_from_slice(&self.id[..33]);
        a
    }

    // equality check
    pub fn equals(&self, other: &ScAddress) -> bool {
        self.id == other.id
    }

    // convert to byte array representation
    pub fn to_bytes(&self) -> &[u8] {
        &self.id
    }

    // human-readable string representation
    pub fn to_string(&self) -> String {
        base58_encode(&self.id)
    }
}

// allow to be used as key in maps
impl MapKey for ScAddress {
    fn get_id(&self) -> Key32 {
        get_key_id_from_bytes(self.to_bytes())
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// container object for 37-byte agent ids
#[derive(PartialEq, Clone)]
pub struct ScAgentId {
    id: [u8; 37],
}

impl ScAgentId {
    // construct from byte array
    pub fn from_bytes(bytes: &[u8]) -> ScAgentId {
        ScAgentId { id: bytes.try_into().expect("invalid agent id lengths") }
    }

    // gets Tangle address from agent id
    pub fn address(&self) -> ScAddress {
        let mut a = ScAddress { id: [0; 33] };
        a.id[..33].copy_from_slice(&self.id[..33]);
        a
    }

    // equality check
    pub fn equals(&self, other: &ScAgentId) -> bool {
        self.id == other.id
    }

    // checks to see if agent id represents a Tangle address
    pub fn is_address(&self) -> bool {
        self.address().as_agent_id().equals(self)
    }

    // convert to byte array representation
    pub fn to_bytes(&self) -> &[u8] {
        &self.id
    }

    // human-readable string representation
    pub fn to_string(&self) -> String {
        base58_encode(&self.id)
    }
}

// allow to be used as key in maps
impl MapKey for ScAgentId {
    fn get_id(&self) -> Key32 {
        get_key_id_from_bytes(self.to_bytes())
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// container object for 33-byte chain ids
#[derive(PartialEq, Clone)]
pub struct ScChainId {
    id: [u8; 33],
}

impl ScChainId {
    // construct from byte array
    pub fn from_bytes(bytes: &[u8]) -> ScChainId {
        ScChainId { id: bytes.try_into().expect("invalid chain id length") }
    }

    // equality check
    pub fn equals(&self, other: &ScChainId) -> bool {
        self.id == other.id
    }

    // convert to byte array representation
    pub fn to_bytes(&self) -> &[u8] {
        &self.id
    }

    // human-readable string representation
    pub fn to_string(&self) -> String {
        base58_encode(&self.id)
    }
}

// allow to be used as key in maps
impl MapKey for ScChainId {
    fn get_id(&self) -> Key32 {
        get_key_id_from_bytes(self.to_bytes())
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// container object for 37-byte contract ids
#[derive(PartialEq, Clone)]
pub struct ScContractId {
    id: [u8; 37],
}

impl ScContractId {
    // construct from chain id and contract name hash
    pub fn new(chain_id: &ScChainId, hname: &ScHname) -> ScContractId {
        let mut c = ScContractId { id: [0; 37] };
        c.id[..33].copy_from_slice(&chain_id.to_bytes());
        c.id[33..].copy_from_slice(&hname.to_bytes());
        c
    }

    // construct from byte array
    pub fn from_bytes(bytes: &[u8]) -> ScContractId {
        ScContractId { id: bytes.try_into().expect("invalid contract id length") }
    }

    // get agent id representation of contract id
    pub fn as_agent_id(&self) -> ScAgentId {
        let mut a = ScAgentId { id: [0x00; 37] };
        a.id[..].copy_from_slice(&self.id[..]);
        a
    }

    // get chain id of chain that contract is on
    pub fn chain_id(&self) -> ScChainId {
        let mut c = ScChainId { id: [0; 33] };
        c.id[..33].copy_from_slice(&self.id[..33]);
        c
    }

    // equality check
    pub fn equals(&self, other: &ScContractId) -> bool {
        self.id == other.id
    }

    // get contract name hash for this contract
    pub fn hname(&self) -> ScHname {
        ScHname::from_bytes(&self.id[33..])
    }

    // convert to byte array representation
    pub fn to_bytes(&self) -> &[u8] {
        &self.id
    }

    // human-readable string representation
    pub fn to_string(&self) -> String {
        base58_encode(&self.id)
    }
}

// allow to be used as key in maps
impl MapKey for ScContractId {
    fn get_id(&self) -> Key32 {
        get_key_id_from_bytes(self.to_bytes())
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// container object for 32-byte token color
#[derive(PartialEq, Clone)]
pub struct ScColor {
    id: [u8; 32],
}

impl ScColor {
    // predefined colors
    pub const IOTA: ScColor = ScColor { id: [0x00; 32] };
    pub const MINT: ScColor = ScColor { id: [0xff; 32] };

    // construct from byte array
    pub fn from_bytes(bytes: &[u8]) -> ScColor {
        ScColor { id: bytes.try_into().expect("invalid color id length") }
    }

    // equality check
    pub fn equals(&self, other: &ScColor) -> bool {
        self.id == other.id
    }

    // convert to byte array representation
    pub fn to_bytes(&self) -> &[u8] {
        &self.id
    }

    // human-readable string representation
    pub fn to_string(&self) -> String {
        base58_encode(&self.id)
    }
}

// allow to be used as key in maps
impl MapKey for ScColor {
    fn get_id(&self) -> Key32 {
        get_key_id_from_bytes(self.to_bytes())
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// container object for 32-byte hash value
#[derive(PartialEq, Clone)]
pub struct ScHash {
    id: [u8; 32],
}

impl ScHash {
    // construct from byte array
    pub fn from_bytes(bytes: &[u8]) -> ScHash {
        ScHash { id: bytes.try_into().expect("invalid hash id length") }
    }

    // equality check
    pub fn equals(&self, other: &ScHash) -> bool {
        self.id == other.id
    }

    // convert to byte array representation
    pub fn to_bytes(&self) -> &[u8] {
        &self.id
    }

    // human-readable string representation
    pub fn to_string(&self) -> String {
        base58_encode(&self.id)
    }
}

// allow to be used as key in maps
impl MapKey for ScHash {
    fn get_id(&self) -> Key32 {
        get_key_id_from_bytes(self.to_bytes())
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// container object for 4-byte name hash
#[derive(Clone, Copy)]
pub struct ScHname(pub u32);

impl ScHname {
    // construct from name string
    pub fn new(name: &str) -> ScHname {
        ScFuncContext {}.utility().hname(name)
    }

    // construct from byte array
    pub fn from_bytes(bytes: &[u8]) -> ScHname {
        if bytes.len() != 4 { panic!("Hname should be 4 bytes"); }
        let val = bytes[3] as u32;
        let val = (val << 8) | (bytes[2] as u32);
        let val = (val << 8) | (bytes[1] as u32);
        let val = (val << 8) | (bytes[0] as u32);
        ScHname(val)
    }

    // equality check
    pub fn equals(&self, other: ScHname) -> bool {
        self.0 == other.0
    }

    // convert to byte array representation
    pub fn to_bytes(&self) -> Vec<u8> {
        let val = self.0;
        let mut bytes: Vec<u8> = Vec::new();
        bytes.push((val >> 0) as u8);
        bytes.push((val >> 8) as u8);
        bytes.push((val >> 16) as u8);
        bytes.push((val >> 24) as u8);
        bytes
    }

    // human-readable string representation
    pub fn to_string(&self) -> String {
        self.0.to_string()
    }
}

// allow to be used as key in maps
impl MapKey for ScHname {
    fn get_id(&self) -> Key32 {
        get_key_id_from_bytes(&self.0.to_ne_bytes())
    }
}
