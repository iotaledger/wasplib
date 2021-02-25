// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.wasmlib.mutable;

import org.iota.wasp.wasmlib.host.*;
import org.iota.wasp.wasmlib.immutable.*;

public class ScMutableChainIdArray {
    int objId;

    public ScMutableChainIdArray(int objId) {
        this.objId = objId;
    }

    public void Clear() {
        Host.Clear(objId);
    }

    public ScMutableChainId GetChainId(int index) {
        return new ScMutableChainId(objId, index);
    }

    public ScImmutableChainIdArray Immutable() {
        return new ScImmutableChainIdArray(objId);
    }

    public int Length() {
        return Host.GetLength(objId);
    }
}
