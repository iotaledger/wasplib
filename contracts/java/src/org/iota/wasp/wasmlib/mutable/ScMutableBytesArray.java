// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.wasmlib.mutable;

import org.iota.wasp.wasmlib.host.Host;
import org.iota.wasp.wasmlib.immutable.ScImmutableBytesArray;

public class ScMutableBytesArray {
	int objId;

	public ScMutableBytesArray(int objId) {
		this.objId = objId;
	}

	public void Clear() {
		Host.SetClear(objId);
	}

	public ScMutableBytes GetBytes(int index) {
		return new ScMutableBytes(objId, index);
	}

	public ScImmutableBytesArray Immutable() {
		return new ScImmutableBytesArray(objId);
	}

	public int Length() {
		return Host.GetLength(objId);
	}
}
