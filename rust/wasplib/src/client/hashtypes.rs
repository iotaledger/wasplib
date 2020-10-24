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
pub struct ScRequestId {
    id: String,
}

impl ScRequestId {
    pub fn from_bytes(bytes: &str) -> ScRequestId {
        ScRequestId { id: bytes.to_string() }
        //ScRequestId { id: bytes.try_into().expect("request id should be 34 bytes") }
    }

    pub fn to_bytes(&self) -> &str {
        &self.id
    }

    pub fn to_string(&self) -> String {
        self.id.to_string()
    }

    // pub fn transaction_id(&self) -> ScTransactionId {
    //     ScTransactionId::from_bytes(&self.id[..32])
    // }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

#[derive(Eq, PartialEq)]
pub struct ScTransactionId {
    id: String,
}

impl ScTransactionId {
    pub fn from_bytes(bytes: &str) -> ScTransactionId {
        ScTransactionId { id: bytes.to_string() }
        //ScTransactionId { id: bytes.try_into().expect("transaction id should be 32 bytes") }
    }

    pub fn to_bytes(&self) -> &str {
        &self.id
    }

    pub fn to_string(&self) -> String {
        self.id.to_string()
    }
}

fn base58_encode(bytes: &[u8]) -> String {
    ScContext::new().utility().base58_encode(bytes)
}