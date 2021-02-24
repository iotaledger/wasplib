// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.wasmlib.immutable;

import org.iota.wasp.wasmlib.host.*;

public class ScImmutableChainIdArray {
	int objId;

	public ScImmutableChainIdArray(int objId) {
		this.objId = objId;
	}

	public ScImmutableChainId GetChainId(int index) {
		return new ScImmutableChainId(objId, index);
	}

	public int Length() {
		return Host.GetLength(objId);
	}
}
