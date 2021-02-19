// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.wasmlib.mutable;

import org.iota.wasp.wasmlib.host.*;
import org.iota.wasp.wasmlib.immutable.*;
import org.iota.wasp.wasmlib.keys.*;

public class ScMutableMap {
	int objId;

	public ScMutableMap(int objId) {
		this.objId = objId;
	}

	public void Clear() {
		Host.SetClear(objId);
	}

	public ScMutableAddress GetAddress(MapKey key) {
		return new ScMutableAddress(objId, key.GetId());
	}

	public ScMutableAddressArray GetAddressArray(MapKey key) {
		int arrId = Host.GetObjectId(objId, key.GetId(), ScType.TYPE_ADDRESS | ScType.TYPE_ARRAY);
		return new ScMutableAddressArray(arrId);
	}

	public ScMutableAgentId GetAgentId(MapKey key) {
		return new ScMutableAgentId(objId, key.GetId());
	}

	public ScMutableAgentIdArray GetAgentIdArray(MapKey key) {
		int arrId = Host.GetObjectId(objId, key.GetId(), ScType.TYPE_AGENT | ScType.TYPE_ARRAY);
		return new ScMutableAgentIdArray(arrId);
	}

	public ScMutableBytes GetBytes(MapKey key) {
		return new ScMutableBytes(objId, key.GetId());
	}

	public ScMutableBytesArray GetBytesArray(MapKey key) {
		int arrId = Host.GetObjectId(objId, key.GetId(), ScType.TYPE_BYTES | ScType.TYPE_ARRAY);
		return new ScMutableBytesArray(arrId);
	}

	public ScMutableChainId GetChainId(MapKey key) {
		return new ScMutableChainId(objId, key.GetId());
	}

	public ScMutableColor GetColor(MapKey key) {
		return new ScMutableColor(objId, key.GetId());
	}

	public ScMutableColorArray GetColorArray(MapKey key) {
		int arrId = Host.GetObjectId(objId, key.GetId(), ScType.TYPE_COLOR | ScType.TYPE_ARRAY);
		return new ScMutableColorArray(arrId);
	}

	public ScMutableContractId GetContractId(MapKey key) {
		return new ScMutableContractId(objId, key.GetId());
	}

	public ScMutableHash GetHash(MapKey key) {
		return new ScMutableHash(objId, key.GetId());
	}

	public ScMutableHname GetHname(MapKey key) {
		return new ScMutableHname(objId, key.GetId());
	}

	public ScMutableHashArray GetHashArray(MapKey key) {
		int arrId = Host.GetObjectId(objId, key.GetId(), ScType.TYPE_HASH | ScType.TYPE_ARRAY);
		return new ScMutableHashArray(arrId);
	}

	public ScMutableInt GetInt(MapKey key) {
		return new ScMutableInt(objId, key.GetId());
	}

	public ScMutableIntArray GetIntArray(MapKey key) {
		int arrId = Host.GetObjectId(objId, key.GetId(), ScType.TYPE_INT | ScType.TYPE_ARRAY);
		return new ScMutableIntArray(arrId);
	}

	public ScMutableMap GetMap(MapKey key) {
		int mapId = Host.GetObjectId(objId, key.GetId(), ScType.TYPE_MAP);
		return new ScMutableMap(mapId);
	}

	public ScMutableMapArray GetMapArray(MapKey key) {
		int arrId = Host.GetObjectId(objId, key.GetId(), ScType.TYPE_MAP | ScType.TYPE_ARRAY);
		return new ScMutableMapArray(arrId);
	}

	public ScMutableString GetString(MapKey key) {
		return new ScMutableString(objId, key.GetId());
	}

	public ScMutableStringArray GetStringArray(MapKey key) {
		int arrId = Host.GetObjectId(objId, key.GetId(), ScType.TYPE_STRING | ScType.TYPE_ARRAY);
		return new ScMutableStringArray(arrId);
	}

	public ScImmutableMap Immutable() {
		return new ScImmutableMap(objId);
	}

	public int Length() {
		return Host.GetLength(objId);
	}
}
