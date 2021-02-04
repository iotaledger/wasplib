// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

#![allow(dead_code)]

use wasplib::client::*;

use super::*;

pub const SC_NAME: &str = "inccounter";
pub const SC_HNAME: ScHname = ScHname(0xaf2438e9);

pub const PARAM_COUNTER: &str = "counter";
pub const PARAM_NUM_REPEATS: &str = "numRepeats";

pub const VAR_COUNTER: &str = "counter";
pub const VAR_INT1: &str = "int1";
pub const VAR_INT_ARRAY1: &str = "intArray1";
pub const VAR_NUM_REPEATS: &str = "numRepeats";
pub const VAR_STRING1: &str = "string1";
pub const VAR_STRING_ARRAY1: &str = "stringArray1";

pub const FUNC_CALL_INCREMENT: &str = "callIncrement";
pub const FUNC_CALL_INCREMENT_RECURSE5X: &str = "callIncrementRecurse5x";
pub const FUNC_INCREMENT: &str = "increment";
pub const FUNC_INIT: &str = "init";
pub const FUNC_LOCAL_STATE_INTERNAL_CALL: &str = "localStateInternalCall";
pub const FUNC_LOCAL_STATE_POST: &str = "localStatePost";
pub const FUNC_LOCAL_STATE_SANDBOX_CALL: &str = "localStateSandboxCall";
pub const FUNC_POST_INCREMENT: &str = "postIncrement";
pub const FUNC_REPEAT_MANY: &str = "repeatMany";
pub const FUNC_RESULTS_TEST: &str = "resultsTest";
pub const FUNC_STATE_TEST: &str = "stateTest";
pub const FUNC_WHEN_MUST_INCREMENT: &str = "whenMustIncrement";
pub const VIEW_GET_COUNTER: &str = "getCounter";
pub const VIEW_RESULTS_CHECK: &str = "resultsCheck";
pub const VIEW_STATE_CHECK: &str = "stateCheck";

pub const HFUNC_CALL_INCREMENT: ScHname = ScHname(0xeb5dcacd);
pub const HFUNC_CALL_INCREMENT_RECURSE5X: ScHname = ScHname(0x8749fbff);
pub const HFUNC_INCREMENT: ScHname = ScHname(0xd351bd12);
pub const HFUNC_INIT: ScHname = ScHname(0x1f44d644);
pub const HFUNC_LOCAL_STATE_INTERNAL_CALL: ScHname = ScHname(0xecfc5d33);
pub const HFUNC_LOCAL_STATE_POST: ScHname = ScHname(0x3fd54d13);
pub const HFUNC_LOCAL_STATE_SANDBOX_CALL: ScHname = ScHname(0x7bd22c53);
pub const HFUNC_POST_INCREMENT: ScHname = ScHname(0x81c772f5);
pub const HFUNC_REPEAT_MANY: ScHname = ScHname(0x4ff450d3);
pub const HFUNC_RESULTS_TEST: ScHname = ScHname(0xd0544634);
pub const HFUNC_STATE_TEST: ScHname = ScHname(0x41830d59);
pub const HFUNC_WHEN_MUST_INCREMENT: ScHname = ScHname(0xb4c3e7a6);
pub const HVIEW_GET_COUNTER: ScHname = ScHname(0xb423e607);
pub const HVIEW_RESULTS_CHECK: ScHname = ScHname(0xa39ac571);
pub const HVIEW_STATE_CHECK: ScHname = ScHname(0xaafeb10a);

#[no_mangle]
fn on_load() {
    let exports = ScExports::new();
    exports.add_call(FUNC_CALL_INCREMENT, func_call_increment);
    exports.add_call(FUNC_CALL_INCREMENT_RECURSE5X, func_call_increment_recurse5x);
    exports.add_call(FUNC_INCREMENT, func_increment);
    exports.add_call(FUNC_INIT, func_init);
    exports.add_call(FUNC_LOCAL_STATE_INTERNAL_CALL, func_local_state_internal_call);
    exports.add_call(FUNC_LOCAL_STATE_POST, func_local_state_post);
    exports.add_call(FUNC_LOCAL_STATE_SANDBOX_CALL, func_local_state_sandbox_call);
    exports.add_call(FUNC_POST_INCREMENT, func_post_increment);
    exports.add_call(FUNC_REPEAT_MANY, func_repeat_many);
    exports.add_call(FUNC_RESULTS_TEST, func_results_test);
    exports.add_call(FUNC_STATE_TEST, func_state_test);
    exports.add_call(FUNC_WHEN_MUST_INCREMENT, func_when_must_increment);
    exports.add_view(VIEW_GET_COUNTER, view_get_counter);
    exports.add_view(VIEW_RESULTS_CHECK, view_results_check);
    exports.add_view(VIEW_STATE_CHECK, view_state_check);
}
