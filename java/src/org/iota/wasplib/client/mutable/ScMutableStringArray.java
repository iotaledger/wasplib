package org.iota.wasplib.client.mutable;

import org.iota.wasplib.client.Host;
import org.iota.wasplib.client.Keys;
import org.iota.wasplib.client.immutable.ScImmutableStringArray;

public class ScMutableStringArray {
	int objId;

	public ScMutableStringArray(int objId) {
		this.objId = objId;
	}

	public void Clear() {
		Host.SetInt(objId, Keys.KeyLength(), 0);
	}

	public ScMutableString GetString(int index) {
		return new ScMutableString(objId, index);
	}

	public ScImmutableStringArray Immutable() {
		return new ScImmutableStringArray(objId);
	}

	public int Length() {
		return (int) Host.GetInt(objId, Keys.KeyLength());
	}
}
