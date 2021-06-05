// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use wasmlib::*;

use crate::*;

const CONTRACT_NAME_DEPLOYED: &str = "exampleDeployTR";
const MSG_FULL_PANIC: &str = "========== panic FULL ENTRY POINT =========";
const MSG_VIEW_PANIC: &str = "========== panic VIEW =========";

pub fn func_call_on_chain(ctx: &ScFuncContext, f: &FuncCallOnChainContext) {
    let param_int = f.params.int_value.value();

    let mut target_contract = ctx.contract();
    if f.params.hname_contract.exists() {
        target_contract = f.params.hname_contract.value();
    }

    let mut target_ep = HFUNC_CALL_ON_CHAIN;
    if f.params.hname_ep.exists() {
        target_ep = f.params.hname_ep.value();
    }

    let counter = f.state.counter();
    ctx.log(&format!("call depth = {}, hnameContract = {}, hnameEP = {}, counter = {}",
                     &f.params.int_value.to_string(),
                     &target_contract.to_string(),
                     &target_ep.to_string(),
                     &counter.to_string()));

    counter.set_value(counter.value() + 1);

    let parms = ScMutableMap::new();
    parms.get_int64(PARAM_INT_VALUE).set_value(param_int);
    let ret = ctx.call(target_contract, target_ep, Some(parms), None);
    let ret_val = ret.get_int64(PARAM_INT_VALUE);
    f.results.int_value.set_value(ret_val.value());
}

pub fn func_check_context_from_full_ep(ctx: &ScFuncContext, f: &FuncCheckContextFromFullEPContext) {
    ctx.require(f.params.agent_id.value() == ctx.account_id(), "fail: agentID");
    ctx.require(f.params.caller.value() == ctx.caller(), "fail: caller");
    ctx.require(f.params.chain_id.value() == ctx.chain_id(), "fail: chainID");
    ctx.require(f.params.chain_owner_id.value() == ctx.chain_owner_id(), "fail: chainOwnerID");
    ctx.require(f.params.contract_creator.value() == ctx.contract_creator(), "fail: contractCreator");
}

pub fn func_do_nothing(ctx: &ScFuncContext, _f: &FuncDoNothingContext) {
    ctx.log("doing nothing...");
}

pub fn func_get_minted_supply(ctx: &ScFuncContext, f: &FuncGetMintedSupplyContext) {
    let minted = ctx.minted();
    let minted_colors = minted.colors();
    ctx.require(minted_colors.length() == 1, "test only supports one minted color");
    let color = minted_colors.get_color(0).value();
    let amount = minted.balance(&color);
    f.results.minted_supply.set_value(amount);
    f.results.minted_color.set_value(&color);
}

pub fn func_inc_counter(_ctx: &ScFuncContext, f: &FuncIncCounterContext) {
    let counter = f.state.counter();
    counter.set_value(counter.value() + 1);
}

pub fn func_init(ctx: &ScFuncContext, _f: &FuncInitContext) {
    ctx.log("doing nothing...");
}

pub fn func_pass_types_full(ctx: &ScFuncContext, f: &FuncPassTypesFullContext) {
    let hash = ctx.utility().hash_blake2b(PARAM_HASH.as_bytes());
    ctx.require(f.params.hash.value() == hash, "Hash wrong");
    ctx.require(f.params.int64.value() == 42, "int64 wrong");
    ctx.require(f.params.int64_zero.value() == 0, "int64-0 wrong");
    ctx.require(f.params.string.value() == PARAM_STRING, "string wrong");
    ctx.require(f.params.string_zero.value() == "", "string-0 wrong");
    ctx.require(f.params.hname.value() == ScHname::new(PARAM_HNAME), "Hname wrong");
    ctx.require(f.params.hname_zero.value() == ScHname(0), "Hname-0 wrong");
}

pub fn func_run_recursion(ctx: &ScFuncContext, f: &FuncRunRecursionContext) {
    let depth = f.params.int_value.value();
    if depth <= 0 {
        return;
    }

    let parms = ScMutableMap::new();
    parms.get_int64(PARAM_INT_VALUE).set_value(depth - 1);
    parms.get_hname(PARAM_HNAME_EP).set_value(HFUNC_RUN_RECURSION);
    ctx.call_self(HFUNC_CALL_ON_CHAIN, Some(parms), None);
    // TODO how would I return result of the call ???
    f.results.int_value.set_value(depth - 1);
}

pub fn func_send_to_address(ctx: &ScFuncContext, f: &FuncSendToAddressContext) {
    let balances = ScTransfers::new_transfers_from_balances(ctx.balances());
    ctx.transfer_to_address(&f.params.address.value(), balances);
}

pub fn func_set_int(ctx: &ScFuncContext, f: &FuncSetIntContext) {
    ctx.state().get_int64(&f.params.name.value()).set_value(f.params.int_value.value());
}

pub fn func_test_call_panic_full_ep(ctx: &ScFuncContext, _f: &FuncTestCallPanicFullEPContext) {
    ctx.call_self(HFUNC_TEST_PANIC_FULL_EP, None, None);
}

pub fn func_test_call_panic_view_ep_from_full(ctx: &ScFuncContext, _f: &FuncTestCallPanicViewEPFromFullContext) {
    ctx.call_self(HVIEW_TEST_PANIC_VIEW_EP, None, None);
}

pub fn func_test_chain_owner_id_full(ctx: &ScFuncContext, f: &FuncTestChainOwnerIDFullContext) {
    f.results.chain_owner_id.set_value(&ctx.chain_owner_id());
}

pub fn func_test_event_log_deploy(ctx: &ScFuncContext, _f: &FuncTestEventLogDeployContext) {
    // deploy the same contract with another name
    let program_hash = ctx.utility().hash_blake2b("test_sandbox".as_bytes());
    ctx.deploy(&program_hash, CONTRACT_NAME_DEPLOYED, "test contract deploy log", None);
}

pub fn func_test_event_log_event_data(ctx: &ScFuncContext, _f: &FuncTestEventLogEventDataContext) {
    ctx.event("[Event] - Testing Event...");
}

pub fn func_test_event_log_generic_data(ctx: &ScFuncContext, f: &FuncTestEventLogGenericDataContext) {
    let event = "[GenericData] Counter Number: ".to_string() + &f.params.counter.to_string();
    ctx.event(&event);
}

pub fn func_test_panic_full_ep(ctx: &ScFuncContext, _f: &FuncTestPanicFullEPContext) {
    ctx.panic(MSG_FULL_PANIC);
}

pub fn func_withdraw_to_chain(ctx: &ScFuncContext, f: &FuncWithdrawToChainContext) {
    let transfer = ScTransfers::iotas(1);
    ctx.post(&f.params.chain_id.value(), CORE_ACCOUNTS, CORE_ACCOUNTS_FUNC_WITHDRAW, None, transfer, 0);
}

pub fn view_check_context_from_view_ep(ctx: &ScViewContext, f: &ViewCheckContextFromViewEPContext) {
    ctx.require(f.params.agent_id.value() == ctx.account_id(), "fail: agentID");
    ctx.require(f.params.chain_id.value() == ctx.chain_id(), "fail: chainID");
    ctx.require(f.params.chain_owner_id.value() == ctx.chain_owner_id(), "fail: chainOwnerID");
    ctx.require(f.params.contract_creator.value() == ctx.contract_creator(), "fail: contractCreator");
}

pub fn view_fibonacci(ctx: &ScViewContext, f: &ViewFibonacciContext) {
    let n = f.params.int_value.value();
    if n == 0 || n == 1 {
        f.results.int_value.set_value(n);
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

    f.results.int_value.set_value(n1 + n2);
}

pub fn view_get_counter(_ctx: &ScViewContext, f: &ViewGetCounterContext) {
    f.results.counter.set_value(f.state.counter().value());
}

pub fn view_get_int(ctx: &ScViewContext, f: &ViewGetIntContext) {
    let name = f.params.name.value();
    let value = ctx.state().get_int64(&name);
    ctx.require(value.exists(), "param 'value' not found");
    ctx.results().get_int64(&name).set_value(value.value());
}

pub fn view_just_view(ctx: &ScViewContext, _f: &ViewJustViewContext) {
    ctx.log("doing nothing...");
}

pub fn view_pass_types_view(ctx: &ScViewContext, f: &ViewPassTypesViewContext) {
    let hash = ctx.utility().hash_blake2b(PARAM_HASH.as_bytes());
    ctx.require(f.params.hash.value() == hash, "Hash wrong");
    ctx.require(f.params.int64.value() == 42, "int64 wrong");
    ctx.require(f.params.int64_zero.value() == 0, "int64-0 wrong");
    ctx.require(f.params.string.value() == PARAM_STRING, "string wrong");
    ctx.require(f.params.string_zero.value() == "", "string-0 wrong");
    ctx.require(f.params.hname.value() == ScHname::new(PARAM_HNAME), "Hname wrong");
    ctx.require(f.params.hname_zero.value() == ScHname(0), "Hname-0 wrong");
}

pub fn view_test_call_panic_view_ep_from_view(ctx: &ScViewContext, _f: &ViewTestCallPanicViewEPFromViewContext) {
    ctx.call_self(HVIEW_TEST_PANIC_VIEW_EP, None);
}

pub fn view_test_chain_owner_id_view(ctx: &ScViewContext, f: &ViewTestChainOwnerIDViewContext) {
    f.results.chain_owner_id.set_value(&ctx.chain_owner_id());
}

pub fn view_test_panic_view_ep(ctx: &ScViewContext, _f: &ViewTestPanicViewEPContext) {
    ctx.panic(MSG_VIEW_PANIC);
}

pub fn view_test_sandbox_call(ctx: &ScViewContext, f: &ViewTestSandboxCallContext) {
    let ret = ctx.call(CORE_ROOT, CORE_ROOT_VIEW_GET_CHAIN_INFO, None);
    let desc = ret.get_string("d").value();
    f.results.sandbox_call.set_value(&desc);
}
