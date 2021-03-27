// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.wasmlib.hashtypes;

import org.iota.wasp.wasmlib.host.*;
import org.iota.wasp.wasmlib.keys.*;

import java.util.*;

public class ScAgentId implements MapKey {
    final byte[] id = new byte[37];

    public ScAgentId(ScChainId chainId, ScHname hContract) {
        System.arraycopy(chainId.id, 0, id, 0, chainId.id.length);
        byte[] bytes = hContract.toBytes();
        System.arraycopy(bytes, 0, id, chainId.id.length, bytes.length);
    }

    public ScAgentId(byte[] bytes) {
        if (bytes == null || bytes.length != id.length) {
            Host.Panic("invalid agent id length");
        }
        System.arraycopy(bytes, 0, id, 0, id.length);
    }

    public ScAddress Address() {
        return new ScAddress(Arrays.copyOf(id, 33));
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        ScAgentId other = (ScAgentId) o;
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

    public ScHname Hname() {
        return new ScHname(Arrays.copyOfRange(id, 33, 37));
    }

    public boolean IsAddress() {
        return Address().AsAgentId().equals(this);
    }

    public byte[] toBytes() {
        return id;
    }

    public String toString() {
        return ScHash.base58Encode(id);
    }
}
