package org.iota.wasp.wasmlib.context;

import org.iota.wasp.wasmlib.hashtypes.*;
import org.iota.wasp.wasmlib.immutable.*;
import org.iota.wasp.wasmlib.mutable.*;

public class ScTransfers {
    final ScMutableMap transfers = new ScMutableMap();

    public static ScTransfers iotas(long amount) {
        return new ScTransfers(ScColor.IOTA, amount);
    }

    public ScTransfers() {
    }

    public ScTransfers(ScColor color, long amount) {
        Set(color, amount);
    }

    public ScTransfers(ScBalances balances) {
        ScImmutableColorArray colors = balances.Colors();
        int length = colors.Length();
        for (int i = 0; i < length; i++) {
            ScColor color = colors.GetColor(i).Value();
            Set(color, balances.Balance(color));
        }
    }

    public void Set(ScColor color, long amount) {
        transfers.GetInt64(color).SetValue(amount);
    }

    public int mapId() {
        return transfers.mapId();
    }
}
