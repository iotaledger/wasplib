// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

#![allow(dead_code)]

use wasplib::client::*;
use super::*;

pub const SC_NAME: &str = "tokenregistry";
pub const SC_HNAME: ScHname = ScHname(0xe1ba0c78);

pub const PARAM_DESCRIPTION: &str = "description";
pub const PARAM_USER_DEFINED: &str = "user_defined";

pub const VAR_COLOR_LIST: &str = "color_list";
pub const VAR_REGISTRY: &str = "registry";

pub const FUNC_MINT_SUPPLY: &str = "mint_supply";
pub const FUNC_TRANSFER_OWNERSHIP: &str = "transfer_ownership";
pub const FUNC_UPDATE_METADATA: &str = "update_metadata";

pub const HFUNC_MINT_SUPPLY: ScHname = ScHname(0x5b0b99b9);
pub const HFUNC_TRANSFER_OWNERSHIP: ScHname = ScHname(0xea337e10);
pub const HFUNC_UPDATE_METADATA: ScHname = ScHname(0xaee46d94);

#[no_mangle]
fn on_load() {
    let exports = ScExports::new();
    exports.add_call(FUNC_MINT_SUPPLY, func_mint_supply);
    exports.add_call(FUNC_TRANSFER_OWNERSHIP, func_transfer_ownership);
    exports.add_call(FUNC_UPDATE_METADATA, func_update_metadata);
}
