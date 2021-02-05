// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package fairauction

import "github.com/iotaledger/wasplib/client"

const ScName = "fairauction"
const ScHname = client.ScHname(0x1b5c43b1)

const ParamColor = client.Key("color")
const ParamDescription = client.Key("description")
const ParamDuration = client.Key("duration")
const ParamMinimumBid = client.Key("minimumBid")
const ParamOwnerMargin = client.Key("ownerMargin")

const VarAuctions = client.Key("auctions")
const VarBidderList = client.Key("bidderList")
const VarBidders = client.Key("bidders")
const VarColor = client.Key("color")
const VarCreator = client.Key("creator")
const VarDeposit = client.Key("deposit")
const VarDescription = client.Key("description")
const VarDuration = client.Key("duration")
const VarHighestBid = client.Key("highestBid")
const VarHighestBidder = client.Key("highestBidder")
const VarInfo = client.Key("info")
const VarMinimumBid = client.Key("minimumBid")
const VarNumTokens = client.Key("numTokens")
const VarOwnerMargin = client.Key("ownerMargin")
const VarWhenStarted = client.Key("whenStarted")

const FuncFinalizeAuction = "finalizeAuction"
const FuncPlaceBid = "placeBid"
const FuncSetOwnerMargin = "setOwnerMargin"
const FuncStartAuction = "startAuction"
const ViewGetInfo = "getInfo"

const HFuncFinalizeAuction = client.ScHname(0x8d534ddc)
const HFuncPlaceBid = client.ScHname(0x9bd72fa9)
const HFuncSetOwnerMargin = client.ScHname(0x1774461a)
const HFuncStartAuction = client.ScHname(0xd5b7bacb)
const HViewGetInfo = client.ScHname(0xcfedba5f)

func OnLoad() {
    exports := client.NewScExports()
    exports.AddCall(FuncFinalizeAuction, funcFinalizeAuction)
    exports.AddCall(FuncPlaceBid, funcPlaceBid)
    exports.AddCall(FuncSetOwnerMargin, funcSetOwnerMargin)
    exports.AddCall(FuncStartAuction, funcStartAuction)
    exports.AddView(ViewGetInfo, viewGetInfo)
}
