// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.wasmlib.mutable;

import org.iota.wasp.wasmlib.host.*;

public class ScMutableInt64 {
    int objId;
    int keyId;

    public ScMutableInt64(int objId, int keyId) {
        this.objId = objId;
        this.keyId = keyId;
    }

    public boolean Exists() {
        return Host.Exists(objId, keyId, ScType.TYPE_INT);
    }

    public void SetValue(long id) {
        byte[] bytes = new byte[8];
        bytes[0] = (byte) id;
        bytes[1] = (byte) (id >> 8);
        bytes[2] = (byte) (id >> 16);
        bytes[3] = (byte) (id >> 24);
        bytes[4] = (byte) (id >> 32);
        bytes[5] = (byte) (id >> 40);
        bytes[6] = (byte) (id >> 48);
        bytes[7] = (byte) (id >> 56);
        Host.SetBytes(objId, keyId, ScType.TYPE_INT, bytes);
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
