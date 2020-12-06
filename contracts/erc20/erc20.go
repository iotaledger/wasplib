// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package erc20

import (
	"github.com/iotaledger/wasplib/client"
)

const (
	varSupply        = client.Key("s")
	varBalances      = client.Key("b")
	varTargetAddress = client.Key("addr")
	varAmount        = client.Key("amount")
)

func OnLoad() {
	exports := client.NewScExports()
	exports.AddCall("init", onInit)
	exports.AddCall("transfer", transfer)
	exports.AddCall("approve", approve)
}

func onInit(sc *client.ScCallContext) {
	sc.Log("initSC")

	scOwner := sc.Contract().Owner()
	if !sc.From(scOwner) {
		sc.Log("Cancel spoofed request")
		return
	}

	state := sc.State()
	supplyState := state.GetInt(varSupply)
	if supplyState.Value() > 0 {
		// already initialized
		sc.Log("initSC.fail: already initialized")
		return
	}
	params := sc.Params()
	supplyParam := params.GetInt(varSupply)
	if supplyParam.Value() == 0 {
		sc.Log("initSC.fail: wrong 'supply' parameter")
		return
	}
	supply := supplyParam.Value()
	supplyState.SetValue(supply)
	state.GetMap(varBalances).GetInt(sc.Contract().Owner()).SetValue(supply)

	sc.Log("initSC: success")
}

func transfer(sc *client.ScCallContext) {
	sc.Log("transfer")

	state := sc.State()
	balances := state.GetMap(varBalances)

	caller := sc.Caller()

	sc.Log("sender address: " + caller.String())

	sourceBalance := balances.GetInt(caller)

	sc.Log("source balance: " + sc.Utility().String(sourceBalance.Value()))

	params := sc.Params()
	amount := params.GetInt(varAmount)
	if amount.Value() == 0 {
		sc.Log("transfer.fail: wrong 'amount' parameter")
		return
	}
	if amount.Value() > sourceBalance.Value() {
		sc.Log("transfer.fail: not enough balance")
		return
	}
	targetAddr := params.GetAgent(varTargetAddress)
	// TODO check if it is a correct address, otherwise won't be possible to transfer from it

	targetBalance := balances.GetInt(targetAddr.Value())
	targetBalance.SetValue(targetBalance.Value() + amount.Value())
	sourceBalance.SetValue(sourceBalance.Value() - amount.Value())

	sc.Log("transfer: success")
}

func approve(sc *client.ScCallContext) {
	// TODO
	sc.Log("approve")
}
