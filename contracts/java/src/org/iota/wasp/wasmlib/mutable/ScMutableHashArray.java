// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.wasmlib.mutable;

import org.iota.wasp.wasmlib.host.*;
import org.iota.wasp.wasmlib.immutable.*;

public class ScMutableHashArray {
	int objId;

	public ScMutableHashArray(int objId) {
		this.objId = objId;
	}

	public void Clear() {
		Host.Clear(objId);
	}

	public ScMutableHash GetHash(int index) {
		return new ScMutableHash(objId, index);
	}

	public ScImmutableHashArray Immutable() {
		return new ScImmutableHashArray(objId);
	}

	public int Length() {
		return Host.GetLength(objId);
	}
}
