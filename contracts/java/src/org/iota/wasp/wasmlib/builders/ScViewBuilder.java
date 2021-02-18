// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.wasmlib.builders;

import org.iota.wasp.wasmlib.immutable.ScImmutableMap;
import org.iota.wasp.wasmlib.mutable.ScMutableMap;

public class ScViewBuilder extends ScRequestBuilder {
	public ScViewBuilder(String function) {
		super("views", function);
	}

	public ScViewBuilder Contract(String contract) {
		contract(contract);
		return this;
	}

	public ScMutableMap Params() {
		return params();
	}

	public ScImmutableMap Results() {
		return results();
	}

	public ScViewBuilder View() {
		exec(-2);
		return this;
	}
}
