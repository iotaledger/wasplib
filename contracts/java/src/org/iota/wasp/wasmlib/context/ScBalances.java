// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.wasmlib.context;

import org.iota.wasp.wasmlib.hashtypes.*;
import org.iota.wasp.wasmlib.immutable.*;
import org.iota.wasp.wasmlib.keys.*;

public class ScBalances {
	ScImmutableMap balances;

	ScBalances(ScImmutableMap balances) {
		this.balances = balances;
	}

	public long Balance(ScColor color) {
		return balances.GetInt(color).Value();
	}

	public ScImmutableColorArray Colors() {
		return balances.GetColorArray(Key.Caller);
	}

	public ScColor Minted() {
		return new ScColor(balances.GetBytes(ScColor.MINT).Value());
	}
}
