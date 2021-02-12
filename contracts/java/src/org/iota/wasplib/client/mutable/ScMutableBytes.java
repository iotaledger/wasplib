// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.mutable;

import org.iota.wasplib.client.context.ScUtility;
import org.iota.wasplib.client.host.Host;

public class ScMutableBytes {
	int objId;
	int keyId;

	public ScMutableBytes(int objId, int keyId) {
		this.objId = objId;
		this.keyId = keyId;
	}

	public boolean Exists() {
		return Host.Exists(objId, keyId);
	}

	public void SetValue(byte[] value) {
		Host.SetBytes(objId, keyId, value);
	}

	@Override
	public String toString() {
		return ScUtility.Base58String(Value());
	}

	public byte[] Value() {
		return Host.GetBytes(objId, keyId);
	}
}
