// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.immutable;

import org.iota.wasplib.client.host.Host;
import org.iota.wasplib.client.keys.Key;

public class ScImmutableColorArray {
	int objId;

	public ScImmutableColorArray(int objId) {
		this.objId = objId;
	}

	public ScImmutableColor GetColor(int index) {
		return new ScImmutableColor(objId, index);
	}

	public int Length() {
		return (int) Host.GetInt(objId, Key.KEY_LENGTH);
	}
}
