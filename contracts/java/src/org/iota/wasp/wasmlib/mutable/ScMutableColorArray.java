// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.wasmlib.mutable;

import org.iota.wasp.wasmlib.host.Host;
import org.iota.wasp.wasmlib.immutable.ScImmutableColorArray;

public class ScMutableColorArray {
	int objId;

	public ScMutableColorArray(int objId) {
		this.objId = objId;
	}

	public void Clear() {
		Host.SetClear(objId);
	}

	public ScMutableColor GetColor(int index) {
		return new ScMutableColor(objId, index);
	}

	public ScImmutableColorArray Immutable() {
		return new ScImmutableColorArray(objId);
	}

	public int Length() {
		return Host.GetLength(objId);
	}
}
