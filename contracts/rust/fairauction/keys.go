// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package fairauction

import "github.com/iotaledger/wasplib/packages/vm/wasmlib"

const IdxParamColor = 0
const IdxParamDescription = 1
const IdxParamDuration = 2
const IdxParamMinimumBid = 3
const IdxParamOwnerMargin = 4
const IdxResultBidders = 5
const IdxResultColor = 6
const IdxResultCreator = 7
const IdxResultDeposit = 8
const IdxResultDescription = 9
const IdxResultDuration = 10
const IdxResultHighestBid = 11
const IdxResultHighestBidder = 12
const IdxResultMinimumBid = 13
const IdxResultNumTokens = 14
const IdxResultOwnerMargin = 15
const IdxResultWhenStarted = 16
const IdxVarAuctions = 17
const IdxVarBidderList = 18
const IdxVarBids = 19
const IdxVarOwnerMargin = 20

var keyMap = [21]string{
	ParamColor,
	ParamDescription,
	ParamDuration,
	ParamMinimumBid,
	ParamOwnerMargin,
	ResultBidders,
	ResultColor,
	ResultCreator,
	ResultDeposit,
	ResultDescription,
	ResultDuration,
	ResultHighestBid,
	ResultHighestBidder,
	ResultMinimumBid,
	ResultNumTokens,
	ResultOwnerMargin,
	ResultWhenStarted,
	VarAuctions,
	VarBidderList,
	VarBids,
	VarOwnerMargin,
}

var idxMap [21]wasmlib.Key32
