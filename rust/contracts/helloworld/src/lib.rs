// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use wasplib::client::*;
use crate::schema::VAR_HELLO_WORLD;

mod schema;

fn func_hello_world(ctx: &ScCallContext) {
    ctx.log("Hello, world!");
}

fn view_get_hello_world(ctx: &ScViewContext) {
    ctx.log("Get Hello world!");
    ctx.results().get_string(VAR_HELLO_WORLD).set_value("Hello, world!");
}