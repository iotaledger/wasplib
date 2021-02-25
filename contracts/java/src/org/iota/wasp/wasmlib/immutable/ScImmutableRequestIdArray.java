// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.wasmlib.immutable;

import org.iota.wasp.wasmlib.host.*;

public class ScImmutableRequestIdArray {
    int objId;

    public ScImmutableRequestIdArray(int objId) {
        this.objId = objId;
    }

    public ScImmutableRequestId GetRequestId(int index) {
        return new ScImmutableRequestId(objId, index);
    }

    public int Length() {
        return Host.GetLength(objId);
    }
}
