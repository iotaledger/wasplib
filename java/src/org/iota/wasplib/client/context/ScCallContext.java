// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.context;

import org.iota.wasplib.client.hashtypes.ScAgent;
import org.iota.wasplib.client.hashtypes.ScColor;
import org.iota.wasplib.client.host.Host;
import org.iota.wasplib.client.keys.Key;
import org.iota.wasplib.client.keys.KeyId;
import org.iota.wasplib.client.keys.Keys;
import org.iota.wasplib.client.mutable.ScMutableMap;
import org.iota.wasplib.client.mutable.ScMutableMapArray;
import org.iota.wasplib.client.request.ScCallInfo;
import org.iota.wasplib.client.request.ScPostInfo;

public class ScCallContext extends ScBaseContext {
	private static final ScMutableMap root = new ScMutableMap(1);

	public ScCallContext() {
	}

	public ScCallInfo Call(String function) {
		return new ScCallInfo(makeRequest(new Key("calls"), function));
	}

	public ScBalances Incoming() {
		return new ScBalances(root.GetMap(new Key("incoming")).Immutable());
	}

	public void Log(String text) {
		Host.SetString(1, Keys.KeyLog(), text);
	}

	public ScPostInfo Post(String function) {
		return new ScPostInfo(makeRequest(new Key("posts"), function));
	}

	public ScMutableMap State() {
		return root.GetMap(new Key("state"));
	}

	public ScLog TimestampedLog(KeyId key) {
		return new ScLog(root.GetMap(new Key("logs")).GetMapArray(key));
	}

	public void Transfer(ScAgent agent, ScColor color, long amount) {
		ScMutableMapArray transfers = root.GetMapArray(new Key("transfers"));
		ScMutableMap transfer = transfers.GetMap(transfers.Length());
		transfer.GetAgent(new Key("agent")).SetValue(agent);
		transfer.GetColor(new Key("color")).SetValue(color);
		transfer.GetInt(new Key("amount")).SetValue(amount);
	}
}
