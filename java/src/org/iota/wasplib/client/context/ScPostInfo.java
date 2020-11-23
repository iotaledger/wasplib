// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.context;

import org.iota.wasplib.client.hashtypes.ScColor;
import org.iota.wasplib.client.mutable.ScMutableMap;
import org.iota.wasplib.client.mutable.ScMutableMapArray;

public class ScPostInfo {
	ScMutableMap post;

	ScPostInfo(ScMutableMap post) {
		this.post = post;
	}

	public ScMutableMap Params() {
		return post.GetMap("params");
	}

	public void Post(long delay) {
		post.GetInt("delay").SetValue(delay);
	}

	public void Transfer(ScColor color, long amount) {
		ScMutableMapArray transfers = post.GetMapArray("transfers");
		ScMutableMap transfer = transfers.GetMap(transfers.Length());
		transfer.GetColor("color").SetValue(color);
		transfer.GetInt("amount").SetValue(amount);
	}
}
