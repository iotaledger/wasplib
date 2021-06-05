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

pub struct IncCounterFuncState {
    pub(crate) state_id: i32,
}

impl IncCounterFuncState {
    pub fn counter(&self) -> ScMutableInt64 {
        ScMutableInt64::new(self.state_id, VAR_COUNTER.get_key_id())
    }

    pub fn num_repeats(&self) -> ScMutableInt64 {
        ScMutableInt64::new(self.state_id, VAR_NUM_REPEATS.get_key_id())
    }
}

pub struct IncCounterViewState {
    pub(crate) state_id: i32,
}

impl IncCounterViewState {
    pub fn counter(&self) -> ScImmutableInt64 {
        ScImmutableInt64::new(self.state_id, VAR_COUNTER.get_key_id())
    }

    pub fn num_repeats(&self) -> ScImmutableInt64 {
        ScImmutableInt64::new(self.state_id, VAR_NUM_REPEATS.get_key_id())
    }
}
