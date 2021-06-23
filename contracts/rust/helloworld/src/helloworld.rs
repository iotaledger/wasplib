// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use wasmlib::*;

use crate::*;

pub fn func_hello_world(ctx: &ScFuncContext, _f: &HelloWorldContext) {
    ctx.log("Hello, world!");
}

pub fn view_get_hello_world(_ctx: &ScViewContext, f: &GetHelloWorldContext) {
    f.results.hello_world().set_value("Hello, world!");
}
