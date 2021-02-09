// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use wasmlib::*;

use crate::*;

pub fn func_init(ctx: &ScFuncContext, params: &FuncInitParams) {
    if params.fail_init_param.exists() {
        ctx.panic("dummy: failing on purpose");
    }
}
