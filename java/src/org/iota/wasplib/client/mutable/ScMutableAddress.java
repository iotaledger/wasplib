// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.mutable;

import org.iota.wasplib.client.Host;
import org.iota.wasplib.client.hashtypes.ScAddress;

public class ScMutableAddress {
	int objId;
	int keyId;

	public ScMutableAddress(int objId, int keyId) {
		this.objId = objId;
		this.keyId = keyId;
	}

	public boolean Exists() {
		return Host.Exists(objId, keyId);
	}

	public void SetValue(ScAddress value) {
		Host.SetBytes(objId, keyId, value.toBytes());
	}

	public ScAddress Value() {
		return new ScAddress(Host.GetBytes(objId, keyId));
	}
}
