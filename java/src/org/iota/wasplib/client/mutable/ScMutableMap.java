// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.mutable;

import org.iota.wasplib.client.Host;
import org.iota.wasplib.client.KeyId;
import org.iota.wasplib.client.Keys;
import org.iota.wasplib.client.ScType;
import org.iota.wasplib.client.immutable.ScImmutableMap;

public class ScMutableMap {
	int objId;

	public ScMutableMap(int objId) {
		this.objId = objId;
	}

	public void Clear() {
		Host.SetInt(objId, Keys.KeyLength(), 0);
	}

	public ScMutableAddress GetAddress(KeyId key) {
		return new ScMutableAddress(objId, key.GetId());
	}

	public ScMutableAddressArray GetAddressArray(KeyId key) {
		int arrId = Host.GetObjectId(objId, key.GetId(), ScType.TYPE_BYTES_ARRAY);
		return new ScMutableAddressArray(arrId);
	}

	public ScMutableAgent GetAgent(KeyId key) {
		return new ScMutableAgent(objId, key.GetId());
	}

	public ScMutableAgentArray GetAgentArray(KeyId key) {
		int arrId = Host.GetObjectId(objId, key.GetId(), ScType.TYPE_BYTES_ARRAY);
		return new ScMutableAgentArray(arrId);
	}

	public ScMutableBytes GetBytes(KeyId key) {
		return new ScMutableBytes(objId, key.GetId());
	}

	public ScMutableBytesArray GetBytesArray(KeyId key) {
		int arrId = Host.GetObjectId(objId, key.GetId(), ScType.TYPE_BYTES_ARRAY);
		return new ScMutableBytesArray(arrId);
	}

	public ScMutableColor GetColor(KeyId key) {
		return new ScMutableColor(objId, key.GetId());
	}

	public ScMutableColorArray GetColorArray(KeyId key) {
		int arrId = Host.GetObjectId(objId, key.GetId(), ScType.TYPE_BYTES_ARRAY);
		return new ScMutableColorArray(arrId);
	}

	public ScMutableInt GetInt(KeyId key) {
		return new ScMutableInt(objId, key.GetId());
	}

	public ScMutableIntArray GetIntArray(KeyId key) {
		int arrId = Host.GetObjectId(objId, key.GetId(), ScType.TYPE_INT_ARRAY);
		return new ScMutableIntArray(arrId);
	}

	public ScMutableMap GetMap(KeyId key) {
		int mapId = Host.GetObjectId(objId, key.GetId(), ScType.TYPE_MAP);
		return new ScMutableMap(mapId);
	}

	public ScMutableMapArray GetMapArray(KeyId key) {
		int arrId = Host.GetObjectId(objId, key.GetId(), ScType.TYPE_MAP_ARRAY);
		return new ScMutableMapArray(arrId);
	}

	public ScMutableString GetString(KeyId key) {
		return new ScMutableString(objId, key.GetId());
	}

	public ScMutableStringArray GetStringArray(KeyId key) {
		int arrId = Host.GetObjectId(objId, key.GetId(), ScType.TYPE_STRING_ARRAY);
		return new ScMutableStringArray(arrId);
	}

	public ScImmutableMap Immutable() {
		return new ScImmutableMap(objId);
	}

	public int Length() {
		return (int) Host.GetInt(objId, Keys.KeyLength());
	}
}
