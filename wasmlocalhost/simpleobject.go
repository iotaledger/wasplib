package wasmlocalhost

import (
	"fmt"
	"github.com/iotaledger/wasp/packages/vm/wasmhost"
)

type VmObject interface {
	wasmhost.HostObject
	InitObj(id int32, ownerId int32)
	Fail(format string, args ...interface{}) bool
	//FindOrMakeObjectId(keyId int32, factory ObjFactory) int32
	Name() string
	Suffix(keyId int32) string
}

type SimpleObject struct {
	host      *SimpleWasmHost
	id        int32
	keyId     int32
	ownerId   int32
	immutable bool
}

func (o *SimpleObject) Fail(format string, args ...interface{}) bool {
	fmt.Printf("FAIL: %s: %s\n", o.Name(), fmt.Sprintf(format, args...))
	return false
}

func (o *SimpleObject) InitObj(id int32, ownerId int32) {
	o.id = id
	o.ownerId = ownerId
}

func (o *SimpleObject) Name() string {
	switch o.id {
	case 0:
		return "null"
	case 1:
		return "root"
	default:
		owner := o.host.FindObject(o.ownerId).(VmObject)
		if o.ownerId == 1 {
			// root sub object, skip the "root." prefix
			return o.host.GetKeyStringFromId(o.keyId)
		}
		return owner.Name() + owner.Suffix(o.keyId)
	}
}

func (o *SimpleObject) Suffix(keyId int32) string {
	return "." + o.host.GetKeyStringFromId(keyId)
}
