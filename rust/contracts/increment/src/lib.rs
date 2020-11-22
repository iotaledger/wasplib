// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

#![allow(dead_code)]
#![allow(non_snake_case)]

use wasplib::client::*;
use wasplib::client::host::*;

#[no_mangle]
pub fn onLoad() {
    let exports = ScExports::new();
    exports.add_call("increment", increment);
    exports.add_call("incrementRepeat1", incrementRepeat1);
    exports.add_call("incrementRepeatMany", incrementRepeatMany);
    exports.add_call("test", test);
    exports.add_call("nothing", ScExports::nothing);
    exports.add_call("init", init);
}

pub fn init(sc: &ScCallContext) {
    let counter = sc.request().params().get_int("counter").value();
    if counter == 0 {
        return;
    }
    sc.state().get_int("counter").set_value(counter);
}

pub fn increment(sc: &ScCallContext) {
    let counter = sc.state().get_int("counter");
    counter.set_value(counter.value() + 1);
}

pub fn incrementRepeat1(sc: &ScCallContext) {
    let counter = sc.state().get_int("counter");
    let value = counter.value();
    counter.set_value(value + 1);
    if value == 0 {
        sc.post_self("increment", 0);
    }
}

pub fn incrementRepeatMany(sc: &ScCallContext) {
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
    sc.post_self("incrementRepeatMany", 0);
}

pub fn test(_sc: &ScCallContext) {
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
