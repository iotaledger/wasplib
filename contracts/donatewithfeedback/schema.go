// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package donatewithfeedback

import "github.com/iotaledger/wasplib/client"

const ScName = "donatewithfeedback"
const ScHname = client.ScHname(0x696d7f66)

const ParamFeedback = client.Key("feedback")
const ParamWithdrawAmount = client.Key("withdraw")

const VarAmount = client.Key("amount")
const VarDonations = client.Key("donations")
const VarDonator = client.Key("donator")
const VarError = client.Key("error")
const VarFeedback = client.Key("feedback")
const VarLog = client.Key("log")
const VarMaxDonation = client.Key("max_donation")
const VarTimestamp = client.Key("timestamp")
const VarTotalDonation = client.Key("total_donation")

const FuncDonate = "donate"
const FuncWithdraw = "withdraw"
const ViewDonations = "view_donations"

const HFuncDonate = client.ScHname(0xdc9b133a)
const HFuncWithdraw = client.ScHname(0x9dcc0f41)
const HViewDonations = client.ScHname(0xc3cc7cb0)

func OnLoad() {
    exports := client.NewScExports()
    exports.AddCall(FuncDonate, funcDonate)
    exports.AddCall(FuncWithdraw, funcWithdraw)
    exports.AddView(ViewDonations, viewDonations)
}
