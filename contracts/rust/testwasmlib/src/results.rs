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

#[derive(Clone, Copy)]
pub struct ImmutableBlockRecordResults {
    pub(crate) id: i32,
}

impl ImmutableBlockRecordResults {
    pub fn record(&self) -> ScImmutableBytes {
        ScImmutableBytes::new(self.id, idx_map(IDX_RESULT_RECORD))
    }
}

#[derive(Clone, Copy)]
pub struct MutableBlockRecordResults {
    pub(crate) id: i32,
}

impl MutableBlockRecordResults {
    pub fn record(&self) -> ScMutableBytes {
        ScMutableBytes::new(self.id, idx_map(IDX_RESULT_RECORD))
    }
}

#[derive(Clone, Copy)]
pub struct ImmutableBlockRecordsResults {
    pub(crate) id: i32,
}

impl ImmutableBlockRecordsResults {
    pub fn count(&self) -> ScImmutableInt32 {
        ScImmutableInt32::new(self.id, idx_map(IDX_RESULT_COUNT))
    }
}

#[derive(Clone, Copy)]
pub struct MutableBlockRecordsResults {
    pub(crate) id: i32,
}

impl MutableBlockRecordsResults {
    pub fn count(&self) -> ScMutableInt32 {
        ScMutableInt32::new(self.id, idx_map(IDX_RESULT_COUNT))
    }
}
