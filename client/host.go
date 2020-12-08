// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package client

const (
	TYPE_BYTES        int32 = 0
	TYPE_BYTES_ARRAY  int32 = 1
	TYPE_INT          int32 = 2
	TYPE_INT_ARRAY    int32 = 3
	TYPE_MAP          int32 = 4
	TYPE_MAP_ARRAY    int32 = 5
	TYPE_STRING       int32 = 6
	TYPE_STRING_ARRAY int32 = 7
)

type ScHost interface {
	Exists(objId int32, keyId int32) bool
	GetBytes(objId int32, keyId int32) []byte
	GetInt(objId int32, keyId int32) int64
	GetKey(bytes []byte) int32
	GetKeyId(key string) int32
	GetObjectId(objId int32, keyId int32, typeId int32) int32
	GetString(objId int32, keyId int32) string
	SetBytes(objId int32, keyId int32, value []byte)
	SetInt(objId int32, keyId int32, value int64)
	SetString(objId int32, keyId int32, value string)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

var host ScHost

func ConnectHost(h ScHost) {
	host = h
}

func Exists(objId int32, keyId int32) bool {
	return host.Exists(objId, keyId)
}

func GetBytes(objId int32, keyId int32) []byte {
	return host.GetBytes(objId, keyId)
}

func GetInt(objId int32, keyId int32) int64 {
	return host.GetInt(objId, keyId)
}

func GetKey(bytes []byte) int32 {
	return host.GetKey(bytes)
}

func GetKeyId(key string) int32 {
	return host.GetKeyId(key)
}

func GetLength(objId int32) int32 {
	return int32(GetInt(objId, int32(KeyLength)))
}

func GetObjectId(objId int32, keyId int32, typeId int32) int32 {
	return host.GetObjectId(objId, keyId, typeId)
}

func GetString(objId int32, keyId int32) string {
	return host.GetString(objId, keyId)
}

func SetBytes(objId int32, keyId int32, value []byte) {
	host.SetBytes(objId, keyId, value)
}

func SetClear(objId int32) {
	SetInt(objId, int32(KeyLength), 0)
}

func SetInt(objId int32, keyId int32, value int64) {
	host.SetInt(objId, keyId, value)
}

func SetString(objId int32, keyId int32, value string) {
	host.SetString(objId, keyId, value)
}
