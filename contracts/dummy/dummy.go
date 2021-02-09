// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package dummy

import "github.com/iotaledger/wasp/packages/vm/wasmlib"

// fails with error if failInitParam exists
func funcInit(ctx *wasmlib.ScFuncContext, params *FuncInitParams) {
	if params.FailInitParam.Exists() {
		ctx.Panic("dummy: failing on purpose")
	}
}
