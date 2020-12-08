// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

#![allow(dead_code)]
#![allow(non_snake_case)]

use wasplib::client::*;

#[no_mangle]
pub fn onLoad() {
    let exports = ScExports::new();
    exports.add_call("init", init);
}

fn init(ctx: &ScCallContext) {
    ctx.log("init call IN");
    let intParam = ctx.params().get_int("intParam");
    if !intParam.exists() {
        ctx.log("intParam does not exist")
    } else {
        ctx.log(&("intParam OK: ".to_string() + &intParam.value().to_string()))
    }
    let intParam = ctx.params().get_agent("agentIDParam");
    if !intParam.exists() {
        ctx.log("agentIDParam does not exist")
    } else {
        ctx.log(&("agentIDParam OK: ".to_string() + &intParam.value().to_string()))
    }
    let failParam = ctx.params().get_bytes("failParam");
    if failParam.exists() {
        // fail on purpose
        ctx.error().set_value("failing on purpose");
        return;
    }
    ctx.log("init call OUT");
}
