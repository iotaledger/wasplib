// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use wasplib::client::*;

const KEY_COUNTER: &str = "counter";

#[no_mangle]
fn on_load() {
    let exports = ScExports::new();
    exports.add_call("hello", hello);
    exports.add_view("getCounter", get_counter);
}

// Function hello implements smart contract entry point "hello".
// Function hello logs the message "Hello, new world!" with the counter and increments the counter
fn hello(ctx: &ScCallContext) {
    let counter = ctx.state().get_int(KEY_COUNTER);
    let msg = "Hello, new world! #".to_string() + &counter.to_string();
    ctx.log(&msg);  // TODO info and debug levels, not events!
    counter.set_value(counter.value() + 1);
}

// Function get_counter implements smart contract VIEW entry point "getCounter".
// It return counter value in the result dictionary with the key "counter"
fn get_counter(ctx: &ScViewContext) {
    let counter = ctx.state().get_int(KEY_COUNTER).value();
    ctx.results().get_int(KEY_COUNTER).set_value(counter);
}
