// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.immutable;

import org.iota.wasplib.client.hashtypes.Hname;
import org.iota.wasplib.client.host.Host;

public class ScImmutableHname {
	int objId;
	int keyId;

	public ScImmutableHname(int objId, int keyId) {
		this.objId = objId;
		this.keyId = keyId;
	}

	public boolean Exists() {
		return Host.Exists(objId, keyId);
	}

	@Override
	public String toString() {
		return Value().toString();
	}

	public Hname Value() {
		return new Hname(Host.GetBytes(objId, keyId));
	}
}
