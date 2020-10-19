#![allow(dead_code)]
#![allow(non_snake_case)]

use wasplib::client::ScContext;
use wasplib::client::ScExports;

#[no_mangle]
pub fn onLoad() {
    let mut exports = ScExports::new();
    exports.add_protected("initSC");
    exports.add("transfer");
    exports.add("approve")
}

#[no_mangle]
pub fn initSC() {
    let sc = ScContext::new();
    // TODO 1
    sc.log("initSC: success");
}

#[no_mangle]
pub fn transfer() {
    let sc = ScContext::new();
    // TODO
    sc.log("transfer");
}

#[no_mangle]
pub fn approve() {
    let sc = ScContext::new();
    // TODO
    sc.log("approve");
}

