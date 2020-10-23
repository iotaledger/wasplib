use super::context::ScContext;

// all const values should exactly match the counterpart values on the host!
pub const TYPE_BYTES: i32 = 0;
pub const TYPE_BYTES_ARRAY: i32 = 1;
pub const TYPE_INT: i32 = 2;
pub const TYPE_INT_ARRAY: i32 = 3;
pub const TYPE_MAP: i32 = 4;
pub const TYPE_MAP_ARRAY: i32 = 5;
pub const TYPE_STRING: i32 = 6;
pub const TYPE_STRING_ARRAY: i32 = 7;

// all token colors must be encoded as a 64-byte hex string,
// except for the following two special cases:
// default color, encoded as client.IOTA (COLOR_IOTA)
// new color, encoded as "new" (COLOR_NEW)

// any host function that gets called once the current request has
// entered an error state will immediately return without action.
// Any return value will be zero or empty string in that case
#[link(wasm_import_module = "wasplib")]
#[no_mangle]
extern {
    pub fn hostGetBytes(obj_id: i32, key_id: i32, value: *mut u8, len: i32) -> i32;
    pub fn hostGetInt(obj_id: i32, key_id: i32) -> i64;
    pub fn hostGetKeyId(key: *const u8, len: i32) -> i32;
    pub fn hostGetObjectId(obj_id: i32, key_id: i32, type_id: i32) -> i32;
    pub fn hostSetBytes(obj_id: i32, key_id: i32, value: *const u8, len: i32);
    pub fn hostSetInt(obj_id: i32, key_id: i32, value: i64);
}


#[no_mangle]
pub fn nothing() {
    let ctx = ScContext::new();
    ctx.log("Doing nothing as requested. Oh, wait...");
}

pub fn get_bytes(obj_id: i32, key_id: i32) -> Vec<u8> {
    unsafe {
        // first query length of bytes array
        let size = hostGetBytes(obj_id, key_id, std::ptr::null_mut(), 0) as usize;
        if size == 0 { return vec![0_u8; 0]; }

        // allocate a byte array in Wasm memory and
        // copy the actual data bytes to Wasm byte array
        let mut bytes = vec![0_u8; size];
        hostGetBytes(obj_id, key_id, bytes.as_mut_ptr(), size as i32);
        return bytes;
    }
}

pub fn get_int(obj_id: i32, key_id: i32) -> i64 {
    unsafe {
        hostGetInt(obj_id, key_id)
    }
}

pub fn get_key(bytes: &[u8]) -> i32 {
    unsafe {
        hostGetKeyId(bytes.as_ptr(), bytes.len() as i32)
    }
}

pub fn get_key_id(key: &str) -> i32 {
    get_key(key.as_bytes())
}

pub fn get_object_id(obj_id: i32, key_id: i32, type_id: i32) -> i32 {
    unsafe {
        hostGetObjectId(obj_id, key_id, type_id)
    }
}

pub fn get_string(obj_id: i32, key_id: i32) -> String {
    // convert UTF8-encoded bytes array to string
    // negative object id indicates to host that this is a string
    unsafe {
        let bytes = get_bytes(-obj_id, key_id);
        return String::from_utf8_unchecked(bytes);
    }
}

pub fn set_bytes(obj_id: i32, key_id: i32, value: &[u8]) {
    unsafe {
        hostSetBytes(obj_id, key_id, value.as_ptr(), value.len() as i32)
    }
}

pub fn set_int(obj_id: i32, key_id: i32, value: i64) {
    unsafe {
        hostSetInt(obj_id, key_id, value)
    }
}

pub fn set_string(obj_id: i32, key_id: i32, value: &str) {
    set_bytes(-obj_id, key_id, value.as_bytes())
}
