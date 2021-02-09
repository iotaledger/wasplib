// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// implementation of ERC-20 smart contract for ISCP
// following https://ethereum.org/en/developers/tutorials/understand-the-erc-20-token-smart-contract/

use wasmlib::*;

use crate::*;

// Sets the allowance value for delegated account
// inputs:
//  - PARAM_DELEGATION: agentID
//  - PARAM_AMOUNT: i64
pub fn func_approve(ctx: &ScFuncContext, params: &FuncApproveParams) {
    ctx.trace("erc20.approve");

    let delegation = params.delegation.value();
    let amount = params.amount.value();
    ctx.require(amount > 0, "erc20.approve.fail: wrong 'amount' parameter");

    // all allowances are in the map under the name of he owner
    let allowances = ctx.state().get_map(&ctx.caller());
    allowances.get_int(&delegation).set_value(amount);
    ctx.log("erc20.approve.success");
}

// on_init is a constructor entry point. It initializes the smart contract with the
// initial value of the token supply and the owner of that supply
// - input:
//   -- PARAM_SUPPLY must be nonzero positive integer. Mandatory
//   -- PARAM_CREATOR is the AgentID where initial supply is placed. Mandatory
pub fn func_init(ctx: &ScFuncContext, params: &FuncInitParams) {
    ctx.trace("erc20.on_init.begin");

    let supply = params.supply.value();
    ctx.require(supply > 0, "erc20.on_init.fail: wrong 'supply' parameter");
    ctx.state().get_int(VAR_SUPPLY).set_value(supply);

    // we cannot use 'caller' here because on_init is always called from the 'root'
    // so, owner of the initial supply must be provided as a parameter PARAM_CREATOR to constructor (on_init)
    // assign the whole supply to creator
    let creator = params.creator.value();
    ctx.state().get_map(VAR_BALANCES).get_int(&creator).set_value(supply);

    let t = "erc20.on_init.success. Supply: ".to_string() + &supply.to_string() +
        &", creator:".to_string() + &creator.to_string();
    ctx.log(&t);
}

// transfer moves tokens from caller's account to target account
// Input:
// - PARAM_ACCOUNT: agentID
// - PARAM_AMOUNT: i64
pub fn func_transfer(ctx: &ScFuncContext, params: &FuncTransferParams) {
    ctx.trace("erc20.transfer");

    let amount = params.amount.value();
    ctx.require(amount > 0, "erc20.transfer.fail: wrong 'amount' parameter");

    let balances = ctx.state().get_map(VAR_BALANCES);
    let source_balance = balances.get_int(&ctx.caller());
    ctx.require(source_balance.value() >= amount, "erc20.transfer.fail: not enough funds");

    let target_addr = params.account.value();
    let target_balance = balances.get_int(&target_addr);
    let result = target_balance.value() + amount;
    ctx.require(result > 0, "erc20.transfer.fail: overflow");

    source_balance.set_value(source_balance.value() - amount);
    target_balance.set_value(target_balance.value() + amount);
    ctx.log("erc20.transfer.success");
}

// Moves the amount of tokens from sender to recipient using the allowance mechanism.
// Amount is then deducted from the caller’s allowance. This function emits the Transfer event.
// Input:
// - PARAM_ACCOUNT: agentID   the spender
// - PARAM_RECIPIENT: agentID   the target
// - PARAM_AMOUNT: i64
pub fn func_transfer_from(ctx: &ScFuncContext, params: &FuncTransferFromParams) {
    ctx.trace("erc20.transfer_from");

    // validate parameters
    let account = params.account.value();
    let recipient = params.recipient.value();
    let amount = params.amount.value();
    ctx.require(amount > 0, "erc20.transfer_from.fail: wrong 'amount' parameter");

    // allowances are in the map under the name of the account
    let allowances = ctx.state().get_map(&account);
    let allowance = allowances.get_int(&recipient);
    ctx.require(allowance.value() >= amount, "erc20.transfer_from.fail: not enough allowance");

    let balances = ctx.state().get_map(VAR_BALANCES);
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

// the view returns max number of tokens the owner PARAM_ACCOUNT of the account
// allowed to retrieve to another party PARAM_DELEGATION
// Input:
// - PARAM_ACCOUNT: agentID
// - PARAM_DELEGATION: agentID
// Output:
// - PARAM_AMOUNT: i64
pub fn view_allowance(ctx: &ScViewContext, params: &ViewAllowanceParams) {
    ctx.trace("erc20.allowance");

    // all allowances of the address 'owner' are stored in the map of the same name
    let allowances = ctx.state().get_map(&params.account.value());
    let allow = allowances.get_int(&params.delegation.value()).value();
    ctx.results().get_int(PARAM_AMOUNT).set_value(allow);
}

// the view returns balance of the token held in the account
// Input:
// - PARAM_ACCOUNT: agentID
pub fn view_balance_of(ctx: &ScViewContext, params: &ViewBalanceOfParams) {
    let balances = ctx.state().get_map(VAR_BALANCES);
    let balance = balances.get_int(&params.account.value()).value();
    ctx.results().get_int(PARAM_AMOUNT).set_value(balance)
}

// the view returns total supply set when creating the contract (a constant).
// Output:
// - PARAM_SUPPLY: i64
pub fn view_total_supply(ctx: &ScViewContext, _params: &ViewTotalSupplyParams) {
    let supply = ctx.state().get_int(VAR_SUPPLY).value();
    ctx.results().get_int(PARAM_SUPPLY).set_value(supply);
}
