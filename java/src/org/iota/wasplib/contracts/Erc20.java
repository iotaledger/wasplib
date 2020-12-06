// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.contracts;

import org.iota.wasplib.client.Key;
import org.iota.wasplib.client.context.ScCallContext;
import org.iota.wasplib.client.exports.ScExports;
import org.iota.wasplib.client.hashtypes.ScAgent;
import org.iota.wasplib.client.immutable.ScImmutableAddress;
import org.iota.wasplib.client.immutable.ScImmutableInt;
import org.iota.wasplib.client.immutable.ScImmutableMap;
import org.iota.wasplib.client.mutable.ScMutableInt;
import org.iota.wasplib.client.mutable.ScMutableMap;

public class Erc20 {
	private static final Key varSupply = new Key("s");
	private static final Key varBalances = new Key("b");
	private static final Key varTargetAddress = new Key("addr");
	private static final Key varAmount = new Key("amount");

	//export onLoad
	public static void onLoad() {
		ScExports exports = new ScExports();
		exports.AddCall("initSC", Erc20::initSC);
		exports.AddCall("transfer", Erc20::transfer);
		exports.AddCall("approve", Erc20::approve);
	}

	public static void initSC(ScCallContext sc) {
		sc.Log("initSC");

		ScMutableMap state = sc.State();
		ScMutableInt supplyState = state.GetInt(varSupply);
		if (supplyState.Value() > 0) {
			// already initialized
			sc.Log("initSC.fail: already initialized");
			return;
		}
		ScImmutableMap params = sc.Params();
		ScImmutableInt supplyParam = params.GetInt(varSupply);
		if (supplyParam.Value() == 0) {
			sc.Log("initSC.fail: wrong 'supply' parameter");
			return;
		}
		long supply = supplyParam.Value();
		supplyState.SetValue(supply);
		state.GetMap(varBalances).GetInt(sc.Contract().Owner()).SetValue(supply);

		sc.Log("initSC: success");
	}

	public static void transfer(ScCallContext sc) {
		sc.Log("transfer");

		ScMutableMap state = sc.State();
		ScMutableMap balances = state.GetMap(varBalances);

		ScAgent caller = sc.Caller();
		sc.Log("caller address: " + caller);

		ScMutableInt sourceBalance = balances.GetInt(caller);
		sc.Log("source balance: " + sourceBalance.Value());

		ScImmutableMap params = sc.Params();
		ScImmutableInt amount = params.GetInt(varAmount);
		if (amount.Value() == 0) {
			sc.Log("transfer.fail: wrong 'amount' parameter");
			return;
		}
		if (amount.Value() > sourceBalance.Value()) {
			sc.Log("transfer.fail: not enough balance");
			return;
		}
		ScImmutableAddress targetAddr = params.GetAddress(varTargetAddress);
		// TODO check if it is a correct address, otherwise won't be possible to transfer from it

		ScMutableInt targetBalance = balances.GetInt(targetAddr.Value());
		targetBalance.SetValue(targetBalance.Value() + amount.Value());
		sourceBalance.SetValue(sourceBalance.Value() - amount.Value());

		sc.Log("transfer: success");
	}

	public static void approve(ScCallContext sc) {
		// TODO
		sc.Log("approve");
	}
}
