// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use wasmlib::*;

const PARAM_INT_PARAM_NAME: &str = "intParamName";
const PARAM_INT_PARAM_VALUE: &str = "intParamValue";
const PARAM_HNAME_CONTRACT: &str = "hnameContract";
const PARAM_HNAME_EP: &str = "hnameEP";

const PARAM_ADDRESS: &str = "address";
const PARAM_CHAIN_OWNER_ID: &str = "chainOwnerID";
const PARAM_CONTRACT_ID: &str = "contractID";
const PARAM_CHAIN_ID: &str = "chainid";
const PARAM_CALLER: &str = "caller";
const PARAM_AGENT_ID: &str = "agentID";
const PARAM_CREATOR: &str = "contractCreator";

const PARAM_INT64: &str = "int64";
const PARAM_INT64_ZERO: &str = "int64-0";
const PARAM_HASH: &str = "Hash";
const PARAM_HNAME: &str = "Hname";
const PARAM_HNAME_ZERO: &str = "Hname-0";
const PARAM_STRING: &str = "string";
const PARAM_STRING_ZERO: &str = "string-0";

const VAR_COUNTER: &str = "counter";
const VAR_CONTRACT_NAME_DEPLOYED: &str = "exampleDeployTR";

const MSG_FULL_PANIC: &str = "========== panic FULL ENTRY POINT =========";
const MSG_VIEW_PANIC: &str = "========== panic VIEW =========";
const MSG_PANIC_UNAUTHORIZED: &str = "============== panic due to unauthorized call";

#[no_mangle]
fn on_load() {
    let exports = ScExports::new();
    exports.add_func("init", on_init);
    exports.add_func("doNothing", do_nothing);
    exports.add_func("callOnChain", call_on_chain);
    exports.add_func("setInt", set_int);
    exports.add_view("getInt", get_int);
    exports.add_view("fibonacci", fibonacci);
    exports.add_view("getCounter", get_counter);
    exports.add_func("runRecursion", run_recursion);

    exports.add_func("testPanicFullEP", test_panic_full_ep);
    exports.add_view("testPanicViewEP", test_panic_view_ep);
    exports.add_func("testCallPanicFullEP", test_call_panic_full_ep);
    exports.add_func("testCallPanicViewEPFromFull", test_call_panic_view_from_full);
    exports.add_view("testCallPanicViewEPFromView", test_call_panic_view_from_view);

    exports.add_view("testChainOwnerIDView", test_chain_owner_id_view);
    exports.add_func("testChainOwnerIDFull", test_chain_owner_id_full);
    exports.add_view("testContractIDView", test_contract_id_view);
    exports.add_func("testContractIDFull", test_contract_id_full);
    exports.add_view("testSandboxCall", test_sandbox_call);

    exports.add_func("passTypesFull", pass_types_full);
    exports.add_view("passTypesView", pass_types_view);
    exports.add_func("checkContextFromFullEP", check_ctx_from_full);
    exports.add_view("checkContextFromViewEP", check_ctx_from_view);

    exports.add_func("sendToAddress", send_to_address);
    exports.add_view("justView", test_just_view);

    exports.add_func("testEventLogGenericData", test_event_log_generic_data);
    exports.add_func("testEventLogEventData", test_event_log_event_data);
    exports.add_func("testEventLogDeploy", test_event_log_deploy);

    exports.add_func("withdrawToChain", withdraw_to_chain);
}

fn on_init(ctx: &ScFuncContext) {
    ctx.log("testcore.on_init.wasm.begin");
}

fn do_nothing(ctx: &ScFuncContext) {
    ctx.log("testcore.do_nothing.begin");
}

fn set_int(ctx: &ScFuncContext) {
    ctx.log("testcore.set_int.begin");
    let param_name = ctx.params().get_string(PARAM_INT_PARAM_NAME);
    ctx.require(param_name.exists(), "param 'name' not found");

    let param_value = ctx.params().get_int(PARAM_INT_PARAM_VALUE);
    ctx.require(param_value.exists(), "param 'value' not found");

    ctx.state().get_int(&param_name.value() as &str).set_value(param_value.value());
}

fn get_int(ctx: &ScViewContext) {
    ctx.log("testcore.get_int.begin");
    let param_name = ctx.params().get_string(PARAM_INT_PARAM_NAME);
    ctx.require(param_name.exists(), "param 'name' not found");

    let param_value = ctx.state().get_int(&param_name.value() as &str);
    ctx.require(param_value.exists(), "param 'value' not found");

    ctx.results().get_int(&param_name.value() as &str).set_value(param_value.value());
}

fn call_on_chain(ctx: &ScFuncContext) {
    let param_value = ctx.params().get_int(PARAM_INT_PARAM_VALUE);
    ctx.require(param_value.exists(), "param 'value' not found");
    let param_in = param_value.value();

    let mut target_contract = ctx.contract_id().hname();
    let param_hname_contract = ctx.params().get_hname(PARAM_HNAME_CONTRACT);
    if param_hname_contract.exists() {
        target_contract = param_hname_contract.value()
    }

    let mut target_ep = ScHname::new("callOnChain");
    let param_hname_ep = ctx.params().get_hname(PARAM_HNAME_EP);
    if param_hname_ep.exists() {
        target_ep = param_hname_ep.value()
    }

    let var_counter = ctx.state().get_int(VAR_COUNTER);
    let mut counter: i64 = 0;
    if var_counter.exists() {
        counter = var_counter.value();
    }
    var_counter.set_value(counter + 1);

    ctx.log(&format!("call depth = {} hnameContract = {} hnameEP = {} counter = {}",
                     param_in, &target_contract.to_string(), &target_ep.to_string(), counter));

    let par = ScMutableMap::new();
    par.get_int(PARAM_INT_PARAM_VALUE).set_value(param_in);
    let ret = ctx.call(target_contract, target_ep, Some(par), None);

    let ret_val = ret.get_int(PARAM_INT_PARAM_VALUE);

    ctx.results().get_int(PARAM_INT_PARAM_VALUE).set_value(ret_val.value());
}

fn get_counter(ctx: &ScViewContext) {
    ctx.log("testcore.get_counter.begin");
    let counter = ctx.state().get_int(VAR_COUNTER);
    ctx.results().get_int(VAR_COUNTER).set_value(counter.value());
}

fn run_recursion(ctx: &ScFuncContext) {
    let param_value = ctx.params().get_int(PARAM_INT_PARAM_VALUE);
    ctx.require(param_value.exists(), "param no found");
    let depth = param_value.value();
    if depth <= 0 {
        return;
    }
    let par = ScMutableMap::new();
    par.get_int(PARAM_INT_PARAM_VALUE).set_value(depth - 1);
    par.get_hname(PARAM_HNAME_EP).set_value(ScHname::new("runRecursion"));
    ctx.call(ctx.contract_id().hname(), ScHname::new("callOnChain"), Some(par), None);
    // TODO how would I return result of the call ???
    ctx.results().get_int(PARAM_INT_PARAM_VALUE).set_value(depth - 1);
}

fn fibonacci(ctx: &ScViewContext) {
    let n = ctx.params().get_int(PARAM_INT_PARAM_VALUE);
    ctx.require(n.exists(), "param 'value' not found");

    let n = n.value();
    if n == 0 || n == 1 {
        ctx.results().get_int(PARAM_INT_PARAM_VALUE).set_value(n);
        return;
    }
    let params1 = ScMutableMap::new();
    params1.get_int(PARAM_INT_PARAM_VALUE).set_value(n - 1);
    let results1 = ctx.call(ctx.contract_id().hname(), ScHname::new("fibonacci"), Some(params1));
    let n1 = results1.get_int(PARAM_INT_PARAM_VALUE).value();

    let params2 = ScMutableMap::new();
    params2.get_int(PARAM_INT_PARAM_VALUE).set_value(n - 2);
    let results2 = ctx.call(ctx.contract_id().hname(), ScHname::new("fibonacci"), Some(params2));
    let n2 = results2.get_int(PARAM_INT_PARAM_VALUE).value();

    ctx.results().get_int(PARAM_INT_PARAM_VALUE).set_value(n1 + n2);
}

fn test_panic_full_ep(ctx: &ScFuncContext) {
    ctx.panic(MSG_FULL_PANIC)
}

fn test_panic_view_ep(ctx: &ScViewContext) {
    ctx.panic(MSG_VIEW_PANIC)
}

fn test_call_panic_full_ep(ctx: &ScFuncContext) {
    ctx.call(ctx.contract_id().hname(), ScHname::new("testPanicFullEP"), None, None);
}

fn test_call_panic_view_from_full(ctx: &ScFuncContext) {
    ctx.call(ctx.contract_id().hname(), ScHname::new("testPanicViewEP"), None, None);
}

fn test_call_panic_view_from_view(ctx: &ScViewContext) {
    ctx.call(ctx.contract_id().hname(), ScHname::new("testPanicViewEP"), None);
}

fn test_just_view(ctx: &ScViewContext) {
    ctx.log("calling empty view entry point")
}

fn send_to_address(ctx: &ScFuncContext) {
    ctx.log("sendToAddress");
    ctx.require(ctx.caller().equals(&ctx.contract_creator()), MSG_PANIC_UNAUTHORIZED);

    let target_addr = ctx.params().get_address(PARAM_ADDRESS);
    ctx.require(target_addr.exists(), "parameter 'address' not found");

    let my_balances = ctx.balances();
    ctx.transfer_to_address(&target_addr.value(), &my_balances);
}

fn test_chain_owner_id_view(ctx: &ScViewContext) {
    ctx.results().get_agent_id(PARAM_CHAIN_OWNER_ID).set_value(&ctx.chain_owner_id())
}

fn test_chain_owner_id_full(ctx: &ScFuncContext) {
    ctx.results().get_agent_id(PARAM_CHAIN_OWNER_ID).set_value(&ctx.chain_owner_id())
}

fn test_contract_id_view(ctx: &ScViewContext) {
    ctx.results().get_contract_id(PARAM_CONTRACT_ID).set_value(&ctx.contract_id());
}

fn test_contract_id_full(ctx: &ScFuncContext) {
    ctx.results().get_contract_id(PARAM_CONTRACT_ID).set_value(&ctx.contract_id());
}

fn test_sandbox_call(ctx: &ScViewContext) {
    let ret = ctx.call(CORE_ROOT, CORE_ROOT_VIEW_GET_CHAIN_INFO, None);
    let desc = ret.get_string("d").value();
    ctx.results().get_string("sandboxCall").set_value(&desc);
}

fn pass_types_full(ctx: &ScFuncContext) {
    ctx.require(ctx.params().get_int(PARAM_INT64).exists(), "!int64.exist");
    ctx.require(ctx.params().get_int(PARAM_INT64).value() == 42, "int64 wrong");

    ctx.require(ctx.params().get_int(PARAM_INT64_ZERO).exists(), "!int64-0.exist");
    ctx.require(ctx.params().get_int(PARAM_INT64_ZERO).value() == 0, "int64-0 wrong");

    ctx.require(ctx.params().get_string(PARAM_STRING).exists(), "!string.exist");
    ctx.require(ctx.params().get_string(PARAM_STRING).value() == "string", "string wrong");

    ctx.require(ctx.params().get_string(PARAM_STRING_ZERO).exists(), "!string-0.exist");
    ctx.require(ctx.params().get_string(PARAM_STRING_ZERO).value() == "", "string-0 wrong");

    ctx.require(ctx.params().get_hash(PARAM_HASH).exists(), "!Hash.exist");

    let hash = ctx.utility().hash_blake2b("Hash".as_bytes());
    ctx.require(ctx.params().get_hash(PARAM_HASH).value().equals(&hash), "Hash wrong");

    ctx.require(ctx.params().get_hname(PARAM_HNAME).exists(), "!Hname.exist");
    ctx.require(ctx.params().get_hname(PARAM_HNAME).value().equals(ScHname::new("Hname")), "Hname wrong");

    ctx.require(ctx.params().get_hname(PARAM_HNAME_ZERO).exists(), "!Hname-0.exist");
    ctx.require(ctx.params().get_hname(PARAM_HNAME_ZERO).value().equals(ScHname(0)), "Hname-0 wrong");
}

fn pass_types_view(ctx: &ScViewContext) {
    ctx.require(ctx.params().get_int(PARAM_INT64).exists(), "!int64.exist");
    ctx.require(ctx.params().get_int(PARAM_INT64).value() == 42, "int64 wrong");

    ctx.require(ctx.params().get_int(PARAM_INT64_ZERO).exists(), "!int64-0.exist");
    ctx.require(ctx.params().get_int(PARAM_INT64_ZERO).value() == 0, "int64-0 wrong");

    ctx.require(ctx.params().get_string(PARAM_STRING).exists(), "!string.exist");
    ctx.require(ctx.params().get_string(PARAM_STRING).value() == "string", "string wrong");

    ctx.require(ctx.params().get_string(PARAM_STRING_ZERO).exists(), "!string-0.exist");
    ctx.require(ctx.params().get_string(PARAM_STRING_ZERO).value() == "", "string-0 wrong");

    ctx.require(ctx.params().get_hash(PARAM_HASH).exists(), "!Hash.exist");

    let hash = ctx.utility().hash_blake2b("Hash".as_bytes());
    ctx.require(ctx.params().get_hash(PARAM_HASH).value().equals(&hash), "Hash wrong");

    ctx.require(ctx.params().get_hname(PARAM_HNAME).exists(), "!Hname.exist");
    ctx.require(ctx.params().get_hname(PARAM_HNAME).value().equals(ScHname::new("Hname")), "Hname wrong");

    ctx.require(ctx.params().get_hname(PARAM_HNAME_ZERO).exists(), "!Hname-0.exist");
    ctx.require(ctx.params().get_hname(PARAM_HNAME_ZERO).value().equals(ScHname(0)), "Hname-0 wrong");
}

fn check_ctx_from_full(ctx: &ScFuncContext) {
    let par = ctx.params();

    let chain_id = par.get_chain_id(PARAM_CHAIN_ID);
    ctx.require(chain_id.exists() && chain_id.value() == ctx.contract_id().chain_id(), "fail: chainID");

    let chain_owner_id = par.get_agent_id(PARAM_CHAIN_OWNER_ID);
    ctx.require(chain_owner_id.exists() && chain_owner_id.value() == ctx.chain_owner_id(), "fail: chainOwnerID");

    let caller = par.get_agent_id(PARAM_CALLER);
    ctx.require(caller.exists() && caller.value() == ctx.caller(), "fail: caller");

    let contract_id = par.get_contract_id(PARAM_CONTRACT_ID);
    ctx.require(contract_id.exists() && contract_id.value() == ctx.contract_id(), "fail: contractID");

    let agent_id = par.get_agent_id(PARAM_AGENT_ID);
    let as_agent_id = ctx.contract_id().as_agent_id();
    ctx.require(agent_id.exists() && agent_id.value() == as_agent_id, "fail: agentID");

    let creator = par.get_agent_id(PARAM_CREATOR);
    ctx.require(creator.exists() && creator.value() == ctx.contract_creator(), "fail: contractCreator");
}

fn check_ctx_from_view(ctx: &ScViewContext) {
    let par = ctx.params();

    let chain_id = par.get_chain_id(PARAM_CHAIN_ID);
    ctx.require(chain_id.exists() && chain_id.value() == ctx.contract_id().chain_id(), "fail: chainID");

    let chain_owner_id = par.get_agent_id(PARAM_CHAIN_OWNER_ID);
    ctx.require(chain_owner_id.exists() && chain_owner_id.value() == ctx.chain_owner_id(), "fail: chainOwnerID");

    let contract_id = par.get_contract_id(PARAM_CONTRACT_ID);
    ctx.require(contract_id.exists() && contract_id.value() == ctx.contract_id(), "fail: contractID");

    let agent_id = par.get_agent_id(PARAM_AGENT_ID);
    let as_agent_id = ctx.contract_id().as_agent_id();
    ctx.require(agent_id.exists() && agent_id.value() == as_agent_id, "fail: agentID");

    let creator = par.get_agent_id(PARAM_CREATOR);
    ctx.require(creator.exists() && creator.value() == ctx.contract_creator(), "fail: contractCreator");
}

fn test_event_log_generic_data(ctx: &ScFuncContext) {
    let counter = ctx.params().get_int(VAR_COUNTER);
    ctx.require(counter.exists(), "!counter.exist");
    let event = "[GenericData] Counter Number: ".to_string() + &counter.to_string();
    ctx.event(&event)
}

fn test_event_log_event_data(ctx: &ScFuncContext) {
    ctx.event("[Event] - Testing Event...");
}

fn test_event_log_deploy(ctx: &ScFuncContext) {
    //Deploy the same contract with another name
    let program_hash = ctx.utility().hash_blake2b("test_sandbox".as_bytes());
    ctx.deploy(&program_hash, VAR_CONTRACT_NAME_DEPLOYED,
               "test contract deploy log", None)
}

fn withdraw_to_chain(ctx: &ScFuncContext) {
    //Deploy the same contract with another name
    let target_chain = ctx.params().get_chain_id(PARAM_CHAIN_ID);
    ctx.require(target_chain.exists(), "chainID not provided");

    let target_contract_id = ScContractId::new(&target_chain.value(), &CORE_ACCOUNTS);
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

