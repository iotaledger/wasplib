// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

#![allow(dead_code)]

use wasplib::client::*;

use super::*;

pub const SC_NAME: &str = "dividend";
pub const SC_HNAME: Hname = Hname(0xcce2e239);

pub const PARAM_ADDRESS: &str = "address";
pub const PARAM_FACTOR: &str = "factor";

pub const VAR_MEMBERS: &str = "members";
pub const VAR_TOTAL_FACTOR: &str = "total_factor";

pub const FUNC_DIVIDE: &str = "divide";
pub const FUNC_MEMBER: &str = "member";

pub const HFUNC_DIVIDE: Hname = Hname(0xc7878107);
pub const HFUNC_MEMBER: Hname = Hname(0xc07da2cb);

#[no_mangle]
fn on_load() {
    let exports = ScExports::new();
    exports.add_call(FUNC_DIVIDE, func_divide);
    exports.add_call(FUNC_MEMBER, func_member);
}
