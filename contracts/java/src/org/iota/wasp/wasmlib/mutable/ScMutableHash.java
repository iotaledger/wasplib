// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.wasmlib.mutable;

import org.iota.wasp.wasmlib.hashtypes.*;
import org.iota.wasp.wasmlib.host.*;

public class ScMutableHash {
    int objId;
    int keyId;

    public ScMutableHash(int objId, int keyId) {
        this.objId = objId;
        this.keyId = keyId;
    }

    public boolean Exists() {
        return Host.Exists(objId, keyId, ScType.TYPE_HASH);
    }

    public void SetValue(ScHash value) {
        Host.SetBytes(objId, keyId, ScType.TYPE_HASH, value.toBytes());
    }

    @Override
    public String toString() {
        return Value().toString();
    }

    public ScHash Value() {
        return new ScHash(Host.GetBytes(objId, keyId, ScType.TYPE_HASH));
    }
}
