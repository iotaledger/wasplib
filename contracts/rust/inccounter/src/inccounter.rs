// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use wasmlib::*;

use crate::*;

static mut LOCAL_STATE_MUST_INCREMENT: bool = false;

pub fn func_call_increment(ctx: &ScFuncContext) {
    let counter = ctx.state().get_int(VAR_COUNTER);
    let value = counter.value();
    counter.set_value(value + 1);
    if value == 0 {
        ctx.call_self(HFUNC_CALL_INCREMENT, None, None);
    }
}

pub fn func_call_increment_recurse5x(ctx: &ScFuncContext) {
    let counter = ctx.state().get_int(VAR_COUNTER);
    let value = counter.value();
    counter.set_value(value + 1);
    if value < 5 {
        ctx.call_self(HFUNC_CALL_INCREMENT_RECURSE5X, None, None);
    }
}

pub fn func_increment(ctx: &ScFuncContext) {
    let counter = ctx.state().get_int(VAR_COUNTER);
    counter.set_value(counter.value() + 1);
}

pub fn func_init(ctx: &ScFuncContext) {
    let p = ctx.params();
    let param_counter = p.get_int(PARAM_COUNTER);
    if !param_counter.exists() {
        return;
    }
    let counter = param_counter.value();
    ctx.state().get_int(VAR_COUNTER).set_value(counter);
}

pub fn func_local_state_internal_call(ctx: &ScFuncContext) {
    unsafe {
        LOCAL_STATE_MUST_INCREMENT = false;
    }
    func_when_must_increment(ctx);
    unsafe {
        LOCAL_STATE_MUST_INCREMENT = true;
    }
    func_when_must_increment(ctx);
    func_when_must_increment(ctx);
    // counter ends up as 2
}

pub fn func_local_state_post(ctx: &ScFuncContext) {
    unsafe {
        LOCAL_STATE_MUST_INCREMENT = false;
    }
    let request = PostRequestParams {
        contract_id: ctx.contract_id(),
        function: HFUNC_WHEN_MUST_INCREMENT,
        params: None,
        transfer: None,
        delay: 0,
    };
    ctx.post(&request);
    unsafe {
        LOCAL_STATE_MUST_INCREMENT = true;
    }
    ctx.post(&request);
    ctx.post(&request);
    // counter ends up as 0
}

pub fn func_local_state_sandbox_call(ctx: &ScFuncContext) {
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

pub fn func_post_increment(ctx: &ScFuncContext) {
    let counter = ctx.state().get_int(VAR_COUNTER);
    let value = counter.value();
    counter.set_value(value + 1);
    if value == 0 {
        ctx.post(&PostRequestParams {
            contract_id: ctx.contract_id(),
            function: HFUNC_POST_INCREMENT,
            params: None,
            transfer: None,
            delay: 0,
        });
    }
}

pub fn func_repeat_many(ctx: &ScFuncContext) {
    let p = ctx.params();
    let param_num_repeats = p.get_int(PARAM_NUM_REPEATS);

    let counter = ctx.state().get_int(VAR_COUNTER);
    let value = counter.value();
    counter.set_value(value + 1);
    let state_repeats = ctx.state().get_int(VAR_NUM_REPEATS);
    let mut repeats = param_num_repeats.value();
    if repeats == 0 {
        repeats = state_repeats.value();
        if repeats == 0 {
            return;
        }
    }
    state_repeats.set_value(repeats - 1);
    ctx.post(&PostRequestParams {
        contract_id: ctx.contract_id(),
        function: HFUNC_REPEAT_MANY,
        params: None,
        transfer: None,
        delay: 0,
    });
}

pub fn func_when_must_increment(ctx: &ScFuncContext) {
    ctx.log("when_must_increment called");
    unsafe {
        if !LOCAL_STATE_MUST_INCREMENT {
            return;
        }
    }
    let counter = ctx.state().get_int(VAR_COUNTER);
    counter.set_value(counter.value() + 1);
}

// note that get_counter mirrors the state of the 'counter' state variable
// which means that if the state variable was not present it also will not be present in the result
pub fn view_get_counter(ctx: &ScViewContext) {
    let counter = ctx.state().get_int(VAR_COUNTER);
    if counter.exists() {
        ctx.results().get_int(VAR_COUNTER).set_value(counter.value());
    }
}
