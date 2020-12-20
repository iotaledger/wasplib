// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use wasplib::client::*;

#[no_mangle]
fn on_load() {
    let exports = ScExports::new();
    exports.add_call("hello", hello);
    exports.add_view("getCounter", get_counter);
}

fn hello(ctx: &ScCallContext) {
    let counter = ctx.state().get_int("counter");
    let msg = "Hello, new world! #".to_string() + &counter.to_string();
    ctx.log(&msg);
    counter.set_value(counter.value()+1);
}

fn get_counter(ctx: &ScViewContext){
    let counter = ctx.state().get_int("counter").value();
    ctx.results().get_int("counter").set_value(counter);
}
