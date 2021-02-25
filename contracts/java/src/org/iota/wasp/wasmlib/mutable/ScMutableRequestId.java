// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.wasmlib.mutable;

import org.iota.wasp.wasmlib.hashtypes.*;
import org.iota.wasp.wasmlib.host.*;

public class ScMutableRequestId {
    int objId;
    int keyId;

    public ScMutableRequestId(int objId, int keyId) {
        this.objId = objId;
        this.keyId = keyId;
    }

    public boolean Exists() {
        return Host.Exists(objId, keyId, ScType.TYPE_REQUEST_ID);
    }

    public void SetValue(ScRequestId value) {
        Host.SetBytes(objId, keyId, ScType.TYPE_REQUEST_ID, value.toBytes());
    }

    @Override
    public String toString() {
        return Value().toString();
    }

    public ScRequestId Value() {
        return new ScRequestId(Host.GetBytes(objId, keyId, ScType.TYPE_REQUEST_ID));
    }
}
