// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.wasmlib.mutable;

import org.iota.wasp.wasmlib.hashtypes.ScChainId;
import org.iota.wasp.wasmlib.hashtypes.ScColor;
import org.iota.wasp.wasmlib.host.Host;

public class ScMutableChainId {
	int objId;
	int keyId;

	public ScMutableChainId(int objId, int keyId) {
		this.objId = objId;
		this.keyId = keyId;
	}

	public boolean Exists() {
		return Host.Exists(objId, keyId);
	}

	public void SetValue(ScColor value) {
		Host.SetBytes(objId, keyId, value.toBytes());
	}

	@Override
	public String toString() {
		return Value().toString();
	}

	public ScChainId Value() {
		return new ScChainId(Host.GetBytes(objId, keyId));
	}
}
