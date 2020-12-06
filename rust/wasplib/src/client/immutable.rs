// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// types encapsulating immutable host objects

use super::hashtypes::*;
use super::host::*;
use super::keys::*;

pub struct ScImmutableAddress {
    obj_id: i32,
    key_id: i32,
}

impl ScImmutableAddress {
    pub(crate) fn new(obj_id: i32, key_id: i32) -> ScImmutableAddress {
        ScImmutableAddress { obj_id, key_id }
    }

    pub fn exists(&self) -> bool {
        exists(self.obj_id, self.key_id)
    }

    pub fn value(&self) -> ScAddress {
        ScAddress::from_bytes(&get_bytes(self.obj_id, self.key_id))
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScImmutableAddressArray {
    obj_id: i32
}

impl ScImmutableAddressArray {
    pub(crate) fn new(obj_id: i32) -> ScImmutableAddressArray {
        ScImmutableAddressArray { obj_id }
    }

    //TODO exists on arrays?

    // index 0..length(), exclusive
    pub fn get_address(&self, index: i32) -> ScImmutableAddress {
        ScImmutableAddress { obj_id: self.obj_id, key_id: index }
    }

    pub fn length(&self) -> i32 {
        get_int(self.obj_id, key_length()) as i32
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScImmutableAgent {
    obj_id: i32,
    key_id: i32,
}

impl ScImmutableAgent {
    pub(crate) fn new(obj_id: i32, key_id: i32) -> ScImmutableAgent {
        ScImmutableAgent { obj_id, key_id }
    }

    pub fn exists(&self) -> bool {
        exists(self.obj_id, self.key_id)
    }

    pub fn value(&self) -> ScAgent {
        ScAgent::from_bytes(&get_bytes(self.obj_id, self.key_id))
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScImmutableAgentArray {
    obj_id: i32
}

impl ScImmutableAgentArray {
    pub(crate) fn new(obj_id: i32) -> ScImmutableAgentArray {
        ScImmutableAgentArray { obj_id }
    }

    //TODO exists on arrays?

    // index 0..length(), exclusive
    pub fn get_agent(&self, index: i32) -> ScImmutableAgent {
        ScImmutableAgent { obj_id: self.obj_id, key_id: index }
    }

    pub fn length(&self) -> i32 {
        get_int(self.obj_id, key_length()) as i32
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScImmutableBytes {
    obj_id: i32,
    key_id: i32,
}

impl ScImmutableBytes {
    pub(crate) fn new(obj_id: i32, key_id: i32) -> ScImmutableBytes {
        ScImmutableBytes { obj_id, key_id }
    }

    pub fn exists(&self) -> bool {
        exists(self.obj_id, self.key_id)
    }

    pub fn value(&self) -> Vec<u8> {
        get_bytes(self.obj_id, self.key_id)
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScImmutableBytesArray {
    obj_id: i32
}

impl ScImmutableBytesArray {
    pub(crate) fn new(obj_id: i32) -> ScImmutableBytesArray {
        ScImmutableBytesArray { obj_id }
    }

    // index 0..length(), exclusive
    pub fn get_bytes(&self, index: i32) -> ScImmutableBytes {
        ScImmutableBytes { obj_id: self.obj_id, key_id: index }
    }

    pub fn length(&self) -> i32 {
        get_int(self.obj_id, key_length()) as i32
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScImmutableColor {
    obj_id: i32,
    key_id: i32,
}

impl ScImmutableColor {
    pub(crate) fn new(obj_id: i32, key_id: i32) -> ScImmutableColor {
        ScImmutableColor { obj_id, key_id }
    }

    pub fn exists(&self) -> bool {
        exists(self.obj_id, self.key_id)
    }

    pub fn value(&self) -> ScColor {
        ScColor::from_bytes(&get_bytes(self.obj_id, self.key_id))
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScImmutableColorArray {
    obj_id: i32
}

impl ScImmutableColorArray {
    pub(crate) fn new(obj_id: i32) -> ScImmutableColorArray {
        ScImmutableColorArray { obj_id }
    }

    // index 0..length(), exclusive
    pub fn get_color(&self, index: i32) -> ScImmutableColor {
        ScImmutableColor { obj_id: self.obj_id, key_id: index }
    }

    pub fn length(&self) -> i32 {
        get_int(self.obj_id, key_length()) as i32
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScImmutableInt {
    obj_id: i32,
    key_id: i32,
}

impl ScImmutableInt {
    pub(crate) fn new(obj_id: i32, key_id: i32) -> ScImmutableInt {
        ScImmutableInt { obj_id, key_id }
    }

    pub fn exists(&self) -> bool {
        exists(self.obj_id, self.key_id)
    }

    pub fn value(&self) -> i64 {
        get_int(self.obj_id, self.key_id)
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScImmutableIntArray {
    obj_id: i32
}

impl ScImmutableIntArray {
    pub(crate) fn new(obj_id: i32) -> ScImmutableIntArray {
        ScImmutableIntArray { obj_id }
    }

    // index 0..length(), exclusive
    pub fn get_int(&self, index: i32) -> ScImmutableInt {
        ScImmutableInt { obj_id: self.obj_id, key_id: index }
    }

    pub fn length(&self) -> i32 {
        get_int(self.obj_id, key_length()) as i32
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScImmutableMap {
    obj_id: i32
}

impl ScImmutableMap {
    pub(crate) const fn new(obj_id: i32) -> ScImmutableMap {
        ScImmutableMap { obj_id }
    }

    pub fn get_address<T: KeyId + ?Sized>(&self, key: &T) -> ScImmutableAddress {
        ScImmutableAddress { obj_id: self.obj_id, key_id: key.get_id() }
    }

    pub fn get_address_array<T: KeyId + ?Sized>(&self, key: &T) -> ScImmutableAddressArray {
        let arr_id = get_object_id(self.obj_id, key.get_id(), TYPE_BYTES_ARRAY);
        ScImmutableAddressArray { obj_id: arr_id }
    }

    pub fn get_agent<T: KeyId + ?Sized>(&self, key: &T) -> ScImmutableAgent {
        ScImmutableAgent { obj_id: self.obj_id, key_id: key.get_id() }
    }

    pub fn get_agent_array<T: KeyId + ?Sized>(&self, key: &T) -> ScImmutableAgentArray {
        let arr_id = get_object_id(self.obj_id, key.get_id(), TYPE_BYTES_ARRAY);
        ScImmutableAgentArray { obj_id: arr_id }
    }

    pub fn get_bytes<T: KeyId + ?Sized>(&self, key: &T) -> ScImmutableBytes {
        ScImmutableBytes { obj_id: self.obj_id, key_id: key.get_id() }
    }

    pub fn get_bytes_array<T: KeyId + ?Sized>(&self, key: &T) -> ScImmutableBytesArray {
        let arr_id = get_object_id(self.obj_id, key.get_id(), TYPE_BYTES_ARRAY);
        ScImmutableBytesArray { obj_id: arr_id }
    }

    pub fn get_color<T: KeyId + ?Sized>(&self, key: &T) -> ScImmutableColor {
        ScImmutableColor { obj_id: self.obj_id, key_id: key.get_id() }
    }

    pub fn get_color_array<T: KeyId + ?Sized>(&self, key: &T) -> ScImmutableColorArray {
        let arr_id = get_object_id(self.obj_id, key.get_id(), TYPE_BYTES_ARRAY);
        ScImmutableColorArray { obj_id: arr_id }
    }

    pub fn get_int<T: KeyId + ?Sized>(&self, key: &T) -> ScImmutableInt {
        ScImmutableInt { obj_id: self.obj_id, key_id: key.get_id() }
    }

    pub fn get_int_array<T: KeyId + ?Sized>(&self, key: &T) -> ScImmutableIntArray {
        let arr_id = get_object_id(self.obj_id, key.get_id(), TYPE_INT_ARRAY);
        ScImmutableIntArray { obj_id: arr_id }
    }

    pub fn get_map<T: KeyId + ?Sized>(&self, key: &T) -> ScImmutableMap {
        let map_id = get_object_id(self.obj_id, key.get_id(), TYPE_MAP);
        ScImmutableMap { obj_id: map_id }
    }

    pub fn get_map_array<T: KeyId + ?Sized>(&self, key: &T) -> ScImmutableMapArray {
        let arr_id = get_object_id(self.obj_id, key.get_id(), TYPE_MAP_ARRAY);
        ScImmutableMapArray { obj_id: arr_id }
    }

    pub fn get_string<T: KeyId + ?Sized>(&self, key: &T) -> ScImmutableString {
        ScImmutableString { obj_id: self.obj_id, key_id: key.get_id() }
    }

    pub fn get_string_array<T: KeyId + ?Sized>(&self, key: &T) -> ScImmutableStringArray {
        let arr_id = get_object_id(self.obj_id, key.get_id(), TYPE_STRING_ARRAY);
        ScImmutableStringArray { obj_id: arr_id }
    }

    pub fn length(&self) -> i32 {
        get_int(self.obj_id, key_length()) as i32
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScImmutableMapArray {
    obj_id: i32
}

impl ScImmutableMapArray {
    pub(crate) fn new(obj_id: i32) -> ScImmutableMapArray {
        ScImmutableMapArray { obj_id }
    }

    // index 0..length(), exclusive
    pub fn get_map(&self, index: i32) -> ScImmutableMap {
        let map_id = get_object_id(self.obj_id, index, TYPE_MAP);
        ScImmutableMap { obj_id: map_id }
    }

    pub fn length(&self) -> i32 {
        get_int(self.obj_id, key_length()) as i32
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScImmutableString {
    obj_id: i32,
    key_id: i32,
}

impl ScImmutableString {
    pub(crate) fn new(obj_id: i32, key_id: i32) -> ScImmutableString {
        ScImmutableString { obj_id, key_id }
    }

    pub fn exists(&self) -> bool {
        exists(self.obj_id, self.key_id)
    }

    pub fn value(&self) -> String {
        get_string(self.obj_id, self.key_id)
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScImmutableStringArray {
    obj_id: i32
}

impl ScImmutableStringArray {
    pub(crate) fn new(obj_id: i32) -> ScImmutableStringArray {
        ScImmutableStringArray { obj_id }
    }

    // index 0..length(), exclusive
    pub fn get_string(&self, index: i32) -> ScImmutableString {
        ScImmutableString { obj_id: self.obj_id, key_id: index }
    }

    pub fn length(&self) -> i32 {
        get_int(self.obj_id, key_length()) as i32
    }
}
