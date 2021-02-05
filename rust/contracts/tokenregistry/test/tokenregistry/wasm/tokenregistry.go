// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// +build wasm

package main

import (
	"github.com/iotaledger/wasplib/rust/contracts/tokenregistry/test/tokenregistry"
	"github.com/iotaledger/wasplib/client/wasm"
)

func main() {
}

//export on_load
func tokenregistryOnLoad() {
	wasmclient.ConnectWasmHost()
	tokenregistry.OnLoad()
}
