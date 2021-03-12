// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package wasmlib

import "encoding/binary"

const (
	// all TYPE_* values should exactly match the counterpart OBJTYPE_* values on the host!
	TYPE_ARRAY int32 = 0x20

	TYPE_ADDRESS     int32 = 1
	TYPE_AGENT_ID    int32 = 2
	TYPE_BYTES       int32 = 3
	TYPE_CHAIN_ID    int32 = 4
	TYPE_COLOR       int32 = 5
	TYPE_CONTRACT_ID int32 = 6
	TYPE_HASH        int32 = 7
	TYPE_HNAME       int32 = 8
	TYPE_INT64       int32 = 9
	TYPE_MAP         int32 = 10
	TYPE_REQUEST_ID  int32 = 11
	TYPE_STRING      int32 = 12
)

var typeSizes = [...]int{0, 33, 37, 0, 33, 32, 37, 32, 4, 8, 0, 34, 0}

type ScHost interface {
	Exists(objId int32, keyId int32, typeId int32) bool
	GetBytes(objId int32, keyId int32, typeId int32) []byte
	GetKeyIdFromBytes(bytes []byte) int32
	GetKeyIdFromString(key string) int32
	GetObjectId(objId int32, keyId int32, typeId int32) int32
	SetBytes(objId int32, keyId int32, typeId int32, value []byte)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

var host ScHost

func ConnectHost(h ScHost) ScHost {
	oldHost := host
	host = h
	return oldHost
}

func Clear(objId int32) {
	SetBytes(objId, KeyLength, TYPE_INT64, make([]byte, 8))
}

func Exists(objId int32, keyId Key32, typeId int32) bool {
	return host.Exists(objId, int32(keyId), typeId)
}

func GetBytes(objId int32, keyId Key32, typeId int32) []byte {
	bytes := host.GetBytes(objId, int32(keyId), typeId)
	if len(bytes) == 0 {
		return make([]byte, typeSizes[typeId])
	}
	return bytes
}

func GetKeyIdFromBytes(bytes []byte) Key32 {
	return Key32(host.GetKeyIdFromBytes(bytes))
}

func GetKeyIdFromString(key string) Key32 {
	return Key32(host.GetKeyIdFromString(key))
}

func GetLength(objId int32) int32 {
	bytes := GetBytes(objId, KeyLength, TYPE_INT64)
	return int32(binary.LittleEndian.Uint64(bytes))
}

func GetObjectId(objId int32, keyId Key32, typeId int32) int32 {
	return host.GetObjectId(objId, int32(keyId), typeId)
}

func Log(text string) {
	SetBytes(1, KeyLog, TYPE_STRING, []byte(text))
}

func Panic(text string) {
	SetBytes(1, KeyPanic, TYPE_STRING, []byte(text))
}

func SetBytes(objId int32, keyId Key32, typeId int32, value []byte) {
	host.SetBytes(objId, int32(keyId), typeId, value)
}

func Trace(text string) {
	SetBytes(1, KeyTrace, TYPE_STRING, []byte(text))
}
