// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.context;

import org.iota.wasplib.client.immutable.ScImmutableMap;
import org.iota.wasplib.client.mutable.ScMutableMap;

public class ScViewInfo {
	ScMutableMap view;

	ScViewInfo(ScMutableMap view) {
		this.view = view;
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
