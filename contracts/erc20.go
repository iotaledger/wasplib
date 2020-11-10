package main

import (
	"github.com/iotaledger/wasplib/client"
	"strconv"
)

const (
	varSupply        = "s"
	varBalances      = "b"
	varTargetAddress = "addr"
	varAmount        = "amount"
)

func main() {
}

//export onLoad
func onLoadERC20() {
	exports := client.NewScExports()
	exports.Add("initSC")
	exports.Add("transfer")
	exports.Add("approve")
}

//export initSC
func initSC() {
	sc := client.NewScContext()
	sc.Log("initSC")

	scOwner := sc.Contract().Owner()
	request := sc.Request()
	if !request.From(scOwner) {
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
	params := sc.Request().Params()
	supplyParam := params.GetInt(varSupply)
	if supplyParam.Value() == 0 {
		sc.Log("initSC.fail: wrong 'supply' parameter")
		return
	}
	supply := supplyParam.Value()
	supplyState.SetValue(supply)
	state.GetKeyMap(varBalances).GetInt(sc.Contract().Owner().Bytes()).SetValue(supply)

	sc.Log("initSC: success")
}

//export transfer
func transfer() {
	sc := client.NewScContext()
	sc.Log("transfer")

	state := sc.State()
	request := sc.Request()
	balances := state.GetKeyMap(varBalances)

	sender := request.Address()

	sc.Log("sender address: " + sender.String())

	sourceBalance := balances.GetInt(sender.Bytes())

	sc.Log("source balance: " + strconv.FormatInt(sourceBalance.Value(), 10))

	params := request.Params()
	amount := params.GetInt(varAmount)
	if amount.Value() == 0 {
		sc.Log("transfer.fail: wrong 'amount' parameter")
		return
	}
	if amount.Value() > sourceBalance.Value() {
		sc.Log("transfer.fail: not enough balance")
		return
	}
	targetAddr := params.GetAddress(varTargetAddress)
	// TODO check if it is a correct address, otherwise won't be possible to transfer from it

	targetBalance := balances.GetInt(targetAddr.Value().Bytes())
	targetBalance.SetValue(targetBalance.Value() + amount.Value())
	sourceBalance.SetValue(sourceBalance.Value() - amount.Value())

	sc.Log("transfer: success")
}

//export approve
func approve() {
	sc := client.NewScContext()
	// TODO
	sc.Log("approve")
}
