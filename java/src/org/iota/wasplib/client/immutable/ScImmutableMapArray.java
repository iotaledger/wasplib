// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.immutable;

import org.iota.wasplib.client.host.Host;
import org.iota.wasplib.client.host.ScType;
import org.iota.wasplib.client.keys.Keys;

public class ScImmutableMapArray {
	int objId;

	public ScImmutableMapArray(int objId) {
		this.objId = objId;
	}

	public ScImmutableMap GetMap(int index) {
		int mapId = Host.GetObjectId(objId, index, ScType.TYPE_MAP);
		return new ScImmutableMap(mapId);
	}

	public int Length() {
		return (int) Host.GetInt(objId, Keys.KeyLength());
	}
}
