// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use wasmlib::*;

use crate::*;
use crate::contract::*;

static mut LOCAL_STATE_MUST_INCREMENT: bool = false;

pub fn func_init(_ctx: &ScFuncContext, f: &InitContext) {
    if f.params.counter().exists() {
        let counter = f.params.counter().value();
        f.state.counter().set_value(counter);
    }
}

pub fn func_call_increment(ctx: &ScFuncContext, f: &CallIncrementContext) {
    let counter = f.state.counter();
    let value = counter.value();
    counter.set_value(value + 1);
    if value == 0 {
        ScFuncs::call_increment(ctx).func.call();
    }
}

pub fn func_call_increment_recurse5x(ctx: &ScFuncContext, f: &CallIncrementRecurse5xContext) {
    let counter = f.state.counter();
    let value = counter.value();
    counter.set_value(value + 1);
    if value < 5 {
        ScFuncs::call_increment_recurse5x(ctx).func.call();
    }
}

pub fn func_endless_loop(_ctx: &ScFuncContext, _f: &EndlessLoopContext) {
    loop {}
}

pub fn func_increment(_ctx: &ScFuncContext, f: &IncrementContext) {
    let counter = f.state.counter();
    counter.set_value(counter.value() + 1);
}

pub fn func_local_state_internal_call(ctx: &ScFuncContext, f: &LocalStateInternalCallContext) {
    unsafe {
        LOCAL_STATE_MUST_INCREMENT = false;
    }
    when_must_increment_state(ctx, &f.state);
    unsafe {
        LOCAL_STATE_MUST_INCREMENT = true;
    }
    when_must_increment_state(ctx, &f.state);
    when_must_increment_state(ctx, &f.state);
    // counter ends up as 2
}

pub fn func_local_state_post(ctx: &ScFuncContext, _f: &LocalStatePostContext) {
    unsafe {
        LOCAL_STATE_MUST_INCREMENT = false;
    }
    // prevent multiple identical posts, need a dummy param to differentiate them
    local_state_post(ctx, 1);
    unsafe {
        LOCAL_STATE_MUST_INCREMENT = true;
    }
    local_state_post(ctx, 2);
    local_state_post(ctx, 3);
    // counter ends up as 0
}

pub fn func_local_state_sandbox_call(ctx: &ScFuncContext, _f: &LocalStateSandboxCallContext) {
    unsafe {
        LOCAL_STATE_MUST_INCREMENT = false;
    }
    ScFuncs::when_must_increment(ctx).func.call();
    unsafe {
        LOCAL_STATE_MUST_INCREMENT = true;
    }
    ScFuncs::when_must_increment(ctx).func.call();
    ScFuncs::when_must_increment(ctx).func.call();
    // counter ends up as 0
}

pub fn func_post_increment(ctx: &ScFuncContext, f: &PostIncrementContext) {
    let counter = f.state.counter();
    let value = counter.value();
    counter.set_value(value + 1);
    if value == 0 {
        ScFuncs::increment(ctx).func.transfer_iotas(1).post();
    }
}

pub fn func_repeat_many(ctx: &ScFuncContext, f: &RepeatManyContext) {
    let counter = f.state.counter();
    let value = counter.value();
    counter.set_value(value + 1);
    let state_repeats = f.state.num_repeats();
    let mut repeats = f.params.num_repeats().value();
    if repeats == 0 {
        repeats = state_repeats.value();
        if repeats == 0 {
            return;
        }
    }
    state_repeats.set_value(repeats - 1);
    ScFuncs::repeat_many(ctx).func.transfer_iotas(1).post();
}

pub fn func_test_leb128(ctx: &ScFuncContext, _f: &TestLeb128Context) {
    leb128_save(ctx, "v-1", -1);
    leb128_save(ctx, "v-2", -2);
    leb128_save(ctx, "v-126", -126);
    leb128_save(ctx, "v-127", -127);
    leb128_save(ctx, "v-128", -128);
    leb128_save(ctx, "v-129", -129);
    leb128_save(ctx, "v0", 0);
    leb128_save(ctx, "v+1", 1);
    leb128_save(ctx, "v+2", 2);
    leb128_save(ctx, "v+126", 126);
    leb128_save(ctx, "v+127", 127);
    leb128_save(ctx, "v+128", 128);
    leb128_save(ctx, "v+129", 129);
}

pub fn func_when_must_increment(ctx: &ScFuncContext, f: &WhenMustIncrementContext) {
    when_must_increment_state(ctx, &f.state);
}

// note that get_counter mirrors the state of the 'counter' state variable
// which means that if the state variable was not present it also will not be present in the result
pub fn view_get_counter(_ctx: &ScViewContext, f: &GetCounterContext) {
    let counter = f.state.counter();
    if counter.exists() {
        f.results.counter().set_value(counter.value());
    }
}

fn leb128_save(ctx: &ScFuncContext, name: &str, value: i64) {
    let mut encoder = BytesEncoder::new();
    encoder.int64(value);
    let spot = ctx.state().get_bytes(name);
    spot.set_value(&encoder.data());

    let bytes = spot.value();
    let mut decoder = BytesDecoder::new(&bytes);
    let retrieved = decoder.int64();
    if retrieved != value {
        ctx.log(&(name.to_string() + " in : " + &value.to_string()));
        ctx.log(&(name.to_string() + " out: " + &retrieved.to_string()));
    }
}

fn local_state_post(ctx: &ScFuncContext, nr: i64) {
    //note: we add a dummy parameter here to prevent "duplicate outputs not allowed" error
    let f = ScFuncs::when_must_increment(ctx);
    f.params.dummy().set_value(nr);
    f.func.transfer_iotas(1).post();
}

fn when_must_increment_state(ctx: &ScFuncContext, state: &MutableIncCounterState) {
    ctx.log("when_must_increment called");
    unsafe {
        if !LOCAL_STATE_MUST_INCREMENT {
            return;
        }
    }
    let counter = state.counter();
    counter.set_value(counter.value() + 1);
}
