// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

use aaa::*;
use schema::*;
use wasmlib::*;

mod aaa;
mod schema;

#[no_mangle]
fn on_load() {
    let exports = ScExports::new();
    exports.add_func(FUNC_MY_FUNC, func_my_func_thunk);
    exports.add_view(VIEW_MY_VIEW, view_my_view_thunk);
}

pub struct FuncMyFuncParams {
    pub text: ScImmutableString, // some mandatory string parameter
}

fn func_my_func_thunk(ctx: &ScFuncContext) {
    let p = ctx.params();
    let params = FuncMyFuncParams {
        text: p.get_string(PARAM_TEXT),
    };
    ctx.require(params.text.exists(), "missing mandatory text");
    func_my_func(ctx, &params);
}

pub struct ViewMyViewParams {
    pub value: ScImmutableInt, // some optional integer parameter
}

fn view_my_view_thunk(ctx: &ScViewContext) {
    let p = ctx.params();
    let params = ViewMyViewParams {
        value: p.get_int(PARAM_VALUE),
    };
    view_my_view(ctx, &params);
}
