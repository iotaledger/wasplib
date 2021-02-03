// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// types encapsulating immutable host objects

use std::convert::TryInto;

use super::context::*;
use super::hashtypes::*;
use super::host::*;
use super::keys::*;

pub struct ScImmutableAddress {
    obj_id: i32,
    key_id: Key32,
}

impl ScImmutableAddress {
    pub(crate) fn new(obj_id: i32, key_id: Key32) -> ScImmutableAddress {
        ScImmutableAddress { obj_id, key_id }
    }

    pub fn exists(&self) -> bool {
        exists(self.obj_id, self.key_id, TYPE_ADDRESS)
    }

    pub fn to_string(&self) -> String {
        self.value().to_string()
    }

    pub fn value(&self) -> ScAddress {
        ScAddress::from_bytes(&get_bytes(self.obj_id, self.key_id, TYPE_ADDRESS))
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

    // index 0..length(), exclusive
    pub fn get_address(&self, index: i32) -> ScImmutableAddress {
        ScImmutableAddress { obj_id: self.obj_id, key_id: Key32(index) }
    }

    pub fn length(&self) -> i32 {
        get_length(self.obj_id)
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScImmutableAgentId {
    obj_id: i32,
    key_id: Key32,
}

impl ScImmutableAgentId {
    pub(crate) fn new(obj_id: i32, key_id: Key32) -> ScImmutableAgentId {
        ScImmutableAgentId { obj_id, key_id }
    }

    pub fn exists(&self) -> bool {
        exists(self.obj_id, self.key_id, TYPE_AGENT_ID)
    }

    pub fn to_string(&self) -> String {
        self.value().to_string()
    }

    pub fn value(&self) -> ScAgentId {
        ScAgentId::from_bytes(&get_bytes(self.obj_id, self.key_id, TYPE_AGENT_ID))
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScImmutableAgentIdArray {
    obj_id: i32
}

impl ScImmutableAgentIdArray {
    pub(crate) fn new(obj_id: i32) -> ScImmutableAgentIdArray {
        ScImmutableAgentIdArray { obj_id }
    }

    // index 0..length(), exclusive
    pub fn get_agent_id(&self, index: i32) -> ScImmutableAgentId {
        ScImmutableAgentId { obj_id: self.obj_id, key_id: Key32(index) }
    }

    pub fn length(&self) -> i32 {
        get_length(self.obj_id)
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScImmutableBytes {
    obj_id: i32,
    key_id: Key32,
}

impl ScImmutableBytes {
    pub(crate) fn new(obj_id: i32, key_id: Key32) -> ScImmutableBytes {
        ScImmutableBytes { obj_id, key_id }
    }

    pub fn exists(&self) -> bool {
        exists(self.obj_id, self.key_id, TYPE_BYTES)
    }

    pub fn to_string(&self) -> String {
        base58_encode(&self.value())
    }

    pub fn value(&self) -> Vec<u8> {
        get_bytes(self.obj_id, self.key_id, TYPE_BYTES)
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
        ScImmutableBytes { obj_id: self.obj_id, key_id: Key32(index) }
    }

    pub fn length(&self) -> i32 {
        get_length(self.obj_id)
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScImmutableChainId {
    obj_id: i32,
    key_id: Key32,
}

impl ScImmutableChainId {
    pub(crate) fn new(obj_id: i32, key_id: Key32) -> ScImmutableChainId {
        ScImmutableChainId { obj_id, key_id }
    }

    pub fn exists(&self) -> bool {
        exists(self.obj_id, self.key_id, TYPE_CHAIN_ID)
    }

    pub fn to_string(&self) -> String {
        self.value().to_string()
    }

    pub fn value(&self) -> ScChainId {
        ScChainId::from_bytes(&get_bytes(self.obj_id, self.key_id, TYPE_CHAIN_ID))
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScImmutableColor {
    obj_id: i32,
    key_id: Key32,
}

impl ScImmutableColor {
    pub(crate) fn new(obj_id: i32, key_id: Key32) -> ScImmutableColor {
        ScImmutableColor { obj_id, key_id }
    }

    pub fn exists(&self) -> bool {
        exists(self.obj_id, self.key_id, TYPE_COLOR)
    }

    pub fn to_string(&self) -> String {
        self.value().to_string()
    }

    pub fn value(&self) -> ScColor {
        ScColor::from_bytes(&get_bytes(self.obj_id, self.key_id, TYPE_COLOR))
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
        ScImmutableColor { obj_id: self.obj_id, key_id: Key32(index) }
    }

    pub fn length(&self) -> i32 {
        get_length(self.obj_id)
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScImmutableContractId {
    obj_id: i32,
    key_id: Key32,
}

impl ScImmutableContractId {
    pub(crate) fn new(obj_id: i32, key_id: Key32) -> ScImmutableContractId {
        ScImmutableContractId { obj_id, key_id }
    }

    pub fn exists(&self) -> bool {
        exists(self.obj_id, self.key_id, TYPE_CONTRACT_ID)
    }

    pub fn to_string(&self) -> String {
        self.value().to_string()
    }

    pub fn value(&self) -> ScContractId {
        ScContractId::from_bytes(&get_bytes(self.obj_id, self.key_id, TYPE_CONTRACT_ID))
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScImmutableHash {
    obj_id: i32,
    key_id: Key32,
}

impl ScImmutableHash {
    pub(crate) fn new(obj_id: i32, key_id: Key32) -> ScImmutableHash {
        ScImmutableHash { obj_id, key_id }
    }

    pub fn exists(&self) -> bool {
        exists(self.obj_id, self.key_id, TYPE_HASH)
    }

    pub fn to_string(&self) -> String {
        self.value().to_string()
    }

    pub fn value(&self) -> ScHash {
        ScHash::from_bytes(&get_bytes(self.obj_id, self.key_id, TYPE_HASH))
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScImmutableHashArray {
    obj_id: i32
}

impl ScImmutableHashArray {
    pub(crate) fn new(obj_id: i32) -> ScImmutableHashArray {
        ScImmutableHashArray { obj_id }
    }

    // index 0..length(), exclusive
    pub fn get_hash(&self, index: i32) -> ScImmutableHash {
        ScImmutableHash { obj_id: self.obj_id, key_id: Key32(index) }
    }

    pub fn length(&self) -> i32 {
        get_length(self.obj_id)
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScImmutableHname {
    obj_id: i32,
    key_id: Key32,
}

impl ScImmutableHname {
    pub(crate) fn new(obj_id: i32, key_id: Key32) -> ScImmutableHname {
        ScImmutableHname { obj_id, key_id }
    }

    pub fn exists(&self) -> bool {
        exists(self.obj_id, self.key_id, TYPE_HNAME)
    }

    pub fn to_string(&self) -> String {
        self.value().to_string()
    }

    pub fn value(&self) -> ScHname {
        ScHname::from_bytes(&get_bytes(self.obj_id, self.key_id, TYPE_HNAME))
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScImmutableInt {
    obj_id: i32,
    key_id: Key32,
}

impl ScImmutableInt {
    pub(crate) fn new(obj_id: i32, key_id: Key32) -> ScImmutableInt {
        ScImmutableInt { obj_id, key_id }
    }

    pub fn exists(&self) -> bool {
        exists(self.obj_id, self.key_id, TYPE_INT)
    }

    pub fn to_string(&self) -> String {
        self.value().to_string()
    }

    pub fn value(&self) -> i64 {
        let bytes = get_bytes(self.obj_id, self.key_id, TYPE_INT);
        i64::from_le_bytes(bytes.try_into().expect("invalid i64 length"))
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
        ScImmutableInt { obj_id: self.obj_id, key_id: Key32(index) }
    }

    pub fn length(&self) -> i32 {
        get_length(self.obj_id)
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScImmutableMap {
    pub(crate) obj_id: i32
}

impl ScImmutableMap {
    pub(crate) const fn new(obj_id: i32) -> ScImmutableMap {
        ScImmutableMap { obj_id }
    }

    pub fn get_address<T: MapKey + ?Sized>(&self, key: &T) -> ScImmutableAddress {
        ScImmutableAddress { obj_id: self.obj_id, key_id: key.get_id() }
    }

    pub fn get_address_array<T: MapKey + ?Sized>(&self, key: &T) -> ScImmutableAddressArray {
        let arr_id = get_object_id(self.obj_id, key.get_id(), TYPE_ADDRESS | TYPE_ARRAY);
        ScImmutableAddressArray { obj_id: arr_id }
    }

    pub fn get_agent_id<T: MapKey + ?Sized>(&self, key: &T) -> ScImmutableAgentId {
        ScImmutableAgentId { obj_id: self.obj_id, key_id: key.get_id() }
    }

    pub fn get_agent_id_array<T: MapKey + ?Sized>(&self, key: &T) -> ScImmutableAgentIdArray {
        let arr_id = get_object_id(self.obj_id, key.get_id(), TYPE_AGENT_ID | TYPE_ARRAY);
        ScImmutableAgentIdArray { obj_id: arr_id }
    }

    pub fn get_bytes<T: MapKey + ?Sized>(&self, key: &T) -> ScImmutableBytes {
        ScImmutableBytes { obj_id: self.obj_id, key_id: key.get_id() }
    }

    pub fn get_bytes_array<T: MapKey + ?Sized>(&self, key: &T) -> ScImmutableBytesArray {
        let arr_id = get_object_id(self.obj_id, key.get_id(), TYPE_BYTES | TYPE_ARRAY);
        ScImmutableBytesArray { obj_id: arr_id }
    }

    pub fn get_chain_id<T: MapKey + ?Sized>(&self, key: &T) -> ScImmutableChainId {
        ScImmutableChainId { obj_id: self.obj_id, key_id: key.get_id() }
    }

    pub fn get_color<T: MapKey + ?Sized>(&self, key: &T) -> ScImmutableColor {
        ScImmutableColor { obj_id: self.obj_id, key_id: key.get_id() }
    }

    pub fn get_color_array<T: MapKey + ?Sized>(&self, key: &T) -> ScImmutableColorArray {
        let arr_id = get_object_id(self.obj_id, key.get_id(), TYPE_COLOR | TYPE_ARRAY);
        ScImmutableColorArray { obj_id: arr_id }
    }

    pub fn get_contract_id<T: MapKey + ?Sized>(&self, key: &T) -> ScImmutableContractId {
        ScImmutableContractId { obj_id: self.obj_id, key_id: key.get_id() }
    }

    pub fn get_hash<T: MapKey + ?Sized>(&self, key: &T) -> ScImmutableHash {
        ScImmutableHash { obj_id: self.obj_id, key_id: key.get_id() }
    }

    pub fn get_hash_array<T: MapKey + ?Sized>(&self, key: &T) -> ScImmutableHashArray {
        let arr_id = get_object_id(self.obj_id, key.get_id(), TYPE_HASH | TYPE_ARRAY);
        ScImmutableHashArray { obj_id: arr_id }
    }

    pub fn get_hname<T: MapKey + ?Sized>(&self, key: &T) -> ScImmutableHname {
        ScImmutableHname { obj_id: self.obj_id, key_id: key.get_id() }
    }

    pub fn get_int<T: MapKey + ?Sized>(&self, key: &T) -> ScImmutableInt {
        ScImmutableInt { obj_id: self.obj_id, key_id: key.get_id() }
    }

    pub fn get_int_array<T: MapKey + ?Sized>(&self, key: &T) -> ScImmutableIntArray {
        let arr_id = get_object_id(self.obj_id, key.get_id(), TYPE_INT | TYPE_ARRAY);
        ScImmutableIntArray { obj_id: arr_id }
    }

    pub fn get_map<T: MapKey + ?Sized>(&self, key: &T) -> ScImmutableMap {
        let map_id = get_object_id(self.obj_id, key.get_id(), TYPE_MAP);
        ScImmutableMap { obj_id: map_id }
    }

    pub fn get_map_array<T: MapKey + ?Sized>(&self, key: &T) -> ScImmutableMapArray {
        let arr_id = get_object_id(self.obj_id, key.get_id(), TYPE_MAP | TYPE_ARRAY);
        ScImmutableMapArray { obj_id: arr_id }
    }

    pub fn get_string<T: MapKey + ?Sized>(&self, key: &T) -> ScImmutableString {
        ScImmutableString { obj_id: self.obj_id, key_id: key.get_id() }
    }

    pub fn get_string_array<T: MapKey + ?Sized>(&self, key: &T) -> ScImmutableStringArray {
        let arr_id = get_object_id(self.obj_id, key.get_id(), TYPE_STRING | TYPE_ARRAY);
        ScImmutableStringArray { obj_id: arr_id }
    }

    pub fn length(&self) -> i32 {
        get_length(self.obj_id)
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
        let map_id = get_object_id(self.obj_id, Key32(index), TYPE_MAP);
        ScImmutableMap { obj_id: map_id }
    }

    pub fn length(&self) -> i32 {
        get_length(self.obj_id)
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScImmutableString {
    obj_id: i32,
    key_id: Key32,
}

impl ScImmutableString {
    pub(crate) fn new(obj_id: i32, key_id: Key32) -> ScImmutableString {
        ScImmutableString { obj_id, key_id }
    }

    pub fn exists(&self) -> bool {
        exists(self.obj_id, self.key_id, TYPE_STRING)
    }

    pub fn to_string(&self) -> String {
        self.value()
    }

    pub fn value(&self) -> String {
        let bytes = get_bytes(self.obj_id, self.key_id, TYPE_STRING);
        unsafe { String::from_utf8_unchecked(bytes) }
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
        ScImmutableString { obj_id: self.obj_id, key_id: Key32(index) }
    }

    pub fn length(&self) -> i32 {
        get_length(self.obj_id)
    }
}
