package org.iota.wasplib.client.immutable;

import org.iota.wasplib.client.Host;
import org.iota.wasplib.client.hashtypes.ScTxHash;

public class ScImmutableTxHash {
	int objId;
	int keyId;

	public ScImmutableTxHash(int objId, int keyId) {
		this.objId = objId;
		this.keyId = keyId;
	}

	public boolean Exists() {
		return Host.Exists(objId, keyId);
	}

	public ScTxHash Value() {
		return new ScTxHash(Host.GetBytes(objId, keyId));
	}
}
