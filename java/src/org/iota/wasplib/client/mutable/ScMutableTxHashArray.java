// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.mutable;

import org.iota.wasplib.client.Host;
import org.iota.wasplib.client.Keys;
import org.iota.wasplib.client.immutable.ScImmutableTxHashArray;

public class ScMutableTxHashArray {
	int objId;

	public ScMutableTxHashArray(int objId) {
		this.objId = objId;
	}

	public void Clear() {
		Host.SetInt(objId, Keys.KeyLength(), 0);
	}

	public ScMutableTxHash GetTxHash(int index) {
		return new ScMutableTxHash(objId, index);
	}

	public ScImmutableTxHashArray Immutable() {
		return new ScImmutableTxHashArray(objId);
	}

	public int Length() {
		return (int) Host.GetInt(objId, Keys.KeyLength());
	}
}
