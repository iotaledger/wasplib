// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package wasmhost

import (
	"bytes"
)

type HostTransfer struct {
	HostMap
	agent []byte
}

func NewHostTransfer(host *SimpleWasmHost, keyId int32) *HostTransfer {
	return &HostTransfer{HostMap: *NewHostMap(host, keyId)}
}

func (a *HostTransfer) SetBytes(keyId int32, value []byte) {
	a.HostMap.SetBytes(keyId, value)
	if keyId == KeyAgent {
		a.agent = value
		return
	}
}

func (a *HostTransfer) SetInt(keyId int32, value int64) {
	a.HostMap.SetInt(keyId, value)
	if keyId == KeyLength {
		return
	}

	balances := a.host.FindSubObject(nil, "balances", OBJTYPE_MAP)
	colorAmount := balances.GetInt(keyId)
	if colorAmount < value {
		a.Error("Insufficient funds")
		return
	}
	// check if compacting, in which case no balance change happens
	contract := a.host.FindSubObject(nil, "contract", OBJTYPE_MAP)
	scId := contract.GetBytes(KeyId)
	if !bytes.Equal(a.agent, scId) {
		balances.SetInt(keyId, colorAmount-value)
	}
}
