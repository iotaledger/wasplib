// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"github.com/iotaledger/wasplib/client"
)

func main() {
}

//export onLoad
func onLoadHelloWorld() {
	exports := client.NewScExports()
	exports.AddCall("helloWorld", helloWorld)
}

func helloWorld(sc *client.ScCallContext) {
	sc.Log("Hello, world!")
}
