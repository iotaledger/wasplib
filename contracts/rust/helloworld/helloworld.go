// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package helloworld

import (
	"github.com/iotaledger/wasplib/packages/vm/wasmlib"
)

func funcHelloWorld(ctx wasmlib.ScFuncContext, f*HelloWorldContext) {
	ctx.Log("Hello, world!")
}

func viewGetHelloWorld(ctx wasmlib.ScViewContext, f*GetHelloWorldContext) {
	f.Results.HelloWorld().SetValue("Hello, world!")
}
