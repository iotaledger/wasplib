// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.request;

import org.iota.wasplib.client.Key;
import org.iota.wasplib.client.immutable.ScImmutableMap;
import org.iota.wasplib.client.mutable.ScMutableMap;

public class ScViewInfo {
	ScMutableMap view;

	public ScViewInfo(ScMutableMap view) {
		this.view = view;
	}

	public ScViewInfo Contract(String contract) {
		view.GetString(new Key("contract")).SetValue(contract);
		return this;
	}

	public ScMutableMap Params() {
		return view.GetMap(new Key("params"));
	}

	public ScImmutableMap Results() {
		return view.GetMap(new Key("results")).Immutable();
	}

	public void View() {
		view.GetInt(new Key("delay")).SetValue(-2);
	}
}
