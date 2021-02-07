// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use wasmlib::*;

const PARAM_FAIL_INIT_PARAM: &str = "failInitParam";

#[no_mangle]
fn on_load() {
    let exports = ScExports::new();
    exports.add_call("init", on_init);
}

// fails with error if failInitParam exists
fn on_init(ctx: &ScCallContext) {
    let fail_param = ctx.params().get_int(PARAM_FAIL_INIT_PARAM);
    if fail_param.exists() {
        ctx.panic("dummy: failing on purpose");
    }
}
