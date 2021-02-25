// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.wasmlib.immutable;

import org.iota.wasp.wasmlib.host.*;

public class ScImmutableBytesArray {
    int objId;

    public ScImmutableBytesArray(int objId) {
        this.objId = objId;
    }

    public ScImmutableBytes GetBytes(int index) {
        return new ScImmutableBytes(objId, index);
    }

    public int Length() {
        return Host.GetLength(objId);
    }
}
