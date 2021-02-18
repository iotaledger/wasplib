// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.wasmlib.context;

import org.iota.wasp.wasmlib.host.Host;
import org.iota.wasp.wasmlib.immutable.ScImmutableMap;
import org.iota.wasp.wasmlib.immutable.ScImmutableMapArray;
import org.iota.wasp.wasmlib.keys.Key;
import org.iota.wasp.wasmlib.keys.MapKey;

public class ScViewContext extends ScBaseContext {
	public ScViewContext() {
	}

	public ScImmutableMap State() {
		return Host.root.GetMap(Key.State).Immutable();
	}
}
