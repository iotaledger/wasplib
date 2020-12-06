// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.request;

import org.iota.wasplib.client.Key;
import org.iota.wasplib.client.hashtypes.ScAddress;
import org.iota.wasplib.client.hashtypes.ScColor;
import org.iota.wasplib.client.mutable.ScMutableMap;

public class ScPostInfo {
	ScMutableMap post;

	public ScPostInfo(ScMutableMap post) {
		this.post = post;
	}

	public ScPostInfo Chain(ScAddress chain) {
		post.GetAddress(new Key("chain")).SetValue(chain);
		return this;
	}

	public ScPostInfo Contract(String contract) {
		post.GetString(new Key("contract")).SetValue(contract);
		return this;
	}

	public ScMutableMap Params() {
		return post.GetMap(new Key("params"));
	}

	public void Post(long delay) {
		post.GetInt(new Key("delay")).SetValue(delay);
	}

	public void Transfer(ScColor color, long amount) {
		ScMutableMap transfers = post.GetMap(new Key("transfers"));
		transfers.GetInt(color).SetValue(amount);
	}
}
