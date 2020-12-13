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
}

func NewSimpleWasmHost(vm WasmVM) (*SimpleWasmHost, error) {
	host := &SimpleWasmHost{}
	host.vm = vm
	host.useBase58Keys = true
	root := NewHostMap(host, 0)
	root.InitObj(1, 0)
	host.Init(NewNullObject(host), root, host)
	err := host.InitVM(vm)
	if err != nil {
		return nil, err
	}
	return host, nil
}

func (host *SimpleWasmHost) Dump(w io.Writer, typeId int32, value interface{}) {
	switch typeId {
	case OBJTYPE_ADDRESS,
		OBJTYPE_AGENT,
		OBJTYPE_BYTES,
		OBJTYPE_COLOR:
		fmt.Fprintf(w, "\"%s\"", base58.Encode(value.([]byte)))
	case OBJTYPE_INT:
		fmt.Fprintf(w, "%d", value.(int64))
	case OBJTYPE_MAP:
		host.FindObject(value.(int32)).(*HostMap).Dump(w)
	case OBJTYPE_STRING:
		fmt.Fprintf(w, "\"%s\"", value.(string))
	default:
		if (typeId & OBJTYPE_ARRAY) == 0 {
			panic("typeId is not an array")
		}
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

func (host *SimpleWasmHost) SetExport(index int32, functionName string) {
	if index < 0 {
		if index != KeyZzzzzzz {
			host.SetError("SetExport: predefined key value mismatch")
		}
		return
	}
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
