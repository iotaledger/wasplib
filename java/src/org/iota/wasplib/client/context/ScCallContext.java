// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.context;

import org.iota.wasplib.client.Host;
import org.iota.wasplib.client.Keys;
import org.iota.wasplib.client.hashtypes.ScAgent;
import org.iota.wasplib.client.hashtypes.ScColor;
import org.iota.wasplib.client.mutable.ScMutableMap;
import org.iota.wasplib.client.mutable.ScMutableMapArray;
import org.iota.wasplib.client.mutable.ScMutableString;

public class ScCallContext {
	ScMutableMap root;

	public ScCallContext() {
		root = new ScMutableMap(1);
	}

	public ScAccount Account() {
		return new ScAccount(root.GetMap("account").Immutable());
	}

	public ScCallInfo Call(String contract, String function) {
		ScMutableMapArray calls = root.GetMapArray("calls");
		ScCallInfo request = new ScCallInfo(calls.GetMap(calls.Length()));
		request.Contract(contract);
		request.Function(function);
		return request;
	}

	public ScCallInfo CallSelf(String function) {
		ScMutableMapArray calls = root.GetMapArray("calls");
		ScCallInfo request = new ScCallInfo(calls.GetMap(calls.Length()));
		request.Function(function);
		return request;
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

	public ScCallInfo Post(String contract, String function, long delay) {
		ScMutableMapArray calls = root.GetMapArray("calls");
		ScCallInfo request = new ScCallInfo(calls.GetMap(calls.Length()));
		request.Contract(contract);
		request.Function(function);
		request.Delay(delay);
		return request;
	}

	public ScCallInfo PostSelf(String function, long delay) {
		ScMutableMapArray calls = root.GetMapArray("calls");
		ScCallInfo request = new ScCallInfo(calls.GetMap(calls.Length()));
		request.Function(function);
		request.Delay(delay);
		return request;
	}

	public ScRequest Request() {
		return new ScRequest(root.GetMap("request").Immutable());
	}

	public ScMutableMap State() {
		return root.GetMap("state");
	}

	public ScLog TimestampedLog(String key) {
		return new ScLog(root.GetMap("logs").GetMapArray(key));
	}

	public void Trace(String text) {
		Host.SetString(1, Keys.KeyTrace(), text);
	}

	public void Transfer(ScAgent agent, ScColor color, long amount) {
		ScMutableMapArray transfers = root.GetMapArray("transfers");
		ScTransfer xfer = new ScTransfer(transfers.GetMap(transfers.Length()));
		xfer.Agent(agent);
		xfer.Color(color);
		xfer.Amount(amount);
	}

	public ScUtility Utility() {
		return new ScUtility(root.GetMap("utility"));
	}
}
