// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

// @formatter:off

#![allow(dead_code)]

use wasmlib::*;

use crate::*;

pub const IDX_PARAM_ADDRESS:      usize = 0;
pub const IDX_PARAM_AGENT_ID:     usize = 1;
pub const IDX_PARAM_BLOCK_INDEX:  usize = 2;
pub const IDX_PARAM_BYTES:        usize = 3;
pub const IDX_PARAM_CHAIN_ID:     usize = 4;
pub const IDX_PARAM_COLOR:        usize = 5;
pub const IDX_PARAM_HASH:         usize = 6;
pub const IDX_PARAM_HNAME:        usize = 7;
pub const IDX_PARAM_INDEX:        usize = 8;
pub const IDX_PARAM_INT16:        usize = 9;
pub const IDX_PARAM_INT32:        usize = 10;
pub const IDX_PARAM_INT64:        usize = 11;
pub const IDX_PARAM_NAME:         usize = 12;
pub const IDX_PARAM_RECORD_INDEX: usize = 13;
pub const IDX_PARAM_REQUEST_ID:   usize = 14;
pub const IDX_PARAM_STRING:       usize = 15;
pub const IDX_PARAM_VALUE:        usize = 16;
pub const IDX_RESULT_COUNT:       usize = 17;
pub const IDX_RESULT_LENGTH:      usize = 18;
pub const IDX_RESULT_RECORD:      usize = 19;
pub const IDX_RESULT_VALUE:       usize = 20;
pub const IDX_STATE_ARRAYS:       usize = 21;

pub const KEY_MAP_LEN: usize = 22;

pub const KEY_MAP: [&str; KEY_MAP_LEN] = [
    PARAM_ADDRESS,
    PARAM_AGENT_ID,
    PARAM_BLOCK_INDEX,
    PARAM_BYTES,
    PARAM_CHAIN_ID,
    PARAM_COLOR,
    PARAM_HASH,
    PARAM_HNAME,
    PARAM_INDEX,
    PARAM_INT16,
    PARAM_INT32,
    PARAM_INT64,
    PARAM_NAME,
    PARAM_RECORD_INDEX,
    PARAM_REQUEST_ID,
    PARAM_STRING,
    PARAM_VALUE,
    RESULT_COUNT,
    RESULT_LENGTH,
    RESULT_RECORD,
    RESULT_VALUE,
    STATE_ARRAYS,
];

pub static mut IDX_MAP: [Key32; KEY_MAP_LEN] = [Key32(0); KEY_MAP_LEN];

pub fn idx_map(idx: usize) -> Key32 {
    unsafe {
        IDX_MAP[idx]
    }
}

// @formatter:on
