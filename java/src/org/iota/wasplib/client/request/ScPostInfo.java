// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.request;

import org.iota.wasplib.client.hashtypes.ScAddress;
import org.iota.wasplib.client.hashtypes.ScColor;
import org.iota.wasplib.client.mutable.ScMutableKeyMap;
import org.iota.wasplib.client.mutable.ScMutableMap;

public class ScPostInfo {
	ScMutableMap post;

	public ScPostInfo(ScMutableMap post) {
		this.post = post;
	}

	public ScPostInfo Chain(ScAddress chain) {
		post.GetAddress("chain").SetValue(chain);
		return this;
	}

	public ScPostInfo Contract(String contract) {
		post.GetString("contract").SetValue(contract);
		return this;
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
