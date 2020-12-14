// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// implementation of ERC-20 smart contract for ISCP
// following https://ethereum.org/en/developers/tutorials/understand-the-erc-20-token-smart-contract/

package erc20

import (
	"github.com/iotaledger/wasplib/client"
)

// state variable
const STATE_VAR_SUPPLY = client.Key("s")

// supply constant
const STATE_VAR_BALANCES = client.Key("b") // name of the map of balances

// params and return variables, used in calls
const PARAM_SUPPLY = client.Key("s")
const PARAM_CREATOR = client.Key("c")
const PARAM_ACCOUNT = client.Key("ac")
const PARAM_DELEGATION = client.Key("d")
const PARAM_AMOUNT = client.Key("am")
const PARAM_RECIPIENT = client.Key("r")

func OnLoad() {
	exports := client.NewScExports()
	exports.AddCall("init", onInit)
	exports.AddView("total_supply", total_supply)
	exports.AddView("balance_of", balance_of)
	exports.AddView("allowance", allowance)
	exports.AddCall("transfer", transfer)
	exports.AddCall("approve", approve)
	exports.AddCall("transfer_from", transfer_from)
}

// TODO would be awesome to have some less syntactically cumbersome way to check and validate parameters.

// onInit is a constructor entry point. It initializes the smart contract with the
// initial value of the token supply and the owner of that supply
// - input:
//   -- PARAM_SUPPLY must be nonzero positive integer
//   -- PARAM_CREATOR is the AgentID where initial supply is placed
func onInit(ctx *client.ScCallContext) {
	ctx.Log("erc20.onInit.begin")
	// validate parameters
	// supply
	supply := ctx.Params().GetInt(PARAM_SUPPLY)
	if !supply.Exists() || supply.Value() <= 0 {
		err := "erc20.onInit.fail: wrong 'supply' parameter"
		ctx.Log(err)
		ctx.Error().SetValue(err)
		return
	}
	// creator (owner)
	// we cannot use 'caller' here because onInit is always called from the 'root'
	// so, owner of the initial supply must be provided as a parameter PARAM_CREATOR to constructor (onInit)
	creator := ctx.Params().GetAgent(PARAM_CREATOR)
	if !creator.Exists() {
		err := "erc20.onInit.fail: wrong 'creator' parameter"
		ctx.Log(err)
		ctx.Error().SetValue(err)
		return
	}
	ctx.State().GetInt(STATE_VAR_SUPPLY).SetValue(supply.Value())

	// assign the whole supply to creator
	ctx.State().GetMap(STATE_VAR_BALANCES).GetInt(creator.Value()).SetValue(supply.Value())

	t := "erc20.onInit.success. Supply: " + supply.String() +
		", creator:" + creator.Value().String()
	ctx.Log(t)
}

// the view returns total supply set when creating the contract (a constant).
// Output:
// - PARAM_SUPPLY: i64
func total_supply(ctx *client.ScViewContext) {
	supply := ctx.State().GetInt(STATE_VAR_SUPPLY).Value()
	ctx.Results().GetInt(PARAM_SUPPLY).SetValue(supply)
}

// the view returns balance of the token held in the account
// Input:
// - PARAM_ACCOUNT: agentID
func balance_of(ctx *client.ScViewContext) {
	account := ctx.Params().GetAgent(PARAM_ACCOUNT)
	if !account.Exists() {
		m := "wrong or non existing parameter: " + account.Value().String()
		ctx.Log(m)
		ctx.Error().SetValue(m)
		return
	}
	balances := ctx.State().GetMap(STATE_VAR_BALANCES)
	balance := balances.GetInt(account.Value()).Value() // 0 if doesn't exist
	ctx.Results().GetInt(PARAM_AMOUNT).SetValue(balance)
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
	owner := ctx.Params().GetAgent(PARAM_ACCOUNT)
	if !owner.Exists() {
		m := "erc20.allowance.fail: wrong 'account' parameter"
		ctx.Log(m)
		ctx.Error().SetValue(m)
		return
	}
	// delegation
	delegation := ctx.Params().GetAgent(PARAM_DELEGATION)
	if !delegation.Exists() {
		m := "erc20.allowance.fail: wrong 'delegation' parameter"
		ctx.Log(m)
		ctx.Error().SetValue(m)
		return
	}
	// all allowances of the address 'owner' are stored in the map of the same name
	allowances := ctx.State().GetMap(owner.Value())
	allow := allowances.GetInt(delegation.Value()).Value()
	ctx.Results().GetInt(PARAM_AMOUNT).SetValue(allow)
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
	target_addrParam := params.GetAgent(PARAM_ACCOUNT)
	if !target_addrParam.Exists() {
		m := "erc20.transfer.fail: wrong 'account' parameter"
		ctx.Log(m)
		ctx.Error().SetValue(m)
		return
	}
	target_addr := target_addrParam.Value()
	// amount
	amount := params.GetInt(PARAM_AMOUNT).Value()
	if amount <= 0 {
		m := "erc20.transfer.fail: wrong 'amount' parameter"
		ctx.Log(m)
		ctx.Error().SetValue(m)
		return
	}
	balances := ctx.State().GetMap(STATE_VAR_BALANCES)
	source_balance := balances.GetInt(ctx.Caller())

	if source_balance.Value() < amount {
		m := "erc20.transfer.fail: not enough funds"
		ctx.Log(m)
		ctx.Error().SetValue(m)
		return
	}
	target_balance := balances.GetInt(target_addr)
	result := target_balance.Value() + amount
	if result <= 0 {
		m := "erc20.transfer.fail: overflow"
		ctx.Log(m)
		ctx.Error().SetValue(m)
		return
	}
	source_balance.SetValue(source_balance.Value() - amount)
	target_balance.SetValue(target_balance.Value() + amount)
	ctx.Log("erc20.transfer.success")
}

// Sets the allowance value for delegated account
// inputs:
//  - PARAM_DELEGATION: agentID
//  - PARAM_AMOUNT: i64
func approve(ctx *client.ScCallContext) {
	ctx.Log("erc20.approve")

	// validate parameters
	delegationParam := ctx.Params().GetAgent(PARAM_DELEGATION)
	if !delegationParam.Exists() {
		m := "erc20.approve.fail: wrong 'delegation' parameter"
		ctx.Log(m)
		ctx.Error().SetValue(m)
		return
	}
	delegation := delegationParam.Value()
	amount := ctx.Params().GetInt(PARAM_AMOUNT).Value()
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
func transfer_from(ctx *client.ScCallContext) {
	ctx.Log("erc20.transfer_from")

	// validate parameters
	accountParam := ctx.Params().GetAgent(PARAM_ACCOUNT)
	if !accountParam.Exists() {
		m := "erc20.transfer_from.fail: wrong 'account' parameter"
		ctx.Log(m)
		ctx.Error().SetValue(m)
		return
	}
	account := accountParam.Value()
	recipientParam := ctx.Params().GetAgent(PARAM_RECIPIENT)
	if !recipientParam.Exists() {
		m := "erc20.transfer_from.fail: wrong 'recipient' parameter"
		ctx.Log(m)
		ctx.Error().SetValue(m)
		return
	}
	recipient := recipientParam.Value()
	amountParam := ctx.Params().GetInt(PARAM_AMOUNT)
	if !amountParam.Exists() {
		m := "erc20.transfer_from.fail: wrong 'amount' parameter"
		ctx.Log(m)
		ctx.Error().SetValue(m)
		return
	}
	amount := amountParam.Value()

	// allowances are in the map under the name of the account
	allowances := ctx.State().GetMap(account)
	allowance := allowances.GetInt(recipient)
	if allowance.Value() < amount {
		m := "erc20.transfer_from.fail: not enough allowance"
		ctx.Log(m)
		ctx.Error().SetValue(m)
		return
	}
	balances := ctx.State().GetMap(STATE_VAR_BALANCES)
	source_balance := balances.GetInt(account)
	if source_balance.Value() < amount {
		m := "erc20.transfer_from.fail: not enough funds"
		ctx.Log(m)
		ctx.Error().SetValue(m)
		return
	}
	recipient_balance := balances.GetInt(recipient)
	result := recipient_balance.Value() + amount
	if result <= 0 {
		m := "erc20.transfer_from.fail: overflow"
		ctx.Log(m)
		ctx.Error().SetValue(m)
		return
	}
	source_balance.SetValue(source_balance.Value() - amount)
	recipient_balance.SetValue(recipient_balance.Value() + amount)
	allowance.SetValue(allowance.Value() - amount)

	ctx.Log("erc20.transfer_from.success")
}
