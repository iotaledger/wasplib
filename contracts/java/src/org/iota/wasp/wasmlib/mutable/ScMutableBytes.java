// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.wasmlib.mutable;

import org.iota.wasp.wasmlib.hashtypes.*;
import org.iota.wasp.wasmlib.host.*;

public class ScMutableBytes {
    int objId;
    int keyId;

    public ScMutableBytes(int objId, int keyId) {
        this.objId = objId;
        this.keyId = keyId;
    }

    public boolean Exists() {
        return Host.Exists(objId, keyId, ScType.TYPE_BYTES);
    }

    public void SetValue(byte[] value) {
        Host.SetBytes(objId, keyId, ScType.TYPE_BYTES, value);
    }

    @Override
    public String toString() {
        return ScHash.base58Encode(Value());
    }

    public byte[] Value() {
        return Host.GetBytes(objId, keyId, ScType.TYPE_BYTES);
    }
}
