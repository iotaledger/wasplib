// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// implementation of ERC-20 smart contract for ISCP
// following https://ethereum.org/en/developers/tutorials/understand-the-erc-20-token-smart-contract/

package erc20

import "github.com/iotaledger/wasp/packages/vm/wasmlib"

// Sets the allowance value for delegated account
// inputs:
//  - PARAM_DELEGATION: agentID
//  - PARAM_AMOUNT: i64
func funcApprove(ctx *wasmlib.ScFuncContext, params *FuncApproveParams) {
	ctx.Trace("erc20.approve")

	delegation := params.Delegation.Value()
	amount := params.Amount.Value()
	ctx.Require(amount > 0, "erc20.approve.fail: wrong 'amount' parameter")

	// all allowances are in the map under the name of he owner
	allowances := ctx.State().GetMap(ctx.Caller())
	allowances.GetInt(delegation).SetValue(amount)
	ctx.Log("erc20.approve.success")
}

// on_init is a constructor entry point. It initializes the smart contract with the
// initial value of the token supply and the owner of that supply
// - input:
//   -- PARAM_SUPPLY must be nonzero positive integer. Mandatory
//   -- PARAM_CREATOR is the AgentID where initial supply is placed. Mandatory
func funcInit(ctx *wasmlib.ScFuncContext, params *FuncInitParams) {
	ctx.Trace("erc20.on_init.begin")

	supply := params.Supply.Value()
	ctx.Require(supply > 0, "erc20.on_init.fail: wrong 'supply' parameter")
	ctx.State().GetInt(VarSupply).SetValue(supply)

	// we cannot use 'caller' here because on_init is always called from the 'root'
	// so, creator/owner of the initial supply must be provided as a parameter PARAM_CREATOR to constructor (on_init)
	// assign the whole supply to creator
	creator := params.Creator.Value()
	ctx.State().GetMap(VarBalances).GetInt(creator).SetValue(supply)

	t := "erc20.on_init.success. Supply: " + params.Supply.String() +
		", creator:" + params.Creator.String()
	ctx.Log(t)
}

// transfer moves tokens from caller's account to target account
// Input:
// - PARAM_ACCOUNT: agentID
// - PARAM_AMOUNT: i64
func funcTransfer(ctx *wasmlib.ScFuncContext, params *FuncTransferParams) {
	ctx.Trace("erc20.transfer")

	amount := params.Amount.Value()
	ctx.Require(amount > 0, "erc20.transfer.fail: wrong 'amount' parameter")

	balances := ctx.State().GetMap(VarBalances)
	sourceBalance := balances.GetInt(ctx.Caller())
	ctx.Require(sourceBalance.Value() >= amount, "erc20.transfer.fail: not enough funds")

	targetAddr := params.Account.Value()
	targetBalance := balances.GetInt(targetAddr)
	result := targetBalance.Value() + amount
	ctx.Require(result > 0, "erc20.transfer.fail: overflow")

	sourceBalance.SetValue(sourceBalance.Value() - amount)
	targetBalance.SetValue(targetBalance.Value() + amount)
	ctx.Log("erc20.transfer.success")
}

// Moves the amount of tokens from sender to recipient using the allowance mechanism.
// Amount is then deducted from the caller’s allowance. This function emits the Transfer event.
// Input:
// - PARAM_ACCOUNT: agentID   the spender
// - PARAM_RECIPIENT: agentID   the target
// - PARAM_AMOUNT: i64
func funcTransferFrom(ctx *wasmlib.ScFuncContext, params *FuncTransferFromParams) {
	ctx.Trace("erc20.transfer_from")

	account := params.Account.Value()
	recipient := params.Recipient.Value()
	amount := params.Amount.Value()
	ctx.Require(amount > 0, "erc20.transfer_from.fail: wrong 'amount' parameter")

	// allowances are in the map under the name of the account
	allowances := ctx.State().GetMap(account)
	allowance := allowances.GetInt(recipient)
	ctx.Require(allowance.Value() >= amount, "erc20.transfer_from.fail: not enough allowance")

	balances := ctx.State().GetMap(VarBalances)
	sourceBalance := balances.GetInt(account)
	ctx.Require(sourceBalance.Value() >= amount, "erc20.transfer_from.fail: not enough funds")

	recipientBalance := balances.GetInt(recipient)
	result := recipientBalance.Value() + amount
	ctx.Require(result > 0, "erc20.transfer_from.fail: overflow")

	sourceBalance.SetValue(sourceBalance.Value() - amount)
	recipientBalance.SetValue(recipientBalance.Value() + amount)
	allowance.SetValue(allowance.Value() - amount)

	ctx.Log("erc20.transfer_from.success")
}

// the view returns max number of tokens the owner PARAM_ACCOUNT of the account
// allowed to retrieve to another party PARAM_DELEGATION
// Input:
// - PARAM_ACCOUNT: agentID
// - PARAM_DELEGATION: agentID
// Output:
// - PARAM_AMOUNT: i64
func viewAllowance(ctx *wasmlib.ScViewContext, params *ViewAllowanceParams) {
	ctx.Trace("erc20.allowance")

	// all allowances of the address 'owner' are stored in the map of the same name
	allowances := ctx.State().GetMap(params.Account.Value())
	allow := allowances.GetInt(params.Delegation.Value()).Value()
	ctx.Results().GetInt(ParamAmount).SetValue(allow)
}

// the view returns balance of the token held in the account
// Input:
// - PARAM_ACCOUNT: agentID
func viewBalanceOf(ctx *wasmlib.ScViewContext, params *ViewBalanceOfParams) {
	balances := ctx.State().GetMap(VarBalances)
	balance := balances.GetInt(params.Account.Value()).Value()
	ctx.Results().GetInt(ParamAmount).SetValue(balance)
}

// the view returns total supply set when creating the contract (a constant).
// Output:
// - PARAM_SUPPLY: i64
func viewTotalSupply(ctx *wasmlib.ScViewContext, params *ViewTotalSupplyParams) {
	supply := ctx.State().GetInt(VarSupply).Value()
	ctx.Results().GetInt(ParamSupply).SetValue(supply)
}
