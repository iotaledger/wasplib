// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use wasmlib::*;

use crate::*;

static mut LOCAL_STATE_MUST_INCREMENT: bool = false;

pub fn func_call_increment(ctx: &ScFuncContext, _params: &FuncCallIncrementParams) {
    let counter = ctx.state().get_int64(VAR_COUNTER);
    let value = counter.value();
    counter.set_value(value + 1);
    if value == 0 {
        ctx.call_self(HFUNC_CALL_INCREMENT, None, None);
    }
}

pub fn func_call_increment_recurse5x(ctx: &ScFuncContext, _params: &FuncCallIncrementRecurse5xParams) {
    let counter = ctx.state().get_int64(VAR_COUNTER);
    let value = counter.value();
    counter.set_value(value + 1);
    if value < 5 {
        ctx.call_self(HFUNC_CALL_INCREMENT_RECURSE5X, None, None);
    }
}

pub fn func_increment(ctx: &ScFuncContext, _params: &FuncIncrementParams) {
    let counter = ctx.state().get_int64(VAR_COUNTER);
    counter.set_value(counter.value() + 1);
}

pub fn func_init(ctx: &ScFuncContext, params: &FuncInitParams) {
    if params.counter.exists() {
        let counter = params.counter.value();
        ctx.state().get_int64(VAR_COUNTER).set_value(counter);
    }
}

pub fn func_local_state_internal_call(ctx: &ScFuncContext, _params: &FuncLocalStateInternalCallParams) {
    unsafe {
        LOCAL_STATE_MUST_INCREMENT = false;
    }
    let par = FuncWhenMustIncrementParams {};
    func_when_must_increment(ctx, &par);
    unsafe {
        LOCAL_STATE_MUST_INCREMENT = true;
    }
    func_when_must_increment(ctx, &par);
    func_when_must_increment(ctx, &par);
    // counter ends up as 2
}

pub fn func_local_state_post(ctx: &ScFuncContext, _params: &FuncLocalStatePostParams) {
    unsafe {
        LOCAL_STATE_MUST_INCREMENT = false;
    }
    ctx.post_self(HFUNC_WHEN_MUST_INCREMENT, None, None, 0);
    unsafe {
        LOCAL_STATE_MUST_INCREMENT = true;
    }
    ctx.post_self(HFUNC_WHEN_MUST_INCREMENT, None, None, 0);
    ctx.post_self(HFUNC_WHEN_MUST_INCREMENT, None, None, 0);
    // counter ends up as 0
}

pub fn func_local_state_sandbox_call(ctx: &ScFuncContext, _params: &FuncLocalStateSandboxCallParams) {
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

pub fn func_post_increment(ctx: &ScFuncContext, _params: &FuncPostIncrementParams) {
    let counter = ctx.state().get_int64(VAR_COUNTER);
    let value = counter.value();
    counter.set_value(value + 1);
    if value == 0 {
        ctx.post_self(HFUNC_POST_INCREMENT, None, None, 0);
    }
}

pub fn func_repeat_many(ctx: &ScFuncContext, params: &FuncRepeatManyParams) {
    let counter = ctx.state().get_int64(VAR_COUNTER);
    let value = counter.value();
    counter.set_value(value + 1);
    let state_repeats = ctx.state().get_int64(VAR_NUM_REPEATS);
    let mut repeats = params.num_repeats.value();
    if repeats == 0 {
        repeats = state_repeats.value();
        if repeats == 0 {
            return;
        }
    }
    state_repeats.set_value(repeats - 1);
    ctx.post_self(HFUNC_REPEAT_MANY, None, None, 0);
}

pub fn func_when_must_increment(ctx: &ScFuncContext, _params: &FuncWhenMustIncrementParams) {
    ctx.log("when_must_increment called");
    unsafe {
        if !LOCAL_STATE_MUST_INCREMENT {
            return;
        }
    }
    let counter = ctx.state().get_int64(VAR_COUNTER);
    counter.set_value(counter.value() + 1);
}

// note that get_counter mirrors the state of the 'counter' state variable
// which means that if the state variable was not present it also will not be present in the result
pub fn view_get_counter(ctx: &ScViewContext, _params: &ViewGetCounterParams) {
    let counter = ctx.state().get_int64(VAR_COUNTER);
    if counter.exists() {
        ctx.results().get_int64(VAR_COUNTER).set_value(counter.value());
    }
}
