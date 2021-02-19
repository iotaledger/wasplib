// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.wasmlib.immutable;

import org.iota.wasp.wasmlib.host.*;

public class ScImmutableAddressArray {
	int objId;

	public ScImmutableAddressArray(int objId) {
		this.objId = objId;
	}

	public ScImmutableAddress GetAddress(int index) {
		return new ScImmutableAddress(objId, index);
	}

	public int Length() {
		return Host.GetLength(objId);
	}
}
