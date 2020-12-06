// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.immutable;

import org.iota.wasplib.client.Host;
import org.iota.wasplib.client.KeyId;
import org.iota.wasplib.client.Keys;
import org.iota.wasplib.client.ScType;

public class ScImmutableMap {
	int objId;

	public ScImmutableMap(int objId) {
		this.objId = objId;
	}

	public ScImmutableAddress GetAddress(KeyId key) {
		return new ScImmutableAddress(objId, key.GetId());
	}

	public ScImmutableAddressArray GetAddressArray(KeyId key) {
		int arrId = Host.GetObjectId(objId, key.GetId(), ScType.TYPE_BYTES_ARRAY);
		return new ScImmutableAddressArray(arrId);
	}

	public ScImmutableAgent GetAgent(KeyId key) {
		return new ScImmutableAgent(objId, key.GetId());
	}

	public ScImmutableAgentArray GetAgentArray(KeyId key) {
		int arrId = Host.GetObjectId(objId, key.GetId(), ScType.TYPE_BYTES_ARRAY);
		return new ScImmutableAgentArray(arrId);
	}

	public ScImmutableBytes GetBytes(KeyId key) {
		return new ScImmutableBytes(objId, key.GetId());
	}

	public ScImmutableBytesArray GetBytesArray(KeyId key) {
		int arrId = Host.GetObjectId(objId, key.GetId(), ScType.TYPE_BYTES_ARRAY);
		return new ScImmutableBytesArray(arrId);
	}

	public ScImmutableColor GetColor(KeyId key) {
		return new ScImmutableColor(objId, key.GetId());
	}

	public ScImmutableColorArray GetColorArray(KeyId key) {
		int arrId = Host.GetObjectId(objId, key.GetId(), ScType.TYPE_BYTES_ARRAY);
		return new ScImmutableColorArray(arrId);
	}

	public ScImmutableInt GetInt(KeyId key) {
		return new ScImmutableInt(objId, key.GetId());
	}

	public ScImmutableIntArray GetIntArray(KeyId key) {
		int arrId = Host.GetObjectId(objId, key.GetId(), ScType.TYPE_INT_ARRAY);
		return new ScImmutableIntArray(arrId);
	}

	public ScImmutableMap GetMap(KeyId key) {
		int mapId = Host.GetObjectId(objId, key.GetId(), ScType.TYPE_MAP);
		return new ScImmutableMap(mapId);
	}

	public ScImmutableMapArray GetMapArray(KeyId key) {
		int arrId = Host.GetObjectId(objId, key.GetId(), ScType.TYPE_MAP_ARRAY);
		return new ScImmutableMapArray(arrId);
	}

	public ScImmutableString GetString(KeyId key) {
		return new ScImmutableString(objId, key.GetId());
	}

	public ScImmutableStringArray GetStringArray(KeyId key) {
		int arrId = Host.GetObjectId(objId, key.GetId(), ScType.TYPE_STRING_ARRAY);
		return new ScImmutableStringArray(arrId);
	}

	public int Length() {
		return (int) Host.GetInt(objId, Keys.KeyLength());
	}
}
