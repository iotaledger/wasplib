// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.context;

import org.iota.wasplib.client.Host;
import org.iota.wasplib.client.Keys;
import org.iota.wasplib.client.immutable.ScImmutableMap;
import org.iota.wasplib.client.mutable.ScMutableMap;
import org.iota.wasplib.client.mutable.ScMutableMapArray;
import org.iota.wasplib.client.mutable.ScMutableString;

public class ScViewContext {
	ScMutableMap root;

	public ScViewContext() {
		root = new ScMutableMap(1);
	}

	public ScAccount Account() {
		return new ScAccount(root.GetMap("account").Immutable());
	}

	public ScCallInfo Call(String contract, String function) {
		ScMutableMapArray calls = root.GetMapArray("calls");
		ScCallInfo call = new ScCallInfo(calls.GetMap(calls.Length()));
		call.Contract(contract);
		call.Function(function);
		return call;
	}

	public ScCallInfo CallSelf(String function) {
		ScMutableMapArray calls = root.GetMapArray("calls");
		ScCallInfo call = new ScCallInfo(calls.GetMap(calls.Length()));
		call.Function(function);
		return call;
	}

	public ScContract Contract() {
		return new ScContract(root.GetMap("contract").Immutable());
	}

	public ScMutableString Error() {
		return root.GetString("error");
	}

	public void Log(String text) {
		Host.SetString(1, Keys.KeyLog(), text);
	}

	public ScRequest Request() {
		return new ScRequest(root.GetMap("request").Immutable());
	}

	public ScImmutableMap State() {
		return root.GetMap("state").Immutable();
	}

	public ScLog TimestampedLog(String key) {
		return new ScLog(root.GetMap("logs").GetMapArray(key));
	}

	public void Trace(String text) {
		Host.SetString(1, Keys.KeyTrace(), text);
	}

	public ScUtility Utility() {
		return new ScUtility(root.GetMap("utility"));
	}
}
