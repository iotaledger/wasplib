// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.wasmlib.mutable;

import org.iota.wasp.wasmlib.hashtypes.*;
import org.iota.wasp.wasmlib.host.*;

public class ScMutableColor {
    int objId;
    int keyId;

    public ScMutableColor(int objId, int keyId) {
        this.objId = objId;
        this.keyId = keyId;
    }

    public boolean Exists() {
        return Host.Exists(objId, keyId, ScType.TYPE_COLOR);
    }

    public void SetValue(ScColor value) {
        Host.SetBytes(objId, keyId, ScType.TYPE_COLOR, value.toBytes());
    }

    @Override
    public String toString() {
        return Value().toString();
    }

    public ScColor Value() {
        return new ScColor(Host.GetBytes(objId, keyId, ScType.TYPE_COLOR));
    }
}
