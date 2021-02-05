// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package helloworld

import "github.com/iotaledger/wasplib/client"

func funcHelloWorld(ctx *client.ScCallContext) {
	ctx.Log("Hello, world!")
}

func viewGetHelloWorld(ctx *client.ScViewContext) {
	ctx.Log("Get Hello world!")
	ctx.Results().GetString(VarHelloWorld).SetValue("Hello, world!")
}
