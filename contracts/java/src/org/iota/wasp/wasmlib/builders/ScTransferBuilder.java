// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.wasmlib.builders;

import org.iota.wasp.wasmlib.hashtypes.ScAddress;
import org.iota.wasp.wasmlib.hashtypes.ScAgentId;
import org.iota.wasp.wasmlib.hashtypes.ScColor;
import org.iota.wasp.wasmlib.host.Host;
import org.iota.wasp.wasmlib.keys.Key;
import org.iota.wasp.wasmlib.mutable.ScMutableMap;
import org.iota.wasp.wasmlib.mutable.ScMutableMapArray;

public class ScTransferBuilder {
	ScMutableMap transfer;

	public ScTransferBuilder(ScAgentId agent) {
		this(Host.root.GetAddress(Key.Chain).Value(), agent);
	}

	public ScTransferBuilder(ScAddress address) {
		this(null, address.AsAgentId());
	}

	public ScTransferBuilder(ScAddress chain, ScAgentId agent) {
		ScMutableMapArray transfers = Host.root.GetMapArray(Key.Transfers);
		transfer = transfers.GetMap(transfers.Length());
		transfer.GetAgentId(Key.Agent).SetValue(agent);
		if (chain != null) {
			transfer.GetAddress(Key.Chain).SetValue(chain);
		}
	}

	public void Send() {
		transfer.GetInt(ScColor.MINT).SetValue(-1);
	}

	public ScTransferBuilder Transfer(ScColor color, long amount) {
		transfer.GetInt(color).SetValue(amount);
		return this;
	}
}
