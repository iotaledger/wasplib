// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package wasmlocalhost

import "github.com/iotaledger/wasp/packages/vm/wasmhost"

type NullObject struct {
	SimpleObject
}

func NewNullObject(host *SimpleWasmHost) wasmhost.HostObject {
	return &NullObject{
		SimpleObject: SimpleObject{host: host},
	}
}

func (n *NullObject) Exists(keyId int32, typeId int32) bool {
	n.host.Error("Null.Exists")
	return false
}

func (n *NullObject) GetBytes(keyId int32, typeId int32) []byte {
	n.host.Error("Null.GetBytes")
	return nil
}

func (n *NullObject) GetObjectId(keyId int32, typeId int32) int32 {
	n.host.Error("Null.GetObjectId")
	return 0
}

func (n *NullObject) GetTypeId(keyId int32) int32 {
	n.host.Error("Null.GetTypeId")
	return -1
}

func (n *NullObject) SetBytes(keyId int32, typeId int32, value []byte) {
	n.host.Error("Null.SetBytes")
}
