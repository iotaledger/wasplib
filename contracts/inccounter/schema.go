// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package inccounter

import "github.com/iotaledger/wasplib/client"

const ScName = "inccounter"
const ScHname = client.ScHname(0xaf2438e9)

const ParamCounter = client.Key("counter")
const ParamNumRepeats = client.Key("numRepeats")

const VarCounter = client.Key("counter")
const VarInt1 = client.Key("int1")
const VarIntArray1 = client.Key("intArray1")
const VarNumRepeats = client.Key("numRepeats")
const VarString1 = client.Key("string1")
const VarStringArray1 = client.Key("stringArray1")

const FuncCallIncrement = "callIncrement"
const FuncCallIncrementRecurse5x = "callIncrementRecurse5x"
const FuncIncrement = "increment"
const FuncInit = "init"
const FuncLocalStateInternalCall = "localStateInternalCall"
const FuncLocalStatePost = "localStatePost"
const FuncLocalStateSandboxCall = "localStateSandboxCall"
const FuncPostIncrement = "postIncrement"
const FuncRepeatMany = "repeatMany"
const FuncResultsTest = "resultsTest"
const FuncStateTest = "stateTest"
const FuncWhenMustIncrement = "whenMustIncrement"
const ViewGetCounter = "getCounter"
const ViewResultsCheck = "resultsCheck"
const ViewStateCheck = "stateCheck"

const HFuncCallIncrement = client.ScHname(0xeb5dcacd)
const HFuncCallIncrementRecurse5x = client.ScHname(0x8749fbff)
const HFuncIncrement = client.ScHname(0xd351bd12)
const HFuncInit = client.ScHname(0x1f44d644)
const HFuncLocalStateInternalCall = client.ScHname(0xecfc5d33)
const HFuncLocalStatePost = client.ScHname(0x3fd54d13)
const HFuncLocalStateSandboxCall = client.ScHname(0x7bd22c53)
const HFuncPostIncrement = client.ScHname(0x81c772f5)
const HFuncRepeatMany = client.ScHname(0x4ff450d3)
const HFuncResultsTest = client.ScHname(0xd0544634)
const HFuncStateTest = client.ScHname(0x41830d59)
const HFuncWhenMustIncrement = client.ScHname(0xb4c3e7a6)
const HViewGetCounter = client.ScHname(0xb423e607)
const HViewResultsCheck = client.ScHname(0xa39ac571)
const HViewStateCheck = client.ScHname(0xaafeb10a)

func OnLoad() {
    exports := client.NewScExports()
    exports.AddCall(FuncCallIncrement, funcCallIncrement)
    exports.AddCall(FuncCallIncrementRecurse5x, funcCallIncrementRecurse5x)
    exports.AddCall(FuncIncrement, funcIncrement)
    exports.AddCall(FuncInit, funcInit)
    exports.AddCall(FuncLocalStateInternalCall, funcLocalStateInternalCall)
    exports.AddCall(FuncLocalStatePost, funcLocalStatePost)
    exports.AddCall(FuncLocalStateSandboxCall, funcLocalStateSandboxCall)
    exports.AddCall(FuncPostIncrement, funcPostIncrement)
    exports.AddCall(FuncRepeatMany, funcRepeatMany)
    exports.AddCall(FuncResultsTest, funcResultsTest)
    exports.AddCall(FuncStateTest, funcStateTest)
    exports.AddCall(FuncWhenMustIncrement, funcWhenMustIncrement)
    exports.AddView(ViewGetCounter, viewGetCounter)
    exports.AddView(ViewResultsCheck, viewResultsCheck)
    exports.AddView(ViewStateCheck, viewStateCheck)
}
