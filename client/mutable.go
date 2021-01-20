// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package client

import "strconv"

var Root = ScMutableMap{objId: 1}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScMutableAddress struct {
	objId int32
	keyId Key32
}

func (o ScMutableAddress) Exists() bool {
	return Exists(o.objId, o.keyId)
}

func (o ScMutableAddress) SetValue(value *ScAddress) {
	SetBytes(o.objId, o.keyId, value.Bytes())
}

func (o ScMutableAddress) String() string {
	return o.Value().String()
}

func (o ScMutableAddress) Value() *ScAddress {
	return NewScAddress(GetBytes(o.objId, o.keyId))
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScMutableAddressArray struct {
	objId int32
}

func (o ScMutableAddressArray) Clear() {
	SetClear(o.objId)
}

func (o ScMutableAddressArray) GetAddress(index int32) ScMutableAddress {
	return ScMutableAddress{objId: o.objId, keyId: Key32(index)}
}

func (o ScMutableAddressArray) Immutable() ScImmutableAddressArray {
	return ScImmutableAddressArray{objId: o.objId}
}

func (o ScMutableAddressArray) Length() int32 {
	return GetLength(o.objId)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScMutableAgent struct {
	objId int32
	keyId Key32
}

func (o ScMutableAgent) Exists() bool {
	return Exists(o.objId, o.keyId)
}

func (o ScMutableAgent) SetValue(value *ScAgent) {
	SetBytes(o.objId, o.keyId, value.Bytes())
}

func (o ScMutableAgent) String() string {
	return o.Value().String()
}

func (o ScMutableAgent) Value() *ScAgent {
	return NewScAgent(GetBytes(o.objId, o.keyId))
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScMutableAgentArray struct {
	objId int32
}

func (o ScMutableAgentArray) Clear() {
	SetClear(o.objId)
}

func (o ScMutableAgentArray) GetAgent(index int32) ScMutableAgent {
	return ScMutableAgent{objId: o.objId, keyId: Key32(index)}
}

func (o ScMutableAgentArray) Immutable() ScImmutableAgentArray {
	return ScImmutableAgentArray{objId: o.objId}
}

func (o ScMutableAgentArray) Length() int32 {
	return GetLength(o.objId)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScMutableBytes struct {
	objId int32
	keyId Key32
}

func (o ScMutableBytes) Exists() bool {
	return Exists(o.objId, o.keyId)
}

func (o ScMutableBytes) SetValue(value []byte) {
	SetBytes(o.objId, o.keyId, value)
}

func (o ScMutableBytes) String() string {
	return base58Encode(o.Value())
}

func (o ScMutableBytes) Value() []byte {
	return GetBytes(o.objId, o.keyId)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScMutableBytesArray struct {
	objId int32
}

func (o ScMutableBytesArray) Clear() {
	SetClear(o.objId)
}

func (o ScMutableBytesArray) GetBytes(index int32) ScMutableBytes {
	return ScMutableBytes{objId: o.objId, keyId: Key32(index)}
}

func (o ScMutableBytesArray) Immutable() ScImmutableBytesArray {
	return ScImmutableBytesArray{objId: o.objId}
}

func (o ScMutableBytesArray) Length() int32 {
	return GetLength(o.objId)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScMutableColor struct {
	objId int32
	keyId Key32
}

func (o ScMutableColor) Exists() bool {
	return Exists(o.objId, o.keyId)
}

func (o ScMutableColor) SetValue(value *ScColor) {
	SetBytes(o.objId, o.keyId, value.Bytes())
}

func (o ScMutableColor) String() string {
	return o.Value().String()
}

func (o ScMutableColor) Value() *ScColor {
	return NewScColor(GetBytes(o.objId, o.keyId))
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScMutableColorArray struct {
	objId int32
}

func (o ScMutableColorArray) Clear() {
	SetClear(o.objId)
}

func (o ScMutableColorArray) GetColor(index int32) ScMutableColor {
	return ScMutableColor{objId: o.objId, keyId: Key32(index)}
}

func (o ScMutableColorArray) Immutable() ScImmutableColorArray {
	return ScImmutableColorArray{objId: o.objId}
}

func (o ScMutableColorArray) Length() int32 {
	return GetLength(o.objId)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScMutableHash struct {
	objId int32
	keyId Key32
}

func (o ScMutableHash) Exists() bool {
	return Exists(o.objId, o.keyId)
}

func (o ScMutableHash) SetValue(value *ScHash) {
	SetBytes(o.objId, o.keyId, value.Bytes())
}

func (o ScMutableHash) String() string {
	return o.Value().String()
}

func (o ScMutableHash) Value() *ScHash {
	return NewScHash(GetBytes(o.objId, o.keyId))
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScMutableHashArray struct {
	objId int32
}

func (o ScMutableHashArray) Clear() {
	SetClear(o.objId)
}

func (o ScMutableHashArray) GetHash(index int32) ScMutableHash {
	return ScMutableHash{objId: o.objId, keyId: Key32(index)}
}

func (o ScMutableHashArray) Immutable() ScImmutableHashArray {
	return ScImmutableHashArray{objId: o.objId}
}

func (o ScMutableHashArray) Length() int32 {
	return GetLength(o.objId)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScMutableHname struct {
	objId int32
	keyId Key32
}

func (o ScMutableHname) Exists() bool {
	return Exists(o.objId, o.keyId)
}

func (o ScMutableHname) SetValue(value Hname) {
	SetInt(o.objId, o.keyId, int64(value))
}

func (o ScMutableHname) String() string {
	return o.Value().String()
}

func (o ScMutableHname) Value() Hname {
	return Hname(GetInt(o.objId, o.keyId))
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScMutableInt struct {
	objId int32
	keyId Key32
}

func (o ScMutableInt) Exists() bool {
	return Exists(o.objId, o.keyId)
}

func (o ScMutableInt) SetValue(value int64) {
	SetInt(o.objId, o.keyId, value)
}

func (o ScMutableInt) String() string {
	return strconv.FormatInt(o.Value(), 10)
}

func (o ScMutableInt) Value() int64 {
	return GetInt(o.objId, o.keyId)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScMutableIntArray struct {
	objId int32
}

func (o ScMutableIntArray) Clear() {
	SetClear(o.objId)
}

func (o ScMutableIntArray) GetInt(index int32) ScMutableInt {
	return ScMutableInt{objId: o.objId, keyId: Key32(index)}
}

func (o ScMutableIntArray) Immutable() ScImmutableIntArray {
	return ScImmutableIntArray{objId: o.objId}
}

func (o ScMutableIntArray) Length() int32 {
	return GetLength(o.objId)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScMutableMap struct {
	objId int32
}

func NewScMutableMap() *ScMutableMap {
	maps := Root.GetMapArray(KeyMaps)
	return &ScMutableMap{objId: maps.GetMap(maps.Length()).objId}
}

func (o ScMutableMap) Clear() {
	SetClear(o.objId)
}

func (o ScMutableMap) GetAddress(key MapKey) ScMutableAddress {
	return ScMutableAddress{objId: o.objId, keyId: key.KeyId()}
}

func (o ScMutableMap) GetAddressArray(key MapKey) ScMutableAddressArray {
	arrId := GetObjectId(o.objId, key.KeyId(), TYPE_ADDRESS|TYPE_ARRAY)
	return ScMutableAddressArray{objId: arrId}
}

func (o ScMutableMap) GetAgent(key MapKey) ScMutableAgent {
	return ScMutableAgent{objId: o.objId, keyId: key.KeyId()}
}

func (o ScMutableMap) GetAgentArray(key MapKey) ScMutableAgentArray {
	arrId := GetObjectId(o.objId, key.KeyId(), TYPE_AGENT|TYPE_ARRAY)
	return ScMutableAgentArray{objId: arrId}
}

func (o ScMutableMap) GetBytes(key MapKey) ScMutableBytes {
	return ScMutableBytes{objId: o.objId, keyId: key.KeyId()}
}

func (o ScMutableMap) GetBytesArray(key MapKey) ScMutableBytesArray {
	arrId := GetObjectId(o.objId, key.KeyId(), TYPE_BYTES|TYPE_ARRAY)
	return ScMutableBytesArray{objId: arrId}
}

func (o ScMutableMap) GetColor(key MapKey) ScMutableColor {
	return ScMutableColor{objId: o.objId, keyId: key.KeyId()}
}

func (o ScMutableMap) GetColorArray(key MapKey) ScMutableColorArray {
	arrId := GetObjectId(o.objId, key.KeyId(), TYPE_COLOR|TYPE_ARRAY)
	return ScMutableColorArray{objId: arrId}
}

func (o ScMutableMap) GetHash(key MapKey) ScMutableHash {
	return ScMutableHash{objId: o.objId, keyId: key.KeyId()}
}

func (o ScMutableMap) GetHashArray(key MapKey) ScMutableHashArray {
	arrId := GetObjectId(o.objId, key.KeyId(), TYPE_HASH|TYPE_ARRAY)
	return ScMutableHashArray{objId: arrId}
}

func (o ScMutableMap) GetHname(key MapKey) ScMutableHname {
	return ScMutableHname{objId: o.objId, keyId: key.KeyId()}
}

func (o ScMutableMap) GetInt(key MapKey) ScMutableInt {
	return ScMutableInt{objId: o.objId, keyId: key.KeyId()}
}

func (o ScMutableMap) GetIntArray(key MapKey) ScMutableIntArray {
	arrId := GetObjectId(o.objId, key.KeyId(), TYPE_INT|TYPE_ARRAY)
	return ScMutableIntArray{objId: arrId}
}

func (o ScMutableMap) GetMap(key MapKey) ScMutableMap {
	mapId := GetObjectId(o.objId, key.KeyId(), TYPE_MAP)
	return ScMutableMap{objId: mapId}
}

func (o ScMutableMap) GetMapArray(key MapKey) ScMutableMapArray {
	arrId := GetObjectId(o.objId, key.KeyId(), TYPE_MAP|TYPE_ARRAY)
	return ScMutableMapArray{objId: arrId}
}

func (o ScMutableMap) GetString(key MapKey) ScMutableString {
	return ScMutableString{objId: o.objId, keyId: key.KeyId()}
}

func (o ScMutableMap) GetStringArray(key MapKey) ScMutableStringArray {
	arrId := GetObjectId(o.objId, key.KeyId(), TYPE_STRING|TYPE_ARRAY)
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
	SetClear(o.objId)
}

func (o ScMutableMapArray) GetMap(index int32) ScMutableMap {
	mapId := GetObjectId(o.objId, Key32(index), TYPE_MAP)
	return ScMutableMap{objId: mapId}
}

func (o ScMutableMapArray) Immutable() ScImmutableMapArray {
	return ScImmutableMapArray{objId: o.objId}
}

func (o ScMutableMapArray) Length() int32 {
	return GetLength(o.objId)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScMutableString struct {
	objId int32
	keyId Key32
}

func (o ScMutableString) Exists() bool {
	return Exists(o.objId, o.keyId)
}

func (o ScMutableString) SetValue(value string) {
	SetString(o.objId, o.keyId, value)
}

func (o ScMutableString) String() string {
	return o.Value()
}

func (o ScMutableString) Value() string {
	return GetString(o.objId, o.keyId)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScMutableStringArray struct {
	objId int32
}

func (o ScMutableStringArray) Clear() {
	SetClear(o.objId)
}

func (o ScMutableStringArray) GetString(index int32) ScMutableString {
	return ScMutableString{objId: o.objId, keyId: Key32(index)}
}

func (o ScMutableStringArray) Immutable() ScImmutableStringArray {
	return ScImmutableStringArray{objId: o.objId}
}

func (o ScMutableStringArray) Length() int32 {
	return GetLength(o.objId)
}
