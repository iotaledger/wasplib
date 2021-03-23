// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package dividend

import "github.com/iotaledger/wasplib/packages/vm/wasmlib"

const ScName = "dividend"
const HScName = wasmlib.ScHname(0xcce2e239)

const ParamAddress = wasmlib.Key("address")
const ParamFactor = wasmlib.Key("factor")
const ParamOwner = wasmlib.Key("owner")

const VarFactor = wasmlib.Key("factor")
const VarMemberList = wasmlib.Key("memberList")
const VarMembers = wasmlib.Key("members")
const VarOwner = wasmlib.Key("owner")
const VarTotalFactor = wasmlib.Key("totalFactor")

const FuncDivide = "divide"
const FuncInit = "init"
const FuncMember = "member"
const FuncSetOwner = "setOwner"
const ViewGetFactor = "getFactor"

const HFuncDivide = wasmlib.ScHname(0xc7878107)
const HFuncInit = wasmlib.ScHname(0x1f44d644)
const HFuncMember = wasmlib.ScHname(0xc07da2cb)
const HFuncSetOwner = wasmlib.ScHname(0x2a15fe7b)
const HViewGetFactor = wasmlib.ScHname(0x0ee668fe)
