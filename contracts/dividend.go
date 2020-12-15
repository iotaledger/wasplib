// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// +build wasm

package main

import (
	"github.com/iotaledger/wasplib/contracts/dividend"
	"github.com/iotaledger/wasplib/wasmclient"
)

func main() {
}

//export on_load
func dividendOnLoad() {
	wasmclient.ConnectWasmHost()
	dividend.OnLoad()
}
