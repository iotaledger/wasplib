// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

use consts::*;
use testcore::*;
use wasmlib::*;

mod consts;
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
    exports.add_func(FUNC_TEST_CONTRACT_ID_FULL, func_test_contract_id_full_thunk);
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
    exports.add_view(VIEW_TEST_CONTRACT_ID_VIEW, view_test_contract_id_view_thunk);
    exports.add_view(VIEW_TEST_PANIC_VIEW_EP, view_test_panic_view_ep_thunk);
    exports.add_view(VIEW_TEST_SANDBOX_CALL, view_test_sandbox_call_thunk);
}

//@formatter:off
pub struct FuncCallOnChainParams {
    pub hname_contract: ScImmutableHname,
    pub hname_ep:       ScImmutableHname,
    pub int_value:      ScImmutableInt64,
}
//@formatter:on

fn func_call_on_chain_thunk(ctx: &ScFuncContext) {
    let p = ctx.params();
    let params = FuncCallOnChainParams {
        hname_contract: p.get_hname(PARAM_HNAME_CONTRACT),
        hname_ep: p.get_hname(PARAM_HNAME_EP),
        int_value: p.get_int64(PARAM_INT_VALUE),
    };
    ctx.require(params.int_value.exists(), "missing mandatory intValue");
    ctx.log("testcore.funcCallOnChain");
    func_call_on_chain(ctx, &params);
    ctx.log("testcore.funcCallOnChain ok");
}

//@formatter:off
pub struct FuncCheckContextFromFullEPParams {
    pub agent_id:         ScImmutableAgentId,
    pub caller:           ScImmutableAgentId,
    pub chain_id:         ScImmutableChainId,
    pub chain_owner_id:   ScImmutableAgentId,
    pub contract_creator: ScImmutableAgentId,
    pub contract_id:      ScImmutableContractId,
}
//@formatter:on

fn func_check_context_from_full_ep_thunk(ctx: &ScFuncContext) {
    let p = ctx.params();
    let params = FuncCheckContextFromFullEPParams {
        agent_id: p.get_agent_id(PARAM_AGENT_ID),
        caller: p.get_agent_id(PARAM_CALLER),
        chain_id: p.get_chain_id(PARAM_CHAIN_ID),
        chain_owner_id: p.get_agent_id(PARAM_CHAIN_OWNER_ID),
        contract_creator: p.get_agent_id(PARAM_CONTRACT_CREATOR),
        contract_id: p.get_contract_id(PARAM_CONTRACT_ID),
    };
    ctx.require(params.agent_id.exists(), "missing mandatory agentId");
    ctx.require(params.caller.exists(), "missing mandatory caller");
    ctx.require(params.chain_id.exists(), "missing mandatory chainId");
    ctx.require(params.chain_owner_id.exists(), "missing mandatory chainOwnerId");
    ctx.require(params.contract_creator.exists(), "missing mandatory contractCreator");
    ctx.require(params.contract_id.exists(), "missing mandatory contractId");
    ctx.log("testcore.funcCheckContextFromFullEP");
    func_check_context_from_full_ep(ctx, &params);
    ctx.log("testcore.funcCheckContextFromFullEP ok");
}

pub struct FuncDoNothingParams {}

fn func_do_nothing_thunk(ctx: &ScFuncContext) {
    let params = FuncDoNothingParams {};
    ctx.log("testcore.funcDoNothing");
    func_do_nothing(ctx, &params);
    ctx.log("testcore.funcDoNothing ok");
}

pub struct FuncGetMintedSupplyParams {}

fn func_get_minted_supply_thunk(ctx: &ScFuncContext) {
    let params = FuncGetMintedSupplyParams {};
    ctx.log("testcore.funcGetMintedSupply");
    func_get_minted_supply(ctx, &params);
    ctx.log("testcore.funcGetMintedSupply ok");
}

pub struct FuncIncCounterParams {}

fn func_inc_counter_thunk(ctx: &ScFuncContext) {
    let params = FuncIncCounterParams {};
    ctx.log("testcore.funcIncCounter");
    func_inc_counter(ctx, &params);
    ctx.log("testcore.funcIncCounter ok");
}

pub struct FuncInitParams {}

fn func_init_thunk(ctx: &ScFuncContext) {
    let params = FuncInitParams {};
    ctx.log("testcore.funcInit");
    func_init(ctx, &params);
    ctx.log("testcore.funcInit ok");
}

//@formatter:off
pub struct FuncPassTypesFullParams {
    pub hash:        ScImmutableHash,
    pub hname:       ScImmutableHname,
    pub hname_zero:  ScImmutableHname,
    pub int64:       ScImmutableInt64,
    pub int64_zero:  ScImmutableInt64,
    pub string:      ScImmutableString,
    pub string_zero: ScImmutableString,
}
//@formatter:on

fn func_pass_types_full_thunk(ctx: &ScFuncContext) {
    let p = ctx.params();
    let params = FuncPassTypesFullParams {
        hash: p.get_hash(PARAM_HASH),
        hname: p.get_hname(PARAM_HNAME),
        hname_zero: p.get_hname(PARAM_HNAME_ZERO),
        int64: p.get_int64(PARAM_INT64),
        int64_zero: p.get_int64(PARAM_INT64_ZERO),
        string: p.get_string(PARAM_STRING),
        string_zero: p.get_string(PARAM_STRING_ZERO),
    };
    ctx.require(params.hash.exists(), "missing mandatory hash");
    ctx.require(params.hname.exists(), "missing mandatory hname");
    ctx.require(params.hname_zero.exists(), "missing mandatory hnameZero");
    ctx.require(params.int64.exists(), "missing mandatory int64");
    ctx.require(params.int64_zero.exists(), "missing mandatory int64Zero");
    ctx.require(params.string.exists(), "missing mandatory string");
    ctx.require(params.string_zero.exists(), "missing mandatory stringZero");
    ctx.log("testcore.funcPassTypesFull");
    func_pass_types_full(ctx, &params);
    ctx.log("testcore.funcPassTypesFull ok");
}

pub struct FuncRunRecursionParams {
    pub int_value: ScImmutableInt64,
}

fn func_run_recursion_thunk(ctx: &ScFuncContext) {
    let p = ctx.params();
    let params = FuncRunRecursionParams {
        int_value: p.get_int64(PARAM_INT_VALUE),
    };
    ctx.require(params.int_value.exists(), "missing mandatory intValue");
    ctx.log("testcore.funcRunRecursion");
    func_run_recursion(ctx, &params);
    ctx.log("testcore.funcRunRecursion ok");
}

pub struct FuncSendToAddressParams {
    pub address: ScImmutableAddress,
}

fn func_send_to_address_thunk(ctx: &ScFuncContext) {
    ctx.require(ctx.caller() == ctx.contract_creator(), "no permission");

    let p = ctx.params();
    let params = FuncSendToAddressParams {
        address: p.get_address(PARAM_ADDRESS),
    };
    ctx.require(params.address.exists(), "missing mandatory address");
    ctx.log("testcore.funcSendToAddress");
    func_send_to_address(ctx, &params);
    ctx.log("testcore.funcSendToAddress ok");
}

//@formatter:off
pub struct FuncSetIntParams {
    pub int_value: ScImmutableInt64,
    pub name:      ScImmutableString,
}
//@formatter:on

fn func_set_int_thunk(ctx: &ScFuncContext) {
    let p = ctx.params();
    let params = FuncSetIntParams {
        int_value: p.get_int64(PARAM_INT_VALUE),
        name: p.get_string(PARAM_NAME),
    };
    ctx.require(params.int_value.exists(), "missing mandatory intValue");
    ctx.require(params.name.exists(), "missing mandatory name");
    ctx.log("testcore.funcSetInt");
    func_set_int(ctx, &params);
    ctx.log("testcore.funcSetInt ok");
}

pub struct FuncTestCallPanicFullEPParams {}

fn func_test_call_panic_full_ep_thunk(ctx: &ScFuncContext) {
    let params = FuncTestCallPanicFullEPParams {};
    ctx.log("testcore.funcTestCallPanicFullEP");
    func_test_call_panic_full_ep(ctx, &params);
    ctx.log("testcore.funcTestCallPanicFullEP ok");
}

pub struct FuncTestCallPanicViewEPFromFullParams {}

fn func_test_call_panic_view_ep_from_full_thunk(ctx: &ScFuncContext) {
    let params = FuncTestCallPanicViewEPFromFullParams {};
    ctx.log("testcore.funcTestCallPanicViewEPFromFull");
    func_test_call_panic_view_ep_from_full(ctx, &params);
    ctx.log("testcore.funcTestCallPanicViewEPFromFull ok");
}

pub struct FuncTestChainOwnerIDFullParams {}

fn func_test_chain_owner_id_full_thunk(ctx: &ScFuncContext) {
    let params = FuncTestChainOwnerIDFullParams {};
    ctx.log("testcore.funcTestChainOwnerIDFull");
    func_test_chain_owner_id_full(ctx, &params);
    ctx.log("testcore.funcTestChainOwnerIDFull ok");
}

pub struct FuncTestContractIDFullParams {}

fn func_test_contract_id_full_thunk(ctx: &ScFuncContext) {
    let params = FuncTestContractIDFullParams {};
    ctx.log("testcore.funcTestContractIDFull");
    func_test_contract_id_full(ctx, &params);
    ctx.log("testcore.funcTestContractIDFull ok");
}

pub struct FuncTestEventLogDeployParams {}

fn func_test_event_log_deploy_thunk(ctx: &ScFuncContext) {
    let params = FuncTestEventLogDeployParams {};
    ctx.log("testcore.funcTestEventLogDeploy");
    func_test_event_log_deploy(ctx, &params);
    ctx.log("testcore.funcTestEventLogDeploy ok");
}

pub struct FuncTestEventLogEventDataParams {}

fn func_test_event_log_event_data_thunk(ctx: &ScFuncContext) {
    let params = FuncTestEventLogEventDataParams {};
    ctx.log("testcore.funcTestEventLogEventData");
    func_test_event_log_event_data(ctx, &params);
    ctx.log("testcore.funcTestEventLogEventData ok");
}

pub struct FuncTestEventLogGenericDataParams {
    pub counter: ScImmutableInt64,
}

fn func_test_event_log_generic_data_thunk(ctx: &ScFuncContext) {
    let p = ctx.params();
    let params = FuncTestEventLogGenericDataParams {
        counter: p.get_int64(PARAM_COUNTER),
    };
    ctx.require(params.counter.exists(), "missing mandatory counter");
    ctx.log("testcore.funcTestEventLogGenericData");
    func_test_event_log_generic_data(ctx, &params);
    ctx.log("testcore.funcTestEventLogGenericData ok");
}

pub struct FuncTestPanicFullEPParams {}

fn func_test_panic_full_ep_thunk(ctx: &ScFuncContext) {
    let params = FuncTestPanicFullEPParams {};
    ctx.log("testcore.funcTestPanicFullEP");
    func_test_panic_full_ep(ctx, &params);
    ctx.log("testcore.funcTestPanicFullEP ok");
}

pub struct FuncWithdrawToChainParams {
    pub chain_id: ScImmutableChainId,
}

fn func_withdraw_to_chain_thunk(ctx: &ScFuncContext) {
    let p = ctx.params();
    let params = FuncWithdrawToChainParams {
        chain_id: p.get_chain_id(PARAM_CHAIN_ID),
    };
    ctx.require(params.chain_id.exists(), "missing mandatory chainId");
    ctx.log("testcore.funcWithdrawToChain");
    func_withdraw_to_chain(ctx, &params);
    ctx.log("testcore.funcWithdrawToChain ok");
}

//@formatter:off
pub struct ViewCheckContextFromViewEPParams {
    pub agent_id:         ScImmutableAgentId,
    pub chain_id:         ScImmutableChainId,
    pub chain_owner_id:   ScImmutableAgentId,
    pub contract_creator: ScImmutableAgentId,
    pub contract_id:      ScImmutableContractId,
}
//@formatter:on

fn view_check_context_from_view_ep_thunk(ctx: &ScViewContext) {
    let p = ctx.params();
    let params = ViewCheckContextFromViewEPParams {
        agent_id: p.get_agent_id(PARAM_AGENT_ID),
        chain_id: p.get_chain_id(PARAM_CHAIN_ID),
        chain_owner_id: p.get_agent_id(PARAM_CHAIN_OWNER_ID),
        contract_creator: p.get_agent_id(PARAM_CONTRACT_CREATOR),
        contract_id: p.get_contract_id(PARAM_CONTRACT_ID),
    };
    ctx.require(params.agent_id.exists(), "missing mandatory agentId");
    ctx.require(params.chain_id.exists(), "missing mandatory chainId");
    ctx.require(params.chain_owner_id.exists(), "missing mandatory chainOwnerId");
    ctx.require(params.contract_creator.exists(), "missing mandatory contractCreator");
    ctx.require(params.contract_id.exists(), "missing mandatory contractId");
    ctx.log("testcore.viewCheckContextFromViewEP");
    view_check_context_from_view_ep(ctx, &params);
    ctx.log("testcore.viewCheckContextFromViewEP ok");
}

pub struct ViewFibonacciParams {
    pub int_value: ScImmutableInt64,
}

fn view_fibonacci_thunk(ctx: &ScViewContext) {
    let p = ctx.params();
    let params = ViewFibonacciParams {
        int_value: p.get_int64(PARAM_INT_VALUE),
    };
    ctx.require(params.int_value.exists(), "missing mandatory intValue");
    ctx.log("testcore.viewFibonacci");
    view_fibonacci(ctx, &params);
    ctx.log("testcore.viewFibonacci ok");
}

pub struct ViewGetCounterParams {}

fn view_get_counter_thunk(ctx: &ScViewContext) {
    let params = ViewGetCounterParams {};
    ctx.log("testcore.viewGetCounter");
    view_get_counter(ctx, &params);
    ctx.log("testcore.viewGetCounter ok");
}

pub struct ViewGetIntParams {
    pub name: ScImmutableString,
}

fn view_get_int_thunk(ctx: &ScViewContext) {
    let p = ctx.params();
    let params = ViewGetIntParams {
        name: p.get_string(PARAM_NAME),
    };
    ctx.require(params.name.exists(), "missing mandatory name");
    ctx.log("testcore.viewGetInt");
    view_get_int(ctx, &params);
    ctx.log("testcore.viewGetInt ok");
}

pub struct ViewJustViewParams {}

fn view_just_view_thunk(ctx: &ScViewContext) {
    let params = ViewJustViewParams {};
    ctx.log("testcore.viewJustView");
    view_just_view(ctx, &params);
    ctx.log("testcore.viewJustView ok");
}

//@formatter:off
pub struct ViewPassTypesViewParams {
    pub hash:        ScImmutableHash,
    pub hname:       ScImmutableHname,
    pub hname_zero:  ScImmutableHname,
    pub int64:       ScImmutableInt64,
    pub int64_zero:  ScImmutableInt64,
    pub string:      ScImmutableString,
    pub string_zero: ScImmutableString,
}
//@formatter:on

fn view_pass_types_view_thunk(ctx: &ScViewContext) {
    let p = ctx.params();
    let params = ViewPassTypesViewParams {
        hash: p.get_hash(PARAM_HASH),
        hname: p.get_hname(PARAM_HNAME),
        hname_zero: p.get_hname(PARAM_HNAME_ZERO),
        int64: p.get_int64(PARAM_INT64),
        int64_zero: p.get_int64(PARAM_INT64_ZERO),
        string: p.get_string(PARAM_STRING),
        string_zero: p.get_string(PARAM_STRING_ZERO),
    };
    ctx.require(params.hash.exists(), "missing mandatory hash");
    ctx.require(params.hname.exists(), "missing mandatory hname");
    ctx.require(params.hname_zero.exists(), "missing mandatory hnameZero");
    ctx.require(params.int64.exists(), "missing mandatory int64");
    ctx.require(params.int64_zero.exists(), "missing mandatory int64Zero");
    ctx.require(params.string.exists(), "missing mandatory string");
    ctx.require(params.string_zero.exists(), "missing mandatory stringZero");
    ctx.log("testcore.viewPassTypesView");
    view_pass_types_view(ctx, &params);
    ctx.log("testcore.viewPassTypesView ok");
}

pub struct ViewTestCallPanicViewEPFromViewParams {}

fn view_test_call_panic_view_ep_from_view_thunk(ctx: &ScViewContext) {
    let params = ViewTestCallPanicViewEPFromViewParams {};
    ctx.log("testcore.viewTestCallPanicViewEPFromView");
    view_test_call_panic_view_ep_from_view(ctx, &params);
    ctx.log("testcore.viewTestCallPanicViewEPFromView ok");
}

pub struct ViewTestChainOwnerIDViewParams {}

fn view_test_chain_owner_id_view_thunk(ctx: &ScViewContext) {
    let params = ViewTestChainOwnerIDViewParams {};
    ctx.log("testcore.viewTestChainOwnerIDView");
    view_test_chain_owner_id_view(ctx, &params);
    ctx.log("testcore.viewTestChainOwnerIDView ok");
}

pub struct ViewTestContractIDViewParams {}

fn view_test_contract_id_view_thunk(ctx: &ScViewContext) {
    let params = ViewTestContractIDViewParams {};
    ctx.log("testcore.viewTestContractIDView");
    view_test_contract_id_view(ctx, &params);
    ctx.log("testcore.viewTestContractIDView ok");
}

pub struct ViewTestPanicViewEPParams {}

fn view_test_panic_view_ep_thunk(ctx: &ScViewContext) {
    let params = ViewTestPanicViewEPParams {};
    ctx.log("testcore.viewTestPanicViewEP");
    view_test_panic_view_ep(ctx, &params);
    ctx.log("testcore.viewTestPanicViewEP ok");
}

pub struct ViewTestSandboxCallParams {}

fn view_test_sandbox_call_thunk(ctx: &ScViewContext) {
    let params = ViewTestSandboxCallParams {};
    ctx.log("testcore.viewTestSandboxCall");
    view_test_sandbox_call(ctx, &params);
    ctx.log("testcore.viewTestSandboxCall ok");
}
