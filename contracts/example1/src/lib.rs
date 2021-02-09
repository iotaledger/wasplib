// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

use example1::*;
use schema::*;
use wasmlib::*;

mod example1;
mod schema;

#[no_mangle]
fn on_load() {
    let exports = ScExports::new();
    exports.add_call(FUNC_STORE_STRING, func_store_string_thunk);
    exports.add_call(FUNC_WITHDRAW_IOTA, func_withdraw_iota_thunk);
    exports.add_view(VIEW_GET_STRING, view_get_string_thunk);
}

pub struct FuncStoreStringParams {
    pub string: ScImmutableString, // string to store
}

fn func_store_string_thunk(ctx: &ScCallContext) {
    let p = ctx.params();
    let params = FuncStoreStringParams {
        string: p.get_string(PARAM_STRING),
    };
    ctx.require(params.string.exists(), "missing mandatory string");
    func_store_string(ctx, &params);
}

pub struct FuncWithdrawIotaParams {
}

fn func_withdraw_iota_thunk(ctx: &ScCallContext) {
    // only the contract creator can withdraw
    ctx.require(ctx.from(&ctx.contract_creator()), "no permission");

    let params = FuncWithdrawIotaParams {
    };
    func_withdraw_iota(ctx, &params);
}

pub struct ViewGetStringParams {
}

fn view_get_string_thunk(ctx: &ScViewContext) {
    let params = ViewGetStringParams {
    };
    view_get_string(ctx, &params);
}
