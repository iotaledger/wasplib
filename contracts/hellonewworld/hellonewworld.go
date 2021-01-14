// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package hellonewworld

import "github.com/iotaledger/wasplib/client"

const KeyCounter = client.Key("counter")

func OnLoad() {
	exports := client.NewScExports()
	exports.AddCall("hello", hello)
	exports.AddView("getCounter", getCounter)
}

// Function hello implement smart contract entry point "hello".
// Function hello logs the message "Hello, new world!" with the counter and increments the counter
func hello(ctx *client.ScCallContext) {
	counter := ctx.State().GetInt(KeyCounter)
	msg := "Hello, new world! #" + counter.String()
	ctx.Log(msg) // todo info and debug levels, not events!
	counter.SetValue(counter.Value() + 1)
}

// Function get_counter implements smart contract VIEW entry point "getCounter".
// It return counter value in the result dictionary with the key "counter"
func getCounter(ctx *client.ScViewContext) {
	counter := ctx.State().GetInt(KeyCounter).Value()
	ctx.Results().GetInt(KeyCounter).SetValue(counter)
}
