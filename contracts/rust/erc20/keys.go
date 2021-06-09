// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package erc20

import "github.com/iotaledger/wasplib/packages/vm/wasmlib"

const IdxParamAccount = 0
const IdxParamAmount = 1
const IdxParamCreator = 2
const IdxParamDelegation = 3
const IdxParamRecipient = 4
const IdxParamSupply = 5
const IdxResultAmount = 6
const IdxResultSupply = 7
const IdxVarAllAllowances = 8
const IdxVarBalances = 9
const IdxVarSupply = 10

var keyMap = [11]wasmlib.Key{
	ParamAccount,
	ParamAmount,
	ParamCreator,
	ParamDelegation,
	ParamRecipient,
	ParamSupply,
	ResultAmount,
	ResultSupply,
	VarAllAllowances,
	VarBalances,
	VarSupply,
}

var idxMap [11]wasmlib.Key32
