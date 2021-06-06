// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

//@formatter:off

#![allow(dead_code)]

use wasmlib::*;
use wasmlib::host::*;

use crate::types::*;

pub type ImmutableBidderList = ArrayOfImmutableAgentId;

pub struct ArrayOfImmutableAgentId {
    pub(crate) obj_id: i32,
}

impl ArrayOfImmutableAgentId {
    pub fn length(&self) -> i32 {
        get_length(self.obj_id)
    }

    pub fn get_agent_id(&self, index: i32) -> ScImmutableAgentId {
        ScImmutableAgentId::new(self.obj_id, Key32(index))
    }
}

pub type MutableBidderList = ArrayOfMutableAgentId;

pub struct ArrayOfMutableAgentId {
    pub(crate) obj_id: i32,
}

impl ArrayOfMutableAgentId {
    pub fn clear(&self) {
        clear(self.obj_id);
    }

    pub fn length(&self) -> i32 {
        get_length(self.obj_id)
    }

    pub fn get_agent_id(&self, index: i32) -> ScMutableAgentId {
        ScMutableAgentId::new(self.obj_id, Key32(index))
    }
}

pub type ImmutableBids = MapAgentIdToImmutableBid;

pub struct MapAgentIdToImmutableBid {
    pub(crate) obj_id: i32,
}

impl MapAgentIdToImmutableBid {
    pub fn get_bid(&self, key: &ScAgentId) -> ImmutableBid {
        ImmutableBid { obj_id: self.obj_id, key_id: key.get_key_id() }
    }
}

pub type MutableBids = MapAgentIdToMutableBid;

pub struct MapAgentIdToMutableBid {
    pub(crate) obj_id: i32,
}

impl MapAgentIdToMutableBid {
    pub fn clear(&self) {
        clear(self.obj_id)
    }

    pub fn get_bid(&self, key: &ScAgentId) -> MutableBid {
        MutableBid { obj_id: self.obj_id, key_id: key.get_key_id() }
    }
}

//@formatter:on
