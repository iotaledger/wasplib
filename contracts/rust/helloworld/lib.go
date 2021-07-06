// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead


package helloworld

import "github.com/iotaledger/wasplib/packages/vm/wasmlib"

func OnLoad() {
	exports := wasmlib.NewScExports()
	exports.AddFunc(FuncHelloWorld, funcHelloWorldThunk)
	exports.AddView(ViewGetHelloWorld, viewGetHelloWorldThunk)

	for i, key := range keyMap {
		idxMap[i] = key.KeyID()
	}
}

type HelloWorldContext struct {
	State MutableHelloWorldState
}

func funcHelloWorldThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("helloworld.funcHelloWorld")
	f := &HelloWorldContext{
		State: MutableHelloWorldState{
			id: wasmlib.OBJ_ID_STATE,
		},
	}
	funcHelloWorld(ctx, f)
	ctx.Log("helloworld.funcHelloWorld ok")
}

type GetHelloWorldContext struct {
	Results MutableGetHelloWorldResults
	State   ImmutableHelloWorldState
}

func viewGetHelloWorldThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("helloworld.viewGetHelloWorld")
	f := &GetHelloWorldContext{
		Results: MutableGetHelloWorldResults{
			id: wasmlib.OBJ_ID_RESULTS,
		},
		State: ImmutableHelloWorldState{
			id: wasmlib.OBJ_ID_STATE,
		},
	}
	viewGetHelloWorld(ctx, f)
	ctx.Log("helloworld.viewGetHelloWorld ok")
}
