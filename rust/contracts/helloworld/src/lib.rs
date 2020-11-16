// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

#![allow(dead_code)]
#![allow(non_snake_case)]

use wasplib::client::*;

#[no_mangle]
pub fn onLoad() {
    let mut exports = ScExports::new();
    exports.add("helloWorld");
}

#[no_mangle]
pub fn helloWorld() {
    let sc = ScContext::new();
    sc.log("Hello, world!");
}
