// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.wasmlib.immutable;

import org.iota.wasp.wasmlib.hashtypes.*;
import org.iota.wasp.wasmlib.host.*;

public class ScImmutableColor {
    int objId;
    int keyId;

    public ScImmutableColor(int objId, int keyId) {
        this.objId = objId;
        this.keyId = keyId;
    }

    public boolean Exists() {
        return Host.Exists(objId, keyId, ScType.TYPE_COLOR);
    }

    @Override
    public String toString() {
        return Value().toString();
    }

    public ScColor Value() {
        return new ScColor(Host.GetBytes(objId, keyId, ScType.TYPE_COLOR));
    }
}
