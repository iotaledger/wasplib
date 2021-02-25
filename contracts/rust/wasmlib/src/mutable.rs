// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// types encapsulating mutable host objects

use std::convert::TryInto;

use crate::context::*;
use crate::hashtypes::*;
use crate::host::*;
use crate::immutable::*;
use crate::keys::*;

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// proxy object for mutable ScAddress in host map
pub struct ScMutableAddress {
    obj_id: i32,
    key_id: Key32,
}

impl ScMutableAddress {
    // check if object exists in host map
    pub fn exists(&self) -> bool {
        exists(self.obj_id, self.key_id, TYPE_ADDRESS)
    }

    // set value in host map
    pub fn set_value(&self, val: &ScAddress) {
        set_bytes(self.obj_id, self.key_id, TYPE_ADDRESS, val.to_bytes());
    }

    // human-readable string representation
    pub fn to_string(&self) -> String {
        self.value().to_string()
    }

    // retrieve value from host map
    pub fn value(&self) -> ScAddress {
        ScAddress::from_bytes(&get_bytes(self.obj_id, self.key_id, TYPE_ADDRESS))
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// mutable array of ScAddress
pub struct ScMutableAddressArray {
    pub(crate) obj_id: i32
}

impl ScMutableAddressArray {
    // empty the array
    pub fn clear(&self) {
        clear(self.obj_id);
    }

    // index 0..length(), when length() a new one is appended
    pub fn get_address(&self, index: i32) -> ScMutableAddress {
        ScMutableAddress { obj_id: self.obj_id, key_id: Key32(index) }
    }

    // get immutable version of array
    pub fn immutable(&self) -> ScImmutableAddressArray {
        ScImmutableAddressArray { obj_id: self.obj_id }
    }

    // number of items in array
    pub fn length(&self) -> i32 {
        get_length(self.obj_id)
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// proxy object for mutable ScAgentId in host map
pub struct ScMutableAgentId {
    obj_id: i32,
    key_id: Key32,
}

impl ScMutableAgentId {
    // check if object exists in host map
    pub fn exists(&self) -> bool {
        exists(self.obj_id, self.key_id, TYPE_AGENT_ID)
    }

    // set value in host map
    pub fn set_value(&self, val: &ScAgentId) {
        set_bytes(self.obj_id, self.key_id, TYPE_AGENT_ID, val.to_bytes());
    }

    // human-readable string representation
    pub fn to_string(&self) -> String {
        self.value().to_string()
    }

    // retrieve value from host map
    pub fn value(&self) -> ScAgentId {
        ScAgentId::from_bytes(&get_bytes(self.obj_id, self.key_id, TYPE_AGENT_ID))
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// mutable array of ScAgentId
pub struct ScMutableAgentIdArray {
    pub(crate) obj_id: i32
}

impl ScMutableAgentIdArray {
    // empty the array
    pub fn clear(&self) {
        clear(self.obj_id);
    }

    // index 0..length(), when length() a new one is appended
    pub fn get_agent_id(&self, index: i32) -> ScMutableAgentId {
        ScMutableAgentId { obj_id: self.obj_id, key_id: Key32(index) }
    }

    // get immutable version of array
    pub fn immutable(&self) -> ScImmutableAgentIdArray {
        ScImmutableAgentIdArray { obj_id: self.obj_id }
    }

    // number of items in array
    pub fn length(&self) -> i32 {
        get_length(self.obj_id)
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// proxy object for mutable bytes array in host map
pub struct ScMutableBytes {
    obj_id: i32,
    key_id: Key32,
}

impl ScMutableBytes {
    // check if object exists in host map
    pub fn exists(&self) -> bool {
        exists(self.obj_id, self.key_id, TYPE_BYTES)
    }

    // set value in host map
    pub fn set_value(&self, val: &[u8]) {
        set_bytes(self.obj_id, self.key_id, TYPE_BYTES, val);
    }

    // human-readable string representation
    pub fn to_string(&self) -> String {
        base58_encode(&self.value())
    }

    // retrieve value from host map
    pub fn value(&self) -> Vec<u8> {
        get_bytes(self.obj_id, self.key_id, TYPE_BYTES)
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// mutable array of byte array
pub struct ScMutableBytesArray {
    pub(crate) obj_id: i32
}

impl ScMutableBytesArray {
    // empty the array
    pub fn clear(&self) {
        clear(self.obj_id);
    }

    // index 0..length(), when length() a new one is appended
    pub fn get_bytes(&self, index: i32) -> ScMutableBytes {
        ScMutableBytes { obj_id: self.obj_id, key_id: Key32(index) }
    }

    // get immutable version of array
    pub fn immutable(&self) -> ScImmutableBytesArray {
        ScImmutableBytesArray { obj_id: self.obj_id }
    }

    // number of items in array
    pub fn length(&self) -> i32 {
        get_length(self.obj_id)
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// proxy object for mutable ScChainId in host map
pub struct ScMutableChainId {
    obj_id: i32,
    key_id: Key32,
}

impl ScMutableChainId {
    // check if object exists in host map
    pub fn exists(&self) -> bool {
        exists(self.obj_id, self.key_id, TYPE_CHAIN_ID)
    }

    // set value in host map
    pub fn set_value(&self, val: &ScChainId) {
        set_bytes(self.obj_id, self.key_id, TYPE_CHAIN_ID, val.to_bytes());
    }

    // human-readable string representation
    pub fn to_string(&self) -> String {
        self.value().to_string()
    }

    // retrieve value from host map
    pub fn value(&self) -> ScChainId {
        ScChainId::from_bytes(&get_bytes(self.obj_id, self.key_id, TYPE_CHAIN_ID))
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// mutable array of ScChainId
pub struct ScMutableChainIdArray {
    pub(crate) obj_id: i32
}

impl ScMutableChainIdArray {
    // empty the array
    pub fn clear(&self) {
        clear(self.obj_id);
    }

    // index 0..length(), when length() a new one is appended
    pub fn get_chain_id(&self, index: i32) -> ScMutableChainId {
        ScMutableChainId { obj_id: self.obj_id, key_id: Key32(index) }
    }

    // get immutable version of array
    pub fn immutable(&self) -> ScImmutableChainIdArray {
        ScImmutableChainIdArray { obj_id: self.obj_id }
    }

    // number of items in array
    pub fn length(&self) -> i32 {
        get_length(self.obj_id)
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// proxy object for mutable ScColor in host map
pub struct ScMutableColor {
    obj_id: i32,
    key_id: Key32,
}

impl ScMutableColor {
    // check if object exists in host map
    pub fn exists(&self) -> bool {
        exists(self.obj_id, self.key_id, TYPE_COLOR)
    }

    // set value in host map
    pub fn set_value(&self, val: &ScColor) {
        set_bytes(self.obj_id, self.key_id, TYPE_COLOR, val.to_bytes());
    }

    // human-readable string representation
    pub fn to_string(&self) -> String {
        self.value().to_string()
    }

    // retrieve value from host map
    pub fn value(&self) -> ScColor {
        ScColor::from_bytes(&get_bytes(self.obj_id, self.key_id, TYPE_COLOR))
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// mutable array of ScColor
pub struct ScMutableColorArray {
    pub(crate) obj_id: i32
}

impl ScMutableColorArray {
    // empty the array
    pub fn clear(&self) {
        clear(self.obj_id);
    }

    // index 0..length(), when length() a new one is appended
    pub fn get_color(&self, index: i32) -> ScMutableColor {
        ScMutableColor { obj_id: self.obj_id, key_id: Key32(index) }
    }

    // get immutable version of array
    pub fn immutable(&self) -> ScImmutableColorArray {
        ScImmutableColorArray { obj_id: self.obj_id }
    }

    // number of items in array
    pub fn length(&self) -> i32 {
        get_length(self.obj_id)
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// proxy object for mutable ScContractId in host map
pub struct ScMutableContractId {
    obj_id: i32,
    key_id: Key32,
}

impl ScMutableContractId {
    // check if object exists in host map
    pub fn exists(&self) -> bool {
        exists(self.obj_id, self.key_id, TYPE_CONTRACT_ID)
    }

    // set value in host map
    pub fn set_value(&self, val: &ScContractId) {
        set_bytes(self.obj_id, self.key_id, TYPE_CONTRACT_ID, val.to_bytes());
    }

    // human-readable string representation
    pub fn to_string(&self) -> String {
        self.value().to_string()
    }

    // retrieve value from host map
    pub fn value(&self) -> ScContractId {
        ScContractId::from_bytes(&get_bytes(self.obj_id, self.key_id, TYPE_CONTRACT_ID))
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// mutable array of ScContractId
pub struct ScMutableContractIdArray {
    pub(crate) obj_id: i32
}

impl ScMutableContractIdArray {
    // empty the array
    pub fn clear(&self) {
        clear(self.obj_id);
    }

    // index 0..length(), when length() a new one is appended
    pub fn get_contract_id(&self, index: i32) -> ScMutableContractId {
        ScMutableContractId { obj_id: self.obj_id, key_id: Key32(index) }
    }

    // get immutable version of array
    pub fn immutable(&self) -> ScImmutableContractIdArray {
        ScImmutableContractIdArray { obj_id: self.obj_id }
    }

    // number of items in array
    pub fn length(&self) -> i32 {
        get_length(self.obj_id)
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// proxy object for mutable ScHash in host map
pub struct ScMutableHash {
    obj_id: i32,
    key_id: Key32,
}

impl ScMutableHash {
    // check if object exists in host map
    pub fn exists(&self) -> bool {
        exists(self.obj_id, self.key_id, TYPE_HASH)
    }

    // set value in host map
    pub fn set_value(&self, val: &ScHash) {
        set_bytes(self.obj_id, self.key_id, TYPE_HASH, val.to_bytes());
    }

    // human-readable string representation
    pub fn to_string(&self) -> String {
        self.value().to_string()
    }

    // retrieve value from host map
    pub fn value(&self) -> ScHash {
        ScHash::from_bytes(&get_bytes(self.obj_id, self.key_id, TYPE_HASH))
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// mutable array of ScHash
pub struct ScMutableHashArray {
    pub(crate) obj_id: i32
}

impl ScMutableHashArray {
    // empty the array
    pub fn clear(&self) {
        clear(self.obj_id);
    }

    // index 0..length(), when length() a new one is appended
    pub fn get_hash(&self, index: i32) -> ScMutableHash {
        ScMutableHash { obj_id: self.obj_id, key_id: Key32(index) }
    }

    // get immutable version of array
    pub fn immutable(&self) -> ScImmutableHashArray {
        ScImmutableHashArray { obj_id: self.obj_id }
    }

    // number of items in array
    pub fn length(&self) -> i32 {
        get_length(self.obj_id)
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// proxy object for mutable ScHname in host map
pub struct ScMutableHname {
    obj_id: i32,
    key_id: Key32,
}

impl ScMutableHname {
    // check if object exists in host map
    pub fn exists(&self) -> bool {
        exists(self.obj_id, self.key_id, TYPE_HNAME)
    }

    // set value in host map
    pub fn set_value(&self, val: ScHname) {
        set_bytes(self.obj_id, self.key_id, TYPE_HNAME, &val.to_bytes());
    }

    // human-readable string representation
    pub fn to_string(&self) -> String {
        self.value().to_string()
    }

    // retrieve value from host map
    pub fn value(&self) -> ScHname {
        ScHname::from_bytes(&get_bytes(self.obj_id, self.key_id, TYPE_HNAME))
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// mutable array of ScHname
pub struct ScMutableHnameArray {
    pub(crate) obj_id: i32
}

impl ScMutableHnameArray {
    // empty the array
    pub fn clear(&self) {
        clear(self.obj_id);
    }

    // index 0..length(), when length() a new one is appended
    pub fn get_hname(&self, index: i32) -> ScMutableHname {
        ScMutableHname { obj_id: self.obj_id, key_id: Key32(index) }
    }

    // get immutable version of array
    pub fn immutable(&self) -> ScImmutableHnameArray {
        ScImmutableHnameArray { obj_id: self.obj_id }
    }

    // number of items in array
    pub fn length(&self) -> i32 {
        get_length(self.obj_id)
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// proxy object for mutable int64 in host map
pub struct ScMutableInt64 {
    obj_id: i32,
    key_id: Key32,
}

impl ScMutableInt64 {
    // check if object exists in host map
    pub fn exists(&self) -> bool {
        exists(self.obj_id, self.key_id, TYPE_INT64)
    }

    // set value in host map
    pub fn set_value(&self, val: i64) {
        set_bytes(self.obj_id, self.key_id, TYPE_INT64, &val.to_le_bytes());
    }

    // human-readable string representation
    pub fn to_string(&self) -> String {
        self.value().to_string()
    }

    // retrieve value from host map
    pub fn value(&self) -> i64 {
        let bytes = get_bytes(self.obj_id, self.key_id, TYPE_INT64);
        i64::from_le_bytes(bytes.try_into().expect("invalid i64 length"))
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// mutable array of int64
pub struct ScMutableInt64Array {
    pub(crate) obj_id: i32
}

impl ScMutableInt64Array {
    // empty the array
    pub fn clear(&self) {
        clear(self.obj_id);
    }

    // index 0..length(), when length() a new one is appended
    pub fn get_int64(&self, index: i32) -> ScMutableInt64 {
        ScMutableInt64 { obj_id: self.obj_id, key_id: Key32(index) }
    }

    // get immutable version of array
    pub fn immutable(&self) -> ScImmutableInt64Array {
        ScImmutableInt64Array { obj_id: self.obj_id }
    }

    // number of items in array
    pub fn length(&self) -> i32 {
        get_length(self.obj_id)
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScMutableMap {
    pub(crate) obj_id: i32
}

impl ScMutableMap {
    pub fn new() -> ScMutableMap {
        let maps = ROOT.get_map_array(&KEY_MAPS);
        maps.get_map(maps.length())
    }

    // empty the map
    pub fn clear(&self) {
        clear(self.obj_id);
    }

    // get proxy for mutable ScAddress field specified by key
    pub fn get_address<T: MapKey + ?Sized>(&self, key: &T) -> ScMutableAddress {
        ScMutableAddress { obj_id: self.obj_id, key_id: key.get_id() }
    }

    // get proxy for ScMutableAddressArray specified by key
    pub fn get_address_array<T: MapKey + ?Sized>(&self, key: &T) -> ScMutableAddressArray {
        let arr_id = get_object_id(self.obj_id, key.get_id(), TYPE_ADDRESS | TYPE_ARRAY);
        ScMutableAddressArray { obj_id: arr_id }
    }

    // get proxy for mutable ScAgentId field specified by key
    pub fn get_agent_id<T: MapKey + ?Sized>(&self, key: &T) -> ScMutableAgentId {
        ScMutableAgentId { obj_id: self.obj_id, key_id: key.get_id() }
    }

    // get proxy for ScMutableAgentIdArray specified by key
    pub fn get_agent_id_array<T: MapKey + ?Sized>(&self, key: &T) -> ScMutableAgentIdArray {
        let arr_id = get_object_id(self.obj_id, key.get_id(), TYPE_AGENT_ID | TYPE_ARRAY);
        ScMutableAgentIdArray { obj_id: arr_id }
    }

    // get proxy for mutable bytes array field specified by key
    pub fn get_bytes<T: MapKey + ?Sized>(&self, key: &T) -> ScMutableBytes {
        ScMutableBytes { obj_id: self.obj_id, key_id: key.get_id() }
    }

    // get proxy for ScMutableBytesArray specified by key
    pub fn get_bytes_array<T: MapKey + ?Sized>(&self, key: &T) -> ScMutableBytesArray {
        let arr_id = get_object_id(self.obj_id, key.get_id(), TYPE_BYTES | TYPE_ARRAY);
        ScMutableBytesArray { obj_id: arr_id }
    }

    // get proxy for mutable ScChainId field specified by key
    pub fn get_chain_id<T: MapKey + ?Sized>(&self, key: &T) -> ScMutableChainId {
        ScMutableChainId { obj_id: self.obj_id, key_id: key.get_id() }
    }

    // get proxy for ScMutableChainIdArray specified by key
    pub fn get_chain_id_array<T: MapKey + ?Sized>(&self, key: &T) -> ScMutableChainIdArray {
        let arr_id = get_object_id(self.obj_id, key.get_id(), TYPE_CHAIN_ID | TYPE_ARRAY);
        ScMutableChainIdArray { obj_id: arr_id }
    }

    // get proxy for mutable ScColor field specified by key
    pub fn get_color<T: MapKey + ?Sized>(&self, key: &T) -> ScMutableColor {
        ScMutableColor { obj_id: self.obj_id, key_id: key.get_id() }
    }

    // get proxy for ScMutableColorArray specified by key
    pub fn get_color_array<T: MapKey + ?Sized>(&self, key: &T) -> ScMutableColorArray {
        let arr_id = get_object_id(self.obj_id, key.get_id(), TYPE_COLOR | TYPE_ARRAY);
        ScMutableColorArray { obj_id: arr_id }
    }

    // get proxy for mutable ScContractId field specified by key
    pub fn get_contract_id<T: MapKey + ?Sized>(&self, key: &T) -> ScMutableContractId {
        ScMutableContractId { obj_id: self.obj_id, key_id: key.get_id() }
    }

    // get proxy for ScMutableContractIdArray specified by key
    pub fn get_contract_id_array<T: MapKey + ?Sized>(&self, key: &T) -> ScMutableContractIdArray {
        let arr_id = get_object_id(self.obj_id, key.get_id(), TYPE_CONTRACT_ID | TYPE_ARRAY);
        ScMutableContractIdArray { obj_id: arr_id }
    }

    // get proxy for mutable ScHash field specified by key
    pub fn get_hash<T: MapKey + ?Sized>(&self, key: &T) -> ScMutableHash {
        ScMutableHash { obj_id: self.obj_id, key_id: key.get_id() }
    }

    // get proxy for ScMutableHashArray specified by key
    pub fn get_hash_array<T: MapKey + ?Sized>(&self, key: &T) -> ScMutableHashArray {
        let arr_id = get_object_id(self.obj_id, key.get_id(), TYPE_HASH | TYPE_ARRAY);
        ScMutableHashArray { obj_id: arr_id }
    }

    // get proxy for mutable ScHname field specified by key
    pub fn get_hname<T: MapKey + ?Sized>(&self, key: &T) -> ScMutableHname {
        ScMutableHname { obj_id: self.obj_id, key_id: key.get_id() }
    }

    // get proxy for ScMutableHnameArray specified by key
    pub fn get_hname_array<T: MapKey + ?Sized>(&self, key: &T) -> ScMutableHnameArray {
        let arr_id = get_object_id(self.obj_id, key.get_id(), TYPE_HNAME | TYPE_ARRAY);
        ScMutableHnameArray { obj_id: arr_id }
    }

    // get proxy for mutable int64 field specified by key
    pub fn get_int64<T: MapKey + ?Sized>(&self, key: &T) -> ScMutableInt64 {
        ScMutableInt64 { obj_id: self.obj_id, key_id: key.get_id() }
    }

    // get proxy for ScMutableInt64Array specified by key
    pub fn get_int64_array<T: MapKey + ?Sized>(&self, key: &T) -> ScMutableInt64Array {
        let arr_id = get_object_id(self.obj_id, key.get_id(), TYPE_INT64 | TYPE_ARRAY);
        ScMutableInt64Array { obj_id: arr_id }
    }

    // get proxy for ScMutableMap specified by key
    pub fn get_map<T: MapKey + ?Sized>(&self, key: &T) -> ScMutableMap {
        let map_id = get_object_id(self.obj_id, key.get_id(), TYPE_MAP);
        ScMutableMap { obj_id: map_id }
    }

    // get proxy for ScMutableMapArray specified by key
    pub fn get_map_array<T: MapKey + ?Sized>(&self, key: &T) -> ScMutableMapArray {
        let arr_id = get_object_id(self.obj_id, key.get_id(), TYPE_MAP | TYPE_ARRAY);
        ScMutableMapArray { obj_id: arr_id }
    }

    // get proxy for mutable ScRequestId field specified by key
    pub fn get_request_id<T: MapKey + ?Sized>(&self, key: &T) -> ScMutableRequestId {
        ScMutableRequestId { obj_id: self.obj_id, key_id: key.get_id() }
    }

    // get proxy for ScMutableRequestIdArray specified by key
    pub fn get_request_id_array<T: MapKey + ?Sized>(&self, key: &T) -> ScMutableRequestIdArray {
        let arr_id = get_object_id(self.obj_id, key.get_id(), TYPE_REQUEST_ID | TYPE_ARRAY);
        ScMutableRequestIdArray { obj_id: arr_id }
    }

    // get proxy for mutable UTF-8 text string field specified by key
    pub fn get_string<T: MapKey + ?Sized>(&self, key: &T) -> ScMutableString {
        ScMutableString { obj_id: self.obj_id, key_id: key.get_id() }
    }

    // get proxy for ScMutableStringArray specified by key
    pub fn get_string_array<T: MapKey + ?Sized>(&self, key: &T) -> ScMutableStringArray {
        let arr_id = get_object_id(self.obj_id, key.get_id(), TYPE_STRING | TYPE_ARRAY);
        ScMutableStringArray { obj_id: arr_id }
    }

    // get immutable version of map
    pub fn immutable(&self) -> ScImmutableMap {
        ScImmutableMap { obj_id: self.obj_id }
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// mutable array of ScMap
pub struct ScMutableMapArray {
    pub(crate) obj_id: i32
}

impl ScMutableMapArray {
    // empty the array
    pub fn clear(&self) {
        clear(self.obj_id);
    }

    // index 0..length(), inclusive, hen length() a new one is appended
    pub fn get_map(&self, index: i32) -> ScMutableMap {
        let map_id = get_object_id(self.obj_id, Key32(index), TYPE_MAP);
        ScMutableMap { obj_id: map_id }
    }

    // get immutable version of array
    pub fn immutable(&self) -> ScImmutableMapArray {
        ScImmutableMapArray { obj_id: self.obj_id }
    }

    // number of items in array
    pub fn length(&self) -> i32 {
        get_length(self.obj_id)
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// proxy object for mutable ScRequestId in host map
pub struct ScMutableRequestId {
    obj_id: i32,
    key_id: Key32,
}

impl ScMutableRequestId {
    // check if object exists in host map
    pub fn exists(&self) -> bool {
        exists(self.obj_id, self.key_id, TYPE_REQUEST_ID)
    }

    // set value in host map
    pub fn set_value(&self, val: &ScRequestId) {
        set_bytes(self.obj_id, self.key_id, TYPE_REQUEST_ID, val.to_bytes());
    }

    // human-readable string representation
    pub fn to_string(&self) -> String {
        self.value().to_string()
    }

    // retrieve value from host map
    pub fn value(&self) -> ScRequestId {
        ScRequestId::from_bytes(&get_bytes(self.obj_id, self.key_id, TYPE_REQUEST_ID))
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// mutable array of ScRequestId
pub struct ScMutableRequestIdArray {
    pub(crate) obj_id: i32
}

impl ScMutableRequestIdArray {
    // empty the array
    pub fn clear(&self) {
        clear(self.obj_id);
    }

    // index 0..length(), when length() a new one is appended
    pub fn get_request_id(&self, index: i32) -> ScMutableRequestId {
        ScMutableRequestId { obj_id: self.obj_id, key_id: Key32(index) }
    }

    // get immutable version of array
    pub fn immutable(&self) -> ScImmutableRequestIdArray {
        ScImmutableRequestIdArray { obj_id: self.obj_id }
    }

    // number of items in array
    pub fn length(&self) -> i32 {
        get_length(self.obj_id)
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// proxy object for mutable UTF-8 text string in host map
pub struct ScMutableString {
    obj_id: i32,
    key_id: Key32,
}

impl ScMutableString {
    // check if object exists in host map
    pub fn exists(&self) -> bool {
        exists(self.obj_id, self.key_id, TYPE_STRING)
    }

    // set value in host map
    pub fn set_value(&self, val: &str) {
        set_bytes(self.obj_id, self.key_id, TYPE_STRING, val.as_bytes());
    }

    // human-readable string representation
    pub fn to_string(&self) -> String {
        self.value()
    }

    // retrieve value from host map
    pub fn value(&self) -> String {
        let bytes = get_bytes(self.obj_id, self.key_id, TYPE_STRING);
        unsafe { String::from_utf8_unchecked(bytes) }
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// mutable array of UTF-8 text string
pub struct ScMutableStringArray {
    pub(crate) obj_id: i32
}

impl ScMutableStringArray {
    // empty the array
    pub fn clear(&self) {
        clear(self.obj_id);
    }

    // index 0..length(), when length() a new one is appended
    pub fn get_string(&self, index: i32) -> ScMutableString {
        ScMutableString { obj_id: self.obj_id, key_id: Key32(index) }
    }

    // get immutable version of array
    pub fn immutable(&self) -> ScImmutableStringArray {
        ScImmutableStringArray { obj_id: self.obj_id }
    }

    // number of items in array
    pub fn length(&self) -> i32 {
        get_length(self.obj_id)
    }
}
