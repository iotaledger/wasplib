package org.iota.wasplib.client.immutable;

import org.iota.wasplib.client.Host;
import org.iota.wasplib.client.Keys;

public class ScImmutableColorArray {
	int objId;

	public ScImmutableColorArray(int objId) {
		this.objId = objId;
	}

	public ScImmutableColor GetColor(int index) {
		return new ScImmutableColor(objId, index);
	}

	public int Length() {
		return (int) Host.GetInt(objId, Keys.KeyLength());
	}
}
