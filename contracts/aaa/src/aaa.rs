// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use wasmlib::*;

use crate::*;

pub fn func_my_func(ctx: &ScFuncContext, params: &FuncMyFuncParams) {
    ctx.log("calling myFunc");
}

pub fn view_my_view(ctx: &ScViewContext, params: &ViewMyViewParams) {
    ctx.log("calling myView");
}
