// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.context;

import org.iota.wasplib.client.Host;
import org.iota.wasplib.client.Keys;
import org.iota.wasplib.client.hashtypes.ScAgent;
import org.iota.wasplib.client.immutable.ScImmutableMap;
import org.iota.wasplib.client.immutable.ScImmutableMapArray;
import org.iota.wasplib.client.mutable.ScMutableMap;
import org.iota.wasplib.client.mutable.ScMutableString;
import org.iota.wasplib.client.request.ScViewInfo;

public class ScViewContext {
	ScMutableMap root;

	public ScViewContext() {
		root = new ScMutableMap(1);
	}

	public ScBalances Balances() {
		return new ScBalances(root.GetKeyMap("balances").Immutable());
	}

	public ScAgent Caller() {
		return root.GetAgent("caller").Value();
	}

	public ScContract Contract() {
		return new ScContract(root.GetMap("contract").Immutable());
	}

	public ScMutableString Error() {
		return root.GetString("error");
	}

	public Boolean From(ScAgent originator) {
		return Caller().equals(originator);
	}

	public void Log(String text) {
		Host.SetString(1, Keys.KeyLog(), text);
	}

	public ScImmutableMap Params() {
		return root.GetMap("params").Immutable();
	}

	public ScMutableMap Results() {
		return root.GetMap("results");
	}

	public ScImmutableMap State() {
		return root.GetMap("state").Immutable();
	}

	public long Timestamp() {
		return root.GetInt("timestamp").Value();
	}

	public ScImmutableMapArray TimestampedLog(String key) {
		return root.GetMap("logs").GetMapArray(key).Immutable();
	}

	public void Trace(String text) {
		Host.SetString(1, Keys.KeyTrace(), text);
	}

	public ScUtility Utility() {
		return new ScUtility(root.GetMap("utility"));
	}

	public ScViewInfo View(String function) {
		return new ScViewInfo(ScCallContext.makeRequest("views", function));
	}
}
