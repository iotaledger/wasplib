package org.iota.wasplib.client.immutable;

import org.iota.wasplib.client.Host;
import org.iota.wasplib.client.Keys;

public class ScImmutableAddressArray {
	int objId;

	public ScImmutableAddressArray(int objId) {
		this.objId = objId;
	}

	public ScImmutableAddress GetAddress(int index) {
		return new ScImmutableAddress(objId, index);
	}

	public int Length() {
		return (int) Host.GetInt(objId, Keys.KeyLength());
	}
}
