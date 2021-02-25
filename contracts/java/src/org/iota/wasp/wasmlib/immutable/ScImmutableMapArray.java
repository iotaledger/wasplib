// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.wasmlib.immutable;

import org.iota.wasp.wasmlib.host.*;

public class ScImmutableMapArray {
    int objId;

    public ScImmutableMapArray(int objId) {
        this.objId = objId;
    }

    public ScImmutableMap GetMap(int index) {
        int mapId = Host.GetObjectId(objId, index, ScType.TYPE_MAP);
        return new ScImmutableMap(mapId);
    }

    public int Length() {
        return Host.GetLength(objId);
    }
}
