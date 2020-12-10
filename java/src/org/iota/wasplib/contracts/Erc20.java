// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// implementation of ERC-20 smart contract for ISCP
// following https://ethereum.org/en/developers/tutorials/understand-the-erc-20-token-smart-contract/

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
	// state variable
	private static final Key STATE_VAR_SUPPLY = new Key("s");

	// supply private static final Keyant
	private static final Key STATE_VAR_BALANCES = new Key("b"); // name of the map of balances

	// params and return variables, used in calls
	private static final Key PARAM_SUPPLY = new Key("s");
	private static final Key PARAM_CREATOR = new Key("c");
	private static final Key PARAM_ACCOUNT = new Key("ac");
	private static final Key PARAM_DELEGATION = new Key("d");
	private static final Key PARAM_AMOUNT = new Key("am");
	private static final Key PARAM_RECIPIENT = new Key("r");

	public static void onLoad() {
		ScExports exports = new ScExports();
		exports.AddCall("init", Erc20::onInit);
		exports.AddView("totalSupply", Erc20::total_supply);
		exports.AddView("balanceOf", Erc20::balance_of);
		exports.AddView("allowance", Erc20::allowance);
		exports.AddCall("transfer", Erc20::transfer);
		exports.AddCall("approve", Erc20::approve);
		exports.AddCall("transferFrom", Erc20::transfer_from);
	}

	// TODO would be awesome to have some less syntactically cumbersome way to check and validate parameters.

	// init is a private static final Keyructor entry point. It initializes the smart contract with the
	// initial value of the token supply and the owner of that supply
	// - input:
	//   -- PARAM_SUPPLY must be nonzero positive integer
	//   -- PARAM_CREATOR is the AgentID where initial supply is placed
	public static void onInit(ScCallContext ctx) {
		ctx.Log("erc20.init.begin");
		// validate parameters
		// supply
		ScImmutableInt supply = ctx.Params().GetInt(PARAM_SUPPLY);
		String err;
		if (!supply.Exists() || supply.Value() <= 0) {
			err = "er20.init.fail: wrong 'supply' parameter";
			ctx.Log(err);
			ctx.Error().SetValue(err);
			return;
		}
		// creator (owner);
		// we cannot use 'caller' here because the init is always called from the 'root'
		// so, owner of the initial supply must be provided as a parameter PARAM_CREATOR to private static final Keyructor (init);
		ScImmutableAgent creator = ctx.Params().GetAgent(PARAM_CREATOR);
		if (!creator.Exists()) {
			err = "er20.init.fail: wrong 'creator' parameter";
			ctx.Log(err);
			ctx.Error().SetValue(err);
			return;
		}
		ctx.State().GetInt(STATE_VAR_SUPPLY).SetValue(supply.Value());

		// assign the whole supply to creator
		ctx.State().GetMap(STATE_VAR_BALANCES).GetInt(creator.Value()).SetValue(supply.Value());

		String t = "erc20.init.success. Supply: " + supply + ", creator:" + creator;
		ctx.Log(t);
	}

	// the view returns total supply set when creating the contract (a private static final Keyant).
	// Output:
	// - PARAM_SUPPLY: i64
	public static void total_supply(ScViewContext ctx) {
		long supply = ctx.State().GetInt(STATE_VAR_SUPPLY).Value();
		ctx.Results().GetInt(PARAM_SUPPLY).SetValue(supply);
	}

	// the view returns balance of the token held in the account
	// Input:
	// - PARAM_ACCOUNT: agentID
	public static void balance_of(ScViewContext ctx) {
		ScImmutableAgent account = ctx.Params().GetAgent(PARAM_ACCOUNT);
		if (!account.Exists()) {
			String m = "wrong or non existing parameter: " + account;
			ctx.Log(m);
			ctx.Error().SetValue(m);
			return;
		}
		ScImmutableMap balances = ctx.State().GetMap(STATE_VAR_BALANCES);
		long balance = balances.GetInt(account.Value()).Value(); // 0 if doesn't exist
		ctx.Results().GetInt(PARAM_AMOUNT).SetValue(balance);
	}

	// the view returns max number of tokens the owner PARAM_ACCOUNT of the account
	// allowed to retrieve to another party PARAM_DELEGATION
	// Input:
	// - PARAM_ACCOUNT: agentID
	// - PARAM_DELEGATION: agentID
	// Output:
	// - PARAM_AMOUNT: i64. 0 if delegation doesn't exists
	public static void allowance(ScViewContext ctx) {
		ctx.Log("erc20.allowance");
		// validate parameters
		// account
		ScImmutableAgent owner = ctx.Params().GetAgent(PARAM_ACCOUNT);
		String m;
		if (!owner.Exists()) {
			m = "er20.allowance.fail: wrong 'account' parameter";
			ctx.Log(m);
			ctx.Error().SetValue(m);
			return;
		}
		// delegation
		ScImmutableAgent delegation = ctx.Params().GetAgent(PARAM_DELEGATION);
		if (!delegation.Exists()) {
			m = "er20.allowance.fail: wrong 'delegation' parameter";
			ctx.Log(m);
			ctx.Error().SetValue(m);
			return;
		}
		// all allowances of the address 'owner' are stored in the map of the same name
		ScImmutableMap allowances = ctx.State().GetMap(owner.Value());
		long allow = allowances.GetInt(delegation.Value()).Value();
		ctx.Results().GetInt(PARAM_AMOUNT).SetValue(allow);
	}

	// transfer moves tokens from caller's account to target account
	// Input:
	// - PARAM_ACCOUNT: agentID
	// - PARAM_AMOUNT: i64
	public static void transfer(ScCallContext ctx) {
		ctx.Log("erc20.transfer");

		// validate params
		ScImmutableMap params = ctx.Params();
		// account
		ScImmutableAgent target_addrParam = params.GetAgent(PARAM_ACCOUNT);
		String m;
		if (!target_addrParam.Exists()) {
			m = "er20.transfer.fail: wrong 'account' parameter";
			ctx.Log(m);
			ctx.Error().SetValue(m);
			return;
		}
		ScAgent target_addr = target_addrParam.Value();
		// amount
		long amount = params.GetInt(PARAM_AMOUNT).Value();
		if (amount <= 0) {
			m = "erc20.transfer.fail: wrong 'amount' parameter";
			ctx.Log(m);
			ctx.Error().SetValue(m);
			return;
		}
		ScMutableMap balances = ctx.State().GetMap(STATE_VAR_BALANCES);
		ScMutableInt source_balance = balances.GetInt(ctx.Caller());

		if (source_balance.Value() < amount) {
			m = "erc20.transfer.fail: not enough funds";
			ctx.Log(m);
			ctx.Error().SetValue(m);
			return;
		}
		ScMutableInt target_balance = balances.GetInt(target_addr);
		long result = target_balance.Value() + amount;
		if (result <= 0) {
			m = "erc20.transfer.fail: overflow";
			ctx.Log(m);
			ctx.Error().SetValue(m);
			return;
		}
		source_balance.SetValue(source_balance.Value() - amount);
		target_balance.SetValue(target_balance.Value() + amount);
		ctx.Log("erc20.transfer.success");
	}

	// Sets the allowance value for delegated account
	// inputs:
	//  - PARAM_DELEGATION: agentID
	//  - PARAM_AMOUNT: i64
	public static void approve(ScCallContext ctx) {
		ctx.Log("erc20.approve");

		// validate parameters
		ScImmutableAgent delegationParam = ctx.Params().GetAgent(PARAM_DELEGATION);
		String m;
		if (!delegationParam.Exists()) {
			m = "erc20.approve.fail: wrong 'delegation' parameter";
			ctx.Log(m);
			ctx.Error().SetValue(m);
			return;
		}
		ScAgent delegation = delegationParam.Value();
		long amount = ctx.Params().GetInt(PARAM_AMOUNT).Value();
		if (amount <= 0) {
			m = "erc20.approve.fail: wrong 'amount' parameter";
			ctx.Log(m);
			ctx.Error().SetValue(m);
			return;
		}
		// all allowances are in the map under the name of he owner
		ScMutableMap allowances = ctx.State().GetMap(ctx.Caller());
		allowances.GetInt(delegation).SetValue(amount);
		ctx.Log("erc20.approve.success");
	}

	// Moves the amount of tokens from sender to recipient using the allowance mechanism.
	// Amount is then deducted from the caller’s allowance. This function emits the Transfer event.
	// Input:
	// - PARAM_ACCOUNT: agentID   the spender
	// - PARAM_RECIPIENT: agentID   the target
	// - PARAM_AMOUNT: i64
	public static void transfer_from(ScCallContext ctx) {
		ctx.Log("erc20.transfer_from");

		// validate parameters
		ScImmutableAgent accountParam = ctx.Params().GetAgent(PARAM_ACCOUNT);
		String m;
		if (!accountParam.Exists()) {
			m = "erc20.transfer_from.fail: wrong 'account' parameter";
			ctx.Log(m);
			ctx.Error().SetValue(m);
			return;
		}
		ScAgent account = accountParam.Value();
		ScImmutableAgent recipientParam = ctx.Params().GetAgent(PARAM_RECIPIENT);
		if (!recipientParam.Exists()) {
			m = "erc20.transfer_from.fail: wrong 'recipient' parameter";
			ctx.Log(m);
			ctx.Error().SetValue(m);
			return;
		}
		ScAgent recipient = recipientParam.Value();
		ScImmutableInt amountParam = ctx.Params().GetInt(PARAM_AMOUNT);
		if (!amountParam.Exists()) {
			m = "erc20.transfer_from.fail: wrong 'amount' parameter";
			ctx.Log(m);
			ctx.Error().SetValue(m);
			return;
		}
		long amount = amountParam.Value();

		// allowances are in the map under the name of the account
		ScMutableMap allowances = ctx.State().GetMap(account);
		ScMutableInt allowance = allowances.GetInt(recipient);
		if (allowance.Value() < amount) {
			m = "erc20.transfer_from.fail: not enough allowance";
			ctx.Log(m);
			ctx.Error().SetValue(m);
			return;
		}
		ScMutableMap balances = ctx.State().GetMap(STATE_VAR_BALANCES);
		ScMutableInt source_balance = balances.GetInt(account);
		if (source_balance.Value() < amount) {
			m = "erc20.transfer_from.fail: not enough funds";
			ctx.Log(m);
			ctx.Error().SetValue(m);
			return;
		}
		ScMutableInt recipient_balance = balances.GetInt(recipient);
		long result = recipient_balance.Value() + amount;
		if (result <= 0) {
			m = "erc20.transfer_from.fail: overflow";
			ctx.Log(m);
			ctx.Error().SetValue(m);
			return;
		}
		source_balance.SetValue(source_balance.Value() - amount);
		recipient_balance.SetValue(recipient_balance.Value() + amount);
		allowance.SetValue(allowance.Value() - amount);

		ctx.Log("erc20.transfer_from.success");
	}
}
