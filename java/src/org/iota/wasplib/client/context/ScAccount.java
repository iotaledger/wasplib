package org.iota.wasplib.client.context;

import org.iota.wasplib.client.hashtypes.ScColor;
import org.iota.wasplib.client.immutable.ScImmutableMap;

public class ScAccount {
	ScImmutableMap account;

	ScAccount(ScImmutableMap account) {
		this.account = account;
	}

	public long Balance(ScColor color) {
		return account.GetMap("balance").GetInt(color.toBytes()).Value();
	}

	public ScColors Colors() {
		return new ScColors(account.GetStringArray("colors"));
	}
}
