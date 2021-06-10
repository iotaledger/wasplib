// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package test

import "github.com/iotaledger/wasp/packages/coretypes"

const (
	ScName = "fairroulette"
	HScName = coretypes.Hname(0xdf79d138)

	ParamNumber = "number"
	ParamPlayPeriod = "playPeriod"

	ResultLastWinningNumber = "lastWinningNumber"

	StateBets = "bets"
	StateLastWinningNumber = "lastWinningNumber"
	StateLockedBets = "lockedBets"
	StatePlayPeriod = "playPeriod"

	FuncLockBets = "lockBets"
	FuncPayWinners = "payWinners"
	FuncPlaceBet = "placeBet"
	FuncPlayPeriod = "playPeriod"
	ViewLastWinningNumber = "lastWinningNumber"

	HFuncLockBets = coretypes.Hname(0xe163b43c)
	HFuncPayWinners = coretypes.Hname(0xfb2b0144)
	HFuncPlaceBet = coretypes.Hname(0xdfba7d1b)
	HFuncPlayPeriod = coretypes.Hname(0xcb94b293)
	HViewLastWinningNumber = coretypes.Hname(0x2f5f09fe)
)
