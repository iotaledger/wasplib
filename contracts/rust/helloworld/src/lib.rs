// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

//@formatter:off

#![allow(dead_code)]

#![allow(unused_imports)]

use helloworld::*;
use wasmlib::*;
use wasmlib::host::*;

use crate::consts::*;
use crate::keys::*;
use crate::state::*;

mod consts;
mod keys;
mod state;
mod helloworld;

#[no_mangle]
fn on_load() {
    let exports = ScExports::new();
    exports.add_func(FUNC_HELLO_WORLD, func_hello_world_thunk);
    exports.add_view(VIEW_GET_HELLO_WORLD, view_get_hello_world_thunk);

    unsafe {
        for i in 0..KEY_MAP_LEN {
            IDX_MAP[i] = get_key_id_from_string(KEY_MAP[i]);
        }
    }
}

pub struct FuncHelloWorldContext {
    state: HelloWorldFuncState,
}

fn func_hello_world_thunk(ctx: &ScFuncContext) {
    ctx.log("helloworld.funcHelloWorld");
    let f = FuncHelloWorldContext {
        state: HelloWorldFuncState {
            state_id: get_object_id(1, KEY_STATE, TYPE_MAP),
        },
    };
    func_hello_world(ctx, &f);
    ctx.log("helloworld.funcHelloWorld ok");
}

pub struct ViewGetHelloWorldResults {
    pub hello_world: ScMutableString,
}

pub struct ViewGetHelloWorldContext {
    results: ViewGetHelloWorldResults,
    state:   HelloWorldViewState,
}

fn view_get_hello_world_thunk(ctx: &ScViewContext) {
    ctx.log("helloworld.viewGetHelloWorld");
    let r = ctx.results().map_id();
    let f = ViewGetHelloWorldContext {
        results: ViewGetHelloWorldResults {
            hello_world: ScMutableString::new(r, idx_map(IDX_RESULT_HELLO_WORLD)),
        },
        state: HelloWorldViewState {
            state_id: get_object_id(1, KEY_STATE, TYPE_MAP),
        },
    };
    view_get_hello_world(ctx, &f);
    ctx.log("helloworld.viewGetHelloWorld ok");
}

//@formatter:on
