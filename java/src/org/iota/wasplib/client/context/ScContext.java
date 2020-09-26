package org.iota.wasplib.client.context;

import org.iota.wasplib.client.Host;
import org.iota.wasplib.client.Keys;
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

	public ScMutableMap Event(String contract, String function, long delay) {
		ScMutableMapArray events = root.GetMapArray("events");
		ScEvent evt = new ScEvent(events.GetMap(events.Length()));
		evt.Contract(contract);
		evt.Function(function);
		evt.Delay(delay);
		return evt.Params();
	}

	public ScMutableMap EventWithCode(String contract, long code, long delay) {
		ScMutableMapArray events = root.GetMapArray("events");
		ScEvent evt = new ScEvent(events.GetMap(events.Length()));
		evt.Contract(contract);
		evt.Code(code);
		evt.Delay(delay);
		return evt.Params();
	}

	public void Log(String text) {
		Host.SetString(1, Keys.KeyLog(), text);
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
		return new ScLog(root.GetMap("logs").GetMap(key));
	}

	public void Trace(String text) {
		Host.SetString(1, Keys.KeyTrace(), text);
	}

	public void Transfer(String address, String color, long amount) {
		ScMutableMapArray transfers = root.GetMapArray("transfers");
		ScTransfer xfer = new ScTransfer(transfers.GetMap(transfers.Length()));
		xfer.Address(address);
		xfer.Color(color);
		xfer.Amount(amount);
	}

	public ScUtility Utility() {
		return new ScUtility(root.GetMap("utility"));
	}
}
