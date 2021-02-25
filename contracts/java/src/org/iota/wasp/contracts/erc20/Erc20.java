// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// implementation of ERC-20 smart contract for ISCP
// following https://ethereum.org/en/developers/tutorials/understand-the-erc-20-token-smart-contract/

package org.iota.wasp.contracts.erc20;

import org.iota.wasp.contracts.erc20.lib.*;
import org.iota.wasp.wasmlib.context.*;
import org.iota.wasp.wasmlib.hashtypes.*;
import org.iota.wasp.wasmlib.immutable.*;
import org.iota.wasp.wasmlib.mutable.*;

public class Erc20 {
    // implementation of ERC-20 smart contract for ISCP
    // following https://ethereum.org/en/developers/tutorials/understand-the-erc-20-token-smart-contract/

    // Sets the allowance value for delegated account
    // inputs:
    //  - PARAM_DELEGATION: agentID
    //  - PARAM_AMOUNT: i64
    public static void funcApprove(ScFuncContext ctx, FuncApproveParams params) {
        ctx.Trace("erc20.approve");

        ScAgentId delegation = params.Delegation.Value();
        long amount = params.Amount.Value();
        ctx.Require(amount > 0, "erc20.approve.fail: wrong 'amount' parameter");

        // all allowances are in the map under the name of he owner
        ScMutableMap allowances = ctx.State().GetMap(ctx.Caller());
        allowances.GetInt64(delegation).SetValue(amount);
        ctx.Log("erc20.approve.success");
    }

    // on_init is a constructor entry point. It initializes the smart contract with the
    // initial value of the token supply and the owner of that supply
    // - input:
    //   -- PARAM_SUPPLY must be nonzero positive integer. Mandatory
    //   -- PARAM_CREATOR is the AgentID where initial supply is placed. Mandatory
    public static void funcInit(ScFuncContext ctx, FuncInitParams params) {
        ctx.Trace("erc20.on_init.begin");

        long supply = params.Supply.Value();
        ctx.Require(supply > 0, "erc20.on_init.fail: wrong 'supply' parameter");
        ctx.State().GetInt64(Consts.VarSupply).SetValue(supply);

        // we cannot use 'caller' here because on_init is always called from the 'root'
        // so, owner of the initial supply must be provided as a parameter PARAM_CREATOR to constructor (on_init)
        // assign the whole supply to creator
        ScAgentId creator = params.Creator.Value();
        ctx.State().GetMap(Consts.VarBalances).GetInt64(creator).SetValue(supply);

        String t = "erc20.on_init.success. Supply: " + supply +
                ", creator:" + creator;
        ctx.Log(t);
    }

    // transfer moves tokens from caller's account to target account
    // Input:
    // - PARAM_ACCOUNT: agentID
    // - PARAM_AMOUNT: i64
    public static void funcTransfer(ScFuncContext ctx, FuncTransferParams params) {
        ctx.Trace("erc20.transfer");

        long amount = params.Amount.Value();
        ctx.Require(amount > 0, "erc20.transfer.fail: wrong 'amount' parameter");

        ScMutableMap balances = ctx.State().GetMap(Consts.VarBalances);
        ScMutableInt64 sourceBalance = balances.GetInt64(ctx.Caller());
        ctx.Require(sourceBalance.Value() >= amount, "erc20.transfer.fail: not enough funds");

        ScAgentId targetAddr = params.Account.Value();
        ScMutableInt64 targetBalance = balances.GetInt64(targetAddr);
        long result = targetBalance.Value() + amount;
        ctx.Require(result > 0, "erc20.transfer.fail: overflow");

        sourceBalance.SetValue(sourceBalance.Value() - amount);
        targetBalance.SetValue(targetBalance.Value() + amount);
        ctx.Log("erc20.transfer.success");
    }

    // Moves the amount of tokens from sender to recipient using the allowance mechanism.
    // Amount is then deducted from the caller’s allowance. This function emits the Transfer event.
    // Input:
    // - PARAM_ACCOUNT: agentID   the spender
    // - PARAM_RECIPIENT: agentID   the target
    // - PARAM_AMOUNT: i64
    public static void funcTransferFrom(ScFuncContext ctx, FuncTransferFromParams params) {
        ctx.Trace("erc20.transfer_from");

        // validate parameters
        ScAgentId account = params.Account.Value();
        ScAgentId recipient = params.Recipient.Value();
        long amount = params.Amount.Value();
        ctx.Require(amount > 0, "erc20.transfer_from.fail: wrong 'amount' parameter");

        // allowances are in the map under the name of the account
        ScMutableMap allowances = ctx.State().GetMap(account);
        ScMutableInt64 allowance = allowances.GetInt64(recipient);
        ctx.Require(allowance.Value() >= amount, "erc20.transfer_from.fail: not enough allowance");

        ScMutableMap balances = ctx.State().GetMap(Consts.VarBalances);
        ScMutableInt64 sourceBalance = balances.GetInt64(account);
        ctx.Require(sourceBalance.Value() >= amount, "erc20.transfer_from.fail: not enough funds");

        ScMutableInt64 recipientBalance = balances.GetInt64(recipient);
        long result = recipientBalance.Value() + amount;
        ctx.Require(result > 0, "erc20.transfer_from.fail: overflow");

        sourceBalance.SetValue(sourceBalance.Value() - amount);
        recipientBalance.SetValue(recipientBalance.Value() + amount);
        allowance.SetValue(allowance.Value() - amount);

        ctx.Log("erc20.transfer_from.success");
    }

    // the view returns max number of tokens the owner PARAM_ACCOUNT of the account
    // allowed to retrieve to another party PARAM_DELEGATION
    // Input:
    // - PARAM_ACCOUNT: agentID
    // - PARAM_DELEGATION: agentID
    // Output:
    // - PARAM_AMOUNT: i64
    public static void viewAllowance(ScViewContext ctx, ViewAllowanceParams params) {
        ctx.Trace("erc20.allowance");

        // all allowances of the address 'owner' are stored in the map of the same name
        ScImmutableMap allowances = ctx.State().GetMap(params.Account.Value());
        long allow = allowances.GetInt64(params.Delegation.Value()).Value();
        ctx.Results().GetInt64(Consts.ParamAmount).SetValue(allow);
    }

    // the view returns balance of the token held in the account
    // Input:
    // - PARAM_ACCOUNT: agentID
    public static void viewBalanceOf(ScViewContext ctx, ViewBalanceOfParams params) {
        ScImmutableMap balances = ctx.State().GetMap(Consts.VarBalances);
        long balance = balances.GetInt64(params.Account.Value()).Value();
        ctx.Results().GetInt64(Consts.ParamAmount).SetValue(balance);
    }

    // the view returns total supply set when creating the contract (a constant).
    // Output:
    // - PARAM_SUPPLY: i64
    public static void viewTotalSupply(ScViewContext ctx, ViewTotalSupplyParams params) {
        long supply = ctx.State().GetInt64(Consts.VarSupply).Value();
        ctx.Results().GetInt64(Consts.ParamSupply).SetValue(supply);
    }
}
