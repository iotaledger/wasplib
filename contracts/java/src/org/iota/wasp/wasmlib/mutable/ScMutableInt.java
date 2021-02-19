// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.wasmlib.mutable;

import org.iota.wasp.wasmlib.host.Host;

public class ScMutableInt {
	int objId;
	int keyId;

	public ScMutableInt(int objId, int keyId) {
		this.objId = objId;
		this.keyId = keyId;
	}

	public boolean Exists() {
		return Host.Exists(objId, keyId);
	}

	public void SetValue(long value) {
		Host.SetInt(objId, keyId, value);
	}

	@Override
	public String toString() {
		return "" + Value();
	}

	public long Value() {
		return Host.GetInt(objId, keyId);
	}
}
