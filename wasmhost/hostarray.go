package wasmhost

import "github.com/mr-tron/base58"

type HostArray struct {
	host      *SimpleWasmHost
	items     []interface{}
	immutable bool
	keyId     int32
	typeId    int32
}

func NewHostArray(host *SimpleWasmHost, keyId int32, typeId int32) *HostArray {
	return &HostArray{host: host, keyId: keyId, typeId: typeId}
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
		a.host.SetError("Map.GetBytes: " + err.Error())
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

func (a *HostArray) SetBytes(keyId int32, value []byte) {
	a.SetString(keyId, base58.Encode(value))
}

func (a *HostArray) SetInt(keyId int32, value int64) {
	if EnableImmutableChecks && a.immutable {
		a.host.SetError("Array.SetInt: Immutable")
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
		a.host.SetError("Array.SetString: Immutable")
		return
	}
	if !a.valid(keyId, OBJTYPE_STRING) {
		return
	}
	a.items[keyId] = value
}

func (a *HostArray) valid(keyId int32, typeId int32) bool {
	if a.typeId != typeId {
		a.host.SetError("Array.valid: Invalid access")
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
			if a.keyId == a.host.TransfersId {
				objId := a.host.TrackObject(NewHostTransfer(a.host, keyId))
				a.items = append(a.items, objId)
				break
			}
			objId := a.host.TrackObject(NewHostMap(a.host, keyId))
			a.items = append(a.items, objId)
		case OBJTYPE_STRING:
			a.items = append(a.items, "")
		default:
			a.host.SetError("Array.valid: Invalid typeId")
			return false
		}
		return true
	}
	if keyId < 0 || keyId >= max {
		a.host.SetError("Array.valid: Invalid index")
		return false
	}
	return true
}
