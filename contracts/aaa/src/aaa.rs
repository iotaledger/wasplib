// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use wasmlib::*;

use crate::*;

//use crate::types::*;

pub fn func_donate(ctx: &ScCallContext, _params: &FuncDonateParams) {
    ctx.log("calling donate");
}

pub fn func_withdraw(ctx: &ScCallContext, _params: &FuncWithdrawParams) {
    ctx.log("calling withdraw");
}

pub fn view_donations(ctx: &ScViewContext, _params: &ViewDonationsParams) {
    ctx.log("calling donations");
}
