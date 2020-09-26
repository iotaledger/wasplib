package org.iota.wasplib.client;

import java.nio.charset.StandardCharsets;

public class Host {
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

	public static byte[] GetBytes(int objId, int keyId) {
		int size = hostGetBytes(objId, keyId, null, 0);
		if (size == 0) {
			return null;
		}
		byte[] bytes = new byte[size];
		hostGetBytes(objId, keyId, bytes, size);
		return bytes;
	}

	public static long GetInt(int objId, int keyId) {
		return hostGetInt(objId, keyId);
	}

	public static int GetKeyId(String key) {
		byte[] bytes = key.getBytes(StandardCharsets.UTF_8);
		return hostGetKeyId(bytes, bytes.length);
	}

	public static int GetObjectId(int objId, int keyId, int typeId) {
		return hostGetObjectId(objId, keyId, typeId);
	}

	public static String GetString(int objId, int keyId) {
		byte[] bytes = GetBytes(-objId, keyId);
		return bytes == null ? "" : new String(bytes, StandardCharsets.UTF_8);
	}

	public static void panic(String text) {
		throw new RuntimeException(text);
	}

	public static void SetBytes(int objId, int keyId, byte[] value) {
		hostSetBytes(objId, keyId, value, value.length);
	}

	public static void SetInt(int objId, int keyId, long value) {
		hostSetInt(objId, keyId, value);
	}

	public static void SetString(int objId, int keyId, String value) {
		SetBytes(-objId, keyId, value.getBytes(StandardCharsets.UTF_8));
	}
}
