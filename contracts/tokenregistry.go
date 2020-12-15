// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// +build wasm

package main

import (
	"github.com/iotaledger/wasplib/contracts/tokenregistry"
	"github.com/iotaledger/wasplib/wasmclient"
)

func main() {
}

//export on_load
func tokenregistryOnLoad() {
	wasmclient.ConnectWasmHost()
	tokenregistry.OnLoad()
}
