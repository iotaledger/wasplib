// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package test

import "github.com/iotaledger/wasp/packages/coretypes"

const (
	ScName  = "fairauction"
	HScName = coretypes.Hname(0x1b5c43b1)
)

const (
	ParamColor       = "color"
	ParamDescription = "description"
	ParamDuration    = "duration"
	ParamMinimumBid  = "minimumBid"
	ParamOwnerMargin = "ownerMargin"
)

const (
	ResultBidders       = "bidders"
	ResultColor         = "color"
	ResultCreator       = "creator"
	ResultDeposit       = "deposit"
	ResultDescription   = "description"
	ResultDuration      = "duration"
	ResultHighestBid    = "highestBid"
	ResultHighestBidder = "highestBidder"
	ResultMinimumBid    = "minimumBid"
	ResultNumTokens     = "numTokens"
	ResultOwnerMargin   = "ownerMargin"
	ResultWhenStarted   = "whenStarted"
)

const (
	StateAuctions    = "auctions"
	StateBidderList  = "bidderList"
	StateBids        = "bids"
	StateOwnerMargin = "ownerMargin"
)

const (
	FuncFinalizeAuction = "finalizeAuction"
	FuncPlaceBid        = "placeBid"
	FuncSetOwnerMargin  = "setOwnerMargin"
	FuncStartAuction    = "startAuction"
	ViewGetInfo         = "getInfo"
)

const (
	HFuncFinalizeAuction = coretypes.Hname(0x8d534ddc)
	HFuncPlaceBid        = coretypes.Hname(0x9bd72fa9)
	HFuncSetOwnerMargin  = coretypes.Hname(0x1774461a)
	HFuncStartAuction    = coretypes.Hname(0xd5b7bacb)
	HViewGetInfo         = coretypes.Hname(0xcfedba5f)
)
