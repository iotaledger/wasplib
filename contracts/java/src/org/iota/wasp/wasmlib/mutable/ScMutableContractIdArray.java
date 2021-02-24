// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.wasmlib.mutable;

import org.iota.wasp.wasmlib.host.*;
import org.iota.wasp.wasmlib.immutable.*;

public class ScMutableContractIdArray {
	int objId;

	public ScMutableContractIdArray(int objId) {
		this.objId = objId;
	}

	public void Clear() {
		Host.Clear(objId);
	}

	public ScMutableContractId GetContractId(int index) {
		return new ScMutableContractId(objId, index);
	}

	public ScImmutableContractIdArray Immutable() {
		return new ScImmutableContractIdArray(objId);
	}

	public int Length() {
		return Host.GetLength(objId);
	}
}
