package org.iota.wasplib.client.context;

import org.iota.wasplib.client.mutable.ScMutableMap;

public class ScLog {
	ScMutableMap log;

	ScLog(ScMutableMap log) {
		this.log = log;
	}

	public void Append(long timestamp, byte[] data) {
		log.GetInt("timestamp").SetValue(timestamp);
		log.GetBytes("data").SetValue(data);
	}

	public int Length() {
		return (int) log.GetInt("length").Value();
	}
}
