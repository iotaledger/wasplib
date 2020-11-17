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

public class ScContext {
	ScMutableMap root;

	public ScContext() {
		root = new ScMutableMap(1);
	}

	public ScAccount Account() {
		return new ScAccount(root.GetMap("account").Immutable());
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

	public ScMutableMap PostRequest(ScAgent contract, String function, long delay) {
		ScMutableMapArray postedRequests = root.GetMapArray("postedRequests");
		ScPostedRequest request = new ScPostedRequest(postedRequests.GetMap(postedRequests.Length()));
		request.Contract(contract);
		request.Function(function);
		request.Delay(delay);
		return request.Params();
	}

	public long Random(long max) {
		long rnd = root.GetInt("random").Value();
		return Long.remainderUnsigned(rnd, max);
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
