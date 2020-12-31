// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package wasmlocalhost

import (
	"bytes"
	"github.com/iotaledger/wasp/packages/vm/wasmhost"
	"github.com/iotaledger/wasplib/client"
)

type HostTransfer struct {
	HostMap
	agent []byte
}

func NewHostTransfer(host *SimpleWasmHost, keyId int32) *HostTransfer {
	return &HostTransfer{HostMap: *NewHostMap(host, keyId)}
}

func (m *HostTransfer) SetBytes(keyId int32, value []byte) {
	m.HostMap.SetBytes(keyId, value)
	if keyId == wasmhost.KeyAgent {
		m.agent = value
		return
	}
}

func (m *HostTransfer) SetInt(keyId int32, value int64) {
	m.HostMap.SetInt(keyId, value)
	if keyId == wasmhost.KeyLength {
		return
	}

	balances := m.host.FindSubObject(nil, wasmhost.KeyBalances, client.TYPE_MAP)
	colorAmount := balances.GetInt(keyId)
	if colorAmount < value {
		m.Error("Insufficient funds")
		return
	}
	// check if compacting, in which case no balance change happens
	contract := m.host.FindSubObject(nil, wasmhost.KeyContract, client.TYPE_MAP)
	scId := contract.GetBytes(wasmhost.KeyId)
	if !bytes.Equal(m.agent, scId) {
		balances.SetInt(keyId, colorAmount-value)
	}
}
