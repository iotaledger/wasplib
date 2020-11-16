// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.mutable;

import org.iota.wasplib.client.Host;
import org.iota.wasplib.client.hashtypes.ScColor;
import org.iota.wasplib.client.hashtypes.ScTxHash;

public class ScMutableTxHash {
	int objId;
	int keyId;

	public ScMutableTxHash(int objId, int keyId) {
		this.objId = objId;
		this.keyId = keyId;
	}

	public boolean Exists() {
		return Host.Exists(objId, keyId);
	}

	public void SetValue(ScColor value) {
		Host.SetBytes(objId, keyId, value.toBytes());
	}

	public ScTxHash Value() {
		return new ScTxHash(Host.GetBytes(objId, keyId));
	}
}
