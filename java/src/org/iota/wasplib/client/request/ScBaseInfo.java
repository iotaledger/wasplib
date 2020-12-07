// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.request;

import org.iota.wasplib.client.hashtypes.ScColor;
import org.iota.wasplib.client.immutable.ScImmutableMap;
import org.iota.wasplib.client.keys.Key;
import org.iota.wasplib.client.mutable.ScMutableMap;

public class ScBaseInfo {
	ScMutableMap request;

	protected ScBaseInfo(ScMutableMap request) {
		this.request = request;
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
