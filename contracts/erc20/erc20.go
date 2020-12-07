// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package erc20

import (
	"github.com/iotaledger/wasplib/client"
)

const (
	keyAccount    = client.Key("ac")
	keyAmount     = client.Key("am")
	keyBalances   = client.Key("b")
	keyCreator    = client.Key("c")
	keyDelegation = client.Key("d")
	keyRecipient  = client.Key("r")
	keySupply     = client.Key("s")
)

func OnLoad() {
	exports := client.NewScExports()
	exports.AddCall("init", onInit)
	exports.AddView("totalSupply", totalSupply)
	exports.AddView("balanceOf", balanceOf)
	exports.AddView("allowance", allowance)
	exports.AddCall("transfer", transfer)
	exports.AddCall("approve", approve)
	exports.AddCall("transferFrom", transferFrom)
}

// TODO would be awesome to have some less syntactically cumbersome way to check and validate parameters.

// init is a constructor entry point. It initializes the smart contract with the initial value of the token supply
// - input:
//   -- keySupply must be nonzero positive integer
//   -- keyCreator is the AgentID where initial supply is placed
func onInit(ctx *client.ScCallContext) {
	ctx.Log("erc20.init. begin")

	// validate parameters
	// supply
	supply := ctx.Params().GetInt(keySupply)
	if !supply.Exists() || supply.Value() <= 0 {
		ctx.Log("er20.init.fail: wrong 'supply' parameter")
		return
	}
	// creator (owner)
	// we cannot use 'caller' here because the init is always called from the 'root'
	// so, owner of the initial supply must be provided as a parameter keyCreator to constructor (init)
	creator := ctx.Params().GetAgent(keyCreator)
	if !creator.Exists() {
		ctx.Log("er20.init.fail: wrong 'creator' parameter")
		return
	}
	ctx.State().GetInt(keySupply).SetValue(supply.Value())

	// assign the whole supply to creator
	ctx.State().GetMap(keyBalances).GetInt(creator.Value()).SetValue(supply.Value())

	ctx.Log("init.success. Supply = " + ctx.Utility().String(supply.Value()))
	ctx.Log("init.success. Owner = " + creator.Value().String())
}

// the view returns total supply set when creating the contract.
// Output:
// - keySupply: i64
func totalSupply(ctx *client.ScViewContext) {
	supply := ctx.State().GetInt(keySupply).Value()
	ctx.Results().GetInt(keySupply).SetValue(supply)
}

// the view returns balance of the token held in the account
// Input:
// - keyAccount: agentID
func balanceOf(ctx *client.ScViewContext) {
	account := ctx.Params().GetAgent(keyAccount)
	if !account.Exists() {
		ctx.Log("wrong or non existing parameter: " + account.Value().String())
		return
	}
	balances := ctx.State().GetMap(keyBalances)
	balance := balances.GetInt(account.Value()).Value() // 0 if doesn't exist
	ctx.Results().GetInt(keyAmount).SetValue(balance)
}

// the view returns max number of tokens the owner PARAM_ACCOUNT of the account
// allowed to retrieve to another party PARAM_DELEGATION
// Input:
// - keyAccount: agentID
// - keyDelegation: agentID
// Output:
// - keyAmount: i64. 0 if delegation doesn't exists
func allowance(ctx *client.ScViewContext) {
	ctx.Log("erc20.allowance")
	// validate parameters
	// account
	owner := ctx.Params().GetAgent(keyAccount)
	if !owner.Exists() {
		ctx.Log("er20.allowance.fail: wrong 'account' parameter")
		return
	}
	// delegation
	delegation := ctx.Params().GetAgent(keyDelegation)
	if !delegation.Exists() {
		ctx.Log("er20.allowance.fail: wrong 'delegation' parameter")
		return
	}
	// all allowances of the address 'owner' are stored in the map of the same name
	allowances := ctx.State().GetMap(owner.Value())
	allow := allowances.GetInt(delegation.Value()).Value()
	ctx.Results().GetInt(keyAmount).SetValue(allow)
}

// transfer moves tokens from caller's account to target account
// Input:
// - keyAccount: agentID
// - keyAmount: i64
func transfer(ctx *client.ScCallContext) {
	ctx.Log("erc20.transfer")

	// validate params
	params := ctx.Params()
	// account
	account := params.GetAgent(keyAccount)
	if !account.Exists() {
		ctx.Log("er20.transfer.fail: wrong 'account' parameter")
		return
	}
	target_addr := account.Value()
	// amount
	amount := params.GetInt(keyAmount).Value()
	if amount <= 0 {
		ctx.Log("erc20.transfer.fail: wrong 'amount' parameter")
		return
	}
	balances := ctx.State().GetMap(keyBalances)
	source_balance := balances.GetInt(ctx.Caller())

	if source_balance.Value() < amount {
		ctx.Log("erc20.transfer.fail: not enough funds")
		return
	}
	target_balance := balances.GetInt(target_addr)
	result := target_balance.Value() + amount
	if result <= 0 {
		ctx.Log("erc20.transfer.fail: overflow")
		return
	}
	source_balance.SetValue(source_balance.Value() - amount)
	target_balance.SetValue(target_balance.Value() + amount)
	ctx.Log("erc20.transfer.success")
}

// Sets the allowance value for delegated account
// inputs:
//  - keyDelegation: agentID
//  - keyAmount: i64
func approve(ctx *client.ScCallContext) {
	ctx.Log("erc20.approve")

	// validate parameters
	accountPar := ctx.Params().GetAgent(keyDelegation)
	if !accountPar.Exists() {
		ctx.Log("erc20.approve.fail: wrong 'delegation' parameter")
		return
	}
	account := accountPar.Value()
	amount := ctx.Params().GetInt(keyAmount).Value()
	if amount <= 0 {
		ctx.Log("erc20.approve.fail: wrong 'amount' parameter")
		return
	}
	allowances := ctx.State().GetMap(ctx.Caller())
	allowances.GetInt(account).SetValue(amount)
	ctx.Log("erc20.approve.success")
}

// Moves the amount of tokens from sender to recipient using the allowance mechanism.
// Amount is then deducted from the caller’s allowance. This function emits the Transfer event.
// Input:
// - keyAccount: agentID   the spender
// - keyRecipient: agentID   the target
// - keyAmount: i64
func transferFrom(ctx *client.ScCallContext) {
	ctx.Log("erc20.transfer_from")

	// validate parameters
	accountPar := ctx.Params().GetAgent(keyAccount)
	if !accountPar.Exists() {
		ctx.Log("erc20.approve.fail: wrong 'account' parameter")
		return
	}
	account := accountPar.Value()
	recipientPar := ctx.Params().GetAgent(keyRecipient)
	if !recipientPar.Exists() {
		ctx.Log("erc20.approve.fail: wrong 'recipient' parameter")
		return
	}
	recipient := recipientPar.Value()
	amountPar := ctx.Params().GetInt(keyAmount)
	if !amountPar.Exists() {
		ctx.Log("erc20.approve.fail: wrong 'amount' parameter")
		return
	}
	amount := amountPar.Value()

	// allowances are in the map under the name of the account
	allowances := ctx.State().GetMap(account)
	allowance := allowances.GetInt(recipient)
	if allowance.Value() < amount {
		ctx.Log("erc20.approve.fail: not enough allowance")
	}

	balances := ctx.State().GetMap(keyBalances)
	source_balance := balances.GetInt(account)
	if source_balance.Value() < amount {
		ctx.Log("erc20.transfer.fail: not enough funds")
		return
	}
	recipient_balance := balances.GetInt(recipient)
	result := recipient_balance.Value() + amount
	if result <= 0 {
		ctx.Log("erc20.transfer.fail: overflow")
		return
	}
	source_balance.SetValue(source_balance.Value() - amount)
	recipient_balance.SetValue(recipient_balance.Value() + amount)
	allowance.SetValue(allowance.Value() - amount)

	ctx.Log("erc20.transfer_from.success")
}
