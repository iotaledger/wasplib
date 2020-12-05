// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.immutable;

import org.iota.wasplib.client.Host;
import org.iota.wasplib.client.Keys;
import org.iota.wasplib.client.ScType;

public class ScImmutableKeyMap {
	int objId;

	public ScImmutableKeyMap(int objId) {
		this.objId = objId;
	}

	public ScImmutableAddress GetAddress(byte[] key) {
		return new ScImmutableAddress(objId, Host.GetKey(key));
	}

	public ScImmutableAddressArray GetAddressArray(byte[] key) {
		int arrId = Host.GetObjectId(objId, Host.GetKey(key), ScType.TYPE_BYTES_ARRAY);
		return new ScImmutableAddressArray(arrId);
	}

	public ScImmutableAgent GetAgent(byte[] key) {
		return new ScImmutableAgent(objId, Host.GetKey(key));
	}

	public ScImmutableAgentArray GetAgentArray(byte[] key) {
		int arrId = Host.GetObjectId(objId, Host.GetKey(key), ScType.TYPE_BYTES_ARRAY);
		return new ScImmutableAgentArray(arrId);
	}

	public ScImmutableBytes GetBytes(byte[] key) {
		return new ScImmutableBytes(objId, Host.GetKey(key));
	}

	public ScImmutableBytesArray GetBytesArray(byte[] key) {
		int arrId = Host.GetObjectId(objId, Host.GetKey(key), ScType.TYPE_BYTES_ARRAY);
		return new ScImmutableBytesArray(arrId);
	}

	public ScImmutableColor GetColor(byte[] key) {
		return new ScImmutableColor(objId, Host.GetKey(key));
	}

	public ScImmutableColorArray GetColorArray(byte[] key) {
		int arrId = Host.GetObjectId(objId, Host.GetKey(key), ScType.TYPE_BYTES_ARRAY);
		return new ScImmutableColorArray(arrId);
	}

	public ScImmutableInt GetInt(byte[] key) {
		return new ScImmutableInt(objId, Host.GetKey(key));
	}

	public ScImmutableIntArray GetIntArray(byte[] key) {
		int arrId = Host.GetObjectId(objId, Host.GetKey(key), ScType.TYPE_INT_ARRAY);
		return new ScImmutableIntArray(arrId);
	}

	public ScImmutableKeyMap GetKeyMap(byte[] key) {
		int mapId = Host.GetObjectId(objId, Host.GetKey(key), ScType.TYPE_MAP);
		return new ScImmutableKeyMap(mapId);
	}

	public ScImmutableMap GetMap(byte[] key) {
		int mapId = Host.GetObjectId(objId, Host.GetKey(key), ScType.TYPE_MAP);
		return new ScImmutableMap(mapId);
	}

	public ScImmutableMapArray GetMapArray(byte[] key) {
		int arrId = Host.GetObjectId(objId, Host.GetKey(key), ScType.TYPE_MAP_ARRAY);
		return new ScImmutableMapArray(arrId);
	}

	public ScImmutableString GetString(byte[] key) {
		return new ScImmutableString(objId, Host.GetKey(key));
	}

	public ScImmutableStringArray GetStringArray(byte[] key) {
		int arrId = Host.GetObjectId(objId, Host.GetKey(key), ScType.TYPE_STRING_ARRAY);
		return new ScImmutableStringArray(arrId);
	}

	public int Length() {
		return (int) Host.GetInt(objId, Keys.KeyLength());
	}
}
