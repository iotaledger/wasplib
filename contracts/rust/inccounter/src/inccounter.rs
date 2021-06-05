// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use wasmlib::*;

use crate::*;

static mut LOCAL_STATE_MUST_INCREMENT: bool = false;

pub fn func_call_increment(ctx: &ScFuncContext, f: &FuncCallIncrementContext) {
    let counter = f.state.counter();
    let value = counter.value();
    counter.set_value(value + 1);
    if value == 0 {
        ctx.call_self(HFUNC_CALL_INCREMENT, None, None);
    }
}

pub fn func_call_increment_recurse5x(ctx: &ScFuncContext, f: &FuncCallIncrementRecurse5xContext) {
    let counter = f.state.counter();
    let value = counter.value();
    counter.set_value(value + 1);
    if value < 5 {
        ctx.call_self(HFUNC_CALL_INCREMENT_RECURSE5X, None, None);
    }
}

pub fn func_increment(_ctx: &ScFuncContext, f: &FuncIncrementContext) {
    let counter = f.state.counter();
    counter.set_value(counter.value() + 1);
}

pub fn func_init(_ctx: &ScFuncContext, f: &FuncInitContext) {
    if f.params.counter.exists() {
        let counter = f.params.counter.value();
        f.state.counter().set_value(counter);
    }
}

pub fn func_local_state_internal_call(ctx: &ScFuncContext, f: &FuncLocalStateInternalCallContext) {
    unsafe {
        LOCAL_STATE_MUST_INCREMENT = false;
    }
    func_when_must_increment_state(ctx, &f.state);
    unsafe {
        LOCAL_STATE_MUST_INCREMENT = true;
    }
    func_when_must_increment_state(ctx, &f.state);
    func_when_must_increment_state(ctx, &f.state);
    // counter ends up as 2
}

pub fn func_local_state_post(ctx: &ScFuncContext, _f: &FuncLocalStatePostContext) {
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

fn local_state_post(ctx: &ScFuncContext, nr: i64) {
    //note: we add a dummy parameter here to prevent "duplicate outputs not allowed" error
    let params = ScMutableMap::new();
    params.get_int64(VAR_COUNTER).set_value(nr);
    let transfer = ScTransfers::iotas(1);
    ctx.post_self(HFUNC_WHEN_MUST_INCREMENT, Some(params), transfer, 0);
}

pub fn func_local_state_sandbox_call(ctx: &ScFuncContext, _f: &FuncLocalStateSandboxCallContext) {
    unsafe {
        LOCAL_STATE_MUST_INCREMENT = false;
    }
    ctx.call_self(HFUNC_WHEN_MUST_INCREMENT, None, None);
    unsafe {
        LOCAL_STATE_MUST_INCREMENT = true;
    }
    ctx.call_self(HFUNC_WHEN_MUST_INCREMENT, None, None);
    ctx.call_self(HFUNC_WHEN_MUST_INCREMENT, None, None);
    // counter ends up as 0
}

pub fn func_loop(_ctx: &ScFuncContext, _f: &FuncLoopContext) {
    loop {}
}

pub fn func_post_increment(ctx: &ScFuncContext, f: &FuncPostIncrementContext) {
    let counter = f.state.counter();
    let value = counter.value();
    counter.set_value(value + 1);
    if value == 0 {
        let transfer = ScTransfers::iotas(1);
        ctx.post_self(HFUNC_POST_INCREMENT, None, transfer, 0);
    }
}

pub fn func_repeat_many(ctx: &ScFuncContext, f: &FuncRepeatManyContext) {
    let counter = f.state.counter();
    let value = counter.value();
    counter.set_value(value + 1);
    let state_repeats = f.state.num_repeats();
    let mut repeats = f.params.num_repeats.value();
    if repeats == 0 {
        repeats = state_repeats.value();
        if repeats == 0 {
            return;
        }
    }
    state_repeats.set_value(repeats - 1);
    let transfer = ScTransfers::iotas(1);
    ctx.post_self(HFUNC_REPEAT_MANY, None, transfer, 0);
}

pub fn func_when_must_increment(ctx: &ScFuncContext, f: &FuncWhenMustIncrementContext) {
    func_when_must_increment_state(ctx, &f.state);
}

pub fn func_when_must_increment_state(ctx: &ScFuncContext, state: &IncCounterFuncState) {
    ctx.log("when_must_increment called");
    unsafe {
        if !LOCAL_STATE_MUST_INCREMENT {
            return;
        }
    }
    let counter = state.counter();
    counter.set_value(counter.value() + 1);
}

// note that get_counter mirrors the state of the 'counter' state variable
// which means that if the state variable was not present it also will not be present in the result
pub fn view_get_counter(_ctx: &ScViewContext, f: &ViewGetCounterContext) {
    let counter = f.state.counter();
    if counter.exists() {
        f.results.counter.set_value(counter.value());
    }
}

pub fn func_test_leb128(ctx: &ScFuncContext, _f: &FuncTestLeb128Context) {
    save(ctx, "v-1", -1);
    save(ctx, "v-2", -2);
    save(ctx, "v-126", -126);
    save(ctx, "v-127", -127);
    save(ctx, "v-128", -128);
    save(ctx, "v-129", -129);
    save(ctx, "v0", 0);
    save(ctx, "v+1", 1);
    save(ctx, "v+2", 2);
    save(ctx, "v+126", 126);
    save(ctx, "v+127", 127);
    save(ctx, "v+128", 128);
    save(ctx, "v+129", 129);
}

fn save(ctx: &ScFuncContext, name: &str, value: i64) {
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
