// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use wasmlib::*;

use crate::*;

pub fn func_divide(ctx: &ScFuncContext, _params: &FuncDivideParams) {
    let amount = ctx.balances().balance(&ScColor::IOTA);
    if amount == 0 {
        ctx.panic("nothing to divide");
    }
    let state = ctx.state();
    let total = state.get_int64(VAR_TOTAL_FACTOR).value();
    let members = state.get_map(VAR_MEMBERS);
    let member_list = state.get_address_array(VAR_MEMBER_LIST);
    let size = member_list.length();
    let mut parts = 0_i64;
    for i in 0..size {
        let address = member_list.get_address(i).value();
        let factor = members.get_int64(&address).value();
        let share = amount * factor / total;
        if share != 0 {
            parts += share;
            let transfers = ScTransfers::new(&ScColor::IOTA, share);
            ctx.transfer_to_address(&address, transfers);
        }
    }
    if parts != amount {
        // note we truncated the calculations down to the nearest integer
        // there could be some small remainder left in the contract, but
        // that will be picked up in the next round as part of the balance
        let remainder = amount - parts;
        ctx.log(&("remainder in contract: ".to_string() + &remainder.to_string()));
    }
}

pub fn func_member(ctx: &ScFuncContext, params: &FuncMemberParams) {
    let state = ctx.state();
    let members = state.get_map(VAR_MEMBERS);
    let address = params.address.value();
    let current_factor = members.get_int64(&address);
    if !current_factor.exists() {
        // add new address to member list
        let member_list = state.get_address_array(VAR_MEMBER_LIST);
        member_list.get_address(member_list.length()).set_value(&address);
    }
    let factor = params.factor.value();
    let total_factor = state.get_int64(VAR_TOTAL_FACTOR);
    let new_total_factor = total_factor.value() - current_factor.value() + factor;
    total_factor.set_value(new_total_factor);
    current_factor.set_value(factor);
}

pub fn view_get_factor(ctx: &ScViewContext, params: &ViewGetFactorParams) {
    let address = params.address.value();
    let members = ctx.state().get_map(VAR_MEMBERS);
    let factor = members.get_int64(&address).value();
    ctx.results().get_int64(VAR_FACTOR).set_value(factor);
}
