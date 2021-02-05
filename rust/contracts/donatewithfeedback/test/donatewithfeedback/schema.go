// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package donatewithfeedback

import "github.com/iotaledger/wasplib/client"

const ScName = "donatewithfeedback"
const ScHname = client.ScHname(0x696d7f66)

const ParamAmount = client.Key("amount")
const ParamFeedback = client.Key("feedback")

const VarAmount = client.Key("amount")
const VarDonations = client.Key("donations")
const VarDonator = client.Key("donator")
const VarError = client.Key("error")
const VarFeedback = client.Key("feedback")
const VarLog = client.Key("log")
const VarMaxDonation = client.Key("maxDonation")
const VarTimestamp = client.Key("timestamp")
const VarTotalDonation = client.Key("totalDonation")

const FuncDonate = "donate"
const FuncWithdraw = "withdraw"
const ViewDonations = "donations"

const HFuncDonate = client.ScHname(0xdc9b133a)
const HFuncWithdraw = client.ScHname(0x9dcc0f41)
const HViewDonations = client.ScHname(0x45686a15)

func OnLoad() {
    exports := client.NewScExports()
    exports.AddCall(FuncDonate, funcDonate)
    exports.AddCall(FuncWithdraw, funcWithdraw)
    exports.AddView(ViewDonations, viewDonations)
}
