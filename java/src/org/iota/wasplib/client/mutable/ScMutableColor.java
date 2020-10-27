package org.iota.wasplib.client.mutable;

import org.iota.wasplib.client.Host;
import org.iota.wasplib.client.hashtypes.ScColor;

public class ScMutableColor {
	int objId;
	int keyId;

	public ScMutableColor(int objId, int keyId) {
		this.objId = objId;
		this.keyId = keyId;
	}

	public boolean Exists() {
		return Host.Exists(objId, keyId);
	}

	public void SetValue(ScColor value) {
		Host.SetBytes(objId, keyId, value.toBytes());
	}

	public ScColor Value() {
		return new ScColor(Host.GetBytes(objId, keyId));
	}
}
