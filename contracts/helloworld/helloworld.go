// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package helloworld

import (
	"github.com/iotaledger/wasplib/client"
)

func OnLoad() {
	exports := client.NewScExports()
	exports.AddCall("hello_world", helloWorld)
}

func helloWorld(sc *client.ScCallContext) {
	sc.Log("Hello, world!")
}
