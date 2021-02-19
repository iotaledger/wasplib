// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.wasmlib.immutable;

import org.iota.wasp.wasmlib.host.*;

public class ScImmutableColorArray {
	int objId;

	public ScImmutableColorArray(int objId) {
		this.objId = objId;
	}

	public ScImmutableColor GetColor(int index) {
		return new ScImmutableColor(objId, index);
	}

	public int Length() {
		return Host.GetLength(objId);
	}
}
