// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

#![allow(dead_code)]
#![allow(unused_imports)]

use wasmlib::*;
use wasmlib::host::*;

use crate::*;
use crate::keys::*;

pub struct ArrayOfImmutableAddress {
    pub(crate) obj_id: i32,
}

impl ArrayOfImmutableAddress {
    pub fn length(&self) -> i32 {
        get_length(self.obj_id)
    }

    pub fn get_address(&self, index: i32) -> ScImmutableAddress {
        ScImmutableAddress::new(self.obj_id, Key32(index))
    }
}

pub struct MapAddressToImmutableInt64 {
    pub(crate) obj_id: i32,
}

impl MapAddressToImmutableInt64 {
    pub fn get_int64(&self, key: &ScAddress) -> ScImmutableInt64 {
        ScImmutableInt64::new(self.obj_id, key.get_key_id())
    }
}

#[derive(Clone, Copy)]
pub struct ImmutableDividendState {
    pub(crate) id: i32,
}

impl ImmutableDividendState {
    pub fn factor(&self) -> ScImmutableInt64 {
        ScImmutableInt64::new(self.id, idx_map(IDX_STATE_FACTOR))
    }

    pub fn member_list(&self) -> ArrayOfImmutableAddress {
        let arr_id = get_object_id(self.id, idx_map(IDX_STATE_MEMBER_LIST), TYPE_ARRAY | TYPE_ADDRESS);
        ArrayOfImmutableAddress { obj_id: arr_id }
    }

    pub fn members(&self) -> MapAddressToImmutableInt64 {
        let map_id = get_object_id(self.id, idx_map(IDX_STATE_MEMBERS), TYPE_MAP);
        MapAddressToImmutableInt64 { obj_id: map_id }
    }

    pub fn owner(&self) -> ScImmutableAgentID {
        ScImmutableAgentID::new(self.id, idx_map(IDX_STATE_OWNER))
    }

    pub fn total_factor(&self) -> ScImmutableInt64 {
        ScImmutableInt64::new(self.id, idx_map(IDX_STATE_TOTAL_FACTOR))
    }
}

pub struct ArrayOfMutableAddress {
    pub(crate) obj_id: i32,
}

impl ArrayOfMutableAddress {
    pub fn clear(&self) {
        clear(self.obj_id);
    }

    pub fn length(&self) -> i32 {
        get_length(self.obj_id)
    }

    pub fn get_address(&self, index: i32) -> ScMutableAddress {
        ScMutableAddress::new(self.obj_id, Key32(index))
    }
}

pub struct MapAddressToMutableInt64 {
    pub(crate) obj_id: i32,
}

impl MapAddressToMutableInt64 {
    pub fn clear(&self) {
        clear(self.obj_id)
    }

    pub fn get_int64(&self, key: &ScAddress) -> ScMutableInt64 {
        ScMutableInt64::new(self.obj_id, key.get_key_id())
    }
}

#[derive(Clone, Copy)]
pub struct MutableDividendState {
    pub(crate) id: i32,
}

impl MutableDividendState {
    pub fn factor(&self) -> ScMutableInt64 {
        ScMutableInt64::new(self.id, idx_map(IDX_STATE_FACTOR))
    }

    pub fn member_list(&self) -> ArrayOfMutableAddress {
        let arr_id = get_object_id(self.id, idx_map(IDX_STATE_MEMBER_LIST), TYPE_ARRAY | TYPE_ADDRESS);
        ArrayOfMutableAddress { obj_id: arr_id }
    }

    pub fn members(&self) -> MapAddressToMutableInt64 {
        let map_id = get_object_id(self.id, idx_map(IDX_STATE_MEMBERS), TYPE_MAP);
        MapAddressToMutableInt64 { obj_id: map_id }
    }

    pub fn owner(&self) -> ScMutableAgentID {
        ScMutableAgentID::new(self.id, idx_map(IDX_STATE_OWNER))
    }

    pub fn total_factor(&self) -> ScMutableInt64 {
        ScMutableInt64::new(self.id, idx_map(IDX_STATE_TOTAL_FACTOR))
    }
}
