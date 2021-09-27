// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use wasmlib::*;

use crate::*;
use crate::contract::*;

const CONTRACT_NAME_DEPLOYED: &str = "exampleDeployTR";
const MSG_CORE_ONLY_PANIC: &str = "========== core only =========";
const MSG_FULL_PANIC: &str = "========== panic FULL ENTRY POINT =========";
const MSG_VIEW_PANIC: &str = "========== panic VIEW =========";

pub fn func_call_on_chain(ctx: &ScFuncContext, f: &CallOnChainContext) {
    let param_int = f.params.int_value().value();

    let mut hname_contract = ctx.contract();
    if f.params.hname_contract().exists() {
        hname_contract = f.params.hname_contract().value();
    }

    let mut hname_ep = HFUNC_CALL_ON_CHAIN;
    if f.params.hname_ep().exists() {
        hname_ep = f.params.hname_ep().value();
    }

    let counter = f.state.counter();
    ctx.log(&format!("call depth = {}, hnameContract = {}, hnameEP = {}, counter = {}",
                     &f.params.int_value().to_string(),
                     &hname_contract.to_string(),
                     &hname_ep.to_string(),
                     &counter.to_string()));

    counter.set_value(counter.value() + 1);

    let parms = ScMutableMap::new();
    parms.get_int64(PARAM_INT_VALUE).set_value(param_int);
    let ret = ctx.call(hname_contract, hname_ep, Some(parms), None);
    let ret_val = ret.get_int64(RESULT_INT_VALUE);
    f.results.int_value().set_value(ret_val.value());
}

pub fn func_check_context_from_full_ep(ctx: &ScFuncContext, f: &CheckContextFromFullEPContext) {
    ctx.require(f.params.agent_id().value() == ctx.account_id(), "fail: agentID");
    ctx.require(f.params.caller().value() == ctx.caller(), "fail: caller");
    ctx.require(f.params.chain_id().value() == ctx.chain_id(), "fail: chainID");
    ctx.require(f.params.chain_owner_id().value() == ctx.chain_owner_id(), "fail: chainOwnerID");
    ctx.require(f.params.contract_creator().value() == ctx.contract_creator(), "fail: contractCreator");
}

pub fn func_do_nothing(ctx: &ScFuncContext, _f: &DoNothingContext) {
    ctx.log("doing nothing...");
}

pub fn func_get_minted_supply(ctx: &ScFuncContext, f: &GetMintedSupplyContext) {
    let minted = ctx.minted();
    let minted_colors = minted.colors();
    ctx.require(minted_colors.length() == 1, "test only supports one minted color");
    let color = minted_colors.get_color(0).value();
    let amount = minted.balance(&color);
    f.results.minted_supply().set_value(amount);
    f.results.minted_color().set_value(&color);
}

pub fn func_inc_counter(_ctx: &ScFuncContext, f: &IncCounterContext) {
    let counter = f.state.counter();
    counter.set_value(counter.value() + 1);
}

pub fn func_init(ctx: &ScFuncContext, f: &InitContext) {
    if f.params.fail().exists() {
        ctx.panic("failing on purpose");
    }
}

pub fn func_pass_types_full(ctx: &ScFuncContext, f: &PassTypesFullContext) {
    let hash = ctx.utility().hash_blake2b(PARAM_HASH.as_bytes());
    ctx.require(f.params.hash().value() == hash, "Hash wrong");
    ctx.require(f.params.int64().value() == 42, "int64 wrong");
    ctx.require(f.params.int64_zero().value() == 0, "int64-0 wrong");
    ctx.require(f.params.string().value() == PARAM_STRING, "string wrong");
    ctx.require(f.params.string_zero().value() == "", "string-0 wrong");
    ctx.require(f.params.hname().value() == ScHname::new(PARAM_HNAME), "Hname wrong");
    ctx.require(f.params.hname_zero().value() == ScHname(0), "Hname-0 wrong");
}

pub fn func_run_recursion(ctx: &ScFuncContext, f: &RunRecursionContext) {
    let depth = f.params.int_value().value();
    if depth <= 0 {
        return;
    }

    let call_on_chain = ScFuncs::call_on_chain(ctx);
    call_on_chain.params.int_value().set_value(depth - 1);
    call_on_chain.params.hname_ep().set_value(HFUNC_RUN_RECURSION);
    call_on_chain.func.call();
    let ret_val = call_on_chain.results.int_value().value();
    f.results.int_value().set_value(ret_val);
}

pub fn func_send_to_address(ctx: &ScFuncContext, f: &SendToAddressContext) {
    let balances = ScTransfers::new_transfers_from_balances(ctx.balances());
    ctx.transfer_to_address(&f.params.address().value(), balances);
}

pub fn func_set_int(_ctx: &ScFuncContext, f: &SetIntContext) {
    f.state.ints().get_int64(&f.params.name().value()).set_value(f.params.int_value().value());
}

pub fn func_test_call_panic_full_ep(ctx: &ScFuncContext, _f: &TestCallPanicFullEPContext) {
    ScFuncs::test_panic_full_ep(ctx).func.call();
}

pub fn func_test_call_panic_view_ep_from_full(ctx: &ScFuncContext, _f: &TestCallPanicViewEPFromFullContext) {
    ScFuncs::test_panic_view_ep(ctx).func.call();
}

pub fn func_test_chain_owner_id_full(ctx: &ScFuncContext, f: &TestChainOwnerIDFullContext) {
    f.results.chain_owner_id().set_value(&ctx.chain_owner_id());
}

pub fn func_test_event_log_deploy(ctx: &ScFuncContext, _f: &TestEventLogDeployContext) {
    // deploy the same contract with another name
    let program_hash = ctx.utility().hash_blake2b("testcore".as_bytes());
    ctx.deploy(&program_hash, CONTRACT_NAME_DEPLOYED, "test contract deploy log", None);
}

pub fn func_test_event_log_event_data(ctx: &ScFuncContext, _f: &TestEventLogEventDataContext) {
    ctx.event("[Event] - Testing Event...");
}

pub fn func_test_event_log_generic_data(ctx: &ScFuncContext, f: &TestEventLogGenericDataContext) {
    let event = "[GenericData] Counter Number: ".to_string() + &f.params.counter().to_string();
    ctx.event(&event);
}

pub fn func_test_panic_full_ep(ctx: &ScFuncContext, _f: &TestPanicFullEPContext) {
    ctx.panic(MSG_FULL_PANIC);
}

pub fn func_withdraw_to_chain(ctx: &ScFuncContext, f: &WithdrawToChainContext) {
    let xx = corecontracts::coreaccounts::ScFuncs::withdraw(ctx);
    xx.func.transfer_iotas(1).post_to_chain(f.params.chain_id().value());
}

pub fn view_check_context_from_view_ep(ctx: &ScViewContext, f: &CheckContextFromViewEPContext) {
    ctx.require(f.params.agent_id().value() == ctx.account_id(), "fail: agentID");
    ctx.require(f.params.chain_id().value() == ctx.chain_id(), "fail: chainID");
    ctx.require(f.params.chain_owner_id().value() == ctx.chain_owner_id(), "fail: chainOwnerID");
    ctx.require(f.params.contract_creator().value() == ctx.contract_creator(), "fail: contractCreator");
}

pub fn view_fibonacci(ctx: &ScViewContext, f: &FibonacciContext) {
    let n = f.params.int_value().value();
    if n == 0 || n == 1 {
        f.results.int_value().set_value(n);
        return;
    }

    let fib = ScFuncs::fibonacci(ctx);
    fib.params.int_value().set_value(n - 1);
    fib.func.call();
    let n1 = fib.results.int_value().value();

    fib.params.int_value().set_value(n - 2);
    fib.func.call();
    let n2 = fib.results.int_value().value();

    f.results.int_value().set_value(n1 + n2);
}

pub fn view_get_counter(_ctx: &ScViewContext, f: &GetCounterContext) {
    f.results.counter().set_value(f.state.counter().value());
}

pub fn view_get_int(ctx: &ScViewContext, f: &GetIntContext) {
    let name = f.params.name().value();
    let value = f.state.ints().get_int64(&name);
    ctx.require(value.exists(), "param 'value' not found");
    f.results.values().get_int64(&name).set_value(value.value());
}

pub fn view_just_view(ctx: &ScViewContext, _f: &JustViewContext) {
    ctx.log("doing nothing...");
}

pub fn view_pass_types_view(ctx: &ScViewContext, f: &PassTypesViewContext) {
    let hash = ctx.utility().hash_blake2b(PARAM_HASH.as_bytes());
    ctx.require(f.params.hash().value() == hash, "Hash wrong");
    ctx.require(f.params.int64().value() == 42, "int64 wrong");
    ctx.require(f.params.int64_zero().value() == 0, "int64-0 wrong");
    ctx.require(f.params.string().value() == PARAM_STRING, "string wrong");
    ctx.require(f.params.string_zero().value() == "", "string-0 wrong");
    ctx.require(f.params.hname().value() == ScHname::new(PARAM_HNAME), "Hname wrong");
    ctx.require(f.params.hname_zero().value() == ScHname(0), "Hname-0 wrong");
}

pub fn view_test_call_panic_view_ep_from_view(ctx: &ScViewContext, _f: &TestCallPanicViewEPFromViewContext) {
    ScFuncs::test_panic_view_ep(ctx).func.call();
}

pub fn view_test_chain_owner_id_view(ctx: &ScViewContext, f: &TestChainOwnerIDViewContext) {
    f.results.chain_owner_id().set_value(&ctx.chain_owner_id());
}

pub fn view_test_panic_view_ep(ctx: &ScViewContext, _f: &TestPanicViewEPContext) {
    ctx.panic(MSG_VIEW_PANIC);
}

pub fn view_test_sandbox_call(ctx: &ScViewContext, f: &TestSandboxCallContext) {
    let get_chain_info = corecontracts::coregovernance::ScFuncs::get_chain_info(ctx);
    get_chain_info.func.call();
    f.results.sandbox_call().set_value(&get_chain_info.results.description().value());
}

pub fn func_test_block_context1(ctx: &ScFuncContext, _f: &TestBlockContext1Context) {
    ctx.panic(MSG_CORE_ONLY_PANIC);
}

pub fn func_test_block_context2(ctx: &ScFuncContext, _f: &TestBlockContext2Context) {
    ctx.panic(MSG_CORE_ONLY_PANIC);
}

pub fn view_get_string_value(ctx: &ScViewContext, _f: &GetStringValueContext) {
    ctx.panic(MSG_CORE_ONLY_PANIC);
}
