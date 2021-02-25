// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.wasmlib.mutable;

import org.iota.wasp.wasmlib.hashtypes.*;
import org.iota.wasp.wasmlib.host.*;

public class ScMutableChainId {
    int objId;
    int keyId;

    public ScMutableChainId(int objId, int keyId) {
        this.objId = objId;
        this.keyId = keyId;
    }

    public boolean Exists() {
        return Host.Exists(objId, keyId, ScType.TYPE_CHAIN_ID);
    }

    public void SetValue(ScColor value) {
        Host.SetBytes(objId, keyId, ScType.TYPE_CHAIN_ID, value.toBytes());
    }

    @Override
    public String toString() {
        return Value().toString();
    }

    public ScChainId Value() {
        return new ScChainId(Host.GetBytes(objId, keyId, ScType.TYPE_CHAIN_ID));
    }
}
