// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.immutable;

import org.iota.wasplib.client.host.Host;
import org.iota.wasplib.client.keys.Key;

public class ScImmutableAddressArray {
	int objId;

	public ScImmutableAddressArray(int objId) {
		this.objId = objId;
	}

	public ScImmutableAddress GetAddress(int index) {
		return new ScImmutableAddress(objId, index);
	}

	public int Length() {
		return (int) Host.GetInt(objId, Key.KEY_LENGTH);
	}
}
