package host

import (
	"github.com/iotaledger/wasplib/host/interfaces"
	"github.com/iotaledger/wasplib/host/interfaces/objtype"
)

type HostArray struct {
	ctx       *HostImpl
	items     []interface{}
	immutable bool
	typeId    int32
}

func NewHostArray(h *HostImpl, typeId int32) *HostArray {
	return &HostArray{ctx: h, typeId: typeId}
}

func (h *HostArray) GetBytes(keyId int32) []byte {
	if !h.valid(keyId, objtype.OBJTYPE_BYTES) {
		return []byte(nil)
	}
	return h.items[keyId].([]byte)
}

func (h *HostArray) GetInt(keyId int32) int64 {
	switch keyId {
	case interfaces.KeyLength:
		return int64(len(h.items))
	}

	if !h.valid(keyId, objtype.OBJTYPE_INT) {
		return 0
	}
	return h.items[keyId].(int64)
}

func (h *HostArray) GetLength() int32 {
	return int32(len(h.items))
}

func (h *HostArray) GetObjectId(keyId int32, typeId int32) int32 {
	if !h.valid(keyId, typeId) {
		return 0
	}
	return h.items[keyId].(int32)
}

func (h *HostArray) GetString(keyId int32) string {
	if !h.valid(keyId, objtype.OBJTYPE_STRING) {
		return ""
	}
	return h.items[keyId].(string)
}

func (h *HostArray) SetBytes(keyId int32, value []byte) {
	if EnableImmutableChecks && h.immutable {
		h.ctx.SetError("Array.SetBytes: Immutable")
		return
	}
	if !h.valid(keyId, objtype.OBJTYPE_BYTES) {
		return
	}
	h.items[keyId] = value
}

func (h *HostArray) SetInt(keyId int32, value int64) {
	if EnableImmutableChecks && h.immutable {
		h.ctx.SetError("Array.SetInt: Immutable")
		return
	}
	if keyId == interfaces.KeyLength {
		if h.typeId == objtype.OBJTYPE_MAP {
			// tell objects to clear themselves
			for i := len(h.items) - 1; i >= 0; i-- {
				h.ctx.SetInt(h.items[i].(int32), keyId, 0)
			}
			//TODO move to pool for reuse of transfers
		}
		h.items = nil
		return
	}
	if !h.valid(keyId, objtype.OBJTYPE_INT) {
		return
	}
	h.items[keyId] = value
}

func (h *HostArray) SetString(keyId int32, value string) {
	if EnableImmutableChecks && h.immutable {
		h.ctx.SetError("Array.SetString: Immutable")
		return
	}
	if !h.valid(keyId, objtype.OBJTYPE_STRING) {
		return
	}
	h.items[keyId] = value
}

func (h *HostArray) valid(keyId int32, typeId int32) bool {
	if h.typeId != typeId {
		h.ctx.SetError("Array.valid: Invalid access")
		return false
	}
	max := int32(len(h.items))
	if keyId == max && !h.immutable {
		switch typeId {
		case objtype.OBJTYPE_BYTES:
			h.items = append(h.items, []byte(nil))
		case objtype.OBJTYPE_INT:
			h.items = append(h.items, int64(0))
		case objtype.OBJTYPE_MAP:
			objId := h.ctx.AddObject(NewHostMap(h.ctx))
			h.items = append(h.items, objId)
		case objtype.OBJTYPE_STRING:
			h.items = append(h.items, "")
		default:
			h.ctx.SetError("Array.valid: Invalid typeId")
			return false
		}
		return true
	}
	if keyId < 0 || keyId >= max {
		h.ctx.SetError("Array.valid: Invalid index")
		return false
	}
	return true
}
