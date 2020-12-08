// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

#![allow(dead_code)]
#![allow(non_snake_case)]

use wasplib::client::*;
use wasplib::client::host::*;

const KEY_COUNTER: &str = "counter";
const KEY_NUM_REPEATS: &str = "numRepeats";

static mut LOCAL_STATE_MUST_INCREMENT: bool = false;

#[no_mangle]
pub fn onLoad() {
    let exports = ScExports::new();
    exports.add_call("init", init);
    exports.add_call("increment", increment);
    exports.add_call("incrementCallIncrement", incrementCallIncrement);
    exports.add_call("incrementCallIncrementRecurse5x", incrementCallIncrementRecurse5x);
    exports.add_call("incrementPostIncrement", incrementPostIncrement);
    exports.add_view("incrementViewCounter", incrementViewCounter);
    exports.add_call("incrementRepeatMany", incrementRepeatMany);
    exports.add_call("incrementWhenMustIncrement", incrementWhenMustIncrement);
    exports.add_call("incrementLocalStateInternalCall", incrementLocalStateInternalCall);
    exports.add_call("incrementLocalStateSandboxCall", incrementLocalStateSandboxCall);
    exports.add_call("incrementLocalStatePost", incrementLocalStatePost);
    exports.add_call("nothing", ScExports::nothing);
    exports.add_call("test", test);
}

fn init(sc: &ScCallContext) {
    let counter = sc.params().get_int(KEY_COUNTER).value();
    if counter == 0 {
        return;
    }
    sc.state().get_int(KEY_COUNTER).set_value(counter);
}

fn increment(sc: &ScCallContext) {
    let counter = sc.state().get_int(KEY_COUNTER);
    counter.set_value(counter.value() + 1);
}

fn incrementCallIncrement(sc: &ScCallContext) {
    let counter = sc.state().get_int(KEY_COUNTER);
    let value = counter.value();
    counter.set_value(value + 1);
    if value == 0 {
        sc.call("incrementCallIncrement").call();
    }
}

fn incrementCallIncrementRecurse5x(sc: &ScCallContext) {
    let counter = sc.state().get_int(KEY_COUNTER);
    let value = counter.value();
    counter.set_value(value + 1);
    if value < 5 {
        sc.call("incrementCallIncrementRecurse5x").call();
    }
}

fn incrementPostIncrement(sc: &ScCallContext) {
    let counter = sc.state().get_int(KEY_COUNTER);
    let value = counter.value();
    counter.set_value(value + 1);
    if value == 0 {
        sc.post("incrementPostIncrement").post(0);
    }
}

fn incrementViewCounter(sc: &ScViewContext) {
    let counter = sc.state().get_int(KEY_COUNTER).value();
    sc.results().get_int(KEY_COUNTER).set_value(counter);
}

fn incrementRepeatMany(sc: &ScCallContext) {
    let counter = sc.state().get_int(KEY_COUNTER);
    let value = counter.value();
    counter.set_value(value + 1);
    let state_repeats = sc.state().get_int(KEY_NUM_REPEATS);
    let mut repeats = sc.params().get_int(KEY_NUM_REPEATS).value();
    if repeats == 0 {
        repeats = state_repeats.value();
        if repeats == 0 {
            return;
        }
    }
    state_repeats.set_value(repeats - 1);
    sc.post("incrementRepeatMany").post(0);
}

fn incrementWhenMustIncrement(sc: &ScCallContext) {
    sc.log("incrementWhenMustIncrement called");
    unsafe {
        if !LOCAL_STATE_MUST_INCREMENT {
            return;
        }
    }
    let counter = sc.state().get_int(KEY_COUNTER);
    counter.set_value(counter.value() + 1);
}

fn incrementLocalStateInternalCall(sc: &ScCallContext) {
    incrementWhenMustIncrement(sc);
    unsafe {
        LOCAL_STATE_MUST_INCREMENT = true;
    }
    incrementWhenMustIncrement(sc);
    incrementWhenMustIncrement(sc);
    // counter ends up as 2
}

fn incrementLocalStateSandboxCall(sc: &ScCallContext) {
    sc.call("incrementWhenMustIncrement").call();
    unsafe {
        LOCAL_STATE_MUST_INCREMENT = true;
    }
    sc.call("incrementWhenMustIncrement").call();
    sc.call("incrementWhenMustIncrement").call();
    // counter ends up as 0
}

fn incrementLocalStatePost(sc: &ScCallContext) {
    sc.post("incrementWhenMustIncrement").post(0);
    unsafe {
        LOCAL_STATE_MUST_INCREMENT = true;
    }
    sc.post("incrementWhenMustIncrement").post(0);
    sc.post("incrementWhenMustIncrement").post(0);
    // counter ends up as 0
}

fn test(_sc: &ScCallContext) {
    let key_id = get_key_id_from_string("timestamp");
    set_int(1, key_id, 123456789);
    let timestamp = get_int(1, key_id);
    set_int(1, key_id, timestamp);
    let key_id2 = get_key_id_from_string("string");
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
