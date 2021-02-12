// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.builders;

import org.iota.wasplib.client.hashtypes.ScAddress;
import org.iota.wasplib.client.hashtypes.ScAgent;
import org.iota.wasplib.client.hashtypes.ScColor;
import org.iota.wasplib.client.host.Host;
import org.iota.wasplib.client.keys.Key;
import org.iota.wasplib.client.mutable.ScMutableMap;
import org.iota.wasplib.client.mutable.ScMutableMapArray;

public class ScTransferBuilder {
	ScMutableMap transfer;

	public ScTransferBuilder(ScAgent agent) {
		this(Host.root.GetAddress(Key.Chain).Value(), agent);
	}

	public ScTransferBuilder(ScAddress address) {
		this(null, address.AsAgent());
	}

	public ScTransferBuilder(ScAddress chain, ScAgent agent) {
		ScMutableMapArray transfers = Host.root.GetMapArray(Key.Transfers);
		transfer = transfers.GetMap(transfers.Length());
		transfer.GetAgent(Key.Agent).SetValue(agent);
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
