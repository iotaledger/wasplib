// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

// @formatter:off

#![allow(dead_code)]

use wasmlib::*;

use crate::*;

pub const IDX_PARAM_COUNTER:     usize = 0;
pub const IDX_PARAM_DUMMY:       usize = 1;
pub const IDX_PARAM_NUM_REPEATS: usize = 2;
pub const IDX_RESULT_COUNTER:    usize = 3;
pub const IDX_STATE_COUNTER:     usize = 4;
pub const IDX_STATE_NUM_REPEATS: usize = 5;

pub const KEY_MAP_LEN: usize = 6;

pub const KEY_MAP: [&str; KEY_MAP_LEN] = [
    PARAM_COUNTER,
    PARAM_DUMMY,
    PARAM_NUM_REPEATS,
    RESULT_COUNTER,
    STATE_COUNTER,
    STATE_NUM_REPEATS,
];

pub static mut IDX_MAP: [Key32; KEY_MAP_LEN] = [Key32(0); KEY_MAP_LEN];

pub fn idx_map(idx: usize) -> Key32 {
    unsafe {
        IDX_MAP[idx]
    }
}

// @formatter:on
