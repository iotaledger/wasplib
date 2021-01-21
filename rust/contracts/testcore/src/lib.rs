// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use wasplib::client::*;

const PARAM_INT_PARAM_NAME: &str = "intParamName";
const PARAM_INT_PARAM_VALUE: &str = "intParamValue";
// const PARAM_HNAME: &str = "hname";
// const PARAM_CALL_OPTION: &str = "callOption";
const PARAM_ADDRESS: &str = "address";
const PARAM_CHAIN_OWNER: &str = "chainOwner";
const PARAM_CONTRACT_ID: &str = "contractID";

const MSG_FULL_PANIC: &str = "========== panic FULL ENTRY POINT =========";
const MSG_VIEW_PANIC: &str = "========== panic VIEW =========";
const MSG_PANIC_UNAUTHORIZED: &str = "============== panic due to unauthorized call";

const SELF_NAME: &str = "test_sandbox";   // temporary, until hname in the call will become available

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

    exports.add_call("sendToAddress", send_to_address);
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
    ctx.log("testcore.call_on_chain.begin");
    //
    // let param_call_option = ctx.params().get_string(PARAM_CALL_OPTION);
    // let param_value = ctx.params().get_int(PARAM_INT_PARAM_VALUE);
    // if !param_value.exists(){
    //     ctx.panic("param value not found")
    // }
    // TODO cannot get hname type
}

fn fibonacci(ctx: &ScViewContext) {
    let n = ctx.params().get_int(PARAM_INT_PARAM_VALUE);
    if !n.exists() {
        ctx.panic("param value not found")
    }
    let n = n.value();
    ctx.log(&("fibonacci: ".to_string() + &n.to_string()));
    if n == 0 || n == 1 {
        ctx.log("return 1");
        ctx.results().get_int(PARAM_INT_PARAM_VALUE).set_value(n);
        return;
    }
    ctx.log("before call 1");
    let params1 = ScMutableMap::new();
    params1.get_int(PARAM_INT_PARAM_VALUE).set_value(n - 1);
    let results1 = ctx.call(Hname::SELF, Hname::new("fibonacci"), params1);
    let n1 = results1.get_int(PARAM_INT_PARAM_VALUE).value();
    ctx.log(&("    fibonacci-1: ".to_string() + &n1.to_string()));

    let params2 = ScMutableMap::new();
    params2.get_int(PARAM_INT_PARAM_VALUE).set_value(n - 2);
    let results2 = ctx.call(Hname::SELF, Hname::new("fibonacci"), params2);
    let n2 = results2.get_int(PARAM_INT_PARAM_VALUE).value();
    ctx.log(&("    fibonacci-2: ".to_string() + &n2.to_string()));

    ctx.results().get_int(PARAM_INT_PARAM_VALUE).set_value(n1 + n2);
}

fn test_panic_full_ep(ctx: &ScCallContext) {
    ctx.panic(MSG_FULL_PANIC)
}

fn test_panic_view_ep(ctx: &ScViewContext) {
    ctx.panic(MSG_VIEW_PANIC)
}

fn test_call_panic_full_ep(ctx: &ScCallContext) {
    ctx.call(Hname::SELF, Hname::new("testPanicFullEP"), ScMutableMap::NONE, ScTransfers::NONE);
}

// FIXME no need for 'view method special'
fn test_call_panic_view_from_full(ctx: &ScCallContext) {
    ctx.call(Hname::SELF, Hname::new("testPanicViewEP"), ScMutableMap::NONE, ScTransfers::NONE);
}

// FIXME no need for 'view method special'
fn test_call_panic_view_from_view(ctx: &ScViewContext) {
    ctx.call(Hname::SELF, Hname::new("testPanicViewEP"), ScMutableMap::NONE);
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
    // let mybalances = ctx.balances();
    // TODO now way of knowing if balances are empty
    // how to transfer all balances
    // ctx.transfer_to_address(&targetAddr.value()).transfer(mybalances).send();
}

fn test_chain_owner_id_view(ctx: &ScViewContext) {
    ctx.results().get_agent(PARAM_CHAIN_OWNER).set_value(&ctx.chain_owner())
}

fn test_chain_owner_id_full(ctx: &ScCallContext) {
    ctx.results().get_agent(PARAM_CHAIN_OWNER).set_value(&ctx.chain_owner())
}

fn test_contract_id_view(_ctx: &ScViewContext) {
    // TODO there's no way to return contact ID
    // ctx.results().(PARAM_CONTRACT_ID).set_value(ctx.chain_owner().value)
}

fn test_contract_id_full(_ctx: &ScCallContext) {
}

fn test_sandbox_call(ctx: &ScViewContext) {
    let ret = ctx.call(CORE_ROOT, VIEW_GET_CHAIN_INFO, ScMutableMap::NONE);
    let desc = ret.get_string("d").value();
    ctx.results().get_string("sandboxCall").set_value(&desc);
}
