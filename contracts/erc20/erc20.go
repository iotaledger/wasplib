// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// implementation of ERC-20 smart contract for ISCP
// following https://ethereum.org/en/developers/tutorials/understand-the-erc-20-token-smart-contract/

package erc20

import "github.com/iotaledger/wasplib/client"

// state variable
const stateVarSupply = client.Key("s")

// supply constant
const stateVarBalances = client.Key("b") // name of the map of balances

// params and return variables, used in calls
const paramSupply = client.Key("s")
const paramCreator = client.Key("c")
const paramAccount = client.Key("ac")
const paramDelegation = client.Key("d")
const paramAmount = client.Key("am")
const paramRecipient = client.Key("r")

func OnLoad() {
	exports := client.NewScExports()
	exports.AddCall("init", onInit)
	exports.AddView("total_supply", totalSupply)
	exports.AddView("balance_of", balanceOf)
	exports.AddView("allowance", allowance)
	exports.AddCall("transfer", transfer)
	exports.AddCall("approve", approve)
	exports.AddCall("transfer_from", transferFrom)
}

// TODO would be awesome to have some less syntactically cumbersome way to check and validate parameters.

// on_init is a constructor entry point. It initializes the smart contract with the
// initial value of the token supply and the owner of that supply
// - input:
//   -- PARAM_SUPPLY must be nonzero positive integer
//   -- PARAM_CREATOR is the AgentID where initial supply is placed
func onInit(ctx *client.ScCallContext) {
	ctx.Log("erc20.onInit.begin")
	// validate parameters
	// supply
	supply := ctx.Params().GetInt(paramSupply)
	if !supply.Exists() || supply.Value() <= 0 {
		err := "erc20.onInit.fail: wrong 'supply' parameter"
		ctx.Log(err)
		ctx.Error().SetValue(err)
		return
	}
	// creator (owner)
	// we cannot use 'caller' here because on_init is always called from the 'root'
	// so, owner of the initial supply must be provided as a parameter PARAM_CREATOR to constructor (on_init)
	creator := ctx.Params().GetAgent(paramCreator)
	if !creator.Exists() {
		err := "erc20.onInit.fail: wrong 'creator' parameter"
		ctx.Log(err)
		ctx.Error().SetValue(err)
		return
	}
	ctx.State().GetInt(stateVarSupply).SetValue(supply.Value())

	// assign the whole supply to creator
	ctx.State().GetMap(stateVarBalances).GetInt(creator.Value()).SetValue(supply.Value())

	t := "erc20.onInit.success. Supply: " + supply.String() +
		", creator:" + creator.String()
	ctx.Log(t)
}

// the view returns total supply set when creating the contract (a constant).
// Output:
// - PARAM_SUPPLY: i64
func totalSupply(ctx *client.ScViewContext) {
	supply := ctx.State().GetInt(stateVarSupply).Value()
	ctx.Results().GetInt(paramSupply).SetValue(supply)
}

// the view returns balance of the token held in the account
// Input:
// - PARAM_ACCOUNT: agentID
func balanceOf(ctx *client.ScViewContext) {
	account := ctx.Params().GetAgent(paramAccount)
	if !account.Exists() {
		m := "wrong or non existing parameter: " + account.String()
		ctx.Log(m)
		ctx.Error().SetValue(m)
		return
	}
	balances := ctx.State().GetMap(stateVarBalances)
	balance := balances.GetInt(account.Value()).Value() // 0 if doesn't exist
	ctx.Results().GetInt(paramAmount).SetValue(balance)
}

// the view returns max number of tokens the owner PARAM_ACCOUNT of the account
// allowed to retrieve to another party PARAM_DELEGATION
// Input:
// - PARAM_ACCOUNT: agentID
// - PARAM_DELEGATION: agentID
// Output:
// - PARAM_AMOUNT: i64. 0 if delegation doesn't exists
func allowance(ctx *client.ScViewContext) {
	ctx.Log("erc20.allowance")
	// validate parameters
	// account
	owner := ctx.Params().GetAgent(paramAccount)
	if !owner.Exists() {
		m := "erc20.allowance.fail: wrong 'account' parameter"
		ctx.Log(m)
		ctx.Error().SetValue(m)
		return
	}
	// delegation
	delegation := ctx.Params().GetAgent(paramDelegation)
	if !delegation.Exists() {
		m := "erc20.allowance.fail: wrong 'delegation' parameter"
		ctx.Log(m)
		ctx.Error().SetValue(m)
		return
	}
	// all allowances of the address 'owner' are stored in the map of the same name
	allowances := ctx.State().GetMap(owner.Value())
	allow := allowances.GetInt(delegation.Value()).Value()
	ctx.Results().GetInt(paramAmount).SetValue(allow)
}

// transfer moves tokens from caller's account to target account
// Input:
// - PARAM_ACCOUNT: agentID
// - PARAM_AMOUNT: i64
func transfer(ctx *client.ScCallContext) {
	ctx.Log("erc20.transfer")

	// validate params
	params := ctx.Params()
	// account
	targetAddrParam := params.GetAgent(paramAccount)
	if !targetAddrParam.Exists() {
		m := "erc20.transfer.fail: wrong 'account' parameter"
		ctx.Log(m)
		ctx.Error().SetValue(m)
		return
	}
	targetAddr := targetAddrParam.Value()
	// amount
	amount := params.GetInt(paramAmount).Value()
	if amount <= 0 {
		m := "erc20.transfer.fail: wrong 'amount' parameter"
		ctx.Log(m)
		ctx.Error().SetValue(m)
		return
	}
	balances := ctx.State().GetMap(stateVarBalances)
	sourceBalance := balances.GetInt(ctx.Caller())

	if sourceBalance.Value() < amount {
		m := "erc20.transfer.fail: not enough funds"
		ctx.Log(m)
		ctx.Error().SetValue(m)
		return
	}
	targetBalance := balances.GetInt(targetAddr)
	result := targetBalance.Value() + amount
	if result <= 0 {
		m := "erc20.transfer.fail: overflow"
		ctx.Log(m)
		ctx.Error().SetValue(m)
		return
	}
	sourceBalance.SetValue(sourceBalance.Value() - amount)
	targetBalance.SetValue(targetBalance.Value() + amount)
	ctx.Log("erc20.transfer.success")
}

// Sets the allowance value for delegated account
// inputs:
//  - PARAM_DELEGATION: agentID
//  - PARAM_AMOUNT: i64
func approve(ctx *client.ScCallContext) {
	ctx.Log("erc20.approve")

	// validate parameters
	delegationParam := ctx.Params().GetAgent(paramDelegation)
	if !delegationParam.Exists() {
		m := "erc20.approve.fail: wrong 'delegation' parameter"
		ctx.Log(m)
		ctx.Error().SetValue(m)
		return
	}
	delegation := delegationParam.Value()
	amount := ctx.Params().GetInt(paramAmount).Value()
	if amount <= 0 {
		m := "erc20.approve.fail: wrong 'amount' parameter"
		ctx.Log(m)
		ctx.Error().SetValue(m)
		return
	}
	// all allowances are in the map under the name of he owner
	allowances := ctx.State().GetMap(ctx.Caller())
	allowances.GetInt(delegation).SetValue(amount)
	ctx.Log("erc20.approve.success")
}

// Moves the amount of tokens from sender to recipient using the allowance mechanism.
// Amount is then deducted from the caller’s allowance. This function emits the Transfer event.
// Input:
// - PARAM_ACCOUNT: agentID   the spender
// - PARAM_RECIPIENT: agentID   the target
// - PARAM_AMOUNT: i64
func transferFrom(ctx *client.ScCallContext) {
	ctx.Log("erc20.transferFrom")

	// validate parameters
	accountParam := ctx.Params().GetAgent(paramAccount)
	if !accountParam.Exists() {
		m := "erc20.transferFrom.fail: wrong 'account' parameter"
		ctx.Log(m)
		ctx.Error().SetValue(m)
		return
	}
	account := accountParam.Value()
	recipientParam := ctx.Params().GetAgent(paramRecipient)
	if !recipientParam.Exists() {
		m := "erc20.transferFrom.fail: wrong 'recipient' parameter"
		ctx.Log(m)
		ctx.Error().SetValue(m)
		return
	}
	recipient := recipientParam.Value()
	amountParam := ctx.Params().GetInt(paramAmount)
	if !amountParam.Exists() {
		m := "erc20.transferFrom.fail: wrong 'amount' parameter"
		ctx.Log(m)
		ctx.Error().SetValue(m)
		return
	}
	amount := amountParam.Value()

	// allowances are in the map under the name of the account
	allowances := ctx.State().GetMap(account)
	allowance := allowances.GetInt(recipient)
	if allowance.Value() < amount {
		m := "erc20.transferFrom.fail: not enough allowance"
		ctx.Log(m)
		ctx.Error().SetValue(m)
		return
	}
	balances := ctx.State().GetMap(stateVarBalances)
	sourceBalance := balances.GetInt(account)
	if sourceBalance.Value() < amount {
		m := "erc20.transferFrom.fail: not enough funds"
		ctx.Log(m)
		ctx.Error().SetValue(m)
		return
	}
	recipientBalance := balances.GetInt(recipient)
	result := recipientBalance.Value() + amount
	if result <= 0 {
		m := "erc20.transferFrom.fail: overflow"
		ctx.Log(m)
		ctx.Error().SetValue(m)
		return
	}
	sourceBalance.SetValue(sourceBalance.Value() - amount)
	recipientBalance.SetValue(recipientBalance.Value() + amount)
	allowance.SetValue(allowance.Value() - amount)

	ctx.Log("erc20.transferFrom.success")
}
