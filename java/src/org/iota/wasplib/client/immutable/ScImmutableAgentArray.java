// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.immutable;

import org.iota.wasplib.client.Host;
import org.iota.wasplib.client.Keys;

public class ScImmutableAgentArray {
	int objId;

	public ScImmutableAgentArray(int objId) {
		this.objId = objId;
	}

	public ScImmutableAgent GetAgent(int index) {
		return new ScImmutableAgent(objId, index);
	}

	public int Length() {
		return (int) Host.GetInt(objId, Keys.KeyLength());
	}
}
