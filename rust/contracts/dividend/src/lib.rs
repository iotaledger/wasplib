// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use types::*;
use wasplib::client::*;

mod types;

const PARAM_ADDRESS: &str = "address";
const PARAM_FACTOR: &str = "factor";

const VAR_MEMBERS: &str = "members";
const VAR_TOTAL_FACTOR: &str = "total_factor";

#[no_mangle]
fn on_load() {
    let exports = ScExports::new();
    exports.add_call("member", member);
    exports.add_call("divide", divide);
}

fn member(ctx: &ScCallContext) {
    if !ctx.from(&ctx.contract_creator()) {
        ctx.panic("Cancel spoofed request");
    }
    let params = ctx.params();
    let address = params.get_address(PARAM_ADDRESS);
    if !address.exists() {
        ctx.panic("Missing address");
    }
    let factor = params.get_int(PARAM_FACTOR);
    if !factor.exists() {
        ctx.panic("Missing factor");
    }
    let member = Member {
        address: address.value(),
        factor: factor.value(),
    };
    let state = ctx.state();
    let total_factor = state.get_int(VAR_TOTAL_FACTOR);
    let mut total = total_factor.value();
    let members = state.get_bytes_array(VAR_MEMBERS);
    let size = members.length();
    for i in 0..size {
        let m = decode_member(&members.get_bytes(i).value());
        if m.address.equals(&member.address) {
            total -= m.factor;
            total += member.factor;
            total_factor.set_value(total);
            members.get_bytes(i).set_value(&encode_member(&member));
            ctx.log(&("Updated: ".to_string() + &member.address.to_string()));
            return;
        }
    }
    total += member.factor;
    total_factor.set_value(total);
    members.get_bytes(size).set_value(&encode_member(&member));
    ctx.log(&("Appended: ".to_string() + &member.address.to_string()));
}

fn divide(ctx: &ScCallContext) {
    let amount = ctx.balances().balance(&ScColor::IOTA);
    if amount == 0 {
        ctx.panic("Nothing to divide");
    }
    let state = ctx.state();
    let total_factor = state.get_int(VAR_TOTAL_FACTOR);
    let total = total_factor.value();
    let members = state.get_bytes_array(VAR_MEMBERS);
    let mut parts = 0_i64;
    let size = members.length();
    for i in 0..size {
        let m = decode_member(&members.get_bytes(i).value());
        let part = amount * m.factor / total;
        if part != 0 {
            parts += part;
            ctx.transfer_to_address(&m.address, &ScTransfers::new(&ScColor::IOTA, part));
        }
    }
    if parts != amount {
        // note we truncated the calculations down to the nearest integer
        // there could be some small remainder left in the contract, but
        // that will be picked up in the next round as part of the balance
        let remainder = amount - parts;
        ctx.log(&("Remainder in contract: ".to_string() + &remainder.to_string()));
    }
}
