// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package wasmlib

import (
	"encoding/binary"
	"strconv"
)

var Root = ScMutableMap{objID: 1}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScMutableAddress struct {
	objID int32
	keyID Key32
}

func NewScMutableAddress(objID int32, keyID Key32) ScMutableAddress {
	return ScMutableAddress{objID: objID, keyID: keyID}
}

func (o ScMutableAddress) Exists() bool {
	return Exists(o.objID, o.keyID, TYPE_ADDRESS)
}

func (o ScMutableAddress) SetValue(value ScAddress) {
	SetBytes(o.objID, o.keyID, TYPE_ADDRESS, value.Bytes())
}

func (o ScMutableAddress) String() string {
	return o.Value().String()
}

func (o ScMutableAddress) Value() ScAddress {
	return NewScAddressFromBytes(GetBytes(o.objID, o.keyID, TYPE_ADDRESS))
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScMutableAddressArray struct {
	objID int32
}

func (o ScMutableAddressArray) Clear() {
	Clear(o.objID)
}

func (o ScMutableAddressArray) GetAddress(index int32) ScMutableAddress {
	return ScMutableAddress{objID: o.objID, keyID: Key32(index)}
}

func (o ScMutableAddressArray) Immutable() ScImmutableAddressArray {
	return ScImmutableAddressArray(o)
}

func (o ScMutableAddressArray) Length() int32 {
	return GetLength(o.objID)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScMutableAgentID struct {
	objID int32
	keyID Key32
}

func NewScMutableAgentID(objID int32, keyID Key32) ScMutableAgentID {
	return ScMutableAgentID{objID: objID, keyID: keyID}
}

func (o ScMutableAgentID) Exists() bool {
	return Exists(o.objID, o.keyID, TYPE_AGENT_ID)
}

func (o ScMutableAgentID) SetValue(value ScAgentID) {
	SetBytes(o.objID, o.keyID, TYPE_AGENT_ID, value.Bytes())
}

func (o ScMutableAgentID) String() string {
	return o.Value().String()
}

func (o ScMutableAgentID) Value() ScAgentID {
	return NewScAgentIDFromBytes(GetBytes(o.objID, o.keyID, TYPE_AGENT_ID))
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScMutableAgentIDArray struct {
	objID int32
}

func (o ScMutableAgentIDArray) Clear() {
	Clear(o.objID)
}

func (o ScMutableAgentIDArray) GetAgentID(index int32) ScMutableAgentID {
	return ScMutableAgentID{objID: o.objID, keyID: Key32(index)}
}

func (o ScMutableAgentIDArray) Immutable() ScImmutableAgentIDArray {
	return ScImmutableAgentIDArray(o)
}

func (o ScMutableAgentIDArray) Length() int32 {
	return GetLength(o.objID)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScMutableBytes struct {
	objID int32
	keyID Key32
}

func NewScMutableBytes(objID int32, keyID Key32) ScMutableBytes {
	return ScMutableBytes{objID: objID, keyID: keyID}
}

func (o ScMutableBytes) Exists() bool {
	return Exists(o.objID, o.keyID, TYPE_BYTES)
}

func (o ScMutableBytes) SetValue(value []byte) {
	SetBytes(o.objID, o.keyID, TYPE_BYTES, value)
}

func (o ScMutableBytes) String() string {
	return base58Encode(o.Value())
}

func (o ScMutableBytes) Value() []byte {
	return GetBytes(o.objID, o.keyID, TYPE_BYTES)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScMutableBytesArray struct {
	objID int32
}

func (o ScMutableBytesArray) Clear() {
	Clear(o.objID)
}

func (o ScMutableBytesArray) GetBytes(index int32) ScMutableBytes {
	return ScMutableBytes{objID: o.objID, keyID: Key32(index)}
}

func (o ScMutableBytesArray) Immutable() ScImmutableBytesArray {
	return ScImmutableBytesArray(o)
}

func (o ScMutableBytesArray) Length() int32 {
	return GetLength(o.objID)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScMutableChainID struct {
	objID int32
	keyID Key32
}

func NewScMutableChainID(objID int32, keyID Key32) ScMutableChainID {
	return ScMutableChainID{objID: objID, keyID: keyID}
}

func (o ScMutableChainID) Exists() bool {
	return Exists(o.objID, o.keyID, TYPE_CHAIN_ID)
}

func (o ScMutableChainID) SetValue(value ScChainID) {
	SetBytes(o.objID, o.keyID, TYPE_CHAIN_ID, value.Bytes())
}

func (o ScMutableChainID) String() string {
	return o.Value().String()
}

func (o ScMutableChainID) Value() ScChainID {
	return NewScChainIDFromBytes(GetBytes(o.objID, o.keyID, TYPE_CHAIN_ID))
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScMutableChainIDArray struct {
	objID int32
}

func (o ScMutableChainIDArray) Clear() {
	Clear(o.objID)
}

func (o ScMutableChainIDArray) GetChainID(index int32) ScMutableChainID {
	return ScMutableChainID{objID: o.objID, keyID: Key32(index)}
}

func (o ScMutableChainIDArray) Immutable() ScImmutableChainIDArray {
	return ScImmutableChainIDArray(o)
}

func (o ScMutableChainIDArray) Length() int32 {
	return GetLength(o.objID)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScMutableColor struct {
	objID int32
	keyID Key32
}

func NewScMutableColor(objID int32, keyID Key32) ScMutableColor {
	return ScMutableColor{objID: objID, keyID: keyID}
}

func (o ScMutableColor) Exists() bool {
	return Exists(o.objID, o.keyID, TYPE_COLOR)
}

func (o ScMutableColor) SetValue(value ScColor) {
	SetBytes(o.objID, o.keyID, TYPE_COLOR, value.Bytes())
}

func (o ScMutableColor) String() string {
	return o.Value().String()
}

func (o ScMutableColor) Value() ScColor {
	return NewScColorFromBytes(GetBytes(o.objID, o.keyID, TYPE_COLOR))
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScMutableColorArray struct {
	objID int32
}

func (o ScMutableColorArray) Clear() {
	Clear(o.objID)
}

func (o ScMutableColorArray) GetColor(index int32) ScMutableColor {
	return ScMutableColor{objID: o.objID, keyID: Key32(index)}
}

func (o ScMutableColorArray) Immutable() ScImmutableColorArray {
	return ScImmutableColorArray(o)
}

func (o ScMutableColorArray) Length() int32 {
	return GetLength(o.objID)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScMutableHash struct {
	objID int32
	keyID Key32
}

func NewScMutableHash(objID int32, keyID Key32) ScMutableHash {
	return ScMutableHash{objID: objID, keyID: keyID}
}

func (o ScMutableHash) Exists() bool {
	return Exists(o.objID, o.keyID, TYPE_HASH)
}

func (o ScMutableHash) SetValue(value ScHash) {
	SetBytes(o.objID, o.keyID, TYPE_HASH, value.Bytes())
}

func (o ScMutableHash) String() string {
	return o.Value().String()
}

func (o ScMutableHash) Value() ScHash {
	return NewScHashFromBytes(GetBytes(o.objID, o.keyID, TYPE_HASH))
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScMutableHashArray struct {
	objID int32
}

func (o ScMutableHashArray) Clear() {
	Clear(o.objID)
}

func (o ScMutableHashArray) GetHash(index int32) ScMutableHash {
	return ScMutableHash{objID: o.objID, keyID: Key32(index)}
}

func (o ScMutableHashArray) Immutable() ScImmutableHashArray {
	return ScImmutableHashArray(o)
}

func (o ScMutableHashArray) Length() int32 {
	return GetLength(o.objID)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScMutableHname struct {
	objID int32
	keyID Key32
}

func NewScMutableHname(objID int32, keyID Key32) ScMutableHname {
	return ScMutableHname{objID: objID, keyID: keyID}
}

func (o ScMutableHname) Exists() bool {
	return Exists(o.objID, o.keyID, TYPE_HNAME)
}

func (o ScMutableHname) SetValue(value ScHname) {
	SetBytes(o.objID, o.keyID, TYPE_HNAME, value.Bytes())
}

func (o ScMutableHname) String() string {
	return o.Value().String()
}

func (o ScMutableHname) Value() ScHname {
	return NewScHnameFromBytes(GetBytes(o.objID, o.keyID, TYPE_HNAME))
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScMutableHnameArray struct {
	objID int32
}

func (o ScMutableHnameArray) Clear() {
	Clear(o.objID)
}

func (o ScMutableHnameArray) GetHname(index int32) ScMutableHname {
	return ScMutableHname{objID: o.objID, keyID: Key32(index)}
}

func (o ScMutableHnameArray) Immutable() ScImmutableHnameArray {
	return ScImmutableHnameArray(o)
}

func (o ScMutableHnameArray) Length() int32 {
	return GetLength(o.objID)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScMutableInt16 struct {
	objID int32
	keyID Key32
}

func NewScMutableInt16(objID int32, keyID Key32) ScMutableInt16 {
	return ScMutableInt16{objID: objID, keyID: keyID}
}

func (o ScMutableInt16) Exists() bool {
	return Exists(o.objID, o.keyID, TYPE_INT16)
}

func (o ScMutableInt16) SetValue(value int16) {
	bytes := make([]byte, 2)
	binary.LittleEndian.PutUint16(bytes, uint16(value))
	SetBytes(o.objID, o.keyID, TYPE_INT16, bytes)
}

func (o ScMutableInt16) String() string {
	return strconv.FormatInt(int64(o.Value()), 10)
}

func (o ScMutableInt16) Value() int16 {
	bytes := GetBytes(o.objID, o.keyID, TYPE_INT16)
	return int16(binary.LittleEndian.Uint16(bytes))
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScMutableInt16Array struct {
	objID int32
}

func (o ScMutableInt16Array) Clear() {
	Clear(o.objID)
}

func (o ScMutableInt16Array) GetInt16(index int32) ScMutableInt16 {
	return ScMutableInt16{objID: o.objID, keyID: Key32(index)}
}

func (o ScMutableInt16Array) Immutable() ScImmutableInt16Array {
	return ScImmutableInt16Array(o)
}

func (o ScMutableInt16Array) Length() int32 {
	return GetLength(o.objID)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScMutableInt32 struct {
	objID int32
	keyID Key32
}

func NewScMutableInt32(objID int32, keyID Key32) ScMutableInt32 {
	return ScMutableInt32{objID: objID, keyID: keyID}
}

func (o ScMutableInt32) Exists() bool {
	return Exists(o.objID, o.keyID, TYPE_INT32)
}

func (o ScMutableInt32) SetValue(value int32) {
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, uint32(value))
	SetBytes(o.objID, o.keyID, TYPE_INT32, bytes)
}

func (o ScMutableInt32) String() string {
	return strconv.FormatInt(int64(o.Value()), 10)
}

func (o ScMutableInt32) Value() int32 {
	bytes := GetBytes(o.objID, o.keyID, TYPE_INT32)
	return int32(binary.LittleEndian.Uint32(bytes))
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScMutableInt32Array struct {
	objID int32
}

func (o ScMutableInt32Array) Clear() {
	Clear(o.objID)
}

func (o ScMutableInt32Array) GetInt32(index int32) ScMutableInt32 {
	return ScMutableInt32{objID: o.objID, keyID: Key32(index)}
}

func (o ScMutableInt32Array) Immutable() ScImmutableInt32Array {
	return ScImmutableInt32Array(o)
}

func (o ScMutableInt32Array) Length() int32 {
	return GetLength(o.objID)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScMutableInt64 struct {
	objID int32
	keyID Key32
}

func NewScMutableInt64(objID int32, keyID Key32) ScMutableInt64 {
	return ScMutableInt64{objID: objID, keyID: keyID}
}

func (o ScMutableInt64) Exists() bool {
	return Exists(o.objID, o.keyID, TYPE_INT64)
}

func (o ScMutableInt64) SetValue(value int64) {
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, uint64(value))
	SetBytes(o.objID, o.keyID, TYPE_INT64, bytes)
}

func (o ScMutableInt64) String() string {
	return strconv.FormatInt(o.Value(), 10)
}

func (o ScMutableInt64) Value() int64 {
	bytes := GetBytes(o.objID, o.keyID, TYPE_INT64)
	return int64(binary.LittleEndian.Uint64(bytes))
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScMutableInt64Array struct {
	objID int32
}

func (o ScMutableInt64Array) Clear() {
	Clear(o.objID)
}

func (o ScMutableInt64Array) GetInt64(index int32) ScMutableInt64 {
	return ScMutableInt64{objID: o.objID, keyID: Key32(index)}
}

func (o ScMutableInt64Array) Immutable() ScImmutableInt64Array {
	return ScImmutableInt64Array(o)
}

func (o ScMutableInt64Array) Length() int32 {
	return GetLength(o.objID)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScMutableMap struct {
	objID int32
}

func NewScMutableMap() *ScMutableMap {
	maps := Root.GetMapArray(KeyMaps)
	return &ScMutableMap{objID: maps.GetMap(maps.Length()).objID}
}

func (o ScMutableMap) CallFunc(keyID Key32, params []byte) []byte {
	return CallFunc(o.objID, keyID, params)
}

func (o ScMutableMap) Clear() {
	Clear(o.objID)
}

func (o ScMutableMap) GetAddress(key MapKey) ScMutableAddress {
	return ScMutableAddress{objID: o.objID, keyID: key.KeyID()}
}

func (o ScMutableMap) GetAddressArray(key MapKey) ScMutableAddressArray {
	arrID := GetObjectID(o.objID, key.KeyID(), TYPE_ADDRESS|TYPE_ARRAY)
	return ScMutableAddressArray{objID: arrID}
}

func (o ScMutableMap) GetAgentID(key MapKey) ScMutableAgentID {
	return ScMutableAgentID{objID: o.objID, keyID: key.KeyID()}
}

func (o ScMutableMap) GetAgentIDArray(key MapKey) ScMutableAgentIDArray {
	arrID := GetObjectID(o.objID, key.KeyID(), TYPE_AGENT_ID|TYPE_ARRAY)
	return ScMutableAgentIDArray{objID: arrID}
}

func (o ScMutableMap) GetBytes(key MapKey) ScMutableBytes {
	return ScMutableBytes{objID: o.objID, keyID: key.KeyID()}
}

func (o ScMutableMap) GetBytesArray(key MapKey) ScMutableBytesArray {
	arrID := GetObjectID(o.objID, key.KeyID(), TYPE_BYTES|TYPE_ARRAY)
	return ScMutableBytesArray{objID: arrID}
}

func (o ScMutableMap) GetChainID(key MapKey) ScMutableChainID {
	return ScMutableChainID{objID: o.objID, keyID: key.KeyID()}
}

func (o ScMutableMap) GetChainIDArray(key MapKey) ScMutableChainIDArray {
	arrID := GetObjectID(o.objID, key.KeyID(), TYPE_CHAIN_ID|TYPE_ARRAY)
	return ScMutableChainIDArray{objID: arrID}
}

func (o ScMutableMap) GetColor(key MapKey) ScMutableColor {
	return ScMutableColor{objID: o.objID, keyID: key.KeyID()}
}

func (o ScMutableMap) GetColorArray(key MapKey) ScMutableColorArray {
	arrID := GetObjectID(o.objID, key.KeyID(), TYPE_COLOR|TYPE_ARRAY)
	return ScMutableColorArray{objID: arrID}
}

func (o ScMutableMap) GetHash(key MapKey) ScMutableHash {
	return ScMutableHash{objID: o.objID, keyID: key.KeyID()}
}

func (o ScMutableMap) GetHashArray(key MapKey) ScMutableHashArray {
	arrID := GetObjectID(o.objID, key.KeyID(), TYPE_HASH|TYPE_ARRAY)
	return ScMutableHashArray{objID: arrID}
}

func (o ScMutableMap) GetHname(key MapKey) ScMutableHname {
	return ScMutableHname{objID: o.objID, keyID: key.KeyID()}
}

func (o ScMutableMap) GetHnameArray(key MapKey) ScMutableHnameArray {
	arrID := GetObjectID(o.objID, key.KeyID(), TYPE_HNAME|TYPE_ARRAY)
	return ScMutableHnameArray{objID: arrID}
}

func (o ScMutableMap) GetInt16(key MapKey) ScMutableInt16 {
	return ScMutableInt16{objID: o.objID, keyID: key.KeyID()}
}

func (o ScMutableMap) GetInt16Array(key MapKey) ScMutableInt16Array {
	arrID := GetObjectID(o.objID, key.KeyID(), TYPE_INT16|TYPE_ARRAY)
	return ScMutableInt16Array{objID: arrID}
}

func (o ScMutableMap) GetInt32(key MapKey) ScMutableInt32 {
	return ScMutableInt32{objID: o.objID, keyID: key.KeyID()}
}

func (o ScMutableMap) GetInt32Array(key MapKey) ScMutableInt32Array {
	arrID := GetObjectID(o.objID, key.KeyID(), TYPE_INT32|TYPE_ARRAY)
	return ScMutableInt32Array{objID: arrID}
}

func (o ScMutableMap) GetInt64(key MapKey) ScMutableInt64 {
	return ScMutableInt64{objID: o.objID, keyID: key.KeyID()}
}

func (o ScMutableMap) GetInt64Array(key MapKey) ScMutableInt64Array {
	arrID := GetObjectID(o.objID, key.KeyID(), TYPE_INT64|TYPE_ARRAY)
	return ScMutableInt64Array{objID: arrID}
}

func (o ScMutableMap) GetMap(key MapKey) ScMutableMap {
	mapID := GetObjectID(o.objID, key.KeyID(), TYPE_MAP)
	return ScMutableMap{objID: mapID}
}

func (o ScMutableMap) GetMapArray(key MapKey) ScMutableMapArray {
	arrID := GetObjectID(o.objID, key.KeyID(), TYPE_MAP|TYPE_ARRAY)
	return ScMutableMapArray{objID: arrID}
}

func (o ScMutableMap) GetRequestID(key MapKey) ScMutableRequestID {
	return ScMutableRequestID{objID: o.objID, keyID: key.KeyID()}
}

func (o ScMutableMap) GetRequestIDArray(key MapKey) ScMutableRequestIDArray {
	arrID := GetObjectID(o.objID, key.KeyID(), TYPE_REQUEST_ID|TYPE_ARRAY)
	return ScMutableRequestIDArray{objID: arrID}
}

func (o ScMutableMap) GetString(key MapKey) ScMutableString {
	return ScMutableString{objID: o.objID, keyID: key.KeyID()}
}

func (o ScMutableMap) GetStringArray(key MapKey) ScMutableStringArray {
	arrID := GetObjectID(o.objID, key.KeyID(), TYPE_STRING|TYPE_ARRAY)
	return ScMutableStringArray{objID: arrID}
}

func (o ScMutableMap) Immutable() ScImmutableMap {
	return ScImmutableMap(o)
}

func (o ScMutableMap) MapID() int32 {
	return o.objID
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScMutableMapArray struct {
	objID int32
}

func (o ScMutableMapArray) Clear() {
	Clear(o.objID)
}

func (o ScMutableMapArray) GetMap(index int32) ScMutableMap {
	mapID := GetObjectID(o.objID, Key32(index), TYPE_MAP)
	return ScMutableMap{objID: mapID}
}

func (o ScMutableMapArray) Immutable() ScImmutableMapArray {
	return ScImmutableMapArray(o)
}

func (o ScMutableMapArray) Length() int32 {
	return GetLength(o.objID)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScMutableRequestID struct {
	objID int32
	keyID Key32
}

func NewScMutableRequestID(objID int32, keyID Key32) ScMutableRequestID {
	return ScMutableRequestID{objID: objID, keyID: keyID}
}

func (o ScMutableRequestID) Exists() bool {
	return Exists(o.objID, o.keyID, TYPE_REQUEST_ID)
}

func (o ScMutableRequestID) SetValue(value ScRequestID) {
	SetBytes(o.objID, o.keyID, TYPE_REQUEST_ID, value.Bytes())
}

func (o ScMutableRequestID) String() string {
	return o.Value().String()
}

func (o ScMutableRequestID) Value() ScRequestID {
	return NewScRequestIDFromBytes(GetBytes(o.objID, o.keyID, TYPE_REQUEST_ID))
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScMutableRequestIDArray struct {
	objID int32
}

func (o ScMutableRequestIDArray) Clear() {
	Clear(o.objID)
}

func (o ScMutableRequestIDArray) GetRequestID(index int32) ScMutableRequestID {
	return ScMutableRequestID{objID: o.objID, keyID: Key32(index)}
}

func (o ScMutableRequestIDArray) Immutable() ScImmutableRequestIDArray {
	return ScImmutableRequestIDArray(o)
}

func (o ScMutableRequestIDArray) Length() int32 {
	return GetLength(o.objID)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScMutableString struct {
	objID int32
	keyID Key32
}

func NewScMutableString(objID int32, keyID Key32) ScMutableString {
	return ScMutableString{objID: objID, keyID: keyID}
}

func (o ScMutableString) Exists() bool {
	return Exists(o.objID, o.keyID, TYPE_STRING)
}

func (o ScMutableString) SetValue(value string) {
	SetBytes(o.objID, o.keyID, TYPE_STRING, []byte(value))
}

func (o ScMutableString) String() string {
	return o.Value()
}

func (o ScMutableString) Value() string {
	bytes := GetBytes(o.objID, o.keyID, TYPE_STRING)
	if bytes == nil {
		return ""
	}
	return string(bytes)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScMutableStringArray struct {
	objID int32
}

func (o ScMutableStringArray) Clear() {
	Clear(o.objID)
}

func (o ScMutableStringArray) GetString(index int32) ScMutableString {
	return ScMutableString{objID: o.objID, keyID: Key32(index)}
}

func (o ScMutableStringArray) Immutable() ScImmutableStringArray {
	return ScImmutableStringArray(o)
}

func (o ScMutableStringArray) Length() int32 {
	return GetLength(o.objID)
}
