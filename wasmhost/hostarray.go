// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package wasmhost

import (
	"fmt"
	"github.com/mr-tron/base58"
	"io"
)

type HostArray struct {
	SimpleObject
	items  []interface{}
	typeId int32
}

func NewHostArray(host *SimpleWasmHost, keyId int32, typeId int32) *HostArray {
	return &HostArray{
		SimpleObject: SimpleObject{host: host, keyId: keyId},
		typeId:       typeId,
	}
}

func (a *HostArray) Dump(w io.Writer) {
	fmt.Fprintf(w, "[\n")
	multiple := false
	for _, item := range a.items {
		if multiple {
			fmt.Fprintf(w, ",\n")
		}
		multiple = true
		if a.keyId == KeyCalls {
			a.host.FindObject(item.(int32)).(*HostCall).Dump(w)
			continue
		}
		if a.keyId == KeyTransfers {
			a.host.FindObject(item.(int32)).(*HostTransfer).Dump(w)
			continue
		}
		a.host.Dump(w, a.typeId, item)
	}
	fmt.Fprintf(w, "]")
}

func (a *HostArray) Error(text string) {
	a.host.SetError(text)
}

func (a *HostArray) Exists(keyId int32) bool {
	return keyId >= 0 && keyId < int32(len(a.items))
}

func (a *HostArray) GetBytes(keyId int32) []byte {
	value := a.GetString(keyId)
	if value == "" {
		return []byte(nil)
	}
	bytes, err := base58.Decode(value)
	if err != nil {
		a.Error("Map.GetBytes: " + err.Error())
		return []byte(nil)
	}
	return bytes
}

func (a *HostArray) GetInt(keyId int32) int64 {
	switch keyId {
	case KeyLength:
		return int64(len(a.items))
	}

	if !a.valid(keyId, OBJTYPE_INT) {
		return 0
	}
	return a.items[keyId].(int64)
}

func (a *HostArray) GetLength() int32 {
	return int32(len(a.items))
}

func (a *HostArray) GetObjectId(keyId int32, typeId int32) int32 {
	if !a.valid(keyId, typeId) {
		return 0
	}
	return a.items[keyId].(int32)
}

func (a *HostArray) GetString(keyId int32) string {
	if !a.valid(keyId, OBJTYPE_STRING) {
		return ""
	}
	return a.items[keyId].(string)
}

func (a *HostArray) GetTypeId(keyId int32) int32 {
	return a.typeId
}

func (a *HostArray) SetBytes(keyId int32, value []byte) {
	a.SetString(keyId, base58.Encode(value))
}

func (a *HostArray) SetInt(keyId int32, value int64) {
	if EnableImmutableChecks && a.immutable {
		a.Error("Array.SetInt: Immutable")
		return
	}
	if keyId == KeyLength {
		if a.typeId == OBJTYPE_MAP {
			// tell objects to clear themselves
			for i := len(a.items) - 1; i >= 0; i-- {
				a.host.SetInt(a.items[i].(int32), keyId, 0)
			}
			//TODO move to pool for reuse of transfers
		}
		a.items = nil
		return
	}
	if !a.valid(keyId, OBJTYPE_INT) {
		return
	}
	a.items[keyId] = value
}

func (a *HostArray) SetString(keyId int32, value string) {
	if EnableImmutableChecks && a.immutable {
		a.Error("Array.SetString: Immutable")
		return
	}
	if !a.valid(keyId, OBJTYPE_STRING) {
		return
	}
	a.items[keyId] = value
}

func (a *HostArray) Suffix(keyId int32) string {
	return fmt.Sprintf("[%d]", keyId)
}

func (a *HostArray) valid(keyId int32, typeId int32) bool {
	if a.typeId != typeId {
		a.Error("Array.valid: Invalid access")
		return false
	}
	max := int32(len(a.items))
	if keyId == max && !a.immutable {
		switch typeId {
		case OBJTYPE_BYTES:
			a.items = append(a.items, []byte(nil))
		case OBJTYPE_INT:
			a.items = append(a.items, int64(0))
		case OBJTYPE_MAP:
			var o VmObject
			switch a.keyId {
			case KeyCalls:
				o = NewHostCall(a.host, keyId)
			case KeyTransfers:
				o = NewHostTransfer(a.host, keyId)
			default:
				o = NewHostMap(a.host, keyId)
			}
			objId := a.host.TrackObject(o)
			a.items = append(a.items, objId)
			o.InitObj(objId, a.id)
		case OBJTYPE_STRING:
			a.items = append(a.items, "")
		default:
			a.Error("Array.valid: Invalid typeId")
			return false
		}
		return true
	}
	if keyId < 0 || keyId >= max {
		a.Error("Array.valid: Invalid index")
		return false
	}
	return true
}
