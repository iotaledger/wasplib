// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use super::host::get_key_id_from_string;

// @formatter:off
pub const KEY_AGENT       : i32 =                 -1;
pub const KEY_AMOUNT      : i32 = KEY_AGENT       -1;
pub const KEY_BALANCES    : i32 = KEY_AMOUNT      -1;
pub const KEY_BASE58      : i32 = KEY_BALANCES    -1;
pub const KEY_CALLER      : i32 = KEY_BASE58      -1;
pub const KEY_CALLS       : i32 = KEY_CALLER      -1;
pub const KEY_CHAIN       : i32 = KEY_CALLS       -1;
pub const KEY_CHAIN_OWNER : i32 = KEY_CHAIN       -1;
pub const KEY_COLOR       : i32 = KEY_CHAIN_OWNER -1;
pub const KEY_CONTRACT    : i32 = KEY_COLOR       -1;
pub const KEY_CREATOR     : i32 = KEY_CONTRACT    -1;
pub const KEY_DATA        : i32 = KEY_CREATOR     -1;
pub const KEY_DELAY       : i32 = KEY_DATA        -1;
pub const KEY_DESCRIPTION : i32 = KEY_DELAY       -1;
pub const KEY_ERROR       : i32 = KEY_DESCRIPTION -1;
pub const KEY_EVENT       : i32 = KEY_ERROR       -1;
pub const KEY_EXPORTS     : i32 = KEY_EVENT       -1;
pub const KEY_FUNCTION    : i32 = KEY_EXPORTS     -1;
pub const KEY_HASH        : i32 = KEY_FUNCTION    -1;
pub const KEY_ID          : i32 = KEY_HASH        -1;
pub const KEY_INCOMING    : i32 = KEY_ID          -1;
pub const KEY_IOTA        : i32 = KEY_INCOMING    -1;
pub const KEY_LENGTH      : i32 = KEY_IOTA        -1;
pub const KEY_LOG         : i32 = KEY_LENGTH      -1;
pub const KEY_LOGS        : i32 = KEY_LOG         -1;
pub const KEY_NAME        : i32 = KEY_LOGS        -1;
pub const KEY_PANIC       : i32 = KEY_NAME        -1;
pub const KEY_PARAMS      : i32 = KEY_PANIC       -1;
pub const KEY_POSTS       : i32 = KEY_PARAMS      -1;
pub const KEY_RANDOM      : i32 = KEY_POSTS       -1;
pub const KEY_RESULTS     : i32 = KEY_RANDOM      -1;
pub const KEY_STATE       : i32 = KEY_RESULTS     -1;
pub const KEY_TIMESTAMP   : i32 = KEY_STATE       -1;
pub const KEY_TRACE       : i32 = KEY_TIMESTAMP   -1;
pub const KEY_TRANSFERS   : i32 = KEY_TRACE       -1;
pub const KEY_UTILITY     : i32 = KEY_TRANSFERS   -1;
pub const KEY_VIEWS       : i32 = KEY_UTILITY     -1;
pub const KEY_WARNING     : i32 = KEY_VIEWS       -1;
pub const KEY_ZZZZZZZ     : i32 = KEY_WARNING     -1;
// @formatter:on

pub trait MapKey {
    fn get_id(&self) -> i32;
}

impl MapKey for str {
    fn get_id(&self) -> i32 {
        get_key_id_from_string(self)
    }
}

impl MapKey for i32 {
    fn get_id(&self) -> i32 {
        *self
    }
}
