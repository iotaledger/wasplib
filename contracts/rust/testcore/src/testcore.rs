// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use wasmlib::*;

use crate::*;

const CONTRACT_NAME_DEPLOYED: &str = "exampleDeployTR";
const MSG_FULL_PANIC: &str = "========== panic FULL ENTRY POINT =========";
const MSG_VIEW_PANIC: &str = "========== panic VIEW =========";

pub fn func_call_on_chain(ctx: &ScFuncContext, params: &FuncCallOnChainParams) {
    let param_int = params.int_value.value();

    let mut target_contract = ctx.contract();
    if params.hname_contract.exists() {
        target_contract = params.hname_contract.value();
    }

    let mut target_ep = HFUNC_CALL_ON_CHAIN;
    if params.hname_ep.exists() {
        target_ep = params.hname_ep.value();
    }

    let var_counter = ctx.state().get_int64(VAR_COUNTER);
    let counter = var_counter.value();
    var_counter.set_value(counter + 1);

    ctx.log(&format!("call depth = {} hnameContract = {} hnameEP = {} counter = {}",
                     param_int, &target_contract.to_string(), &target_ep.to_string(), counter));

    let parms = ScMutableMap::new();
    parms.get_int64(PARAM_INT_VALUE).set_value(param_int);
    let ret = ctx.call(target_contract, target_ep, Some(parms), None);

    let ret_val = ret.get_int64(PARAM_INT_VALUE);
    ctx.results().get_int64(PARAM_INT_VALUE).set_value(ret_val.value());
}

pub fn func_check_context_from_full_ep(ctx: &ScFuncContext, params: &FuncCheckContextFromFullEPParams) {
    ctx.require(params.agent_id.value() == ctx.account_id(), "fail: agentID");
    ctx.require(params.caller.value() == ctx.caller(), "fail: caller");
    ctx.require(params.chain_id.value() == ctx.chain_id(), "fail: chainID");
    ctx.require(params.chain_owner_id.value() == ctx.chain_owner_id(), "fail: chainOwnerID");
    ctx.require(params.contract_creator.value() == ctx.contract_creator(), "fail: contractCreator");
}

pub fn func_do_nothing(ctx: &ScFuncContext, _params: &FuncDoNothingParams) {
    ctx.log("doing nothing...");
}

pub fn func_get_minted_supply(ctx: &ScFuncContext, _params: &FuncGetMintedSupplyParams) {
    let minted = ctx.minted();
    let minted_colors = minted.colors();
    ctx.require(minted_colors.length() == 1, "test only supports one minted color");
    let color = minted_colors.get_color(0).value();
    let amount = minted.balance(&color);
    ctx.results().get_int64(VAR_MINTED_SUPPLY).set_value(amount);
    ctx.results().get_color(VAR_MINTED_COLOR).set_value(&color);
}

pub fn func_inc_counter(ctx: &ScFuncContext, _params: &FuncIncCounterParams) {
    ctx.state().get_int64(VAR_COUNTER).set_value(ctx.state().get_int64(VAR_COUNTER).value() + 1);
}

pub fn func_init(ctx: &ScFuncContext, _params: &FuncInitParams) {
    ctx.log("doing nothing...");
}

pub fn func_pass_types_full(ctx: &ScFuncContext, params: &FuncPassTypesFullParams) {
    let hash = ctx.utility().hash_blake2b(PARAM_HASH.as_bytes());
    ctx.require(params.hash.value() == hash, "Hash wrong");
    ctx.require(params.int64.value() == 42, "int64 wrong");
    ctx.require(params.int64_zero.value() == 0, "int64-0 wrong");
    ctx.require(params.string.value() == PARAM_STRING, "string wrong");
    ctx.require(params.string_zero.value() == "", "string-0 wrong");
    ctx.require(params.hname.value() == ScHname::new(PARAM_HNAME), "Hname wrong");
    ctx.require(params.hname_zero.value() == ScHname(0), "Hname-0 wrong");
}

pub fn func_run_recursion(ctx: &ScFuncContext, params: &FuncRunRecursionParams) {
    let depth = params.int_value.value();
    if depth <= 0 {
        return;
    }

    let parms = ScMutableMap::new();
    parms.get_int64(PARAM_INT_VALUE).set_value(depth - 1);
    parms.get_hname(PARAM_HNAME_EP).set_value(HFUNC_RUN_RECURSION);
    ctx.call_self(HFUNC_CALL_ON_CHAIN, Some(parms), None);
    // TODO how would I return result of the call ???
    ctx.results().get_int64(PARAM_INT_VALUE).set_value(depth - 1);
}

pub fn func_send_to_address(ctx: &ScFuncContext, params: &FuncSendToAddressParams) {
    let balances = ScTransfers::new_transfers_from_balances(ctx.balances());
    ctx.transfer_to_address(&params.address.value(), balances);
}

pub fn func_set_int(ctx: &ScFuncContext, params: &FuncSetIntParams) {
    ctx.state().get_int64(&params.name.value()).set_value(params.int_value.value());
}

pub fn func_test_call_panic_full_ep(ctx: &ScFuncContext, _params: &FuncTestCallPanicFullEPParams) {
    ctx.call_self(HFUNC_TEST_PANIC_FULL_EP, None, None);
}

pub fn func_test_call_panic_view_ep_from_full(ctx: &ScFuncContext, _params: &FuncTestCallPanicViewEPFromFullParams) {
    ctx.call_self(HVIEW_TEST_PANIC_VIEW_EP, None, None);
}

pub fn func_test_chain_owner_id_full(ctx: &ScFuncContext, _params: &FuncTestChainOwnerIDFullParams) {
    ctx.results().get_agent_id(PARAM_CHAIN_OWNER_ID).set_value(&ctx.chain_owner_id());
}

pub fn func_test_event_log_deploy(ctx: &ScFuncContext, _params: &FuncTestEventLogDeployParams) {
    //Deploy the same contract with another name
    let program_hash = ctx.utility().hash_blake2b("test_sandbox".as_bytes());
    ctx.deploy(&program_hash, CONTRACT_NAME_DEPLOYED, "test contract deploy log", None);
}

pub fn func_test_event_log_event_data(ctx: &ScFuncContext, _params: &FuncTestEventLogEventDataParams) {
    ctx.event("[Event] - Testing Event...");
}

pub fn func_test_event_log_generic_data(ctx: &ScFuncContext, params: &FuncTestEventLogGenericDataParams) {
    let event = "[GenericData] Counter Number: ".to_string() + &params.counter.to_string();
    ctx.event(&event);
}

pub fn func_test_panic_full_ep(ctx: &ScFuncContext, _params: &FuncTestPanicFullEPParams) {
    ctx.panic(MSG_FULL_PANIC);
}

pub fn func_withdraw_to_chain(ctx: &ScFuncContext, params: &FuncWithdrawToChainParams) {
    let transfer = ScTransfers::iotas(1);
    ctx.post(&params.chain_id.value(), CORE_ACCOUNTS, CORE_ACCOUNTS_FUNC_WITHDRAW, None, transfer, 0);
    ctx.log("====  success ====");
}

pub fn view_check_context_from_view_ep(ctx: &ScViewContext, params: &ViewCheckContextFromViewEPParams) {
    ctx.require(params.agent_id.value() == ctx.account_id(), "fail: agentID");
    ctx.require(params.chain_id.value() == ctx.chain_id(), "fail: chainID");
    ctx.require(params.chain_owner_id.value() == ctx.chain_owner_id(), "fail: chainOwnerID");
    ctx.require(params.contract_creator.value() == ctx.contract_creator(), "fail: contractCreator");
}

pub fn view_fibonacci(ctx: &ScViewContext, params: &ViewFibonacciParams) {
    let n = params.int_value.value();
    if n == 0 || n == 1 {
        ctx.results().get_int64(PARAM_INT_VALUE).set_value(n);
        return;
    }
    let parms1 = ScMutableMap::new();
    parms1.get_int64(PARAM_INT_VALUE).set_value(n - 1);
    let results1 = ctx.call_self(HVIEW_FIBONACCI, Some(parms1));
    let n1 = results1.get_int64(PARAM_INT_VALUE).value();

    let parms2 = ScMutableMap::new();
    parms2.get_int64(PARAM_INT_VALUE).set_value(n - 2);
    let results2 = ctx.call_self(HVIEW_FIBONACCI, Some(parms2));
    let n2 = results2.get_int64(PARAM_INT_VALUE).value();

    ctx.results().get_int64(PARAM_INT_VALUE).set_value(n1 + n2);
}

pub fn view_get_counter(ctx: &ScViewContext, _params: &ViewGetCounterParams) {
    let counter = ctx.state().get_int64(VAR_COUNTER);
    ctx.results().get_int64(VAR_COUNTER).set_value(counter.value());
}

pub fn view_get_int(ctx: &ScViewContext, params: &ViewGetIntParams) {
    let name = params.name.value();
    let value = ctx.state().get_int64(&name);
    ctx.require(value.exists(), "param 'value' not found");
    ctx.results().get_int64(&name).set_value(value.value());
}

pub fn view_just_view(ctx: &ScViewContext, _params: &ViewJustViewParams) {
    ctx.log("doing nothing...");
}

pub fn view_pass_types_view(ctx: &ScViewContext, params: &ViewPassTypesViewParams) {
    let hash = ctx.utility().hash_blake2b(PARAM_HASH.as_bytes());
    ctx.require(params.hash.value() == hash, "Hash wrong");
    ctx.require(params.int64.value() == 42, "int64 wrong");
    ctx.require(params.int64_zero.value() == 0, "int64-0 wrong");
    ctx.require(params.string.value() == PARAM_STRING, "string wrong");
    ctx.require(params.string_zero.value() == "", "string-0 wrong");
    ctx.require(params.hname.value() == ScHname::new(PARAM_HNAME), "Hname wrong");
    ctx.require(params.hname_zero.value() == ScHname(0), "Hname-0 wrong");
}

pub fn view_test_call_panic_view_ep_from_view(ctx: &ScViewContext, _params: &ViewTestCallPanicViewEPFromViewParams) {
    ctx.call_self(HVIEW_TEST_PANIC_VIEW_EP, None);
}

pub fn view_test_chain_owner_id_view(ctx: &ScViewContext, _params: &ViewTestChainOwnerIDViewParams) {
    ctx.results().get_agent_id(PARAM_CHAIN_OWNER_ID).set_value(&ctx.chain_owner_id());
}

pub fn view_test_panic_view_ep(ctx: &ScViewContext, _params: &ViewTestPanicViewEPParams) {
    ctx.panic(MSG_VIEW_PANIC);
}

pub fn view_test_sandbox_call(ctx: &ScViewContext, _params: &ViewTestSandboxCallParams) {
    let ret = ctx.call(CORE_ROOT, CORE_ROOT_VIEW_GET_CHAIN_INFO, None);
    let desc = ret.get_string("d").value();
    ctx.results().get_string("sandboxCall").set_value(&desc);
}
