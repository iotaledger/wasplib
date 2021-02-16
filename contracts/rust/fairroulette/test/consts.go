// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package test

import (
	"github.com/iotaledger/wasp/packages/coretypes"
)

const ScName = "fairroulette"
const ScHname = coretypes.Hname(0xdf79d138)

const ParamNumber = "number"
const ParamPlayPeriod = "playPeriod"

const VarBets = "bets"
const VarLastWinningNumber = "lastWinningNumber"
const VarLockedBets = "lockedBets"
const VarPlayPeriod = "playPeriod"

const FuncLockBets = "lockBets"
const FuncPayWinners = "payWinners"
const FuncPlaceBet = "placeBet"
const FuncPlayPeriod = "playPeriod"

const HFuncLockBets = coretypes.Hname(0xe163b43c)
const HFuncPayWinners = coretypes.Hname(0xfb2b0144)
const HFuncPlaceBet = coretypes.Hname(0xdfba7d1b)
const HFuncPlayPeriod = coretypes.Hname(0xcb94b293)
