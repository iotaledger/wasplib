package org.iota.wasplib.client.context;

import org.iota.wasplib.client.mutable.ScMutableMap;
import org.iota.wasplib.client.mutable.ScMutableStringArray;

public class ScExports {
	ScMutableStringArray exports;
	int next;

	public ScExports() {
		ScMutableMap root = new ScMutableMap(1);
		exports = root.GetStringArray("exports");
	}

	public void Add(String name) {
		next++;
		exports.GetString(next).SetValue(name);
	}

	public void AddProtected(String name) {
		next++;
		exports.GetString(next | 0x4000).SetValue(name);
	}
}
