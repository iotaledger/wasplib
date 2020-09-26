package host

import (
	"github.com/iotaledger/wasplib/host/interfaces"
	"github.com/iotaledger/wasplib/host/interfaces/objtype"
)

var EnableImmutableChecks = true

type HostImpl struct {
	HostBase
}

func NewHostImpl() *HostImpl {
	h := &HostImpl{}
	h.Init(h, NewHostMap(h), nil)
	return h
}

func (h *HostImpl) AddBalance(obj interfaces.HostObject, color string, amount int64) {
	colors := h.Object(obj, "colors", objtype.OBJTYPE_STRING_ARRAY)
	length := colors.GetInt(interfaces.KeyLength)
	colors.SetString(int32(length), color)
	colorId := h.GetKeyId(color)
	balance := h.Object(obj, "balance", objtype.OBJTYPE_MAP)
	balance.SetInt(colorId, amount)
}

func (h *HostImpl) Object(obj interfaces.HostObject, key string, typeId int32) interfaces.HostObject {
	if obj == nil {
		// use root object
		obj = h.GetObject(1)
	}
	return h.GetObject(obj.GetObjectId(h.GetKeyId(key), typeId))
}
