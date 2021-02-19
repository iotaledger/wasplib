// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.wasmlib.builders;

import org.iota.wasp.wasmlib.hashtypes.*;
import org.iota.wasp.wasmlib.keys.*;
import org.iota.wasp.wasmlib.mutable.*;

public class ScPostBuilder extends ScRequestBuilder {
	public ScPostBuilder(String function) {
		super("posts", function);
	}

	public ScPostBuilder Chain(ScAddress chain) {
		request.GetAddress(Key.Chain).SetValue(chain);
		return this;
	}

	public ScPostBuilder Contract(String contract) {
		contract(contract);
		return this;
	}

	public ScMutableMap Params() {
		return params();
	}

	public void Post(long delay) {
		exec(delay);
	}

	public ScPostBuilder Transfer(ScColor color, long amount) {
		transfer(color, amount);
		return this;
	}
}
