// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.wasmlib.mutable;

import org.iota.wasp.wasmlib.host.*;
import org.iota.wasp.wasmlib.immutable.*;

public class ScMutableInt64Array {
    int objId;

    public ScMutableInt64Array(int objId) {
        this.objId = objId;
    }

    public void Clear() {
        Host.Clear(objId);
    }

    public ScMutableInt64 GetInt64(int index) {
        return new ScMutableInt64(objId, index);
    }

    public ScImmutableInt64Array Immutable() {
        return new ScImmutableInt64Array(objId);
    }

    public int Length() {
        return Host.GetLength(objId);
    }
}
