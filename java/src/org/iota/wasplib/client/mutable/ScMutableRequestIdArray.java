package org.iota.wasplib.client.mutable;

import org.iota.wasplib.client.Host;
import org.iota.wasplib.client.Keys;
import org.iota.wasplib.client.immutable.ScImmutableRequestIdArray;

public class ScMutableRequestIdArray {
	int objId;

	public ScMutableRequestIdArray(int objId) {
		this.objId = objId;
	}

	public void Clear() {
		Host.SetInt(objId, Keys.KeyLength(), 0);
	}

	public ScMutableRequestId GetRequestId(int index) {
		return new ScMutableRequestId(objId, index);
	}

	public ScImmutableRequestIdArray Immutable() {
		return new ScImmutableRequestIdArray(objId);
	}

	public int Length() {
		return (int) Host.GetInt(objId, Keys.KeyLength());
	}
}
