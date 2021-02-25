// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.wasmlib.mutable;

import org.iota.wasp.wasmlib.host.*;

import java.nio.charset.*;

public class ScMutableString {
    int objId;
    int keyId;

    public ScMutableString(int objId, int keyId) {
        this.objId = objId;
        this.keyId = keyId;
    }

    public boolean Exists() {
        return Host.Exists(objId, keyId, ScType.TYPE_STRING);
    }

    public void SetValue(String value) {
        // convert string to UTF8-encoded bytes array
        byte[] bytes = value.getBytes(StandardCharsets.UTF_8);
        Host.SetBytes(objId, keyId, ScType.TYPE_STRING, bytes);
    }

    @Override
    public String toString() {
        return Value();
    }

    public String Value() {
        // convert UTF8-encoded bytes array to string
        byte[] bytes = Host.GetBytes(objId, keyId, ScType.TYPE_STRING);
        return bytes == null ? "" : new String(bytes, StandardCharsets.UTF_8);
    }
}
