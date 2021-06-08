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

pub struct HelloWorldFuncState {
    pub(crate) state_id: i32,
}

impl HelloWorldFuncState {
    pub fn dummy(&self) -> ScMutableString {
        ScMutableString::new(self.state_id, idx_map(IDX_VAR_DUMMY))
    }
}

pub struct HelloWorldViewState {
    pub(crate) state_id: i32,
}

impl HelloWorldViewState {
    pub fn dummy(&self) -> ScImmutableString {
        ScImmutableString::new(self.state_id, idx_map(IDX_VAR_DUMMY))
    }
}