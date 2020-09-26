package org.iota.wasplib.client.context;

import org.iota.wasplib.client.mutable.ScMutableMap;
import org.iota.wasplib.client.mutable.ScMutableString;

public class ScUtility {
	ScMutableMap utility;

	ScUtility(ScMutableMap utility) {
		this.utility = utility;
	}

	public String Hash(String value) {
		ScMutableString hash = utility.GetString("hash");
		hash.SetValue(value);
		return hash.Value();
	}
}
