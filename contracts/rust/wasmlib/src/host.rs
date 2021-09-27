// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// interface WasmLib to the VM host

use std::convert::TryInto;

use crate::keys::*;

// all type id values should exactly match their counterpart values on the host!
pub const TYPE_ARRAY: i32 = 0x20;
pub const TYPE_CALL: i32 = 0x40;

pub const TYPE_ADDRESS: i32 = 1;
pub const TYPE_AGENT_ID: i32 = 2;
pub const TYPE_BYTES: i32 = 3;
pub const TYPE_CHAIN_ID: i32 = 4;
pub const TYPE_COLOR: i32 = 5;
pub const TYPE_HASH: i32 = 6;
pub const TYPE_HNAME: i32 = 7;
pub const TYPE_INT16: i32 = 8;
pub const TYPE_INT32: i32 = 9;
pub const TYPE_INT64: i32 = 10;
pub const TYPE_MAP: i32 = 11;
pub const TYPE_REQUEST_ID: i32 = 12;
pub const TYPE_STRING: i32 = 13;

pub const OBJ_ID_NULL: i32 = 0;
pub const OBJ_ID_ROOT: i32 = 1;
pub const OBJ_ID_STATE: i32 = 2;
pub const OBJ_ID_PARAMS: i32 = 3;
pub const OBJ_ID_RESULTS: i32 = 4;

// size in bytes of predefined types, indexed by the TYPE_* consts
const TYPE_SIZES: &[u8] = &[0, 33, 37, 0, 33, 32, 32, 4, 2, 4, 8, 0, 34, 0];

// These 4 external functions are funneling the entire WasmLib functionality
// to their counterparts on the host.
#[link(wasm_import_module = "WasmLib")]
extern {
    // Copy the value data bytes of type <type_id> stored in the host container object <obj_id>,
    // under key <key_id>, into the pre-allocated <buffer> which can hold len bytes.
    // Returns the actual length of the value data bytes on the host.
    pub fn hostGetBytes(obj_id: i32, key_id: i32, type_id: i32, buffer: *const u8, len: i32) -> i32;

    // Retrieve the key id associated with the <key> data bytes of length <len>.
    // A negative length indicates a bytes key, positive indicates a string key
    // We discern between the two for better readable logging purposes
    pub fn hostGetKeyID(key: *const u8, len: i32) -> i32;

    // Retrieve the id of the container sub-object of type <type_id> stored in
    // the host container object <obj_id>, under key <key_id>.
    pub fn hostGetObjectID(obj_id: i32, key_id: i32, type_id: i32) -> i32;

    // copy the <len> value data bytes of type <type_id> from the <buffer>
    // into the host container object <obj_id>, under key <key_id>.
    pub fn hostSetBytes(obj_id: i32, key_id: i32, type_id: i32, buffer: *const u8, len: i32);
}

pub fn call_func(obj_id: i32, key_id: Key32, params: &[u8]) -> Vec<u8> {
    unsafe {
        let mut args = std::ptr::null();
        let mut size = params.len() as i32;
        if size != 0 {
            args = params.as_ptr();
        }

        // variable-sized type, first query expected length of bytes array
        // (pass zero-length buffer)
        size = hostGetBytes(obj_id, key_id.0, TYPE_CALL, args, size);

        // -1 means non-existent, so return default value for type
        if size <= 0 {
            return vec![0_u8; 0];
        }

        // allocate a sufficient length byte array in Wasm memory
        // and let the host copy the actual data bytes into this Wasm byte array
        let mut result = vec![0_u8; size as usize];
        hostGetBytes(obj_id, key_id.0, TYPE_CALL + 1, result.as_mut_ptr(), size);
        return result;
    }
}

// Clear the entire contents of the specified container object.
// Removes all its sub-objects as well.
pub fn clear(obj_id: i32) {
    // special key "length" is used with integer value zero
    set_bytes(obj_id, KEY_LENGTH, TYPE_INT32, &0_i32.to_le_bytes())
}

// Check if the specified container object contains a value with the specified key and type.
pub fn exists(obj_id: i32, key_id: Key32, type_id: i32) -> bool {
    unsafe {
        // negative length (-1) means only test for existence
        // returned size -1 indicates keyID not found (or error)
        // this removes the need for a separate hostExists function
        hostGetBytes(obj_id, key_id.0, type_id, std::ptr::null_mut(), -1) >= 0
    }
}

// Retrieve the bytes stored in the specified container object under the specified key
// and with specified type. Note that if the key does not exist this function will
// return the default value for the specified type.
pub fn get_bytes(obj_id: i32, key_id: Key32, type_id: i32) -> Vec<u8> {
    unsafe {
        let mut size = TYPE_SIZES[type_id as usize] as i32;
        if size == 0 {
            // variable-sized type, first query expected length of bytes array
            // (pass zero-length buffer)
            size = hostGetBytes(obj_id, key_id.0, type_id, std::ptr::null_mut(), 0);

            // -1 means non-existent, so return default value for type
            if size < 0 {
                return vec![0_u8; 0];
            }
        }

        // allocate a sufficient length byte array in Wasm memory
        // and let the host copy the actual data bytes into this Wasm byte array
        let mut result = vec![0_u8; size as usize];
        hostGetBytes(obj_id, key_id.0, type_id, result.as_mut_ptr(), size);
        return result;
    }
}

// Retrieve the key id that the host has associated with the specified bytes key
pub fn get_key_id_from_bytes(bytes: &[u8]) -> Key32 {
    unsafe {
        let size = bytes.len() as i32;
        // negative size indicates this is a bytes key
        Key32(hostGetKeyID(bytes.as_ptr(), -size - 1))
    }
}

// Retrieve the key id that the host has associated with the specified string key
pub fn get_key_id_from_string(key: &str) -> Key32 {
    let bytes = key.as_bytes();
    unsafe {
        // non-negative size indicates this is a string key
        Key32(hostGetKeyID(bytes.as_ptr(), bytes.len() as i32))
    }
}

// Retrieve the length of an array container object on the host
pub fn get_length(obj_id: i32) -> i32 {
    // special integer key "length" is used
    let bytes = get_bytes(obj_id, KEY_LENGTH, TYPE_INT32);
    i32::from_le_bytes(bytes.try_into().unwrap())
}

// Retrieve the id of the specified container sub-object
pub fn get_object_id(obj_id: i32, key_id: Key32, type_id: i32) -> i32 {
    unsafe {
        hostGetObjectID(obj_id, key_id.0, type_id)
    }
}

// Direct logging of informational text to host log
pub fn log(text: &str) {
    set_bytes(1, KEY_LOG, TYPE_STRING, text.as_bytes())
}

// Direct logging of error to host log, followed by panicking out of the Wasm code
pub fn panic(text: &str) {
    set_bytes(1, KEY_PANIC, TYPE_STRING, text.as_bytes())
}

// Store the provided value bytes of specified type in the specified container object
// under the specified key. Note that if the key does not exist this function will
// create it first.
pub fn set_bytes(obj_id: i32, key_id: Key32, type_id: i32, value: &[u8]) {
    unsafe {
        hostSetBytes(obj_id, key_id.0, type_id, value.as_ptr(), value.len() as i32)
    }
}

// Direct logging of debug trace text to host log
pub fn trace(text: &str) {
    set_bytes(1, KEY_TRACE, TYPE_STRING, text.as_bytes())
}
