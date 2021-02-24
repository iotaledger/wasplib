// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.wasmlib.immutable;

import org.iota.wasp.wasmlib.host.*;

public class ScImmutableHnameArray {
	int objId;

	public ScImmutableHnameArray(int objId) {
		this.objId = objId;
	}

	public ScImmutableHname GetHname(int index) {
		return new ScImmutableHname(objId, index);
	}

	public int Length() {
		return Host.GetLength(objId);
	}
}
