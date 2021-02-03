// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package fairroulette

import "github.com/iotaledger/wasplib/client"

const ScName = "fairroulette"
const ScHname = client.ScHname(0xdf79d138)

const ParamColor = client.Key("color")
const ParamPlayPeriod = client.Key("play_period")

const VarBets = client.Key("bets")
const VarLastWinningColor = client.Key("last_winning_color")
const VarLockedBets = client.Key("locked_bets")
const VarPlayPeriod = client.Key("play_period")

const FuncLockBets = "lock_bets"
const FuncPayWinners = "pay_winners"
const FuncPlaceBet = "place_bet"
const FuncPlayPeriod = "play_period"

const HFuncLockBets = client.ScHname(0x853da2a7)
const HFuncPayWinners = client.ScHname(0x3df139de)
const HFuncPlaceBet = client.ScHname(0x575b51d2)
const HFuncPlayPeriod = client.ScHname(0xf534dac1)

func OnLoad() {
    exports := client.NewScExports()
    exports.AddCall(FuncLockBets, funcLockBets)
    exports.AddCall(FuncPayWinners, funcPayWinners)
    exports.AddCall(FuncPlaceBet, funcPlaceBet)
    exports.AddCall(FuncPlayPeriod, funcPlayPeriod)
}
