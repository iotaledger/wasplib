// types encapsulating mutable host objects
// ScMutableAddress      : refers to mutable address on host
// ScMutableAddressArray : refers to mutable array of mutable address on host
// ScMutableBytes        : refers to mutable byte array on host
// ScMutableBytesArray   : refers to mutable array of mutable byte arrays on host
// ScMutableColor        : refers to mutable color on host
// ScMutableColorArray   : refers to mutable array of mutable color on host
// ScMutableInt          : refers to mutable integer on host
// ScMutableIntArray     : refers to mutable array of mutable integers on host
// ScMutableKeyMap       : refers to mutable map of mutable values on host
// ScMutableMap          : refers to mutable map of mutable values on host
// ScMutableMapArray     : refers to mutable array of mutable maps of mutable values on host
// ScMutableString       : refers to mutable string on host
// ScMutableStringArray  : refers to mutable array of mutable strings on host

use super::hashtypes::*;
use super::host::*;
use super::immutable::*;
use super::keys::key_length;

pub struct ScMutableAddress {
    obj_id: i32,
    key_id: i32,
}

impl ScMutableAddress {
    pub(crate) fn new(obj_id: i32, key_id: i32) -> ScMutableAddress {
        ScMutableAddress { obj_id, key_id }
    }

    pub fn exists(&self) -> bool {
        exists(self.obj_id, self.key_id)
    }

    pub fn set_value(&self, val: &ScAddress) {
        set_bytes(self.obj_id, self.key_id, val.to_bytes());
    }

    pub fn value(&self) -> ScAddress {
        ScAddress::from_bytes(&get_bytes(self.obj_id, self.key_id))
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScMutableAddressArray {
    obj_id: i32
}

impl ScMutableAddressArray {
    pub(crate) fn new(obj_id: i32) -> ScMutableAddressArray {
        ScMutableAddressArray { obj_id }
    }

    pub fn clear(&self) {
        set_int(self.obj_id, key_length(), 0);
    }

    //TODO exists on arrays?

    // index 0..length(), when length() a new one is appended
    pub fn get_address(&self, index: i32) -> ScMutableAddress {
        ScMutableAddress { obj_id: self.obj_id, key_id: index }
    }

    pub fn immutable(&self) -> ScImmutableAddressArray {
        ScImmutableAddressArray::new(self.obj_id)
    }

    pub fn length(&self) -> i32 {
        get_int(self.obj_id, key_length()) as i32
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScMutableBytes {
    obj_id: i32,
    key_id: i32,
}

impl ScMutableBytes {
    pub(crate) fn new(obj_id: i32, key_id: i32) -> ScMutableBytes {
        ScMutableBytes { obj_id, key_id }
    }

    pub fn exists(&self) -> bool {
        exists(self.obj_id, self.key_id)
    }

    pub fn set_value(&self, val: &[u8]) {
        set_bytes(self.obj_id, self.key_id, val);
    }

    pub fn value(&self) -> Vec<u8> {
        get_bytes(self.obj_id, self.key_id)
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

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

pub struct ScMutableColor {
    obj_id: i32,
    key_id: i32,
}

impl ScMutableColor {
    pub(crate) fn new(obj_id: i32, key_id: i32) -> ScMutableColor {
        ScMutableColor { obj_id, key_id }
    }

    pub fn exists(&self) -> bool {
        exists(self.obj_id, self.key_id)
    }

    pub fn set_value(&self, val: &ScColor) {
        set_bytes(self.obj_id, self.key_id, val.to_bytes());
    }

    pub fn value(&self) -> ScColor {
        ScColor::from_bytes(&get_bytes(self.obj_id, self.key_id))
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScMutableColorArray {
    obj_id: i32
}

impl ScMutableColorArray {
    pub(crate) fn new(obj_id: i32) -> ScMutableColorArray {
        ScMutableColorArray { obj_id }
    }

    pub fn clear(&self) {
        set_int(self.obj_id, key_length(), 0);
    }

    // index 0..length(), when length() a new one is appended
    pub fn get_color(&self, index: i32) -> ScMutableColor {
        ScMutableColor { obj_id: self.obj_id, key_id: index }
    }

    pub fn immutable(&self) -> ScImmutableColorArray {
        ScImmutableColorArray::new(self.obj_id)
    }

    pub fn length(&self) -> i32 {
        get_int(self.obj_id, key_length()) as i32
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScMutableInt {
    obj_id: i32,
    key_id: i32,
}

impl ScMutableInt {
    pub(crate) fn new(obj_id: i32, key_id: i32) -> ScMutableInt {
        ScMutableInt { obj_id, key_id }
    }

    //TODO exists?

    pub fn set_value(&self, val: i64) {
        set_int(self.obj_id, self.key_id, val);
    }

    pub fn value(&self) -> i64 {
        get_int(self.obj_id, self.key_id)
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

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

pub struct ScMutableKeyMap {
    obj_id: i32
}

impl ScMutableKeyMap {
    pub(crate) fn new(obj_id: i32) -> ScMutableKeyMap {
        ScMutableKeyMap { obj_id }
    }

    pub fn clear(&self) {
        set_int(self.obj_id, key_length(), 0);
    }

    pub fn get_address(&self, key: &[u8]) -> ScMutableAddress {
        ScMutableAddress { obj_id: self.obj_id, key_id: get_key(key) }
    }

    pub fn get_address_array(&self, key: &[u8]) -> ScMutableAddressArray {
        let arr_id = get_object_id(self.obj_id, get_key(key), TYPE_BYTES_ARRAY);
        ScMutableAddressArray { obj_id: arr_id }
    }

    pub fn get_bytes(&self, key: &[u8]) -> ScMutableBytes {
        ScMutableBytes { obj_id: self.obj_id, key_id: get_key(key) }
    }

    pub fn get_bytes_array(&self, key: &[u8]) -> ScMutableBytesArray {
        let arr_id = get_object_id(self.obj_id, get_key(key), TYPE_BYTES_ARRAY);
        ScMutableBytesArray { obj_id: arr_id }
    }

    pub fn get_color(&self, key: &[u8]) -> ScMutableColor {
        ScMutableColor { obj_id: self.obj_id, key_id: get_key(key) }
    }

    pub fn get_color_array(&self, key: &[u8]) -> ScMutableColorArray {
        let arr_id = get_object_id(self.obj_id, get_key(key), TYPE_BYTES_ARRAY);
        ScMutableColorArray { obj_id: arr_id }
    }

    pub fn get_int(&self, key: &[u8]) -> ScMutableInt {
        ScMutableInt { obj_id: self.obj_id, key_id: get_key(key) }
    }

    pub fn get_int_array(&self, key: &[u8]) -> ScMutableIntArray {
        let arr_id = get_object_id(self.obj_id, get_key(key), TYPE_INT_ARRAY);
        ScMutableIntArray { obj_id: arr_id }
    }

    pub fn get_key_map(&self, key: &[u8]) -> ScMutableKeyMap {
        let map_id = get_object_id(self.obj_id, get_key(key), TYPE_MAP);
        ScMutableKeyMap { obj_id: map_id }
    }

    pub fn get_map(&self, key: &[u8]) -> ScMutableMap {
        let map_id = get_object_id(self.obj_id, get_key(key), TYPE_MAP);
        ScMutableMap { obj_id: map_id }
    }

    pub fn get_map_array(&self, key: &[u8]) -> ScMutableMapArray {
        let arr_id = get_object_id(self.obj_id, get_key(key), TYPE_MAP_ARRAY);
        ScMutableMapArray { obj_id: arr_id }
    }

    pub fn get_string(&self, key: &[u8]) -> ScMutableString {
        ScMutableString { obj_id: self.obj_id, key_id: get_key(key) }
    }

    pub fn get_string_array(&self, key: &[u8]) -> ScMutableStringArray {
        let arr_id = get_object_id(self.obj_id, get_key(key), TYPE_STRING_ARRAY);
        ScMutableStringArray { obj_id: arr_id }
    }

    pub fn immutable(&self) -> ScImmutableKeyMap {
        ScImmutableKeyMap::new(self.obj_id)
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

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

    pub fn get_address(&self, key: &str) -> ScMutableAddress {
        ScMutableAddress { obj_id: self.obj_id, key_id: get_key_id(key) }
    }

    pub fn get_address_array(&self, key: &str) -> ScMutableAddressArray {
        let arr_id = get_object_id(self.obj_id, get_key_id(key), TYPE_BYTES_ARRAY);
        ScMutableAddressArray { obj_id: arr_id }
    }

    pub fn get_bytes(&self, key: &str) -> ScMutableBytes {
        ScMutableBytes { obj_id: self.obj_id, key_id: get_key_id(key) }
    }

    pub fn get_bytes_array(&self, key: &str) -> ScMutableBytesArray {
        let arr_id = get_object_id(self.obj_id, get_key_id(key), TYPE_BYTES_ARRAY);
        ScMutableBytesArray { obj_id: arr_id }
    }

    pub fn get_color(&self, key: &str) -> ScMutableColor {
        ScMutableColor { obj_id: self.obj_id, key_id: get_key_id(key) }
    }

    pub fn get_color_array(&self, key: &str) -> ScMutableColorArray {
        let arr_id = get_object_id(self.obj_id, get_key_id(key), TYPE_BYTES_ARRAY);
        ScMutableColorArray { obj_id: arr_id }
    }

    pub fn get_int(&self, key: &str) -> ScMutableInt {
        ScMutableInt { obj_id: self.obj_id, key_id: get_key_id(key) }
    }

    pub fn get_int_array(&self, key: &str) -> ScMutableIntArray {
        let arr_id = get_object_id(self.obj_id, get_key_id(key), TYPE_INT_ARRAY);
        ScMutableIntArray { obj_id: arr_id }
    }

    pub fn get_key_map(&self, key: &str) -> ScMutableKeyMap {
        let map_id = get_object_id(self.obj_id, get_key_id(key), TYPE_MAP);
        ScMutableKeyMap { obj_id: map_id }
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

    // index 0..length(), inclusive, when length() a new one is appended
    pub fn get_key_map(&self, index: i32) -> ScMutableKeyMap {
        let map_id = get_object_id(self.obj_id, index, TYPE_MAP);
        ScMutableKeyMap { obj_id: map_id }
    }

    // index 0..length(), inclusive, hen length() a new one is appended
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

pub struct ScMutableString {
    obj_id: i32,
    key_id: i32,
}

impl ScMutableString {
    pub(crate) fn new(obj_id: i32, key_id: i32) -> ScMutableString {
        ScMutableString { obj_id, key_id }
    }

    pub fn exists(&self) -> bool {
        exists(self.obj_id, self.key_id)
    }

    pub fn set_value(&self, val: &str) {
        set_string(self.obj_id, self.key_id, val);
    }

    pub fn value(&self) -> String {
        get_string(self.obj_id, self.key_id)
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

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
