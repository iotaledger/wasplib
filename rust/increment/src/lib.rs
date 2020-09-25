#![allow(dead_code)]
#![allow(non_snake_case)]

use wasplib::client::host::{get_int, get_key_id, get_string, set_int, set_string};
use wasplib::client::ScContext;

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

#[no_mangle]
pub fn nothing() {
    let ctx = ScContext::new();
    ctx.log("Doing nothing as requested. Oh, wait...");
}

#[no_mangle]
pub fn increment() {
    let ctx = ScContext::new();
    let counter = ctx.state().get_int("counter");
    counter.set_value(counter.value() + 1);
}

#[no_mangle]
pub fn incrementRepeat1() {
    let ctx = ScContext::new();
    let counter = ctx.state().get_int("counter");
    let value = counter.value();
    counter.set_value(value + 1);
    if value == 0 {
        ctx.event("", "increment", 5);
    }
}

#[no_mangle]
pub fn incrementRepeatMany() {
    let ctx = ScContext::new();
    let counter = ctx.state().get_int("counter");
    let value = counter.value();
    counter.set_value(value + 1);
    let mut repeats = ctx.request().params().get_int("numrepeats").value();
    let state_repeats = ctx.state().get_int("numrepeats");
    if repeats == 0 {
        repeats = state_repeats.value();
        if repeats == 0 {
            return;
        }
    }
    state_repeats.set_value(repeats - 1);
    ctx.event("", "incrementRepeatMany", 3);
}
