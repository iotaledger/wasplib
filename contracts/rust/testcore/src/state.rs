// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

#![allow(dead_code)]
#![allow(unused_imports)]

use wasmlib::*;
use wasmlib::host::*;

use crate::*;
use crate::keys::*;

pub struct MutableTestCoreState {
    pub(crate) id: i32,
}

impl MutableTestCoreState {
    pub fn counter(&self) -> ScMutableInt64 {
        ScMutableInt64::new(self.id, idx_map(IDX_STATE_COUNTER))
    }

    pub fn hname_ep(&self) -> ScMutableHname {
        ScMutableHname::new(self.id, idx_map(IDX_STATE_HNAME_EP))
    }

    pub fn minted_color(&self) -> ScMutableColor {
        ScMutableColor::new(self.id, idx_map(IDX_STATE_MINTED_COLOR))
    }

    pub fn minted_supply(&self) -> ScMutableInt64 {
        ScMutableInt64::new(self.id, idx_map(IDX_STATE_MINTED_SUPPLY))
    }
}

pub struct ImmutableTestCoreState {
    pub(crate) id: i32,
}

impl ImmutableTestCoreState {
    pub fn counter(&self) -> ScImmutableInt64 {
        ScImmutableInt64::new(self.id, idx_map(IDX_STATE_COUNTER))
    }

    pub fn hname_ep(&self) -> ScImmutableHname {
        ScImmutableHname::new(self.id, idx_map(IDX_STATE_HNAME_EP))
    }

    pub fn minted_color(&self) -> ScImmutableColor {
        ScImmutableColor::new(self.id, idx_map(IDX_STATE_MINTED_COLOR))
    }

    pub fn minted_supply(&self) -> ScImmutableInt64 {
        ScImmutableInt64::new(self.id, idx_map(IDX_STATE_MINTED_SUPPLY))
    }
}
