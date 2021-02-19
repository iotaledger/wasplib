// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.wasmlib.context;

import org.iota.wasp.wasmlib.builders.ScCallBuilder;
import org.iota.wasp.wasmlib.builders.ScDeployBuilder;
import org.iota.wasp.wasmlib.builders.ScPostBuilder;
import org.iota.wasp.wasmlib.builders.ScTransferBuilder;
import org.iota.wasp.wasmlib.hashtypes.ScAddress;
import org.iota.wasp.wasmlib.hashtypes.ScAgentId;
import org.iota.wasp.wasmlib.hashtypes.ScColor;
import org.iota.wasp.wasmlib.host.Host;
import org.iota.wasp.wasmlib.keys.Key;
import org.iota.wasp.wasmlib.mutable.ScMutableMap;

public class ScFuncContext extends ScBaseContext {
	public ScFuncContext() {
	}

	public ScCallBuilder Call(String function) {
		return new ScCallBuilder(function);
	}

	public ScDeployBuilder Deploy(String name, String description) {
		return new ScDeployBuilder(name, description);
	}

	public ScBalances Incoming() {
		return new ScBalances(Host.root.GetMap(Key.Incoming).Immutable());
	}

	public ScPostBuilder Post(String function) {
		return new ScPostBuilder(function);
	}

	public void Event(String text) {
		Host.root.GetString(Key.Event).SetValue(text);
	}

	public ScMutableMap State() {
		return Host.root.GetMap(Key.State);
	}

	public void Transfer(ScAgentId agent, ScColor color, long amount) {
		new ScTransferBuilder(agent).Transfer(color, amount).Send();
	}

	public ScTransferBuilder TransferToAddress(ScAddress address) {
		return new ScTransferBuilder(address);
	}

	public ScTransferBuilder TransferCrossChain(ScAddress chain, ScAgentId agent) {
		return new ScTransferBuilder(chain, agent);
	}
}
