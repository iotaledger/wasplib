// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.context;

import org.iota.wasplib.client.immutable.ScImmutableMap;
import org.iota.wasplib.client.immutable.ScImmutableMapArray;
import org.iota.wasplib.client.keys.Key;
import org.iota.wasplib.client.keys.KeyId;

public class ScViewContext extends ScBaseContext {
	public ScViewContext() {
	}

	public ScImmutableMap State() {
		return root.GetMap(new Key("state")).Immutable();
	}

	public ScImmutableMapArray TimestampedLog(KeyId key) {
		return root.GetMap(new Key("logs")).GetMapArray(key).Immutable();
	}
}
