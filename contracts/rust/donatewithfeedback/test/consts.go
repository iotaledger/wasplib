// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package test

import (
	"github.com/iotaledger/wasp/packages/coretypes"
)

const ScName = "donatewithfeedback"
const ScHname = coretypes.Hname(0x696d7f66)

const ParamAmount = "amount"
const ParamFeedback = "feedback"

const VarAmount = "amount"
const VarDonations = "donations"
const VarDonator = "donator"
const VarError = "error"
const VarFeedback = "feedback"
const VarLog = "log"
const VarMaxDonation = "maxDonation"
const VarTimestamp = "timestamp"
const VarTotalDonation = "totalDonation"

const FuncDonate = "donate"
const FuncWithdraw = "withdraw"
const ViewDonations = "donations"

const HFuncDonate = coretypes.Hname(0xdc9b133a)
const HFuncWithdraw = coretypes.Hname(0x9dcc0f41)
const HViewDonations = coretypes.Hname(0x45686a15)
