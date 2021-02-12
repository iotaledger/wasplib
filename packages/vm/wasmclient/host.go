// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// +build wasm

package wasmclient

import "github.com/iotaledger/wasp/packages/vm/wasmlib"

//go:wasm-module wasplib
//export hostGetBytes
func hostGetBytes(objId int32, keyId int32, typeId int32, value *byte, size int32) int32

//go:wasm-module wasplib
//export hostGetKeyId
func hostGetKeyId(key *byte, size int32) int32

//go:wasm-module wasplib
//export hostGetObjectId
func hostGetObjectId(objId int32, keyId int32, typeId int32) int32

//go:wasm-module wasplib
//export hostSetBytes
func hostSetBytes(objId int32, keyId int32, typeId int32, value *byte, size int32)

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// implements wasmlib.ScHost interface
type WasmVmHost struct{}

func ConnectWasmHost() {
	wasmlib.ConnectHost(WasmVmHost{})
}

func (w WasmVmHost) Exists(objId int32, keyId int32, typeId int32) bool {
	// negative length (-1) means only test for existence
	// returned size -1 indicates keyId not found (or error)
	// this removes the need for a separate hostExists function
	return hostGetBytes(objId, keyId, typeId, nil, -1) >= 0
}

func (w WasmVmHost) GetBytes(objId int32, keyId int32, typeId int32) []byte {
	// first query expected length of bytes array
	size := hostGetBytes(objId, keyId, typeId, nil, 0)
	if size <= 0 {
		return []byte(nil)
	}

	// allocate a byte array in Wasm memory and
	// copy the actual data bytes to Wasm byte array
	bytes := make([]byte, size)
	hostGetBytes(objId, keyId, typeId, &bytes[0], size)
	return bytes
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

func (w WasmVmHost) SetBytes(objId int32, keyId int32, typeId int32, value []byte) {
	// &bytes[0] will panic on zero length slice, so use nil instead
	size := int32(len(value))
	if size == 0 {
		hostSetBytes(objId, keyId, typeId, nil, size)
		return
	}
	hostSetBytes(objId, keyId, typeId, &value[0], size)
}
