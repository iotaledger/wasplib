// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use wasmlib::*;

use crate::*;

const CONTRACT_NAME_DEPLOYED: &str = "exampleDeployTR";
const MSG_FULL_PANIC: &str = "========== panic FULL ENTRY POINT =========";
const MSG_VIEW_PANIC: &str = "========== panic VIEW =========";

pub fn func_call_on_chain(ctx: &ScFuncContext, params: &FuncCallOnChainParams) {
    ctx.log("calling callOnChain");
    let param_in = params.int_value.value();

    let mut target_contract = ctx.contract_id().hname();
    if params.hname_contract.exists() {
        target_contract = params.hname_contract.value()
    }

    let mut target_ep = HFUNC_CALL_ON_CHAIN;
    let param_hname_ep = ctx.params().get_hname(PARAM_HNAME_EP);
    if param_hname_ep.exists() {
        target_ep = param_hname_ep.value()
    }

    let var_counter = ctx.state().get_int(VAR_COUNTER);
    let counter = var_counter.value();
    var_counter.set_value(counter + 1);

    ctx.log(&format!("call depth = {} hnameContract = {} hnameEP = {} counter = {}",
                     param_in, &target_contract.to_string(), &target_ep.to_string(), counter));

    let par = ScMutableMap::new();
    par.get_int(PARAM_INT_VALUE).set_value(param_in);
    let ret = ctx.call(target_contract, target_ep, Some(par), None);

    let ret_val = ret.get_int(PARAM_INT_VALUE);

    ctx.results().get_int(PARAM_INT_VALUE).set_value(ret_val.value());
}

pub fn func_check_context_from_full_ep(ctx: &ScFuncContext, params: &FuncCheckContextFromFullEPParams) {
    ctx.log("calling checkContextFromFullEP");

    ctx.require(params.chain_id.value() == ctx.contract_id().chain_id(), "fail: chainID");
    ctx.require(params.chain_owner_id.value() == ctx.chain_owner_id(), "fail: chainOwnerID");
    ctx.require(params.caller.value() == ctx.caller(), "fail: caller");
    ctx.require(params.contract_id.value() == ctx.contract_id(), "fail: contractID");
    ctx.require(params.agent_id.value() == ctx.contract_id().as_agent_id(), "fail: agentID");
    ctx.require(params.contract_creator.value() == ctx.contract_creator(), "fail: contractCreator");
}

pub fn func_do_nothing(ctx: &ScFuncContext, _params: &FuncDoNothingParams) {
    ctx.log("calling doNothing");
}

pub fn func_init(ctx: &ScFuncContext, _params: &FuncInitParams) {
    ctx.log("calling init");
}

pub fn func_pass_types_full(ctx: &ScFuncContext, params: &FuncPassTypesFullParams) {
    ctx.log("calling passTypesFull");

    ctx.require(params.int64.value() == 42, "int64 wrong");
    ctx.require(params.int64_zero.value() == 0, "int64-0 wrong");
    ctx.require(params.string.value() == PARAM_STRING, "string wrong");
    ctx.require(params.string_zero.value() == "", "string-0 wrong");

    let hash = ctx.utility().hash_blake2b(PARAM_HASH.as_bytes());
    ctx.require(params.hash.value().equals(&hash), "Hash wrong");

    ctx.require(params.hname.value().equals(ScHname::new(PARAM_HNAME)), "Hname wrong");
    ctx.require(params.hname_zero.value().equals(ScHname(0)), "Hname-0 wrong");
}

pub fn func_run_recursion(ctx: &ScFuncContext, params: &FuncRunRecursionParams) {
    ctx.log("calling runRecursion");
    let depth = params.int_value.value();
    if depth <= 0 {
        return;
    }
    let par = ScMutableMap::new();
    par.get_int(PARAM_INT_VALUE).set_value(depth - 1);
    par.get_hname(PARAM_HNAME_EP).set_value(HFUNC_RUN_RECURSION);
    ctx.call_self(HFUNC_CALL_ON_CHAIN, Some(par), None);
    // TODO how would I return result of the call ???
    ctx.results().get_int(PARAM_INT_VALUE).set_value(depth - 1);
}

pub fn func_send_to_address(ctx: &ScFuncContext, params: &FuncSendToAddressParams) {
    ctx.log("calling sendToAddress");
    ctx.transfer_to_address(&params.address.value(), &ctx.balances());
}

pub fn func_set_int(ctx: &ScFuncContext, params: &FuncSetIntParams) {
    ctx.log("calling setInt");
    ctx.state().get_int(&params.name.value() as &str).set_value(params.int_value.value());
}

pub fn func_test_call_panic_full_ep(ctx: &ScFuncContext, _params: &FuncTestCallPanicFullEPParams) {
    ctx.log("calling testCallPanicFullEP");
    ctx.call_self(HFUNC_TEST_PANIC_FULL_EP, None, None);
}

pub fn func_test_call_panic_view_epfrom_full(ctx: &ScFuncContext, _params: &FuncTestCallPanicViewEPFromFullParams) {
    ctx.log("calling testCallPanicViewEPFromFull");
    ctx.call_self(HVIEW_TEST_PANIC_VIEW_EP, None, None);
}

pub fn func_test_chain_owner_idfull(ctx: &ScFuncContext, _params: &FuncTestChainOwnerIDFullParams) {
    ctx.log("calling testChainOwnerIDFull");
    ctx.results().get_agent_id(PARAM_CHAIN_OWNER_ID).set_value(&ctx.chain_owner_id())
}

pub fn func_test_contract_idfull(ctx: &ScFuncContext, _params: &FuncTestContractIDFullParams) {
    ctx.log("calling testContractIDFull");
    ctx.results().get_contract_id(PARAM_CONTRACT_ID).set_value(&ctx.contract_id());
}

pub fn func_test_event_log_deploy(ctx: &ScFuncContext, _params: &FuncTestEventLogDeployParams) {
    ctx.log("calling testEventLogDeploy");
    //Deploy the same contract with another name
    let program_hash = ctx.utility().hash_blake2b("test_sandbox".as_bytes());
    ctx.deploy(&program_hash, CONTRACT_NAME_DEPLOYED,
               "test contract deploy log", None)
}

pub fn func_test_event_log_event_data(ctx: &ScFuncContext, _params: &FuncTestEventLogEventDataParams) {
    ctx.log("calling testEventLogEventData");
    ctx.event("[Event] - Testing Event...");
}

pub fn func_test_event_log_generic_data(ctx: &ScFuncContext, params: &FuncTestEventLogGenericDataParams) {
    ctx.log("calling testEventLogGenericData");
    let event = "[GenericData] Counter Number: ".to_string() + &params.counter.to_string();
    ctx.event(&event)
}

pub fn func_test_panic_full_ep(ctx: &ScFuncContext, _params: &FuncTestPanicFullEPParams) {
    ctx.log("calling testPanicFullEP");
    ctx.panic(MSG_FULL_PANIC)
}

pub fn func_withdraw_to_chain(ctx: &ScFuncContext, params: &FuncWithdrawToChainParams) {
    ctx.log("calling withdrawToChain");
    //Deploy the same contract with another name
    let target_contract_id = ScContractId::new(&params.chain_id.value(), &CORE_ACCOUNTS);
    ctx.post(&PostRequestParams {
        contract_id: target_contract_id,
        function: CORE_ACCOUNTS_FUNC_WITHDRAW_TO_CHAIN,
        params: None,
        transfer: Some(Box::new(ScTransfers::new(&ScColor::IOTA, 2))),
        delay: 0,
    });
    ctx.log("====  success ====");
    // TODO how to check if post was successful
}

pub fn view_check_context_from_view_ep(ctx: &ScViewContext, params: &ViewCheckContextFromViewEPParams) {
    ctx.log("calling checkContextFromViewEP");

    ctx.require(params.chain_id.value() == ctx.contract_id().chain_id(), "fail: chainID");
    ctx.require(params.chain_owner_id.value() == ctx.chain_owner_id(), "fail: chainOwnerID");
    ctx.require(params.contract_id.value() == ctx.contract_id(), "fail: contractID");
    ctx.require(params.agent_id.value() == ctx.contract_id().as_agent_id(), "fail: agentID");
    ctx.require(params.contract_creator.value() == ctx.contract_creator(), "fail: contractCreator");
}

pub fn view_fibonacci(ctx: &ScViewContext, params: &ViewFibonacciParams) {
    ctx.log("calling fibonacci");
    let n = params.int_value.value();
    if n == 0 || n == 1 {
        ctx.results().get_int(PARAM_INT_VALUE).set_value(n);
        return;
    }
    let params1 = ScMutableMap::new();
    params1.get_int(PARAM_INT_VALUE).set_value(n - 1);
    let results1 = ctx.call_self(HVIEW_FIBONACCI, Some(params1));
    let n1 = results1.get_int(PARAM_INT_VALUE).value();

    let params2 = ScMutableMap::new();
    params2.get_int(PARAM_INT_VALUE).set_value(n - 2);
    let results2 = ctx.call_self(HVIEW_FIBONACCI, Some(params2));
    let n2 = results2.get_int(PARAM_INT_VALUE).value();

    ctx.results().get_int(PARAM_INT_VALUE).set_value(n1 + n2);
}

pub fn view_get_counter(ctx: &ScViewContext, _params: &ViewGetCounterParams) {
    ctx.log("calling getCounter");
    let counter = ctx.state().get_int(VAR_COUNTER);
    ctx.results().get_int(VAR_COUNTER).set_value(counter.value());
}

pub fn view_get_int(ctx: &ScViewContext, params: &ViewGetIntParams) {
    ctx.log("calling getInt");
    let name = params.name.value();
    let value = ctx.state().get_int(&name);
    ctx.require(value.exists(), "param 'value' not found");
    ctx.results().get_int(&name).set_value(value.value());
}

pub fn view_just_view(ctx: &ScViewContext, _params: &ViewJustViewParams) {
    ctx.log("calling justView");
}

pub fn view_pass_types_view(ctx: &ScViewContext, params: &ViewPassTypesViewParams) {
    ctx.log("calling passTypesView");

    ctx.require(params.int64.value() == 42, "int64 wrong");
    ctx.require(params.int64_zero.value() == 0, "int64-0 wrong");
    ctx.require(params.string.value() == PARAM_STRING, "string wrong");
    ctx.require(params.string_zero.value() == "", "string-0 wrong");

    let hash = ctx.utility().hash_blake2b(PARAM_HASH.as_bytes());
    ctx.require(params.hash.value().equals(&hash), "Hash wrong");

    ctx.require(params.hname.value().equals(ScHname::new(PARAM_HNAME)), "Hname wrong");
    ctx.require(params.hname_zero.value().equals(ScHname(0)), "Hname-0 wrong");
}

pub fn view_test_call_panic_view_epfrom_view(ctx: &ScViewContext, _params: &ViewTestCallPanicViewEPFromViewParams) {
    ctx.log("calling testCallPanicViewEPFromView");
    ctx.call_self(HVIEW_TEST_PANIC_VIEW_EP, None);
}

pub fn view_test_chain_owner_idview(ctx: &ScViewContext, _params: &ViewTestChainOwnerIDViewParams) {
    ctx.log("calling testChainOwnerIDView");
    ctx.results().get_agent_id(PARAM_CHAIN_OWNER_ID).set_value(&ctx.chain_owner_id())
}

pub fn view_test_contract_idview(ctx: &ScViewContext, _params: &ViewTestContractIDViewParams) {
    ctx.log("calling testContractIDView");
    ctx.results().get_contract_id(PARAM_CONTRACT_ID).set_value(&ctx.contract_id());
}

pub fn view_test_panic_view_ep(ctx: &ScViewContext, _params: &ViewTestPanicViewEPParams) {
    ctx.log("calling testPanicViewEP");
    ctx.panic(MSG_VIEW_PANIC)
}

pub fn view_test_sandbox_call(ctx: &ScViewContext, _params: &ViewTestSandboxCallParams) {
    ctx.log("calling testSandboxCall");
    let ret = ctx.call(CORE_ROOT, CORE_ROOT_VIEW_GET_CHAIN_INFO, None);
    let desc = ret.get_string("d").value();
    ctx.results().get_string("sandboxCall").set_value(&desc);
}
