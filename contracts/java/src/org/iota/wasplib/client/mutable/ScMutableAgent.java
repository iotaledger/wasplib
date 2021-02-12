// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.mutable;

import org.iota.wasplib.client.hashtypes.ScAgent;
import org.iota.wasplib.client.host.Host;

public class ScMutableAgent {
	int objId;
	int keyId;

	public ScMutableAgent(int objId, int keyId) {
		this.objId = objId;
		this.keyId = keyId;
	}

	public boolean Exists() {
		return Host.Exists(objId, keyId);
	}

	public void SetValue(ScAgent value) {
		Host.SetBytes(objId, keyId, value.toBytes());
	}

	@Override
	public String toString() {
		return Value().toString();
	}

	public ScAgent Value() {
		return new ScAgent(Host.GetBytes(objId, keyId));
	}
}
