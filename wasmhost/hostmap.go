// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package wasmhost

import (
	"fmt"
	"github.com/mr-tron/base58/base58"
	"io"
)

type HostMap struct {
	host      *SimpleWasmHost
	fields    map[int32]interface{}
	immutable bool
	keyId     int32
	types     map[int32]int32
}

func NewHostMap(host *SimpleWasmHost, keyId int32) *HostMap {
	return &HostMap{
		host:   host,
		keyId:  keyId,
		fields: make(map[int32]interface{}),
		types:  make(map[int32]int32),
	}
}

func (m *HostMap) Dump(w io.Writer) {
	fmt.Fprintf(w, "{\n")
	multiple := false
	for keyId, value := range m.fields {
		if multiple {
			fmt.Fprintf(w, ",\n")
		}
		multiple = true
		key := m.host.GetKey(keyId)
		fmt.Fprintf(w, "\"%s\": ", string(key))
		if keyId == m.host.ExportsId {
			m.host.FindObject(value.(int32)).(*HostExports).Dump(w)
			continue
		}
		m.host.Dump(w, m.types[keyId], value)
	}
	fmt.Fprintf(w, "}")
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
		m.host.SetError("Map.GetBytes: " + err.Error())
		return []byte(nil)
	}
	return bytes
}

func (m *HostMap) GetInt(keyId int32) int64 {
	switch keyId {
	case KeyLength:
		return int64(len(m.fields))
	}

	if !m.valid(keyId, OBJTYPE_INT) {
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

	var o HostObject
	switch typeId {
	case OBJTYPE_BYTES_ARRAY:
		o = NewHostArray(m.host, keyId, OBJTYPE_STRING)
	case OBJTYPE_INT_ARRAY:
		o = NewHostArray(m.host, keyId, OBJTYPE_INT)
	case OBJTYPE_MAP:
		o = NewHostMap(m.host, keyId)
	case OBJTYPE_MAP_ARRAY:
		o = NewHostArray(m.host, keyId, OBJTYPE_MAP)
	case OBJTYPE_STRING_ARRAY:
		if keyId == m.host.ExportsId {
			o = NewHostExports(m.host, keyId)
			break
		}
		o = NewHostArray(m.host, keyId, OBJTYPE_STRING)
	default:
		m.host.SetError("Map.GetObjectId: Invalid type id")
		return 0
	}
	objId := m.host.TrackObject(o)
	m.fields[keyId] = objId
	return objId
}

func (m *HostMap) GetString(keyId int32) string {
	if !m.valid(keyId, OBJTYPE_STRING) {
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
		m.host.SetError("Map.SetInt: Immutable")
		return
	}
	if keyId == KeyLength {
		for k, v := range m.types {
			switch v {
			case OBJTYPE_MAP,
				OBJTYPE_BYTES_ARRAY,
				OBJTYPE_INT_ARRAY,
				OBJTYPE_STRING_ARRAY:
				field, ok := m.fields[k]
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
	if !m.valid(keyId, OBJTYPE_INT) {
		return
	}
	m.fields[keyId] = value
}

func (m *HostMap) SetString(keyId int32, value string) {
	if EnableImmutableChecks && m.immutable {
		m.host.SetError("Map.SetString: Immutable")
		return
	}
	if !m.valid(keyId, OBJTYPE_STRING) {
		return
	}
	m.fields[keyId] = value
}

func (m *HostMap) valid(keyId int32, typeId int32) bool {
	fieldType, ok := m.types[keyId]
	if !ok {
		if EnableImmutableChecks && m.immutable {
			m.host.SetError("Map.valid: Immutable")
			return false
		}
		m.types[keyId] = typeId
		return true
	}
	if fieldType != typeId {
		m.host.SetError("Map.valid: Invalid typeId")
		return false
	}
	return true
}

func (m *HostMap) CopyDataTo(other HostObject) {
	for k, v := range m.fields {
		switch m.types[k] {
		case OBJTYPE_BYTES:
			other.SetBytes(k, v.([]byte))
		case OBJTYPE_INT:
			other.SetInt(k, v.(int64))
		case OBJTYPE_STRING:
			other.SetString(k, v.(string))
		default:
			//TODO what about recursion?
			panic("Implement types")
		}
	}
}
