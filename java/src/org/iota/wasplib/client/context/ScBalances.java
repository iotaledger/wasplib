// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.context;

import org.iota.wasplib.client.hashtypes.ScColor;
import org.iota.wasplib.client.immutable.ScImmutableKeyMap;

public class ScBalances {
	ScImmutableKeyMap balances;

	ScBalances(ScImmutableKeyMap balances) {
		this.balances = balances;
	}

	public long Balance(ScColor color) {
		return balances.GetInt(color.toBytes()).Value();
	}

	public ScColor Minted() {
		byte[] mintKey = ScColor.MINT.toBytes();
		return new ScColor(balances.GetBytes(mintKey).Value());
	}
}
