// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use wasplib::client::*;

const PARAM_INT_PARAM_NAME: &str = "intParamName";
const PARAM_INT_PARAM_VALUE: &str = "intParamValue";
const PARAM_HNAME: &str = "hname";
const PARAM_CALL_OPTION: &str = "callOption";
const PARAM_ADDRESS: &str = "address";
const PARAM_CHAIN_OWNER: &str = "chainOwner";
const PARAM_CONTRACT_ID: &str = "contractID";

const VAR_COUNTER: &str = "counter";

const MSG_FULL_PANIC: &str = "========== panic FULL ENTRY POINT =========";
const MSG_VIEW_PANIC: &str = "========== panic VIEW =========";
const MSG_PANIC_UNAUTHORIZED: &str = "============== panic due to unauthorized call";

const CALL_OPTION_FORWARD: &str = "forward";

#[no_mangle]
fn on_load() {
    let exports = ScExports::new();
    exports.add_call("init", on_init);
    exports.add_call("doNothing", do_nothing);
    exports.add_call("callOnChain", call_on_chain);
    exports.add_call("setInt", set_int);
    exports.add_view("getInt", get_int);
    exports.add_view("fibonacci", fibonacci);

    exports.add_call("testPanicFullEP", test_panic_full_ep);
    exports.add_view("testPanicViewEP", test_panic_view_ep);
    exports.add_call("testCallPanicFullEP", test_call_panic_full_ep);
    exports.add_call("testCallPanicViewEPFromFull", test_call_panic_view_from_full);
    exports.add_view("testCallPanicViewEPFromView", test_call_panic_view_from_view);

    exports.add_view("testChainOwnerIDView", test_chain_owner_id_view);
    exports.add_call("testChainOwnerIDFull", test_chain_owner_id_full);
    exports.add_view("testContractIDView", test_contract_id_view);
    exports.add_call("testContractIDFull", test_contract_id_full);
    exports.add_view("testSandboxCall", test_sandbox_call);

    exports.add_call("passTypesFull", pass_types_full);
    exports.add_view("passTypesView", pass_types_view);

    exports.add_call("sendToAddress", send_to_address);
    exports.add_view("justView", test_just_view);
}

fn on_init(ctx: &ScCallContext) {
    ctx.log("testcore.on_init.wasm.begin");
}

fn do_nothing(ctx: &ScCallContext) {
    ctx.log("testcore.do_nothing.begin");
}

fn set_int(ctx: &ScCallContext) {
    ctx.log("testcore.set_int.begin");
    let param_name = ctx.params().get_string(PARAM_INT_PARAM_NAME);
    if !param_name.exists() {
        ctx.panic("param name not found")
    }
    let param_value = ctx.params().get_int(PARAM_INT_PARAM_VALUE);
    if !param_value.exists() {
        ctx.panic("param value not found")
    }
    ctx.state().get_int(&param_name.value() as &str).set_value(param_value.value());
}

fn get_int(ctx: &ScViewContext) {
    ctx.log("testcore.get_int.begin");
    let param_name = ctx.params().get_string(PARAM_INT_PARAM_NAME);
    if !param_name.exists() {
        ctx.panic("param name not found")
    }
    let param_value = ctx.state().get_int(&param_name.value() as &str);
    if !param_value.exists() {
        ctx.panic("param value is not in state")
    }
    ctx.results().get_int(&param_name.value() as &str).set_value(param_value.value());
}

fn call_on_chain(ctx: &ScCallContext) {
    let param_call_option = ctx.params().get_string(PARAM_CALL_OPTION);
    if !param_call_option.exists() {
        ctx.panic("'callOption' not specified")
    }
    let call_option = param_call_option.value();

    let param_value = ctx.params().get_int(PARAM_INT_PARAM_VALUE);
    if !param_value.exists() {
        ctx.panic("param value not found")
    }
    let mut call_depth = param_value.value();

    let mut target = Hname::SELF;
    let param_hname = ctx.params().get_hname(PARAM_HNAME);
    if param_hname.exists() {
        target = param_hname.value();
    }

    let var_counter = ctx.state().get_int(VAR_COUNTER);
    let mut counter: i64 = 0;
    if var_counter.exists() {
        counter = var_counter.value();
    }

    // TODO ctx.contract_id() ContactID is not an AgentID type.
    //  should be

    ctx.log(&format!("call depth = {} option = '{}' hname = {} counter = {}",
                     call_depth, call_option, &target.to_string(), counter));


    if call_depth <= 0 {
        ctx.results().get_int(VAR_COUNTER).set_value(var_counter.value());
        return;
    }

    var_counter.set_value(counter + 1);
    call_depth = call_depth - 1;
    if call_option == CALL_OPTION_FORWARD {
        let par = ScMutableMap::new();
        par.get_string(PARAM_CALL_OPTION).set_value(CALL_OPTION_FORWARD);
        par.get_int(PARAM_INT_PARAM_VALUE).set_value(call_depth);
        let ret = ctx.call(target, Hname::new("callOnChain"), par, &ScTransfers::NONE);
        ctx.results().get_int(VAR_COUNTER).set_value(ret.get_int(VAR_COUNTER).value());
    } else {
        ctx.panic("unknown call option")
    }
}


fn fibonacci(ctx: &ScViewContext) {
    let n = ctx.params().get_int(PARAM_INT_PARAM_VALUE);
    if !n.exists() {
        ctx.panic("param value not found")
    }
    let n = n.value();
    // ctx.log(&("fibonacci: ".to_string() + &n.to_string()));
    if n == 0 || n == 1 {
        ctx.results().get_int(PARAM_INT_PARAM_VALUE).set_value(n);
        return;
    }
    let params1 = ScMutableMap::new();
    params1.get_int(PARAM_INT_PARAM_VALUE).set_value(n - 1);
    let results1 = ctx.call(Hname::SELF, Hname::new("fibonacci"), params1);
    let n1 = results1.get_int(PARAM_INT_PARAM_VALUE).value();
    // ctx.log(&("    fibonacci-1: ".to_string() + &n1.to_string()));

    let params2 = ScMutableMap::new();
    params2.get_int(PARAM_INT_PARAM_VALUE).set_value(n - 2);
    let results2 = ctx.call(Hname::SELF, Hname::new("fibonacci"), params2);
    let n2 = results2.get_int(PARAM_INT_PARAM_VALUE).value();
    // ctx.log(&("    fibonacci-2: ".to_string() + &n2.to_string()));

    ctx.results().get_int(PARAM_INT_PARAM_VALUE).set_value(n1 + n2);
}

fn test_panic_full_ep(ctx: &ScCallContext) {
    ctx.panic(MSG_FULL_PANIC)
}

fn test_panic_view_ep(ctx: &ScViewContext) {
    ctx.panic(MSG_VIEW_PANIC)
}

fn test_call_panic_full_ep(ctx: &ScCallContext) {
    ctx.call(Hname::SELF, Hname::new("testPanicFullEP"), ScMutableMap::NONE, &ScTransfers::NONE);
}

// FIXME no need for 'view method special'
fn test_call_panic_view_from_full(ctx: &ScCallContext) {
    ctx.call(Hname::SELF, Hname::new("testPanicViewEP"), ScMutableMap::NONE, &ScTransfers::NONE);
}

// FIXME no need for 'view method special'
fn test_call_panic_view_from_view(ctx: &ScViewContext) {
    ctx.call(Hname::SELF, Hname::new("testPanicViewEP"), ScMutableMap::NONE);
}

fn test_just_view(ctx: &ScViewContext) {
    ctx.log("calling empty view entry point")
}

fn send_to_address(ctx: &ScCallContext) {
    ctx.log("sendToAddress");
    if !ctx.caller().equals(&ctx.contract_creator()) {
        ctx.panic(MSG_PANIC_UNAUTHORIZED);
    }
    let target_addr = ctx.params().get_address(PARAM_ADDRESS);
    if !target_addr.exists() {
        ctx.panic("parameter 'address' not provided")
    }
    let my_balances = ctx.balances();
    ctx.transfer_to_address(&target_addr.value(), &my_balances);
}

fn test_chain_owner_id_view(ctx: &ScViewContext) {
    ctx.results().get_agent(PARAM_CHAIN_OWNER).set_value(&ctx.chain_owner())
}

fn test_chain_owner_id_full(ctx: &ScCallContext) {
    ctx.results().get_agent(PARAM_CHAIN_OWNER).set_value(&ctx.chain_owner())
}

fn test_contract_id_view(ctx: &ScViewContext) {
    //TODO discussion about using ChainID vs ContractID because one of those seems redundant
    ctx.results().get_agent(PARAM_CONTRACT_ID).set_value(&ctx.contract_id());
    // alternatively do not use agent but bytes instead for now:
    // ctx.results().get_bytes(PARAM_CONTRACT_ID).set_value(ctx.contract_id().to_bytes());
}

fn test_contract_id_full(ctx: &ScCallContext) {
    ctx.results().get_agent(PARAM_CONTRACT_ID).set_value(&ctx.contract_id());
    // alternatively do not use agent but bytes instead for now:
    // ctx.results().get_bytes(PARAM_CONTRACT_ID).set_value(ctx.contract_id().to_bytes());
}

fn test_sandbox_call(ctx: &ScViewContext) {
    let ret = ctx.call(CORE_ROOT, VIEW_GET_CHAIN_INFO, ScMutableMap::NONE);
    let desc = ret.get_string("d").value();
    ctx.results().get_string("sandboxCall").set_value(&desc);
}

fn pass_types_full(ctx: &ScCallContext) {
    if !ctx.params().get_int("int64").exists() {
        ctx.panic("!int64.exist")
    }
    if ctx.params().get_int("int64").value() != 42 {
        ctx.panic("int64 wrong")
    }
    if !ctx.params().get_int("int64-0").exists() {
        ctx.panic("!int64-0.exist")
    }
    if ctx.params().get_int("int64-0").value() != 0 {
        ctx.panic("int64-0 wrong")
    }
    if !ctx.params().get_hash("Hash").exists() {
        ctx.panic("!Hash.exist")
    }
    let hash = ctx.utility().hash("Hash".as_bytes());
    if !ctx.params().get_hash("Hash").value().equals(&hash) {
        ctx.panic("Hash wrong")
    }
    if !ctx.params().get_hname("Hname").exists() {
        ctx.panic("!Hname.exist")
    }
    if !ctx.params().get_hname("Hname").value().equals(Hname::new("Hname")) {
        ctx.panic("Hname wrong")
    }
    if !ctx.params().get_hname("Hname-0").exists() {
        ctx.panic("!Hname-0.exist")
    }
    if !ctx.params().get_hname("Hname-0").value().equals(Hname(0)) {
        ctx.panic("Hname-0 wrong")
    }
}

fn pass_types_view(ctx: &ScViewContext) {
    if !ctx.params().get_int("int64").exists() {
        ctx.panic("!int64.exist")
    }
    if ctx.params().get_int("int64").value() != 42 {
        ctx.panic("int64 wrong")
    }
    if !ctx.params().get_int("int64-0").exists() {
        ctx.panic("!int64-0.exist")
    }
    if ctx.params().get_int("int64-0").value() != 0 {
        ctx.panic("int64-0 wrong")
    }
    if !ctx.params().get_hash("Hash").exists() {
        ctx.panic("!Hash.exist")
    }
    let hash = ctx.utility().hash("Hash".as_bytes());
    if !ctx.params().get_hash("Hash").value().equals(&hash) {
        ctx.panic("Hash wrong")
    }
    if !ctx.params().get_hname("Hname").exists() {
        ctx.panic("!Hname.exist")
    }
    if !ctx.params().get_hname("Hname").value().equals(Hname::new("Hname")) {
        ctx.panic("Hname wrong")
    }
    if !ctx.params().get_hname("Hname-0").exists() {
        ctx.panic("!Hname-0.exist")
    }
    if !ctx.params().get_hname("Hname-0").value().equals(Hname(0)) {
        ctx.panic("Hname-0 wrong")
    }
}
