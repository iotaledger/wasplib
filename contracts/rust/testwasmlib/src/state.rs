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

pub struct TestWasmLibFuncState {
    pub(crate) state_id: i32,
}

impl TestWasmLibFuncState {
    pub fn dummy(&self) -> ScMutableAgentId {
        ScMutableAgentId::new(self.state_id, idx_map(IDX_VAR_DUMMY))
    }
}

pub struct TestWasmLibViewState {
    pub(crate) state_id: i32,
}

impl TestWasmLibViewState {
    pub fn dummy(&self) -> ScImmutableAgentId {
        ScImmutableAgentId::new(self.state_id, idx_map(IDX_VAR_DUMMY))
    }
}