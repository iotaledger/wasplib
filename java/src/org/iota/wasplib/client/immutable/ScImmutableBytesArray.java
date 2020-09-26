package org.iota.wasplib.client.immutable;

import org.iota.wasplib.client.Host;
import org.iota.wasplib.client.Keys;

public class ScImmutableBytesArray {
	int objId;

	public ScImmutableBytesArray(int objId) {
		this.objId = objId;
	}

	public ScImmutableBytes GetBytes(int index) {
		return new ScImmutableBytes(objId, index);
	}

	public int Length() {
		return (int) Host.GetInt(objId, Keys.KeyLength());
	}
}
