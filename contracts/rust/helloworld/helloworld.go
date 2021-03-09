// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package helloworld

import (
	"github.com/iotaledger/wasplib/packages/vm/wasmlib"
)

func funcHelloWorld(ctx wasmlib.ScFuncContext, params *FuncHelloWorldParams) {
	ctx.Log("Hello, world!")
}

func viewGetHelloWorld(ctx wasmlib.ScViewContext, params *ViewGetHelloWorldParams) {
	ctx.Results().GetString(VarHelloWorld).SetValue("Hello, world!")
}