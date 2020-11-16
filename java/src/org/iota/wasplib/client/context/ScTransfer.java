// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.context;

import org.iota.wasplib.client.hashtypes.ScAgent;
import org.iota.wasplib.client.hashtypes.ScColor;
import org.iota.wasplib.client.mutable.ScMutableMap;

public class ScTransfer {
	ScMutableMap transfer;

	ScTransfer(ScMutableMap transfer) {
		this.transfer = transfer;
	}

	public void Agent(ScAgent agent) {
		transfer.GetAgent("agent").SetValue(agent);
	}

	public void Amount(long amount) {
		transfer.GetInt("amount").SetValue(amount);
	}

	public void Color(ScColor color) {
		transfer.GetColor("color").SetValue(color);
	}
}
