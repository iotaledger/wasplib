// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

//@formatter:off

#![allow(dead_code)]

use wasmlib::*;

use crate::*;

pub const IDX_PARAM_ADDRESS:          usize = 0;
pub const IDX_PARAM_AGENT_ID:         usize = 1;
pub const IDX_PARAM_CALLER:           usize = 2;
pub const IDX_PARAM_CHAIN_ID:         usize = 3;
pub const IDX_PARAM_CHAIN_OWNER_ID:   usize = 4;
pub const IDX_PARAM_CONTRACT_CREATOR: usize = 5;
pub const IDX_PARAM_COUNTER:          usize = 6;
pub const IDX_PARAM_HASH:             usize = 7;
pub const IDX_PARAM_HNAME:            usize = 8;
pub const IDX_PARAM_HNAME_CONTRACT:   usize = 9;
pub const IDX_PARAM_HNAME_EP:         usize = 10;
pub const IDX_PARAM_HNAME_ZERO:       usize = 11;
pub const IDX_PARAM_INT64:            usize = 12;
pub const IDX_PARAM_INT64_ZERO:       usize = 13;
pub const IDX_PARAM_INT_VALUE:        usize = 14;
pub const IDX_PARAM_NAME:             usize = 15;
pub const IDX_PARAM_STRING:           usize = 16;
pub const IDX_PARAM_STRING_ZERO:      usize = 17;
pub const IDX_RESULT_CHAIN_OWNER_ID:  usize = 18;
pub const IDX_RESULT_COUNTER:         usize = 19;
pub const IDX_RESULT_INT_VALUE:       usize = 20;
pub const IDX_RESULT_MINTED_COLOR:    usize = 21;
pub const IDX_RESULT_MINTED_SUPPLY:   usize = 22;
pub const IDX_RESULT_SANDBOX_CALL:    usize = 23;
pub const IDX_STATE_COUNTER:          usize = 24;
pub const IDX_STATE_HNAME_EP:         usize = 25;
pub const IDX_STATE_INTS:             usize = 26;
pub const IDX_STATE_MINTED_COLOR:     usize = 27;
pub const IDX_STATE_MINTED_SUPPLY:    usize = 28;

pub const KEY_MAP_LEN: usize = 29;

pub const KEY_MAP: [&str; KEY_MAP_LEN] = [
    PARAM_ADDRESS,
    PARAM_AGENT_ID,
    PARAM_CALLER,
    PARAM_CHAIN_ID,
    PARAM_CHAIN_OWNER_ID,
    PARAM_CONTRACT_CREATOR,
    PARAM_COUNTER,
    PARAM_HASH,
    PARAM_HNAME,
    PARAM_HNAME_CONTRACT,
    PARAM_HNAME_EP,
    PARAM_HNAME_ZERO,
    PARAM_INT64,
    PARAM_INT64_ZERO,
    PARAM_INT_VALUE,
    PARAM_NAME,
    PARAM_STRING,
    PARAM_STRING_ZERO,
    RESULT_CHAIN_OWNER_ID,
    RESULT_COUNTER,
    RESULT_INT_VALUE,
    RESULT_MINTED_COLOR,
    RESULT_MINTED_SUPPLY,
    RESULT_SANDBOX_CALL,
    STATE_COUNTER,
    STATE_HNAME_EP,
    STATE_INTS,
    STATE_MINTED_COLOR,
    STATE_MINTED_SUPPLY,
];

pub static mut IDX_MAP: [Key32; KEY_MAP_LEN] = [Key32(0); KEY_MAP_LEN];

pub fn idx_map(idx: usize) -> Key32 {
    unsafe {
        IDX_MAP[idx]
    }
}

//@formatter:on
