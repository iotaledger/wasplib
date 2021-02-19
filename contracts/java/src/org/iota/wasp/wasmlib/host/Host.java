// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package org.iota.wasp.wasmlib.host;

import org.iota.wasp.wasmlib.keys.*;
import org.iota.wasp.wasmlib.mutable.*;

import java.nio.charset.*;

public class Host {
	public static final ScMutableMap root = new ScMutableMap(1);

	//TODO figure out how to specify extern hostXxxx functions for each
	// of the functions below to call in Wasm module "waspJava"

	// #[link(wasm_import_module = "wasplib")]
	// #[no_mangle]
	public static int hostGetBytes(int objId, int keyId, byte[] value, int size) {
		return 0;
	}

	public static long hostGetInt(int objId, int keyId) {
		return 0;
	}

	public static int hostGetKeyId(byte[] key, int size) {
		return 0;
	}

	public static int hostGetObjectId(int objId, int keyId, int typeId) {
		return 0;
	}

	public static void hostSetBytes(int objId, int keyId, byte[] value, int size) {
	}

	public static void hostSetInt(int objId, int keyId, long value) {
	}

	public static boolean Exists(int objId, int keyId) {
		// negative length (-1) means only test for existence
		// returned size -1 indicates keyId not found (or error)
		// this removes the need for a separate hostExists function
		return hostGetBytes(objId, keyId, null, -1) >= 0;
	}

	public static byte[] GetBytes(int objId, int keyId) {
		// first query length of bytes array
		int size = hostGetBytes(objId, keyId, null, 0);
		if (size <= 0) {
			return null;
		}

		// allocate a byte array in Wasm memory and
		// copy the actual data bytes to Wasm byte array
		byte[] bytes = new byte[size];
		//noinspection ResultOfMethodCallIgnored
		hostGetBytes(objId, keyId, bytes, size);
		return bytes;
	}

	public static long GetInt(int objId, int keyId) {
		return hostGetInt(objId, keyId);
	}

	public static int GetKeyIdFromBytes(byte[] key) {
		int size = key.length;
		// negative size indicates this was from bytes
		return hostGetKeyId(key, -size - 1);
	}

	public static int GetKeyIdFromString(String key) {
		byte[] bytes = key.getBytes(StandardCharsets.UTF_8);
		int size = bytes.length;
		// non-negative size indicates this was from string
		return hostGetKeyId(bytes, size);
	}

	public static int GetLength(int objId) {
		return (int) GetInt(objId, Key.Length.GetId());
	}

	public static int GetObjectId(int objId, int keyId, int typeId) {
		return hostGetObjectId(objId, keyId, typeId);
	}

	public static String GetString(int objId, int keyId) {
		// convert UTF8-encoded bytes array to string
		// negative object id indicates to host that this is a string
		// this removes the need for a separate hostGetString function
		byte[] bytes = GetBytes(-objId, keyId);
		return bytes == null ? "" : new String(bytes, StandardCharsets.UTF_8);
	}

	public static void panic(String text) {
		throw new RuntimeException(text);
	}

	public static void SetBytes(int objId, int keyId, byte[] value) {
		hostSetBytes(objId, keyId, value, value.length);
	}

	public static void SetClear(int objId) {
		SetInt(objId, Key.Length.GetId(), 0);
	}

	public static void SetInt(int objId, int keyId, long value) {
		hostSetInt(objId, keyId, value);
	}

	public static void SetString(int objId, int keyId, String value) {
		// convert string to UTF8-encoded bytes array
		// negative object id indicates to host that this is a string
		// this removes the need for a separate hostSetString function
		SetBytes(-objId, keyId, value.getBytes(StandardCharsets.UTF_8));
	}
}
