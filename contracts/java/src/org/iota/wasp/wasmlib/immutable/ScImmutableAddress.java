// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.wasmlib.immutable;

import org.iota.wasp.wasmlib.hashtypes.ScAddress;
import org.iota.wasp.wasmlib.host.Host;

public class ScImmutableAddress {
	int objId;
	int keyId;

	public ScImmutableAddress(int objId, int keyId) {
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

	public ScAddress Value() {
		return new ScAddress(Host.GetBytes(objId, keyId));
	}
}
