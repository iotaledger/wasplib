package org.iota.wasplib.client.immutable;

import org.iota.wasplib.client.Host;
import org.iota.wasplib.client.hashtypes.ScAddress;

public class ScImmutableAddress {
	int objId;
	int keyId;

	public ScImmutableAddress(int objId, int keyId) {
		this.objId = objId;
		this.keyId = keyId;
	}

	public boolean Exists() {
		return Host.Exists(objId, keyId);
	}

	public ScAddress Value() {
		return new ScAddress(Host.GetBytes(objId, keyId));
	}
}
