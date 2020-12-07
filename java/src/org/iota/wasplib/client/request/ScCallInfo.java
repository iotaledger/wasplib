// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.request;

import org.iota.wasplib.client.hashtypes.ScColor;
import org.iota.wasplib.client.immutable.ScImmutableMap;
import org.iota.wasplib.client.mutable.ScMutableMap;

public class ScCallInfo extends ScBaseInfo {
	public ScCallInfo(ScMutableMap request) {
		super(request);
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
