// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package wasmlib

import "encoding/binary"

//nolint:revive
const (
	// all TYPE_* values should exactly match the counterpart OBJTYPE_* values on the host!
	TYPE_ARRAY int32 = 0x20
	TYPE_CALL  int32 = 0x40

	TYPE_ADDRESS    int32 = 1
	TYPE_AGENT_ID   int32 = 2
	TYPE_BYTES      int32 = 3
	TYPE_CHAIN_ID   int32 = 4
	TYPE_COLOR      int32 = 5
	TYPE_HASH       int32 = 6
	TYPE_HNAME      int32 = 7
	TYPE_INT16      int32 = 8
	TYPE_INT32      int32 = 9
	TYPE_INT64      int32 = 10
	TYPE_MAP        int32 = 11
	TYPE_REQUEST_ID int32 = 12
	TYPE_STRING     int32 = 13

	OBJ_ID_NULL    int32 = 0
	OBJ_ID_ROOT    int32 = 1
	OBJ_ID_STATE   int32 = 2
	OBJ_ID_PARAMS  int32 = 3
	OBJ_ID_RESULTS int32 = 4
)

var TypeSizes = [...]uint8{0, 33, 37, 0, 33, 32, 32, 4, 2, 4, 8, 0, 34, 0}

type ScHost interface {
	CallFunc(objID, keyID int32, params []byte) []byte
	Exists(objID, keyID, typeID int32) bool
	GetBytes(objID, keyID, typeID int32) []byte
	GetKeyIDFromBytes(bytes []byte) int32
	GetKeyIDFromString(key string) int32
	GetObjectID(objID, keyID, typeID int32) int32
	SetBytes(objID, keyID, typeID int32, value []byte)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\
var host ScHost

func ConnectHost(h ScHost) ScHost {
	oldHost := host
	host = h
	return oldHost
}

func CallFunc(objID int32, keyID Key32, params []byte) []byte {
	return host.CallFunc(objID, int32(keyID), params)
}

func Clear(objID int32) {
	var zero [4]byte
	SetBytes(objID, KeyLength, TYPE_INT32, zero[:])
}

func Exists(objID int32, keyID Key32, typeID int32) bool {
	return host.Exists(objID, int32(keyID), typeID)
}

func GetBytes(objID int32, keyID Key32, typeID int32) []byte {
	bytes := host.GetBytes(objID, int32(keyID), typeID)
	if len(bytes) == 0 {
		return make([]byte, TypeSizes[typeID])
	}
	return bytes
}

func GetKeyIDFromBytes(bytes []byte) Key32 {
	return Key32(host.GetKeyIDFromBytes(bytes))
}

func GetKeyIDFromString(key string) Key32 {
	return Key32(host.GetKeyIDFromString(key))
}

func GetLength(objID int32) int32 {
	bytes := GetBytes(objID, KeyLength, TYPE_INT32)
	return int32(binary.LittleEndian.Uint32(bytes))
}

func GetObjectID(objID int32, keyID Key32, typeID int32) int32 {
	return host.GetObjectID(objID, int32(keyID), typeID)
}

func Log(text string) {
	SetBytes(1, KeyLog, TYPE_STRING, []byte(text))
}

func Panic(text string) {
	SetBytes(1, KeyPanic, TYPE_STRING, []byte(text))
}

func SetBytes(objID int32, keyID Key32, typeID int32, value []byte) {
	host.SetBytes(objID, int32(keyID), typeID, value)
}

func Trace(text string) {
	SetBytes(1, KeyTrace, TYPE_STRING, []byte(text))
}
