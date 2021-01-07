// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use wasplib::client::*;

#[no_mangle]
fn on_load() {
    let exports = ScExports::new();
    exports.add_call("storeString", store_string);
    exports.add_view("getString", get_string);
}

fn store_string(ctx: &ScCallContext) {
    let par = ctx.params().get_string("paramString");
    if !par.exists(){
        ctx.panic("string parameter not found")
    }
    ctx.state().get_string("storedString").set_value(&par.value());
    let msg = "Message stored: ".to_string() + &par.value();
    ctx.log(&msg);
}

fn get_string(ctx: &ScViewContext) {
    let s = ctx.state().get_string("storedString").value();
    ctx.results().get_string("paramString").set_value(&s);
}

