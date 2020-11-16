// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

#![allow(dead_code)]
#![allow(non_snake_case)]

use wasplib::client::*;
use wasplib::client::host::*;

#[no_mangle]
pub fn onLoad() {
    let mut exports = ScExports::new();
    exports.add("increment");
    exports.add("incrementRepeat1");
    exports.add("incrementRepeatMany");
    exports.add("test");
    exports.add("nothing");
    exports.add("init");
}

#[no_mangle]
pub fn init() {
    let sc = ScContext::new();
    let counter = sc.request().params().get_int("counter").value();
    if counter == 0 {
        return;
    }
    sc.state().get_int("counter").set_value(counter);
}

#[no_mangle]
pub fn increment() {
    let sc = ScContext::new();
    let counter = sc.state().get_int("counter");
    counter.set_value(counter.value() + 1);
}

#[no_mangle]
pub fn incrementRepeat1() {
    let sc = ScContext::new();
    let counter = sc.state().get_int("counter");
    let value = counter.value();
    counter.set_value(value + 1);
    if value == 0 {
        sc.post_request(&sc.contract().id(), "increment", 0);
    }
}

#[no_mangle]
pub fn incrementRepeatMany() {
    let sc = ScContext::new();
    let counter = sc.state().get_int("counter");
    let value = counter.value();
    counter.set_value(value + 1);
    let state_repeats = sc.state().get_int("numRepeats");
    let mut repeats = sc.request().params().get_int("numRepeats").value();
    if repeats == 0 {
        repeats = state_repeats.value();
        if repeats == 0 {
            return;
        }
    }
    state_repeats.set_value(repeats - 1);
    sc.post_request(&sc.contract().id(), "incrementRepeatMany", 0);
}

#[no_mangle]
pub fn test() {
    let key_id = get_key_id("timestamp");
    set_int(1, key_id, 123456789);
    let timestamp = get_int(1, key_id);
    set_int(1, key_id, timestamp);
    let key_id2 = get_key_id("string");
    set_string(1, key_id2, "Test");
    let s1 = get_string(1, key_id2);
    set_string(1, key_id2, "Bleep");
    let s2 = get_string(1, key_id2);
    set_string(1, key_id2, "Klunky");
    let s3 = get_string(1, key_id2);
    set_string(1, key_id2, &s1);
    set_string(1, key_id2, &s2);
    set_string(1, key_id2, &s3);
}
