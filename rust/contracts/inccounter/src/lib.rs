// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use wasplib::client::*;
use wasplib::client::host::*;

const KEY_COUNTER: &str = "counter";
const KEY_NUM_REPEATS: &str = "num_repeats";

static mut LOCAL_STATE_MUST_INCREMENT: bool = false;

#[no_mangle]
fn on_load() {
    let exports = ScExports::new();
    exports.add_call("init", on_init);
    exports.add_call("increment", increment);
    exports.add_call("increment_call_increment", increment_call_increment);
    exports.add_call("increment_call_increment_recurse5x", increment_call_increment_recurse5x);
    exports.add_call("increment_post_increment", increment_post_increment);
    exports.add_view("increment_view_counter", increment_view_counter);
    exports.add_call("increment_repeat_many", increment_repeat_many);
    exports.add_call("increment_when_must_increment", increment_when_must_increment);
    exports.add_call("increment_local_state_internal_call", increment_local_state_internal_call);
    exports.add_call("increment_local_state_sandbox_call", increment_local_state_sandbox_call);
    exports.add_call("increment_local_state_post", increment_local_state_post);
    exports.add_call("nothing", ScExports::nothing);
    exports.add_call("test", test);
}

fn on_init(sc: &ScCallContext) {
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

fn increment_call_increment(sc: &ScCallContext) {
    let counter = sc.state().get_int(KEY_COUNTER);
    let value = counter.value();
    counter.set_value(value + 1);
    if value == 0 {
        sc.call("increment_call_increment").call();
    }
}

fn increment_call_increment_recurse5x(sc: &ScCallContext) {
    let counter = sc.state().get_int(KEY_COUNTER);
    let value = counter.value();
    counter.set_value(value + 1);
    if value < 5 {
        sc.call("increment_call_increment_recurse5x").call();
    }
}

fn increment_post_increment(sc: &ScCallContext) {
    let counter = sc.state().get_int(KEY_COUNTER);
    let value = counter.value();
    counter.set_value(value + 1);
    if value == 0 {
        sc.post("increment_post_increment").post(0);
    }
}

fn increment_view_counter(sc: &ScViewContext) {
    let counter = sc.state().get_int(KEY_COUNTER).value();
    sc.results().get_int(KEY_COUNTER).set_value(counter);
}

fn increment_repeat_many(sc: &ScCallContext) {
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
    sc.post("increment_repeat_many").post(0);
}

fn increment_when_must_increment(sc: &ScCallContext) {
    sc.log("increment_when_must_increment called");
    unsafe {
        if !LOCAL_STATE_MUST_INCREMENT {
            return;
        }
    }
    let counter = sc.state().get_int(KEY_COUNTER);
    counter.set_value(counter.value() + 1);
}

fn increment_local_state_internal_call(sc: &ScCallContext) {
    increment_when_must_increment(sc);
    unsafe {
        LOCAL_STATE_MUST_INCREMENT = true;
    }
    increment_when_must_increment(sc);
    increment_when_must_increment(sc);
    // counter ends up as 2
}

fn increment_local_state_sandbox_call(sc: &ScCallContext) {
    sc.call("increment_when_must_increment").call();
    unsafe {
        LOCAL_STATE_MUST_INCREMENT = true;
    }
    sc.call("increment_when_must_increment").call();
    sc.call("increment_when_must_increment").call();
    // counter ends up as 0
}

fn increment_local_state_post(sc: &ScCallContext) {
    sc.post("increment_when_must_increment").post(0);
    unsafe {
        LOCAL_STATE_MUST_INCREMENT = true;
    }
    sc.post("increment_when_must_increment").post(0);
    sc.post("increment_when_must_increment").post(0);
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
