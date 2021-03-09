// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.wasmlib.hashtypes;

import org.iota.wasp.wasmlib.context.*;
import org.iota.wasp.wasmlib.host.*;
import org.iota.wasp.wasmlib.keys.*;

import java.util.*;

public class ScHash implements MapKey {
    final byte[] id = new byte[32];

    public ScHash(byte[] bytes) {
        if (bytes == null || bytes.length != id.length) {
            throw new RuntimeException("invalid hash id length");
        }
        System.arraycopy(bytes, 0, id, 0, id.length);
    }

    public static String base58Encode(byte[] bytes) {
        return new ScFuncContext().Utility().Base58Encode(bytes);
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        ScHash other = (ScHash) o;
        return Arrays.equals(id, other.id);
    }

    @Override
    public int KeyId() {
        return Host.GetKeyIdFromBytes(id);
    }

    @Override
    public int hashCode() {
        return Arrays.hashCode(id);
    }

    public byte[] toBytes() {
        return id;
    }

    @Override

    public String toString() {
        return base58Encode(id);
    }
}
