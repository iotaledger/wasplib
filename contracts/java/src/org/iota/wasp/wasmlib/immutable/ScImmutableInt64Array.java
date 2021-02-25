// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.wasmlib.immutable;

import org.iota.wasp.wasmlib.host.*;

public class ScImmutableInt64Array {
    int objId;

    public ScImmutableInt64Array(int objId) {
        this.objId = objId;
    }

    public ScImmutableInt64 GetInt64(int index) {
        return new ScImmutableInt64(objId, index);
    }

    public int Length() {
        return Host.GetLength(objId);
    }
}
