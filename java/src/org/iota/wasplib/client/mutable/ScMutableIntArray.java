// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.mutable;

import org.iota.wasplib.client.host.Host;
import org.iota.wasplib.client.immutable.ScImmutableIntArray;
import org.iota.wasplib.client.keys.Key;

public class ScMutableIntArray {
	int objId;

	public ScMutableIntArray(int objId) {
		this.objId = objId;
	}

	public void Clear() {
		Host.SetInt(objId, Key.KEY_LENGTH, 0);
	}

	public ScMutableInt GetInt(int index) {
		return new ScMutableInt(objId, index);
	}

	public ScImmutableIntArray Immutable() {
		return new ScImmutableIntArray(objId);
	}

	public int Length() {
		return (int) Host.GetInt(objId, Key.KEY_LENGTH);
	}
}
