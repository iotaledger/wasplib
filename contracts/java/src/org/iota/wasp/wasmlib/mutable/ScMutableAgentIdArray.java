// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.wasmlib.mutable;

import org.iota.wasp.wasmlib.host.*;
import org.iota.wasp.wasmlib.immutable.*;

public class ScMutableAgentIdArray {
	int objId;

	public ScMutableAgentIdArray(int objId) {
		this.objId = objId;
	}

	public void Clear() {
		Host.SetClear(objId);
	}

	public ScMutableAgentId GetAgentId(int index) {
		return new ScMutableAgentId(objId, index);
	}

	public ScImmutableAgentIdArray Immutable() {
		return new ScImmutableAgentIdArray(objId);
	}

	public int Length() {
		return Host.GetLength(objId);
	}
}
