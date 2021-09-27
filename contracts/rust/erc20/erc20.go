// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// implementation of ERC-20 smart contract for ISCP
// following https:// ethereum.org/en/developers/tutorials/understand-the-erc-20-token-smart-contract/

package erc20

import (
	"github.com/iotaledger/wasp/packages/vm/wasmlib"
)

// Sets the allowance value for delegated account
// inputs:
//  - PARAM_DELEGATION: agentID
//  - PARAM_AMOUNT: i64
func funcApprove(ctx wasmlib.ScFuncContext, f *ApproveContext) {
	delegation := f.Params.Delegation().Value()
	amount := f.Params.Amount().Value()
	ctx.Require(amount > 0, "erc20.approve.fail: wrong 'amount' parameter")

	// all allowances are in the map under the name of he owner
	allowances := f.State.AllAllowances().GetAllowancesForAgent(ctx.Caller())
	allowances.GetInt64(delegation).SetValue(amount)
}

// on_init is a constructor entry point. It initializes the smart contract with the
// initial value of the token supply and the owner of that supply
// - input:
//   -- PARAM_SUPPLY must be nonzero positive integer. Mandatory
//   -- PARAM_CREATOR is the AgentID where initial supply is placed. Mandatory
func funcInit(ctx wasmlib.ScFuncContext, f *InitContext) {
	supply := f.Params.Supply().Value()
	ctx.Require(supply > 0, "erc20.on_init.fail: wrong 'supply' parameter")
	f.State.Supply().SetValue(supply)

	// we cannot use 'caller' here because on_init is always called from the 'root'
	// so, owner of the initial supply must be provided as a parameter PARAM_CREATOR to constructor (on_init)
	// assign the whole supply to creator
	creator := f.Params.Creator().Value()
	f.State.Balances().GetInt64(creator).SetValue(supply)

	t := "erc20.on_init.success. Supply: " + f.Params.Supply().String() +
		", creator:" + creator.String()
	ctx.Log(t)
}

// transfer moves tokens from caller's account to target account
// Input:
// - PARAM_ACCOUNT: agentID
// - PARAM_AMOUNT: i64
func funcTransfer(ctx wasmlib.ScFuncContext, f *TransferContext) {
	amount := f.Params.Amount().Value()
	ctx.Require(amount > 0, "erc20.transfer.fail: wrong 'amount' parameter")

	balances := f.State.Balances()
	sourceBalance := balances.GetInt64(ctx.Caller())
	ctx.Require(sourceBalance.Value() >= amount, "erc20.transfer.fail: not enough funds")

	targetAddr := f.Params.Account().Value()
	targetBalance := balances.GetInt64(targetAddr)
	result := targetBalance.Value() + amount
	ctx.Require(result > 0, "erc20.transfer.fail: overflow")

	sourceBalance.SetValue(sourceBalance.Value() - amount)
	targetBalance.SetValue(targetBalance.Value() + amount)
}

// Moves the amount of tokens from sender to recipient using the allowance mechanism.
// Amount is then deducted from the caller’s allowance. This function emits the Transfer event.
// Input:
// - PARAM_ACCOUNT: agentID   the spender
// - PARAM_RECIPIENT: agentID   the target
// - PARAM_AMOUNT: i64
func funcTransferFrom(ctx wasmlib.ScFuncContext, f *TransferFromContext) {
	// validate parameters
	account := f.Params.Account().Value()
	recipient := f.Params.Recipient().Value()
	amount := f.Params.Amount().Value()
	ctx.Require(amount > 0, "erc20.transfer_from.fail: wrong 'amount' parameter")

	// allowances are in the map under the name of the account
	allowances := f.State.AllAllowances().GetAllowancesForAgent(account)
	allowance := allowances.GetInt64(recipient)
	ctx.Require(allowance.Value() >= amount, "erc20.transfer_from.fail: not enough allowance")

	balances := f.State.Balances()
	sourceBalance := balances.GetInt64(account)
	ctx.Require(sourceBalance.Value() >= amount, "erc20.transfer_from.fail: not enough funds")

	recipientBalance := balances.GetInt64(recipient)
	result := recipientBalance.Value() + amount
	ctx.Require(result > 0, "erc20.transfer_from.fail: overflow")

	sourceBalance.SetValue(sourceBalance.Value() - amount)
	recipientBalance.SetValue(recipientBalance.Value() + amount)
	allowance.SetValue(allowance.Value() - amount)
}

// the view returns max number of tokens the owner PARAM_ACCOUNT of the account
// allowed to retrieve to another party PARAM_DELEGATION
// Input:
// - PARAM_ACCOUNT: agentID
// - PARAM_DELEGATION: agentID
// Output:
// - PARAM_AMOUNT: i64
func viewAllowance(ctx wasmlib.ScViewContext, f *AllowanceContext) {
	// all allowances of the address 'owner' are stored in the map of the same name
	allowances := f.State.AllAllowances().GetAllowancesForAgent(f.Params.Account().Value())
	allow := allowances.GetInt64(f.Params.Delegation().Value()).Value()
	f.Results.Amount().SetValue(allow)
}

// the view returns balance of the token held in the account
// Input:
// - PARAM_ACCOUNT: agentID
func viewBalanceOf(ctx wasmlib.ScViewContext, f *BalanceOfContext) {
	balances := f.State.Balances()
	balance := balances.GetInt64(f.Params.Account().Value())
	f.Results.Amount().SetValue(balance.Value())
}

// the view returns total supply set when creating the contract (a constant).
// Output:
// - PARAM_SUPPLY: i64
func viewTotalSupply(ctx wasmlib.ScViewContext, f *TotalSupplyContext) {
	f.Results.Supply().SetValue(f.State.Supply().Value())
}
