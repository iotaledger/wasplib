// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.context;

import org.iota.wasplib.client.hashtypes.ScAgent;
import org.iota.wasplib.client.host.Host;
import org.iota.wasplib.client.immutable.ScImmutableMap;
import org.iota.wasplib.client.keys.Key;
import org.iota.wasplib.client.keys.Keys;
import org.iota.wasplib.client.mutable.ScMutableMap;
import org.iota.wasplib.client.mutable.ScMutableString;
import org.iota.wasplib.client.request.ScViewInfo;

public class ScBaseContext {
	protected static final ScMutableMap root = new ScMutableMap(1);

	protected ScBaseContext() {
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

	public long Timestamp() {
		return root.GetInt(new Key("timestamp")).Value();
	}

	public void Trace(String text) {
		Host.SetString(1, Keys.KeyTrace(), text);
	}

	public ScUtility Utility() {
		return new ScUtility(root.GetMap(new Key("utility")));
	}

	public ScViewInfo View(String function) {
		return new ScViewInfo(function);
	}
}
