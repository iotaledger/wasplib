// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.mutable;

import org.iota.wasplib.client.hashtypes.ScHash;
import org.iota.wasplib.client.host.Host;

public class ScMutableHash {
	int objId;
	int keyId;

	public ScMutableHash(int objId, int keyId) {
		this.objId = objId;
		this.keyId = keyId;
	}

	public boolean Exists() {
		return Host.Exists(objId, keyId);
	}

	public void SetValue(ScHash value) {
		Host.SetBytes(objId, keyId, value.toBytes());
	}

	@Override
	public String toString() {
		return Value().toString();
	}

	public ScHash Value() {
		return new ScHash(Host.GetBytes(objId, keyId));
	}
}
