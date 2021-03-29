// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package testwasmlib

import "github.com/iotaledger/wasplib/packages/vm/wasmlib"

const ScName = "testwasmlib"
const ScDescription = "Exercise all aspects of WasmLib"
const HScName = wasmlib.ScHname(0x89703a45)

const ParamAddress = wasmlib.Key("address")
const ParamAgentId = wasmlib.Key("agentId")
const ParamBytes = wasmlib.Key("bytes")
const ParamChainId = wasmlib.Key("chainId")
const ParamColor = wasmlib.Key("color")
const ParamHash = wasmlib.Key("hash")
const ParamHname = wasmlib.Key("hname")
const ParamInt64 = wasmlib.Key("int64")
const ParamRequestId = wasmlib.Key("requestId")
const ParamString = wasmlib.Key("string")

const FuncParamTypes = "paramTypes"

const HFuncParamTypes = wasmlib.ScHname(0x6921c4cd)