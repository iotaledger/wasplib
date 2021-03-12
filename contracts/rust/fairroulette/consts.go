// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package fairroulette

import "github.com/iotaledger/wasplib/packages/vm/wasmlib"

const ScName = "fairroulette"
const HScName = wasmlib.ScHname(0xdf79d138)

const ParamNumber = wasmlib.Key("number")
const ParamPlayPeriod = wasmlib.Key("playPeriod")

const VarBets = wasmlib.Key("bets")
const VarLastWinningNumber = wasmlib.Key("lastWinningNumber")
const VarLockedBets = wasmlib.Key("lockedBets")
const VarPlayPeriod = wasmlib.Key("playPeriod")

const FuncLockBets = "lockBets"
const FuncPayWinners = "payWinners"
const FuncPlaceBet = "placeBet"
const FuncPlayPeriod = "playPeriod"
const ViewLastWinningNumber = "lastWinningNumber"

const HFuncLockBets = wasmlib.ScHname(0xe163b43c)
const HFuncPayWinners = wasmlib.ScHname(0xfb2b0144)
const HFuncPlaceBet = wasmlib.ScHname(0xdfba7d1b)
const HFuncPlayPeriod = wasmlib.ScHname(0xcb94b293)
const HViewLastWinningNumber = wasmlib.ScHname(0x2f5f09fe)
