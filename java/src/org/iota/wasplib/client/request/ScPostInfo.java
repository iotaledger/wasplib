// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.request;

import org.iota.wasplib.client.hashtypes.ScAddress;
import org.iota.wasplib.client.hashtypes.ScColor;
import org.iota.wasplib.client.keys.Key;

public class ScPostInfo extends ScBaseInfo {
	public ScPostInfo(String function) {
		super("posts", function);
	}

	public ScPostInfo Chain(ScAddress chain) {
		request.GetAddress(new Key("chain")).SetValue(chain);
		return this;
	}

	public void Post(long delay) {
		exec(delay);
	}

	public void Transfer(ScColor color, long amount) {
		transfer(color, amount);
	}
}
