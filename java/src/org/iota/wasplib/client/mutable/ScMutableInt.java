package org.iota.wasplib.client.mutable;

import org.iota.wasplib.client.Host;

public class ScMutableInt {
	int objId;
	int keyId;

	public ScMutableInt(int objId, int keyId) {
		this.objId = objId;
		this.keyId = keyId;
	}

	public void SetValue(long value) {
		Host.SetInt(objId, keyId, value);
	}

	public long Value() {
		return Host.GetInt(objId, keyId);
	}
}
