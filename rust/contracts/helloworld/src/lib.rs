// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

#![allow(dead_code)]
#![allow(non_snake_case)]

use wasplib::client::*;

#[no_mangle]
pub fn onLoad() {
    let exports = ScExports::new();
    exports.add_call("helloWorld", helloWorld);
}

pub fn helloWorld(sc: &ScCallContext) {
    sc.log("Hello, world!");
}
