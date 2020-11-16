// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.mutable;

import org.iota.wasplib.client.Host;
import org.iota.wasplib.client.Keys;
import org.iota.wasplib.client.immutable.ScImmutableIntArray;

public class ScMutableIntArray {
	int objId;

	public ScMutableIntArray(int objId) {
		this.objId = objId;
	}

	public void Clear() {
		Host.SetInt(objId, Keys.KeyLength(), 0);
	}

	public ScMutableInt GetInt(int index) {
		return new ScMutableInt(objId, index);
	}

	public ScImmutableIntArray Immutable() {
		return new ScImmutableIntArray(objId);
	}

	public int Length() {
		return (int) Host.GetInt(objId, Keys.KeyLength());
	}
}
