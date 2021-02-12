// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package wasmlib

import (
	"encoding/binary"
	"strconv"
)

type ScImmutableAddress struct {
	objId int32
	keyId Key32
}

func (o ScImmutableAddress) Exists() bool {
	return Exists(o.objId, o.keyId, TYPE_ADDRESS)
}

func (o ScImmutableAddress) String() string {
	return o.Value().String()
}

func (o ScImmutableAddress) Value() *ScAddress {
	return NewScAddressFromBytes(GetBytes(o.objId, o.keyId, TYPE_ADDRESS))
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

type ScImmutableAgentId struct {
	objId int32
	keyId Key32
}

func (o ScImmutableAgentId) Exists() bool {
	return Exists(o.objId, o.keyId, TYPE_AGENT_ID)
}

func (o ScImmutableAgentId) String() string {
	return o.Value().String()
}

func (o ScImmutableAgentId) Value() *ScAgentId {
	return NewScAgentIdFromBytes(GetBytes(o.objId, o.keyId, TYPE_AGENT_ID))
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScImmutableAgentArray struct {
	objId int32
}

func (o ScImmutableAgentArray) GetAgentId(index int32) ScImmutableAgentId {
	return ScImmutableAgentId{objId: o.objId, keyId: Key32(index)}
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
	return Exists(o.objId, o.keyId, TYPE_BYTES)
}

func (o ScImmutableBytes) String() string {
	return base58Encode(o.Value())
}

func (o ScImmutableBytes) Value() []byte {
	return GetBytes(o.objId, o.keyId, TYPE_BYTES)
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

type ScImmutableChainId struct {
	objId int32
	keyId Key32
}

func (o ScImmutableChainId) Exists() bool {
	return Exists(o.objId, o.keyId, TYPE_CHAIN_ID)
}

func (o ScImmutableChainId) String() string {
	return o.Value().String()
}

func (o ScImmutableChainId) Value() *ScChainId {
	return NewScChainIdFromBytes(GetBytes(o.objId, o.keyId, TYPE_CHAIN_ID))
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScImmutableColor struct {
	objId int32
	keyId Key32
}

func (o ScImmutableColor) Exists() bool {
	return Exists(o.objId, o.keyId, TYPE_COLOR)
}

func (o ScImmutableColor) String() string {
	return o.Value().String()
}

func (o ScImmutableColor) Value() *ScColor {
	return NewScColorFromBytes(GetBytes(o.objId, o.keyId, TYPE_COLOR))
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

type ScImmutableContractId struct {
	objId int32
	keyId Key32
}

func (o ScImmutableContractId) Exists() bool {
	return Exists(o.objId, o.keyId, TYPE_CONTRACT_ID)
}

func (o ScImmutableContractId) String() string {
	return o.Value().String()
}

func (o ScImmutableContractId) Value() *ScContractId {
	return NewScContractIdFromBytes(GetBytes(o.objId, o.keyId, TYPE_CONTRACT_ID))
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScImmutableHash struct {
	objId int32
	keyId Key32
}

func (o ScImmutableHash) Exists() bool {
	return Exists(o.objId, o.keyId, TYPE_HASH)
}

func (o ScImmutableHash) String() string {
	return o.Value().String()
}

func (o ScImmutableHash) Value() *ScHash {
	return NewScHashFromBytes(GetBytes(o.objId, o.keyId, TYPE_HASH))
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
	return Exists(o.objId, o.keyId, TYPE_HNAME)
}

func (o ScImmutableHname) String() string {
	return o.Value().String()
}

func (o ScImmutableHname) Value() ScHname {
	return NewScHnameFromBytes(GetBytes(o.objId, o.keyId, TYPE_HNAME))
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScImmutableInt struct {
	objId int32
	keyId Key32
}

func (o ScImmutableInt) Exists() bool {
	return Exists(o.objId, o.keyId, TYPE_INT)
}

func (o ScImmutableInt) String() string {
	return strconv.FormatInt(o.Value(), 10)
}

func (o ScImmutableInt) Value() int64 {
	bytes := GetBytes(o.objId, o.keyId, TYPE_INT)
	return int64(binary.LittleEndian.Uint64(bytes))
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

func (o ScImmutableMap) GetAgentId(key MapKey) ScImmutableAgentId {
	return ScImmutableAgentId{objId: o.objId, keyId: key.KeyId()}
}

func (o ScImmutableMap) GetAgentIdArray(key MapKey) ScImmutableAgentArray {
	arrId := GetObjectId(o.objId, key.KeyId(), TYPE_AGENT_ID|TYPE_ARRAY)
	return ScImmutableAgentArray{objId: arrId}
}

func (o ScImmutableMap) GetBytes(key MapKey) ScImmutableBytes {
	return ScImmutableBytes{objId: o.objId, keyId: key.KeyId()}
}

func (o ScImmutableMap) GetBytesArray(key MapKey) ScImmutableBytesArray {
	arrId := GetObjectId(o.objId, key.KeyId(), TYPE_BYTES|TYPE_ARRAY)
	return ScImmutableBytesArray{objId: arrId}
}

func (o ScImmutableMap) GetChainId(key MapKey) ScImmutableChainId {
	return ScImmutableChainId{objId: o.objId, keyId: key.KeyId()}
}

func (o ScImmutableMap) GetColor(key MapKey) ScImmutableColor {
	return ScImmutableColor{objId: o.objId, keyId: key.KeyId()}
}

func (o ScImmutableMap) GetColorArray(key MapKey) ScImmutableColorArray {
	arrId := GetObjectId(o.objId, key.KeyId(), TYPE_COLOR|TYPE_ARRAY)
	return ScImmutableColorArray{objId: arrId}
}

func (o ScImmutableMap) GetContractId(key MapKey) ScImmutableContractId {
	return ScImmutableContractId{objId: o.objId, keyId: key.KeyId()}
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
	return Exists(o.objId, o.keyId, TYPE_STRING)
}

func (o ScImmutableString) String() string {
	return o.Value()
}

func (o ScImmutableString) Value() string {
	bytes := GetBytes(o.objId, o.keyId, TYPE_STRING)
	if bytes == nil {
		return ""
	}
	return string(bytes)
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
