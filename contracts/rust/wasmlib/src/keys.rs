// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use crate::host::*;

pub trait MapKey {
    fn get_id(&self) -> Key32;
}

impl MapKey for str {
    fn get_id(&self) -> Key32 {
        get_key_id_from_string(self)
    }
}

impl MapKey for String {
    fn get_id(&self) -> Key32 {
        get_key_id_from_string(self)
    }
}

#[derive(Clone, Copy)]
pub struct Key32(pub i32);

impl MapKey for Key32 {
    fn get_id(&self) -> Key32 {
        *self
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// @formatter:off
pub const KEY_ADDRESS          : Key32 = Key32(-1);
pub const KEY_BALANCES         : Key32 = Key32(-2);
pub const KEY_BASE58_BYTES     : Key32 = Key32(-3);
pub const KEY_BASE58_STRING    : Key32 = Key32(-4);
pub const KEY_BLS_ADDRESS      : Key32 = Key32(-5);
pub const KEY_BLS_AGGREGATE    : Key32 = Key32(-6);
pub const KEY_BLS_VALID        : Key32 = Key32(-7);
pub const KEY_CALL             : Key32 = Key32(-8);
pub const KEY_CALLER           : Key32 = Key32(-9);
pub const KEY_CHAIN_OWNER_ID   : Key32 = Key32(-10);
pub const KEY_COLOR            : Key32 = Key32(-11);
pub const KEY_CONTRACT_CREATOR : Key32 = Key32(-12);
pub const KEY_CONTRACT_ID      : Key32 = Key32(-13);
pub const KEY_DEPLOY           : Key32 = Key32(-14);
pub const KEY_ED25519_ADDRESS  : Key32 = Key32(-15);
pub const KEY_ED25519_VALID    : Key32 = Key32(-16);
pub const KEY_EVENT            : Key32 = Key32(-17);
pub const KEY_EXPORTS          : Key32 = Key32(-18);
pub const KEY_HASH_BLAKE2B     : Key32 = Key32(-19);
pub const KEY_HASH_SHA3        : Key32 = Key32(-20);
pub const KEY_HNAME            : Key32 = Key32(-21);
pub const KEY_INCOMING         : Key32 = Key32(-22);
pub const KEY_LENGTH           : Key32 = Key32(-23);
pub const KEY_LOG              : Key32 = Key32(-24);
pub const KEY_MAPS             : Key32 = Key32(-25);
pub const KEY_MINTED           : Key32 = Key32(-26);
pub const KEY_NAME             : Key32 = Key32(-27);
pub const KEY_PANIC            : Key32 = Key32(-28);
pub const KEY_PARAMS           : Key32 = Key32(-29);
pub const KEY_POST             : Key32 = Key32(-30);
pub const KEY_RANDOM           : Key32 = Key32(-31);
pub const KEY_REQUEST_ID       : Key32 = Key32(-32);
pub const KEY_RESULTS          : Key32 = Key32(-33);
pub const KEY_RETURN           : Key32 = Key32(-34);
pub const KEY_STATE            : Key32 = Key32(-35);
pub const KEY_TIMESTAMP        : Key32 = Key32(-36);
pub const KEY_TRACE            : Key32 = Key32(-37);
pub const KEY_TRANSFERS        : Key32 = Key32(-38);
pub const KEY_UTILITY          : Key32 = Key32(-39);
pub const KEY_VALID            : Key32 = Key32(-40);
pub const KEY_ZZZZZZZ          : Key32 = Key32(-41);
// @formatter:on
