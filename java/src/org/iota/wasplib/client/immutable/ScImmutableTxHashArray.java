package org.iota.wasplib.client.immutable;

import org.iota.wasplib.client.Host;
import org.iota.wasplib.client.Keys;

public class ScImmutableTxHashArray {
	int objId;

	public ScImmutableTxHashArray(int objId) {
		this.objId = objId;
	}

	public ScImmutableTxHash GetTxHash(int index) {
		return new ScImmutableTxHash(objId, index);
	}

	public int Length() {
		return (int) Host.GetInt(objId, Keys.KeyLength());
	}
}
