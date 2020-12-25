// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package wasmlocalhost

import (
	"fmt"
	"github.com/iotaledger/wasp/packages/vm/wasmhost"
	"github.com/mr-tron/base58"
	"io"
)

var EnableImmutableChecks = true

type SimpleWasmHost struct {
	wasmhost.WasmHost
}

func NewSimpleWasmHost(vm wasmhost.WasmVM) (*SimpleWasmHost, error) {
	host := &SimpleWasmHost{}
	err := host.InitVM(vm, true)
	if err != nil {
		return nil, err
	}
	root := NewHostMap(host, 0)
	root.InitObj(1, 0)
	host.Init(NewNullObject(host), root, host)
	return host, nil
}

func (host *SimpleWasmHost) Dump(w io.Writer, typeId int32, value interface{}) {
	switch typeId {
	case wasmhost.OBJTYPE_ADDRESS,
		wasmhost.OBJTYPE_AGENT,
		wasmhost.OBJTYPE_BYTES,
		wasmhost.OBJTYPE_COLOR:
		fmt.Fprintf(w, "\"%s\"", base58.Encode(value.([]byte)))
	case wasmhost.OBJTYPE_INT:
		fmt.Fprintf(w, "%d", value.(int64))
	case wasmhost.OBJTYPE_MAP:
		host.FindObject(value.(int32)).(*HostMap).Dump(w)
	case wasmhost.OBJTYPE_STRING:
		fmt.Fprintf(w, "\"%s\"", value.(string))
	default:
		if (typeId & wasmhost.OBJTYPE_ARRAY) == 0 {
			panic("typeId is not an array")
		}
		host.FindObject(value.(int32)).(*HostArray).Dump(w)
	}
}

func (host *SimpleWasmHost) Log(logLevel int32, text string) {
	switch logLevel {
	case wasmhost.KeyTraceAll:
		//fmt.Println(text)
	case wasmhost.KeyTrace:
		//fmt.Println(text)
	case wasmhost.KeyLog:
		fmt.Println(text)
	case wasmhost.KeyWarning:
		fmt.Println(text)
	case wasmhost.KeyError:
		fmt.Println(text)
	}
}
