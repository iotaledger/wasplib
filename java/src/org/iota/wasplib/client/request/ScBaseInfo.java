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
		request.GetString(Key.Function).SetValue(function);
	}

	public ScBaseInfo Contract(String contract) {
		request.GetString(Key.Contract).SetValue(contract);
		return this;
	}

	protected void exec(long delay) {
		request.GetInt(Key.Delay).SetValue(delay);
	}

	public ScMutableMap Params() {
		return request.GetMap(Key.Params);
	}

	protected ScImmutableMap results() {
		return request.GetMap(Key.Results).Immutable();
	}

	protected void transfer(ScColor color, long amount) {
		ScMutableMap transfers = request.GetMap(Key.Transfers);
		transfers.GetInt(color).SetValue(amount);
	}
}
