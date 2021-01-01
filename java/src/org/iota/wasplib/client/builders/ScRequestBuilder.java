// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.builders;

import org.iota.wasplib.client.hashtypes.ScColor;
import org.iota.wasplib.client.immutable.ScImmutableMap;
import org.iota.wasplib.client.keys.Key;
import org.iota.wasplib.client.mutable.ScMutableMap;
import org.iota.wasplib.client.mutable.ScMutableMapArray;

public abstract class ScRequestBuilder {
	protected static final ScMutableMap root = new ScMutableMap(1);

	ScMutableMap request;

	protected ScRequestBuilder(String key, String function) {
		ScMutableMapArray requests = root.GetMapArray(new Key(key));
		request = requests.GetMap(requests.Length());
		request.GetString(Key.Function).SetValue(function);
	}

	protected void contract(String contract) {
		request.GetString(Key.Contract).SetValue(contract);
	}

	protected void exec(long delay) {
		request.GetInt(Key.Delay).SetValue(delay);
	}

	protected ScMutableMap params() {
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
