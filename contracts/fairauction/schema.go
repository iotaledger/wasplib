// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package fairauction

import "github.com/iotaledger/wasplib/client"

const ScName = "fairauction"
const ScHname = client.ScHname(0x1b5c43b1)

const ParamColor = client.Key("color")
const ParamDescription = client.Key("description")
const ParamDuration = client.Key("duration")
const ParamMinimumBid = client.Key("minimum")
const ParamOwnerMargin = client.Key("owner_margin")

const VarAuctions = client.Key("auctions")
const VarBidderList = client.Key("bidder_list")
const VarBidders = client.Key("bidders")
const VarColor = client.Key("color")
const VarCreator = client.Key("creator")
const VarDeposit = client.Key("deposit")
const VarDescription = client.Key("description")
const VarDuration = client.Key("duration")
const VarHighestBid = client.Key("highest_bid")
const VarHighestBidder = client.Key("highest_bidder")
const VarInfo = client.Key("info")
const VarMinimumBid = client.Key("minimum")
const VarNumTokens = client.Key("num_tokens")
const VarOwnerMargin = client.Key("owner_margin")
const VarWhenStarted = client.Key("when_started")

const FuncFinalizeAuction = "finalize_auction"
const FuncPlaceBid = "place_bid"
const FuncSetOwnerMargin = "set_owner_margin"
const FuncStartAuction = "start_auction"
const ViewGetInfo = "get_info"

const HFuncFinalizeAuction = client.ScHname(0xb427dd28)
const HFuncPlaceBid = client.ScHname(0xf2cc1c44)
const HFuncSetOwnerMargin = client.ScHname(0x65402dca)
const HFuncStartAuction = client.ScHname(0x7ee53d08)
const HViewGetInfo = client.ScHname(0x2b9d8867)

func OnLoad() {
    exports := client.NewScExports()
    exports.AddCall(FuncFinalizeAuction, funcFinalizeAuction)
    exports.AddCall(FuncPlaceBid, funcPlaceBid)
    exports.AddCall(FuncSetOwnerMargin, funcSetOwnerMargin)
    exports.AddCall(FuncStartAuction, funcStartAuction)
    exports.AddView(ViewGetInfo, viewGetInfo)
}
