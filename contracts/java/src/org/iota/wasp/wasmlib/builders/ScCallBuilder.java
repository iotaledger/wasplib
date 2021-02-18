// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.wasmlib.builders;

import org.iota.wasp.wasmlib.hashtypes.ScColor;
import org.iota.wasp.wasmlib.immutable.ScImmutableMap;
import org.iota.wasp.wasmlib.mutable.ScMutableMap;

public class ScCallBuilder extends ScRequestBuilder {
	public ScCallBuilder(String function) {
		super("calls", function);
	}

	public ScCallBuilder Call() {
		exec(-1);
		return this;
	}

	public ScCallBuilder Contract(String contract) {
		contract(contract);
		return this;
	}

	public ScMutableMap Params() {
		return params();
	}

	public ScImmutableMap Results() {
		return results();
	}

	public ScCallBuilder Transfer(ScColor color, long amount) {
		transfer(color, amount);
		return this;
	}
}
