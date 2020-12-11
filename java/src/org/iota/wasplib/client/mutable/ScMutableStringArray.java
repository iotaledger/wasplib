// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.mutable;

import org.iota.wasplib.client.host.Host;
import org.iota.wasplib.client.immutable.ScImmutableStringArray;

public class ScMutableStringArray {
	int objId;

	public ScMutableStringArray(int objId) {
		this.objId = objId;
	}

	public void Clear() {
		Host.SetClear(objId);
	}

	public ScMutableString GetString(int index) {
		return new ScMutableString(objId, index);
	}

	public ScImmutableStringArray Immutable() {
		return new ScImmutableStringArray(objId);
	}

	public int Length() {
		return Host.GetLength(objId);
	}
}
