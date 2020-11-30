// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package wasmhost

import (
	"fmt"
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

func (a *HostCall) call() {
	t := a.host
	request := t.FindSubObject(nil, "request", OBJTYPE_MAP)
	reqParams := t.FindSubObject(request, "params", OBJTYPE_MAP)

	scId := t.FindSubObject(nil, "contract", OBJTYPE_MAP).GetString(t.GetKeyId("id"))
	request.SetString(t.GetKeyId("sender"), scId)
	request.SetString(t.GetKeyId("function"), a.function)
	savedParams := NewHostMap(a.host, 0)
	reqParams.(*HostMap).CopyDataTo(savedParams)
	reqParams.SetInt(KeyLength, 0)
	params := t.FindSubObject(a, "params", OBJTYPE_MAP)
	params.(*HostMap).CopyDataTo(reqParams)
	fmt.Printf("    Call function: %v\n", a.function)
	err := t.RunScFunction(a.function)
	if err != nil {
		fmt.Printf("FAIL: Request function %s: %v\n", a.function, err)
		a.Error(err.Error())
	}
	reqParams.SetInt(KeyLength, 0)
	savedParams.CopyDataTo(reqParams)
}

func (a *HostCall) SetBytes(keyId int32, value []byte) {
	s := string(a.host.GetKeyFromId(keyId))
	//fmt.Printf("Call.SetBytes %s = %s\n", s, base58.Encode(value))
	a.HostMap.SetBytes(keyId, value)
	if s == "chain" {
		a.chain = value
		return
	}
}

func (a *HostCall) SetInt(keyId int32, value int64) {
	s := string(a.host.GetKeyFromId(keyId))
	//fmt.Printf("Call.SetInt %s = %d\n", s, value)
	a.HostMap.SetInt(keyId, value)
	if s != "delay" {
		return
	}
	if a.contract == "" {
		// call to self, immediately executed
		a.call()
		return
	}
	panic("Call.SetInt: call to other contract not implemented yet")
	//TODO take return values from json
}

func (a *HostCall) SetString(keyId int32, value string) {
	s := string(a.host.GetKeyFromId(keyId))
	//fmt.Printf("Call.SetString %s = %s\n", s, value)
	a.HostMap.SetString(keyId, value)
	if s == "contract" {
		a.contract = value
		return
	}
	if s == "function" {
		a.function = value
		return
	}
}
