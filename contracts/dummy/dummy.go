// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package dummy

import "github.com/iotaledger/wasp/packages/vm/wasmlib"

const ParamFailInitParam = wasmlib.Key("failInitParam")

func OnLoad() {
	exports := wasmlib.NewScExports()
	exports.AddCall("init", onInit)
}

// fails with error if failInitParam exists
func onInit(ctx *wasmlib.ScCallContext) {
	failParam := ctx.Params().GetInt(ParamFailInitParam)
	if failParam.Exists() {
		ctx.Panic("dummy: failing on purpose")
	}
}
