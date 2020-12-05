// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.immutable;

import org.iota.wasplib.client.Host;
import org.iota.wasplib.client.Keys;
import org.iota.wasplib.client.ScType;

public class ScImmutableMap {
	int objId;

	public ScImmutableMap(int objId) {
		this.objId = objId;
	}

	public ScImmutableAddress GetAddress(String key) {
		return new ScImmutableAddress(objId, Host.GetKeyId(key));
	}

	public ScImmutableAddressArray GetAddressArray(String key) {
		int arrId = Host.GetObjectId(objId, Host.GetKeyId(key), ScType.TYPE_BYTES_ARRAY);
		return new ScImmutableAddressArray(arrId);
	}

	public ScImmutableAgent GetAgent(String key) {
		return new ScImmutableAgent(objId, Host.GetKeyId(key));
	}

	public ScImmutableAgentArray GetAgentArray(String key) {
		int arrId = Host.GetObjectId(objId, Host.GetKeyId(key), ScType.TYPE_BYTES_ARRAY);
		return new ScImmutableAgentArray(arrId);
	}

	public ScImmutableBytes GetBytes(String key) {
		return new ScImmutableBytes(objId, Host.GetKeyId(key));
	}

	public ScImmutableBytesArray GetBytesArray(String key) {
		int arrId = Host.GetObjectId(objId, Host.GetKeyId(key), ScType.TYPE_BYTES_ARRAY);
		return new ScImmutableBytesArray(arrId);
	}

	public ScImmutableColor GetColor(String key) {
		return new ScImmutableColor(objId, Host.GetKeyId(key));
	}

	public ScImmutableColorArray GetColorArray(String key) {
		int arrId = Host.GetObjectId(objId, Host.GetKeyId(key), ScType.TYPE_BYTES_ARRAY);
		return new ScImmutableColorArray(arrId);
	}

	public ScImmutableInt GetInt(String key) {
		return new ScImmutableInt(objId, Host.GetKeyId(key));
	}

	public ScImmutableIntArray GetIntArray(String key) {
		int arrId = Host.GetObjectId(objId, Host.GetKeyId(key), ScType.TYPE_INT_ARRAY);
		return new ScImmutableIntArray(arrId);
	}

	public ScImmutableKeyMap GetKeyMap(String key) {
		int mapId = Host.GetObjectId(objId, Host.GetKeyId(key), ScType.TYPE_MAP);
		return new ScImmutableKeyMap(mapId);
	}

	public ScImmutableMap GetMap(String key) {
		int mapId = Host.GetObjectId(objId, Host.GetKeyId(key), ScType.TYPE_MAP);
		return new ScImmutableMap(mapId);
	}

	public ScImmutableMapArray GetMapArray(String key) {
		int arrId = Host.GetObjectId(objId, Host.GetKeyId(key), ScType.TYPE_MAP_ARRAY);
		return new ScImmutableMapArray(arrId);
	}

	public ScImmutableString GetString(String key) {
		return new ScImmutableString(objId, Host.GetKeyId(key));
	}

	public ScImmutableStringArray GetStringArray(String key) {
		int arrId = Host.GetObjectId(objId, Host.GetKeyId(key), ScType.TYPE_STRING_ARRAY);
		return new ScImmutableStringArray(arrId);
	}

	public int Length() {
		return (int) Host.GetInt(objId, Keys.KeyLength());
	}
}
