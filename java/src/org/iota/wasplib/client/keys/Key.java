package org.iota.wasplib.client.keys;

import org.iota.wasplib.client.host.Host;

public class Key implements KeyId {
	String key;

	public Key(String key) {
		this.key = key;
	}

	@Override
	public int GetId() {
		return Host.GetKeyId(key);
	}
}
