// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.wasmlib.immutable;

import org.iota.wasp.wasmlib.host.*;

public class ScImmutableContractIdArray {
    int objId;

    public ScImmutableContractIdArray(int objId) {
        this.objId = objId;
    }

    public ScImmutableContractId GetContractId(int index) {
        return new ScImmutableContractId(objId, index);
    }

    public int Length() {
        return Host.GetLength(objId);
    }
}
