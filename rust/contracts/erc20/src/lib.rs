// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// implementation of ERC-20 smart contract for ISCP
// following https://ethereum.org/en/developers/tutorials/understand-the-erc-20-token-smart-contract/

#![allow(dead_code)]
#![allow(non_snake_case)]

use wasplib::client::*;

// state variable
const STATE_VAR_SUPPLY: &str = "s";
// supply constant
const STATE_VAR_BALANCES: &str = "b";     // map of balances

// params and return variables, used in calls
const PARAM_SUPPLY: &str = "s";
const PARAM_CREATOR: &str = "c";
const PARAM_ACCOUNT: &str = "ac";
const PARAM_DELEGATION: &str = "d";
const PARAM_AMOUNT: &str = "am";
const PARAM_RECIPIENT: &str = "r";

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

// TODO would be awesome to have some less syntactically cumbersome way to check and validate parameters.

// init is a constructor entry point. It initializes the smart contract with the initial value of the token supply
// - input:
//   -- PARAM_SUPPLY must be nonzero positive integer
//   -- PARAM_CREATOR is the AgentID where initial supply is placed
fn init(ctx: &ScCallContext) {
    ctx.log("erc20.init. begin");

    // validate parameters
    // supply
    let supply = ctx.params().get_int(PARAM_SUPPLY);
    if !supply.exists() || supply.value() <= 0 {
        ctx.log("er20.init.fail: wrong 'supply' parameter");
        return;
    }
    // creator (owner)
    // we cannot use 'caller' here because the init is always called from the 'root'
    // so, owner of the initial supply must be provided as a parameter PARAM_CREATOR to constructor (init)
    let creator = ctx.params().get_agent(PARAM_CREATOR);
    if !creator.exists() {
        ctx.log("er20.init.fail: wrong 'creator' parameter");
        return;
    }
    ctx.state().get_int(STATE_VAR_SUPPLY).set_value(supply.value());

    // assign the whole supply to creator
    ctx.state().get_key_map(STATE_VAR_BALANCES).get_int(creator.value().to_bytes()).set_value(supply.value());

    ctx.log(&("init.success. Supply = ".to_string() + &supply.value().to_string()));
    ctx.log(&("init.success. Owner = ".to_string() + &creator.value().to_string()));
}

// the view returns total supply set when creating the contract.
// Output:
// - PARAM_SUPPLY: i64
fn total_supply(ctx: &ScViewContext) {
    let supply = ctx.state().get_int(STATE_VAR_SUPPLY).value();
    ctx.results().get_int(PARAM_SUPPLY).set_value(supply);
}

// the view returns balance of the token held in the account
// Input:
// - PARAM_ACCOUNT: agentID
fn balance_of(ctx: &ScViewContext) {
    let account = ctx.params().get_agent(PARAM_ACCOUNT);
    if !account.exists() {
        ctx.log(&("wrong or non existing parameter: ".to_string() + &account.value().to_string()));
        return;
    }
    let balances = ctx.state().get_key_map(STATE_VAR_BALANCES);
    let balance = balances.get_int(account.value().to_bytes()).value();  // 0 if doesn't exist
    ctx.results().get_int(PARAM_AMOUNT).set_value(balance)
}

// the view returns max number of tokens the owner PARAM_ACCOUNT of the account
// allowed to retrieve to another party PARAM_DELEGATION
// Input:
// - PARAM_ACCOUNT: agentID
// - PARAM_DELEGATION: agentID
// Output:
// - PARAM_AMOUNT: i64. 0 if delegation doesn't exists
fn allowance(ctx: &ScViewContext) {
    ctx.log("erc20.allowance");
    // validate parameters
    // account
    let owner = ctx.params().get_agent(PARAM_ACCOUNT);
    if !owner.exists() {
        ctx.log("er20.allowance.fail: wrong 'account' parameter");
        return;
    }
    // delegation
    let delegation = ctx.params().get_agent(PARAM_DELEGATION);
    if !delegation.exists() {
        ctx.log("er20.allowance.fail: wrong 'delegation' parameter");
        return;
    }
    // all allowances of the address 'owner' are stored in the map of the same name
    let allowances = ctx.state().get_key_map(&owner.value().to_string());
    let allow = allowances.get_int(delegation.value().to_bytes()).value();
    ctx.results().get_int(PARAM_AMOUNT).set_value(allow);
}

// transfer moves tokens from caller's account to target account
// Input:
// - PARAM_ACCOUNT: agentID
// - PARAM_AMOUNT: i64
fn transfer(ctx: &ScCallContext) {
    ctx.log("erc20.transfer");

    // validate params
    let params = ctx.params();
    // account
    let target_addr = params.get_agent(PARAM_ACCOUNT);
    if !target_addr.exists() {
        ctx.log("er20.transfer.fail: wrong 'account' parameter");
        return;
    }
    let target_addr = target_addr.value();
    // amount
    let amount = params.get_int(PARAM_AMOUNT).value();
    if amount <= 0 {
        ctx.log("erc20.transfer.fail: wrong 'amount' parameter");
        return;
    }
    let balances = ctx.state().get_key_map(STATE_VAR_BALANCES);
    let source_balance = balances.get_int(ctx.caller().to_bytes());

    if source_balance.value() < amount {
        ctx.log("erc20.transfer.fail: not enough funds");
        return;
    }
    let target_balance = balances.get_int(target_addr.to_bytes());
    let result = target_balance.value() + amount;
    if result <= 0 {
        ctx.log("erc20.transfer.fail: overflow");
        return;
    }
    source_balance.set_value(source_balance.value() - amount);
    target_balance.set_value(target_balance.value() + amount);
    ctx.log("erc20.transfer.success");
}

// Sets the allowance value for delegated account
// inputs:
//  - PARAM_DELEGATION: agentID
//  - PARAM_AMOUNT: i64
fn approve(ctx: &ScCallContext) {
    ctx.log("erc20.approve");

    // validate parameters
    let account = ctx.params().get_agent(PARAM_DELEGATION);
    if !account.exists() {
        ctx.log("erc20.approve.fail: wrong 'delegation' parameter");
        return;
    }
    let account = account.value();
    let amount = ctx.params().get_int(PARAM_AMOUNT).value();
    if amount <= 0 {
        ctx.log("erc20.approve.fail: wrong 'amount' parameter");
        return;
    }
    let caller = ctx.caller().to_string();
    let allowances = ctx.state().get_key_map(&caller);
    allowances.get_int(account.to_bytes()).set_value(amount);
    ctx.log("erc20.approve.success");
}

// Moves the amount of tokens from sender to recipient using the allowance mechanism.
// Amount is then deducted from the caller’s allowance. This function emits the Transfer event.
// Input:
// - PARAM_ACCOUNT: agentID   the spender
// - PARAM_RECIPIENT: agentID   the target
// - PARAM_AMOUNT: i64
fn transfer_from(ctx: &ScCallContext) {
    ctx.log("erc20.transfer_from");

    // validate parameters
    let account = ctx.params().get_agent(PARAM_ACCOUNT);
    if !account.exists() {
        ctx.log("erc20.approve.fail: wrong 'account' parameter");
        return;
    }
    let account = account.value();
    let recipient = ctx.params().get_agent(PARAM_RECIPIENT);
    if !recipient.exists() {
        ctx.log("erc20.approve.fail: wrong 'recipient' parameter");
        return;
    }
    let recipient = recipient.value();
    let amount = ctx.params().get_int(PARAM_AMOUNT);
    if !amount.exists() {
        ctx.log("erc20.approve.fail: wrong 'amount' parameter");
        return;
    }
    let amount = amount.value();

    // allowances are in the map under the name of the account
    let allowances = ctx.state().get_key_map(&account.to_string());
    let allowance = allowances.get_int(recipient.to_bytes());
    if allowance.value() < amount {
        ctx.log("erc20.approve.fail: not enough allowance");
    }

    let balances = ctx.state().get_key_map(STATE_VAR_BALANCES);
    let source_balance = balances.get_int(account.to_bytes());
    if source_balance.value() < amount {
        ctx.log("erc20.transfer.fail: not enough funds");
        return;
    }
    let recipient_balance = balances.get_int(recipient.to_bytes());
    let result = recipient_balance.value() + amount;
    if result <= 0 {
        ctx.log("erc20.transfer.fail: overflow");
        return;
    }
    source_balance.set_value(source_balance.value() - amount);
    recipient_balance.set_value(recipient_balance.value() + amount);
    allowance.set_value(allowance.value() - amount);

    ctx.log("erc20.transfer_from.success");
}
