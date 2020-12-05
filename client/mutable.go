// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package client

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

type ScMutableKeyMap struct {
	objId int32
}

func (o ScMutableKeyMap) Clear() {
	SetInt(o.objId, KeyLength(), 0)
}

func (o ScMutableKeyMap) GetAddress(key []byte) ScMutableAddress {
	return ScMutableAddress{objId: o.objId, keyId: GetKey(key)}
}

func (o ScMutableKeyMap) GetAddressArray(key []byte) ScMutableAddressArray {
	arrId := GetObjectId(o.objId, GetKey(key), TYPE_BYTES_ARRAY)
	return ScMutableAddressArray{objId: arrId}
}

func (o ScMutableKeyMap) GetAgent(key []byte) ScMutableAgent {
	return ScMutableAgent{objId: o.objId, keyId: GetKey(key)}
}

func (o ScMutableKeyMap) GetAgentArray(key []byte) ScMutableAgentArray {
	arrId := GetObjectId(o.objId, GetKey(key), TYPE_BYTES_ARRAY)
	return ScMutableAgentArray{objId: arrId}
}

func (o ScMutableKeyMap) GetBytes(key []byte) ScMutableBytes {
	return ScMutableBytes{objId: o.objId, keyId: GetKey(key)}
}

func (o ScMutableKeyMap) GetBytesArray(key []byte) ScMutableBytesArray {
	arrId := GetObjectId(o.objId, GetKey(key), TYPE_BYTES_ARRAY)
	return ScMutableBytesArray{objId: arrId}
}

func (o ScMutableKeyMap) GetColor(key []byte) ScMutableColor {
	return ScMutableColor{objId: o.objId, keyId: GetKey(key)}
}

func (o ScMutableKeyMap) GetColorArray(key []byte) ScMutableColorArray {
	arrId := GetObjectId(o.objId, GetKey(key), TYPE_BYTES_ARRAY)
	return ScMutableColorArray{objId: arrId}
}

func (o ScMutableKeyMap) GetInt(key []byte) ScMutableInt {
	return ScMutableInt{objId: o.objId, keyId: GetKey(key)}
}

func (o ScMutableKeyMap) GetIntArray(key []byte) ScMutableIntArray {
	arrId := GetObjectId(o.objId, GetKey(key), TYPE_INT_ARRAY)
	return ScMutableIntArray{objId: arrId}
}

func (o ScMutableKeyMap) GetKeyMap(key []byte) ScMutableKeyMap {
	mapId := GetObjectId(o.objId, GetKey(key), TYPE_MAP)
	return ScMutableKeyMap{objId: mapId}
}

func (o ScMutableKeyMap) GetMap(key []byte) ScMutableMap {
	mapId := GetObjectId(o.objId, GetKey(key), TYPE_MAP)
	return ScMutableMap{objId: mapId}
}

func (o ScMutableKeyMap) GetMapArray(key []byte) ScMutableMapArray {
	arrId := GetObjectId(o.objId, GetKey(key), TYPE_MAP_ARRAY)
	return ScMutableMapArray{objId: arrId}
}

func (o ScMutableKeyMap) GetString(key []byte) ScMutableString {
	return ScMutableString{objId: o.objId, keyId: GetKey(key)}
}

func (o ScMutableKeyMap) GetStringArray(key []byte) ScMutableStringArray {
	arrId := GetObjectId(o.objId, GetKey(key), TYPE_STRING_ARRAY)
	return ScMutableStringArray{objId: arrId}
}

func (o ScMutableKeyMap) Immutable() ScImmutableKeyMap {
	return ScImmutableKeyMap{objId: o.objId}
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScMutableMap struct {
	objId int32
}

func (o ScMutableMap) Clear() {
	SetInt(o.objId, KeyLength(), 0)
}

func (o ScMutableMap) GetAddress(key string) ScMutableAddress {
	return ScMutableAddress{objId: o.objId, keyId: GetKeyId(key)}
}

func (o ScMutableMap) GetAddressArray(key string) ScMutableAddressArray {
	arrId := GetObjectId(o.objId, GetKeyId(key), TYPE_BYTES_ARRAY)
	return ScMutableAddressArray{objId: arrId}
}

func (o ScMutableMap) GetAgent(key string) ScMutableAgent {
	return ScMutableAgent{objId: o.objId, keyId: GetKeyId(key)}
}

func (o ScMutableMap) GetAgentArray(key string) ScMutableAgentArray {
	arrId := GetObjectId(o.objId, GetKeyId(key), TYPE_BYTES_ARRAY)
	return ScMutableAgentArray{objId: arrId}
}

func (o ScMutableMap) GetBytes(key string) ScMutableBytes {
	return ScMutableBytes{objId: o.objId, keyId: GetKeyId(key)}
}

func (o ScMutableMap) GetBytesArray(key string) ScMutableBytesArray {
	arrId := GetObjectId(o.objId, GetKeyId(key), TYPE_BYTES_ARRAY)
	return ScMutableBytesArray{objId: arrId}
}

func (o ScMutableMap) GetColor(key string) ScMutableColor {
	return ScMutableColor{objId: o.objId, keyId: GetKeyId(key)}
}

func (o ScMutableMap) GetColorArray(key string) ScMutableColorArray {
	arrId := GetObjectId(o.objId, GetKeyId(key), TYPE_BYTES_ARRAY)
	return ScMutableColorArray{objId: arrId}
}

func (o ScMutableMap) GetInt(key string) ScMutableInt {
	return ScMutableInt{objId: o.objId, keyId: GetKeyId(key)}
}

func (o ScMutableMap) GetIntArray(key string) ScMutableIntArray {
	arrId := GetObjectId(o.objId, GetKeyId(key), TYPE_INT_ARRAY)
	return ScMutableIntArray{objId: arrId}
}

func (o ScMutableMap) GetKeyMap(key string) ScMutableKeyMap {
	mapId := GetObjectId(o.objId, GetKeyId(key), TYPE_MAP)
	return ScMutableKeyMap{objId: mapId}
}

func (o ScMutableMap) GetMap(key string) ScMutableMap {
	mapId := GetObjectId(o.objId, GetKeyId(key), TYPE_MAP)
	return ScMutableMap{objId: mapId}
}

func (o ScMutableMap) GetMapArray(key string) ScMutableMapArray {
	arrId := GetObjectId(o.objId, GetKeyId(key), TYPE_MAP_ARRAY)
	return ScMutableMapArray{objId: arrId}
}

func (o ScMutableMap) GetString(key string) ScMutableString {
	return ScMutableString{objId: o.objId, keyId: GetKeyId(key)}
}

func (o ScMutableMap) GetStringArray(key string) ScMutableStringArray {
	arrId := GetObjectId(o.objId, GetKeyId(key), TYPE_STRING_ARRAY)
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

func (o ScMutableMapArray) GetKeyMap(index int32) ScMutableKeyMap {
	mapId := GetObjectId(o.objId, index, TYPE_MAP)
	return ScMutableKeyMap{objId: mapId}
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
