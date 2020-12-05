// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.request;

import org.iota.wasplib.client.hashtypes.ScColor;
import org.iota.wasplib.client.immutable.ScImmutableMap;
import org.iota.wasplib.client.mutable.ScMutableKeyMap;
import org.iota.wasplib.client.mutable.ScMutableMap;

public class ScCallInfo {
	ScMutableMap call;

	public ScCallInfo(ScMutableMap call) {
		this.call = call;
	}

	public void Call() {
		call.GetInt("delay").SetValue(-1);
	}

	public ScCallInfo Contract(String contract) {
		call.GetString("contract").SetValue(contract);
		return this;
	}

	public ScMutableMap Params() {
		return call.GetMap("params");
	}

	public ScImmutableMap Results() {
		return call.GetMap("results").Immutable();
	}

	public void Transfer(ScColor color, long amount) {
		ScMutableKeyMap transfers = call.GetKeyMap("transfers");
		transfers.GetInt(color.toBytes()).SetValue(amount);
	}
}
