// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.context;

import org.iota.wasplib.client.Host;
import org.iota.wasplib.client.Key;
import org.iota.wasplib.client.KeyId;
import org.iota.wasplib.client.Keys;
import org.iota.wasplib.client.hashtypes.ScAgent;
import org.iota.wasplib.client.immutable.ScImmutableMap;
import org.iota.wasplib.client.immutable.ScImmutableMapArray;
import org.iota.wasplib.client.mutable.ScMutableMap;
import org.iota.wasplib.client.mutable.ScMutableString;
import org.iota.wasplib.client.request.ScViewInfo;

public class ScViewContext {
	private static final ScMutableMap root = new ScMutableMap(1);

	public ScViewContext() {
	}

	public ScBalances Balances() {
		return new ScBalances(root.GetMap(new Key("balances")).Immutable());
	}

	public ScAgent Caller() {
		return root.GetAgent(new Key("caller")).Value();
	}

	public ScContract Contract() {
		return new ScContract(root.GetMap(new Key("contract")).Immutable());
	}

	public ScMutableString Error() {
		return root.GetString(new Key("error"));
	}

	public Boolean From(ScAgent originator) {
		return Caller().equals(originator);
	}

	public void Log(String text) {
		Host.SetString(1, Keys.KeyLog(), text);
	}

	public ScImmutableMap Params() {
		return root.GetMap(new Key("params")).Immutable();
	}

	public ScMutableMap Results() {
		return root.GetMap(new Key("results"));
	}

	public ScImmutableMap State() {
		return root.GetMap(new Key("state")).Immutable();
	}

	public long Timestamp() {
		return root.GetInt(new Key("timestamp")).Value();
	}

	public ScImmutableMapArray TimestampedLog(KeyId key) {
		return root.GetMap(new Key("logs")).GetMapArray(key).Immutable();
	}

	public void Trace(String text) {
		Host.SetString(1, Keys.KeyTrace(), text);
	}

	public ScUtility Utility() {
		return new ScUtility(root.GetMap(new Key("utility")));
	}

	public ScViewInfo View(String function) {
		return new ScViewInfo(ScCallContext.makeRequest(new Key("views"), function));
	}
}
