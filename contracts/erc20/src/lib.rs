// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// implementation of ERC-20 smart contract for ISCP
// following https://ethereum.org/en/developers/tutorials/understand-the-erc-20-token-smart-contract/

use wasmlib::*;

// state variable
const STATE_VAR_SUPPLY: &str = "s";
// supply constant
const STATE_VAR_BALANCES: &str = "b";     // name of the map of balances

// params and return variables, used in calls
const PARAM_SUPPLY: &str = "s";
const PARAM_CREATOR: &str = "c";
const PARAM_ACCOUNT: &str = "ac";
const PARAM_DELEGATION: &str = "d";
const PARAM_AMOUNT: &str = "am";
const PARAM_RECIPIENT: &str = "r";

#[no_mangle]
fn on_load() {
    let exports = ScExports::new();
    exports.add_call("init", on_init);
    exports.add_view("total_supply", total_supply);
    exports.add_view("balance_of", balance_of);
    exports.add_view("allowance", allowance);
    exports.add_call("transfer", transfer);
    exports.add_call("approve", approve);
    exports.add_call("transfer_from", transfer_from);
}

// TODO would be awesome to have some less syntactically cumbersome way to check and validate parameters.

// on_init is a constructor entry point. It initializes the smart contract with the
// initial value of the token supply and the owner of that supply
// - input:
//   -- PARAM_SUPPLY must be nonzero positive integer. Mandatory
//   -- PARAM_CREATOR is the AgentID where initial supply is placed. Mandatory
fn on_init(ctx: &ScCallContext) {
    ctx.trace("erc20.on_init.begin");
    // validate parameters
    // supply
    let supply = ctx.params().get_int(PARAM_SUPPLY);
    ctx.require(supply.exists() && supply.value() > 0, "erc20.on_init.fail: wrong 'supply' parameter");
    // creator (owner)
    // we cannot use 'caller' here because on_init is always called from the 'root'
    // so, owner of the initial supply must be provided as a parameter PARAM_CREATOR to constructor (on_init)
    let creator = ctx.params().get_agent_id(PARAM_CREATOR);
    ctx.require(creator.exists(), "erc20.on_init.fail: wrong 'creator' parameter");

    ctx.state().get_int(STATE_VAR_SUPPLY).set_value(supply.value());

    // assign the whole supply to creator
    ctx.state().get_map(STATE_VAR_BALANCES).get_int(&creator.value()).set_value(supply.value());

    let t = "erc20.on_init.success. Supply: ".to_string() + &supply.value().to_string() +
        &", creator:".to_string() + &creator.value().to_string();
    ctx.log(&t);
}

// the view returns total supply set when creating the contract (a constant).
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
    let account = ctx.params().get_agent_id(PARAM_ACCOUNT);
    ctx.require(account.exists(), &("wrong or non existing parameter: ".to_string() + &account.value().to_string()));

    let balances = ctx.state().get_map(STATE_VAR_BALANCES);
    let balance = balances.get_int(&account.value()).value();  // 0 if doesn't exist
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
    ctx.trace("erc20.allowance");
    // validate parameters
    // account
    let owner = ctx.params().get_agent_id(PARAM_ACCOUNT);
    ctx.require(owner.exists(), "erc20.allowance.fail: wrong 'account' parameter");
    // delegation
    let delegation = ctx.params().get_agent_id(PARAM_DELEGATION);
    ctx.require(delegation.exists(), "erc20.allowance.fail: wrong 'delegation' parameter");

    // all allowances of the address 'owner' are stored in the map of the same name
    let allowances = ctx.state().get_map(&owner.value());
    let allow = allowances.get_int(&delegation.value()).value();
    ctx.results().get_int(PARAM_AMOUNT).set_value(allow);
}

// transfer moves tokens from caller's account to target account
// Input:
// - PARAM_ACCOUNT: agentID
// - PARAM_AMOUNT: i64
fn transfer(ctx: &ScCallContext) {
    ctx.trace("erc20.transfer");

    // validate params
    let params = ctx.params();
    // account
    let target_addr = params.get_agent_id(PARAM_ACCOUNT);
    ctx.require(target_addr.exists(), "erc20.transfer.fail: wrong 'account' parameter");

    let target_addr = target_addr.value();
    // amount
    let amount = params.get_int(PARAM_AMOUNT).value();
    ctx.require(amount > 0, "erc20.transfer.fail: wrong 'amount' parameter");

    let balances = ctx.state().get_map(STATE_VAR_BALANCES);
    let source_balance = balances.get_int(&ctx.caller());

    ctx.require(source_balance.value() >= amount, "erc20.transfer.fail: not enough funds");

    let target_balance = balances.get_int(&target_addr);
    let result = target_balance.value() + amount;
    ctx.require(result > 0, "erc20.transfer.fail: overflow");

    source_balance.set_value(source_balance.value() - amount);
    target_balance.set_value(target_balance.value() + amount);
    ctx.log("erc20.transfer.success");
}

// Sets the allowance value for delegated account
// inputs:
//  - PARAM_DELEGATION: agentID
//  - PARAM_AMOUNT: i64
fn approve(ctx: &ScCallContext) {
    ctx.trace("erc20.approve");

    // validate parameters
    let delegation = ctx.params().get_agent_id(PARAM_DELEGATION);
    ctx.require(delegation.exists(), "erc20.approve.fail: wrong 'delegation' parameter");

    let delegation = delegation.value();
    let amount = ctx.params().get_int(PARAM_AMOUNT).value();
    ctx.require(amount > 0, "erc20.approve.fail: wrong 'amount' parameter");

    // all allowances are in the map under the name of he owner
    let allowances = ctx.state().get_map(&ctx.caller());
    allowances.get_int(&delegation).set_value(amount);
    ctx.log("erc20.approve.success");
}

// Moves the amount of tokens from sender to recipient using the allowance mechanism.
// Amount is then deducted from the caller’s allowance. This function emits the Transfer event.
// Input:
// - PARAM_ACCOUNT: agentID   the spender
// - PARAM_RECIPIENT: agentID   the target
// - PARAM_AMOUNT: i64
fn transfer_from(ctx: &ScCallContext) {
    ctx.trace("erc20.transfer_from");

    // validate parameters
    let account = ctx.params().get_agent_id(PARAM_ACCOUNT);
    ctx.require(account.exists(), "erc20.transfer_from.fail: wrong 'account' parameter");

    let account = account.value();
    let recipient = ctx.params().get_agent_id(PARAM_RECIPIENT);
    ctx.require(recipient.exists(), "erc20.transfer_from.fail: wrong 'recipient' parameter");

    let recipient = recipient.value();
    let amount = ctx.params().get_int(PARAM_AMOUNT);
    ctx.require(amount.exists(), "erc20.transfer_from.fail: wrong 'amount' parameter");
    let amount = amount.value();

    // allowances are in the map under the name of the account
    let allowances = ctx.state().get_map(&account);
    let allowance = allowances.get_int(&recipient);
    ctx.require(allowance.value() >= amount, "erc20.transfer_from.fail: not enough allowance");

    let balances = ctx.state().get_map(STATE_VAR_BALANCES);
    let source_balance = balances.get_int(&account);
    ctx.require(source_balance.value() >= amount, "erc20.transfer_from.fail: not enough funds");

    let recipient_balance = balances.get_int(&recipient);
    let result = recipient_balance.value() + amount;
    ctx.require(result > 0, "erc20.transfer_from.fail: overflow");

    source_balance.set_value(source_balance.value() - amount);
    recipient_balance.set_value(recipient_balance.value() + amount);
    allowance.set_value(allowance.value() - amount);

    ctx.log("erc20.transfer_from.success");
}
