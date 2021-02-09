// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use wasmlib::*;

use crate::*;

// storeString entry point stores a string provided as parameters
// in the state as a value of the key 'storedString'
// panics if parameter is not provided
pub fn func_store_string(ctx: &ScFuncContext, params: &FuncStoreStringParams) {
    // take parameter paramString
    let param_string = params.string.value();

    // store the string in "storedString" variable
    ctx.state().get_string(VAR_STRING).set_value(&param_string);
    // log the text
    let msg = "Message stored: ".to_string() + &param_string;
    ctx.log(&msg);
}

// withdraw_iota sends all iotas contained in the contract's account
// to the caller's L1 address.
// Panics of the caller is not an address
// Panics if the address is not the creator of the contract is the caller
// The caller will be address only if request is sent from the wallet on the L1, not a smart contract
pub fn func_withdraw_iota(ctx: &ScFuncContext, _params: &FuncWithdrawIotaParams) {
    let caller = ctx.caller();
    ctx.require(caller.is_address(), "caller must be an address");

    let bal = ctx.balances().balance(&ScColor::IOTA);
    if bal > 0 {
        ctx.transfer_to_address(&caller.address(), &ScTransfers::new(&ScColor::IOTA, bal))
    }
}

// getString view returns the string value of the key 'storedString'
// The call return result as a key/value dictionary.
// the returned value in the result is under key 'paramString'
pub fn view_get_string(ctx: &ScViewContext, _params: &ViewGetStringParams) {
    // take the stored string
    let s = ctx.state().get_string(VAR_STRING).value();
    // return the string value in the result dictionary
    ctx.results().get_string(PARAM_STRING).set_value(&s);
}
