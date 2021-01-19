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
    exports.add_call("state_test", state_test);
    exports.add_view("state_check", state_check);
    exports.add_call("results_test", results_test);
    exports.add_view("results_check", results_check);
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

fn results_test(sc: &ScCallContext) {
    test_map(sc.results());
    check_map(sc.results().immutable());
    //sc.call("results_check");
}

fn state_test(sc: &ScCallContext) {
    test_map(sc.state());
    sc.call("state_check");
}

fn results_check(sc: &ScViewContext) {
    check_map(sc.results().immutable());
}

fn state_check(sc: &ScViewContext) {
    check_map(sc.state());
}

fn test_map(kvstore: ScMutableMap) {
    let int1 = kvstore.get_int("int1");
    check(int1.value() == 0);
    int1.set_value(1);

    let string1 = kvstore.get_string("string1");
    check(string1.value() == "");
    string1.set_value("a");

    let ia1 = kvstore.get_int_array("ia1");
    let int2 = ia1.get_int(0);
    check(int2.value() == 0);
    int2.set_value(2);
    let int3 = ia1.get_int(1);
    check(int3.value() == 0);
    int3.set_value(3);

    let sa1 = kvstore.get_string_array("sa1");
    let string2 = sa1.get_string(0);
    check(string2.value() == "");
    string2.set_value("bc");
    let string3 = sa1.get_string(1);
    check(string3.value() == "");
    string3.set_value("def");
}

fn check_map(kvstore: ScImmutableMap) {
    let int1 = kvstore.get_int("int1");
    check(int1.value() == 1);

    let string1 = kvstore.get_string("string1");
    check(string1.value() == "a");

    let ia1 = kvstore.get_int_array("ia1");
    let int2 = ia1.get_int(0);
    check(int2.value() == 2);
    let int3 = ia1.get_int(1);
    check(int3.value() == 3);

    let sa1 = kvstore.get_string_array("sa1");
    let string2 = sa1.get_string(0);
    check(string2.value() == "bc");
    let string3 = sa1.get_string(1);
    check(string3.value() == "def");
}

// fn check_map_rev(kvstore: ScImmutableMap) {
//     let sa1 = kvstore.get_string_array("sa1");
//     let string3 = sa1.get_string(1);
//     check(string3.value() == "def");
//     let string2 = sa1.get_string(0);
//     check(string2.value() == "bc");
//
//     let ia1 = kvstore.get_int_array("ia1");
//     let int3 = ia1.get_int(1);
//     check(int3.value() == 3);
//     let int2 = ia1.get_int(0);
//     check(int2.value() == 2);
//
//     let string1 = kvstore.get_string("string1");
//     check(string1.value() == "a");
//
//     let int1 = kvstore.get_int("int1");
//     check(int1.value() == 1);
// }

fn check(condition: bool) {
    if !condition {
        panic!("Check failed!")
    }
}
