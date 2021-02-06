// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// +build wasm

package main

import (
	"github.com/iotaledger/wasp/packages/vm/wasmclient"
	"github.com/iotaledger/wasplib/contracts/dummy"
)

func main() {
}

//export on_load
func dummyOnLoad() {
	wasmclient.ConnectWasmHost()
	dummy.OnLoad()
}
