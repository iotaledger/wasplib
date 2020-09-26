package org.iota.wasplib.client.immutable;

import org.iota.wasplib.client.Host;
import org.iota.wasplib.client.Keys;

public class ScImmutableStringArray {
	int objId;

	public ScImmutableStringArray(int objId) {
		this.objId = objId;
	}

	public ScImmutableString GetString(int index) {
		return new ScImmutableString(objId, index);
	}

	public int Length() {
		return (int) Host.GetInt(objId, Keys.KeyLength());
	}
}
