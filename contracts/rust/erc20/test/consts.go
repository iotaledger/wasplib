// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package test

import (
	"github.com/iotaledger/wasp/packages/coretypes"
)

const ScName = "erc20"
const ScDescription = "ERC-20 PoC for IOTA Smart Contracts"
const ScHname = coretypes.Hname(0x200e3733)

const ParamAccount = "ac"
const ParamAmount = "am"
const ParamCreator = "c"
const ParamDelegation = "d"
const ParamRecipient = "r"
const ParamSupply = "s"

const VarBalances = "b"
const VarSupply = "s"

const FuncApprove = "approve"
const FuncInit = "init"
const FuncTransfer = "transfer"
const FuncTransferFrom = "transferFrom"
const ViewAllowance = "allowance"
const ViewBalanceOf = "balanceOf"
const ViewTotalSupply = "totalSupply"

const HFuncApprove = coretypes.Hname(0xa0661268)
const HFuncInit = coretypes.Hname(0x1f44d644)
const HFuncTransfer = coretypes.Hname(0xa15da184)
const HFuncTransferFrom = coretypes.Hname(0xd5e0a602)
const HViewAllowance = coretypes.Hname(0x5e16006a)
const HViewBalanceOf = coretypes.Hname(0x67ef8df4)
const HViewTotalSupply = coretypes.Hname(0x9505e6ca)
