// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package aaa

import "github.com/iotaledger/wasp/packages/vm/wasmlib"


func funcMyFunc(ctx *wasmlib.ScFuncContext, params *FuncMyFuncParams) {
    ctx.Log("calling myFunc")
}

func viewMyView(ctx *wasmlib.ScViewContext, params *ViewMyViewParams) {
    ctx.Log("calling myView")
}
