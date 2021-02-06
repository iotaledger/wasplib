// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package _go

import "github.com/iotaledger/wasplib/client"

const ScName = "dividend"
const ScHname = client.ScHname(0xcce2e239)

const ParamAddress = client.Key("address")
const ParamFactor = client.Key("factor")

const VarMembers = client.Key("members")
const VarTotalFactor = client.Key("totalFactor")

const FuncDivide = "divide"
const FuncMember = "member"

const HFuncDivide = client.ScHname(0xc7878107)
const HFuncMember = client.ScHname(0xc07da2cb)
