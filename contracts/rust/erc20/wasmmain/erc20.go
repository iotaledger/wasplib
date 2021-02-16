// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

// +build wasm

package main

import "github.com/iotaledger/wasplib/packages/vm/wasmclient"
import "github.com/iotaledger/wasplib/contracts/rust/erc20"

func main() {
}

//export on_load
func OnLoad() {
	wasmclient.ConnectWasmHost()
	erc20.OnLoad()
}
