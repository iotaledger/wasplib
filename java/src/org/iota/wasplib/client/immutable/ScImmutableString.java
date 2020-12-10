// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.immutable;

import org.iota.wasplib.client.host.Host;

public class ScImmutableString {
	int objId;
	int keyId;

	public ScImmutableString(int objId, int keyId) {
		this.objId = objId;
		this.keyId = keyId;
	}

	@Override
	public String toString() {
		return Value();
	}

	public String Value() {
		return Host.GetString(objId, keyId);
	}
}
