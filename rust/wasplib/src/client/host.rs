// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use crate::client::KEY_LENGTH;

// all TYPE_* values should exactly match the counterpart OBJTYPE_* values on the host!
pub const TYPE_ARRAY: i32 = 0x20;

pub const TYPE_ADDRESS: i32 = 1;
pub const TYPE_AGENT: i32 = 2;
pub const TYPE_BYTES: i32 = 3;
pub const TYPE_COLOR: i32 = 4;
pub const TYPE_HASH: i32 = 5;
pub const TYPE_INT: i32 = 6;
pub const TYPE_MAP: i32 = 7;
pub const TYPE_STRING: i32 = 8;

// any host function that gets called once the current request has
// entered an error state will immediately return without action.
// Any return value will be zero or empty string in that case
#[link(wasm_import_module = "wasplib")]
extern {
    pub fn hostGetBytes(obj_id: i32, key_id: i32, value: *mut u8, len: i32) -> i32;
    pub fn hostGetInt(obj_id: i32, key_id: i32) -> i64;
    pub fn hostGetKeyId(key: *const u8, len: i32) -> i32;
    pub fn hostGetObjectId(obj_id: i32, key_id: i32, type_id: i32) -> i32;
    pub fn hostSetBytes(obj_id: i32, key_id: i32, value: *const u8, len: i32);
    pub fn hostSetInt(obj_id: i32, key_id: i32, value: i64);
}

pub fn exists(obj_id: i32, key_id: i32) -> bool {
    unsafe {
        // negative length (-1) means only test for existence
        // returned size -1 indicates keyId not found (or error)
        // this removes the need for a separate hostExists function
        hostGetBytes(obj_id, key_id, std::ptr::null_mut(), -1) >= 0_i32
    }
}

pub fn get_bytes(obj_id: i32, key_id: i32) -> Vec<u8> {
    unsafe {
        // first query length of bytes array
        let size = hostGetBytes(obj_id, key_id, std::ptr::null_mut(), 0);
        if size <= 0 { return vec![0_u8; 0]; }

        // allocate a byte array in Wasm memory and
        // copy the actual data bytes to Wasm byte array
        let mut bytes = vec![0_u8; size as usize];
        hostGetBytes(obj_id, key_id, bytes.as_mut_ptr(), size);
        return bytes;
    }
}

pub fn get_int(obj_id: i32, key_id: i32) -> i64 {
    unsafe {
        hostGetInt(obj_id, key_id)
    }
}

pub fn get_key_id_from_bytes(bytes: &[u8]) -> i32 {
    unsafe {
        let size = bytes.len() as i32;
        // negative size indicates this was from bytes
        hostGetKeyId(bytes.as_ptr(), -size - 1)
    }
}

pub fn get_key_id_from_string(key: &str) -> i32 {
    let bytes = key.as_bytes();
    unsafe {
        // non-negative size indicates this was from string
        hostGetKeyId(bytes.as_ptr(), bytes.len() as i32)
    }
}

pub fn get_length(obj_id: i32) -> i32 {
    get_int(obj_id, KEY_LENGTH) as i32
}

pub fn get_object_id(obj_id: i32, key_id: i32, type_id: i32) -> i32 {
    unsafe {
        hostGetObjectId(obj_id, key_id, type_id)
    }
}

pub fn get_string(obj_id: i32, key_id: i32) -> String {
    // convert UTF8-encoded bytes array to string
    // negative object id indicates to host that this is a string
    // this removes the need for a separate hostGetString function
    unsafe {
        let bytes = get_bytes(-obj_id, key_id);
        return String::from_utf8_unchecked(bytes);
    }
}

pub fn set_bytes(obj_id: i32, key_id: i32, value: &[u8]) {
    unsafe {
        hostSetBytes(obj_id, key_id, value.as_ptr(), value.len() as i32)
    }
}

pub fn set_clear(obj_id: i32) {
    set_int(obj_id, KEY_LENGTH, 0)
}

pub fn set_int(obj_id: i32, key_id: i32, value: i64) {
    unsafe {
        hostSetInt(obj_id, key_id, value)
    }
}

pub fn set_string(obj_id: i32, key_id: i32, value: &str) {
    // convert string to UTF8-encoded bytes array
    // negative object id indicates to host that this is a string
    // this removes the need for a separate hostSetString function
    set_bytes(-obj_id, key_id, value.as_bytes())
}
