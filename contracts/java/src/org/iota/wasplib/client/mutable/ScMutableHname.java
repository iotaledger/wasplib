// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.mutable;

import org.iota.wasplib.client.hashtypes.Hname;
import org.iota.wasplib.client.hashtypes.ScColor;
import org.iota.wasplib.client.host.Host;

public class ScMutableHname {
	int objId;
	int keyId;

	public ScMutableHname(int objId, int keyId) {
		this.objId = objId;
		this.keyId = keyId;
	}

	public boolean Exists() {
		return Host.Exists(objId, keyId);
	}

	public void SetValue(ScColor value) {
		Host.SetBytes(objId, keyId, value.toBytes());
	}

	@Override
	public String toString() {
		return Value().toString();
	}

	public Hname Value() {
		return new Hname(Host.GetBytes(objId, keyId));
	}
}
