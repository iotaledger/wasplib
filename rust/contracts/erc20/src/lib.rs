// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

#![allow(dead_code)]
#![allow(non_snake_case)]

use wasplib::client::*;

// state vars
const STATE_VAR_SUPPLY: &str = "s";
const STATE_VAR_BALANCES: &str = "b";
const STATE_VAR_APPROVALS: &str = "a";
// params and return variables

const PARAM_SUPPLY: &str = "supply";
const PARAM_CREATOR: &str = "creator";
const PARAM_ACCOUNT: &str = "acc";
const PARAM_DELEGATION: &str = "delegation";
const PARAM_AMOUNT: &str = "amount";
const PARAM_RECIPIENT: &str = "rec";

#[no_mangle]
pub fn onLoad() {
    let exports = ScExports::new();
    exports.add_call("init", init);
    exports.add_view("totalSupply", total_supply);
    exports.add_view("balanceOf", balance_of);
    exports.add_view("allowance", allowance);
    exports.add_call("transfer", transfer);
    exports.add_call("approve", approve);
    exports.add_call("transferFrom", transfer_from);
}

// init is a constructor entry point. It initializes the smart contract with the initial value of the token supply
// - input:
//   -- PARAM_SUPPLY must be nonzero positive integer
//   -- PARAM_CREATOR is the AgentID where initial supply is placed
// TODO would be awesome to have some less syntactically cumbersome way to check and validate parameters.
fn init(call: &ScCallContext) {
    call.log("erc20.init. begin");
    let supply = call.request().params().get_int(PARAM_SUPPLY);
    if !supply.exists() || supply.value() <= 0{
        call.log("er20.init.fail: wrong 'supply' parameter");
        return;
    }
    let creator = call.request().params().get_agent(PARAM_CREATOR);
    if !creator.exists(){
        call.log("er20.init.fail: wrong 'creator' parameter");
        return;
    }
    call.state().get_int(STATE_VAR_SUPPLY).set_value(supply.value());

    // assign the whole supply to creator
    // we cannot use 'caller' here because the init is always called from the 'root'
    call.state().get_key_map(STATE_VAR_BALANCES).get_int(creator.value().to_bytes()).set_value(supply.value());
    call.log(&("init.success. Supply = ".to_string() + &supply.value().to_string()));
    call.log(&("init.success. Owner = ".to_string() + &creator.value().to_string()));
}

// the view returns total supply set when creating the contract
// - Output:
//   -- PARAM_SUPPLY: i64
fn total_supply(call: &ScViewContext){
    let supply = call.state().get_int(STATE_VAR_SUPPLY).value();
    call.results().get_int(PARAM_SUPPLY).set_value(supply);
}

// the view returns balance of the token held in the account
// - Input:
//    -- PARAM_ACCOUNT: agentID
fn balance_of(call: &ScViewContext){
    let params = call.request().params();
    let account = params.get_agent(PARAM_ACCOUNT);
    if !account.exists(){
        call.log(&("wrong or non existing parameter: ".to_string() + &account.value().to_string()));
        return;
    }
    let balances = call.state().get_key_map(STATE_VAR_BALANCES);
    let balance = balances.get_int(account.value().to_bytes()).value();  // 0 if doesn't exist
    call.results().get_int(PARAM_AMOUNT).set_value(balance)
}

// the view returns max number of tokens the owner of the account allowed to retrieve to another party
// Input:
// - PARAM_ACCOUNT: agentID
// - PARAM_DELEGATION: agentID
// Output:
// - PARAM_AMOUNT: i64. 0 if delegation doesn't exists
fn allowance(call: &ScViewContext){
    call.log("erc20.allowance");
    let owner = call.request().params().get_agent(PARAM_ACCOUNT);
    if !owner.exists(){
        call.log("er20.allowance.fail: wrong 'account' parameter");
        return;
    }
    let delegation = call.request().params().get_agent(PARAM_DELEGATION);
    if !delegation.exists(){
        call.log("er20.allowance.fail: wrong 'delegation' parameter");
        return;
    }
    let allowances = call.state().get_key_map(&(owner.value().to_string()+"_allow"));
    let allow = allowances.get_int(delegation.value().to_bytes()).value();
    call.results().get_int(PARAM_AMOUNT).set_value(allow);
}

// transfer moves tokens from caller's account to target account
// Input:
// - PARAM_ACCOUNT: agentID
// - PARAM_AMOUNT: i64
fn transfer(call: &ScCallContext) {
    call.log("erc20.transfer");

    let request = call.request();
    let params = request.params();

    let source_addr = request.sender();   // cannot take immutable agent
    let target_addr = params.get_agent(PARAM_ACCOUNT);
    if !target_addr.exists(){
        call.log("er20.transfer.fail: wrong 'account' parameter");
        return;
    }
    let target_addr = target_addr.value();
    let amount = params.get_int(PARAM_AMOUNT).value();
    if amount <= 0 {
        call.log("erc20.transfer.fail: wrong 'amount' parameter");
        return;
    }
    let balances = call.state().get_key_map(STATE_VAR_BALANCES);
    let source_balance = balances.get_int(source_addr.to_bytes());
    if source_balance.value() < amount{
        call.log("erc20.transfer.fail: not enough funds");
        return;
    }
    let target_balance = balances.get_int(target_addr.to_bytes());
    let result = target_balance.value() + amount;
    if result <= 0{
        call.log("erc20.transfer.fail: overflow");
        return;
    }
    source_balance.set_value(source_balance.value() - amount);
    target_balance.set_value(target_balance.value() + amount);
    call.log("erc20.transfer.success");
}

// Sets the allowance value for delegated account
// inputs:
//  - PARAM_DELEGATION: agentID
//  - PARAM_AMOUNT: i64
fn approve(call: &ScCallContext) {
    call.log("erc20.approve");

    let account = call.request().params().get_agent(PARAM_DELEGATION);
    if !account.exists(){
        call.log("erc20.approve.fail: wrong 'delegation' parameter");
        return;
    }
    let account = account.value();
    let amount = call.request().params().get_int(PARAM_AMOUNT).value();
    if amount <= 0 {
        call.log("erc20.approve.fail: wrong 'amount' parameter");
        return;
    }
    let caller = call.request().sender().to_string();
    let allowances = call.state().get_key_map(&(caller + "_allow"));
    allowances.get_int(account.to_bytes()).set_value(amount);
    call.log("erc20.approve.success");
}

// Moves the amount of tokens from sender to recipient using the allowance mechanism.
// Amount is then deducted from the caller’s allowance. This function emits the Transfer event.
// Input:
// - PARAM_ACCOUNT: agentID   the spender
// - PARAM_RECIPIENT: agentID   the target
// - PARAM_AMOUNT: i64
fn transfer_from(call: &ScCallContext) {
    call.log("erc20.transfer_from");

    let account = call.request().params().get_agent(PARAM_ACCOUNT);
    if !account.exists(){
        call.log("erc20.approve.fail: wrong 'account' parameter");
        return;
    }
    let account = account.value();
    let recipient = call.request().params().get_agent(PARAM_RECIPIENT);
    if !recipient.exists(){
        call.log("erc20.approve.fail: wrong 'recipient' parameter");
        return;
    }
    let recipient = recipient.value();
    let amount = call.request().params().get_int(PARAM_AMOUNT);
    if !amount.exists(){
        call.log("erc20.approve.fail: wrong 'amount' parameter");
        return;
    }
    let amount = amount.value();

    let allowances = call.state().get_key_map(&(account.to_string()+"_allow"));
    let allowance = allowances.get_int(recipient.to_bytes());
    if allowance.value() < amount{
        call.log("erc20.approve.fail: not enough allowance");
    }

    let balances = call.state().get_key_map(STATE_VAR_BALANCES);
    let source_balance = balances.get_int(account.to_bytes());
    if source_balance.value() < amount{
        call.log("erc20.transfer.fail: not enough funds");
        return;
    }
    let recipient_balance = balances.get_int(recipient.to_bytes());
    let result = recipient_balance.value() + amount;
    if result <= 0{
        call.log("erc20.transfer.fail: overflow");
        return;
    }
    source_balance.set_value(source_balance.value() - amount);
    recipient_balance.set_value(recipient_balance.value() + amount);
    allowance.set_value(allowance.value()-amount);

    call.log("erc20.transfer_from.success");
}
