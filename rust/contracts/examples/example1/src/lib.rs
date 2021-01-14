// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use wasplib::client::*;

#[no_mangle]
fn on_load() {
    // declare entry points of the smart contract
    let exports = ScExports::new();
    exports.add_call("storeString", store_string);
    exports.add_view("getString", get_string);
}

// storeString entry point
fn store_string(ctx: &ScCallContext) {
    // take parameter paramString
    let par = ctx.params().get_string("paramString");
    if !par.exists() {
        ctx.panic("string parameter not found") // panic if parameter does not exist
    }
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

