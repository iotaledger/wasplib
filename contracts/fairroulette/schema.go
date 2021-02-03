// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package fairroulette

import "github.com/iotaledger/wasplib/client"

const ScName = "fairroulette"
const ScHname = client.ScHname(0xdf79d138)

const ParamNumber = client.Key("number")
const ParamPlayPeriod = client.Key("playPeriod")

const VarBets = client.Key("bets")
const VarLastWinningNumber = client.Key("lastWinningNumber")
const VarLockedBets = client.Key("lockedBets")
const VarPlayPeriod = client.Key("playPeriod")

const FuncLockBets = "lockBets"
const FuncPayWinners = "payWinners"
const FuncPlaceBet = "placeBet"
const FuncPlayPeriod = "playPeriod"

const HFuncLockBets = client.ScHname(0xe163b43c)
const HFuncPayWinners = client.ScHname(0xfb2b0144)
const HFuncPlaceBet = client.ScHname(0xdfba7d1b)
const HFuncPlayPeriod = client.ScHname(0xcb94b293)

func OnLoad() {
    exports := client.NewScExports()
    exports.AddCall(FuncLockBets, funcLockBets)
    exports.AddCall(FuncPayWinners, funcPayWinners)
    exports.AddCall(FuncPlaceBet, funcPlaceBet)
    exports.AddCall(FuncPlayPeriod, funcPlayPeriod)
}
