// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.wasmlib.mutable;

import org.iota.wasp.wasmlib.hashtypes.*;
import org.iota.wasp.wasmlib.host.*;

public class ScMutableContractId {
	int objId;
	int keyId;

	public ScMutableContractId(int objId, int keyId) {
		this.objId = objId;
		this.keyId = keyId;
	}

	public boolean Exists() {
		return Host.Exists(objId, keyId, ScType.TYPE_CONTRACT_ID);
	}

	public void SetValue(ScColor value) {
		Host.SetBytes(objId, keyId, ScType.TYPE_CONTRACT_ID, value.toBytes());
	}

	@Override
	public String toString() {
		return Value().toString();
	}

	public ScContractId Value() {
		return new ScContractId(Host.GetBytes(objId, keyId, ScType.TYPE_CONTRACT_ID));
	}
}
