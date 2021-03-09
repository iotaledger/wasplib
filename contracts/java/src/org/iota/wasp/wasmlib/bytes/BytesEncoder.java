// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.wasmlib.bytes;

import org.iota.wasp.wasmlib.hashtypes.*;

import java.io.*;
import java.nio.charset.*;

public class BytesEncoder {
    ByteArrayOutputStream data;

    public BytesEncoder() {
        data = new ByteArrayOutputStream();
    }

    public BytesEncoder Address(ScAddress value) {
        return Bytes(value.toBytes());
    }

    public BytesEncoder AgentId(ScAgentId value) {
        return Bytes(value.toBytes());
    }

    public BytesEncoder Bytes(byte[] value) {
        Int64(value.length);
        data.writeBytes(value);
        return this;
    }

    public BytesEncoder ChainId(ScChainId value) {
        return Bytes(value.toBytes());
    }

    public BytesEncoder Color(ScColor value) {
        return Bytes(value.toBytes());
    }

    public BytesEncoder ContractId(ScContractId value) {
        return Bytes(value.toBytes());
    }

    public byte[] Data() {
        return data.toByteArray();
    }

    public BytesEncoder Hash(ScHash value) {
        return Bytes(value.toBytes());
    }

    public BytesEncoder Hname(ScHname value) {
        return Bytes(value.toBytes());
    }

    public BytesEncoder Int64(long val) {
        int value = (int)val;
        for (; ; ) {
            byte b = (byte) value;
            byte s = (byte) (b & 0x40);
            value >>= 7;
            if ((value == 0 && s == 0) || (value == -1 && s != 0)) {
                data.write((byte) (b & 0x7f));
                return this;
            }
            data.write((byte) (b | 0x80));
        }
    }

    public BytesEncoder String(String value) {
        return Bytes(value.getBytes(StandardCharsets.UTF_8));
    }
}
