package org.iota.wasplib.client;

public class Host {
	//TODO figure out how to specify extern hostXxxx functions for each
	// of the functions below to call in Wasm module "waspJava"

	// #[link(wasm_import_module = "waspJava")]
	// #[no_mangle]
	// extern {
	//     pub fn hostGetInt(obj_id: i32, key_id: i32) -> i64;
	//     pub fn hostGetKeyId(key: &str) -> i32;
	//     pub fn hostGetObjectId(obj_id: i32, key_id: i32, type_id: i32) -> i32;
	//     pub fn hostGetString(obj_id: i32, key_id: i32) -> &'static str;
	//     pub fn hostSetInt(obj_id: i32, key_id: i32, value: i64);
	//     pub fn hostSetString(obj_id: i32, key_id: i32, value: &str);
	// }

	public static long GetInt(int objId, int keyId) {
		return 0;
	}

	public static int GetKeyId(String key) {
		return 0;
	}

	public static int GetObjectId(int objId, int keyId, int typeId) {
		return 0;
	}

	public static String GetString(int objId, int keyId) {
		return "";
	}

	public static void SetInt(int objId, int keyId, long value) {
	}

	public static void SetString(int objId, int keyId, String value) {
	}
}
