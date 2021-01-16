// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use wasplib::client::*;

const PARAM_INT_PARAM_NAME: &str = "intParamName";
const PARAM_INT_PARAM_VALUE: &str = "intParamValue";
// const PARAM_HNAME: &str = "hname";
// const PARAM_CALL_OPTION: &str = "callOption";

const MSG_FULL_PANIC: &str = "========== panic FULL ENTRY POINT =========";
const MSG_VIEW_PANIC: &str = "========== panic VIEW =========";

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
    if !param_name.exists(){
        ctx.panic("param name not found")
    }
    let param_value = ctx.params().get_int(PARAM_INT_PARAM_VALUE);
    if !param_value.exists(){
        ctx.panic("param value not found")
    }
    ctx.state().get_int(&param_name.value() as &str).set_value(param_value.value());
}

fn get_int(ctx: &ScViewContext) {
    ctx.log("testcore.get_int.begin");
    let param_name = ctx.params().get_string(PARAM_INT_PARAM_NAME);
    if !param_name.exists(){
        ctx.panic("param name not found")
    }
    let param_value = ctx.state().get_int(&param_name.value() as &str);
    if !param_value.exists(){
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
    if !n.exists(){
        ctx.panic("param value not found")
    }
    let n = n.value();
    ctx.log(&("fibonacci: ".to_string() + &n.to_string()));
    if n == 0 || n == 1{
        ctx.log("return 1");
        ctx.results().get_int(PARAM_INT_PARAM_VALUE).set_value(n);
        return;
    }
    ctx.log("before call 1");
    let view_call1 = ctx.view("fibonacci");
    view_call1.contract(SELF_NAME);
    view_call1.params().get_int(PARAM_INT_PARAM_VALUE).set_value(n-1);
    view_call1.view();
    let n1 = view_call1.results().get_int(PARAM_INT_PARAM_VALUE).value();
    ctx.log(&("    fibonacci-1: ".to_string() + &n1.to_string()));

    let view_call2 = ctx.view("fibonacci");
    view_call2.contract(SELF_NAME);
    view_call2.params().get_int(PARAM_INT_PARAM_VALUE).set_value(n-2);
    view_call2.view();
    let n2 = view_call2.results().get_int(PARAM_INT_PARAM_VALUE).value();
    ctx.log(&("    fibonacci-2: ".to_string() + &n2.to_string()));

    ctx.results().get_int(PARAM_INT_PARAM_VALUE).set_value(n1+n2);
}

fn test_panic_full_ep(ctx: &ScCallContext){
    ctx.panic(MSG_FULL_PANIC)
}

fn test_panic_view_ep(ctx: &ScViewContext){
    ctx.panic(MSG_VIEW_PANIC)
}

fn test_call_panic_full_ep(ctx: &ScCallContext){
    ctx.call("testPanicFullEP").contract(SELF_NAME).call();  // FIXME need self.hname
}

// FIXME no need for 'view method special'
fn test_call_panic_view_from_full(ctx: &ScCallContext){
    ctx.view("testPanicViewEP").contract(SELF_NAME).view();  // FIXME need self.hname
}

// FIXME no need for 'view method special'
fn test_call_panic_view_from_view(ctx: &ScViewContext){
    ctx.view("testPanicViewEP").contract(SELF_NAME).view();
}
