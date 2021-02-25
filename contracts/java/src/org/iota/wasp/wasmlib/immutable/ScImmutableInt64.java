// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.wasmlib.immutable;

import org.iota.wasp.wasmlib.host.*;

public class ScImmutableInt64 {
    int objId;
    int keyId;

    public ScImmutableInt64(int objId, int keyId) {
        this.objId = objId;
        this.keyId = keyId;
    }

    public boolean Exists() {
        return Host.Exists(objId, keyId, ScType.TYPE_INT);
    }

    @Override
    public String toString() {
        return "" + Value();
    }

    public long Value() {
        byte[] bytes = Host.GetBytes(objId, keyId, ScType.TYPE_INT);
        return (bytes[0] & 0xffL) |
                ((bytes[1] & 0xffL) << 8) |
                ((bytes[2] & 0xffL) << 16) |
                ((bytes[3] & 0xffL) << 24) |
                ((bytes[4] & 0xffL) << 32) |
                ((bytes[5] & 0xffL) << 40) |
                ((bytes[6] & 0xffL) << 48) |
                ((bytes[7] & 0xffL) << 56);
    }
}
