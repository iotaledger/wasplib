// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// immutable proxies to host objects

use std::convert::TryInto;

use crate::context::*;
use crate::hashtypes::*;
use crate::host::*;
use crate::keys::*;

// value proxy for immutable ScAddress in host container
pub struct ScImmutableAddress {
    obj_id: i32,
    key_id: Key32,
}

impl ScImmutableAddress {
    // check if value exists in host container
    pub fn exists(&self) -> bool {
        exists(self.obj_id, self.key_id, TYPE_ADDRESS)
    }

    // human-readable string representation
    pub fn to_string(&self) -> String {
        self.value().to_string()
    }

    // get value from host container
    pub fn value(&self) -> ScAddress {
        ScAddress::from_bytes(&get_bytes(self.obj_id, self.key_id, TYPE_ADDRESS))
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// array proxy for immutable array of ScAddress
pub struct ScImmutableAddressArray {
    pub(crate) obj_id: i32
}

impl ScImmutableAddressArray {
    // get value proxy for item at index, index can be 0..length()-1
    pub fn get_address(&self, index: i32) -> ScImmutableAddress {
        ScImmutableAddress { obj_id: self.obj_id, key_id: Key32(index) }
    }

    // number of items in array
    pub fn length(&self) -> i32 {
        get_length(self.obj_id)
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// value proxy for immutable ScAgentId in host container
pub struct ScImmutableAgentId {
    obj_id: i32,
    key_id: Key32,
}

impl ScImmutableAgentId {
    // check if value exists in host container
    pub fn exists(&self) -> bool {
        exists(self.obj_id, self.key_id, TYPE_AGENT_ID)
    }

    // human-readable string representation
    pub fn to_string(&self) -> String {
        self.value().to_string()
    }

    // get value from host container
    pub fn value(&self) -> ScAgentId {
        ScAgentId::from_bytes(&get_bytes(self.obj_id, self.key_id, TYPE_AGENT_ID))
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// array proxy for immutable array of ScAgentId
pub struct ScImmutableAgentIdArray {
    pub(crate) obj_id: i32
}

impl ScImmutableAgentIdArray {
    // get value proxy for item at index, index can be 0..length()-1
    pub fn get_agent_id(&self, index: i32) -> ScImmutableAgentId {
        ScImmutableAgentId { obj_id: self.obj_id, key_id: Key32(index) }
    }

    // number of items in array
    pub fn length(&self) -> i32 {
        get_length(self.obj_id)
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// value proxy for immutable bytes array in host container
pub struct ScImmutableBytes {
    obj_id: i32,
    key_id: Key32,
}

impl ScImmutableBytes {
    // check if value exists in host container
    pub fn exists(&self) -> bool {
        exists(self.obj_id, self.key_id, TYPE_BYTES)
    }

    // human-readable string representation
    pub fn to_string(&self) -> String {
        base58_encode(&self.value())
    }

    // get value from host container
    pub fn value(&self) -> Vec<u8> {
        get_bytes(self.obj_id, self.key_id, TYPE_BYTES)
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// array proxy for immutable array of byte array
pub struct ScImmutableBytesArray {
    pub(crate) obj_id: i32
}

impl ScImmutableBytesArray {
    // get value proxy for item at index, index can be 0..length()-1
    pub fn get_bytes(&self, index: i32) -> ScImmutableBytes {
        ScImmutableBytes { obj_id: self.obj_id, key_id: Key32(index) }
    }

    // number of items in array
    pub fn length(&self) -> i32 {
        get_length(self.obj_id)
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// value proxy for immutable ScChainId in host container
pub struct ScImmutableChainId {
    obj_id: i32,
    key_id: Key32,
}

impl ScImmutableChainId {
    // check if value exists in host container
    pub fn exists(&self) -> bool {
        exists(self.obj_id, self.key_id, TYPE_CHAIN_ID)
    }

    // human-readable string representation
    pub fn to_string(&self) -> String {
        self.value().to_string()
    }

    // get value from host container
    pub fn value(&self) -> ScChainId {
        ScChainId::from_bytes(&get_bytes(self.obj_id, self.key_id, TYPE_CHAIN_ID))
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// array proxy for immutable array of ScChainId
pub struct ScImmutableChainIdArray {
    pub(crate) obj_id: i32
}

impl ScImmutableChainIdArray {
    // get value proxy for item at index, index can be 0..length()-1
    pub fn get_chain_id(&self, index: i32) -> ScImmutableChainId {
        ScImmutableChainId { obj_id: self.obj_id, key_id: Key32(index) }
    }

    // number of items in array
    pub fn length(&self) -> i32 {
        get_length(self.obj_id)
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// value proxy for immutable ScColor in host container
pub struct ScImmutableColor {
    obj_id: i32,
    key_id: Key32,
}

impl ScImmutableColor {
    // check if value exists in host container
    pub fn exists(&self) -> bool {
        exists(self.obj_id, self.key_id, TYPE_COLOR)
    }

    // human-readable string representation
    pub fn to_string(&self) -> String {
        self.value().to_string()
    }

    // get value from host container
    pub fn value(&self) -> ScColor {
        ScColor::from_bytes(&get_bytes(self.obj_id, self.key_id, TYPE_COLOR))
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// array proxy for immutable array of ScColor
pub struct ScImmutableColorArray {
    pub(crate) obj_id: i32
}

impl ScImmutableColorArray {
    // get value proxy for item at index, index can be 0..length()-1
    pub fn get_color(&self, index: i32) -> ScImmutableColor {
        ScImmutableColor { obj_id: self.obj_id, key_id: Key32(index) }
    }

    // number of items in array
    pub fn length(&self) -> i32 {
        get_length(self.obj_id)
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// value proxy for immutable ScContractId in host container
pub struct ScImmutableContractId {
    obj_id: i32,
    key_id: Key32,
}

impl ScImmutableContractId {
    // check if value exists in host container
    pub fn exists(&self) -> bool {
        exists(self.obj_id, self.key_id, TYPE_CONTRACT_ID)
    }

    // human-readable string representation
    pub fn to_string(&self) -> String {
        self.value().to_string()
    }

    // get value from host container
    pub fn value(&self) -> ScContractId {
        ScContractId::from_bytes(&get_bytes(self.obj_id, self.key_id, TYPE_CONTRACT_ID))
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// array proxy for immutable array of ScContractId
pub struct ScImmutableContractIdArray {
    pub(crate) obj_id: i32
}

impl ScImmutableContractIdArray {
    // get value proxy for item at index, index can be 0..length()-1
    pub fn get_contract_id(&self, index: i32) -> ScImmutableContractId {
        ScImmutableContractId { obj_id: self.obj_id, key_id: Key32(index) }
    }

    // number of items in array
    pub fn length(&self) -> i32 {
        get_length(self.obj_id)
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// value proxy for immutable ScHash in host container
pub struct ScImmutableHash {
    obj_id: i32,
    key_id: Key32,
}

impl ScImmutableHash {
    // check if value exists in host container
    pub fn exists(&self) -> bool {
        exists(self.obj_id, self.key_id, TYPE_HASH)
    }

    // human-readable string representation
    pub fn to_string(&self) -> String {
        self.value().to_string()
    }

    // get value from host container
    pub fn value(&self) -> ScHash {
        ScHash::from_bytes(&get_bytes(self.obj_id, self.key_id, TYPE_HASH))
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// array proxy for immutable array of ScHash
pub struct ScImmutableHashArray {
    pub(crate) obj_id: i32
}

impl ScImmutableHashArray {
    // get value proxy for item at index, index can be 0..length()-1
    pub fn get_hash(&self, index: i32) -> ScImmutableHash {
        ScImmutableHash { obj_id: self.obj_id, key_id: Key32(index) }
    }

    // number of items in array
    pub fn length(&self) -> i32 {
        get_length(self.obj_id)
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// value proxy for immutable ScHname in host container
pub struct ScImmutableHname {
    obj_id: i32,
    key_id: Key32,
}

impl ScImmutableHname {
    // check if value exists in host container
    pub fn exists(&self) -> bool {
        exists(self.obj_id, self.key_id, TYPE_HNAME)
    }

    // human-readable string representation
    pub fn to_string(&self) -> String {
        self.value().to_string()
    }

    // get value from host container
    pub fn value(&self) -> ScHname {
        ScHname::from_bytes(&get_bytes(self.obj_id, self.key_id, TYPE_HNAME))
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// array proxy for immutable array of ScHname
pub struct ScImmutableHnameArray {
    pub(crate) obj_id: i32
}

impl ScImmutableHnameArray {
    // get value proxy for item at index, index can be 0..length()-1
    pub fn get_hname(&self, index: i32) -> ScImmutableHname {
        ScImmutableHname { obj_id: self.obj_id, key_id: Key32(index) }
    }

    // number of items in array
    pub fn length(&self) -> i32 {
        get_length(self.obj_id)
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// value proxy for immutable int64 in host container
pub struct ScImmutableInt64 {
    obj_id: i32,
    key_id: Key32,
}

impl ScImmutableInt64 {
    // check if value exists in host container
    pub fn exists(&self) -> bool {
        exists(self.obj_id, self.key_id, TYPE_INT64)
    }

    // human-readable string representation
    pub fn to_string(&self) -> String {
        self.value().to_string()
    }

    // get value from host container
    pub fn value(&self) -> i64 {
        let bytes = get_bytes(self.obj_id, self.key_id, TYPE_INT64);
        i64::from_le_bytes(bytes.try_into().expect("invalid i64 length"))
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// array proxy for immutable array of int64
pub struct ScImmutableInt64Array {
    pub(crate) obj_id: i32
}

impl ScImmutableInt64Array {
    // get value proxy for item at index, index can be 0..length()-1
    pub fn get_int64(&self, index: i32) -> ScImmutableInt64 {
        ScImmutableInt64 { obj_id: self.obj_id, key_id: Key32(index) }
    }

    // number of items in array
    pub fn length(&self) -> i32 {
        get_length(self.obj_id)
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// map proxy for immutable map
pub struct ScImmutableMap {
    pub(crate) obj_id: i32
}

impl ScImmutableMap {
    // get value proxy for immutable ScAddress field specified by key
    pub fn get_address<T: MapKey + ?Sized>(&self, key: &T) -> ScImmutableAddress {
        ScImmutableAddress { obj_id: self.obj_id, key_id: key.get_key_id() }
    }

    // get array proxy for ScImmutableAddressArray specified by key
    pub fn get_address_array<T: MapKey + ?Sized>(&self, key: &T) -> ScImmutableAddressArray {
        let arr_id = get_object_id(self.obj_id, key.get_key_id(), TYPE_ADDRESS | TYPE_ARRAY);
        ScImmutableAddressArray { obj_id: arr_id }
    }

    // get value proxy for immutable ScAgentId field specified by key
    pub fn get_agent_id<T: MapKey + ?Sized>(&self, key: &T) -> ScImmutableAgentId {
        ScImmutableAgentId { obj_id: self.obj_id, key_id: key.get_key_id() }
    }

    // get array proxy for ScImmutableAgentIdArray specified by key
    pub fn get_agent_id_array<T: MapKey + ?Sized>(&self, key: &T) -> ScImmutableAgentIdArray {
        let arr_id = get_object_id(self.obj_id, key.get_key_id(), TYPE_AGENT_ID | TYPE_ARRAY);
        ScImmutableAgentIdArray { obj_id: arr_id }
    }

    // get value proxy for immutable bytes array field specified by key
    pub fn get_bytes<T: MapKey + ?Sized>(&self, key: &T) -> ScImmutableBytes {
        ScImmutableBytes { obj_id: self.obj_id, key_id: key.get_key_id() }
    }

    // get array proxy for ScImmutableBytesArray specified by key
    pub fn get_bytes_array<T: MapKey + ?Sized>(&self, key: &T) -> ScImmutableBytesArray {
        let arr_id = get_object_id(self.obj_id, key.get_key_id(), TYPE_BYTES | TYPE_ARRAY);
        ScImmutableBytesArray { obj_id: arr_id }
    }

    // get value proxy for immutable ScChainId field specified by key
    pub fn get_chain_id<T: MapKey + ?Sized>(&self, key: &T) -> ScImmutableChainId {
        ScImmutableChainId { obj_id: self.obj_id, key_id: key.get_key_id() }
    }

    // get array proxy for ScImmutableChainIdArray specified by key
    pub fn get_chain_id_array<T: MapKey + ?Sized>(&self, key: &T) -> ScImmutableChainIdArray {
        let arr_id = get_object_id(self.obj_id, key.get_key_id(), TYPE_CHAIN_ID | TYPE_ARRAY);
        ScImmutableChainIdArray { obj_id: arr_id }
    }

    // get value proxy for immutable ScColor field specified by key
    pub fn get_color<T: MapKey + ?Sized>(&self, key: &T) -> ScImmutableColor {
        ScImmutableColor { obj_id: self.obj_id, key_id: key.get_key_id() }
    }

    // get array proxy for ScImmutableColorArray specified by key
    pub fn get_color_array<T: MapKey + ?Sized>(&self, key: &T) -> ScImmutableColorArray {
        let arr_id = get_object_id(self.obj_id, key.get_key_id(), TYPE_COLOR | TYPE_ARRAY);
        ScImmutableColorArray { obj_id: arr_id }
    }

    // get value proxy for immutable ScContractId field specified by key
    pub fn get_contract_id<T: MapKey + ?Sized>(&self, key: &T) -> ScImmutableContractId {
        ScImmutableContractId { obj_id: self.obj_id, key_id: key.get_key_id() }
    }

    // get array proxy for ScImmutableContractIdArray specified by key
    pub fn get_contract_id_array<T: MapKey + ?Sized>(&self, key: &T) -> ScImmutableContractIdArray {
        let arr_id = get_object_id(self.obj_id, key.get_key_id(), TYPE_CONTRACT_ID | TYPE_ARRAY);
        ScImmutableContractIdArray { obj_id: arr_id }
    }

    // get value proxy for immutable ScHash field specified by key
    pub fn get_hash<T: MapKey + ?Sized>(&self, key: &T) -> ScImmutableHash {
        ScImmutableHash { obj_id: self.obj_id, key_id: key.get_key_id() }
    }

    // get array proxy for ScImmutableHashArray specified by key
    pub fn get_hash_array<T: MapKey + ?Sized>(&self, key: &T) -> ScImmutableHashArray {
        let arr_id = get_object_id(self.obj_id, key.get_key_id(), TYPE_HASH | TYPE_ARRAY);
        ScImmutableHashArray { obj_id: arr_id }
    }

    // get value proxy for immutable ScHname field specified by key
    pub fn get_hname<T: MapKey + ?Sized>(&self, key: &T) -> ScImmutableHname {
        ScImmutableHname { obj_id: self.obj_id, key_id: key.get_key_id() }
    }

    // get array proxy for ScImmutableHnameArray specified by key
    pub fn get_hname_array<T: MapKey + ?Sized>(&self, key: &T) -> ScImmutableHnameArray {
        let arr_id = get_object_id(self.obj_id, key.get_key_id(), TYPE_HNAME | TYPE_ARRAY);
        ScImmutableHnameArray { obj_id: arr_id }
    }

    // get value proxy for immutable int64 field specified by key
    pub fn get_int64<T: MapKey + ?Sized>(&self, key: &T) -> ScImmutableInt64 {
        ScImmutableInt64 { obj_id: self.obj_id, key_id: key.get_key_id() }
    }

    // get array proxy for ScImmutableInt64Array specified by key
    pub fn get_int64_array<T: MapKey + ?Sized>(&self, key: &T) -> ScImmutableInt64Array {
        let arr_id = get_object_id(self.obj_id, key.get_key_id(), TYPE_INT64 | TYPE_ARRAY);
        ScImmutableInt64Array { obj_id: arr_id }
    }

    // get map proxy for ScImmutableMap specified by key
    pub fn get_map<T: MapKey + ?Sized>(&self, key: &T) -> ScImmutableMap {
        let map_id = get_object_id(self.obj_id, key.get_key_id(), TYPE_MAP);
        ScImmutableMap { obj_id: map_id }
    }

    // get array proxy for ScImmutableMapArray specified by key
    pub fn get_map_array<T: MapKey + ?Sized>(&self, key: &T) -> ScImmutableMapArray {
        let arr_id = get_object_id(self.obj_id, key.get_key_id(), TYPE_MAP | TYPE_ARRAY);
        ScImmutableMapArray { obj_id: arr_id }
    }

    // get value proxy for immutable ScRequestId field specified by key
    pub fn get_request_id<T: MapKey + ?Sized>(&self, key: &T) -> ScImmutableRequestId {
        ScImmutableRequestId { obj_id: self.obj_id, key_id: key.get_key_id() }
    }

    // get array proxy for ScImmutableRequestIdArray specified by key
    pub fn get_request_id_array<T: MapKey + ?Sized>(&self, key: &T) -> ScImmutableRequestIdArray {
        let arr_id = get_object_id(self.obj_id, key.get_key_id(), TYPE_REQUEST_ID | TYPE_ARRAY);
        ScImmutableRequestIdArray { obj_id: arr_id }
    }

    // get value proxy for immutable UTF-8 text string field specified by key
    pub fn get_string<T: MapKey + ?Sized>(&self, key: &T) -> ScImmutableString {
        ScImmutableString { obj_id: self.obj_id, key_id: key.get_key_id() }
    }

    // get array proxy for ScImmutableStringArray specified by key
    pub fn get_string_array<T: MapKey + ?Sized>(&self, key: &T) -> ScImmutableStringArray {
        let arr_id = get_object_id(self.obj_id, key.get_key_id(), TYPE_STRING | TYPE_ARRAY);
        ScImmutableStringArray { obj_id: arr_id }
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// array proxy for immutable array of maps
pub struct ScImmutableMapArray {
    pub(crate) obj_id: i32
}

impl ScImmutableMapArray {
    // get value proxy for item at index, index can be 0..length()-1
    pub fn get_map(&self, index: i32) -> ScImmutableMap {
        let map_id = get_object_id(self.obj_id, Key32(index), TYPE_MAP);
        ScImmutableMap { obj_id: map_id }
    }

    // number of items in array
    pub fn length(&self) -> i32 {
        get_length(self.obj_id)
    }
}

// value proxy for immutable ScRequestId in host container
pub struct ScImmutableRequestId {
    obj_id: i32,
    key_id: Key32,
}

impl ScImmutableRequestId {
    // check if value exists in host container
    pub fn exists(&self) -> bool {
        exists(self.obj_id, self.key_id, TYPE_REQUEST_ID)
    }

    // human-readable string representation
    pub fn to_string(&self) -> String {
        self.value().to_string()
    }

    // get value from host container
    pub fn value(&self) -> ScRequestId {
        ScRequestId::from_bytes(&get_bytes(self.obj_id, self.key_id, TYPE_REQUEST_ID))
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// array proxy for immutable array of ScRequestId
pub struct ScImmutableRequestIdArray {
    pub(crate) obj_id: i32
}

impl ScImmutableRequestIdArray {
    // get value proxy for item at index, index can be 0..length()-1
    pub fn get_request_id(&self, index: i32) -> ScImmutableRequestId {
        ScImmutableRequestId { obj_id: self.obj_id, key_id: Key32(index) }
    }

    // number of items in array
    pub fn length(&self) -> i32 {
        get_length(self.obj_id)
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// value proxy for immutable UTF-8 text string in host container
pub struct ScImmutableString {
    obj_id: i32,
    key_id: Key32,
}

impl ScImmutableString {
    // check if value exists in host container
    pub fn exists(&self) -> bool {
        exists(self.obj_id, self.key_id, TYPE_STRING)
    }

    // human-readable string representation
    pub fn to_string(&self) -> String {
        self.value()
    }

    // get value from host container
    pub fn value(&self) -> String {
        let bytes = get_bytes(self.obj_id, self.key_id, TYPE_STRING);
        unsafe { String::from_utf8_unchecked(bytes) }
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// array proxy for immutable array of UTF-8 text string
pub struct ScImmutableStringArray {
    pub(crate) obj_id: i32
}

impl ScImmutableStringArray {
    // get value proxy for item at index, index can be 0..length()-1
    pub fn get_string(&self, index: i32) -> ScImmutableString {
        ScImmutableString { obj_id: self.obj_id, key_id: Key32(index) }
    }

    // number of items in array
    pub fn length(&self) -> i32 {
        get_length(self.obj_id)
    }
}
