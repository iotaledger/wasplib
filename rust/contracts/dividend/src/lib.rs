// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use types::*;
use wasplib::client::*;

mod types;

const KEY_ADDRESS: &str = "address";
const KEY_FACTOR: &str = "factor";
const KEY_MEMBERS: &str = "members";
const KEY_TOTAL_FACTOR: &str = "total_factor";

#[no_mangle]
fn on_load() {
    let exports = ScExports::new();
    exports.add_call("member", member);
    exports.add_call("dividend", dividend);
}

fn member(sc: &ScCallContext) {
    if !sc.from(&sc.contract_creator()) {
        sc.panic("Cancel spoofed request");
    }
    let params = sc.params();
    let address = params.get_address(KEY_ADDRESS);
    if !address.exists() {
        sc.panic("Missing address");
    }
    let factor = params.get_int(KEY_FACTOR);
    if !factor.exists() {
        sc.panic("Missing factor");
    }
    let member = Member {
        address: address.value(),
        factor: factor.value(),
    };
    let state = sc.state();
    let total_factor = state.get_int(KEY_TOTAL_FACTOR);
    let mut total = total_factor.value();
    let members = state.get_bytes_array(KEY_MEMBERS);
    let size = members.length();
    for i in 0..size {
        let m = decode_member(&members.get_bytes(i).value());
        if m.address.equals(&member.address) {
            total -= m.factor;
            total += member.factor;
            total_factor.set_value(total);
            members.get_bytes(i).set_value(&encode_member(&member));
            sc.log(&("Updated: ".to_string() + &member.address.to_string()));
            return;
        }
    }
    total += member.factor;
    total_factor.set_value(total);
    members.get_bytes(size).set_value(&encode_member(&member));
    sc.log(&("Appended: ".to_string() + &member.address.to_string()));
}

fn dividend(sc: &ScCallContext) {
    let amount = sc.balances().balance(&ScColor::IOTA);
    if amount == 0 {
        sc.panic("Nothing to divide");
    }
    let state = sc.state();
    let total_factor = state.get_int(KEY_TOTAL_FACTOR);
    let total = total_factor.value();
    let members = state.get_bytes_array(KEY_MEMBERS);
    let mut parts = 0_i64;
    let size = members.length();
    for i in 0..size {
        let m = decode_member(&members.get_bytes(i).value());
        let part = amount * m.factor / total;
        if part != 0 {
            parts += part;
            sc.transfer_to_address(&m.address, &ScColor::IOTA, part);
        }
    }
    if parts != amount {
        // note we truncated the calculations down to the nearest integer
        // there could be some small remainder left in the contract, but
        // that will be picked up in the next round as part of the balance
        let remainder = amount - parts;
        sc.log(&("Remainder in contract: ".to_string() + &remainder.to_string()));
    }
}
