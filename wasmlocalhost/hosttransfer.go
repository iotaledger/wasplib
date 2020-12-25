// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package wasmlocalhost

import (
	"bytes"
	"github.com/iotaledger/wasp/packages/vm/wasmhost"
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
	if keyId == wasmhost.KeyAgent {
		a.agent = value
		return
	}
}

func (a *HostTransfer) SetInt(keyId int32, value int64) {
	a.HostMap.SetInt(keyId, value)
	if keyId == wasmhost.KeyLength {
		return
	}

	balances := a.host.FindSubObject(nil, wasmhost.KeyBalances, wasmhost.OBJTYPE_MAP)
	colorAmount := balances.GetInt(keyId)
	if colorAmount < value {
		a.Error("Insufficient funds")
		return
	}
	// check if compacting, in which case no balance change happens
	contract := a.host.FindSubObject(nil, wasmhost.KeyContract, wasmhost.OBJTYPE_MAP)
	scId := contract.GetBytes(wasmhost.KeyId)
	if !bytes.Equal(a.agent, scId) {
		balances.SetInt(keyId, colorAmount-value)
	}
}
