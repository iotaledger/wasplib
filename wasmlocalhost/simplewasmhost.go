// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package wasmlocalhost

import (
	"fmt"
	"github.com/iotaledger/hive.go/logger"
	"github.com/iotaledger/wasp/packages/vm/wasmhost"
	"github.com/iotaledger/wasplib/client"
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
	case client.TYPE_ADDRESS,
		client.TYPE_AGENT_ID,
		client.TYPE_BYTES,
		client.TYPE_COLOR:
		fmt.Fprintf(w, "\"%s\"", base58.Encode(value.([]byte)))
	case client.TYPE_INT:
		fmt.Fprintf(w, "%d", value.(int64))
	case client.TYPE_MAP:
		obj := host.FindObject(value.(int32))
		switch obj.(type) {
		case *HostMap:
			obj.(*HostMap).Dump(w)
		}
	case client.TYPE_STRING:
		fmt.Fprintf(w, "\"%s\"", value.(string))
	default:
		if (typeId & client.TYPE_ARRAY) == 0 {
			panic("typeId is not an array")
		}
		host.FindObject(value.(int32)).(*HostArray).Dump(w)
	}
}

func (host *SimpleWasmHost) Error(text string) {
	panic(text)
}
