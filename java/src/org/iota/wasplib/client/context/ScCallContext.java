// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.context;

import org.iota.wasplib.client.builders.ScCallBuilder;
import org.iota.wasplib.client.builders.ScDeployBuilder;
import org.iota.wasplib.client.builders.ScPostBuilder;
import org.iota.wasplib.client.builders.ScTransferBuilder;
import org.iota.wasplib.client.hashtypes.ScAddress;
import org.iota.wasplib.client.hashtypes.ScAgent;
import org.iota.wasplib.client.hashtypes.ScColor;
import org.iota.wasplib.client.keys.Key;
import org.iota.wasplib.client.keys.MapKey;
import org.iota.wasplib.client.mutable.ScMutableMap;

public class ScCallContext extends ScBaseContext {
	private static final ScMutableMap root = new ScMutableMap(1);

	public ScCallContext() {
	}

	public ScCallBuilder Call(String function) {
		return new ScCallBuilder(function);
	}

	public ScDeployBuilder Deploy(String name, String description) {
		return new ScDeployBuilder(name, description);
	}

	public ScBalances Incoming() {
		return new ScBalances(root.GetMap(Key.Incoming).Immutable());
	}

	public ScPostBuilder Post(String function) {
		return new ScPostBuilder(function);
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
		new ScTransferBuilder(agent).Transfer(color, amount).Send();
	}

	public ScTransferBuilder TransferToAddress(ScAddress address) {
		return new ScTransferBuilder(address);
	}

	public ScTransferBuilder TransferCrossChain(ScAddress chain, ScAgent agent) {
		return new ScTransferBuilder(chain, agent);
	}
}
