// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package client

import "strconv"

type ScImmutableAddress struct {
	objId int32
	keyId Key32
}

func (o ScImmutableAddress) Exists() bool {
	return Exists(o.objId, o.keyId)
}

func (o ScImmutableAddress) String() string {
	return o.Value().String()
}

func (o ScImmutableAddress) Value() *ScAddress {
	return NewScAddress(GetBytes(o.objId, o.keyId))
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScImmutableAddressArray struct {
	objId int32
}

func (o ScImmutableAddressArray) GetAddress(index int32) ScImmutableAddress {
	return ScImmutableAddress{objId: o.objId, keyId: Key32(index)}
}

func (o ScImmutableAddressArray) Length() int32 {
	return GetLength(o.objId)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScImmutableAgent struct {
	objId int32
	keyId Key32
}

func (o ScImmutableAgent) Exists() bool {
	return Exists(o.objId, o.keyId)
}

func (o ScImmutableAgent) String() string {
	return o.Value().String()
}

func (o ScImmutableAgent) Value() *ScAgent {
	return NewScAgent(GetBytes(o.objId, o.keyId))
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScImmutableAgentArray struct {
	objId int32
}

func (o ScImmutableAgentArray) GetAgent(index int32) ScImmutableAgent {
	return ScImmutableAgent{objId: o.objId, keyId: Key32(index)}
}

func (o ScImmutableAgentArray) Length() int32 {
	return GetLength(o.objId)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScImmutableBytes struct {
	objId int32
	keyId Key32
}

func (o ScImmutableBytes) Exists() bool {
	return Exists(o.objId, o.keyId)
}

func (o ScImmutableBytes) String() string {
	return base58Encode(o.Value())
}

func (o ScImmutableBytes) Value() []byte {
	return GetBytes(o.objId, o.keyId)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScImmutableBytesArray struct {
	objId int32
}

func (o ScImmutableBytesArray) GetBytes(index int32) ScImmutableBytes {
	return ScImmutableBytes{objId: o.objId, keyId: Key32(index)}
}

func (o ScImmutableBytesArray) Length() int32 {
	return GetLength(o.objId)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScImmutableColor struct {
	objId int32
	keyId Key32
}

func (o ScImmutableColor) Exists() bool {
	return Exists(o.objId, o.keyId)
}

func (o ScImmutableColor) String() string {
	return o.Value().String()
}

func (o ScImmutableColor) Value() *ScColor {
	return NewScColor(GetBytes(o.objId, o.keyId))
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScImmutableColorArray struct {
	objId int32
}

func (o ScImmutableColorArray) GetColor(index int32) ScImmutableColor {
	return ScImmutableColor{objId: o.objId, keyId: Key32(index)}
}

func (o ScImmutableColorArray) Length() int32 {
	return GetLength(o.objId)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScImmutableHash struct {
	objId int32
	keyId Key32
}

func (o ScImmutableHash) Exists() bool {
	return Exists(o.objId, o.keyId)
}

func (o ScImmutableHash) String() string {
	return o.Value().String()
}

func (o ScImmutableHash) Value() *ScHash {
	return NewScHash(GetBytes(o.objId, o.keyId))
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScImmutableHashArray struct {
	objId int32
}

func (o ScImmutableHashArray) GetHash(index int32) ScImmutableHash {
	return ScImmutableHash{objId: o.objId, keyId: Key32(index)}
}

func (o ScImmutableHashArray) Length() int32 {
	return GetLength(o.objId)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScImmutableHname struct {
	objId int32
	keyId Key32
}

func (o ScImmutableHname) Exists() bool {
	return Exists(o.objId, o.keyId)
}

func (o ScImmutableHname) String() string {
	return o.Value().String()
}

func (o ScImmutableHname) Value() Hname {
	return Hname(GetInt(o.objId, o.keyId))
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScImmutableInt struct {
	objId int32
	keyId Key32
}

func (o ScImmutableInt) Exists() bool {
	return Exists(o.objId, o.keyId)
}

func (o ScImmutableInt) String() string {
	return strconv.FormatInt(o.Value(), 10)
}

func (o ScImmutableInt) Value() int64 {
	return GetInt(o.objId, o.keyId)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScImmutableIntArray struct {
	objId int32
}

func (o ScImmutableIntArray) GetInt(index int32) ScImmutableInt {
	return ScImmutableInt{objId: o.objId, keyId: Key32(index)}
}

func (o ScImmutableIntArray) Length() int32 {
	return GetLength(o.objId)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScImmutableMap struct {
	objId int32
}

func (o ScImmutableMap) GetAddress(key MapKey) ScImmutableAddress {
	return ScImmutableAddress{objId: o.objId, keyId: key.KeyId()}
}

func (o ScImmutableMap) GetAddressArray(key MapKey) ScImmutableAddressArray {
	arrId := GetObjectId(o.objId, key.KeyId(), TYPE_ADDRESS|TYPE_ARRAY)
	return ScImmutableAddressArray{objId: arrId}
}

func (o ScImmutableMap) GetAgent(key MapKey) ScImmutableAgent {
	return ScImmutableAgent{objId: o.objId, keyId: key.KeyId()}
}

func (o ScImmutableMap) GetAgentArray(key MapKey) ScImmutableAgentArray {
	arrId := GetObjectId(o.objId, key.KeyId(), TYPE_AGENT|TYPE_ARRAY)
	return ScImmutableAgentArray{objId: arrId}
}

func (o ScImmutableMap) GetBytes(key MapKey) ScImmutableBytes {
	return ScImmutableBytes{objId: o.objId, keyId: key.KeyId()}
}

func (o ScImmutableMap) GetBytesArray(key MapKey) ScImmutableBytesArray {
	arrId := GetObjectId(o.objId, key.KeyId(), TYPE_BYTES|TYPE_ARRAY)
	return ScImmutableBytesArray{objId: arrId}
}

func (o ScImmutableMap) GetColor(key MapKey) ScImmutableColor {
	return ScImmutableColor{objId: o.objId, keyId: key.KeyId()}
}

func (o ScImmutableMap) GetColorArray(key MapKey) ScImmutableColorArray {
	arrId := GetObjectId(o.objId, key.KeyId(), TYPE_COLOR|TYPE_ARRAY)
	return ScImmutableColorArray{objId: arrId}
}

func (o ScImmutableMap) GetHash(key MapKey) ScImmutableHash {
	return ScImmutableHash{objId: o.objId, keyId: key.KeyId()}
}

func (o ScImmutableMap) GetHashArray(key MapKey) ScImmutableHashArray {
	arrId := GetObjectId(o.objId, key.KeyId(), TYPE_HASH|TYPE_ARRAY)
	return ScImmutableHashArray{objId: arrId}
}

func (o ScImmutableMap) GetHname(key MapKey) ScImmutableHname {
	return ScImmutableHname{objId: o.objId, keyId: key.KeyId()}
}

func (o ScImmutableMap) GetInt(key MapKey) ScImmutableInt {
	return ScImmutableInt{objId: o.objId, keyId: key.KeyId()}
}

func (o ScImmutableMap) GetIntArray(key MapKey) ScImmutableIntArray {
	arrId := GetObjectId(o.objId, key.KeyId(), TYPE_INT|TYPE_ARRAY)
	return ScImmutableIntArray{objId: arrId}
}

func (o ScImmutableMap) GetMap(key MapKey) ScImmutableMap {
	mapId := GetObjectId(o.objId, key.KeyId(), TYPE_MAP)
	return ScImmutableMap{objId: mapId}
}

func (o ScImmutableMap) GetMapArray(key MapKey) ScImmutableMapArray {
	arrId := GetObjectId(o.objId, key.KeyId(), TYPE_MAP|TYPE_ARRAY)
	return ScImmutableMapArray{objId: arrId}
}

func (o ScImmutableMap) GetString(key MapKey) ScImmutableString {
	return ScImmutableString{objId: o.objId, keyId: key.KeyId()}
}

func (o ScImmutableMap) GetStringArray(key MapKey) ScImmutableStringArray {
	arrId := GetObjectId(o.objId, key.KeyId(), TYPE_STRING|TYPE_ARRAY)
	return ScImmutableStringArray{objId: arrId}
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScImmutableMapArray struct {
	objId int32
}

func (o ScImmutableMapArray) GetMap(index int32) ScImmutableMap {
	mapId := GetObjectId(o.objId, Key32(index), TYPE_MAP)
	return ScImmutableMap{objId: mapId}
}

func (o ScImmutableMapArray) Length() int32 {
	return GetLength(o.objId)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScImmutableString struct {
	objId int32
	keyId Key32
}

func (o ScImmutableString) Exists() bool {
	return Exists(o.objId, o.keyId)
}

func (o ScImmutableString) String() string {
	return o.Value()
}

func (o ScImmutableString) Value() string {
	return GetString(o.objId, o.keyId)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScImmutableStringArray struct {
	objId int32
}

func (o ScImmutableStringArray) GetString(index int32) ScImmutableString {
	return ScImmutableString{objId: o.objId, keyId: Key32(index)}
}

func (o ScImmutableStringArray) Length() int32 {
	return GetLength(o.objId)
}
