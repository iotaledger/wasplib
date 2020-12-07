// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.immutable;

import org.iota.wasplib.client.host.Host;
import org.iota.wasplib.client.keys.Keys;

public class ScImmutableStringArray {
	int objId;

	public ScImmutableStringArray(int objId) {
		this.objId = objId;
	}

	public ScImmutableString GetString(int index) {
		return new ScImmutableString(objId, index);
	}

	public int Length() {
		return (int) Host.GetInt(objId, Keys.KeyLength());
	}
}
