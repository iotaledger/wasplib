// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.wasmlib.host;

import org.iota.wasp.wasmlib.keys.*;
import org.iota.wasp.wasmlib.mutable.*;

import java.nio.charset.*;

public class Host {
	public static final ScMutableMap root = new ScMutableMap(1);

	private static final byte[] TYPE_SIZES = {0, 33, 37, 0, 33, 32, 37, 32, 4, 8, 0, 0};

	//TODO figure out how to specify extern hostXxxx functions for each
	// of the functions below to call in Wasm module "waspJava"

	// #[link(wasm_import_module = "wasplib")]
	// #[no_mangle]
	public static int hostGetBytes(int objId, int keyId, int typeId, byte[] value, int size) {
		return 0;
	}

	public static int hostGetKeyId(byte[] key, int size) {
		return 0;
	}

	public static int hostGetObjectId(int objId, int keyId, int typeId) {
		return 0;
	}

	public static void hostSetBytes(int objId, int keyId, int typeId, byte[] value, int size) {
	}

	public static void Clear(int objId) {
		SetBytes(objId, Key.Length.KeyId(), ScType.TYPE_INT, new byte[8]);
	}

	public static boolean Exists(int objId, int keyId, int typeId) {
		// negative length (-1) means only test for existence
		// returned size -1 indicates keyId not found (or error)
		// this removes the need for a separate hostExists function
		return hostGetBytes(objId, keyId, typeId, null, -1) >= 0;
	}

	public static byte[] GetBytes(int objId, int keyId, int typeId) {
		// first query length of bytes array
		int size = hostGetBytes(objId, keyId, typeId, null, 0);
		if (size <= 0) {
			return new byte[TYPE_SIZES[typeId]];
		}

		// allocate a byte array in Wasm memory and
		// copy the actual data bytes to Wasm byte array
		byte[] bytes = new byte[size];
		//noinspection ResultOfMethodCallIgnored
		hostGetBytes(objId, keyId, typeId, bytes, size);
		return bytes;
	}

	public static int GetKeyIdFromBytes(byte[] bytes) {
		int size = bytes.length;
		// negative size indicates this was from bytes
		return hostGetKeyId(bytes, -size - 1);
	}

	public static int GetKeyIdFromString(String key) {
		byte[] bytes = key.getBytes(StandardCharsets.UTF_8);
		// non-negative size indicates this was from string
		return hostGetKeyId(bytes, bytes.length);
	}

	public static int GetLength(int objId) {
		byte[] bytes = GetBytes(objId, Key.Length.KeyId(), ScType.TYPE_INT);
		return (bytes[0] & 0xff) | ((bytes[1] & 0xff) << 8) | ((bytes[2] & 0xff) << 16) | ((bytes[3] & 0xff) << 24);
	}

	public static int GetObjectId(int objId, int keyId, int typeId) {
		return hostGetObjectId(objId, keyId, typeId);
	}

	public static void panic(String text) {
		throw new RuntimeException(text);
	}

	public static void SetBytes(int objId, int keyId, int typeId, byte[] value) {
		hostSetBytes(objId, keyId, typeId, value, value.length);
	}
}
