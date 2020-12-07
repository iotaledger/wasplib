// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.immutable;

import org.iota.wasplib.client.hashtypes.ScColor;
import org.iota.wasplib.client.host.Host;

public class ScImmutableColor {
	int objId;
	int keyId;

	public ScImmutableColor(int objId, int keyId) {
		this.objId = objId;
		this.keyId = keyId;
	}

	public boolean Exists() {
		return Host.Exists(objId, keyId);
	}

	public ScColor Value() {
		return new ScColor(Host.GetBytes(objId, keyId));
	}
}
