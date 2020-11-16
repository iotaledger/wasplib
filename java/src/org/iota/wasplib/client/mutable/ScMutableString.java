// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.mutable;

import org.iota.wasplib.client.Host;

public class ScMutableString {
	int objId;
	int keyId;

	public ScMutableString(int objId, int keyId) {
		this.objId = objId;
		this.keyId = keyId;
	}

	public void SetValue(String value) {
		Host.SetString(objId, keyId, value);
	}

	public String Value() {
		return Host.GetString(objId, keyId);
	}
}
