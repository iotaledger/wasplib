// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.wasmlib.immutable;

import org.iota.wasp.wasmlib.hashtypes.ScHash;
import org.iota.wasp.wasmlib.host.Host;

public class ScImmutableHash {
	int objId;
	int keyId;

	public ScImmutableHash(int objId, int keyId) {
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

	public ScHash Value() {
		return new ScHash(Host.GetBytes(objId, keyId));
	}
}
