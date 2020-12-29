// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.context;

import org.iota.wasplib.client.hashtypes.ScAgent;
import org.iota.wasplib.client.hashtypes.ScColor;
import org.iota.wasplib.client.keys.Key;
import org.iota.wasplib.client.keys.MapKey;
import org.iota.wasplib.client.mutable.ScMutableMap;
import org.iota.wasplib.client.mutable.ScMutableMapArray;
import org.iota.wasplib.client.request.ScCallInfo;
import org.iota.wasplib.client.request.ScPostInfo;

public class ScCallContext extends ScBaseContext {
	private static final ScMutableMap root = new ScMutableMap(1);

	public ScCallContext() {
	}

	public ScCallInfo Call(String function) {
		return new ScCallInfo(function);
	}

	public ScBalances Incoming() {
		return new ScBalances(root.GetMap(Key.Incoming).Immutable());
	}

	public ScPostInfo Post(String function) {
		return new ScPostInfo(function);
	}

	public void SignalEvent(String text) {
		root.GetString(Key.Event).SetValue(text);
	}

	public ScMutableMap State() {
		return root.GetMap(Key.State);
	}

	public ScLog TimestampedLog(MapKey key) {
		return new ScLog(root.GetMap(Key.Logs).GetMapArray(key));
	}

	public void Transfer(ScAgent agent, ScColor color, long amount) {
		ScMutableMapArray transfers = root.GetMapArray(Key.Transfers);
		ScMutableMap transfer = transfers.GetMap(transfers.Length());
		transfer.GetAgent(Key.Agent).SetValue(agent);
		transfer.GetInt(color).SetValue(amount);
	}
}
