// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.immutable;

import org.iota.wasplib.client.host.Host;

public class ScImmutableIntArray {
	int objId;

	public ScImmutableIntArray(int objId) {
		this.objId = objId;
	}

	public ScImmutableInt GetInt(int index) {
		return new ScImmutableInt(objId, index);
	}

	public int Length() {
		return Host.GetLength(objId);
	}
}
