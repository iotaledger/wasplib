// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.wasmlib.mutable;

import org.iota.wasp.wasmlib.host.*;
import org.iota.wasp.wasmlib.immutable.*;

public class ScMutableHnameArray {
	int objId;

	public ScMutableHnameArray(int objId) {
		this.objId = objId;
	}

	public void Clear() {
		Host.Clear(objId);
	}

	public ScMutableHname GetHname(int index) {
		return new ScMutableHname(objId, index);
	}

	public ScImmutableHnameArray Immutable() {
		return new ScImmutableHnameArray(objId);
	}

	public int Length() {
		return Host.GetLength(objId);
	}
}
