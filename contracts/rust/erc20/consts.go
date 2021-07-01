// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package erc20

import "github.com/iotaledger/wasplib/packages/vm/wasmlib"

const (
	ScName        = "erc20"
	ScDescription = "ERC-20 PoC for IOTA Smart Contracts"
	HScName       = wasmlib.ScHname(0x200e3733)
)

const (
	ParamAccount    = wasmlib.Key("ac")
	ParamAmount     = wasmlib.Key("am")
	ParamCreator    = wasmlib.Key("c")
	ParamDelegation = wasmlib.Key("d")
	ParamRecipient  = wasmlib.Key("r")
	ParamSupply     = wasmlib.Key("s")
)

const (
	ResultAmount = wasmlib.Key("am")
	ResultSupply = wasmlib.Key("s")
)

const (
	StateAllAllowances = wasmlib.Key("a")
	StateBalances      = wasmlib.Key("b")
	StateSupply        = wasmlib.Key("s")
)

const (
	FuncApprove      = "approve"
	FuncInit         = "init"
	FuncTransfer     = "transfer"
	FuncTransferFrom = "transferFrom"
	ViewAllowance    = "allowance"
	ViewBalanceOf    = "balanceOf"
	ViewTotalSupply  = "totalSupply"
)

const (
	HFuncApprove      = wasmlib.ScHname(0xa0661268)
	HFuncInit         = wasmlib.ScHname(0x1f44d644)
	HFuncTransfer     = wasmlib.ScHname(0xa15da184)
	HFuncTransferFrom = wasmlib.ScHname(0xd5e0a602)
	HViewAllowance    = wasmlib.ScHname(0x5e16006a)
	HViewBalanceOf    = wasmlib.ScHname(0x67ef8df4)
	HViewTotalSupply  = wasmlib.ScHname(0x9505e6ca)
)
