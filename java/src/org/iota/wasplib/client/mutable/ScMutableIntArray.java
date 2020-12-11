// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.mutable;

import org.iota.wasplib.client.host.Host;
import org.iota.wasplib.client.immutable.ScImmutableIntArray;

public class ScMutableIntArray {
	int objId;

	public ScMutableIntArray(int objId) {
		this.objId = objId;
	}

	public void Clear() {
		Host.SetClear(objId);
	}

	public ScMutableInt GetInt(int index) {
		return new ScMutableInt(objId, index);
	}

	public ScImmutableIntArray Immutable() {
		return new ScImmutableIntArray(objId);
	}

	public int Length() {
		return Host.GetLength(objId);
	}
}
