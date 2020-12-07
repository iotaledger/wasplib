// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.request;

import org.iota.wasplib.client.hashtypes.ScColor;
import org.iota.wasplib.client.immutable.ScImmutableMap;
import org.iota.wasplib.client.keys.Key;
import org.iota.wasplib.client.mutable.ScMutableMap;
import org.iota.wasplib.client.mutable.ScMutableMapArray;

public class ScBaseInfo {
	protected static final ScMutableMap root = new ScMutableMap(1);

	ScMutableMap request;

	protected ScBaseInfo(String key, String function) {
		ScMutableMapArray requests = root.GetMapArray(new Key(key));
		request = requests.GetMap(requests.Length());
		request.GetString(new Key("function")).SetValue(function);
	}

	public ScBaseInfo Contract(String contract) {
		request.GetString(new Key("contract")).SetValue(contract);
		return this;
	}

	protected void exec(long delay) {
		request.GetInt(new Key("delay")).SetValue(delay);
	}

	public ScMutableMap Params() {
		return request.GetMap(new Key("params"));
	}

	protected ScImmutableMap results() {
		return request.GetMap(new Key("results")).Immutable();
	}

	protected void transfer(ScColor color, long amount) {
		ScMutableMap transfers = request.GetMap(new Key("transfers"));
		transfers.GetInt(color).SetValue(amount);
	}
}
