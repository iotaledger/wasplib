// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// +build wasm

package main

import (
	"github.com/iotaledger/wasp/packages/vm/wasmclient"
	"github.com/iotaledger/wasplib/contracts/example1"
)

func main() {
}

//export on_load
func example1OnLoad() {
	wasmclient.ConnectWasmHost()
	example1.OnLoad()
}
