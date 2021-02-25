// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use std::convert::TryInto;

use crate::keys::*;

// all TYPE_* values should exactly match the counterpart OBJTYPE_* values on the host!
pub const TYPE_ARRAY: i32 = 0x20;

pub const TYPE_ADDRESS: i32 = 1;
pub const TYPE_AGENT_ID: i32 = 2;
pub const TYPE_BYTES: i32 = 3;
pub const TYPE_CHAIN_ID: i32 = 4;
pub const TYPE_COLOR: i32 = 5;
pub const TYPE_CONTRACT_ID: i32 = 6;
pub const TYPE_HASH: i32 = 7;
pub const TYPE_HNAME: i32 = 8;
pub const TYPE_INT64: i32 = 9;
pub const TYPE_MAP: i32 = 10;
pub const TYPE_REQUEST_ID: i32 = 11;
pub const TYPE_STRING: i32 = 12;

const TYPE_SIZES: &[usize] = &[0, 33, 37, 0, 33, 32, 37, 32, 4, 8, 0, 34, 0];

// any host function that gets called once the current request has
// entered an error state will immediately return without action.
// Any return value will be zero or empty string in that case
#[link(wasm_import_module = "wasplib")]
extern {
    pub fn hostGetBytes(obj_id: i32, key_id: i32, type_id: i32, value: *mut u8, len: i32) -> i32;
    pub fn hostGetKeyId(key: *const u8, len: i32) -> i32;
    pub fn hostGetObjectId(obj_id: i32, key_id: i32, type_id: i32) -> i32;
    pub fn hostSetBytes(obj_id: i32, key_id: i32, type_id: i32, value: *const u8, len: i32);
}

pub fn clear(obj_id: i32) {
    set_bytes(obj_id, KEY_LENGTH, TYPE_INT64, &0_i64.to_le_bytes())
}

pub fn exists(obj_id: i32, key_id: Key32, type_id: i32) -> bool {
    unsafe {
        // negative length (-1) means only test for existence
        // returned size -1 indicates keyId not found (or error)
        // this removes the need for a separate hostExists function
        hostGetBytes(obj_id, key_id.0, type_id, std::ptr::null_mut(), -1) >= 0
    }
}

pub fn get_bytes(obj_id: i32, key_id: Key32, type_id: i32) -> Vec<u8> {
    unsafe {
        // first query length of bytes array
        let size = hostGetBytes(obj_id, key_id.0, type_id, std::ptr::null_mut(), 0);
        if size <= 0 { return vec![0_u8; TYPE_SIZES[type_id as usize]]; }

        // allocate a byte array in Wasm memory and
        // copy the actual data bytes to Wasm byte array
        let mut bytes = vec![0_u8; size as usize];
        hostGetBytes(obj_id, key_id.0, type_id, bytes.as_mut_ptr(), size);
        return bytes;
    }
}

pub fn get_key_id_from_bytes(bytes: &[u8]) -> Key32 {
    unsafe {
        let size = bytes.len() as i32;
        // negative size indicates this was from bytes
        Key32(hostGetKeyId(bytes.as_ptr(), -size - 1))
    }
}

pub fn get_key_id_from_string(key: &str) -> Key32 {
    let bytes = key.as_bytes();
    unsafe {
        // non-negative size indicates this was from string
        Key32(hostGetKeyId(bytes.as_ptr(), bytes.len() as i32))
    }
}

pub fn get_length(obj_id: i32) -> i32 {
    let bytes = get_bytes(obj_id, KEY_LENGTH, TYPE_INT64);
    i64::from_le_bytes(bytes.try_into().unwrap()) as i32
}

pub fn get_object_id(obj_id: i32, key_id: Key32, type_id: i32) -> i32 {
    unsafe {
        hostGetObjectId(obj_id, key_id.0, type_id)
    }
}

pub fn set_bytes(obj_id: i32, key_id: Key32, type_id: i32, value: &[u8]) {
    unsafe {
        hostSetBytes(obj_id, key_id.0, type_id, value.as_ptr(), value.len() as i32)
    }
}
