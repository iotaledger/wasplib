// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasplib.client.context;

import org.iota.wasplib.client.hashtypes.ScColor;
import org.iota.wasplib.client.immutable.ScImmutableMap;

public class ScBalances {
	ScImmutableMap balances;

	ScBalances(ScImmutableMap balances) {
		this.balances = balances;
	}

	public long Balance(ScColor color) {
		return balances.GetInt(color).Value();
	}

	public ScColor Minted() {
		return new ScColor(balances.GetBytes(ScColor.MINT).Value());
	}
}
