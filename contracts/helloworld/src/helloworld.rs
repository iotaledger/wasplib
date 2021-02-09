// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use wasmlib::*;

use crate::*;

pub fn func_hello_world(ctx: &ScFuncContext, _params: &FuncHelloWorldParams) {
    ctx.log("Hello, world!");
}

pub fn view_get_hello_world(ctx: &ScViewContext, _params: &ViewGetHelloWorldParams) {
    ctx.log("Get Hello world!");
    ctx.results().get_string(VAR_HELLO_WORLD).set_value("Hello, world!");
}
