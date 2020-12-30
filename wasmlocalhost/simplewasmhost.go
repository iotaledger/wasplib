// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package wasmlocalhost

import (
	"fmt"
	"github.com/iotaledger/hive.go/logger"
	"github.com/iotaledger/wasp/packages/vm/wasmhost"
	"github.com/mr-tron/base58"
	"io"
)

var EnableImmutableChecks = true
var cfgDefault = logger.Config{
	Level:         "warn", // warn, info, or debug
	Encoding:      "console",
	OutputPaths:   []string{"stdout"},
	DisableEvents: true,
}

type SimpleWasmHost struct {
	wasmhost.WasmHost
	panicked bool
}

func NewSimpleWasmHost(vm wasmhost.WasmVM) (*SimpleWasmHost, error) {
	host := &SimpleWasmHost{}
	err := host.InitVM(vm, true)
	if err != nil {
		return nil, err
	}
	root := NewHostMap(host, 0)
	root.InitObj(1, 0)

	rootLogger, err := logger.NewRootLogger(cfgDefault)
	host.Init(NewNullObject(host), root, rootLogger)
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

func (host *SimpleWasmHost) Error(text string) {
	panic(text)
}
