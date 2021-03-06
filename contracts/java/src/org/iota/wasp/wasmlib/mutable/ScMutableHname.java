// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.wasmlib.mutable;

import org.iota.wasp.wasmlib.hashtypes.*;
import org.iota.wasp.wasmlib.host.*;

public class ScMutableHname {
    int objId;
    int keyId;

    public ScMutableHname(int objId, int keyId) {
        this.objId = objId;
        this.keyId = keyId;
    }

    public boolean Exists() {
        return Host.Exists(objId, keyId, ScType.TYPE_HNAME);
    }

    public void SetValue(ScHname value) {
        Host.SetBytes(objId, keyId, ScType.TYPE_HNAME, value.toBytes());
    }

    @Override
    public String toString() {
        return Value().toString();
    }

    public ScHname Value() {
        return new ScHname(Host.GetBytes(objId, keyId, ScType.TYPE_HNAME));
    }
}
