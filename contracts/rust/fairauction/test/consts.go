// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package test

import (
	"github.com/iotaledger/wasp/packages/coretypes"
)

const ScName = "fairauction"
const ScHname = coretypes.Hname(0x1b5c43b1)

const ParamColor = "color"
const ParamDescription = "description"
const ParamDuration = "duration"
const ParamMinimumBid = "minimumBid"
const ParamOwnerMargin = "ownerMargin"

const VarAuctions = "auctions"
const VarBidderList = "bidderList"
const VarBidders = "bidders"
const VarColor = "color"
const VarCreator = "creator"
const VarDeposit = "deposit"
const VarDescription = "description"
const VarDuration = "duration"
const VarHighestBid = "highestBid"
const VarHighestBidder = "highestBidder"
const VarInfo = "info"
const VarMinimumBid = "minimumBid"
const VarNumTokens = "numTokens"
const VarOwnerMargin = "ownerMargin"
const VarWhenStarted = "whenStarted"

const FuncFinalizeAuction = "finalizeAuction"
const FuncPlaceBid = "placeBid"
const FuncSetOwnerMargin = "setOwnerMargin"
const FuncStartAuction = "startAuction"
const ViewGetInfo = "getInfo"

const HFuncFinalizeAuction = coretypes.Hname(0x8d534ddc)
const HFuncPlaceBid = coretypes.Hname(0x9bd72fa9)
const HFuncSetOwnerMargin = coretypes.Hname(0x1774461a)
const HFuncStartAuction = coretypes.Hname(0xd5b7bacb)
const HViewGetInfo = coretypes.Hname(0xcfedba5f)
