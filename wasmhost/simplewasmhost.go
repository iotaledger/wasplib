// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package wasmhost

import (
	"fmt"
	"github.com/mr-tron/base58"
	"hash/fnv"
	"io"
)

var EnableImmutableChecks = true

type SimpleWasmHost struct {
	WasmHost
	CallsId     int32
	ExportsId   int32
	TransfersId int32
}

func NewSimpleWasmHost() (*SimpleWasmHost, error) {
	host := &SimpleWasmHost{}
	host.useBase58Keys = true
	err := host.Init(NewNullObject(host), NewHostMap(host, 0), nil, host)
	if err != nil {
		return nil, err
	}
	host.CallsId = host.GetKeyId("calls")
	host.ExportsId = host.GetKeyId("exports")
	host.TransfersId = host.GetKeyId("transfers")
	return host, nil
}

func (host *SimpleWasmHost) Dump(w io.Writer, typeId int32, value interface{}) {
	switch typeId {
	case OBJTYPE_BYTES:
		fmt.Fprintf(w, "\"%s\"", base58.Encode(value.([]byte)))
	case OBJTYPE_INT:
		fmt.Fprintf(w, "%d", value.(int64))
	case OBJTYPE_MAP:
		host.FindObject(value.(int32)).(*HostMap).Dump(w)
	case OBJTYPE_STRING:
		fmt.Fprintf(w, "\"%s\"", value.(string))
	case OBJTYPE_BYTES_ARRAY, OBJTYPE_INT_ARRAY, OBJTYPE_MAP_ARRAY, OBJTYPE_STRING_ARRAY:
		host.FindObject(value.(int32)).(*HostArray).Dump(w)
	}
}

func (host *SimpleWasmHost) Log(logLevel int32, text string) {
	switch logLevel {
	case KeyTraceHost:
		//fmt.Println(text)
	case KeyTrace:
		//fmt.Println(text)
	case KeyLog:
		fmt.Println(text)
	case KeyWarning:
		fmt.Println(text)
	case KeyError:
		fmt.Println(text)
	}
}

//func (host *SimpleWasmHost) RunScFunction(functionName string) error {
//	fmt.Printf("Simple function: %v\n", functionName)
//	index, ok := host.funcToIndex[functionName]
//	if !ok {
//		return errors.New("unknown SC function name: " + functionName)
//	}
//
//	client.ScCallEntrypoint(index)
//	return nil
//}

func (host *SimpleWasmHost) SetExport(index int32, functionName string) {
	_, ok := host.funcToCode[functionName]
	if ok {
		host.SetError("SetExport: duplicate function name")
		return
	}
	h := fnv.New32a()
	h.Write([]byte(functionName))
	hashedName := h.Sum32()
	_, ok = host.codeToFunc[hashedName]
	if ok {
		host.SetError("SetExport: duplicate hashed name")
		return
	}
	host.codeToFunc[hashedName] = functionName
	host.funcToCode[functionName] = hashedName
	host.funcToIndex[functionName] = index
}
