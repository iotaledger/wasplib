package client

const (
	OBJTYPE_BYTES        int32 = 0
	OBJTYPE_BYTES_ARRAY  int32 = 1
	OBJTYPE_INT          int32 = 2
	OBJTYPE_INT_ARRAY    int32 = 3
	OBJTYPE_MAP          int32 = 4
	OBJTYPE_MAP_ARRAY    int32 = 5
	OBJTYPE_STRING       int32 = 6
	OBJTYPE_STRING_ARRAY int32 = 7
)

//go:wasm-module wasplib
//export hostGetBytes
func hostGetBytes(objId int32, keyId int32, value *byte, size int32) int32

//go:wasm-module wasplib
//export hostGetIntRef
func hostGetIntRef(objId int32, keyId int32, value *int64)

//go:wasm-module wasplib
//export hostGetKeyId
func hostGetKeyId(key *byte, size int32) int32

//go:wasm-module wasplib
//export hostGetObjectId
func hostGetObjectId(objId int32, keyId int32, typeId int32) int32

//go:wasm-module wasplib
//export hostSetBytes
func hostSetBytes(objId int32, keyId int32, value *byte, size int32)

//go:wasm-module wasplib
//export hostSetIntRef
func hostSetIntRef(objId int32, keyId int32, value *int64)

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

//export nothing
func nothing() {
	ctx := NewScContext()
	ctx.Log("Doing nothing as requested. Oh, wait...")
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

func Exists(objId int32, keyId int32) bool {
	return hostGetBytes(objId, keyId, nil, 0) >= 0
}

func GetBytes(objId int32, keyId int32) []byte {
	// first query length of bytes array
	size := hostGetBytes(objId, keyId, nil, 0)
	if size <= 0 {
		return []byte(nil)
	}

	// allocate a byte array in Wasm memory and
	// copy the actual data bytes to Wasm byte array
	bytes := make([]byte, size)
	hostGetBytes(objId, keyId, &bytes[0], size)
	return bytes
}

func GetInt(objId int32, keyId int32) int64 {
	// Go's Wasm implementation is still geared towards Javascript,
	// which does not know int64. So instead of calling hostGetInt()
	// we call hostGetIntRef() with a 32-bit reference to an int64
	value := int64(0)
	hostGetIntRef(objId, keyId, &value)
	return value
}

func GetKey(bytes []byte) int32 {
	size := int32(len(bytes))
	if size == 0 {
		return hostGetKeyId(nil, -1)
	}
	return hostGetKeyId(&bytes[0], -size-1)
}

func GetKeyId(key string) int32 {
	bytes := []byte(key)
	size := int32(len(bytes))
	if size == 0 {
		return hostGetKeyId(nil, 0)
	}
	return hostGetKeyId(&bytes[0], size)
}

func GetObjectId(objId int32, keyId int32, typeId int32) int32 {
	return hostGetObjectId(objId, keyId, typeId)
}

func GetString(objId int32, keyId int32) string {
	// convert UTF8-encoded bytes array to string
	// negative object id indicates to host that this is a string
	bytes := GetBytes(-objId, keyId)
	if bytes == nil {
		return ""
	}
	return string(bytes)
}

func SetBytes(objId int32, keyId int32, value []byte) {
	var ptr *byte = nil
	if len(value) != 0 {
		ptr = &value[0]
	}
	hostSetBytes(objId, keyId, ptr, int32(len(value)))
}

func SetInt(objId int32, keyId int32, value int64) {
	// Go's Wasm implementation is still geared towards Javascript,
	// which does not know int64. So instead of calling hostSetInt()
	// we call hostSetIntRef() with a 32-bit reference to the int64
	hostSetIntRef(objId, keyId, &value)
}

func SetString(objId int32, keyId int32, value string) {
	// convert string to UTF8-encoded bytes array
	// negative object id indicates to host that this is a string
	SetBytes(-objId, keyId, []byte(value))
}
