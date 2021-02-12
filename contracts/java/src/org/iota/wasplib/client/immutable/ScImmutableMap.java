// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.immutable;

import org.iota.wasplib.client.host.Host;
import org.iota.wasplib.client.host.ScType;
import org.iota.wasplib.client.keys.MapKey;

public class ScImmutableMap {
	int objId;

	public ScImmutableMap(int objId) {
		this.objId = objId;
	}

	public ScImmutableAddress GetAddress(MapKey key) {
		return new ScImmutableAddress(objId, key.GetId());
	}

	public ScImmutableAddressArray GetAddressArray(MapKey key) {
		int arrId = Host.GetObjectId(objId, key.GetId(), ScType.TYPE_ADDRESS | ScType.TYPE_ARRAY);
		return new ScImmutableAddressArray(arrId);
	}

	public ScImmutableAgent GetAgent(MapKey key) {
		return new ScImmutableAgent(objId, key.GetId());
	}

	public ScImmutableAgentArray GetAgentArray(MapKey key) {
		int arrId = Host.GetObjectId(objId, key.GetId(), ScType.TYPE_AGENT | ScType.TYPE_ARRAY);
		return new ScImmutableAgentArray(arrId);
	}

	public ScImmutableBytes GetBytes(MapKey key) {
		return new ScImmutableBytes(objId, key.GetId());
	}

	public ScImmutableBytesArray GetBytesArray(MapKey key) {
		int arrId = Host.GetObjectId(objId, key.GetId(), ScType.TYPE_BYTES | ScType.TYPE_ARRAY);
		return new ScImmutableBytesArray(arrId);
	}

	public ScImmutableChainId GetChainId(MapKey key) {
		return new ScImmutableChainId(objId, key.GetId());
	}

	public ScImmutableColor GetColor(MapKey key) {
		return new ScImmutableColor(objId, key.GetId());
	}

	public ScImmutableColorArray GetColorArray(MapKey key) {
		int arrId = Host.GetObjectId(objId, key.GetId(), ScType.TYPE_COLOR | ScType.TYPE_ARRAY);
		return new ScImmutableColorArray(arrId);
	}

	public ScImmutableContractId GetContractId(MapKey key) {
		return new ScImmutableContractId(objId, key.GetId());
	}

	public ScImmutableHash GetHash(MapKey key) {
		return new ScImmutableHash(objId, key.GetId());
	}

	public ScImmutableHname GetHname(MapKey key) {
		return new ScImmutableHname(objId, key.GetId());
	}

	public ScImmutableHashArray GetHashArray(MapKey key) {
		int arrId = Host.GetObjectId(objId, key.GetId(), ScType.TYPE_HASH | ScType.TYPE_ARRAY);
		return new ScImmutableHashArray(arrId);
	}

	public ScImmutableInt GetInt(MapKey key) {
		return new ScImmutableInt(objId, key.GetId());
	}

	public ScImmutableIntArray GetIntArray(MapKey key) {
		int arrId = Host.GetObjectId(objId, key.GetId(), ScType.TYPE_INT | ScType.TYPE_ARRAY);
		return new ScImmutableIntArray(arrId);
	}

	public ScImmutableMap GetMap(MapKey key) {
		int mapId = Host.GetObjectId(objId, key.GetId(), ScType.TYPE_MAP);
		return new ScImmutableMap(mapId);
	}

	public ScImmutableMapArray GetMapArray(MapKey key) {
		int arrId = Host.GetObjectId(objId, key.GetId(), ScType.TYPE_MAP | ScType.TYPE_ARRAY);
		return new ScImmutableMapArray(arrId);
	}

	public ScImmutableString GetString(MapKey key) {
		return new ScImmutableString(objId, key.GetId());
	}

	public ScImmutableStringArray GetStringArray(MapKey key) {
		int arrId = Host.GetObjectId(objId, key.GetId(), ScType.TYPE_STRING | ScType.TYPE_ARRAY);
		return new ScImmutableStringArray(arrId);
	}

	public int Length() {
		return Host.GetLength(objId);
	}
}
