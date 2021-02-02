// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

#![allow(dead_code)]

use wasplib::client::*;

use super::*;

pub const SC_NAME: &str = "donatewithfeedback";
pub const SC_HNAME: Hname = Hname(0x696d7f66);

pub const PARAM_FEEDBACK: &str = "feedback";
pub const PARAM_WITHDRAW_AMOUNT: &str = "withdraw";

pub const VAR_AMOUNT: &str = "amount";
pub const VAR_DONATIONS: &str = "donations";
pub const VAR_DONATOR: &str = "donator";
pub const VAR_ERROR: &str = "error";
pub const VAR_FEEDBACK: &str = "feedback";
pub const VAR_LOG: &str = "log";
pub const VAR_MAX_DONATION: &str = "max_donation";
pub const VAR_TIMESTAMP: &str = "timestamp";
pub const VAR_TOTAL_DONATION: &str = "total_donation";

pub const FUNC_DONATE: &str = "donate";
pub const FUNC_WITHDRAW: &str = "withdraw";
pub const VIEW_DONATIONS: &str = "view_donations";

pub const HFUNC_DONATE: Hname = Hname(0xdc9b133a);
pub const HFUNC_WITHDRAW: Hname = Hname(0x9dcc0f41);
pub const HVIEW_DONATIONS: Hname = Hname(0xc3cc7cb0);

#[no_mangle]
fn on_load() {
    let exports = ScExports::new();
    exports.add_call(FUNC_DONATE, func_donate);
    exports.add_call(FUNC_WITHDRAW, func_withdraw);
    exports.add_view(VIEW_DONATIONS, view_donations);
}
