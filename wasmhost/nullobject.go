package wasmhost

type NullObject struct {
	host *WasmHost
}

func NewNullObject(host *WasmHost) HostObject {
	return &NullObject{host: host}
}

func (n *NullObject) GetBytes(keyId int32) []byte {
	n.host.SetError("Null.GetBytes")
	return nil
}

func (n *NullObject) GetInt(keyId int32) int64 {
	n.host.SetError("Null.GetInt")
	return 0
}

func (n *NullObject) GetString(keyId int32) string {
	n.host.SetError("Null.GetString")
	return ""
}

func (n *NullObject) GetObjectId(keyId int32, typeId int32) int32 {
	n.host.SetError("Null.GetObjectId")
	return 0
}

func (n *NullObject) SetBytes(keyId int32, value []byte) {
	n.host.SetError("Null.SetBytes")
}

func (n *NullObject) SetInt(keyId int32, value int64) {
	n.host.SetError("Null.SetInt")
}

func (n *NullObject) SetString(keyId int32, value string) {
	n.host.SetError("Null.SetString")
}
