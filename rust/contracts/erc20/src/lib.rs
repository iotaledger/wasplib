#![allow(dead_code)]
#![allow(non_snake_case)]

use wasplib::client::ScContext;
use wasplib::client::ScExports;

const VAR_SUPPLY: &str = "s";
const VAR_BALANCES: &str = "b";
const VAR_TARGET_ADDRESS: &str = "addr";
const VAR_AMOUNT: &str = "amount";

#[no_mangle]
pub fn onLoad() {
    let mut exports = ScExports::new();
    exports.add_protected("initSC");
    exports.add("transfer");
    exports.add("approve")
}

#[no_mangle]
pub fn initSC() {
    let sc = ScContext::new();
    sc.log("initSC");

    let state = sc.state();
    let supplyState = state.get_int(VAR_SUPPLY);
    if supplyState.value() > 0 {
        // already initialized
        sc.log("initSC.fail: already initialized");
        return;
    }
    let params = sc.request().params();
    let supplyParam = params.get_int(VAR_SUPPLY);
    if supplyParam.value() == 0 {
        sc.log("initSC.fail: wrong 'supply' parameter");
        return;
    }
    let supply = supplyParam.value();
    supplyState.set_value(supply);
    state.get_map(VAR_BALANCES).get_int(&sc.contract().owner()).set_value(supply);

    sc.log("initSC: success");
}

#[no_mangle]
pub fn transfer() {
    let sc = ScContext::new();
    sc.log("transfer");

    let state = sc.state();
    let request = sc.request();
    let balances = state.get_map(VAR_BALANCES);

    let sender = request.address();

    sc.log(&("sender address: ".to_string() + &sender));

    let source_balance = balances.get_int(&sender);

    sc.log(&("source balance: ".to_string() + &source_balance.value().to_string()));

    let params = request.params();
    let amount = params.get_int(VAR_AMOUNT);

    if amount.value() == 0 {
        sc.log("transfer.fail: wrong 'amount' parameter");
        return;
    }
    if amount.value() > source_balance.value() {
        sc.log("transfer.fail: not enough balance");
        return;
    }
    let target_addr = params.get_string(VAR_TARGET_ADDRESS);
    // TODO check if it is a correct address, otherwise won't be possible to transfer from it

    let target_balance = balances.get_int(&target_addr.value());

    target_balance.set_value(target_balance.value() + amount.value());
    source_balance.set_value(source_balance.value() - amount.value());

    sc.log("transfer: success");
}

#[no_mangle]
pub fn approve() {
    let sc = ScContext::new();
    // TODO
    sc.log("approve");
}

