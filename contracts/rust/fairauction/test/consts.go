// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package test

import "github.com/iotaledger/wasplib/packages/vm/wasmlib"

const ScName = "fairauction"
const HScName = wasmlib.ScHname(0x1b5c43b1)

const ParamColor = "color"
const ParamDescription = "description"
const ParamDuration = "duration"
const ParamMinimumBid = "minimumBid"
const ParamOwnerMargin = "ownerMargin"

const ResultBidders = "bidders"
const ResultColor = "color"
const ResultCreator = "creator"
const ResultDeposit = "deposit"
const ResultDescription = "description"
const ResultDuration = "duration"
const ResultHighestBid = "highestBid"
const ResultHighestBidder = "highestBidder"
const ResultMinimumBid = "minimumBid"
const ResultNumTokens = "numTokens"
const ResultOwnerMargin = "ownerMargin"
const ResultWhenStarted = "whenStarted"

const VarAuctions = "auctions"
const VarBidderList = "bidderList"
const VarBids = "bids"
const VarOwnerMargin = "ownerMargin"

const FuncFinalizeAuction = "finalizeAuction"
const FuncPlaceBid = "placeBid"
const FuncSetOwnerMargin = "setOwnerMargin"
const FuncStartAuction = "startAuction"
const ViewGetInfo = "getInfo"

const HFuncFinalizeAuction = wasmlib.ScHname(0x8d534ddc)
const HFuncPlaceBid = wasmlib.ScHname(0x9bd72fa9)
const HFuncSetOwnerMargin = wasmlib.ScHname(0x1774461a)
const HFuncStartAuction = wasmlib.ScHname(0xd5b7bacb)
const HViewGetInfo = wasmlib.ScHname(0xcfedba5f)
