// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// implementation of ERC-20 smart contract for ISCP
// following https://ethereum.org/en/developers/tutorials/understand-the-erc-20-token-smart-contract/

package erc20

import "github.com/iotaledger/wasplib/client"

// state variable
const StateVarSupply = client.Key("s")

// supply constant
const StateVarBalances = client.Key("b") // name of the map of balances

// params and return variables, used in calls
const ParamSupply = client.Key("s")
const ParamCreator = client.Key("c")
const ParamAccount = client.Key("ac")
const ParamDelegation = client.Key("d")
const ParamAmount = client.Key("am")
const ParamRecipient = client.Key("r")

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
	ctx.Log("erc20.OnInit.Begin")
	// validate parameters
	// supply
	supply := ctx.Params().GetInt(ParamSupply)
	if !supply.Exists() || supply.Value() <= 0 {
		ctx.Panic("erc20.OnInit.Fail: wrong 'supply' parameter")
	}
	// creator (owner)
	// we cannot use 'caller' here because on_init is always called from the 'root'
	// so, owner of the initial supply must be provided as a parameter PARAM_CREATOR to constructor (on_init)
	creator := ctx.Params().GetAgent(ParamCreator)
	if !creator.Exists() {
		ctx.Panic("erc20.OnInit.Fail: wrong 'creator' parameter")
	}
	ctx.State().GetInt(StateVarSupply).SetValue(supply.Value())

	// assign the whole supply to creator
	ctx.State().GetMap(StateVarBalances).GetInt(creator.Value()).SetValue(supply.Value())

	t := "erc20.OnInit.Success. Supply: " + supply.String() +
		", creator:" + creator.String()
	ctx.Log(t)
}

// the view returns total supply set when creating the contract (a constant).
// Output:
// - PARAM_SUPPLY: i64
func totalSupply(ctx *client.ScViewContext) {
	supply := ctx.State().GetInt(StateVarSupply).Value()
	ctx.Results().GetInt(ParamSupply).SetValue(supply)
}

// the view returns balance of the token held in the account
// Input:
// - PARAM_ACCOUNT: agentID
func balanceOf(ctx *client.ScViewContext) {
	account := ctx.Params().GetAgent(ParamAccount)
	if !account.Exists() {
		ctx.Panic("wrong or non existing parameter: " + account.String())
	}
	balances := ctx.State().GetMap(StateVarBalances)
	balance := balances.GetInt(account.Value()).Value() // 0 if doesn't exist
	ctx.Results().GetInt(ParamAmount).SetValue(balance)
}

// the view returns max number of tokens the owner PARAM_ACCOUNT of the account
// allowed to retrieve to another party PARAM_DELEGATION
// Input:
// - PARAM_ACCOUNT: agentID
// - PARAM_DELEGATION: agentID
// Output:
// - PARAM_AMOUNT: i64. 0 if delegation doesn't exists
func allowance(ctx *client.ScViewContext) {
	ctx.Log("erc20.Allowance")
	// validate parameters
	// account
	owner := ctx.Params().GetAgent(ParamAccount)
	if !owner.Exists() {
		ctx.Panic("erc20.Allowance.Fail: wrong 'account' parameter")
	}
	// delegation
	delegation := ctx.Params().GetAgent(ParamDelegation)
	if !delegation.Exists() {
		ctx.Panic("erc20.Allowance.Fail: wrong 'delegation' parameter")
	}
	// all allowances of the address 'owner' are stored in the map of the same name
	allowances := ctx.State().GetMap(owner.Value())
	allow := allowances.GetInt(delegation.Value()).Value()
	ctx.Results().GetInt(ParamAmount).SetValue(allow)
}

// transfer moves tokens from caller's account to target account
// Input:
// - PARAM_ACCOUNT: agentID
// - PARAM_AMOUNT: i64
func transfer(ctx *client.ScCallContext) {
	ctx.Log("erc20.Transfer")

	// validate params
	params := ctx.Params()
	// account
	targetAddrParam := params.GetAgent(ParamAccount)
	if !targetAddrParam.Exists() {
		ctx.Panic("erc20.Transfer.Fail: wrong 'account' parameter")
	}
	targetAddr := targetAddrParam.Value()
	// amount
	amount := params.GetInt(ParamAmount).Value()
	if amount <= 0 {
		ctx.Panic("erc20.Transfer.Fail: wrong 'amount' parameter")
	}
	balances := ctx.State().GetMap(StateVarBalances)
	sourceBalance := balances.GetInt(ctx.Caller())

	if sourceBalance.Value() < amount {
		ctx.Panic("erc20.Transfer.Fail: not enough funds")
	}
	targetBalance := balances.GetInt(targetAddr)
	result := targetBalance.Value() + amount
	if result <= 0 {
		ctx.Panic("erc20.Transfer.Fail: overflow")
	}
	sourceBalance.SetValue(sourceBalance.Value() - amount)
	targetBalance.SetValue(targetBalance.Value() + amount)
	ctx.Log("erc20.Transfer.Success")
}

// Sets the allowance value for delegated account
// inputs:
//  - PARAM_DELEGATION: agentID
//  - PARAM_AMOUNT: i64
func approve(ctx *client.ScCallContext) {
	ctx.Log("erc20.Approve")

	// validate parameters
	delegationParam := ctx.Params().GetAgent(ParamDelegation)
	if !delegationParam.Exists() {
		ctx.Panic("erc20.Approve.Fail: wrong 'delegation' parameter")
	}
	delegation := delegationParam.Value()
	amount := ctx.Params().GetInt(ParamAmount).Value()
	if amount <= 0 {
		ctx.Panic("erc20.Approve.Fail: wrong 'amount' parameter")
	}
	// all allowances are in the map under the name of he owner
	allowances := ctx.State().GetMap(ctx.Caller())
	allowances.GetInt(delegation).SetValue(amount)
	ctx.Log("erc20.Approve.Success")
}

// Moves the amount of tokens from sender to recipient using the allowance mechanism.
// Amount is then deducted from the caller’s allowance. This function emits the Transfer event.
// Input:
// - PARAM_ACCOUNT: agentID   the spender
// - PARAM_RECIPIENT: agentID   the target
// - PARAM_AMOUNT: i64
func transferFrom(ctx *client.ScCallContext) {
	ctx.Log("erc20.TransferFrom")

	// validate parameters
	accountParam := ctx.Params().GetAgent(ParamAccount)
	if !accountParam.Exists() {
		ctx.Panic("erc20.TransferFrom.Fail: wrong 'account' parameter")
	}
	account := accountParam.Value()
	recipientParam := ctx.Params().GetAgent(ParamRecipient)
	if !recipientParam.Exists() {
		ctx.Panic("erc20.TransferFrom.Fail: wrong 'recipient' parameter")
	}
	recipient := recipientParam.Value()
	amountParam := ctx.Params().GetInt(ParamAmount)
	if !amountParam.Exists() {
		ctx.Panic("erc20.TransferFrom.Fail: wrong 'amount' parameter")
	}
	amount := amountParam.Value()

	// allowances are in the map under the name of the account
	allowances := ctx.State().GetMap(account)
	allowance := allowances.GetInt(recipient)
	if allowance.Value() < amount {
		ctx.Panic("erc20.TransferFrom.Fail: not enough allowance")
	}
	balances := ctx.State().GetMap(StateVarBalances)
	sourceBalance := balances.GetInt(account)
	if sourceBalance.Value() < amount {
		ctx.Panic("erc20.TransferFrom.Fail: not enough funds")
	}
	recipientBalance := balances.GetInt(recipient)
	result := recipientBalance.Value() + amount
	if result <= 0 {
		ctx.Panic("erc20.TransferFrom.Fail: overflow")
	}
	sourceBalance.SetValue(sourceBalance.Value() - amount)
	recipientBalance.SetValue(recipientBalance.Value() + amount)
	allowance.SetValue(allowance.Value() - amount)

	ctx.Log("erc20.TransferFrom.Success")
}
