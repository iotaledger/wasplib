// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use crate::hashtypes::*;
use crate::host::*;

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

    // decodes an ScAgentID from the byte buffer
    pub fn agent_id(&mut self) -> ScAgentID {
        ScAgentID::from_bytes(self.bytes())
    }

    // decodes the next substring of bytes from the byte buffer
    pub fn bytes(&mut self) -> &[u8] {
        let size = self.int32() as usize;
        if self.data.len() < size {
            panic("insufficient bytes");
        }
        let value = &self.data[..size];
        self.data = &self.data[size..];
        value
    }

    // decodes an ScChainID from the byte buffer
    pub fn chain_id(&mut self) -> ScChainID {
        ScChainID::from_bytes(self.bytes())
    }

    // decodes an ScColor from the byte buffer
    pub fn color(&mut self) -> ScColor {
        ScColor::from_bytes(self.bytes())
    }

    // decodes an ScHash from the byte buffer
    pub fn hash(&mut self) -> ScHash {
        ScHash::from_bytes(self.bytes())
    }

    // decodes an ScHname from the byte buffer
    pub fn hname(&mut self) -> ScHname {
        ScHname::from_bytes(self.bytes())
    }

    // decodes an int16 from the byte buffer
    // note that these are encoded using leb128 encoding to conserve space
    pub fn int16(&mut self) -> i16 {
        self.leb128_decode(16) as i16
    }

    // decodes an int32 from the byte buffer
    // note that these are encoded using leb128 encoding to conserve space
    pub fn int32(&mut self) -> i32 {
        self.leb128_decode(32) as i32
    }

    // decodes an int64 from the byte buffer
    // note that these are encoded using leb128 encoding to conserve space
    pub fn int64(&mut self) -> i64 {
        self.leb128_decode(64)
    }

    // leb128 decoder
    fn leb128_decode(&mut self, bits: i32) -> i64 {
        let mut val = 0_i64;
        let mut s = 0;
        loop {
            if self.data.len() == 0 {
                panic("insufficient bytes");
            }
            let mut b = self.data[0] as i8;
            self.data = &self.data[1..];
            val |= ((b & 0x7f) as i64) << s;

            // termination bit set?
            if (b & -0x80) == 0 {
                if ((val >> s) as i8) & 0x7f != b & 0x7f {
                    panic("integer too large");
                }

                // extend int7 sign to int8
                b |= (b & 0x40) << 1;

                // extend int8 sign to int64
                return val | ((b as i64) << s);
            }
            s += 7;
            if s >= bits {
                panic("integer representation too long");
            }
        }
    }

    // decodes an ScRequestID from the byte buffer
    pub fn request_id(&mut self) -> ScRequestID {
        ScRequestID::from_bytes(self.bytes())
    }

    // decodes an UTF-8 text string from the byte buffer
    pub fn string(&mut self) -> String {
        String::from_utf8_lossy(self.bytes()).to_string()
    }
}

impl Drop for BytesDecoder<'_> {
    fn drop(&mut self) {
        if self.data.len() != 0 {
            panic("extra bytes");
        }
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

    // encodes an ScAgentID into the byte buffer
    pub fn agent_id(&mut self, value: &ScAgentID) -> &BytesEncoder {
        self.bytes(value.to_bytes());
        self
    }

    // encodes a substring of bytes into the byte buffer
    pub fn bytes(&mut self, value: &[u8]) -> &BytesEncoder {
        self.int32(value.len() as i32);
        self.data.extend_from_slice(value);
        self
    }

    // encodes an ScChainID into the byte buffer
    pub fn chain_id(&mut self, value: &ScChainID) -> &BytesEncoder {
        self.bytes(value.to_bytes());
        self
    }

    // encodes an ScColor into the byte buffer
    pub fn color(&mut self, value: &ScColor) -> &BytesEncoder {
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

    // encodes an int16 into the byte buffer
    // note that these are encoded using leb128 encoding to conserve space
    pub fn int16(&mut self, val: i16) -> &BytesEncoder {
        self.leb128_encode(val as i64)
    }

    // encodes an int32 into the byte buffer
    // note that these are encoded using leb128 encoding to conserve space
    pub fn int32(&mut self, val: i32) -> &BytesEncoder {
        self.leb128_encode(val as i64)
    }

    // encodes an int64 into the byte buffer
    // note that these are encoded using leb128 encoding to conserve space
    pub fn int64(&mut self, val: i64) -> &BytesEncoder {
        self.leb128_encode(val)
    }

    // leb128 encoder
    fn leb128_encode(&mut self, mut val: i64) -> &BytesEncoder {
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

    // encodes an ScRequestID into the byte buffer
    pub fn request_id(&mut self, value: &ScRequestID) -> &BytesEncoder {
        self.bytes(value.to_bytes());
        self
    }

    // encodes an UTF-8 text string into the byte buffer
    pub fn string(&mut self, value: &str) -> &BytesEncoder {
        self.bytes(value.as_bytes());
        self
    }
}
