// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package client

var root = ScMutableMap{objId: 1}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScMutableAddress struct {
	objId int32
	keyId int32
}

func (o ScMutableAddress) Exists() bool {
	return Exists(o.objId, o.keyId)
}

func (o ScMutableAddress) SetValue(value *ScAddress) {
	SetBytes(o.objId, o.keyId, value.Bytes())
}

func (o ScMutableAddress) Value() *ScAddress {
	return NewScAddress(GetBytes(o.objId, o.keyId))
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScMutableAddressArray struct {
	objId int32
}

func (o ScMutableAddressArray) Clear() {
	SetInt(o.objId, KeyLength(), 0)
}

func (o ScMutableAddressArray) GetAddress(index int32) ScMutableAddress {
	return ScMutableAddress{objId: o.objId, keyId: index}
}

func (o ScMutableAddressArray) Immutable() ScImmutableAddressArray {
	return ScImmutableAddressArray{objId: o.objId}
}

func (o ScMutableAddressArray) Length() int32 {
	return int32(GetInt(o.objId, KeyLength()))
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScMutableAgent struct {
	objId int32
	keyId int32
}

func (o ScMutableAgent) Exists() bool {
	return Exists(o.objId, o.keyId)
}

func (o ScMutableAgent) SetValue(value *ScAgent) {
	SetBytes(o.objId, o.keyId, value.Bytes())
}

func (o ScMutableAgent) Value() *ScAgent {
	return NewScAgent(GetBytes(o.objId, o.keyId))
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScMutableAgentArray struct {
	objId int32
}

func (o ScMutableAgentArray) Clear() {
	SetInt(o.objId, KeyLength(), 0)
}

func (o ScMutableAgentArray) GetAgent(index int32) ScMutableAgent {
	return ScMutableAgent{objId: o.objId, keyId: index}
}

func (o ScMutableAgentArray) Immutable() ScImmutableAgentArray {
	return ScImmutableAgentArray{objId: o.objId}
}

func (o ScMutableAgentArray) Length() int32 {
	return int32(GetInt(o.objId, KeyLength()))
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScMutableBytes struct {
	objId int32
	keyId int32
}

func (o ScMutableBytes) Exists() bool {
	return Exists(o.objId, o.keyId)
}

func (o ScMutableBytes) SetValue(value []byte) {
	SetBytes(o.objId, o.keyId, value)
}

func (o ScMutableBytes) Value() []byte {
	return GetBytes(o.objId, o.keyId)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScMutableBytesArray struct {
	objId int32
}

func (o ScMutableBytesArray) Clear() {
	SetInt(o.objId, KeyLength(), 0)
}

func (o ScMutableBytesArray) GetBytes(index int32) ScMutableBytes {
	return ScMutableBytes{objId: o.objId, keyId: index}
}

func (o ScMutableBytesArray) Immutable() ScImmutableBytesArray {
	return ScImmutableBytesArray{objId: o.objId}
}

func (o ScMutableBytesArray) Length() int32 {
	return int32(GetInt(o.objId, KeyLength()))
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScMutableColor struct {
	objId int32
	keyId int32
}

func (o ScMutableColor) Exists() bool {
	return Exists(o.objId, o.keyId)
}

func (o ScMutableColor) SetValue(value *ScColor) {
	SetBytes(o.objId, o.keyId, value.Bytes())
}

func (o ScMutableColor) Value() *ScColor {
	return NewScColor(GetBytes(o.objId, o.keyId))
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScMutableColorArray struct {
	objId int32
}

func (o ScMutableColorArray) Clear() {
	SetInt(o.objId, KeyLength(), 0)
}

func (o ScMutableColorArray) GetColor(index int32) ScMutableColor {
	return ScMutableColor{objId: o.objId, keyId: index}
}

func (o ScMutableColorArray) Immutable() ScImmutableColorArray {
	return ScImmutableColorArray{objId: o.objId}
}

func (o ScMutableColorArray) Length() int32 {
	return int32(GetInt(o.objId, KeyLength()))
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScMutableInt struct {
	objId int32
	keyId int32
}

func (o ScMutableInt) Exists() bool {
	return Exists(o.objId, o.keyId)
}

func (o ScMutableInt) SetValue(value int64) {
	SetInt(o.objId, o.keyId, value)
}

func (o ScMutableInt) Value() int64 {
	return GetInt(o.objId, o.keyId)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScMutableIntArray struct {
	objId int32
}

func (o ScMutableIntArray) Clear() {
	SetInt(o.objId, KeyLength(), 0)
}

func (o ScMutableIntArray) GetInt(index int32) ScMutableInt {
	return ScMutableInt{objId: o.objId, keyId: index}
}

func (o ScMutableIntArray) Immutable() ScImmutableIntArray {
	return ScImmutableIntArray{objId: o.objId}
}

func (o ScMutableIntArray) Length() int32 {
	return int32(GetInt(o.objId, KeyLength()))
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScMutableMap struct {
	objId int32
}

func (o ScMutableMap) Clear() {
	SetInt(o.objId, KeyLength(), 0)
}

func (o ScMutableMap) GetAddress(key KeyId) ScMutableAddress {
	return ScMutableAddress{objId: o.objId, keyId: key.GetId()}
}

func (o ScMutableMap) GetAddressArray(key KeyId) ScMutableAddressArray {
	arrId := GetObjectId(o.objId, key.GetId(), TYPE_BYTES_ARRAY)
	return ScMutableAddressArray{objId: arrId}
}

func (o ScMutableMap) GetAgent(key KeyId) ScMutableAgent {
	return ScMutableAgent{objId: o.objId, keyId: key.GetId()}
}

func (o ScMutableMap) GetAgentArray(key KeyId) ScMutableAgentArray {
	arrId := GetObjectId(o.objId, key.GetId(), TYPE_BYTES_ARRAY)
	return ScMutableAgentArray{objId: arrId}
}

func (o ScMutableMap) GetBytes(key KeyId) ScMutableBytes {
	return ScMutableBytes{objId: o.objId, keyId: key.GetId()}
}

func (o ScMutableMap) GetBytesArray(key KeyId) ScMutableBytesArray {
	arrId := GetObjectId(o.objId, key.GetId(), TYPE_BYTES_ARRAY)
	return ScMutableBytesArray{objId: arrId}
}

func (o ScMutableMap) GetColor(key KeyId) ScMutableColor {
	return ScMutableColor{objId: o.objId, keyId: key.GetId()}
}

func (o ScMutableMap) GetColorArray(key KeyId) ScMutableColorArray {
	arrId := GetObjectId(o.objId, key.GetId(), TYPE_BYTES_ARRAY)
	return ScMutableColorArray{objId: arrId}
}

func (o ScMutableMap) GetInt(key KeyId) ScMutableInt {
	return ScMutableInt{objId: o.objId, keyId: key.GetId()}
}

func (o ScMutableMap) GetIntArray(key KeyId) ScMutableIntArray {
	arrId := GetObjectId(o.objId, key.GetId(), TYPE_INT_ARRAY)
	return ScMutableIntArray{objId: arrId}
}

func (o ScMutableMap) GetMap(key KeyId) ScMutableMap {
	mapId := GetObjectId(o.objId, key.GetId(), TYPE_MAP)
	return ScMutableMap{objId: mapId}
}

func (o ScMutableMap) GetMapArray(key KeyId) ScMutableMapArray {
	arrId := GetObjectId(o.objId, key.GetId(), TYPE_MAP_ARRAY)
	return ScMutableMapArray{objId: arrId}
}

func (o ScMutableMap) GetString(key KeyId) ScMutableString {
	return ScMutableString{objId: o.objId, keyId: key.GetId()}
}

func (o ScMutableMap) GetStringArray(key KeyId) ScMutableStringArray {
	arrId := GetObjectId(o.objId, key.GetId(), TYPE_STRING_ARRAY)
	return ScMutableStringArray{objId: arrId}
}

func (o ScMutableMap) Immutable() ScImmutableMap {
	return ScImmutableMap{objId: o.objId}
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScMutableMapArray struct {
	objId int32
}

func (o ScMutableMapArray) Clear() {
	SetInt(o.objId, KeyLength(), 0)
}

func (o ScMutableMapArray) GetMap(index int32) ScMutableMap {
	mapId := GetObjectId(o.objId, index, TYPE_MAP)
	return ScMutableMap{objId: mapId}
}

func (o ScMutableMapArray) Immutable() ScImmutableMapArray {
	return ScImmutableMapArray{objId: o.objId}
}

func (o ScMutableMapArray) Length() int32 {
	return int32(GetInt(o.objId, KeyLength()))
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScMutableString struct {
	objId int32
	keyId int32
}

func (o ScMutableString) Exists() bool {
	return Exists(o.objId, o.keyId)
}

func (o ScMutableString) SetValue(value string) {
	SetString(o.objId, o.keyId, value)
}

func (o ScMutableString) Value() string {
	return GetString(o.objId, o.keyId)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScMutableStringArray struct {
	objId int32
}

func (o ScMutableStringArray) Clear() {
	SetInt(o.objId, KeyLength(), 0)
}

func (o ScMutableStringArray) GetString(index int32) ScMutableString {
	return ScMutableString{objId: o.objId, keyId: index}
}

func (o ScMutableStringArray) Immutable() ScImmutableStringArray {
	return ScImmutableStringArray{objId: o.objId}
}

func (o ScMutableStringArray) Length() int32 {
	return int32(GetInt(o.objId, KeyLength()))
}
