// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package inccounter

import "github.com/iotaledger/wasplib/client"

const ScName = "inccounter"
const ScHname = client.Hname(0xaf2438e9)

const ParamCounter = client.Key("counter")
const ParamNumRepeats = client.Key("num_repeats")

const VarCounter = client.Key("counter")
const VarInt1 = client.Key("i1")
const VarIntArray1 = client.Key("ia1")
const VarNumRepeats = client.Key("num_repeats")
const VarString1 = client.Key("s1")
const VarStringArray1 = client.Key("sa1")

const FuncCallIncrement = "call_increment"
const FuncCallIncrementRecurse5x = "call_increment_recurse5x"
const FuncIncrement = "increment"
const FuncInit = "init"
const FuncLocalStateInternalCall = "local_state_internal_call"
const FuncLocalStatePost = "local_state_post"
const FuncLocalStateSandboxCall = "local_state_sandbox_call"
const FuncPostIncrement = "post_increment"
const FuncRepeatMany = "repeat_many"
const FuncResultsTest = "results_test"
const FuncStateTest = "state_test"
const FuncWhenMustIncrement = "when_must_increment"
const ViewGetCounter = "get_counter"
const ViewResultsCheck = "results_check"
const ViewStateCheck = "state_check"

const HFuncCallIncrement = client.Hname(0x96b915f2)
const HFuncCallIncrementRecurse5x = client.Hname(0x30319639)
const HFuncIncrement = client.Hname(0xd351bd12)
const HFuncInit = client.Hname(0x1f44d644)
const HFuncLocalStateInternalCall = client.Hname(0xc4e9cbef)
const HFuncLocalStatePost = client.Hname(0x90051958)
const HFuncLocalStateSandboxCall = client.Hname(0x07431bc8)
const HFuncPostIncrement = client.Hname(0xb775b58a)
const HFuncRepeatMany = client.Hname(0x020e669e)
const HFuncResultsTest = client.Hname(0xf73a0ee0)
const HFuncStateTest = client.Hname(0x5691431b)
const HFuncWhenMustIncrement = client.Hname(0x28a49492)
const HViewGetCounter = client.Hname(0xb8e70081)
const HViewResultsCheck = client.Hname(0xfaf7081b)
const HViewStateCheck = client.Hname(0xc760249d)

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
