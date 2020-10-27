package org.iota.wasplib.client.mutable;

import org.iota.wasplib.client.Host;
import org.iota.wasplib.client.Keys;
import org.iota.wasplib.client.immutable.ScImmutableColorArray;

public class ScMutableColorArray {
	int objId;

	public ScMutableColorArray(int objId) {
		this.objId = objId;
	}

	public void Clear() {
		Host.SetInt(objId, Keys.KeyLength(), 0);
	}

	public ScMutableColor GetColor(int index) {
		return new ScMutableColor(objId, index);
	}

	public ScImmutableColorArray Immutable() {
		return new ScImmutableColorArray(objId);
	}

	public int Length() {
		return (int) Host.GetInt(objId, Keys.KeyLength());
	}
}
