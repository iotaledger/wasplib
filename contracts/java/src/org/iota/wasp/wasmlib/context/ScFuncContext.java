// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.wasmlib.context;

import org.iota.wasp.wasmlib.builders.*;
import org.iota.wasp.wasmlib.hashtypes.*;
import org.iota.wasp.wasmlib.host.*;
import org.iota.wasp.wasmlib.keys.*;
import org.iota.wasp.wasmlib.mutable.*;

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
