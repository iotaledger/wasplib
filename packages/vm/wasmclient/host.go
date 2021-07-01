// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// +build wasm

package wasmclient

import "github.com/iotaledger/wasplib/packages/vm/wasmlib"

//go:wasm-module WasmLib
//export hostGetBytes
func hostGetBytes(objID, keyID, typeID int32, value *byte, size int32) int32

//go:wasm-module WasmLib
//export hostGetKeyID
func hostGetKeyID(key *byte, size int32) int32

//go:wasm-module WasmLib
//export hostGetObjectID
func hostGetObjectID(objID, keyID, typeID int32) int32

//go:wasm-module WasmLib
//export hostSetBytes
func hostSetBytes(objID, keyID, typeID int32, value *byte, size int32)

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type WasmVMHost struct{}

// implements wasmlib.ScHost interface
var _ wasmlib.ScHost = &WasmVMHost{}

func ConnectWasmHost() {
	wasmlib.ConnectHost(WasmVMHost{})
}

func (w WasmVMHost) CallFunc(objID, keyID int32, params []byte) []byte {
	args := (*byte)(nil)
	size := int32(len(params))
	if size != 0 {
		args = &params[0]
	}

	// pass params and query expected length of result
	size = hostGetBytes(objID, keyID, wasmlib.TYPE_CALL, args, size)

	// -1 means non-existent, so return default value for type
	if size <= 0 {
		return []byte(nil)
	}

	// allocate a sufficient length byte array in Wasm memory
	// and let the host copy the actual result into this Wasm byte array
	result := make([]byte, size)
	hostGetBytes(objID, keyID, wasmlib.TYPE_CALL+1, &result[0], size)
	return result
}

func (w WasmVMHost) Exists(objID, keyID, typeID int32) bool {
	// negative length (-1) means only test for existence
	// returned size -1 indicates keyID not found (or error)
	// this removes the need for a separate hostExists function
	return hostGetBytes(objID, keyID, typeID, nil, -1) >= 0
}

func (w WasmVMHost) GetBytes(objID, keyID, typeID int32) []byte {
	size := int32(wasmlib.TypeSizes[typeID])
	if size == 0 {
		// variable-sized type, first query expected length of bytes array
		// (pass zero-length buffer)
		size = hostGetBytes(objID, keyID, typeID, nil, 0)

		// -1 means non-existent, so return default value for type
		if size <= 0 {
			return []byte(nil)
		}
	}

	// allocate a sufficient length byte array in Wasm memory
	// and let the host copy the actual data bytes into this Wasm byte array
	result := make([]byte, size)
	hostGetBytes(objID, keyID, typeID, &result[0], size)
	return result
}

func (w WasmVMHost) GetKeyIDFromBytes(bytes []byte) int32 {
	size := int32(len(bytes))
	// &bytes[0] will panic on zero length slice, so use nil instead
	// negative size indicates this was from bytes
	if size == 0 {
		return hostGetKeyID(nil, -1)
	}
	return hostGetKeyID(&bytes[0], -size-1)
}

func (w WasmVMHost) GetKeyIDFromString(key string) int32 {
	bytes := []byte(key)
	size := int32(len(bytes))
	// &bytes[0] will panic on zero length slice, so use nil instead
	// non-negative size indicates this was from string
	if size == 0 {
		return hostGetKeyID(nil, 0)
	}
	return hostGetKeyID(&bytes[0], size)
}

func (w WasmVMHost) GetObjectID(objID, keyID, typeID int32) int32 {
	return hostGetObjectID(objID, keyID, typeID)
}

func (w WasmVMHost) SetBytes(objID, keyID, typeID int32, value []byte) {
	// &bytes[0] will panic on zero length slice, so use nil instead
	size := int32(len(value))
	if size == 0 {
		hostSetBytes(objID, keyID, typeID, nil, size)
		return
	}
	hostSetBytes(objID, keyID, typeID, &value[0], size)
}
