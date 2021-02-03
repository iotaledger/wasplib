// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package wasmlocalhost

import (
	"bytes"
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/balance"
	"github.com/iotaledger/wasp/packages/vm/wasmhost"
	"github.com/iotaledger/wasplib/client"
)

type HostTransfer struct {
	HostMap
	agent []byte
	chain []byte
}

func NewHostTransfer(host *SimpleWasmHost, keyId int32) *HostTransfer {
	return &HostTransfer{HostMap: *NewHostMap(host, keyId)}
}

func (m *HostTransfer) SetBytes(keyId int32, typeId int32, value []byte) {
	m.HostMap.SetBytes(keyId, typeId, value)
}

func (m *HostTransfer) SetInt(keyId int32, value int64) {
	m.HostMap.SetInt(keyId, value)
	if keyId == wasmhost.KeyLength {
		return
	}

	if keyId == m.host.GetKeyIdFromBytes(balance.ColorNew[:]) && value == -1 {
		return
	}

	balances := m.host.FindSubObject(nil, wasmhost.KeyBalances, client.TYPE_MAP)
	colorAmount := BytesToInt(balances.GetBytes(keyId, client.TYPE_INT))
	if colorAmount < value {
		m.Error("Insufficient funds")
		return
	}
	// check if compacting, in which case no balance change happens
	root := m.host.FindObject(1)
	scId := root.GetBytes(wasmhost.KeyContractId, client.TYPE_CONTRACT_ID)
	if !bytes.Equal(m.agent, scId) {
		balances.SetBytes(keyId, client.TYPE_INT, IntToBytes(colorAmount-value))
	}
}
