// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use wasplib::client::*;

const PARAM_INT_PARAM_NAME: &str = "intParamName";
const PARAM_INT_PARAM_VALUE: &str = "intParamValue";
// const PARAM_HNAME: &str = "hname";
// const PARAM_CALL_OPTION: &str = "callOption";


#[no_mangle]
fn on_load() {
    let exports = ScExports::new();
    exports.add_call("init", on_init);
    exports.add_call("doNothing", do_nothing);
    exports.add_call("callOnChain", call_on_chain);
    exports.add_call("setInt", set_int);
    exports.add_view("getInt", get_int);
    exports.add_view("fibonacci", fibonacci);
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
    ctx.log("testcore.fibonacci.begin");
    //
    // let n = ctx.params().get_int(PARAM_INT_PARAM_VALUE);
    // if !n.exists(){
    //     ctx.panic("param value not found")
    // }
    // let n = n.value();
    // if n == 0 || n == 1{
    //     ctx.results().get_int(PARAM_INT_PARAM_VALUE).set_value(n);
    //     return;
    // }
}
