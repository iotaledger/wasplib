// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"github.com/iotaledger/wasplib/contracts/inccounter"
	"github.com/iotaledger/wasplib/wasmclient"
)

func main() {
}

//export onLoad
func onLoadIncCounter() {
	wasmclient.ConnectWasmHost()
	inccounter.OnLoad()
}
