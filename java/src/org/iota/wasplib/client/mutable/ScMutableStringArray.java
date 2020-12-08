// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.mutable;

import org.iota.wasplib.client.host.Host;
import org.iota.wasplib.client.immutable.ScImmutableStringArray;
import org.iota.wasplib.client.keys.Key;

public class ScMutableStringArray {
	int objId;

	public ScMutableStringArray(int objId) {
		this.objId = objId;
	}

	public void Clear() {
		Host.SetInt(objId, Key.KEY_LENGTH, 0);
	}

	public ScMutableString GetString(int index) {
		return new ScMutableString(objId, index);
	}

	public ScImmutableStringArray Immutable() {
		return new ScImmutableStringArray(objId);
	}

	public int Length() {
		return (int) Host.GetInt(objId, Key.KEY_LENGTH);
	}
}
