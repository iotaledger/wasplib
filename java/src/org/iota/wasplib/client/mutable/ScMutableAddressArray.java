// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.mutable;

import org.iota.wasplib.client.Host;
import org.iota.wasplib.client.Keys;
import org.iota.wasplib.client.immutable.ScImmutableAddressArray;

public class ScMutableAddressArray {
	int objId;

	public ScMutableAddressArray(int objId) {
		this.objId = objId;
	}

	public void Clear() {
		Host.SetInt(objId, Keys.KeyLength(), 0);
	}

	public ScMutableAddress GetAddress(int index) {
		return new ScMutableAddress(objId, index);
	}

	public ScImmutableAddressArray Immutable() {
		return new ScImmutableAddressArray(objId);
	}

	public int Length() {
		return (int) Host.GetInt(objId, Keys.KeyLength());
	}
}
