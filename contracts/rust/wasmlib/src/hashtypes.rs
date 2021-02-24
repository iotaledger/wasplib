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
        let mut agent_id = ScAgentId { id: [0; 37] };
        agent_id.id[..33].copy_from_slice(&self.id[..33]);
        agent_id
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
        let mut address = ScAddress { id: [0; 33] };
        address.id[..33].copy_from_slice(&self.id[..33]);
        address
    }

    // checks to see if agent id represents a Tangle address
    pub fn is_address(&self) -> bool {
        self.address().as_agent_id() == *self
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

    pub fn from_request_id(request_id: &ScRequestId) -> ScColor {
        let mut color = ScColor { id: [0x00; 32] };
        color.id[..].copy_from_slice(&request_id.id[..]);
        color
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

// container object for 37-byte contract ids
#[derive(PartialEq, Clone)]
pub struct ScContractId {
    id: [u8; 37],
}

impl ScContractId {
    // construct from chain id and contract name hash
    pub fn new(chain_id: &ScChainId, hname: &ScHname) -> ScContractId {
        let mut contract_id = ScContractId { id: [0; 37] };
        contract_id.id[..33].copy_from_slice(&chain_id.to_bytes());
        contract_id.id[33..].copy_from_slice(&hname.to_bytes());
        contract_id
    }

    // construct from byte array
    pub fn from_bytes(bytes: &[u8]) -> ScContractId {
        ScContractId { id: bytes.try_into().expect("invalid contract id length") }
    }

    // get agent id representation of contract id
    pub fn as_agent_id(&self) -> ScAgentId {
        let mut agent_id = ScAgentId { id: [0x00; 37] };
        agent_id.id[..].copy_from_slice(&self.id[..]);
        agent_id
    }

    // get chain id of chain that contract is on
    pub fn chain_id(&self) -> ScChainId {
        let mut chain_id = ScChainId { id: [0; 33] };
        chain_id.id[..33].copy_from_slice(&self.id[..33]);
        chain_id
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
#[derive(PartialEq, Clone)]
pub struct ScHname(pub u32);

impl ScHname {
    // construct from name string
    pub fn new(name: &str) -> ScHname {
        ScFuncContext {}.utility().hname(name)
    }

    // construct from byte array
    pub fn from_bytes(bytes: &[u8]) -> ScHname {
        let val = u32::from_le_bytes(bytes.try_into().expect("invalid hname length"));
        ScHname(val)
    }

    // convert to byte array representation
    pub fn to_bytes(&self) -> Vec<u8> {
        self.0.to_le_bytes().to_vec()
    }

    // human-readable string representation
    pub fn to_string(&self) -> String {
        self.0.to_string()
    }
}

// allow to be used as key in maps
impl MapKey for ScHname {
    fn get_id(&self) -> Key32 {
        get_key_id_from_bytes(&self.0.to_le_bytes())
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// container object for 34-byte transaction request ids
#[derive(PartialEq, Clone)]
pub struct ScRequestId {
    id: [u8; 34],
}

impl ScRequestId {
    // construct from byte array
    pub fn from_bytes(bytes: &[u8]) -> ScRequestId {
        ScRequestId { id: bytes.try_into().expect("invalid request id length") }
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
impl MapKey for ScRequestId {
    fn get_id(&self) -> Key32 {
        get_key_id_from_bytes(self.to_bytes())
    }
}

