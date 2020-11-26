// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.context;

import org.iota.wasplib.client.hashtypes.ScColor;
import org.iota.wasplib.client.mutable.ScMutableKeyMap;
import org.iota.wasplib.client.mutable.ScMutableMap;

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
		ScMutableKeyMap transfers = post.GetKeyMap("transfers");
		transfers.GetInt(color.toBytes()).SetValue(amount);
	}
}
