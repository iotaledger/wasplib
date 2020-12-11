// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.immutable;

import org.iota.wasplib.client.host.Host;

public class ScImmutableStringArray {
	int objId;

	public ScImmutableStringArray(int objId) {
		this.objId = objId;
	}

	public ScImmutableString GetString(int index) {
		return new ScImmutableString(objId, index);
	}

	public int Length() {
		return Host.GetLength(objId);
	}
}
