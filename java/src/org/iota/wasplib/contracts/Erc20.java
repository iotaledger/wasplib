// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.contracts;

import org.iota.wasplib.client.context.ScCallContext;
import org.iota.wasplib.client.context.ScRequest;
import org.iota.wasplib.client.exports.ScExports;
import org.iota.wasplib.client.hashtypes.ScAgent;
import org.iota.wasplib.client.immutable.ScImmutableAddress;
import org.iota.wasplib.client.immutable.ScImmutableInt;
import org.iota.wasplib.client.immutable.ScImmutableMap;
import org.iota.wasplib.client.mutable.ScMutableInt;
import org.iota.wasplib.client.mutable.ScMutableKeyMap;
import org.iota.wasplib.client.mutable.ScMutableMap;

public class Erc20 {
	private static final String varSupply = "s";
	private static final String varBalances = "b";
	private static final String varTargetAddress = "addr";
	private static final String varAmount = "amount";

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
		ScImmutableMap params = sc.Request().Params();
		ScImmutableInt supplyParam = params.GetInt(varSupply);
		if (supplyParam.Value() == 0) {
			sc.Log("initSC.fail: wrong 'supply' parameter");
			return;
		}
		long supply = supplyParam.Value();
		supplyState.SetValue(supply);
		state.GetKeyMap(varBalances).GetInt(sc.Contract().Owner().toBytes()).SetValue(supply);

		sc.Log("initSC: success");
	}

	public static void transfer(ScCallContext sc) {
		sc.Log("transfer");

		ScMutableMap state = sc.State();
		ScRequest request = sc.Request();
		ScMutableKeyMap balances = state.GetKeyMap(varBalances);

		ScAgent sender = request.Sender();
		sc.Log("sender address: " + sender);

		ScMutableInt sourceBalance = balances.GetInt(sender.toBytes());
		sc.Log("source balance: " + sourceBalance.Value());

		ScImmutableMap params = request.Params();
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

		ScMutableInt targetBalance = balances.GetInt(targetAddr.Value().toBytes());
		targetBalance.SetValue(targetBalance.Value() + amount.Value());
		sourceBalance.SetValue(sourceBalance.Value() - amount.Value());

		sc.Log("transfer: success");
	}

	public static void approve(ScCallContext sc) {
		// TODO
		sc.Log("approve");
	}
}
