// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.request;

import org.iota.wasplib.client.hashtypes.ScColor;
import org.iota.wasplib.client.immutable.ScImmutableMap;

public class ScCallInfo extends ScBaseInfo {
	public ScCallInfo(String function) {
		super("calls", function);
	}

	public void Call() {
		exec(-1);
	}

	public ScImmutableMap Results() {
		return results();
	}

	public void Transfer(ScColor color, long amount) {
		transfer(color, amount);
	}
}
