// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package helloworld

import "github.com/iotaledger/wasp/packages/vm/wasmlib"

func funcHelloWorld(ctx *wasmlib.ScCallContext) {
	ctx.Log("Hello, world!")
}

func viewGetHelloWorld(ctx *wasmlib.ScViewContext) {
	ctx.Log("Get Hello world!")
	ctx.Results().GetString(VarHelloWorld).SetValue("Hello, world!")
}
