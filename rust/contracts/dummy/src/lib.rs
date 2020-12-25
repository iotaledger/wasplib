// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use wasplib::client::*;

#[no_mangle]
fn on_load() {
    let exports = ScExports::new();
    exports.add_call("init", on_init);
}

// fails with error if failInitParam exists
fn on_init(ctx: &ScCallContext) {
    let fail_param = ctx.params().get_int("failInitParam");
    if fail_param.exists(){
        let err = "dummy: failing on purpose";
        ctx.log(err);
        ctx.error().set_value(err);
        return;
    }
}
