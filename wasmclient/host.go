// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// +build wasm

package wasmclient

import "github.com/iotaledger/wasplib/client"

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

// implements client.ScHost interface
type WasmVmHost struct{}

func ConnectWasmHost() {
	client.ConnectHost(WasmVmHost{})
}

func (w WasmVmHost) Exists(objId int32, keyId int32) bool {
	// negative length (-1) means only test for existence
	// returned size -1 indicates keyId not found (or error)
	// this removes the need for a separate hostExists function
	return hostGetBytes(objId, keyId, nil, -1) >= 0
}

func (w WasmVmHost) GetBytes(objId int32, keyId int32) []byte {
	// first query expected length of bytes array
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

func (w WasmVmHost) GetInt(objId int32, keyId int32) int64 {
	// Go's Wasm implementation is still geared towards Javascript,
	// which does not know int64. So instead of calling hostGetInt()
	// we call hostGetIntRef() with a 32-bit reference to an int64
	value := int64(0)
	hostGetIntRef(objId, keyId, &value)
	return value
}

func (w WasmVmHost) GetKeyIdFromBytes(bytes []byte) int32 {
	size := int32(len(bytes))
	// &bytes[0] will panic on zero length slice, so use nil instead
	// negative size indicates this was from bytes
	if size == 0 {
		return hostGetKeyId(nil, -1)
	}
	return hostGetKeyId(&bytes[0], -size-1)
}

func (w WasmVmHost) GetKeyIdFromString(key string) int32 {
	bytes := []byte(key)
	size := int32(len(bytes))
	// &bytes[0] will panic on zero length slice, so use nil instead
	// non-negative size indicates this was from string
	if size == 0 {
		return hostGetKeyId(nil, 0)
	}
	return hostGetKeyId(&bytes[0], size)
}

func (w WasmVmHost) GetObjectId(objId int32, keyId int32, typeId int32) int32 {
	return hostGetObjectId(objId, keyId, typeId)
}

func (w WasmVmHost) GetString(objId int32, keyId int32) string {
	// convert UTF8-encoded bytes array to string
	// negative object id indicates to host that this is a string
	// this removes the need for a separate hostGetString function
	bytes := w.GetBytes(-objId, keyId)
	if bytes == nil {
		return ""
	}
	return string(bytes)
}

func (w WasmVmHost) SetBytes(objId int32, keyId int32, value []byte) {
	// &bytes[0] will panic on zero length slice, so use nil instead
	size := int32(len(value))
	if size == 0 {
		hostSetBytes(objId, keyId, nil, size)
		return
	}
	hostSetBytes(objId, keyId, &value[0], size)
}

func (w WasmVmHost) SetInt(objId int32, keyId int32, value int64) {
	// Go's Wasm implementation is still geared towards Javascript,
	// which does not know int64. So instead of calling hostSetInt()
	// we call hostSetIntRef() with a 32-bit reference to the int64
	hostSetIntRef(objId, keyId, &value)
}

func (w WasmVmHost) SetString(objId int32, keyId int32, value string) {
	// convert string to UTF8-encoded bytes array
	// negative object id indicates to host that this is a string
	// this removes the need for a separate hostSetString function
	w.SetBytes(-objId, keyId, []byte(value))
}
