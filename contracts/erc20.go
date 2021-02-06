// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// +build wasm

package main

import (
	"github.com/iotaledger/wasp/packages/vm/wasmclient"
	"github.com/iotaledger/wasplib/contracts/erc20"
)

func main() {
}

//export on_load
func erc20OnLoad() {
	wasmclient.ConnectWasmHost()
	erc20.OnLoad()
}
