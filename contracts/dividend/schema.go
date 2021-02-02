// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package dividend

import "github.com/iotaledger/wasplib/client"

const ScName = "dividend"
const ScHname = client.Hname(0xcce2e239)

const ParamAddress = client.Key("address")
const ParamFactor = client.Key("factor")

const VarMembers = client.Key("members")
const VarTotalFactor = client.Key("total_factor")

const FuncDivide = "divide"
const FuncMember = "member"

const HFuncDivide = client.Hname(0xc7878107)
const HFuncMember = client.Hname(0xc07da2cb)

func OnLoad() {
	exports := client.NewScExports()
	exports.AddCall(FuncDivide, funcDivide)
	exports.AddCall(FuncMember, funcMember)
}
