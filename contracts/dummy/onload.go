// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package dummy

import "github.com/iotaledger/wasp/packages/vm/wasmlib"

func OnLoad() {
    exports := wasmlib.NewScExports()
    exports.AddFunc(FuncInit, funcInitThunk)
}

type FuncInitParams struct {
    FailInitParam wasmlib.ScImmutableInt // when present fail on purpose
}

func funcInitThunk(ctx *wasmlib.ScFuncContext) {
    p := ctx.Params()
    params := &FuncInitParams {
        FailInitParam: p.GetInt(ParamFailInitParam),
    }
    funcInit(ctx, params)
}
