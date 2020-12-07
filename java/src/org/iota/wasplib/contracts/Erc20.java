// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.contracts;

import org.iota.wasplib.client.context.ScCallContext;
import org.iota.wasplib.client.context.ScViewContext;
import org.iota.wasplib.client.exports.ScExports;
import org.iota.wasplib.client.hashtypes.ScAgent;
import org.iota.wasplib.client.immutable.ScImmutableAgent;
import org.iota.wasplib.client.immutable.ScImmutableInt;
import org.iota.wasplib.client.immutable.ScImmutableMap;
import org.iota.wasplib.client.keys.Key;
import org.iota.wasplib.client.mutable.ScMutableInt;
import org.iota.wasplib.client.mutable.ScMutableMap;

public class Erc20 {
	private static final Key keyAccount = new Key("ac");
	private static final Key keyAmount = new Key("am");
	private static final Key keyBalances = new Key("b");
	private static final Key keyCreator = new Key("c");
	private static final Key keyDelegation = new Key("d");
	private static final Key keyRecipient = new Key("r");
	private static final Key keySupply = new Key("s");

	//export onLoad
	public static void onLoad() {
		ScExports exports = new ScExports();
		exports.AddCall("init", Erc20::onInit);
		exports.AddView("totalSupply", Erc20::totalSupply);
		exports.AddView("balanceOf", Erc20::balanceOf);
		exports.AddView("allowance", Erc20::allowance);
		exports.AddCall("transfer", Erc20::transfer);
		exports.AddCall("approve", Erc20::approve);
		exports.AddCall("transferFrom", Erc20::transferFrom);
	}

	// TODO would be awesome to have some less syntactically cumbersome way to check and validate parameters.

	// init is a constructor entry point. It initializes the smart contract with the initial value of the token supply
	// - input:
	//   -- keySupply must be nonzero positive integer
	//   -- keyCreator is the AgentID where initial supply is placed
	public static void onInit(ScCallContext ctx) {
		ctx.Log("erc20.init. begin");

		// validate parameters
		// supply
		ScImmutableInt supply = ctx.Params().GetInt(keySupply);
		if (!supply.Exists() || supply.Value() <= 0) {
			ctx.Log("er20.init.fail: wrong 'supply' parameter");
			return;
		}
		// creator (owner);
		// we cannot use 'caller' here because the init is always called from the 'root'
		// so, owner of the initial supply must be provided as a parameter keyCreator to constructor (init);
		ScImmutableAgent creator = ctx.Params().GetAgent(keyCreator);
		if (!creator.Exists()) {
			ctx.Log("er20.init.fail: wrong 'creator' parameter");
			return;
		}
		ctx.State().GetInt(keySupply).SetValue(supply.Value());

		// assign the whole supply to creator
		ctx.State().GetMap(keyBalances).GetInt(creator.Value()).SetValue(supply.Value());

		ctx.Log("init.success. Supply = " + ctx.Utility().String(supply.Value()));
		ctx.Log("init.success. Owner = " + creator.Value());
	}

	// the view returns total supply set when creating the contract.
	// Output:
	// - keySupply: i64
	public static void totalSupply(ScViewContext ctx) {
		long supply = ctx.State().GetInt(keySupply).Value();
		ctx.Results().GetInt(keySupply).SetValue(supply);
	}

	// the view returns balance of the token held in the account
	// Input:
	// - keyAccount: agentID
	public static void balanceOf(ScViewContext ctx) {
		ScImmutableAgent account = ctx.Params().GetAgent(keyAccount);
		if (!account.Exists()) {
			ctx.Log("wrong or non existing parameter: " + account.Value());
			return;
		}
		ScImmutableMap balances = ctx.State().GetMap(keyBalances);
		long balance = balances.GetInt(account.Value()).Value(); // 0 if doesn't exist
		ctx.Results().GetInt(keyAmount).SetValue(balance);
	}

	// the view returns max number of tokens the owner PARAM_ACCOUNT of the account
	// allowed to retrieve to another party PARAM_DELEGATION
	// Input:
	// - keyAccount: agentID
	// - keyDelegation: agentID
	// Output:
	// - keyAmount: i64. 0 if delegation doesn't exists
	public static void allowance(ScViewContext ctx) {
		ctx.Log("erc20.allowance");
		// validate parameters
		// account
		ScImmutableAgent owner = ctx.Params().GetAgent(keyAccount);
		if (!owner.Exists()) {
			ctx.Log("er20.allowance.fail: wrong 'account' parameter");
			return;
		}
		// delegation
		ScImmutableAgent delegation = ctx.Params().GetAgent(keyDelegation);
		if (!delegation.Exists()) {
			ctx.Log("er20.allowance.fail: wrong 'delegation' parameter");
			return;
		}
		// all allowances of the address 'owner' are stored in the map of the same name
		ScImmutableMap allowances = ctx.State().GetMap(owner.Value());
		long allow = allowances.GetInt(delegation.Value()).Value();
		ctx.Results().GetInt(keyAmount).SetValue(allow);
	}

	// transfer moves tokens from caller's account to target account
	// Input:
	// - keyAccount: agentID
	// - keyAmount: i64
	public static void transfer(ScCallContext ctx) {
		ctx.Log("erc20.transfer");

		// validate params
		ScImmutableMap params = ctx.Params();
		// account
		ScImmutableAgent account = params.GetAgent(keyAccount);
		if (!account.Exists()) {
			ctx.Log("er20.transfer.fail: wrong 'account' parameter");
			return;
		}
		ScAgent target_addr = account.Value();
		// amount
		long amount = params.GetInt(keyAmount).Value();
		if (amount <= 0) {
			ctx.Log("erc20.transfer.fail: wrong 'amount' parameter");
			return;
		}
		ScMutableMap balances = ctx.State().GetMap(keyBalances);
		ScMutableInt source_balance = balances.GetInt(ctx.Caller());

		if (source_balance.Value() < amount) {
			ctx.Log("erc20.transfer.fail: not enough funds");
			return;
		}
		ScMutableInt target_balance = balances.GetInt(target_addr);
		long result = target_balance.Value() + amount;
		if (result <= 0) {
			ctx.Log("erc20.transfer.fail: overflow");
			return;
		}
		source_balance.SetValue(source_balance.Value() - amount);
		target_balance.SetValue(target_balance.Value() + amount);
		ctx.Log("erc20.transfer.success");
	}

	// Sets the allowance value for delegated account
	// inputs:
	//  - keyDelegation: agentID
	//  - keyAmount: i64
	public static void approve(ScCallContext ctx) {
		ctx.Log("erc20.approve");

		// validate parameters
		ScImmutableAgent accountPar = ctx.Params().GetAgent(keyDelegation);
		if (!accountPar.Exists()) {
			ctx.Log("erc20.approve.fail: wrong 'delegation' parameter");
			return;
		}
		ScAgent account = accountPar.Value();
		long amount = ctx.Params().GetInt(keyAmount).Value();
		if (amount <= 0) {
			ctx.Log("erc20.approve.fail: wrong 'amount' parameter");
			return;
		}
		ScMutableMap allowances = ctx.State().GetMap(ctx.Caller());
		allowances.GetInt(account).SetValue(amount);
		ctx.Log("erc20.approve.success");
	}

	// Moves the amount of tokens from sender to recipient using the allowance mechanism.
	// Amount is then deducted from the caller’s allowance. This function emits the Transfer event.
	// Input:
	// - keyAccount: agentID   the spender
	// - keyRecipient: agentID   the target
	// - keyAmount: i64
	public static void transferFrom(ScCallContext ctx) {
		ctx.Log("erc20.transfer_from");

		// validate parameters
		ScImmutableAgent accountPar = ctx.Params().GetAgent(keyAccount);
		if (!accountPar.Exists()) {
			ctx.Log("erc20.approve.fail: wrong 'account' parameter");
			return;
		}
		ScAgent account = accountPar.Value();
		ScImmutableAgent recipientPar = ctx.Params().GetAgent(keyRecipient);
		if (!recipientPar.Exists()) {
			ctx.Log("erc20.approve.fail: wrong 'recipient' parameter");
			return;
		}
		ScAgent recipient = recipientPar.Value();
		ScImmutableInt amountPar = ctx.Params().GetInt(keyAmount);
		if (!amountPar.Exists()) {
			ctx.Log("erc20.approve.fail: wrong 'amount' parameter");
			return;
		}
		long amount = amountPar.Value();

		// allowances are in the map under the name of the account
		ScMutableMap allowances = ctx.State().GetMap(account);
		ScMutableInt allowance = allowances.GetInt(recipient);
		if (allowance.Value() < amount) {
			ctx.Log("erc20.approve.fail: not enough allowance");
		}

		ScMutableMap balances = ctx.State().GetMap(keyBalances);
		ScMutableInt source_balance = balances.GetInt(account);
		if (source_balance.Value() < amount) {
			ctx.Log("erc20.transfer.fail: not enough funds");
			return;
		}
		ScMutableInt recipient_balance = balances.GetInt(recipient);
		long result = recipient_balance.Value() + amount;
		if (result <= 0) {
			ctx.Log("erc20.transfer.fail: overflow");
			return;
		}
		source_balance.SetValue(source_balance.Value() - amount);
		recipient_balance.SetValue(recipient_balance.Value() + amount);
		allowance.SetValue(allowance.Value() - amount);

		ctx.Log("erc20.transfer_from.success");
	}
}
