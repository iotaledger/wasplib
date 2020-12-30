// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package wasmlocalhost

import (
	"fmt"
	"github.com/iotaledger/wasp/packages/vm/wasmhost"
	"github.com/mr-tron/base58"
)

type HostCall struct {
	HostMap
	chain    []byte
	contract string
	function string
	delay    int64
}

func NewHostCall(host *SimpleWasmHost, keyId int32) *HostCall {
	return &HostCall{HostMap: *NewHostMap(host, keyId)}
}

func (m *HostCall) call() {
	host := m.host

	root := host.FindObject(1)
	savedCaller := root.GetString(wasmhost.KeyCaller)
	scId := host.FindSubObject(nil, wasmhost.KeyContract, wasmhost.OBJTYPE_MAP).GetString(wasmhost.KeyId)
	root.SetString(wasmhost.KeyCaller, scId)

	requestParams := host.FindSubObject(nil, wasmhost.KeyParams, wasmhost.OBJTYPE_MAP)
	savedParams := NewHostMap(m.host, 0)
	requestParams.(*HostMap).CopyDataTo(savedParams)
	requestParams.SetInt(wasmhost.KeyLength, 0)
	params := host.FindSubObject(m, wasmhost.KeyParams, wasmhost.OBJTYPE_MAP)
	params.(*HostMap).CopyDataTo(requestParams)

	fmt.Printf("    Call function: %v\n", m.function)
	err := host.RunScFunction(m.function)
	if err != nil {
		fmt.Printf("FAIL: Request function %s: %v\n", m.function, err)
		m.Error(err.Error())
	}

	requestParams.SetInt(wasmhost.KeyLength, 0)
	savedParams.CopyDataTo(requestParams)
	root.SetString(wasmhost.KeyCaller, savedCaller)
}

func (m *HostCall) SetBytes(keyId int32, value []byte) {
	key := string(m.host.GetKeyFromId(keyId))
	m.host.TraceAll("Call.SetBytes %s = %s", key, base58.Encode(value))
	m.HostMap.SetBytes(keyId, value)
	if key == "chain" {
		m.chain = value
		return
	}
}

func (m *HostCall) SetInt(keyId int32, value int64) {
	key := string(m.host.GetKeyFromId(keyId))
	m.host.TraceAll("Call.SetInt %s = %d\n", key, value)
	m.HostMap.SetInt(keyId, value)
	if key != "delay" {
		return
	}
	if m.contract == "" {
		// call to self, immediately executed
		m.call()
		return
	}
	panic("Call.SetInt: call to other contract not implemented yet")
	//TODO take return values from json
}

func (m *HostCall) SetString(keyId int32, value string) {
	key := string(m.host.GetKeyFromId(keyId))
	m.host.TraceAll("Call.SetString %s = %s\n", key, value)
	m.HostMap.SetString(keyId, value)
	if key == "contract" {
		m.contract = value
		return
	}
	if key == "function" {
		m.function = value
		return
	}
}
