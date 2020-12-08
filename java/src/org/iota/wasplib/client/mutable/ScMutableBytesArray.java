// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.mutable;

import org.iota.wasplib.client.host.Host;
import org.iota.wasplib.client.immutable.ScImmutableBytesArray;
import org.iota.wasplib.client.keys.Key;

public class ScMutableBytesArray {
	int objId;

	public ScMutableBytesArray(int objId) {
		this.objId = objId;
	}

	public void Clear() {
		Host.SetInt(objId, Key.KEY_LENGTH, 0);
	}

	public ScMutableBytes GetBytes(int index) {
		return new ScMutableBytes(objId, index);
	}

	public ScImmutableBytesArray Immutable() {
		return new ScImmutableBytesArray(objId);
	}

	public int Length() {
		return (int) Host.GetInt(objId, Key.KEY_LENGTH);
	}
}
