// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package aaa

import "github.com/iotaledger/wasp/packages/vm/wasmlib"

func funcDonate(ctx *wasmlib.ScCallContext, params *FuncDonateParams) {
	ctx.Log("calling donate")
}

func funcWithdraw(ctx *wasmlib.ScCallContext, params *FuncWithdrawParams) {
	ctx.Log("calling withdraw")
}

func viewDonations(ctx *wasmlib.ScViewContext, params *ViewDonationsParams) {
	ctx.Log("calling donations")
}
