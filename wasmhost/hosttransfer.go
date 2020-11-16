// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package wasmhost

import (
	"bytes"
	"github.com/mr-tron/base58"
)

type HostTransfer struct {
	HostMap
	address []byte
	color   []byte
}

func NewHostTransfer(host *SimpleWasmHost, keyId int32) *HostTransfer {
	return &HostTransfer{HostMap: *NewHostMap(host, keyId)}
}

func (a *HostTransfer) SetBytes(keyId int32, value []byte) {
	s := string(a.host.GetKey(keyId))
	//fmt.Printf("Transfer.SetBytes %s = %s\n", s, base58.Encode(value))
	a.HostMap.SetBytes(keyId, value)
	if s == "address" {
		a.address = value
		return
	}
	if s == "color" {
		a.color = value
		return
	}
}

func (a *HostTransfer) SetInt(keyId int32, value int64) {
	s := string(a.host.GetKey(keyId))
	//fmt.Printf("Transfer.SetInt %s = %d\n", s, value)
	a.HostMap.SetInt(keyId, value)
	if s != "amount" {
		return
	}
	account := a.host.FindSubObject(nil, "account", OBJTYPE_MAP)
	balance := a.host.FindSubObject(account, "balance", OBJTYPE_MAP)
	colorKeyId := a.host.GetKeyIdFromBytes([]byte(base58.Encode(a.color)))
	colorAmount := balance.GetInt(colorKeyId)
	if colorAmount < value {
		a.host.SetError("Insufficient funds")
		return
	}
	// check if compacting, in which case no balance change happens
	contract := a.host.FindSubObject(nil, "contract", OBJTYPE_MAP)
	if !bytes.Equal(a.address, contract.GetBytes(a.host.GetKeyId("address"))) {
		balance.SetInt(colorKeyId, colorAmount-value)
	}
}

func (a *HostTransfer) SetString(keyId int32, value string) {
	//s := string(a.host.GetKey(keyId))
	//fmt.Printf("Transfer.SetString %s = %s\n", s, value)
	a.HostMap.SetString(keyId, value)
}
