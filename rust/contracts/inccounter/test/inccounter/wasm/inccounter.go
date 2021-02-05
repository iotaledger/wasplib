// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// +build wasm

package main

import (
	"github.com/iotaledger/wasplib/rust/contracts/inccounter/test/inccounter"
	"github.com/iotaledger/wasplib/client/wasm"
)

func main() {
}

//export on_load
func inccounterOnLoad() {
	wasmclient.ConnectWasmHost()
	inccounter.OnLoad()
}
