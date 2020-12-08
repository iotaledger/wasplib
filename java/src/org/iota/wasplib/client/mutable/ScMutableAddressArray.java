// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.mutable;

import org.iota.wasplib.client.host.Host;
import org.iota.wasplib.client.immutable.ScImmutableAddressArray;
import org.iota.wasplib.client.keys.Key;

public class ScMutableAddressArray {
	int objId;

	public ScMutableAddressArray(int objId) {
		this.objId = objId;
	}

	public void Clear() {
		Host.SetInt(objId, Key.KEY_LENGTH, 0);
	}

	public ScMutableAddress GetAddress(int index) {
		return new ScMutableAddress(objId, index);
	}

	public ScImmutableAddressArray Immutable() {
		return new ScImmutableAddressArray(objId);
	}

	public int Length() {
		return (int) Host.GetInt(objId, Key.KEY_LENGTH);
	}
}
