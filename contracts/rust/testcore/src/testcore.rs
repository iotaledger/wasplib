// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use wasmlib::*;

use crate::*;

const CONTRACT_NAME_DEPLOYED: &str = "exampleDeployTR";
const MSG_FULL_PANIC: &str = "========== panic FULL ENTRY POINT =========";
const MSG_VIEW_PANIC: &str = "========== panic VIEW =========";

pub fn func_call_on_chain(ctx: &ScFuncContext) {
    ctx.log("calling callOnChain");

    let p = ctx.params();
    let param_hname_contract = p.get_hname(PARAM_HNAME_CONTRACT);
    let param_hname_ep = p.get_hname(PARAM_HNAME_EP);
    let param_int_value = p.get_int(PARAM_INT_VALUE);

    ctx.require(param_int_value.exists(), "missing mandatory intValue");

    let param_int = param_int_value.value();

    let mut target_contract = ctx.contract_id().hname();
    if param_hname_contract.exists() {
        target_contract = param_hname_contract.value()
    }

    let mut target_ep = HFUNC_CALL_ON_CHAIN;
    if param_hname_ep.exists() {
        target_ep = param_hname_ep.value()
    }

    let var_counter = ctx.state().get_int(VAR_COUNTER);
    let counter = var_counter.value();
    var_counter.set_value(counter + 1);

    ctx.log(&format!("call depth = {} hnameContract = {} hnameEP = {} counter = {}",
                     param_int, &target_contract.to_string(), &target_ep.to_string(), counter));

    let params = ScMutableMap::new();
    params.get_int(PARAM_INT_VALUE).set_value(param_int);
    let ret = ctx.call(target_contract, target_ep, Some(params), None);

    let ret_val = ret.get_int(PARAM_INT_VALUE);
    ctx.results().get_int(PARAM_INT_VALUE).set_value(ret_val.value());
}

pub fn func_check_context_from_full_ep(ctx: &ScFuncContext) {
    ctx.log("calling checkContextFromFullEP");

    let p = ctx.params();
    let param_agent_id = p.get_agent_id(PARAM_AGENT_ID);
    let param_caller = p.get_agent_id(PARAM_CALLER);
    let param_chain_id = p.get_chain_id(PARAM_CHAIN_ID);
    let param_chain_owner_id = p.get_agent_id(PARAM_CHAIN_OWNER_ID);
    let param_contract_creator = p.get_agent_id(PARAM_CONTRACT_CREATOR);
    let param_contract_id = p.get_contract_id(PARAM_CONTRACT_ID);

    ctx.require(param_chain_id.value() == ctx.contract_id().chain_id(), "fail: chainID");
    ctx.require(param_chain_owner_id.value() == ctx.chain_owner_id(), "fail: chainOwnerID");
    ctx.require(param_caller.value() == ctx.caller(), "fail: caller");
    ctx.require(param_contract_id.value() == ctx.contract_id(), "fail: contractID");
    ctx.require(param_agent_id.value() == ctx.contract_id().as_agent_id(), "fail: agentID");
    ctx.require(param_contract_creator.value() == ctx.contract_creator(), "fail: contractCreator");
}

pub fn func_do_nothing(ctx: &ScFuncContext) {
    ctx.log("calling doNothing");
}

pub fn func_init(ctx: &ScFuncContext) {
    ctx.log("calling init");
}

pub fn func_pass_types_full(ctx: &ScFuncContext) {
    ctx.log("calling passTypesFull");

    let p = ctx.params();
    let param_hash = p.get_hash(PARAM_HASH);
    let param_hname = p.get_hname(PARAM_HNAME);
    let param_hname_zero = p.get_hname(PARAM_HNAME_ZERO);
    let param_int64 = p.get_int(PARAM_INT64);
    let param_int64_zero = p.get_int(PARAM_INT64_ZERO);
    let param_string = p.get_string(PARAM_STRING);
    let param_string_zero = p.get_string(PARAM_STRING_ZERO);

    ctx.require(param_hash.exists(), "missing mandatory hash");
    ctx.require(param_hname.exists(), "missing mandatory hname");
    ctx.require(param_hname_zero.exists(), "missing mandatory hnameZero");
    ctx.require(param_int64.exists(), "missing mandatory int64");
    ctx.require(param_int64_zero.exists(), "missing mandatory int64Zero");
    ctx.require(param_string.exists(), "missing mandatory string");
    ctx.require(param_string_zero.exists(), "missing mandatory stringZero");

    let hash = ctx.utility().hash_blake2b(PARAM_HASH.as_bytes());
    ctx.require(param_hash.value() == hash, "Hash wrong");
    ctx.require(param_int64.value() == 42, "int64 wrong");
    ctx.require(param_int64_zero.value() == 0, "int64-0 wrong");
    ctx.require(param_string.value() == PARAM_STRING, "string wrong");
    ctx.require(param_hname.value() == ScHname::new(PARAM_HNAME), "Hname wrong");
    ctx.require(param_hname_zero.value() == ScHname(0), "Hname-0 wrong");
}

pub fn func_run_recursion(ctx: &ScFuncContext) {
    ctx.log("calling runRecursion");

    let p = ctx.params();
    let param_int_value = p.get_int(PARAM_INT_VALUE);

    ctx.require(param_int_value.exists(), "missing mandatory intValue");

    let depth = param_int_value.value();
    if depth <= 0 {
        return;
    }

    let params = ScMutableMap::new();
    params.get_int(PARAM_INT_VALUE).set_value(depth - 1);
    params.get_hname(PARAM_HNAME_EP).set_value(HFUNC_RUN_RECURSION);
    ctx.call_self(HFUNC_CALL_ON_CHAIN, Some(params), None);
    // TODO how would I return result of the call ???
    ctx.results().get_int(PARAM_INT_VALUE).set_value(depth - 1);
}

pub fn func_send_to_address(ctx: &ScFuncContext) {
    ctx.log("calling sendToAddress");

    ctx.require(ctx.caller() == ctx.contract_creator(), "no permission");

    let p = ctx.params();
    let param_address = p.get_address(PARAM_ADDRESS);

    ctx.require(param_address.exists(), "missing mandatory address");

    ctx.transfer_to_address(&param_address.value(), &ctx.balances());
}

pub fn func_set_int(ctx: &ScFuncContext) {
    ctx.log("calling setInt");

    let p = ctx.params();
    let param_int_value = p.get_int(PARAM_INT_VALUE);
    let param_name = p.get_string(PARAM_NAME);

    ctx.require(param_int_value.exists(), "missing mandatory intValue");
    ctx.require(param_name.exists(), "missing mandatory name");

    ctx.state().get_int(&param_name.value()).set_value(param_int_value.value());
}

pub fn func_test_call_panic_full_ep(ctx: &ScFuncContext) {
    ctx.log("calling testCallPanicFullEP");
    ctx.call_self(HFUNC_TEST_PANIC_FULL_EP, None, None);
}

pub fn func_test_call_panic_view_epfrom_full(ctx: &ScFuncContext) {
    ctx.log("calling testCallPanicViewEPFromFull");
    ctx.call_self(HVIEW_TEST_PANIC_VIEW_EP, None, None);
}

pub fn func_test_chain_owner_idfull(ctx: &ScFuncContext) {
    ctx.log("calling testChainOwnerIDFull");
    ctx.results().get_agent_id(PARAM_CHAIN_OWNER_ID).set_value(&ctx.chain_owner_id())
}

pub fn func_test_contract_idfull(ctx: &ScFuncContext) {
    ctx.log("calling testContractIDFull");
    ctx.results().get_contract_id(PARAM_CONTRACT_ID).set_value(&ctx.contract_id());
}

pub fn func_test_event_log_deploy(ctx: &ScFuncContext) {
    ctx.log("calling testEventLogDeploy");
    //Deploy the same contract with another name
    let program_hash = ctx.utility().hash_blake2b("test_sandbox".as_bytes());
    ctx.deploy(&program_hash, CONTRACT_NAME_DEPLOYED,
               "test contract deploy log", None)
}

pub fn func_test_event_log_event_data(ctx: &ScFuncContext) {
    ctx.log("calling testEventLogEventData");
    ctx.event("[Event] - Testing Event...");
}

pub fn func_test_event_log_generic_data(ctx: &ScFuncContext) {
    ctx.log("calling testEventLogGenericData");

    let p = ctx.params();
    let param_counter = p.get_int(PARAM_COUNTER);

    ctx.require(param_counter.exists(), "missing mandatory counter");

    let event = "[GenericData] Counter Number: ".to_string() + &param_counter.to_string();
    ctx.event(&event)
}

pub fn func_test_panic_full_ep(ctx: &ScFuncContext) {
    ctx.log("calling testPanicFullEP");
    ctx.panic(MSG_FULL_PANIC)
}

pub fn func_withdraw_to_chain(ctx: &ScFuncContext) {
    ctx.log("calling withdrawToChain");

    let p = ctx.params();
    let param_chain_id = p.get_chain_id(PARAM_CHAIN_ID);

    ctx.require(param_chain_id.exists(), "missing mandatory chainId");

    //Deploy the same contract with another name
    let target_contract_id = ScContractId::new(&param_chain_id.value(), &CORE_ACCOUNTS);
    let transfers = ScTransfers::new(&ScColor::IOTA, 2);
    ctx.post(&PostRequestParams {
        contract_id: target_contract_id,
        function: CORE_ACCOUNTS_FUNC_WITHDRAW_TO_CHAIN,
        params: None,
        transfer: Some(Box::new(transfers)),
        delay: 0,
    });
    ctx.log("====  success ====");
    // TODO how to check if post was successful
}

pub fn view_check_context_from_view_ep(ctx: &ScViewContext) {
    ctx.log("calling checkContextFromViewEP");

    let p = ctx.params();
    let param_agent_id = p.get_agent_id(PARAM_AGENT_ID);
    let param_chain_id = p.get_chain_id(PARAM_CHAIN_ID);
    let param_chain_owner_id = p.get_agent_id(PARAM_CHAIN_OWNER_ID);
    let param_contract_creator = p.get_agent_id(PARAM_CONTRACT_CREATOR);
    let param_contract_id = p.get_contract_id(PARAM_CONTRACT_ID);

    ctx.require(param_chain_id.value() == ctx.contract_id().chain_id(), "fail: chainID");
    ctx.require(param_chain_owner_id.value() == ctx.chain_owner_id(), "fail: chainOwnerID");
    ctx.require(param_contract_id.value() == ctx.contract_id(), "fail: contractID");
    ctx.require(param_agent_id.value() == ctx.contract_id().as_agent_id(), "fail: agentID");
    ctx.require(param_contract_creator.value() == ctx.contract_creator(), "fail: contractCreator");
}

pub fn view_fibonacci(ctx: &ScViewContext) {
    ctx.log("calling fibonacci");

    let p = ctx.params();
    let param_int_value = p.get_int(PARAM_INT_VALUE);

    ctx.require(param_int_value.exists(), "missing mandatory intValue");

    let n = param_int_value.value();
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

pub fn view_get_counter(ctx: &ScViewContext) {
    ctx.log("calling getCounter");
    let counter = ctx.state().get_int(VAR_COUNTER);
    ctx.results().get_int(VAR_COUNTER).set_value(counter.value());
}

pub fn view_get_int(ctx: &ScViewContext) {
    ctx.log("calling getInt");

    let p = ctx.params();
    let param_name = p.get_string(PARAM_NAME);

    ctx.require(param_name.exists(), "missing mandatory name");

    let name = param_name.value();
    let value = ctx.state().get_int(&name);
    ctx.require(value.exists(), "param 'value' not found");
    ctx.results().get_int(&name).set_value(value.value());
}

pub fn view_just_view(ctx: &ScViewContext) {
    ctx.log("calling justView");
}

pub fn view_pass_types_view(ctx: &ScViewContext) {
    ctx.log("calling passTypesView");

    let p = ctx.params();
    let param_hash = p.get_hash(PARAM_HASH);
    let param_hname = p.get_hname(PARAM_HNAME);
    let param_hname_zero = p.get_hname(PARAM_HNAME_ZERO);
    let param_int64 = p.get_int(PARAM_INT64);
    let param_int64_zero = p.get_int(PARAM_INT64_ZERO);
    let param_string = p.get_string(PARAM_STRING);
    let param_string_zero = p.get_string(PARAM_STRING_ZERO);

    ctx.require(param_hash.exists(), "missing mandatory hash");
    ctx.require(param_hname.exists(), "missing mandatory hname");
    ctx.require(param_hname_zero.exists(), "missing mandatory hnameZero");
    ctx.require(param_int64.exists(), "missing mandatory int64");
    ctx.require(param_int64_zero.exists(), "missing mandatory int64Zero");
    ctx.require(param_string.exists(), "missing mandatory string");
    ctx.require(param_string_zero.exists(), "missing mandatory stringZero");

    let hash = ctx.utility().hash_blake2b(PARAM_HASH.as_bytes());
    ctx.require(param_hash.value() == hash, "Hash wrong");
    ctx.require(param_int64.value() == 42, "int64 wrong");
    ctx.require(param_int64_zero.value() == 0, "int64-0 wrong");
    ctx.require(param_string.value() == PARAM_STRING, "string wrong");
    ctx.require(param_string_zero.value() == "", "string-0 wrong");
    ctx.require(param_hname.value() == ScHname::new(PARAM_HNAME), "Hname wrong");
    ctx.require(param_hname_zero.value() == ScHname(0), "Hname-0 wrong");
}

pub fn view_test_call_panic_view_epfrom_view(ctx: &ScViewContext) {
    ctx.log("calling testCallPanicViewEPFromView");
    ctx.call_self(HVIEW_TEST_PANIC_VIEW_EP, None);
}

pub fn view_test_chain_owner_idview(ctx: &ScViewContext) {
    ctx.log("calling testChainOwnerIDView");
    ctx.results().get_agent_id(PARAM_CHAIN_OWNER_ID).set_value(&ctx.chain_owner_id())
}

pub fn view_test_contract_idview(ctx: &ScViewContext) {
    ctx.log("calling testContractIDView");
    ctx.results().get_contract_id(PARAM_CONTRACT_ID).set_value(&ctx.contract_id());
}

pub fn view_test_panic_view_ep(ctx: &ScViewContext) {
    ctx.log("calling testPanicViewEP");
    ctx.panic(MSG_VIEW_PANIC)
}

pub fn view_test_sandbox_call(ctx: &ScViewContext) {
    ctx.log("calling testSandboxCall");
    let ret = ctx.call(CORE_ROOT, CORE_ROOT_VIEW_GET_CHAIN_INFO, None);
    let desc = ret.get_string("d").value();
    ctx.results().get_string("sandboxCall").set_value(&desc);
}
