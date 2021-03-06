// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.wasmlib.host;

public class ScType {
    // all TYPE_* values should exactly match the counterpart OBJTYPE_* values on the host!
    public static int TYPE_ARRAY = 0x20;

    public static int TYPE_ADDRESS = 1;
    public static int TYPE_AGENT_ID = 2;
    public static int TYPE_BYTES = 3;
    public static int TYPE_CHAIN_ID = 4;
    public static int TYPE_COLOR = 5;
    public static int TYPE_HASH = 6;
    public static int TYPE_HNAME = 7;
    public static int TYPE_INT64 = 8;
    public static int TYPE_MAP = 9;
    public static int TYPE_REQUEST_ID = 10;
    public static int TYPE_STRING = 11;
}
