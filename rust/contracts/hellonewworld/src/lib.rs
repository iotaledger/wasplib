// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use wasplib::client::*;

const KEY_COUNTER: &str = "counter";
const KEY_PANIC: &str = "panic";

#[no_mangle]
fn on_load() {
    let exports = ScExports::new();
    exports.add_call("hello", hello);
    exports.add_view("getCounter", get_counter);
}

fn hello(ctx: &ScCallContext) {
    if ctx.params().get_bytes(KEY_PANIC).exists() {
        ctx.panic("panic instead of Hello");
    }
    let counter = ctx.state().get_int(KEY_COUNTER);
    let msg = "Hello, new world! #".to_string() + &counter.to_string();
    ctx.log(&msg);
    counter.set_value(counter.value() + 1);
}

fn get_counter(ctx: &ScViewContext) {
    let counter = ctx.state().get_int(KEY_COUNTER).value();
    ctx.results().get_int(KEY_COUNTER).set_value(counter);
}

