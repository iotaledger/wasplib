package wasmproc

import (
	"github.com/iotaledger/wasp/packages/vm/wasmhost"
)

type ScSandboxObject struct {
	ScDict
	vm *wasmProcessor
}

func (o *ScSandboxObject) invalidKey(keyId int32) {
	o.Panic("Invalid key: %d", keyId)
}

func (o *ScSandboxObject) GetBytes(keyId int32, typeId int32) []byte {
	if keyId == wasmhost.KeyLength && (o.typeId&wasmhost.OBJTYPE_ARRAY) != 0 {
		return o.Int64Bytes(int64(o.length))
	}
	o.invalidKey(keyId)
	return nil
}

func (o *ScSandboxObject) GetObjectId(keyId int32, typeId int32) int32 {
	o.invalidKey(keyId)
	return 0
}

func (o *ScSandboxObject) SetBytes(keyId int32, typeId int32, bytes []byte) {
	o.invalidKey(keyId)
}
