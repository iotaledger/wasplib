// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.mutable;

import org.iota.wasplib.client.host.Host;
import org.iota.wasplib.client.host.ScType;
import org.iota.wasplib.client.immutable.ScImmutableMapArray;

public class ScMutableMapArray {
	int objId;

	public ScMutableMapArray(int objId) {
		this.objId = objId;
	}

	public void Clear() {
		Host.SetClear(objId);
	}

	public ScMutableMap GetMap(int index) {
		int mapId = Host.GetObjectId(objId, index, ScType.TYPE_MAP);
		return new ScMutableMap(mapId);
	}

	public ScImmutableMapArray Immutable() {
		return new ScImmutableMapArray(objId);
	}

	public int Length() {
		return Host.GetLength(objId);
	}
}
