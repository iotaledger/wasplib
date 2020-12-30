// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package wasmlocalhost

import (
	"bytes"
	"fmt"
	"github.com/iotaledger/wasp/packages/vm/wasmhost"
	"github.com/mr-tron/base58/base58"
	"io"
	"sort"
)

type HostMap struct {
	SimpleObject
	fields map[int32]interface{}
	types  map[int32]int32
}

func NewHostMap(host *SimpleWasmHost, keyId int32) *HostMap {
	return &HostMap{
		SimpleObject: SimpleObject{host: host, keyId: keyId},
		fields:       make(map[int32]interface{}),
		types:        make(map[int32]int32),
	}
}

func (m *HostMap) Dump(w io.Writer) {
	keys := make([]int32, len(m.fields))
	i := 0
	for k := range m.fields {
		keys[i] = k
		i++
	}
	sort.Slice(keys, func(i int, j int) bool {
		lhs := keys[i]
		rhs := keys[j]
		lhsFromString := (lhs & wasmhost.KeyFromString) != 0
		rhsFromString := (rhs & wasmhost.KeyFromString) != 0
		if lhsFromString != rhsFromString {
			// strings sort smaller than bytes
			return lhsFromString
		}
		lhsKey := m.host.GetKeyFromId(lhs)
		rhsKey := m.host.GetKeyFromId(rhs)
		return bytes.Compare(lhsKey, rhsKey) < 0
	})

	fmt.Fprintf(w, "{\n")
	multiple := false
	for _, keyId := range keys {
		value := m.fields[keyId]
		if multiple {
			fmt.Fprintf(w, ",\n")
		}
		multiple = true
		key := m.host.GetKeyFromId(keyId)
		fmt.Fprintf(w, "\"%s\": ", string(key))
		if keyId == wasmhost.KeyExports {
			m.host.FindObject(value.(int32)).(*HostExports).Dump(w)
			continue
		}
		m.host.Dump(w, m.types[keyId], value)
	}
	fmt.Fprintf(w, "}")
}

func (m *HostMap) Error(text string) {
	m.host.Error(text)
}

func (m *HostMap) Exists(keyId int32) bool {
	_, ok := m.fields[keyId]
	return ok
}

func (m *HostMap) GetBytes(keyId int32) []byte {
	value := m.GetString(keyId)
	if value == "" {
		return []byte(nil)
	}
	bytes, err := base58.Decode(value)
	if err != nil {
		m.Error("Map.GetBytes: " + err.Error())
		return []byte(nil)
	}
	return bytes
}

func (m *HostMap) GetInt(keyId int32) int64 {
	switch keyId {
	case wasmhost.KeyLength:
		return int64(len(m.fields))
	}

	if !m.valid(keyId, wasmhost.OBJTYPE_INT) {
		return 0
	}

	value, ok := m.fields[keyId]
	if !ok {
		return 0
	}
	return value.(int64)
}

func (m *HostMap) GetObjectId(keyId int32, typeId int32) int32 {
	if !m.valid(keyId, typeId) {
		return 0
	}
	value, ok := m.fields[keyId]
	if ok {
		return value.(int32)
	}

	var o VmObject
	switch typeId {
	case wasmhost.OBJTYPE_INT | wasmhost.OBJTYPE_ARRAY:
		o = NewHostArray(m.host, keyId, wasmhost.OBJTYPE_INT)
	case wasmhost.OBJTYPE_MAP:
		o = NewHostMap(m.host, keyId)
	case wasmhost.OBJTYPE_MAP | wasmhost.OBJTYPE_ARRAY:
		o = NewHostArray(m.host, keyId, wasmhost.OBJTYPE_MAP)
	default:
		if keyId == wasmhost.KeyExports {
			o = NewHostExports(m.host, keyId)
			break
		}
		if (typeId & wasmhost.OBJTYPE_ARRAY) != 0 {
			// all bytes types are treated as string
			o = NewHostArray(m.host, keyId, wasmhost.OBJTYPE_STRING)
			break
		}
		m.Error("Map.GetObjectId: Invalid type id")
		return 0
	}
	objId := m.host.TrackObject(o)
	o.InitObj(objId, m.id)
	m.fields[keyId] = objId
	return objId
}

func (m *HostMap) GetString(keyId int32) string {
	if !m.valid(keyId, wasmhost.OBJTYPE_STRING) {
		return ""
	}
	value, ok := m.fields[keyId]
	if !ok {
		return ""
	}
	return value.(string)
}

func (m *HostMap) GetTypeId(keyId int32) int32 {
	typeId, ok := m.types[keyId]
	if !ok {
		return -1
	}
	return typeId
}

func (m *HostMap) SetBytes(keyId int32, value []byte) {
	m.SetString(keyId, base58.Encode(value))
}

func (m *HostMap) SetInt(keyId int32, value int64) {
	if EnableImmutableChecks && m.immutable {
		m.Error("Map.SetInt: Immutable")
		return
	}
	if keyId == wasmhost.KeyLength {
		for fieldId, typeId := range m.types {
			if typeId == wasmhost.OBJTYPE_MAP || (typeId&wasmhost.OBJTYPE_ARRAY) != 0 {
				field, ok := m.fields[fieldId]
				if ok {
					// tell object to clear itself
					m.host.SetInt(field.(int32), keyId, 0)
				}
				//TODO move to pool for reuse of transfers
			}
		}
		m.fields = make(map[int32]interface{})
		return
	}
	if !m.valid(keyId, wasmhost.OBJTYPE_INT) {
		return
	}
	m.fields[keyId] = value
}

func (m *HostMap) SetString(keyId int32, value string) {
	if EnableImmutableChecks && m.immutable {
		m.Error("Map.SetString: Immutable")
		return
	}
	if !m.valid(keyId, wasmhost.OBJTYPE_STRING) {
		return
	}
	m.fields[keyId] = value
	if keyId == wasmhost.KeyPanic {
		m.host.panicked = true
		m.host.Error(value)
	}
}

func (m *HostMap) valid(keyId int32, typeId int32) bool {
	fieldType, ok := m.types[keyId]
	if !ok {
		if EnableImmutableChecks && m.immutable {
			m.Error("Map.valid: Immutable")
			return false
		}
		m.types[keyId] = typeId
		return true
	}
	if fieldType != typeId {
		m.Error("Map.valid: Invalid typeId")
		return false
	}
	return true
}

func (m *HostMap) CopyDataTo(other wasmhost.HostObject) {
	for k, v := range m.fields {
		switch m.types[k] {
		case wasmhost.OBJTYPE_BYTES:
			other.SetBytes(k, v.([]byte))
		case wasmhost.OBJTYPE_INT:
			other.SetInt(k, v.(int64))
		case wasmhost.OBJTYPE_STRING:
			other.SetString(k, v.(string))
		default:
			//TODO what about recursion?
			panic("Implement types")
		}
	}
}
