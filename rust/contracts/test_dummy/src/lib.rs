// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use wasplib::client::*;

#[no_mangle]
fn on_load() {
    let exports = ScExports::new();
    exports.add_call("init", on_init);
}

fn on_init(ctx: &ScCallContext) {
    ctx.log("on_init call IN");
    let int_param = ctx.params().get_int("intParam");
    if !int_param.exists() {
        ctx.log("intParam does not exist")
    } else {
        ctx.log(&("intParam OK: ".to_string() + &int_param.value().to_string()))
    }
    let agent_id_param = ctx.params().get_agent("agentIDParam");
    if !agent_id_param.exists() {
        ctx.log("agentIDParam does not exist")
    } else {
        ctx.log(&("agentIDParam OK: ".to_string() + &agent_id_param.value().to_string()))
    }
    let fail_param = ctx.params().get_bytes("failParam");
    if fail_param.exists() {
        // fail on purpose
        ctx.error().set_value("failing on purpose");
        return;
    }
    ctx.log("on_init call OUT");
}
