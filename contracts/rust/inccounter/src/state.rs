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

pub struct MutableIncCounterState {
    pub(crate) id: i32,
}

impl MutableIncCounterState {
    pub fn counter(&self) -> ScMutableInt64 {
        ScMutableInt64::new(self.id, idx_map(IDX_STATE_COUNTER))
    }

    pub fn num_repeats(&self) -> ScMutableInt64 {
        ScMutableInt64::new(self.id, idx_map(IDX_STATE_NUM_REPEATS))
    }
}

pub struct ImmutableIncCounterState {
    pub(crate) id: i32,
}

impl ImmutableIncCounterState {
    pub fn counter(&self) -> ScImmutableInt64 {
        ScImmutableInt64::new(self.id, idx_map(IDX_STATE_COUNTER))
    }

    pub fn num_repeats(&self) -> ScImmutableInt64 {
        ScImmutableInt64::new(self.id, idx_map(IDX_STATE_NUM_REPEATS))
    }
}
