// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use crate::hashtypes::*;

// decodes separate entities from a byte buffer
pub struct BytesDecoder<'a> {
    data: &'a [u8],
}

impl BytesDecoder<'_> {
    // constructs a decoder
    pub fn new(data: &[u8]) -> BytesDecoder {
        BytesDecoder { data: data }
    }

    // decodes an ScAddress from the byte buffer
    pub fn address(&mut self) -> ScAddress {
        ScAddress::from_bytes(self.bytes())
    }

    // decodes an ScAgentId from the byte buffer
    pub fn agent_id(&mut self) -> ScAgentId {
        ScAgentId::from_bytes(self.bytes())
    }

    // decodes the next substring of bytes from the byte buffer
    pub fn bytes(&mut self) -> &[u8] {
        let size = self.int64() as usize;
        if self.data.len() < size {
            panic!("Cannot decode bytes");
        }
        let value = &self.data[..size];
        self.data = &self.data[size..];
        value
    }

    // decodes an ScChainId from the byte buffer
    pub fn chain_id(&mut self) -> ScChainId {
        ScChainId::from_bytes(self.bytes())
    }

    // decodes an ScColor from the byte buffer
    pub fn color(&mut self) -> ScColor {
        ScColor::from_bytes(self.bytes())
    }

    // decodes an ScContractId from the byte buffer
    pub fn contract_id(&mut self) -> ScContractId {
        ScContractId::from_bytes(self.bytes())
    }

    // decodes an ScHash from the byte buffer
    pub fn hash(&mut self) -> ScHash {
        ScHash::from_bytes(self.bytes())
    }

    // decodes an ScHname from the byte buffer
    pub fn hname(&mut self) -> ScHname {
        ScHname::from_bytes(self.bytes())
    }

    // decodes an int64 from the byte buffer
    // note that ints are encoded using leb128 encoding
    pub fn int64(&mut self) -> i64 {
        // leb128 decoder
        let mut val = 0_i64;
        let mut s = 0;
        loop {
            let mut b = self.data[0] as i8;
            self.data = &self.data[1..];
            val |= ((b & 0x7f) as i64) << s;
            if b >= 0 {
                if ((val >> s) as i8) & 0x7f != b & 0x7f {
                    panic!("Integer too large");
                }
                // extend int7 sign to int8
                if (b & 0x40) != 0 {
                    b |= -0x80
                }
                // extend int8 sign to int64
                return val | ((b as i64) << s);
            }
            s += 7;
            if s >= 64 {
                panic!("integer representation too long");
            }
        }
    }

    // decodes an ScRequestId from the byte buffer
    pub fn request_id(&mut self) -> ScRequestId {
        ScRequestId::from_bytes(self.bytes())
    }

    // decodes an UTF-8 text string from the byte buffer
    pub fn string(&mut self) -> String {
        String::from_utf8_lossy(self.bytes()).to_string()
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// encodes separate entities into a byte buffer
pub struct BytesEncoder {
    data: Vec<u8>,
}

impl BytesEncoder {
    // constructs an encoder
    pub fn new() -> BytesEncoder {
        BytesEncoder { data: Vec::new() }
    }

    // encodes an ScAddress into the byte buffer
    pub fn address(&mut self, value: &ScAddress) -> &BytesEncoder {
        self.bytes(value.to_bytes());
        self
    }

    // encodes an ScAgentId into the byte buffer
    pub fn agent_id(&mut self, value: &ScAgentId) -> &BytesEncoder {
        self.bytes(value.to_bytes());
        self
    }

    // encodes a substring of bytes into the byte buffer
    pub fn bytes(&mut self, value: &[u8]) -> &BytesEncoder {
        self.int64(value.len() as i64);
        self.data.extend_from_slice(value);
        self
    }

    // encodes an ScChainId into the byte buffer
    pub fn chain_id(&mut self, value: &ScChainId) -> &BytesEncoder {
        self.bytes(value.to_bytes());
        self
    }

    // encodes an ScColor into the byte buffer
    pub fn color(&mut self, value: &ScColor) -> &BytesEncoder {
        self.bytes(value.to_bytes());
        self
    }

    // encodes an ScContractId into the byte buffer
    pub fn contract_id(&mut self, value: &ScContractId) -> &BytesEncoder {
        self.bytes(value.to_bytes());
        self
    }

    // retrieve the encoded byte buffer
    pub fn data(&self) -> Vec<u8> {
        self.data.clone()
    }

    // encodes an ScHash into the byte buffer
    pub fn hash(&mut self, value: &ScHash) -> &BytesEncoder {
        self.bytes(value.to_bytes());
        self
    }

    // encodes an ScHname into the byte buffer
    pub fn hname(&mut self, value: &ScHname) -> &BytesEncoder {
        self.bytes(&value.to_bytes());
        self
    }

    // encodes an int64 into the byte buffer
    // note that ints are encoded using leb128 encoding
    pub fn int64(&mut self, mut val: i64) -> &BytesEncoder {
        // leb128 encoder
        loop {
            let b = val as u8;
            let s = b & 0x40;
            val >>= 7;
            if (val == 0 && s == 0) || (val == -1 && s != 0) {
                self.data.push(b & 0x7f);
                return self;
            }
            self.data.push(b | 0x80)
        }
    }

    // encodes an ScRequestId into the byte buffer
    pub fn request_id(&mut self, value: &ScRequestId) -> &BytesEncoder {
        self.bytes(value.to_bytes());
        self
    }

    // encodes an UTF-8 text string into the byte buffer
    pub fn string(&mut self, value: &str) -> &BytesEncoder {
        self.bytes(value.as_bytes());
        self
    }
}
