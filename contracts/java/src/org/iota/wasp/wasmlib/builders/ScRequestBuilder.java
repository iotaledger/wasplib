// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.wasmlib.builders;

import org.iota.wasp.wasmlib.hashtypes.*;
import org.iota.wasp.wasmlib.host.*;
import org.iota.wasp.wasmlib.immutable.*;
import org.iota.wasp.wasmlib.keys.*;
import org.iota.wasp.wasmlib.mutable.*;

public abstract class ScRequestBuilder {
	ScMutableMap request;

	protected ScRequestBuilder(String key, String function) {
		ScMutableMapArray requests = Host.root.GetMapArray(new Key(key));
		request = requests.GetMap(requests.Length());
		request.GetString(Key.Function).SetValue(function);
	}

	protected void contract(String contract) {
		request.GetString(Key.Contract).SetValue(contract);
	}

	protected void exec(long delay) {
		request.GetInt(Key.Delay).SetValue(delay);
	}

	protected ScMutableMap params() {
		return request.GetMap(Key.Params);
	}

	protected ScImmutableMap results() {
		return request.GetMap(Key.Results).Immutable();
	}

	protected void transfer(ScColor color, long amount) {
		ScMutableMap transfers = request.GetMap(Key.Transfers);
		transfers.GetInt(color).SetValue(amount);
	}
}
