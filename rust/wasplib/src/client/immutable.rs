// types encapsulating immutable host objects
// ScImmutableBytes        : refers to immutable byte array on host
// ScImmutableBytesArray   : refers to immutable array of immutable byte arrays on host
// ScImmutableInt          : refers to immutable integer on host
// ScImmutableIntArray     : refers to immutable array of immutable integers on host
// ScImmutableKeyMap       : refers to immutable map of immutable values on host
// ScImmutableMap          : refers to immutable map of immutable values on host
// ScImmutableMapArray     : refers to immutable array of immutable maps of immutable values on host
// ScImmutableString       : refers to immutable string on host
// ScImmutableStringArray  : refers to immutable array of immutable strings on host

use super::host::{TYPE_BYTES_ARRAY, TYPE_INT_ARRAY, TYPE_MAP, TYPE_MAP_ARRAY, TYPE_STRING_ARRAY};
use super::host::{get_bytes, get_int, get_key, get_key_id, get_object_id, get_string};
use super::keys::key_length;

pub struct ScImmutableBytes {
    obj_id: i32,
    key_id: i32,
}

impl ScImmutableBytes {
    pub(crate) fn new(obj_id: i32, key_id: i32) -> ScImmutableBytes {
        ScImmutableBytes { obj_id, key_id }
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

pub struct ScImmutableInt {
    obj_id: i32,
    key_id: i32,
}

impl ScImmutableInt {
    pub(crate) fn new(obj_id: i32, key_id: i32) -> ScImmutableInt {
        ScImmutableInt { obj_id, key_id }
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

pub struct ScImmutableKeyMap {
    obj_id: i32
}

impl ScImmutableKeyMap {
    pub(crate) fn new(obj_id: i32) -> ScImmutableKeyMap {
        ScImmutableKeyMap { obj_id }
    }

    pub fn get_bytes(&self, key: &[u8]) -> ScImmutableBytes {
        ScImmutableBytes { obj_id: self.obj_id, key_id: get_key(key) }
    }

    pub fn get_bytes_array(&self, key: &[u8]) -> ScImmutableBytesArray {
        let arr_id = get_object_id(self.obj_id, get_key(key), TYPE_BYTES_ARRAY);
        ScImmutableBytesArray { obj_id: arr_id }
    }

    pub fn get_int(&self, key: &[u8]) -> ScImmutableInt {
        ScImmutableInt { obj_id: self.obj_id, key_id: get_key(key) }
    }

    pub fn get_int_array(&self, key: &[u8]) -> ScImmutableIntArray {
        let arr_id = get_object_id(self.obj_id, get_key(key), TYPE_INT_ARRAY);
        ScImmutableIntArray { obj_id: arr_id }
    }

    pub fn get_key_map(&self, key: &[u8]) -> ScImmutableKeyMap {
        let map_id = get_object_id(self.obj_id, get_key(key), TYPE_MAP);
        ScImmutableKeyMap { obj_id: map_id }
    }

    pub fn get_map(&self, key: &[u8]) -> ScImmutableMap {
        let map_id = get_object_id(self.obj_id, get_key(key), TYPE_MAP);
        ScImmutableMap { obj_id: map_id }
    }

    pub fn get_map_array(&self, key: &[u8]) -> ScImmutableMapArray {
        let arr_id = get_object_id(self.obj_id, get_key(key), TYPE_MAP_ARRAY);
        ScImmutableMapArray { obj_id: arr_id }
    }

    pub fn get_string(&self, key: &[u8]) -> ScImmutableString {
        ScImmutableString { obj_id: self.obj_id, key_id: get_key(key) }
    }

    pub fn get_string_array(&self, key: &[u8]) -> ScImmutableStringArray {
        let arr_id = get_object_id(self.obj_id, get_key(key), TYPE_STRING_ARRAY);
        ScImmutableStringArray { obj_id: arr_id }
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
    pub(crate) fn new(obj_id: i32) -> ScImmutableMap {
        ScImmutableMap { obj_id }
    }

    pub fn get_bytes(&self, key: &str) -> ScImmutableBytes {
        ScImmutableBytes { obj_id: self.obj_id, key_id: get_key_id(key) }
    }

    pub fn get_bytes_array(&self, key: &str) -> ScImmutableBytesArray {
        let arr_id = get_object_id(self.obj_id, get_key_id(key), TYPE_BYTES_ARRAY);
        ScImmutableBytesArray { obj_id: arr_id }
    }

    pub fn get_int(&self, key: &str) -> ScImmutableInt {
        ScImmutableInt { obj_id: self.obj_id, key_id: get_key_id(key) }
    }

    pub fn get_int_array(&self, key: &str) -> ScImmutableIntArray {
        let arr_id = get_object_id(self.obj_id, get_key_id(key), TYPE_INT_ARRAY);
        ScImmutableIntArray { obj_id: arr_id }
    }

    pub fn get_key_map(&self, key: &str) -> ScImmutableKeyMap {
        let map_id = get_object_id(self.obj_id, get_key_id(key), TYPE_MAP);
        ScImmutableKeyMap { obj_id: map_id }
    }

    pub fn get_map(&self, key: &str) -> ScImmutableMap {
        let map_id = get_object_id(self.obj_id, get_key_id(key), TYPE_MAP);
        ScImmutableMap { obj_id: map_id }
    }

    pub fn get_map_array(&self, key: &str) -> ScImmutableMapArray {
        let arr_id = get_object_id(self.obj_id, get_key_id(key), TYPE_MAP_ARRAY);
        ScImmutableMapArray { obj_id: arr_id }
    }

    pub fn get_string(&self, key: &str) -> ScImmutableString {
        ScImmutableString { obj_id: self.obj_id, key_id: get_key_id(key) }
    }

    pub fn get_string_array(&self, key: &str) -> ScImmutableStringArray {
        let arr_id = get_object_id(self.obj_id, get_key_id(key), TYPE_STRING_ARRAY);
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
    pub fn get_key_map(&self, index: i32) -> ScImmutableKeyMap {
        let map_id = get_object_id(self.obj_id, index, TYPE_MAP);
        ScImmutableKeyMap { obj_id: map_id }
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
