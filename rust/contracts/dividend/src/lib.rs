// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

#![allow(dead_code)]
#![allow(non_snake_case)]

use wasplib::client::*;

struct Member {
    address: ScAddress,
    factor: i64,
}

#[no_mangle]
pub fn onLoad() {
    let exports = ScExports::new();
    exports.add_call("member", member);
    exports.add_call("dividend", dividend);
}

fn member(sc: &ScCallContext) {
    let request = sc.request();
    if !request.from(&sc.contract().owner()) {
        sc.log("Cancel spoofed request");
        return;
    }
    let params = request.params();
    let address = params.get_address("address");
    if !address.exists() {
        sc.log("Missing address");
        return;
    }
    let factor = params.get_int("factor");
    if !factor.exists() {
        sc.log("Missing factor");
        return;
    }
    let member = Member {
        address: address.value(),
        factor: factor.value(),
    };
    let state = sc.state();
    let totalFactor = state.get_int("totalFactor");
    let mut total = totalFactor.value();
    let members = state.get_bytes_array("members");
    let size = members.length();
    for i in 0..size {
        let bytes = members.get_bytes(i).value();
        let m = decodeMember(&bytes);
        if m.address == member.address {
            total -= m.factor;
            total += member.factor;
            totalFactor.set_value(total);
            let bytes = encodeMember(&member);
            members.get_bytes(i).set_value(&bytes);
            sc.log(&("Updated: ".to_string() + &member.address.to_string()));
            return;
        }
    }
    total += member.factor;
    totalFactor.set_value(total);
    let bytes = encodeMember(&member);
    members.get_bytes(size).set_value(&bytes);
    sc.log(&("Appended: ".to_string() + &member.address.to_string()));
}

fn dividend(sc: &ScCallContext) {
    let amount = sc.account().balance(&ScColor::IOTA);
    if amount == 0 {
        sc.log("Nothing to divide");
        return;
    }
    let state = sc.state();
    let totalFactor = state.get_int("totalFactor");
    let total = totalFactor.value();
    let members = state.get_bytes_array("members");
    let size = members.length();
    let mut parts = 0_i64;
    for i in 0..size {
        let bytes = members.get_bytes(i).value();
        let m = decodeMember(&bytes);
        let part = amount * m.factor / total;
        if part != 0 {
            parts += part;
            sc.transfer(&m.address.as_agent(), &ScColor::IOTA, part);
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

fn decodeMember(bytes: &[u8]) -> Member {
    let mut decoder = BytesDecoder::new(bytes);
    Member {
        address: decoder.address(),
        factor: decoder.int(),
    }
}

fn encodeMember(member: &Member) -> Vec<u8> {
    let mut encoder = BytesEncoder::new();
    encoder.address(&member.address);
    encoder.int(member.factor);
    encoder.data()
}
