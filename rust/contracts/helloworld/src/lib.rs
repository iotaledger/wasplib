// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use wasplib::client::*;

#[no_mangle]
fn on_load() {
    let exports = ScExports::new();
    exports.add_call("hello_world", hello_world);
}

fn hello_world(ctx: &ScCallContext) {
    ctx.log("Hello, world!");
}
