// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.wasmlib.mutable;

import org.iota.wasp.wasmlib.hashtypes.ScColor;
import org.iota.wasp.wasmlib.hashtypes.ScContractId;
import org.iota.wasp.wasmlib.host.Host;

public class ScMutableContractId {
	int objId;
	int keyId;

	public ScMutableContractId(int objId, int keyId) {
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

	public ScContractId Value() {
		return new ScContractId(Host.GetBytes(objId, keyId));
	}
}
