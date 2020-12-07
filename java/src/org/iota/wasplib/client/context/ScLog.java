// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.context;

import org.iota.wasplib.client.keys.Key;
import org.iota.wasplib.client.mutable.ScMutableMap;
import org.iota.wasplib.client.mutable.ScMutableMapArray;

public class ScLog {
	ScMutableMapArray log;

	ScLog(ScMutableMapArray log) {
		this.log = log;
	}

	public void Append(long timestamp, byte[] data) {
		ScMutableMap logEntry = log.GetMap(log.Length());
		logEntry.GetInt(new Key("timestamp")).SetValue(timestamp);
		logEntry.GetBytes(new Key("data")).SetValue(data);
	}

	public int Length() {
		return log.Length();
	}
}
