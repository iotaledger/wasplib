// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.wasmlib.context;

import org.iota.wasp.wasmlib.host.*;
import org.iota.wasp.wasmlib.immutable.*;
import org.iota.wasp.wasmlib.keys.*;

public class ScViewContext extends ScBaseContext {
	public ScViewContext() {
	}

	public ScImmutableMap State() {
		return Host.root.GetMap(Key.State).Immutable();
	}
}
