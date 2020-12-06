// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use super::host::get_key_id;

static mut ID_LENGTH: i32 = 0;
static mut ID_LOG: i32 = 0;
static mut ID_TRACE: i32 = 0;

pub trait KeyId {
    fn get_id(&self) -> i32;
}

impl KeyId for str {
    fn get_id(&self) -> i32 {
        get_key_id(self)
    }
}

pub fn key_length() -> i32 {
    unsafe {
        if ID_LENGTH == 0 {
            ID_LENGTH = get_key_id("length");
        }
        ID_LENGTH
    }
}

pub fn key_log() -> i32 {
    unsafe {
        if ID_LOG == 0 {
            ID_LOG = get_key_id("log");
        }
        ID_LOG
    }
}

pub fn key_trace() -> i32 {
    unsafe {
        if ID_TRACE == 0 {
            ID_TRACE = get_key_id("trace");
        }
        ID_TRACE
    }
}
