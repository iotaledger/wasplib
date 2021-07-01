// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
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
mod contract;
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

pub struct CallOnChainContext {
    params:  ImmutableCallOnChainParams,
    results: MutableCallOnChainResults,
    state:   MutableTestCoreState,
}

fn func_call_on_chain_thunk(ctx: &ScFuncContext) {
    ctx.log("testcore.funcCallOnChain");
    let f = CallOnChainContext {
        params: ImmutableCallOnChainParams {
            id: OBJ_ID_PARAMS,
        },
        results: MutableCallOnChainResults {
            id: OBJ_ID_RESULTS,
        },
        state: MutableTestCoreState {
            id: OBJ_ID_STATE,
        },
    };
    ctx.require(f.params.int_value().exists(), "missing mandatory intValue");
    func_call_on_chain(ctx, &f);
    ctx.log("testcore.funcCallOnChain ok");
}

pub struct CheckContextFromFullEPContext {
    params: ImmutableCheckContextFromFullEPParams,
    state:  MutableTestCoreState,
}

fn func_check_context_from_full_ep_thunk(ctx: &ScFuncContext) {
    ctx.log("testcore.funcCheckContextFromFullEP");
    let f = CheckContextFromFullEPContext {
        params: ImmutableCheckContextFromFullEPParams {
            id: OBJ_ID_PARAMS,
        },
        state: MutableTestCoreState {
            id: OBJ_ID_STATE,
        },
    };
    ctx.require(f.params.agent_id().exists(), "missing mandatory agentID");
    ctx.require(f.params.caller().exists(), "missing mandatory caller");
    ctx.require(f.params.chain_id().exists(), "missing mandatory chainID");
    ctx.require(f.params.chain_owner_id().exists(), "missing mandatory chainOwnerID");
    ctx.require(f.params.contract_creator().exists(), "missing mandatory contractCreator");
    func_check_context_from_full_ep(ctx, &f);
    ctx.log("testcore.funcCheckContextFromFullEP ok");
}

pub struct DoNothingContext {
    state: MutableTestCoreState,
}

fn func_do_nothing_thunk(ctx: &ScFuncContext) {
    ctx.log("testcore.funcDoNothing");
    let f = DoNothingContext {
        state: MutableTestCoreState {
            id: OBJ_ID_STATE,
        },
    };
    func_do_nothing(ctx, &f);
    ctx.log("testcore.funcDoNothing ok");
}

pub struct GetMintedSupplyContext {
    results: MutableGetMintedSupplyResults,
    state:   MutableTestCoreState,
}

fn func_get_minted_supply_thunk(ctx: &ScFuncContext) {
    ctx.log("testcore.funcGetMintedSupply");
    let f = GetMintedSupplyContext {
        results: MutableGetMintedSupplyResults {
            id: OBJ_ID_RESULTS,
        },
        state: MutableTestCoreState {
            id: OBJ_ID_STATE,
        },
    };
    func_get_minted_supply(ctx, &f);
    ctx.log("testcore.funcGetMintedSupply ok");
}

pub struct IncCounterContext {
    state: MutableTestCoreState,
}

fn func_inc_counter_thunk(ctx: &ScFuncContext) {
    ctx.log("testcore.funcIncCounter");
    let f = IncCounterContext {
        state: MutableTestCoreState {
            id: OBJ_ID_STATE,
        },
    };
    func_inc_counter(ctx, &f);
    ctx.log("testcore.funcIncCounter ok");
}

pub struct InitContext {
    state: MutableTestCoreState,
}

fn func_init_thunk(ctx: &ScFuncContext) {
    ctx.log("testcore.funcInit");
    let f = InitContext {
        state: MutableTestCoreState {
            id: OBJ_ID_STATE,
        },
    };
    func_init(ctx, &f);
    ctx.log("testcore.funcInit ok");
}

pub struct PassTypesFullContext {
    params: ImmutablePassTypesFullParams,
    state:  MutableTestCoreState,
}

fn func_pass_types_full_thunk(ctx: &ScFuncContext) {
    ctx.log("testcore.funcPassTypesFull");
    let f = PassTypesFullContext {
        params: ImmutablePassTypesFullParams {
            id: OBJ_ID_PARAMS,
        },
        state: MutableTestCoreState {
            id: OBJ_ID_STATE,
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

pub struct RunRecursionContext {
    params:  ImmutableRunRecursionParams,
    results: MutableRunRecursionResults,
    state:   MutableTestCoreState,
}

fn func_run_recursion_thunk(ctx: &ScFuncContext) {
    ctx.log("testcore.funcRunRecursion");
    let f = RunRecursionContext {
        params: ImmutableRunRecursionParams {
            id: OBJ_ID_PARAMS,
        },
        results: MutableRunRecursionResults {
            id: OBJ_ID_RESULTS,
        },
        state: MutableTestCoreState {
            id: OBJ_ID_STATE,
        },
    };
    ctx.require(f.params.int_value().exists(), "missing mandatory intValue");
    func_run_recursion(ctx, &f);
    ctx.log("testcore.funcRunRecursion ok");
}

pub struct SendToAddressContext {
    params: ImmutableSendToAddressParams,
    state:  MutableTestCoreState,
}

fn func_send_to_address_thunk(ctx: &ScFuncContext) {
    ctx.log("testcore.funcSendToAddress");
    ctx.require(ctx.caller() == ctx.contract_creator(), "no permission");

    let f = SendToAddressContext {
        params: ImmutableSendToAddressParams {
            id: OBJ_ID_PARAMS,
        },
        state: MutableTestCoreState {
            id: OBJ_ID_STATE,
        },
    };
    ctx.require(f.params.address().exists(), "missing mandatory address");
    func_send_to_address(ctx, &f);
    ctx.log("testcore.funcSendToAddress ok");
}

pub struct SetIntContext {
    params: ImmutableSetIntParams,
    state:  MutableTestCoreState,
}

fn func_set_int_thunk(ctx: &ScFuncContext) {
    ctx.log("testcore.funcSetInt");
    let f = SetIntContext {
        params: ImmutableSetIntParams {
            id: OBJ_ID_PARAMS,
        },
        state: MutableTestCoreState {
            id: OBJ_ID_STATE,
        },
    };
    ctx.require(f.params.int_value().exists(), "missing mandatory intValue");
    ctx.require(f.params.name().exists(), "missing mandatory name");
    func_set_int(ctx, &f);
    ctx.log("testcore.funcSetInt ok");
}

pub struct TestCallPanicFullEPContext {
    state: MutableTestCoreState,
}

fn func_test_call_panic_full_ep_thunk(ctx: &ScFuncContext) {
    ctx.log("testcore.funcTestCallPanicFullEP");
    let f = TestCallPanicFullEPContext {
        state: MutableTestCoreState {
            id: OBJ_ID_STATE,
        },
    };
    func_test_call_panic_full_ep(ctx, &f);
    ctx.log("testcore.funcTestCallPanicFullEP ok");
}

pub struct TestCallPanicViewEPFromFullContext {
    state: MutableTestCoreState,
}

fn func_test_call_panic_view_ep_from_full_thunk(ctx: &ScFuncContext) {
    ctx.log("testcore.funcTestCallPanicViewEPFromFull");
    let f = TestCallPanicViewEPFromFullContext {
        state: MutableTestCoreState {
            id: OBJ_ID_STATE,
        },
    };
    func_test_call_panic_view_ep_from_full(ctx, &f);
    ctx.log("testcore.funcTestCallPanicViewEPFromFull ok");
}

pub struct TestChainOwnerIDFullContext {
    results: MutableTestChainOwnerIDFullResults,
    state:   MutableTestCoreState,
}

fn func_test_chain_owner_id_full_thunk(ctx: &ScFuncContext) {
    ctx.log("testcore.funcTestChainOwnerIDFull");
    let f = TestChainOwnerIDFullContext {
        results: MutableTestChainOwnerIDFullResults {
            id: OBJ_ID_RESULTS,
        },
        state: MutableTestCoreState {
            id: OBJ_ID_STATE,
        },
    };
    func_test_chain_owner_id_full(ctx, &f);
    ctx.log("testcore.funcTestChainOwnerIDFull ok");
}

pub struct TestEventLogDeployContext {
    state: MutableTestCoreState,
}

fn func_test_event_log_deploy_thunk(ctx: &ScFuncContext) {
    ctx.log("testcore.funcTestEventLogDeploy");
    let f = TestEventLogDeployContext {
        state: MutableTestCoreState {
            id: OBJ_ID_STATE,
        },
    };
    func_test_event_log_deploy(ctx, &f);
    ctx.log("testcore.funcTestEventLogDeploy ok");
}

pub struct TestEventLogEventDataContext {
    state: MutableTestCoreState,
}

fn func_test_event_log_event_data_thunk(ctx: &ScFuncContext) {
    ctx.log("testcore.funcTestEventLogEventData");
    let f = TestEventLogEventDataContext {
        state: MutableTestCoreState {
            id: OBJ_ID_STATE,
        },
    };
    func_test_event_log_event_data(ctx, &f);
    ctx.log("testcore.funcTestEventLogEventData ok");
}

pub struct TestEventLogGenericDataContext {
    params: ImmutableTestEventLogGenericDataParams,
    state:  MutableTestCoreState,
}

fn func_test_event_log_generic_data_thunk(ctx: &ScFuncContext) {
    ctx.log("testcore.funcTestEventLogGenericData");
    let f = TestEventLogGenericDataContext {
        params: ImmutableTestEventLogGenericDataParams {
            id: OBJ_ID_PARAMS,
        },
        state: MutableTestCoreState {
            id: OBJ_ID_STATE,
        },
    };
    ctx.require(f.params.counter().exists(), "missing mandatory counter");
    func_test_event_log_generic_data(ctx, &f);
    ctx.log("testcore.funcTestEventLogGenericData ok");
}

pub struct TestPanicFullEPContext {
    state: MutableTestCoreState,
}

fn func_test_panic_full_ep_thunk(ctx: &ScFuncContext) {
    ctx.log("testcore.funcTestPanicFullEP");
    let f = TestPanicFullEPContext {
        state: MutableTestCoreState {
            id: OBJ_ID_STATE,
        },
    };
    func_test_panic_full_ep(ctx, &f);
    ctx.log("testcore.funcTestPanicFullEP ok");
}

pub struct WithdrawToChainContext {
    params: ImmutableWithdrawToChainParams,
    state:  MutableTestCoreState,
}

fn func_withdraw_to_chain_thunk(ctx: &ScFuncContext) {
    ctx.log("testcore.funcWithdrawToChain");
    let f = WithdrawToChainContext {
        params: ImmutableWithdrawToChainParams {
            id: OBJ_ID_PARAMS,
        },
        state: MutableTestCoreState {
            id: OBJ_ID_STATE,
        },
    };
    ctx.require(f.params.chain_id().exists(), "missing mandatory chainID");
    func_withdraw_to_chain(ctx, &f);
    ctx.log("testcore.funcWithdrawToChain ok");
}

pub struct CheckContextFromViewEPContext {
    params: ImmutableCheckContextFromViewEPParams,
    state:  ImmutableTestCoreState,
}

fn view_check_context_from_view_ep_thunk(ctx: &ScViewContext) {
    ctx.log("testcore.viewCheckContextFromViewEP");
    let f = CheckContextFromViewEPContext {
        params: ImmutableCheckContextFromViewEPParams {
            id: OBJ_ID_PARAMS,
        },
        state: ImmutableTestCoreState {
            id: OBJ_ID_STATE,
        },
    };
    ctx.require(f.params.agent_id().exists(), "missing mandatory agentID");
    ctx.require(f.params.chain_id().exists(), "missing mandatory chainID");
    ctx.require(f.params.chain_owner_id().exists(), "missing mandatory chainOwnerID");
    ctx.require(f.params.contract_creator().exists(), "missing mandatory contractCreator");
    view_check_context_from_view_ep(ctx, &f);
    ctx.log("testcore.viewCheckContextFromViewEP ok");
}

pub struct FibonacciContext {
    params:  ImmutableFibonacciParams,
    results: MutableFibonacciResults,
    state:   ImmutableTestCoreState,
}

fn view_fibonacci_thunk(ctx: &ScViewContext) {
    ctx.log("testcore.viewFibonacci");
    let f = FibonacciContext {
        params: ImmutableFibonacciParams {
            id: OBJ_ID_PARAMS,
        },
        results: MutableFibonacciResults {
            id: OBJ_ID_RESULTS,
        },
        state: ImmutableTestCoreState {
            id: OBJ_ID_STATE,
        },
    };
    ctx.require(f.params.int_value().exists(), "missing mandatory intValue");
    view_fibonacci(ctx, &f);
    ctx.log("testcore.viewFibonacci ok");
}

pub struct GetCounterContext {
    results: MutableGetCounterResults,
    state:   ImmutableTestCoreState,
}

fn view_get_counter_thunk(ctx: &ScViewContext) {
    ctx.log("testcore.viewGetCounter");
    let f = GetCounterContext {
        results: MutableGetCounterResults {
            id: OBJ_ID_RESULTS,
        },
        state: ImmutableTestCoreState {
            id: OBJ_ID_STATE,
        },
    };
    view_get_counter(ctx, &f);
    ctx.log("testcore.viewGetCounter ok");
}

pub struct GetIntContext {
    params:  ImmutableGetIntParams,
    results: MutableGetIntResults,
    state:   ImmutableTestCoreState,
}

fn view_get_int_thunk(ctx: &ScViewContext) {
    ctx.log("testcore.viewGetInt");
    let f = GetIntContext {
        params: ImmutableGetIntParams {
            id: OBJ_ID_PARAMS,
        },
        results: MutableGetIntResults {
            id: OBJ_ID_RESULTS,
        },
        state: ImmutableTestCoreState {
            id: OBJ_ID_STATE,
        },
    };
    ctx.require(f.params.name().exists(), "missing mandatory name");
    view_get_int(ctx, &f);
    ctx.log("testcore.viewGetInt ok");
}

pub struct JustViewContext {
    state: ImmutableTestCoreState,
}

fn view_just_view_thunk(ctx: &ScViewContext) {
    ctx.log("testcore.viewJustView");
    let f = JustViewContext {
        state: ImmutableTestCoreState {
            id: OBJ_ID_STATE,
        },
    };
    view_just_view(ctx, &f);
    ctx.log("testcore.viewJustView ok");
}

pub struct PassTypesViewContext {
    params: ImmutablePassTypesViewParams,
    state:  ImmutableTestCoreState,
}

fn view_pass_types_view_thunk(ctx: &ScViewContext) {
    ctx.log("testcore.viewPassTypesView");
    let f = PassTypesViewContext {
        params: ImmutablePassTypesViewParams {
            id: OBJ_ID_PARAMS,
        },
        state: ImmutableTestCoreState {
            id: OBJ_ID_STATE,
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

pub struct TestCallPanicViewEPFromViewContext {
    state: ImmutableTestCoreState,
}

fn view_test_call_panic_view_ep_from_view_thunk(ctx: &ScViewContext) {
    ctx.log("testcore.viewTestCallPanicViewEPFromView");
    let f = TestCallPanicViewEPFromViewContext {
        state: ImmutableTestCoreState {
            id: OBJ_ID_STATE,
        },
    };
    view_test_call_panic_view_ep_from_view(ctx, &f);
    ctx.log("testcore.viewTestCallPanicViewEPFromView ok");
}

pub struct TestChainOwnerIDViewContext {
    results: MutableTestChainOwnerIDViewResults,
    state:   ImmutableTestCoreState,
}

fn view_test_chain_owner_id_view_thunk(ctx: &ScViewContext) {
    ctx.log("testcore.viewTestChainOwnerIDView");
    let f = TestChainOwnerIDViewContext {
        results: MutableTestChainOwnerIDViewResults {
            id: OBJ_ID_RESULTS,
        },
        state: ImmutableTestCoreState {
            id: OBJ_ID_STATE,
        },
    };
    view_test_chain_owner_id_view(ctx, &f);
    ctx.log("testcore.viewTestChainOwnerIDView ok");
}

pub struct TestPanicViewEPContext {
    state: ImmutableTestCoreState,
}

fn view_test_panic_view_ep_thunk(ctx: &ScViewContext) {
    ctx.log("testcore.viewTestPanicViewEP");
    let f = TestPanicViewEPContext {
        state: ImmutableTestCoreState {
            id: OBJ_ID_STATE,
        },
    };
    view_test_panic_view_ep(ctx, &f);
    ctx.log("testcore.viewTestPanicViewEP ok");
}

pub struct TestSandboxCallContext {
    results: MutableTestSandboxCallResults,
    state:   ImmutableTestCoreState,
}

fn view_test_sandbox_call_thunk(ctx: &ScViewContext) {
    ctx.log("testcore.viewTestSandboxCall");
    let f = TestSandboxCallContext {
        results: MutableTestSandboxCallResults {
            id: OBJ_ID_RESULTS,
        },
        state: ImmutableTestCoreState {
            id: OBJ_ID_STATE,
        },
    };
    view_test_sandbox_call(ctx, &f);
    ctx.log("testcore.viewTestSandboxCall ok");
}

//@formatter:on
