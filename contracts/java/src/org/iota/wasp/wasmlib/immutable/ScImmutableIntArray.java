// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.wasmlib.immutable;

import org.iota.wasp.wasmlib.host.*;

public class ScImmutableIntArray {
	int objId;

	public ScImmutableIntArray(int objId) {
		this.objId = objId;
	}

	public ScImmutableInt GetInt(int index) {
		return new ScImmutableInt(objId, index);
	}

	public int Length() {
		return Host.GetLength(objId);
	}
}
