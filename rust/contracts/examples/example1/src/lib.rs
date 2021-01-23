// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use wasplib::client::*;

#[no_mangle]
fn on_load() {
    // declare entry points of the smart contract
    let exports = ScExports::new();
    exports.add_call("storeString", store_string);
    exports.add_view("getString", get_string);
    exports.add_call("withdraw_iota", withdraw_iota);
}

// storeString entry point
fn store_string(ctx: &ScCallContext) {
    // take parameter paramString
    let par = ctx.params().get_string("paramString");
    ctx.require(par.exists(), "string parameter not found");

    // store the string in "storedString" variable
    ctx.state().get_string("storedString").set_value(&par.value());
    // log the text
    let msg = "Message stored: ".to_string() + &par.value();
    ctx.log(&msg);
}

// getString view
fn get_string(ctx: &ScViewContext) {
    // take the stored string
    let s = ctx.state().get_string("storedString").value();
    // return the string value in the result dictionary
    ctx.results().get_string("paramString").set_value(&s);
}

fn withdraw_iota(ctx: &ScCallContext) {
    let creator = ctx.contract_creator();
    let caller = ctx.caller();

    ctx.require(creator.equals(&caller), "not authorised");
    ctx.require(caller.is_address(), "caller must be an address");

    let bal = ctx.balances().balance(&ScColor::IOTA);
    if bal > 0 {
        ctx.transfer_to_address(&caller.address(), &ScTransfers::new(&ScColor::IOTA, bal))
    }
}
