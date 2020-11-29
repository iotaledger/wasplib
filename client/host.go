// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

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

type Host interface {
	GetBytes(objId int32, keyId int32, value *byte, size int32) int32
	GetIntRef(objId int32, keyId int32, value *int64)
	GetKeyId(key *byte, size int32) int32
	GetObjectId(objId int32, keyId int32, typeId int32) int32
	SetBytes(objId int32, keyId int32, value *byte, size int32)
	SetIntRef(objId int32, keyId int32, value *int64)
}

var host Host

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

func ConnectHost(h Host) {
	host = h
}

func Exists(objId int32, keyId int32) bool {
	return host.GetBytes(objId, keyId, nil, 0) >= 0
}

func GetBytes(objId int32, keyId int32) []byte {
	// first query length of bytes array
	size := host.GetBytes(objId, keyId, nil, 0)
	if size <= 0 {
		return []byte(nil)
	}

	// allocate a byte array in Wasm memory and
	// copy the actual data bytes to Wasm byte array
	bytes := make([]byte, size)
	host.GetBytes(objId, keyId, &bytes[0], size)
	return bytes
}

func GetInt(objId int32, keyId int32) int64 {
	// Go's Wasm implementation is still geared towards Javascript,
	// which does not know int64. So instead of calling host.GetInt()
	// we call host.GetIntRef() with a 32-bit reference to an int64
	value := int64(0)
	host.GetIntRef(objId, keyId, &value)
	return value
}

func GetKey(bytes []byte) int32 {
	size := int32(len(bytes))
	if size == 0 {
		return host.GetKeyId(nil, -1)
	}
	return host.GetKeyId(&bytes[0], -size-1)
}

func GetKeyId(key string) int32 {
	bytes := []byte(key)
	size := int32(len(bytes))
	if size == 0 {
		return host.GetKeyId(nil, 0)
	}
	return host.GetKeyId(&bytes[0], size)
}

func GetObjectId(objId int32, keyId int32, typeId int32) int32 {
	return host.GetObjectId(objId, keyId, typeId)
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
	host.SetBytes(objId, keyId, ptr, int32(len(value)))
}

func SetInt(objId int32, keyId int32, value int64) {
	// Go's Wasm implementation is still geared towards Javascript,
	// which does not know int64. So instead of calling host.SetInt()
	// we call host.SetIntRef() with a 32-bit reference to the int64
	host.SetIntRef(objId, keyId, &value)
}

func SetString(objId int32, keyId int32, value string) {
	// convert string to UTF8-encoded bytes array
	// negative object id indicates to host that this is a string
	SetBytes(-objId, keyId, []byte(value))
}
