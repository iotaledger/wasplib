// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.request;

import org.iota.wasplib.client.immutable.ScImmutableMap;
import org.iota.wasplib.client.mutable.ScMutableMap;

public class ScViewInfo extends ScBaseInfo {
	public ScViewInfo(ScMutableMap request) {
		super(request);
	}

	public ScImmutableMap Results() {
		return results();
	}

	public void View() {
		exec(-2);
	}
}
