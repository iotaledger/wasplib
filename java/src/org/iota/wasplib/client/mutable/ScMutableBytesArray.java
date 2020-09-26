package org.iota.wasplib.client.mutable;

import org.iota.wasplib.client.Host;
import org.iota.wasplib.client.Keys;
import org.iota.wasplib.client.immutable.ScImmutableBytesArray;

public class ScMutableBytesArray {
	int objId;

	public ScMutableBytesArray(int objId) {
		this.objId = objId;
	}

	public void Clear() {
		Host.SetInt(objId, Keys.KeyLength(), 0);
	}

	public ScMutableBytes GetBytes(int index) {
		return new ScMutableBytes(objId, index);
	}

	public ScImmutableBytesArray Immutable() {
		return new ScImmutableBytesArray(objId);
	}

	public int Length() {
		return (int) Host.GetInt(objId, Keys.KeyLength());
	}
}
