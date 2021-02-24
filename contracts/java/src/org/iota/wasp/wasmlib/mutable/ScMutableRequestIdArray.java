// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.wasmlib.mutable;

import org.iota.wasp.wasmlib.host.*;
import org.iota.wasp.wasmlib.immutable.*;

public class ScMutableRequestIdArray {
	int objId;

	public ScMutableRequestIdArray(int objId) {
		this.objId = objId;
	}

	public void Clear() {
		Host.Clear(objId);
	}

	public ScMutableRequestId GetRequestId(int index) {
		return new ScMutableRequestId(objId, index);
	}

	public ScImmutableRequestIdArray Immutable() {
		return new ScImmutableRequestIdArray(objId);
	}

	public int Length() {
		return Host.GetLength(objId);
	}
}
