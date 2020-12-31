// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.immutable;

import org.iota.wasplib.client.host.Host;

public class ScImmutableHashArray {
	int objId;

	public ScImmutableHashArray(int objId) {
		this.objId = objId;
	}

	public ScImmutableHash GetHash(int index) {
		return new ScImmutableHash(objId, index);
	}

	public int Length() {
		return Host.GetLength(objId);
	}
}
