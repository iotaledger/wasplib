// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.immutable;

import org.iota.wasplib.client.host.Host;

public class ScImmutableBytes {
	int objId;
	int keyId;

	public ScImmutableBytes(int objId, int keyId) {
		this.objId = objId;
		this.keyId = keyId;
	}

	public boolean Exists() {
		return Host.Exists(objId, keyId);
	}

	public byte[] Value() {
		return Host.GetBytes(objId, keyId);
	}
}
