// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.immutable;

import org.iota.wasplib.client.Host;
import org.iota.wasplib.client.Keys;

public class ScImmutableRequestIdArray {
	int objId;

	public ScImmutableRequestIdArray(int objId) {
		this.objId = objId;
	}

	public ScImmutableRequestId GetRequestId(int index) {
		return new ScImmutableRequestId(objId, index);
	}

	public int Length() {
		return (int) Host.GetInt(objId, Keys.KeyLength());
	}
}
