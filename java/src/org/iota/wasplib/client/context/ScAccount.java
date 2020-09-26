package org.iota.wasplib.client.context;

import org.iota.wasplib.client.immutable.ScImmutableMap;
import org.iota.wasplib.client.immutable.ScImmutableStringArray;

public class ScAccount {
	ScImmutableMap account;

	ScAccount(ScImmutableMap account) {
		this.account = account;
	}

	public long Balance(String color) {
		return account.GetMap("balance").GetInt(color).Value();
	}

	public ScImmutableStringArray Colors() {
		return account.GetStringArray("colors");
	}
}
