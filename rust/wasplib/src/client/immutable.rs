// types encapsulating immutable host objects

use super::hashtypes::*;
use super::host::*;
use super::keys::key_length;

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

    //TODO exists?

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

    pub fn get_address(&self, key: &[u8]) -> ScImmutableAddress {
        ScImmutableAddress { obj_id: self.obj_id, key_id: get_key(key) }
    }

    pub fn get_address_array(&self, key: &[u8]) -> ScImmutableAddressArray {
        let arr_id = get_object_id(self.obj_id, get_key(key), TYPE_BYTES_ARRAY);
        ScImmutableAddressArray { obj_id: arr_id }
    }

    pub fn get_bytes(&self, key: &[u8]) -> ScImmutableBytes {
        ScImmutableBytes { obj_id: self.obj_id, key_id: get_key(key) }
    }

    pub fn get_bytes_array(&self, key: &[u8]) -> ScImmutableBytesArray {
        let arr_id = get_object_id(self.obj_id, get_key(key), TYPE_BYTES_ARRAY);
        ScImmutableBytesArray { obj_id: arr_id }
    }

    pub fn get_color(&self, key: &[u8]) -> ScImmutableColor {
        ScImmutableColor { obj_id: self.obj_id, key_id: get_key(key) }
    }

    pub fn get_color_array(&self, key: &[u8]) -> ScImmutableColorArray {
        let arr_id = get_object_id(self.obj_id, get_key(key), TYPE_BYTES_ARRAY);
        ScImmutableColorArray { obj_id: arr_id }
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

    pub fn get_request_id(&self, key: &[u8]) -> ScImmutableRequestId {
        ScImmutableRequestId { obj_id: self.obj_id, key_id: get_key(key) }
    }

    pub fn get_request_id_array(&self, key: &[u8]) -> ScImmutableRequestIdArray {
        let arr_id = get_object_id(self.obj_id, get_key(key), TYPE_BYTES_ARRAY);
        ScImmutableRequestIdArray { obj_id: arr_id }
    }

    pub fn get_string(&self, key: &[u8]) -> ScImmutableString {
        ScImmutableString { obj_id: self.obj_id, key_id: get_key(key) }
    }

    pub fn get_string_array(&self, key: &[u8]) -> ScImmutableStringArray {
        let arr_id = get_object_id(self.obj_id, get_key(key), TYPE_STRING_ARRAY);
        ScImmutableStringArray { obj_id: arr_id }
    }

    pub fn get_tx_hash(&self, key: &[u8]) -> ScImmutableTxHash {
        ScImmutableTxHash { obj_id: self.obj_id, key_id: get_key(key) }
    }

    pub fn get_tx_hash_array(&self, key: &[u8]) -> ScImmutableTxHashArray {
        let arr_id = get_object_id(self.obj_id, get_key(key), TYPE_BYTES_ARRAY);
        ScImmutableTxHashArray { obj_id: arr_id }
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

    pub fn get_address(&self, key: &str) -> ScImmutableAddress {
        ScImmutableAddress { obj_id: self.obj_id, key_id: get_key_id(key) }
    }

    pub fn get_address_array(&self, key: &str) -> ScImmutableAddressArray {
        let arr_id = get_object_id(self.obj_id, get_key_id(key), TYPE_BYTES_ARRAY);
        ScImmutableAddressArray { obj_id: arr_id }
    }

    pub fn get_bytes(&self, key: &str) -> ScImmutableBytes {
        ScImmutableBytes { obj_id: self.obj_id, key_id: get_key_id(key) }
    }

    pub fn get_bytes_array(&self, key: &str) -> ScImmutableBytesArray {
        let arr_id = get_object_id(self.obj_id, get_key_id(key), TYPE_BYTES_ARRAY);
        ScImmutableBytesArray { obj_id: arr_id }
    }

    pub fn get_color(&self, key: &str) -> ScImmutableColor {
        ScImmutableColor { obj_id: self.obj_id, key_id: get_key_id(key) }
    }

    pub fn get_color_array(&self, key: &str) -> ScImmutableColorArray {
        let arr_id = get_object_id(self.obj_id, get_key_id(key), TYPE_BYTES_ARRAY);
        ScImmutableColorArray { obj_id: arr_id }
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

    pub fn get_request_id(&self, key: &str) -> ScImmutableRequestId {
        ScImmutableRequestId { obj_id: self.obj_id, key_id: get_key_id(key) }
    }

    pub fn get_request_id_array(&self, key: &str) -> ScImmutableRequestIdArray {
        let arr_id = get_object_id(self.obj_id, get_key_id(key), TYPE_BYTES_ARRAY);
        ScImmutableRequestIdArray { obj_id: arr_id }
    }

    pub fn get_string(&self, key: &str) -> ScImmutableString {
        ScImmutableString { obj_id: self.obj_id, key_id: get_key_id(key) }
    }

    pub fn get_string_array(&self, key: &str) -> ScImmutableStringArray {
        let arr_id = get_object_id(self.obj_id, get_key_id(key), TYPE_STRING_ARRAY);
        ScImmutableStringArray { obj_id: arr_id }
    }

    pub fn get_tx_hash(&self, key: &str) -> ScImmutableTxHash {
        ScImmutableTxHash { obj_id: self.obj_id, key_id: get_key_id(key) }
    }

    pub fn get_tx_hash_array(&self, key: &str) -> ScImmutableTxHashArray {
        let arr_id = get_object_id(self.obj_id, get_key_id(key), TYPE_BYTES_ARRAY);
        ScImmutableTxHashArray { obj_id: arr_id }
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

pub struct ScImmutableRequestId {
    obj_id: i32,
    key_id: i32,
}

impl ScImmutableRequestId {
    pub(crate) fn new(obj_id: i32, key_id: i32) -> ScImmutableRequestId {
        ScImmutableRequestId { obj_id, key_id }
    }

    pub fn exists(&self) -> bool {
        exists(self.obj_id, self.key_id)
    }

    pub fn value(&self) -> ScRequestId {
        ScRequestId::from_bytes(&get_bytes(self.obj_id, self.key_id))
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScImmutableRequestIdArray {
    obj_id: i32
}

impl ScImmutableRequestIdArray {
    pub(crate) fn new(obj_id: i32) -> ScImmutableRequestIdArray {
        ScImmutableRequestIdArray { obj_id }
    }

    // index 0..length(), exclusive
    pub fn get_request_id(&self, index: i32) -> ScImmutableRequestId {
        ScImmutableRequestId { obj_id: self.obj_id, key_id: index }
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

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScImmutableTxHash {
    obj_id: i32,
    key_id: i32,
}

impl ScImmutableTxHash {
    pub(crate) fn new(obj_id: i32, key_id: i32) -> ScImmutableTxHash {
        ScImmutableTxHash { obj_id, key_id }
    }

    pub fn exists(&self) -> bool {
        exists(self.obj_id, self.key_id)
    }

    pub fn value(&self) -> ScTxHash {
        ScTxHash::from_bytes(&get_bytes(self.obj_id, self.key_id))
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScImmutableTxHashArray {
    obj_id: i32
}

impl ScImmutableTxHashArray {
    pub(crate) fn new(obj_id: i32) -> ScImmutableTxHashArray {
        ScImmutableTxHashArray { obj_id }
    }

    // index 0..length(), exclusive
    pub fn get_tx_hash(&self, index: i32) -> ScImmutableTxHash {
        ScImmutableTxHash { obj_id: self.obj_id, key_id: index }
    }

    pub fn length(&self) -> i32 {
        get_int(self.obj_id, key_length()) as i32
    }
}
