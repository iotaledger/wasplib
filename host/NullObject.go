package host

import "github.com/iotaledger/wasplib/host/interfaces"

type NullObject struct {
	ctx interfaces.HostInterface
}

func NewNullObject(h interfaces.HostInterface) interfaces.HostObject {
	return &NullObject{ctx: h}
}

func (n *NullObject) error(text string) {
	n.ctx.SetBytes(1, interfaces.KeyError, []byte(text))
}

func (n *NullObject) GetBytes(keyId int32) []byte {
	n.error("Null.GetBytes")
	return nil
}

func (n *NullObject) GetInt(keyId int32) int64 {
	n.error("Null.GetInt")
	return 0
}

func (n *NullObject) GetString(keyId int32) string {
	n.error("Null.GetString")
	return ""
}

func (n *NullObject) GetObjectId(keyId int32, typeId int32) int32 {
	n.error("Null.GetObjectId")
	return 0
}

func (n *NullObject) SetBytes(keyId int32, value []byte) {
	n.error("Null.SetBytes")
}

func (n *NullObject) SetInt(keyId int32, value int64) {
	n.error("Null.SetInt")
}

func (n *NullObject) SetString(keyId int32, value string) {
	n.error("Null.SetString")
}
