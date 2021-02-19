// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// implementation of ERC-20 smart contract for ISCP
// following https://ethereum.org/en/developers/tutorials/understand-the-erc-20-token-smart-contract/

package org.iota.wasp.contracts.erc20;

import org.iota.wasp.contracts.erc20.lib.*;
import org.iota.wasp.wasmlib.context.*;
import org.iota.wasp.wasmlib.hashtypes.*;
import org.iota.wasp.wasmlib.immutable.*;
import org.iota.wasp.wasmlib.keys.*;
import org.iota.wasp.wasmlib.mutable.*;

public class Erc20 {
	// implementation of ERC-20 smart contract for ISCP
	// following https://ethereum.org/en/developers/tutorials/understand-the-erc-20-token-smart-contract/

	// state variable
	private static final Key stateVarSupply = new Key("s");
	// supply constant
	private static final Key stateVarBalances = new Key("b");     // name of the map of balances

	// params and return variables, used in calls
	private static final Key paramSupply = new Key("s");
	private static final Key paramCreator = new Key("c");
	private static final Key paramAccount = new Key("ac");
	private static final Key paramDelegation = new Key("d");
	private static final Key paramAmount = new Key("am");
	private static final Key paramRecipient = new Key("r");

	// Sets the allowance value for delegated account
	// inputs:
	//  - PARAM_DELEGATION: agentID
	//  - PARAM_AMOUNT: i64
	public static void FuncApprove(ScFuncContext ctx, FuncApproveParams params) {
		ctx.Log("erc20.approve");

		// validate parameters
		ScImmutableAgentId delegationParam = ctx.Params().GetAgentId(paramDelegation);
		if (!delegationParam.Exists()) {
			ctx.Panic("erc20.approve.fail: wrong 'delegation' parameter");
		}
		ScAgentId delegation = delegationParam.Value();
		long amount = ctx.Params().GetInt(paramAmount).Value();
		if (amount <= 0) {
			ctx.Panic("erc20.approve.fail: wrong 'amount' parameter");
		}
		// all allowances are in the map under the name of he owner
		ScMutableMap allowances = ctx.State().GetMap(ctx.Caller());
		allowances.GetInt(delegation).SetValue(amount);
		ctx.Log("erc20.approve.success");
	}

	// on_init is a constructor entry point. It initializes the smart contract with the
	// initial value of the token supply and the owner of that supply
	// - input:
	//   -- PARAM_SUPPLY must be nonzero positive integer
	//   -- PARAM_CREATOR is the AgentID where initial supply is placed
	public static void FuncInit(ScFuncContext ctx, FuncInitParams params) {
		ctx.Log("erc20.onInit.begin");
		// validate parameters
		// supply
		ScImmutableInt supply = ctx.Params().GetInt(paramSupply);
		if (!supply.Exists() || supply.Value() <= 0) {
			ctx.Panic("erc20.onInit.fail: wrong 'supply' parameter");
		}
		// creator (owner)
		// we cannot use 'caller' here because on_init is always called from the 'root'
		// so, owner of the initial supply must be provided as a parameter PARAM_CREATOR to constructor (on_init)
		ScImmutableAgentId creator = ctx.Params().GetAgentId(paramCreator);
		if (!creator.Exists()) {
			ctx.Panic("erc20.onInit.fail: wrong 'creator' parameter");
		}
		ctx.State().GetInt(stateVarSupply).SetValue(supply.Value());

		// assign the whole supply to creator
		ctx.State().GetMap(stateVarBalances).GetInt(creator.Value()).SetValue(supply.Value());

		String t = "erc20.onInit.success. Supply: "
				+ supply.Value()
				+ ", creator:"
				+ creator.Value();
		ctx.Log(t);
	}

	// transfer moves tokens from caller's account to target account
	// Input:
	// - PARAM_ACCOUNT: agentID
	// - PARAM_AMOUNT: i64
	public static void FuncTransfer(ScFuncContext ctx, FuncTransferParams params) {
		ctx.Log("erc20.transfer");

		// validate params
		ScImmutableMap p = ctx.Params();
		// account
		ScImmutableAgentId targetAddrParam = p.GetAgentId(paramAccount);
		if (!targetAddrParam.Exists()) {
			ctx.Panic("erc20.transfer.fail: wrong 'account' parameter");
		}
		ScAgentId targetAddr = targetAddrParam.Value();
		// amount
		long amount = p.GetInt(paramAmount).Value();
		if (amount <= 0) {
			ctx.Panic("erc20.transfer.fail: wrong 'amount' parameter");
		}
		ScMutableMap balances = ctx.State().GetMap(stateVarBalances);
		ScMutableInt sourceBalance = balances.GetInt(ctx.Caller());

		if (sourceBalance.Value() < amount) {
			ctx.Panic("erc20.transfer.fail: not enough funds");
		}
		ScMutableInt targetBalance = balances.GetInt(targetAddr);
		long result = targetBalance.Value() + amount;
		if (result <= 0) {
			ctx.Panic("erc20.transfer.fail: overflow");
		}
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
	public static void FuncTransferFrom(ScFuncContext ctx, FuncTransferFromParams params) {
		ctx.Log("erc20.transferFrom");

		// validate parameters
		ScImmutableAgentId accountParam = ctx.Params().GetAgentId(paramAccount);
		if (!accountParam.Exists()) {
			ctx.Panic("erc20.transferFrom.fail: wrong 'account' parameter");
		}
		ScAgentId account = accountParam.Value();
		ScImmutableAgentId recipientParam = ctx.Params().GetAgentId(paramRecipient);
		if (!recipientParam.Exists()) {
			ctx.Panic("erc20.transferFrom.fail: wrong 'recipient' parameter");
		}
		ScAgentId recipient = recipientParam.Value();
		ScImmutableInt amountParam = ctx.Params().GetInt(paramAmount);
		if (!amountParam.Exists()) {
			ctx.Panic("erc20.transferFrom.fail: wrong 'amount' parameter");
		}
		long amount = amountParam.Value();

		// allowances are in the map under the name of the account
		ScMutableMap allowances = ctx.State().GetMap(account);
		ScMutableInt allowance = allowances.GetInt(recipient);
		if (allowance.Value() < amount) {
			ctx.Panic("erc20.transferFrom.fail: not enough allowance");
		}
		ScMutableMap balances = ctx.State().GetMap(stateVarBalances);
		ScMutableInt sourceBalance = balances.GetInt(account);
		if (sourceBalance.Value() < amount) {
			ctx.Panic("erc20.transferFrom.fail: not enough funds");
		}
		ScMutableInt recipientBalance = balances.GetInt(recipient);
		long result = recipientBalance.Value() + amount;
		if (result <= 0) {
			ctx.Panic("erc20.transferFrom.fail: overflow");
		}
		sourceBalance.SetValue(sourceBalance.Value() - amount);
		recipientBalance.SetValue(recipientBalance.Value() + amount);
		allowance.SetValue(allowance.Value() - amount);

		ctx.Log("erc20.transferFrom.success");
	}

	// the view returns max number of tokens the owner PARAM_ACCOUNT of the account
	// allowed to retrieve to another party PARAM_DELEGATION
	// Input:
	// - PARAM_ACCOUNT: agentID
	// - PARAM_DELEGATION: agentID
	// Output:
	// - PARAM_AMOUNT: i64. 0 if delegation doesn't exists
	public static void ViewAllowance(ScViewContext ctx, ViewAllowanceParams params) {
		ctx.Log("erc20.allowance");
		// validate parameters
		// account
		ScImmutableAgentId owner = ctx.Params().GetAgentId(paramAccount);
		if (!owner.Exists()) {
			ctx.Panic("erc20.allowance.fail: wrong 'account' parameter");
		}
		// delegation
		ScImmutableAgentId delegation = ctx.Params().GetAgentId(paramDelegation);
		if (!delegation.Exists()) {
			ctx.Panic("erc20.allowance.fail: wrong 'delegation' parameter");
		}
		// all allowances of the address 'owner' are stored in the map of the same name
		ScImmutableMap allowances = ctx.State().GetMap(owner.Value());
		long allow = allowances.GetInt(delegation.Value()).Value();
		ctx.Results().GetInt(paramAmount).SetValue(allow);
	}

	// the view returns balance of the token held in the account
	// Input:
	// - PARAM_ACCOUNT: agentID
	public static void ViewBalanceOf(ScViewContext ctx, ViewBalanceOfParams params) {
		ScImmutableAgentId account = ctx.Params().GetAgentId(paramAccount);
		if (!account.Exists()) {
			ctx.Panic("wrong or non existing parameter: " + account.Value());
		}
		ScImmutableMap balances = ctx.State().GetMap(stateVarBalances);
		long balance = balances.GetInt(account.Value()).Value();  // 0 if doesn't exist
		ctx.Results().GetInt(paramAmount).SetValue(balance);
	}

	// the view returns total supply set when creating the contract (a constant).
	// Output:
	// - PARAM_SUPPLY: i64
	public static void ViewTotalSupply(ScViewContext ctx, ViewTotalSupplyParams params) {
		long supply = ctx.State().GetInt(stateVarSupply).Value();
		ctx.Results().GetInt(paramSupply).SetValue(supply);
	}
}
