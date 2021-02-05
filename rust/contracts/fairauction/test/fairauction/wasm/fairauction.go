// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// +build wasm

package main

import (
	"github.com/iotaledger/wasplib/rust/contracts/fairauction/test/fairauction"
	"github.com/iotaledger/wasplib/client/wasm"
)

func main() {
}

//export on_load
func fairauctionOnLoad() {
	wasmclient.ConnectWasmHost()
	fairauction.OnLoad()
}
