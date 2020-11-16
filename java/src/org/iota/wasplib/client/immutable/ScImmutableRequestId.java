// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.immutable;

import org.iota.wasplib.client.Host;
import org.iota.wasplib.client.hashtypes.ScRequestId;

public class ScImmutableRequestId {
	int objId;
	int keyId;

	public ScImmutableRequestId(int objId, int keyId) {
		this.objId = objId;
		this.keyId = keyId;
	}

	public boolean Exists() {
		return Host.Exists(objId, keyId);
	}

	public ScRequestId Value() {
		return new ScRequestId(Host.GetBytes(objId, keyId));
	}
}
