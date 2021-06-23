// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package donatewithfeedback

import "github.com/iotaledger/wasplib/packages/vm/wasmlib"

const (
	ScName  = "donatewithfeedback"
	HScName = wasmlib.ScHname(0x696d7f66)
)

const (
	ParamAmount   = wasmlib.Key("amount")
	ParamFeedback = wasmlib.Key("feedback")
	ParamNr       = wasmlib.Key("nr")
)

const (
	ResultAmount        = wasmlib.Key("amount")
	ResultCount         = wasmlib.Key("count")
	ResultDonator       = wasmlib.Key("donator")
	ResultError         = wasmlib.Key("error")
	ResultFeedback      = wasmlib.Key("feedback")
	ResultMaxDonation   = wasmlib.Key("maxDonation")
	ResultTimestamp     = wasmlib.Key("timestamp")
	ResultTotalDonation = wasmlib.Key("totalDonation")
)

const (
	StateLog           = wasmlib.Key("log")
	StateMaxDonation   = wasmlib.Key("maxDonation")
	StateTotalDonation = wasmlib.Key("totalDonation")
)

const (
	FuncDonate       = "donate"
	FuncWithdraw     = "withdraw"
	ViewDonation     = "donation"
	ViewDonationInfo = "donationInfo"
)

const (
	HFuncDonate       = wasmlib.ScHname(0xdc9b133a)
	HFuncWithdraw     = wasmlib.ScHname(0x9dcc0f41)
	HViewDonation     = wasmlib.ScHname(0xbdb245ba)
	HViewDonationInfo = wasmlib.ScHname(0xc8f7c726)
)
