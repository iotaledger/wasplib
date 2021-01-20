// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package dummy

import "github.com/iotaledger/wasplib/client"

const KeyFailInitParam = client.Key("failInitParam")

func OnLoad() {
	exports := client.NewScExports()
	exports.AddCall("init", onInit)
}

// fails with error if failInitParam exists
func onInit(ctx *client.ScCallContext) {
	failParam := ctx.Params().GetInt(KeyFailInitParam)
	if failParam.Exists() {
		ctx.Panic("dummy: failing on purpose")
	}
}
