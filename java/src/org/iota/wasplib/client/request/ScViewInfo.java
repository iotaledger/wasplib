// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.request;

import org.iota.wasplib.client.immutable.ScImmutableMap;

public class ScViewInfo extends ScBaseInfo {
	public ScViewInfo(String function) {
		super("views", function);
	}

	public ScImmutableMap Results() {
		return results();
	}

	public void View() {
		exec(-2);
	}
}
