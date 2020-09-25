package org.iota.wasplib.client.context;

import org.iota.wasplib.client.mutable.ScMutableMap;

public class ScEvent {
	ScMutableMap event;

	ScEvent(ScMutableMap event) {
		this.event = event;
	}

	public void Code(long code) {
		event.GetInt("code").SetValue(code);
	}

	public void Contract(String contract) {
		event.GetString("contract").SetValue(contract);
	}

	public void Delay(long delay) {
		event.GetInt("delay").SetValue(delay);
	}

	public void Function(String function) {
		event.GetString("function").SetValue(function);
	}

	public ScMutableMap Params() {
		return event.GetMap("params");
	}
}
