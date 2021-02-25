// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.wasmlib.mutable;

import org.iota.wasp.wasmlib.host.*;
import org.iota.wasp.wasmlib.immutable.*;

public class ScMutableMapArray {
    int objId;

    public ScMutableMapArray(int objId) {
        this.objId = objId;
    }

    public void Clear() {
        Host.Clear(objId);
    }

    public ScMutableMap GetMap(int index) {
        int mapId = Host.GetObjectId(objId, index, ScType.TYPE_MAP);
        return new ScMutableMap(mapId);
    }

    public ScImmutableMapArray Immutable() {
        return new ScImmutableMapArray(objId);
    }

    public int Length() {
        return Host.GetLength(objId);
    }
}
