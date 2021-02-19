// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.wasmlib.mutable;

import org.iota.wasp.wasmlib.hashtypes.*;
import org.iota.wasp.wasmlib.host.*;

public class ScMutableAddress {
	int objId;
	int keyId;

	public ScMutableAddress(int objId, int keyId) {
		this.objId = objId;
		this.keyId = keyId;
	}

	public boolean Exists() {
		return Host.Exists(objId, keyId, ScType.TYPE_ADDRESS);
	}

	public void SetValue(ScAddress value) {
		Host.SetBytes(objId, keyId, ScType.TYPE_ADDRESS, value.toBytes());
	}

	@Override
	public String toString() {
		return Value().toString();
	}

	public ScAddress Value() {
		return new ScAddress(Host.GetBytes(objId, keyId, ScType.TYPE_ADDRESS));
	}
}
