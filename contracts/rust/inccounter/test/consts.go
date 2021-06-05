// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package test

import "github.com/iotaledger/wasplib/packages/vm/wasmlib"

const ScName = "inccounter"
const HScName = wasmlib.ScHname(0xaf2438e9)

const ParamCounter = "counter"
const ParamNumRepeats = "numRepeats"

const ResultCounter = "counter"

const VarCounter = "counter"
const VarNumRepeats = "numRepeats"

const FuncCallIncrement = "callIncrement"
const FuncCallIncrementRecurse5x = "callIncrementRecurse5x"
const FuncIncrement = "increment"
const FuncInit = "init"
const FuncLocalStateInternalCall = "localStateInternalCall"
const FuncLocalStatePost = "localStatePost"
const FuncLocalStateSandboxCall = "localStateSandboxCall"
const FuncLoop = "loop"
const FuncPostIncrement = "postIncrement"
const FuncRepeatMany = "repeatMany"
const FuncTestLeb128 = "testLeb128"
const FuncWhenMustIncrement = "whenMustIncrement"
const ViewGetCounter = "getCounter"

const HFuncCallIncrement = wasmlib.ScHname(0xeb5dcacd)
const HFuncCallIncrementRecurse5x = wasmlib.ScHname(0x8749fbff)
const HFuncIncrement = wasmlib.ScHname(0xd351bd12)
const HFuncInit = wasmlib.ScHname(0x1f44d644)
const HFuncLocalStateInternalCall = wasmlib.ScHname(0xecfc5d33)
const HFuncLocalStatePost = wasmlib.ScHname(0x3fd54d13)
const HFuncLocalStateSandboxCall = wasmlib.ScHname(0x7bd22c53)
const HFuncLoop = wasmlib.ScHname(0xa9a20fa9)
const HFuncPostIncrement = wasmlib.ScHname(0x81c772f5)
const HFuncRepeatMany = wasmlib.ScHname(0x4ff450d3)
const HFuncTestLeb128 = wasmlib.ScHname(0xd8364cb9)
const HFuncWhenMustIncrement = wasmlib.ScHname(0xb4c3e7a6)
const HViewGetCounter = wasmlib.ScHname(0xb423e607)
