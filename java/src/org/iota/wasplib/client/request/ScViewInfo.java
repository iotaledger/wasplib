// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.request;

import org.iota.wasplib.client.immutable.ScImmutableMap;
import org.iota.wasplib.client.mutable.ScMutableMap;

public class ScViewInfo {
	ScMutableMap view;

	public ScViewInfo(ScMutableMap view) {
		this.view = view;
	}

	public ScViewInfo Contract(String contract) {
		view.GetString("contract").SetValue(contract);
		return this;
	}

	public ScMutableMap Params() {
		return view.GetMap("params");
	}

	public ScImmutableMap Results() {
		return view.GetMap("results").Immutable();
	}

	public void View() {
		view.GetInt("delay").SetValue(-2);
	}
}
