package org.iota.wasplib.client.immutable;

import org.iota.wasplib.client.Host;

public class ScImmutableInt {
	int objId;
	int keyId;

	public ScImmutableInt(int objId, int keyId) {
		this.objId = objId;
		this.keyId = keyId;
	}

	public long Value() {
		return Host.GetInt(objId, keyId);
	}
}
