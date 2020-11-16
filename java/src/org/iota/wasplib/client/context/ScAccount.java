// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.context;

import org.iota.wasplib.client.hashtypes.ScColor;
import org.iota.wasplib.client.immutable.ScImmutableColorArray;
import org.iota.wasplib.client.immutable.ScImmutableMap;

public class ScAccount {
	ScImmutableMap account;

	ScAccount(ScImmutableMap account) {
		this.account = account;
	}

	public long Balance(ScColor color) {
		return account.GetKeyMap("balance").GetInt(color.toBytes()).Value();
	}

	public ScImmutableColorArray Colors() {
		return account.GetColorArray("colors");
	}
}
