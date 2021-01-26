// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// types encapsulating mutable host objects

use super::context::*;
use super::hashtypes::*;
use super::host::*;
use super::immutable::*;
use super::keys::*;

pub(crate) static ROOT: ScMutableMap = ScMutableMap { obj_id: 1 };

pub struct ScMutableAddress {
    obj_id: i32,
    key_id: Key32,
}

impl ScMutableAddress {
    pub(crate) fn new(obj_id: i32, key_id: Key32) -> ScMutableAddress {
        ScMutableAddress { obj_id, key_id }
    }

    pub fn exists(&self) -> bool {
        exists(self.obj_id, self.key_id)
    }

    pub fn set_value(&self, val: &ScAddress) {
        set_bytes(self.obj_id, self.key_id, val.to_bytes());
    }

    pub fn to_string(&self) -> String {
        self.value().to_string()
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
        set_clear(self.obj_id);
    }

    //TODO exists on arrays?

    // index 0..length(), when length() a new one is appended
    pub fn get_address(&self, index: i32) -> ScMutableAddress {
        ScMutableAddress { obj_id: self.obj_id, key_id: Key32(index) }
    }

    pub fn immutable(&self) -> ScImmutableAddressArray {
        ScImmutableAddressArray::new(self.obj_id)
    }

    pub fn length(&self) -> i32 {
        get_length(self.obj_id)
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScMutableAgent {
    obj_id: i32,
    key_id: Key32,
}

impl ScMutableAgent {
    pub(crate) fn new(obj_id: i32, key_id: Key32) -> ScMutableAgent {
        ScMutableAgent { obj_id, key_id }
    }

    pub fn exists(&self) -> bool {
        exists(self.obj_id, self.key_id)
    }

    pub fn set_value(&self, val: &ScAgent) {
        set_bytes(self.obj_id, self.key_id, val.to_bytes());
    }

    pub fn to_string(&self) -> String {
        self.value().to_string()
    }

    pub fn value(&self) -> ScAgent {
        ScAgent::from_bytes(&get_bytes(self.obj_id, self.key_id))
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScMutableAgentArray {
    obj_id: i32
}

impl ScMutableAgentArray {
    pub(crate) fn new(obj_id: i32) -> ScMutableAgentArray {
        ScMutableAgentArray { obj_id }
    }

    pub fn clear(&self) {
        set_clear(self.obj_id);
    }

    //TODO exists on arrays?

    // index 0..length(), when length() a new one is appended
    pub fn get_agent(&self, index: i32) -> ScMutableAgent {
        ScMutableAgent { obj_id: self.obj_id, key_id: Key32(index) }
    }

    pub fn immutable(&self) -> ScImmutableAgentArray {
        ScImmutableAgentArray::new(self.obj_id)
    }

    pub fn length(&self) -> i32 {
        get_length(self.obj_id)
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScMutableBytes {
    obj_id: i32,
    key_id: Key32,
}

impl ScMutableBytes {
    pub(crate) fn new(obj_id: i32, key_id: Key32) -> ScMutableBytes {
        ScMutableBytes { obj_id, key_id }
    }

    pub fn exists(&self) -> bool {
        exists(self.obj_id, self.key_id)
    }

    pub fn set_value(&self, val: &[u8]) {
        set_bytes(self.obj_id, self.key_id, val);
    }

    pub fn to_string(&self) -> String {
        base58_encode(&self.value())
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
        set_clear(self.obj_id);
    }

    // index 0..length(), when length() a new one is appended
    pub fn get_bytes(&self, index: i32) -> ScMutableBytes {
        ScMutableBytes { obj_id: self.obj_id, key_id: Key32(index) }
    }

    pub fn immutable(&self) -> ScImmutableBytesArray {
        ScImmutableBytesArray::new(self.obj_id)
    }

    pub fn length(&self) -> i32 {
        get_length(self.obj_id)
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScMutableChainId {
    obj_id: i32,
    key_id: Key32,
}

impl ScMutableChainId {
    pub(crate) fn new(obj_id: i32, key_id: Key32) -> ScMutableChainId {
        ScMutableChainId { obj_id, key_id }
    }

    pub fn exists(&self) -> bool {
        exists(self.obj_id, self.key_id)
    }

    pub fn set_value(&self, val: &ScChainId) {
        set_bytes(self.obj_id, self.key_id, val.to_bytes());
    }

    pub fn to_string(&self) -> String {
        self.value().to_string()
    }

    pub fn value(&self) -> ScChainId {
        ScChainId::from_bytes(&get_bytes(self.obj_id, self.key_id))
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScMutableColor {
    obj_id: i32,
    key_id: Key32,
}

impl ScMutableColor {
    pub(crate) fn new(obj_id: i32, key_id: Key32) -> ScMutableColor {
        ScMutableColor { obj_id, key_id }
    }

    pub fn exists(&self) -> bool {
        exists(self.obj_id, self.key_id)
    }

    pub fn set_value(&self, val: &ScColor) {
        set_bytes(self.obj_id, self.key_id, val.to_bytes());
    }

    pub fn to_string(&self) -> String {
        self.value().to_string()
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
        set_clear(self.obj_id);
    }

    // index 0..length(), when length() a new one is appended
    pub fn get_color(&self, index: i32) -> ScMutableColor {
        ScMutableColor { obj_id: self.obj_id, key_id: Key32(index) }
    }

    pub fn immutable(&self) -> ScImmutableColorArray {
        ScImmutableColorArray::new(self.obj_id)
    }

    pub fn length(&self) -> i32 {
        get_length(self.obj_id)
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScMutableContractId {
    obj_id: i32,
    key_id: Key32,
}

impl ScMutableContractId {
    pub(crate) fn new(obj_id: i32, key_id: Key32) -> ScMutableContractId {
        ScMutableContractId { obj_id, key_id }
    }

    pub fn exists(&self) -> bool {
        exists(self.obj_id, self.key_id)
    }

    pub fn set_value(&self, val: &ScContractId) {
        set_bytes(self.obj_id, self.key_id, val.to_bytes());
    }

    pub fn to_string(&self) -> String {
        self.value().to_string()
    }

    pub fn value(&self) -> ScContractId {
        ScContractId::from_bytes(&get_bytes(self.obj_id, self.key_id))
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScMutableHash {
    obj_id: i32,
    key_id: Key32,
}

impl ScMutableHash {
    pub(crate) fn new(obj_id: i32, key_id: Key32) -> ScMutableHash {
        ScMutableHash { obj_id, key_id }
    }

    pub fn exists(&self) -> bool {
        exists(self.obj_id, self.key_id)
    }

    pub fn set_value(&self, val: &ScHash) {
        set_bytes(self.obj_id, self.key_id, val.to_bytes());
    }

    pub fn to_string(&self) -> String {
        self.value().to_string()
    }

    pub fn value(&self) -> ScHash {
        ScHash::from_bytes(&get_bytes(self.obj_id, self.key_id))
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScMutableHashArray {
    obj_id: i32
}

impl ScMutableHashArray {
    pub(crate) fn new(obj_id: i32) -> ScMutableHashArray {
        ScMutableHashArray { obj_id }
    }

    pub fn clear(&self) {
        set_clear(self.obj_id);
    }

    // index 0..length(), when length() a new one is appended
    pub fn get_hash(&self, index: i32) -> ScMutableHash {
        ScMutableHash { obj_id: self.obj_id, key_id: Key32(index) }
    }

    pub fn immutable(&self) -> ScImmutableHashArray {
        ScImmutableHashArray::new(self.obj_id)
    }

    pub fn length(&self) -> i32 {
        get_length(self.obj_id)
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScMutableHname {
    obj_id: i32,
    key_id: Key32,
}

impl ScMutableHname {
    pub(crate) fn new(obj_id: i32, key_id: Key32) -> ScMutableHname {
        ScMutableHname { obj_id, key_id }
    }

    pub fn exists(&self) -> bool {
        exists(self.obj_id, self.key_id)
    }

    pub fn set_value(&self, val: Hname) {
        set_bytes(self.obj_id, self.key_id, &val.to_bytes());
    }

    pub fn to_string(&self) -> String {
        self.value().to_string()
    }

    pub fn value(&self) -> Hname {
        Hname::from_bytes(&get_bytes(self.obj_id, self.key_id))
    }
}


// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScMutableInt {
    obj_id: i32,
    key_id: Key32,
}

impl ScMutableInt {
    pub(crate) fn new(obj_id: i32, key_id: Key32) -> ScMutableInt {
        ScMutableInt { obj_id, key_id }
    }

    pub fn exists(&self) -> bool {
        exists(self.obj_id, self.key_id)
    }

    pub fn set_value(&self, val: i64) {
        set_int(self.obj_id, self.key_id, val);
    }

    pub fn to_string(&self) -> String {
        self.value().to_string()
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
        set_clear(self.obj_id);
    }

    // index 0..length(), when length() a new one is appended
    pub fn get_int(&self, index: i32) -> ScMutableInt {
        ScMutableInt { obj_id: self.obj_id, key_id: Key32(index) }
    }

    pub fn immutable(&self) -> ScImmutableIntArray {
        ScImmutableIntArray::new(self.obj_id)
    }

    pub fn length(&self) -> i32 {
        get_length(self.obj_id)
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScMutableMap {
    pub(crate) obj_id: i32
}

impl ScMutableMap {
    pub const NONE: ScMutableMap = ScMutableMap { obj_id: 0 };

    pub fn new() -> ScMutableMap {
        let maps = ROOT.get_map_array(&KEY_MAPS);
        maps.get_map(maps.length())
    }

    pub fn clear(&self) {
        set_clear(self.obj_id);
    }

    pub fn get_address<T: MapKey + ?Sized>(&self, key: &T) -> ScMutableAddress {
        ScMutableAddress { obj_id: self.obj_id, key_id: key.get_id() }
    }

    pub fn get_address_array<T: MapKey + ?Sized>(&self, key: &T) -> ScMutableAddressArray {
        let arr_id = get_object_id(self.obj_id, key.get_id(), TYPE_ADDRESS | TYPE_ARRAY);
        ScMutableAddressArray { obj_id: arr_id }
    }

    pub fn get_agent<T: MapKey + ?Sized>(&self, key: &T) -> ScMutableAgent {
        ScMutableAgent { obj_id: self.obj_id, key_id: key.get_id() }
    }

    pub fn get_agent_array<T: MapKey + ?Sized>(&self, key: &T) -> ScMutableAgentArray {
        let arr_id = get_object_id(self.obj_id, key.get_id(), TYPE_AGENT | TYPE_ARRAY);
        ScMutableAgentArray { obj_id: arr_id }
    }

    pub fn get_bytes<T: MapKey + ?Sized>(&self, key: &T) -> ScMutableBytes {
        ScMutableBytes { obj_id: self.obj_id, key_id: key.get_id() }
    }

    pub fn get_bytes_array<T: MapKey + ?Sized>(&self, key: &T) -> ScMutableBytesArray {
        let arr_id = get_object_id(self.obj_id, key.get_id(), TYPE_BYTES | TYPE_ARRAY);
        ScMutableBytesArray { obj_id: arr_id }
    }

    pub fn get_chain_id<T: MapKey + ?Sized>(&self, key: &T) -> ScMutableChainId {
        ScMutableChainId { obj_id: self.obj_id, key_id: key.get_id() }
    }

    pub fn get_color<T: MapKey + ?Sized>(&self, key: &T) -> ScMutableColor {
        ScMutableColor { obj_id: self.obj_id, key_id: key.get_id() }
    }

    pub fn get_color_array<T: MapKey + ?Sized>(&self, key: &T) -> ScMutableColorArray {
        let arr_id = get_object_id(self.obj_id, key.get_id(), TYPE_COLOR | TYPE_ARRAY);
        ScMutableColorArray { obj_id: arr_id }
    }

    pub fn get_contract_id<T: MapKey + ?Sized>(&self, key: &T) -> ScMutableContractId {
        ScMutableContractId { obj_id: self.obj_id, key_id: key.get_id() }
    }

    pub fn get_hash<T: MapKey + ?Sized>(&self, key: &T) -> ScMutableHash {
        ScMutableHash { obj_id: self.obj_id, key_id: key.get_id() }
    }

    pub fn get_hash_array<T: MapKey + ?Sized>(&self, key: &T) -> ScMutableHashArray {
        let arr_id = get_object_id(self.obj_id, key.get_id(), TYPE_HASH | TYPE_ARRAY);
        ScMutableHashArray { obj_id: arr_id }
    }

    pub fn get_hname<T: MapKey + ?Sized>(&self, key: &T) -> ScMutableHname {
        ScMutableHname { obj_id: self.obj_id, key_id: key.get_id() }
    }

    pub fn get_int<T: MapKey + ?Sized>(&self, key: &T) -> ScMutableInt {
        ScMutableInt { obj_id: self.obj_id, key_id: key.get_id() }
    }

    pub fn get_int_array<T: MapKey + ?Sized>(&self, key: &T) -> ScMutableIntArray {
        let arr_id = get_object_id(self.obj_id, key.get_id(), TYPE_INT | TYPE_ARRAY);
        ScMutableIntArray { obj_id: arr_id }
    }

    pub fn get_map<T: MapKey + ?Sized>(&self, key: &T) -> ScMutableMap {
        let map_id = get_object_id(self.obj_id, key.get_id(), TYPE_MAP);
        ScMutableMap { obj_id: map_id }
    }

    pub fn get_map_array<T: MapKey + ?Sized>(&self, key: &T) -> ScMutableMapArray {
        let arr_id = get_object_id(self.obj_id, key.get_id(), TYPE_MAP | TYPE_ARRAY);
        ScMutableMapArray { obj_id: arr_id }
    }

    pub fn get_string<T: MapKey + ?Sized>(&self, key: &T) -> ScMutableString {
        ScMutableString { obj_id: self.obj_id, key_id: key.get_id() }
    }

    pub fn get_string_array<T: MapKey + ?Sized>(&self, key: &T) -> ScMutableStringArray {
        let arr_id = get_object_id(self.obj_id, key.get_id(), TYPE_STRING | TYPE_ARRAY);
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
        set_clear(self.obj_id);
    }


    // index 0..length(), inclusive, hen length() a new one is appended
    pub fn get_map(&self, index: i32) -> ScMutableMap {
        let map_id = get_object_id(self.obj_id, Key32(index), TYPE_MAP);
        ScMutableMap { obj_id: map_id }
    }

    pub fn immutable(&self) -> ScImmutableMapArray {
        ScImmutableMapArray::new(self.obj_id)
    }

    pub fn length(&self) -> i32 {
        get_length(self.obj_id)
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScMutableString {
    obj_id: i32,
    key_id: Key32,
}

impl ScMutableString {
    pub(crate) fn new(obj_id: i32, key_id: Key32) -> ScMutableString {
        ScMutableString { obj_id, key_id }
    }

    pub fn exists(&self) -> bool {
        exists(self.obj_id, self.key_id)
    }

    pub fn set_value(&self, val: &str) {
        set_string(self.obj_id, self.key_id, val);
    }

    pub fn to_string(&self) -> String {
        self.value()
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
        set_clear(self.obj_id);
    }

    // index 0..length(), when length() a new one is appended
    pub fn get_string(&self, index: i32) -> ScMutableString {
        ScMutableString { obj_id: self.obj_id, key_id: Key32(index) }
    }

    pub fn immutable(&self) -> ScImmutableStringArray {
        ScImmutableStringArray::new(self.obj_id)
    }

    pub fn length(&self) -> i32 {
        get_length(self.obj_id)
    }
}
