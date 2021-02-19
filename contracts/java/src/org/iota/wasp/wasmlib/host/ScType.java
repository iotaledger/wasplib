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
	public static int TYPE_CONTRACT_ID = 6;
	public static int TYPE_HASH = 7;
	public static int TYPE_HNAME = 8;
	public static int TYPE_INT = 9;
	public static int TYPE_MAP = 10;
	public static int TYPE_STRING = 11;
}
