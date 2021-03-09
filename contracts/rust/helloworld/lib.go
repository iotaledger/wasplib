// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package helloworld

import "github.com/iotaledger/wasplib/packages/vm/wasmlib"

func OnLoad() {
	exports := wasmlib.NewScExports()
	exports.AddFunc(FuncHelloWorld, funcHelloWorldThunk)
	exports.AddView(ViewGetHelloWorld, viewGetHelloWorldThunk)
}

type FuncHelloWorldParams struct {
}

func funcHelloWorldThunk(ctx wasmlib.ScFuncContext) {
	params := &FuncHelloWorldParams{
	}
	ctx.Log("helloworld.funcHelloWorld")
	funcHelloWorld(ctx, params)
	ctx.Log("helloworld.funcHelloWorld ok")
}

type ViewGetHelloWorldParams struct {
}

func viewGetHelloWorldThunk(ctx wasmlib.ScViewContext) {
	params := &ViewGetHelloWorldParams{
	}
	ctx.Log("helloworld.viewGetHelloWorld")
	viewGetHelloWorld(ctx, params)
	ctx.Log("helloworld.viewGetHelloWorld ok")
}
