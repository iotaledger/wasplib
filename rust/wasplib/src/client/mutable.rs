// types encapsulating mutable host objects
// ScMutableBytes        : refers to mutable byte array on host
// ScMutableBytesArray   : refers to mutable array of mutable byte arrays on host
// ScMutableInt          : refers to mutable integer on host
// ScMutableIntArray     : refers to mutable array of mutable integers on host
// ScMutableMap          : refers to mutable map of mutable values on host
// ScMutableMapArray     : refers to mutable array of mutable maps of mutable values on host
// ScMutableString       : refers to mutable string on host
// ScMutableStringArray  : refers to mutable array of mutable strings on host

use super::host::{TYPE_BYTES_ARRAY, TYPE_INT_ARRAY, TYPE_MAP, TYPE_MAP_ARRAY, TYPE_STRING_ARRAY};
use super::host::{get_bytes, get_int, get_key_id, get_object_id, get_string, set_bytes, set_int, set_string};
use super::immutable::{ScImmutableBytesArray, ScImmutableIntArray, ScImmutableMap, ScImmutableMapArray, ScImmutableStringArray};
use super::keys::key_length;

#[derive(Copy, Clone)]
pub struct ScMutableBytes {
    obj_id: i32,
    key_id: i32,
}

impl ScMutableBytes {
    pub(crate) fn new(obj_id: i32, key_id: i32) -> ScMutableBytes {
        ScMutableBytes { obj_id, key_id }
    }

    pub fn set_value(&self, val: &[u8]) {
        set_bytes(self.obj_id, self.key_id, val);
    }

    pub fn value(&self) -> Vec<u8> {
        get_bytes(self.obj_id, self.key_id)
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

#[derive(Copy, Clone)]
pub struct ScMutableBytesArray {
    obj_id: i32
}

impl ScMutableBytesArray {
    pub(crate) fn new(obj_id: i32) -> ScMutableBytesArray {
        ScMutableBytesArray { obj_id }
    }

    pub fn clear(&self) {
        set_int(self.obj_id, key_length(), 0);
    }

    // index 0..length(), when length() a new one is appended
    pub fn get_bytes(&self, index: i32) -> ScMutableBytes {
        ScMutableBytes { obj_id: self.obj_id, key_id: index }
    }

    pub fn immutable(&self) -> ScImmutableBytesArray {
        ScImmutableBytesArray::new(self.obj_id)
    }

    pub fn length(&self) -> i32 {
        get_int(self.obj_id, key_length()) as i32
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

#[derive(Copy, Clone)]
pub struct ScMutableInt {
    obj_id: i32,
    key_id: i32,
}

impl ScMutableInt {
    pub(crate) fn new(obj_id: i32, key_id: i32) -> ScMutableInt {
        ScMutableInt { obj_id, key_id }
    }

    pub fn set_value(&self, val: i64) {
        set_int(self.obj_id, self.key_id, val);
    }

    pub fn value(&self) -> i64 {
        get_int(self.obj_id, self.key_id)
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

#[derive(Copy, Clone)]
pub struct ScMutableIntArray {
    obj_id: i32
}

impl ScMutableIntArray {
    pub(crate) fn new(obj_id: i32) -> ScMutableIntArray {
        ScMutableIntArray { obj_id }
    }

    pub fn clear(&self) {
        set_int(self.obj_id, key_length(), 0);
    }

    // index 0..length(), when length() a new one is appended
    pub fn get_int(&self, index: i32) -> ScMutableInt {
        ScMutableInt { obj_id: self.obj_id, key_id: index }
    }

    pub fn immutable(&self) -> ScImmutableIntArray {
        ScImmutableIntArray::new(self.obj_id)
    }

    pub fn length(&self) -> i32 {
        get_int(self.obj_id, key_length()) as i32
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

#[derive(Copy, Clone)]
pub struct ScMutableMap {
    obj_id: i32
}

impl ScMutableMap {
    pub(crate) fn new(obj_id: i32) -> ScMutableMap {
        ScMutableMap { obj_id }
    }

    pub fn clear(&self) {
        set_int(self.obj_id, key_length(), 0);
    }

    pub fn get_bytes(&self, key: &str) -> ScMutableBytes {
        ScMutableBytes { obj_id: self.obj_id, key_id: get_key_id(key) }
    }

    pub fn get_bytes_array(&self, key: &str) -> ScMutableBytesArray {
        let arr_id = get_object_id(self.obj_id, get_key_id(key), TYPE_BYTES_ARRAY);
        ScMutableBytesArray { obj_id: arr_id }
    }

    pub fn get_int(&self, key: &str) -> ScMutableInt {
        ScMutableInt { obj_id: self.obj_id, key_id: get_key_id(key) }
    }

    pub fn get_int_array(&self, key: &str) -> ScMutableIntArray {
        let arr_id = get_object_id(self.obj_id, get_key_id(key), TYPE_INT_ARRAY);
        ScMutableIntArray { obj_id: arr_id }
    }

    pub fn get_map(&self, key: &str) -> ScMutableMap {
        let map_id = get_object_id(self.obj_id, get_key_id(key), TYPE_MAP);
        ScMutableMap { obj_id: map_id }
    }

    pub fn get_map_array(&self, key: &str) -> ScMutableMapArray {
        let arr_id = get_object_id(self.obj_id, get_key_id(key), TYPE_MAP_ARRAY);
        ScMutableMapArray { obj_id: arr_id }
    }

    pub fn get_string(&self, key: &str) -> ScMutableString {
        ScMutableString { obj_id: self.obj_id, key_id: get_key_id(key) }
    }

    pub fn get_string_array(&self, key: &str) -> ScMutableStringArray {
        let arr_id = get_object_id(self.obj_id, get_key_id(key), TYPE_STRING_ARRAY);
        ScMutableStringArray { obj_id: arr_id }
    }

    pub fn immutable(&self) -> ScImmutableMap {
        ScImmutableMap::new(self.obj_id)
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

#[derive(Copy, Clone)]
pub struct ScMutableMapArray {
    obj_id: i32
}

impl ScMutableMapArray {
    pub(crate) fn new(obj_id: i32) -> ScMutableMapArray {
        ScMutableMapArray { obj_id }
    }

    pub fn clear(&self) {
        set_int(self.obj_id, key_length(), 0);
    }

    // index 0..length(), when length() a new one is appended
    pub fn get_map(&self, index: i32) -> ScMutableMap {
        let map_id = get_object_id(self.obj_id, index, TYPE_MAP);
        ScMutableMap { obj_id: map_id }
    }

    pub fn immutable(&self) -> ScImmutableMapArray {
        ScImmutableMapArray::new(self.obj_id)
    }

    pub fn length(&self) -> i32 {
        get_int(self.obj_id, key_length()) as i32
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

#[derive(Copy, Clone)]
pub struct ScMutableString {
    obj_id: i32,
    key_id: i32,
}

impl ScMutableString {
    pub(crate) fn new(obj_id: i32, key_id: i32) -> ScMutableString {
        ScMutableString { obj_id, key_id }
    }

    pub fn set_value(&self, val: &str) {
        set_string(self.obj_id, self.key_id, val);
    }

    pub fn value(&self) -> String {
        get_string(self.obj_id, self.key_id)
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

#[derive(Copy, Clone)]
pub struct ScMutableStringArray {
    obj_id: i32
}

impl ScMutableStringArray {
    pub(crate) fn new(obj_id: i32) -> ScMutableStringArray {
        ScMutableStringArray { obj_id }
    }

    pub fn clear(&self) {
        set_int(self.obj_id, key_length(), 0);
    }

    // index 0..length(), when length() a new one is appended
    pub fn get_string(&self, index: i32) -> ScMutableString {
        ScMutableString { obj_id: self.obj_id, key_id: index }
    }

    pub fn immutable(&self) -> ScImmutableStringArray {
        ScImmutableStringArray::new(self.obj_id)
    }

    pub fn length(&self) -> i32 {
        get_int(self.obj_id, key_length()) as i32
    }
}
