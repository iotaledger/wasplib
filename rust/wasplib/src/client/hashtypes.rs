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