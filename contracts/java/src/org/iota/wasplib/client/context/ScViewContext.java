// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.context;

import org.iota.wasplib.client.host.Host;
import org.iota.wasplib.client.immutable.ScImmutableMap;
import org.iota.wasplib.client.immutable.ScImmutableMapArray;
import org.iota.wasplib.client.keys.Key;
import org.iota.wasplib.client.keys.MapKey;

public class ScViewContext extends ScBaseContext {
	public ScViewContext() {
	}

	public ScImmutableMap State() {
		return Host.root.GetMap(Key.State).Immutable();
	}

	public ScImmutableMapArray TimestampedLog(MapKey key) {
		return Host.root.GetMap(Key.Logs).GetMapArray(key).Immutable();
	}
}
