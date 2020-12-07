// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.immutable;

import org.iota.wasplib.client.hashtypes.ScAgent;
import org.iota.wasplib.client.host.Host;

public class ScImmutableAgent {
	int objId;
	int keyId;

	public ScImmutableAgent(int objId, int keyId) {
		this.objId = objId;
		this.keyId = keyId;
	}

	public boolean Exists() {
		return Host.Exists(objId, keyId);
	}

	public ScAgent Value() {
		return new ScAgent(Host.GetBytes(objId, keyId));
	}
}
