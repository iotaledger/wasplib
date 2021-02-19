package org.iota.wasp.wasmlib.context;

import org.iota.wasp.wasmlib.hashtypes.*;
import org.iota.wasp.wasmlib.immutable.*;
import org.iota.wasp.wasmlib.mutable.*;

public class ScTransfers {
	final ScMutableMap transfers = new ScMutableMap();

	public ScTransfers() {
	}

	public ScTransfers(ScColor color, long amount) {
		Add(color, amount);
	}

	public ScTransfers(ScBalances balances) {
		ScImmutableColorArray colors = balances.Colors();
		int length = colors.Length();
		for (int i = 0; i < length; i++) {
			ScColor color = colors.GetColor(i).Value();
			Add(color, balances.Balance(color));
		}
	}

	public void Add(ScColor color, long amount) {
		transfers.GetInt(color).SetValue(amount);
	}

	public int mapId() {
		return transfers.mapId();
	}
}
