// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.wasmlib.immutable;

import org.iota.wasp.wasmlib.hashtypes.ScAgentId;
import org.iota.wasp.wasmlib.host.Host;

public class ScImmutableAgentId {
	int objId;
	int keyId;

	public ScImmutableAgentId(int objId, int keyId) {
		this.objId = objId;
		this.keyId = keyId;
	}

	public boolean Exists() {
		return Host.Exists(objId, keyId);
	}

	@Override
	public String toString() {
		return Value().toString();
	}

	public ScAgentId Value() {
		return new ScAgentId(Host.GetBytes(objId, keyId));
	}
}
