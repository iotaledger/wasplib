// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

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

type WasmHost struct{}

func ConnectWasmHost() {
	client.ConnectHost(WasmHost{})
}

func (w WasmHost) GetBytes(objId int32, keyId int32, value *byte, size int32) int32 {
	return hostGetBytes(objId, keyId, value, size)
}

func (w WasmHost) GetIntRef(objId int32, keyId int32, value *int64) {
	hostGetIntRef(objId, keyId, value)
}

func (w WasmHost) GetKeyId(key *byte, size int32) int32 {
	return hostGetKeyId(key, size)
}

func (w WasmHost) GetObjectId(objId int32, keyId int32, typeId int32) int32 {
	return hostGetObjectId(objId, keyId, typeId)
}

func (w WasmHost) SetBytes(objId int32, keyId int32, value *byte, size int32) {
	hostSetBytes(objId, keyId, value, size)
}

func (w WasmHost) SetIntRef(objId int32, keyId int32, value *int64) {
	hostSetIntRef(objId, keyId, value)
}
