// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

//@formatter:off

#![allow(dead_code)]

#![allow(unused_imports)]

use testcore::*;
use wasmlib::*;
use wasmlib::host::*;

use crate::consts::*;
use crate::keys::*;
use crate::params::*;
use crate::results::*;
use crate::state::*;

mod consts;
mod keys;
mod params;
mod results;
mod state;
mod testcore;

#[no_mangle]
fn on_load() {
    let exports = ScExports::new();
    exports.add_func(FUNC_CALL_ON_CHAIN, func_call_on_chain_thunk);
    exports.add_func(FUNC_CHECK_CONTEXT_FROM_FULL_EP, func_check_context_from_full_ep_thunk);
    exports.add_func(FUNC_DO_NOTHING, func_do_nothing_thunk);
    exports.add_func(FUNC_GET_MINTED_SUPPLY, func_get_minted_supply_thunk);
    exports.add_func(FUNC_INC_COUNTER, func_inc_counter_thunk);
    exports.add_func(FUNC_INIT, func_init_thunk);
    exports.add_func(FUNC_PASS_TYPES_FULL, func_pass_types_full_thunk);
    exports.add_func(FUNC_RUN_RECURSION, func_run_recursion_thunk);
    exports.add_func(FUNC_SEND_TO_ADDRESS, func_send_to_address_thunk);
    exports.add_func(FUNC_SET_INT, func_set_int_thunk);
    exports.add_func(FUNC_TEST_CALL_PANIC_FULL_EP, func_test_call_panic_full_ep_thunk);
    exports.add_func(FUNC_TEST_CALL_PANIC_VIEW_EP_FROM_FULL, func_test_call_panic_view_ep_from_full_thunk);
    exports.add_func(FUNC_TEST_CHAIN_OWNER_ID_FULL, func_test_chain_owner_id_full_thunk);
    exports.add_func(FUNC_TEST_EVENT_LOG_DEPLOY, func_test_event_log_deploy_thunk);
    exports.add_func(FUNC_TEST_EVENT_LOG_EVENT_DATA, func_test_event_log_event_data_thunk);
    exports.add_func(FUNC_TEST_EVENT_LOG_GENERIC_DATA, func_test_event_log_generic_data_thunk);
    exports.add_func(FUNC_TEST_PANIC_FULL_EP, func_test_panic_full_ep_thunk);
    exports.add_func(FUNC_WITHDRAW_TO_CHAIN, func_withdraw_to_chain_thunk);
    exports.add_view(VIEW_CHECK_CONTEXT_FROM_VIEW_EP, view_check_context_from_view_ep_thunk);
    exports.add_view(VIEW_FIBONACCI, view_fibonacci_thunk);
    exports.add_view(VIEW_GET_COUNTER, view_get_counter_thunk);
    exports.add_view(VIEW_GET_INT, view_get_int_thunk);
    exports.add_view(VIEW_JUST_VIEW, view_just_view_thunk);
    exports.add_view(VIEW_PASS_TYPES_VIEW, view_pass_types_view_thunk);
    exports.add_view(VIEW_TEST_CALL_PANIC_VIEW_EP_FROM_VIEW, view_test_call_panic_view_ep_from_view_thunk);
    exports.add_view(VIEW_TEST_CHAIN_OWNER_ID_VIEW, view_test_chain_owner_id_view_thunk);
    exports.add_view(VIEW_TEST_PANIC_VIEW_EP, view_test_panic_view_ep_thunk);
    exports.add_view(VIEW_TEST_SANDBOX_CALL, view_test_sandbox_call_thunk);

    unsafe {
        for i in 0..KEY_MAP_LEN {
            IDX_MAP[i] = get_key_id_from_string(KEY_MAP[i]);
        }
    }
}

pub struct FuncCallOnChainContext {
    params:  ImmutableFuncCallOnChainParams,
    results: MutableFuncCallOnChainResults,
    state:   MutableTestCoreState,
}

fn func_call_on_chain_thunk(ctx: &ScFuncContext) {
    ctx.log("testcore.funcCallOnChain");
    let f = FuncCallOnChainContext {
        params: ImmutableFuncCallOnChainParams {
            id: get_object_id(1, KEY_PARAMS, TYPE_MAP),
        },
        results: MutableFuncCallOnChainResults {
            id: get_object_id(1, KEY_RESULTS, TYPE_MAP),
        },
        state: MutableTestCoreState {
            id: get_object_id(1, KEY_STATE, TYPE_MAP),
        },
    };
    ctx.require(f.params.int_value().exists(), "missing mandatory intValue");
    func_call_on_chain(ctx, &f);
    ctx.log("testcore.funcCallOnChain ok");
}

pub struct FuncCheckContextFromFullEPContext {
    params: ImmutableFuncCheckContextFromFullEPParams,
    state:  MutableTestCoreState,
}

fn func_check_context_from_full_ep_thunk(ctx: &ScFuncContext) {
    ctx.log("testcore.funcCheckContextFromFullEP");
    let f = FuncCheckContextFromFullEPContext {
        params: ImmutableFuncCheckContextFromFullEPParams {
            id: get_object_id(1, KEY_PARAMS, TYPE_MAP),
        },
        state: MutableTestCoreState {
            id: get_object_id(1, KEY_STATE, TYPE_MAP),
        },
    };
    ctx.require(f.params.agent_id().exists(), "missing mandatory agentId");
    ctx.require(f.params.caller().exists(), "missing mandatory caller");
    ctx.require(f.params.chain_id().exists(), "missing mandatory chainId");
    ctx.require(f.params.chain_owner_id().exists(), "missing mandatory chainOwnerId");
    ctx.require(f.params.contract_creator().exists(), "missing mandatory contractCreator");
    func_check_context_from_full_ep(ctx, &f);
    ctx.log("testcore.funcCheckContextFromFullEP ok");
}

pub struct FuncDoNothingContext {
    state: MutableTestCoreState,
}

fn func_do_nothing_thunk(ctx: &ScFuncContext) {
    ctx.log("testcore.funcDoNothing");
    let f = FuncDoNothingContext {
        state: MutableTestCoreState {
            id: get_object_id(1, KEY_STATE, TYPE_MAP),
        },
    };
    func_do_nothing(ctx, &f);
    ctx.log("testcore.funcDoNothing ok");
}

pub struct FuncGetMintedSupplyContext {
    results: MutableFuncGetMintedSupplyResults,
    state:   MutableTestCoreState,
}

fn func_get_minted_supply_thunk(ctx: &ScFuncContext) {
    ctx.log("testcore.funcGetMintedSupply");
    let f = FuncGetMintedSupplyContext {
        results: MutableFuncGetMintedSupplyResults {
            id: get_object_id(1, KEY_RESULTS, TYPE_MAP),
        },
        state: MutableTestCoreState {
            id: get_object_id(1, KEY_STATE, TYPE_MAP),
        },
    };
    func_get_minted_supply(ctx, &f);
    ctx.log("testcore.funcGetMintedSupply ok");
}

pub struct FuncIncCounterContext {
    state: MutableTestCoreState,
}

fn func_inc_counter_thunk(ctx: &ScFuncContext) {
    ctx.log("testcore.funcIncCounter");
    let f = FuncIncCounterContext {
        state: MutableTestCoreState {
            id: get_object_id(1, KEY_STATE, TYPE_MAP),
        },
    };
    func_inc_counter(ctx, &f);
    ctx.log("testcore.funcIncCounter ok");
}

pub struct FuncInitContext {
    state: MutableTestCoreState,
}

fn func_init_thunk(ctx: &ScFuncContext) {
    ctx.log("testcore.funcInit");
    let f = FuncInitContext {
        state: MutableTestCoreState {
            id: get_object_id(1, KEY_STATE, TYPE_MAP),
        },
    };
    func_init(ctx, &f);
    ctx.log("testcore.funcInit ok");
}

pub struct FuncPassTypesFullContext {
    params: ImmutableFuncPassTypesFullParams,
    state:  MutableTestCoreState,
}

fn func_pass_types_full_thunk(ctx: &ScFuncContext) {
    ctx.log("testcore.funcPassTypesFull");
    let f = FuncPassTypesFullContext {
        params: ImmutableFuncPassTypesFullParams {
            id: get_object_id(1, KEY_PARAMS, TYPE_MAP),
        },
        state: MutableTestCoreState {
            id: get_object_id(1, KEY_STATE, TYPE_MAP),
        },
    };
    ctx.require(f.params.hash().exists(), "missing mandatory hash");
    ctx.require(f.params.hname().exists(), "missing mandatory hname");
    ctx.require(f.params.hname_zero().exists(), "missing mandatory hnameZero");
    ctx.require(f.params.int64().exists(), "missing mandatory int64");
    ctx.require(f.params.int64_zero().exists(), "missing mandatory int64Zero");
    ctx.require(f.params.string().exists(), "missing mandatory string");
    ctx.require(f.params.string_zero().exists(), "missing mandatory stringZero");
    func_pass_types_full(ctx, &f);
    ctx.log("testcore.funcPassTypesFull ok");
}

pub struct FuncRunRecursionContext {
    params:  ImmutableFuncRunRecursionParams,
    results: MutableFuncRunRecursionResults,
    state:   MutableTestCoreState,
}

fn func_run_recursion_thunk(ctx: &ScFuncContext) {
    ctx.log("testcore.funcRunRecursion");
    let f = FuncRunRecursionContext {
        params: ImmutableFuncRunRecursionParams {
            id: get_object_id(1, KEY_PARAMS, TYPE_MAP),
        },
        results: MutableFuncRunRecursionResults {
            id: get_object_id(1, KEY_RESULTS, TYPE_MAP),
        },
        state: MutableTestCoreState {
            id: get_object_id(1, KEY_STATE, TYPE_MAP),
        },
    };
    ctx.require(f.params.int_value().exists(), "missing mandatory intValue");
    func_run_recursion(ctx, &f);
    ctx.log("testcore.funcRunRecursion ok");
}

pub struct FuncSendToAddressContext {
    params: ImmutableFuncSendToAddressParams,
    state:  MutableTestCoreState,
}

fn func_send_to_address_thunk(ctx: &ScFuncContext) {
    ctx.log("testcore.funcSendToAddress");
    ctx.require(ctx.caller() == ctx.contract_creator(), "no permission");

    let f = FuncSendToAddressContext {
        params: ImmutableFuncSendToAddressParams {
            id: get_object_id(1, KEY_PARAMS, TYPE_MAP),
        },
        state: MutableTestCoreState {
            id: get_object_id(1, KEY_STATE, TYPE_MAP),
        },
    };
    ctx.require(f.params.address().exists(), "missing mandatory address");
    func_send_to_address(ctx, &f);
    ctx.log("testcore.funcSendToAddress ok");
}

pub struct FuncSetIntContext {
    params: ImmutableFuncSetIntParams,
    state:  MutableTestCoreState,
}

fn func_set_int_thunk(ctx: &ScFuncContext) {
    ctx.log("testcore.funcSetInt");
    let f = FuncSetIntContext {
        params: ImmutableFuncSetIntParams {
            id: get_object_id(1, KEY_PARAMS, TYPE_MAP),
        },
        state: MutableTestCoreState {
            id: get_object_id(1, KEY_STATE, TYPE_MAP),
        },
    };
    ctx.require(f.params.int_value().exists(), "missing mandatory intValue");
    ctx.require(f.params.name().exists(), "missing mandatory name");
    func_set_int(ctx, &f);
    ctx.log("testcore.funcSetInt ok");
}

pub struct FuncTestCallPanicFullEPContext {
    state: MutableTestCoreState,
}

fn func_test_call_panic_full_ep_thunk(ctx: &ScFuncContext) {
    ctx.log("testcore.funcTestCallPanicFullEP");
    let f = FuncTestCallPanicFullEPContext {
        state: MutableTestCoreState {
            id: get_object_id(1, KEY_STATE, TYPE_MAP),
        },
    };
    func_test_call_panic_full_ep(ctx, &f);
    ctx.log("testcore.funcTestCallPanicFullEP ok");
}

pub struct FuncTestCallPanicViewEPFromFullContext {
    state: MutableTestCoreState,
}

fn func_test_call_panic_view_ep_from_full_thunk(ctx: &ScFuncContext) {
    ctx.log("testcore.funcTestCallPanicViewEPFromFull");
    let f = FuncTestCallPanicViewEPFromFullContext {
        state: MutableTestCoreState {
            id: get_object_id(1, KEY_STATE, TYPE_MAP),
        },
    };
    func_test_call_panic_view_ep_from_full(ctx, &f);
    ctx.log("testcore.funcTestCallPanicViewEPFromFull ok");
}

pub struct FuncTestChainOwnerIDFullContext {
    results: MutableFuncTestChainOwnerIDFullResults,
    state:   MutableTestCoreState,
}

fn func_test_chain_owner_id_full_thunk(ctx: &ScFuncContext) {
    ctx.log("testcore.funcTestChainOwnerIDFull");
    let f = FuncTestChainOwnerIDFullContext {
        results: MutableFuncTestChainOwnerIDFullResults {
            id: get_object_id(1, KEY_RESULTS, TYPE_MAP),
        },
        state: MutableTestCoreState {
            id: get_object_id(1, KEY_STATE, TYPE_MAP),
        },
    };
    func_test_chain_owner_id_full(ctx, &f);
    ctx.log("testcore.funcTestChainOwnerIDFull ok");
}

pub struct FuncTestEventLogDeployContext {
    state: MutableTestCoreState,
}

fn func_test_event_log_deploy_thunk(ctx: &ScFuncContext) {
    ctx.log("testcore.funcTestEventLogDeploy");
    let f = FuncTestEventLogDeployContext {
        state: MutableTestCoreState {
            id: get_object_id(1, KEY_STATE, TYPE_MAP),
        },
    };
    func_test_event_log_deploy(ctx, &f);
    ctx.log("testcore.funcTestEventLogDeploy ok");
}

pub struct FuncTestEventLogEventDataContext {
    state: MutableTestCoreState,
}

fn func_test_event_log_event_data_thunk(ctx: &ScFuncContext) {
    ctx.log("testcore.funcTestEventLogEventData");
    let f = FuncTestEventLogEventDataContext {
        state: MutableTestCoreState {
            id: get_object_id(1, KEY_STATE, TYPE_MAP),
        },
    };
    func_test_event_log_event_data(ctx, &f);
    ctx.log("testcore.funcTestEventLogEventData ok");
}

pub struct FuncTestEventLogGenericDataContext {
    params: ImmutableFuncTestEventLogGenericDataParams,
    state:  MutableTestCoreState,
}

fn func_test_event_log_generic_data_thunk(ctx: &ScFuncContext) {
    ctx.log("testcore.funcTestEventLogGenericData");
    let f = FuncTestEventLogGenericDataContext {
        params: ImmutableFuncTestEventLogGenericDataParams {
            id: get_object_id(1, KEY_PARAMS, TYPE_MAP),
        },
        state: MutableTestCoreState {
            id: get_object_id(1, KEY_STATE, TYPE_MAP),
        },
    };
    ctx.require(f.params.counter().exists(), "missing mandatory counter");
    func_test_event_log_generic_data(ctx, &f);
    ctx.log("testcore.funcTestEventLogGenericData ok");
}

pub struct FuncTestPanicFullEPContext {
    state: MutableTestCoreState,
}

fn func_test_panic_full_ep_thunk(ctx: &ScFuncContext) {
    ctx.log("testcore.funcTestPanicFullEP");
    let f = FuncTestPanicFullEPContext {
        state: MutableTestCoreState {
            id: get_object_id(1, KEY_STATE, TYPE_MAP),
        },
    };
    func_test_panic_full_ep(ctx, &f);
    ctx.log("testcore.funcTestPanicFullEP ok");
}

pub struct FuncWithdrawToChainContext {
    params: ImmutableFuncWithdrawToChainParams,
    state:  MutableTestCoreState,
}

fn func_withdraw_to_chain_thunk(ctx: &ScFuncContext) {
    ctx.log("testcore.funcWithdrawToChain");
    let f = FuncWithdrawToChainContext {
        params: ImmutableFuncWithdrawToChainParams {
            id: get_object_id(1, KEY_PARAMS, TYPE_MAP),
        },
        state: MutableTestCoreState {
            id: get_object_id(1, KEY_STATE, TYPE_MAP),
        },
    };
    ctx.require(f.params.chain_id().exists(), "missing mandatory chainId");
    func_withdraw_to_chain(ctx, &f);
    ctx.log("testcore.funcWithdrawToChain ok");
}

pub struct ViewCheckContextFromViewEPContext {
    params: ImmutableViewCheckContextFromViewEPParams,
    state:  ImmutableTestCoreState,
}

fn view_check_context_from_view_ep_thunk(ctx: &ScViewContext) {
    ctx.log("testcore.viewCheckContextFromViewEP");
    let f = ViewCheckContextFromViewEPContext {
        params: ImmutableViewCheckContextFromViewEPParams {
            id: get_object_id(1, KEY_PARAMS, TYPE_MAP),
        },
        state: ImmutableTestCoreState {
            id: get_object_id(1, KEY_STATE, TYPE_MAP),
        },
    };
    ctx.require(f.params.agent_id().exists(), "missing mandatory agentId");
    ctx.require(f.params.chain_id().exists(), "missing mandatory chainId");
    ctx.require(f.params.chain_owner_id().exists(), "missing mandatory chainOwnerId");
    ctx.require(f.params.contract_creator().exists(), "missing mandatory contractCreator");
    view_check_context_from_view_ep(ctx, &f);
    ctx.log("testcore.viewCheckContextFromViewEP ok");
}

pub struct ViewFibonacciContext {
    params:  ImmutableViewFibonacciParams,
    results: MutableViewFibonacciResults,
    state:   ImmutableTestCoreState,
}

fn view_fibonacci_thunk(ctx: &ScViewContext) {
    ctx.log("testcore.viewFibonacci");
    let f = ViewFibonacciContext {
        params: ImmutableViewFibonacciParams {
            id: get_object_id(1, KEY_PARAMS, TYPE_MAP),
        },
        results: MutableViewFibonacciResults {
            id: get_object_id(1, KEY_RESULTS, TYPE_MAP),
        },
        state: ImmutableTestCoreState {
            id: get_object_id(1, KEY_STATE, TYPE_MAP),
        },
    };
    ctx.require(f.params.int_value().exists(), "missing mandatory intValue");
    view_fibonacci(ctx, &f);
    ctx.log("testcore.viewFibonacci ok");
}

pub struct ViewGetCounterContext {
    results: MutableViewGetCounterResults,
    state:   ImmutableTestCoreState,
}

fn view_get_counter_thunk(ctx: &ScViewContext) {
    ctx.log("testcore.viewGetCounter");
    let f = ViewGetCounterContext {
        results: MutableViewGetCounterResults {
            id: get_object_id(1, KEY_RESULTS, TYPE_MAP),
        },
        state: ImmutableTestCoreState {
            id: get_object_id(1, KEY_STATE, TYPE_MAP),
        },
    };
    view_get_counter(ctx, &f);
    ctx.log("testcore.viewGetCounter ok");
}

pub struct ViewGetIntContext {
    params: ImmutableViewGetIntParams,
    state:  ImmutableTestCoreState,
}

fn view_get_int_thunk(ctx: &ScViewContext) {
    ctx.log("testcore.viewGetInt");
    let f = ViewGetIntContext {
        params: ImmutableViewGetIntParams {
            id: get_object_id(1, KEY_PARAMS, TYPE_MAP),
        },
        state: ImmutableTestCoreState {
            id: get_object_id(1, KEY_STATE, TYPE_MAP),
        },
    };
    ctx.require(f.params.name().exists(), "missing mandatory name");
    view_get_int(ctx, &f);
    ctx.log("testcore.viewGetInt ok");
}

pub struct ViewJustViewContext {
    state: ImmutableTestCoreState,
}

fn view_just_view_thunk(ctx: &ScViewContext) {
    ctx.log("testcore.viewJustView");
    let f = ViewJustViewContext {
        state: ImmutableTestCoreState {
            id: get_object_id(1, KEY_STATE, TYPE_MAP),
        },
    };
    view_just_view(ctx, &f);
    ctx.log("testcore.viewJustView ok");
}

pub struct ViewPassTypesViewContext {
    params: ImmutableViewPassTypesViewParams,
    state:  ImmutableTestCoreState,
}

fn view_pass_types_view_thunk(ctx: &ScViewContext) {
    ctx.log("testcore.viewPassTypesView");
    let f = ViewPassTypesViewContext {
        params: ImmutableViewPassTypesViewParams {
            id: get_object_id(1, KEY_PARAMS, TYPE_MAP),
        },
        state: ImmutableTestCoreState {
            id: get_object_id(1, KEY_STATE, TYPE_MAP),
        },
    };
    ctx.require(f.params.hash().exists(), "missing mandatory hash");
    ctx.require(f.params.hname().exists(), "missing mandatory hname");
    ctx.require(f.params.hname_zero().exists(), "missing mandatory hnameZero");
    ctx.require(f.params.int64().exists(), "missing mandatory int64");
    ctx.require(f.params.int64_zero().exists(), "missing mandatory int64Zero");
    ctx.require(f.params.string().exists(), "missing mandatory string");
    ctx.require(f.params.string_zero().exists(), "missing mandatory stringZero");
    view_pass_types_view(ctx, &f);
    ctx.log("testcore.viewPassTypesView ok");
}

pub struct ViewTestCallPanicViewEPFromViewContext {
    state: ImmutableTestCoreState,
}

fn view_test_call_panic_view_ep_from_view_thunk(ctx: &ScViewContext) {
    ctx.log("testcore.viewTestCallPanicViewEPFromView");
    let f = ViewTestCallPanicViewEPFromViewContext {
        state: ImmutableTestCoreState {
            id: get_object_id(1, KEY_STATE, TYPE_MAP),
        },
    };
    view_test_call_panic_view_ep_from_view(ctx, &f);
    ctx.log("testcore.viewTestCallPanicViewEPFromView ok");
}

pub struct ViewTestChainOwnerIDViewContext {
    results: MutableViewTestChainOwnerIDViewResults,
    state:   ImmutableTestCoreState,
}

fn view_test_chain_owner_id_view_thunk(ctx: &ScViewContext) {
    ctx.log("testcore.viewTestChainOwnerIDView");
    let f = ViewTestChainOwnerIDViewContext {
        results: MutableViewTestChainOwnerIDViewResults {
            id: get_object_id(1, KEY_RESULTS, TYPE_MAP),
        },
        state: ImmutableTestCoreState {
            id: get_object_id(1, KEY_STATE, TYPE_MAP),
        },
    };
    view_test_chain_owner_id_view(ctx, &f);
    ctx.log("testcore.viewTestChainOwnerIDView ok");
}

pub struct ViewTestPanicViewEPContext {
    state: ImmutableTestCoreState,
}

fn view_test_panic_view_ep_thunk(ctx: &ScViewContext) {
    ctx.log("testcore.viewTestPanicViewEP");
    let f = ViewTestPanicViewEPContext {
        state: ImmutableTestCoreState {
            id: get_object_id(1, KEY_STATE, TYPE_MAP),
        },
    };
    view_test_panic_view_ep(ctx, &f);
    ctx.log("testcore.viewTestPanicViewEP ok");
}

pub struct ViewTestSandboxCallContext {
    results: MutableViewTestSandboxCallResults,
    state:   ImmutableTestCoreState,
}

fn view_test_sandbox_call_thunk(ctx: &ScViewContext) {
    ctx.log("testcore.viewTestSandboxCall");
    let f = ViewTestSandboxCallContext {
        results: MutableViewTestSandboxCallResults {
            id: get_object_id(1, KEY_RESULTS, TYPE_MAP),
        },
        state: ImmutableTestCoreState {
            id: get_object_id(1, KEY_STATE, TYPE_MAP),
        },
    };
    view_test_sandbox_call(ctx, &f);
    ctx.log("testcore.viewTestSandboxCall ok");
}

//@formatter:on
