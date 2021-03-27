// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.wasmlib.bytes;

import org.iota.wasp.wasmlib.hashtypes.*;
import org.iota.wasp.wasmlib.host.*;

import java.nio.charset.*;
import java.util.*;

public class BytesDecoder {
    byte[] data;

    public BytesDecoder(byte[] data) {
        this.data = data;
    }

    public ScAddress Address() {
        return new ScAddress(Bytes());
    }

    public ScAgentId AgentId() {
        return new ScAgentId(Bytes());
    }

    public byte[] Bytes() {
        int size = (int) Int64();
        if (data.length < size) {
            Host.Panic("cannot decode bytes");
        }
        byte[] value = Arrays.copyOfRange(data, 0, size);
        data = Arrays.copyOfRange(data, size, data.length);
        return value;
    }

    public ScChainId ChainId() {
        return new ScChainId(Bytes());
    }

    public ScColor Color() {
        return new ScColor(Bytes());
    }

    public ScHash Hash() {
        return new ScHash(Bytes());
    }

    public ScHname Hname() {
        return new ScHname(Bytes());
    }

    public long Int64() {
        long val = 0;
        int s = 0;
        for (; ; ) {
            byte b = data[0];
            data = Arrays.copyOfRange(data, 1, data.length);
            val |= ((long) (b & 0x7f)) << s;
            if ((b & 0x80) == 0) {
                if (((byte) (val >> s) & 0x7f) != (b & 0x7f)) {
                    Host.Panic("integer too large");
                    return 0;
                }
                // positive value?
                if ((b & 0x40) == 0) {
                    // extend positive sign to int64
                    return val | (((long) b) << s);
                }
                // extend negative sign to int64
                return val | ((0xffffffffffffff80L | b) << s);
            }
            s += 7;
            if (s >= 64) {
                Host.Panic("integer representation too long");
                return 0;
            }
        }
    }

    public String String() {
        return new String(Bytes(), StandardCharsets.UTF_8);
    }
}
