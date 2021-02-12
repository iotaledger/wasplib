// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use wasmlib::*;

use crate::*;

static mut LOCAL_STATE_MUST_INCREMENT: bool = false;

pub fn func_call_increment(ctx: &ScFuncContext, _params: &FuncCallIncrementParams) {
    let counter = ctx.state().get_int(VAR_COUNTER);
    let value = counter.value();
    counter.set_value(value + 1);
    if value == 0 {
        ctx.call_self(HFUNC_CALL_INCREMENT, None, None);
    }
}

pub fn func_call_increment_recurse5x(ctx: &ScFuncContext, _params: &FuncCallIncrementRecurse5xParams) {
    let counter = ctx.state().get_int(VAR_COUNTER);
    let value = counter.value();
    counter.set_value(value + 1);
    if value < 5 {
        ctx.call_self(HFUNC_CALL_INCREMENT_RECURSE5X, None, None);
    }
}

pub fn func_increment(ctx: &ScFuncContext, _params: &FuncIncrementParams) {
    let counter = ctx.state().get_int(VAR_COUNTER);
    counter.set_value(counter.value() + 1);
}

pub fn func_init(ctx: &ScFuncContext, params: &FuncInitParams) {
    let counter = params.counter.value();
    if counter == 0 {
        return;
    }
    ctx.state().get_int(VAR_COUNTER).set_value(counter);
}

pub fn func_local_state_internal_call(ctx: &ScFuncContext, _params: &FuncLocalStateInternalCallParams) {
    unsafe {
        LOCAL_STATE_MUST_INCREMENT = false;
    }
    let par = &FuncWhenMustIncrementParams {};
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

pub fn func_repeat_many(ctx: &ScFuncContext, params: &FuncRepeatManyParams) {
    let counter = ctx.state().get_int(VAR_COUNTER);
    let value = counter.value();
    counter.set_value(value + 1);
    let state_repeats = ctx.state().get_int(VAR_NUM_REPEATS);
    let mut repeats = params.num_repeats.value();
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

pub fn func_results_test(ctx: &ScFuncContext, _params: &FuncResultsTestParams) {
    test_map(ctx.results());
    check_map(ctx.results().immutable());
    //ctx.call_self(HFUNC_RESULTS_CHECK, None, None);
}

pub fn func_state_test(ctx: &ScFuncContext, _params: &FuncStateTestParams) {
    test_map(ctx.state());
    ctx.call_self(HVIEW_STATE_CHECK, None, None);
}

pub fn func_when_must_increment(ctx: &ScFuncContext, _params: &FuncWhenMustIncrementParams) {
    ctx.log("when_must_increment called");
    unsafe {
        if !LOCAL_STATE_MUST_INCREMENT {
            return;
        }
    }
    let counter = ctx.state().get_int(VAR_COUNTER);
    counter.set_value(counter.value() + 1);
}

pub fn view_get_counter(ctx: &ScViewContext, _params: &ViewGetCounterParams) {
    let counter = ctx.state().get_int(VAR_COUNTER).value();
    ctx.results().get_int(VAR_COUNTER).set_value(counter);
}

pub fn view_results_check(ctx: &ScViewContext, _params: &ViewResultsCheckParams) {
    check_map(ctx.results().immutable());
}

pub fn view_state_check(ctx: &ScViewContext, _params: &ViewStateCheckParams) {
    check_map(ctx.state());
}

pub fn test_map(kvstore: ScMutableMap) {
    let int1 = kvstore.get_int(VAR_INT1);
    check(int1.value() == 0);
    int1.set_value(1);

    let string1 = kvstore.get_string(VAR_STRING1);
    check(string1.value() == "");
    string1.set_value("a");

    let ia1 = kvstore.get_int_array(VAR_INT_ARRAY1);
    let int2 = ia1.get_int(0);
    check(int2.value() == 0);
    int2.set_value(2);
    let int3 = ia1.get_int(1);
    check(int3.value() == 0);
    int3.set_value(3);

    let sa1 = kvstore.get_string_array(VAR_STRING_ARRAY1);
    let string2 = sa1.get_string(0);
    check(string2.value() == "");
    string2.set_value("bc");
    let string3 = sa1.get_string(1);
    check(string3.value() == "");
    string3.set_value("def");
}

pub fn check_map(kvstore: ScImmutableMap) {
    let int1 = kvstore.get_int(VAR_INT1);
    check(int1.value() == 1);

    let string1 = kvstore.get_string(VAR_STRING1);
    check(string1.value() == "a");

    let ia1 = kvstore.get_int_array(VAR_INT_ARRAY1);
    let int2 = ia1.get_int(0);
    check(int2.value() == 2);
    let int3 = ia1.get_int(1);
    check(int3.value() == 3);

    let sa1 = kvstore.get_string_array(VAR_STRING_ARRAY1);
    let string2 = sa1.get_string(0);
    check(string2.value() == "bc");
    let string3 = sa1.get_string(1);
    check(string3.value() == "def");
}

pub fn check(condition: bool) {
    if !condition {
        panic!("Check failed!")
    }
}
