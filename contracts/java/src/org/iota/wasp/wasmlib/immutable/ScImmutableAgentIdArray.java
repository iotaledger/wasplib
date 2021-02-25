// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.wasmlib.immutable;

import org.iota.wasp.wasmlib.host.*;

public class ScImmutableAgentIdArray {
    int objId;

    public ScImmutableAgentIdArray(int objId) {
        this.objId = objId;
    }

    public ScImmutableAgentId GetAgentId(int index) {
        return new ScImmutableAgentId(objId, index);
    }

    public int Length() {
        return Host.GetLength(objId);
    }
}
