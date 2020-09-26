package org.iota.wasplib.client.immutable;

import org.iota.wasplib.client.Host;

public class ScImmutableBytes {
	int objId;
	int keyId;

	public ScImmutableBytes(int objId, int keyId) {
		this.objId = objId;
		this.keyId = keyId;
	}

	public byte[] Value() {
		return Host.GetBytes(objId, keyId);
	}
}
