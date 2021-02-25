// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.wasmlib.immutable;

import org.iota.wasp.wasmlib.host.*;

public class ScImmutableHashArray {
    int objId;

    public ScImmutableHashArray(int objId) {
        this.objId = objId;
    }

    public ScImmutableHash GetHash(int index) {
        return new ScImmutableHash(objId, index);
    }

    public int Length() {
        return Host.GetLength(objId);
    }
}
