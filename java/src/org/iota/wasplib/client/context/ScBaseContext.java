// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.context;

import org.iota.wasplib.client.hashtypes.ScAgent;
import org.iota.wasplib.client.host.Host;
import org.iota.wasplib.client.immutable.ScImmutableMap;
import org.iota.wasplib.client.keys.Key;
import org.iota.wasplib.client.mutable.ScMutableMap;
import org.iota.wasplib.client.mutable.ScMutableString;
import org.iota.wasplib.client.request.ScViewInfo;

public class ScBaseContext {
	protected static final ScMutableMap root = new ScMutableMap(1);

	protected ScBaseContext() {
	}

	public ScBalances Balances() {
		return new ScBalances(root.GetMap(Key.Balances).Immutable());
	}

	public ScAgent Caller() {
		return root.GetAgent(Key.Caller).Value();
	}

	public ScContract Contract() {
		return new ScContract(root.GetMap(Key.Contract).Immutable());
	}

	public ScMutableString Error() {
		return root.GetString(Key.Error);
	}

	public Boolean From(ScAgent originator) {
		return Caller().equals(originator);
	}

	public void Log(String text) {
		Host.SetString(1, Key.KEY_LOG, text);
	}

	public ScImmutableMap Params() {
		return root.GetMap(Key.Params).Immutable();
	}

	public ScMutableMap Results() {
		return root.GetMap(Key.Results);
	}

	public long Timestamp() {
		return root.GetInt(Key.Timestamp).Value();
	}

	public void Trace(String text) {
		Host.SetString(1, Key.KEY_TRACE, text);
	}

	public ScUtility Utility() {
		return new ScUtility(root.GetMap(Key.Utility));
	}

	public ScViewInfo View(String function) {
		return new ScViewInfo(function);
	}
}
