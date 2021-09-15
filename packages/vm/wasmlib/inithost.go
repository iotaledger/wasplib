// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package wasmlib

type InitHost struct {
	params map[int32][]byte
}

var _ ScHost = &InitHost{}

func NewInitHost() *InitHost {
	return &InitHost{params: make(map[int32][]byte)}
}

func (h InitHost) CallFunc(objID, keyID int32, params []byte) []byte {
	Panic("InitHost::CallFunc")
	return nil
}

func (h InitHost) Exists(objID, keyID, typeID int32) bool {
	Panic("InitHost::Exists")
	return false
}

func (h InitHost) GetBytes(objID, keyID, typeID int32) []byte {
	if objID == int32(KeyMaps) && keyID == int32(KeyLength) {
		return nil
	}
	Panic("InitHost::GetBytes")
	return nil
}

func (h InitHost) GetKeyIDFromBytes(bytes []byte) int32 {
	Panic("InitHost::GetKeyIDFromBytes")
	return 0
}

func (h InitHost) GetKeyIDFromString(key string) int32 {
	Panic("InitHost::GetKeyIDFromString")
	return 0
}

func (h InitHost) GetObjectID(objID, keyID, typeID int32) int32 {
	if objID == 1 && keyID == int32(KeyMaps) {
		return keyID
	}
	if objID == int32(KeyMaps) && keyID == 0 {
		return int32(KeyParams)
	}
	Panic("InitHost::GetObjectID")
	return 0
}

func (h InitHost) SetBytes(objID, keyID, typeID int32, value []byte) {
	if objID == int32(KeyParams) {
		h.params[keyID] = value
		return
	}
	Panic("InitHost::SetBytes")
}
