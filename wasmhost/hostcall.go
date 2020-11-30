// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package wasmhost

import (
	"fmt"
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

func (a *HostCall) call() {
	host := a.host
	request := host.FindSubObject(nil, "request", OBJTYPE_MAP)
	reqParams := host.FindSubObject(request, "params", OBJTYPE_MAP)

	scId := host.FindSubObject(nil, "contract", OBJTYPE_MAP).GetString(host.GetKeyId("id"))
	request.SetString(host.GetKeyId("sender"), scId)
	request.SetString(host.GetKeyId("function"), a.function)
	savedParams := NewHostMap(a.host, 0)
	reqParams.(*HostMap).CopyDataTo(savedParams)
	reqParams.SetInt(KeyLength, 0)
	params := host.FindSubObject(a, "params", OBJTYPE_MAP)
	params.(*HostMap).CopyDataTo(reqParams)
	fmt.Printf("    Call function: %v\n", a.function)
	err := host.RunScFunction(a.function)
	if err != nil {
		fmt.Printf("FAIL: Request function %s: %v\n", a.function, err)
		a.Error(err.Error())
	}
	reqParams.SetInt(KeyLength, 0)
	savedParams.CopyDataTo(reqParams)
}

func (a *HostCall) SetBytes(keyId int32, value []byte) {
	key := string(a.host.GetKeyFromId(keyId))
	a.host.TraceHost("Call.SetBytes %s = %s", key, base58.Encode(value))
	a.HostMap.SetBytes(keyId, value)
	if key == "chain" {
		a.chain = value
		return
	}
}

func (a *HostCall) SetInt(keyId int32, value int64) {
	key := string(a.host.GetKeyFromId(keyId))
	a.host.TraceHost("Call.SetInt %s = %d\n", key, value)
	a.HostMap.SetInt(keyId, value)
	if key != "delay" {
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
	key := string(a.host.GetKeyFromId(keyId))
	a.host.TraceHost("Call.SetString %s = %s\n", key, value)
	a.HostMap.SetString(keyId, value)
	if key == "contract" {
		a.contract = value
		return
	}
	if key == "function" {
		a.function = value
		return
	}
}
