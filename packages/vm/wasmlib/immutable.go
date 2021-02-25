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

func (o ScImmutableAddress) Value() ScAddress {
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

func (o ScImmutableAgentId) Value() ScAgentId {
	return NewScAgentIdFromBytes(GetBytes(o.objId, o.keyId, TYPE_AGENT_ID))
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScImmutableAgentIdArray struct {
	objId int32
}

func (o ScImmutableAgentIdArray) GetAgentId(index int32) ScImmutableAgentId {
	return ScImmutableAgentId{objId: o.objId, keyId: Key32(index)}
}

func (o ScImmutableAgentIdArray) Length() int32 {
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

func (o ScImmutableChainId) Value() ScChainId {
	return NewScChainIdFromBytes(GetBytes(o.objId, o.keyId, TYPE_CHAIN_ID))
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScImmutableChainIdArray struct {
	objId int32
}

func (o ScImmutableChainIdArray) GetChainId(index int32) ScImmutableChainId {
	return ScImmutableChainId{objId: o.objId, keyId: Key32(index)}
}

func (o ScImmutableChainIdArray) Length() int32 {
	return GetLength(o.objId)
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

func (o ScImmutableColor) Value() ScColor {
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

func (o ScImmutableContractId) Value() ScContractId {
	return NewScContractIdFromBytes(GetBytes(o.objId, o.keyId, TYPE_CONTRACT_ID))
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScImmutableContractIdArray struct {
	objId int32
}

func (o ScImmutableContractIdArray) GetContractId(index int32) ScImmutableContractId {
	return ScImmutableContractId{objId: o.objId, keyId: Key32(index)}
}

func (o ScImmutableContractIdArray) Length() int32 {
	return GetLength(o.objId)
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

func (o ScImmutableHash) Value() ScHash {
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

type ScImmutableHnameArray struct {
	objId int32
}

func (o ScImmutableHnameArray) GetHname(index int32) ScImmutableHname {
	return ScImmutableHname{objId: o.objId, keyId: Key32(index)}
}

func (o ScImmutableHnameArray) Length() int32 {
	return GetLength(o.objId)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScImmutableInt64 struct {
	objId int32
	keyId Key32
}

func (o ScImmutableInt64) Exists() bool {
	return Exists(o.objId, o.keyId, TYPE_INT64)
}

func (o ScImmutableInt64) String() string {
	return strconv.FormatInt(o.Value(), 10)
}

func (o ScImmutableInt64) Value() int64 {
	bytes := GetBytes(o.objId, o.keyId, TYPE_INT64)
	return int64(binary.LittleEndian.Uint64(bytes))
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScImmutableInt64Array struct {
	objId int32
}

func (o ScImmutableInt64Array) GetInt64(index int32) ScImmutableInt64 {
	return ScImmutableInt64{objId: o.objId, keyId: Key32(index)}
}

func (o ScImmutableInt64Array) Length() int32 {
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

func (o ScImmutableMap) GetAgentIdArray(key MapKey) ScImmutableAgentIdArray {
	arrId := GetObjectId(o.objId, key.KeyId(), TYPE_AGENT_ID|TYPE_ARRAY)
	return ScImmutableAgentIdArray{objId: arrId}
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

func (o ScImmutableMap) GetChainIdArray(key MapKey) ScImmutableChainIdArray {
	arrId := GetObjectId(o.objId, key.KeyId(), TYPE_CHAIN_ID|TYPE_ARRAY)
	return ScImmutableChainIdArray{objId: arrId}
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

func (o ScImmutableMap) GetContractIdArray(key MapKey) ScImmutableContractIdArray {
	arrId := GetObjectId(o.objId, key.KeyId(), TYPE_CONTRACT_ID|TYPE_ARRAY)
	return ScImmutableContractIdArray{objId: arrId}
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

func (o ScImmutableMap) GetHnameArray(key MapKey) ScImmutableHnameArray {
	arrId := GetObjectId(o.objId, key.KeyId(), TYPE_HNAME|TYPE_ARRAY)
	return ScImmutableHnameArray{objId: arrId}
}

func (o ScImmutableMap) GetInt64(key MapKey) ScImmutableInt64 {
	return ScImmutableInt64{objId: o.objId, keyId: key.KeyId()}
}

func (o ScImmutableMap) GetInt64Array(key MapKey) ScImmutableInt64Array {
	arrId := GetObjectId(o.objId, key.KeyId(), TYPE_INT64|TYPE_ARRAY)
	return ScImmutableInt64Array{objId: arrId}
}

func (o ScImmutableMap) GetMap(key MapKey) ScImmutableMap {
	mapId := GetObjectId(o.objId, key.KeyId(), TYPE_MAP)
	return ScImmutableMap{objId: mapId}
}

func (o ScImmutableMap) GetMapArray(key MapKey) ScImmutableMapArray {
	arrId := GetObjectId(o.objId, key.KeyId(), TYPE_MAP|TYPE_ARRAY)
	return ScImmutableMapArray{objId: arrId}
}

func (o ScImmutableMap) GetRequestId(key MapKey) ScImmutableRequestId {
	return ScImmutableRequestId{objId: o.objId, keyId: key.KeyId()}
}

func (o ScImmutableMap) GetRequestIdArray(key MapKey) ScImmutableRequestIdArray {
	arrId := GetObjectId(o.objId, key.KeyId(), TYPE_REQUEST_ID|TYPE_ARRAY)
	return ScImmutableRequestIdArray{objId: arrId}
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

type ScImmutableRequestId struct {
	objId int32
	keyId Key32
}

func (o ScImmutableRequestId) Exists() bool {
	return Exists(o.objId, o.keyId, TYPE_REQUEST_ID)
}

func (o ScImmutableRequestId) String() string {
	return o.Value().String()
}

func (o ScImmutableRequestId) Value() ScRequestId {
	return NewScRequestIdFromBytes(GetBytes(o.objId, o.keyId, TYPE_REQUEST_ID))
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScImmutableRequestIdArray struct {
	objId int32
}

func (o ScImmutableRequestIdArray) GetRequestId(index int32) ScImmutableRequestId {
	return ScImmutableRequestId{objId: o.objId, keyId: Key32(index)}
}

func (o ScImmutableRequestIdArray) Length() int32 {
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
