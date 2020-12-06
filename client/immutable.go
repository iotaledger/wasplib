// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package client

type ScImmutableAddress struct {
	objId int32
	keyId int32
}

func (o ScImmutableAddress) Exists() bool {
	return Exists(o.objId, o.keyId)
}

func (o ScImmutableAddress) Value() *ScAddress {
	return NewScAddress(GetBytes(o.objId, o.keyId))
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScImmutableAddressArray struct {
	objId int32
}

func (o ScImmutableAddressArray) GetAddress(index int32) ScImmutableAddress {
	return ScImmutableAddress{objId: o.objId, keyId: index}
}

func (o ScImmutableAddressArray) Length() int32 {
	return int32(GetInt(o.objId, KeyLength()))
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScImmutableAgent struct {
	objId int32
	keyId int32
}

func (o ScImmutableAgent) Exists() bool {
	return Exists(o.objId, o.keyId)
}

func (o ScImmutableAgent) Value() *ScAgent {
	return NewScAgent(GetBytes(o.objId, o.keyId))
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScImmutableAgentArray struct {
	objId int32
}

func (o ScImmutableAgentArray) GetAgent(index int32) ScImmutableAgent {
	return ScImmutableAgent{objId: o.objId, keyId: index}
}

func (o ScImmutableAgentArray) Length() int32 {
	return int32(GetInt(o.objId, KeyLength()))
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScImmutableBytes struct {
	objId int32
	keyId int32
}

func (o ScImmutableBytes) Exists() bool {
	return Exists(o.objId, o.keyId)
}

func (o ScImmutableBytes) Value() []byte {
	return GetBytes(o.objId, o.keyId)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScImmutableBytesArray struct {
	objId int32
}

func (o ScImmutableBytesArray) GetBytes(index int32) ScImmutableBytes {
	return ScImmutableBytes{objId: o.objId, keyId: index}
}

func (o ScImmutableBytesArray) Length() int32 {
	return int32(GetInt(o.objId, KeyLength()))
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScImmutableColor struct {
	objId int32
	keyId int32
}

func (o ScImmutableColor) Exists() bool {
	return Exists(o.objId, o.keyId)
}

func (o ScImmutableColor) Value() *ScColor {
	return NewScColor(GetBytes(o.objId, o.keyId))
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScImmutableColorArray struct {
	objId int32
}

func (o ScImmutableColorArray) GetColor(index int32) ScImmutableColor {
	return ScImmutableColor{objId: o.objId, keyId: index}
}

func (o ScImmutableColorArray) Length() int32 {
	return int32(GetInt(o.objId, KeyLength()))
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScImmutableInt struct {
	objId int32
	keyId int32
}

func (o ScImmutableInt) Exists() bool {
	return Exists(o.objId, o.keyId)
}

func (o ScImmutableInt) Value() int64 {
	return GetInt(o.objId, o.keyId)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScImmutableIntArray struct {
	objId int32
}

func (o ScImmutableIntArray) GetInt(index int32) ScImmutableInt {
	return ScImmutableInt{objId: o.objId, keyId: index}
}

func (o ScImmutableIntArray) Length() int32 {
	return int32(GetInt(o.objId, KeyLength()))
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScImmutableMap struct {
	objId int32
}

func (o ScImmutableMap) GetAddress(key KeyId) ScImmutableAddress {
	return ScImmutableAddress{objId: o.objId, keyId: key.GetId()}
}

func (o ScImmutableMap) GetAddressArray(key KeyId) ScImmutableAddressArray {
	arrId := GetObjectId(o.objId, key.GetId(), TYPE_BYTES_ARRAY)
	return ScImmutableAddressArray{objId: arrId}
}

func (o ScImmutableMap) GetAgent(key KeyId) ScImmutableAgent {
	return ScImmutableAgent{objId: o.objId, keyId: key.GetId()}
}

func (o ScImmutableMap) GetAgentArray(key KeyId) ScImmutableAgentArray {
	arrId := GetObjectId(o.objId, key.GetId(), TYPE_BYTES_ARRAY)
	return ScImmutableAgentArray{objId: arrId}
}

func (o ScImmutableMap) GetBytes(key KeyId) ScImmutableBytes {
	return ScImmutableBytes{objId: o.objId, keyId: key.GetId()}
}

func (o ScImmutableMap) GetBytesArray(key KeyId) ScImmutableBytesArray {
	arrId := GetObjectId(o.objId, key.GetId(), TYPE_BYTES_ARRAY)
	return ScImmutableBytesArray{objId: arrId}
}

func (o ScImmutableMap) GetColor(key KeyId) ScImmutableColor {
	return ScImmutableColor{objId: o.objId, keyId: key.GetId()}
}

func (o ScImmutableMap) GetColorArray(key KeyId) ScImmutableColorArray {
	arrId := GetObjectId(o.objId, key.GetId(), TYPE_BYTES_ARRAY)
	return ScImmutableColorArray{objId: arrId}
}

func (o ScImmutableMap) GetInt(key KeyId) ScImmutableInt {
	return ScImmutableInt{objId: o.objId, keyId: key.GetId()}
}

func (o ScImmutableMap) GetIntArray(key KeyId) ScImmutableIntArray {
	arrId := GetObjectId(o.objId, key.GetId(), TYPE_INT_ARRAY)
	return ScImmutableIntArray{objId: arrId}
}

func (o ScImmutableMap) GetMap(key KeyId) ScImmutableMap {
	mapId := GetObjectId(o.objId, key.GetId(), TYPE_MAP)
	return ScImmutableMap{objId: mapId}
}

func (o ScImmutableMap) GetMapArray(key KeyId) ScImmutableMapArray {
	arrId := GetObjectId(o.objId, key.GetId(), TYPE_MAP_ARRAY)
	return ScImmutableMapArray{objId: arrId}
}

func (o ScImmutableMap) GetString(key KeyId) ScImmutableString {
	return ScImmutableString{objId: o.objId, keyId: key.GetId()}
}

func (o ScImmutableMap) GetStringArray(key KeyId) ScImmutableStringArray {
	arrId := GetObjectId(o.objId, key.GetId(), TYPE_STRING_ARRAY)
	return ScImmutableStringArray{objId: arrId}
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScImmutableMapArray struct {
	objId int32
}

func (o ScImmutableMapArray) GetMap(index int32) ScImmutableMap {
	mapId := GetObjectId(o.objId, index, TYPE_MAP)
	return ScImmutableMap{objId: mapId}
}

func (o ScImmutableMapArray) Length() int32 {
	return int32(GetInt(o.objId, KeyLength()))
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScImmutableString struct {
	objId int32
	keyId int32
}

func (o ScImmutableString) Exists() bool {
	return Exists(o.objId, o.keyId)
}

func (o ScImmutableString) Value() string {
	return GetString(o.objId, o.keyId)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScImmutableStringArray struct {
	objId int32
}

func (o ScImmutableStringArray) GetString(index int32) ScImmutableString {
	return ScImmutableString{objId: o.objId, keyId: index}
}

func (o ScImmutableStringArray) Length() int32 {
	return int32(GetInt(o.objId, KeyLength()))
}
