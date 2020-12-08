// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.mutable;

import org.iota.wasplib.client.host.Host;
import org.iota.wasplib.client.immutable.ScImmutableAgentArray;
import org.iota.wasplib.client.keys.Key;

public class ScMutableAgentArray {
	int objId;

	public ScMutableAgentArray(int objId) {
		this.objId = objId;
	}

	public void Clear() {
		Host.SetInt(objId, Key.KEY_LENGTH, 0);
	}

	public ScMutableAgent GetAgent(int index) {
		return new ScMutableAgent(objId, index);
	}

	public ScImmutableAgentArray Immutable() {
		return new ScImmutableAgentArray(objId);
	}

	public int Length() {
		return (int) Host.GetInt(objId, Key.KEY_LENGTH);
	}
}
