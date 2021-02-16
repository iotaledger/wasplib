// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package test

import (
	"github.com/iotaledger/wasp/packages/coretypes"
)

const ScName = "testcore"
const ScDescription = "Core test for ISCP wasmlib Rust/Wasm library"
const ScHname = coretypes.Hname(0x370d33ad)

const ParamAddress = "address"
const ParamAgentId = "agentID"
const ParamCaller = "caller"
const ParamChainId = "chainid"
const ParamChainOwnerId = "chainOwnerID"
const ParamContractCreator = "contractCreator"
const ParamContractId = "contractID"
const ParamCounter = "counter"
const ParamHash = "Hash"
const ParamHname = "Hname"
const ParamHnameContract = "hnameContract"
const ParamHnameEP = "hnameEP"
const ParamHnameZero = "Hname-0"
const ParamInt64 = "int64"
const ParamInt64Zero = "int64-0"
const ParamIntValue = "intParamValue"
const ParamName = "intParamName"
const ParamString = "string"
const ParamStringZero = "string-0"

const VarCounter = "counter"
const VarHnameEP = "hnameEP"

const FuncCallOnChain = "callOnChain"
const FuncCheckContextFromFullEP = "checkContextFromFullEP"
const FuncDoNothing = "doNothing"
const FuncInit = "init"
const FuncPassTypesFull = "passTypesFull"
const FuncRunRecursion = "runRecursion"
const FuncSendToAddress = "sendToAddress"
const FuncSetInt = "setInt"
const FuncTestCallPanicFullEP = "testCallPanicFullEP"
const FuncTestCallPanicViewEPFromFull = "testCallPanicViewEPFromFull"
const FuncTestChainOwnerIDFull = "testChainOwnerIDFull"
const FuncTestContractIDFull = "testContractIDFull"
const FuncTestEventLogDeploy = "testEventLogDeploy"
const FuncTestEventLogEventData = "testEventLogEventData"
const FuncTestEventLogGenericData = "testEventLogGenericData"
const FuncTestPanicFullEP = "testPanicFullEP"
const FuncWithdrawToChain = "withdrawToChain"
const ViewCheckContextFromViewEP = "checkContextFromViewEP"
const ViewFibonacci = "fibonacci"
const ViewGetCounter = "getCounter"
const ViewGetInt = "getInt"
const ViewJustView = "justView"
const ViewPassTypesView = "passTypesView"
const ViewTestCallPanicViewEPFromView = "testCallPanicViewEPFromView"
const ViewTestChainOwnerIDView = "testChainOwnerIDView"
const ViewTestContractIDView = "testContractIDView"
const ViewTestPanicViewEP = "testPanicViewEP"
const ViewTestSandboxCall = "testSandboxCall"

const HFuncCallOnChain = coretypes.Hname(0x95a3d123)
const HFuncCheckContextFromFullEP = coretypes.Hname(0xa56c24ba)
const HFuncDoNothing = coretypes.Hname(0xdda4a6de)
const HFuncInit = coretypes.Hname(0x1f44d644)
const HFuncPassTypesFull = coretypes.Hname(0x733ea0ea)
const HFuncRunRecursion = coretypes.Hname(0x833425fd)
const HFuncSendToAddress = coretypes.Hname(0x63ce4634)
const HFuncSetInt = coretypes.Hname(0x62056f74)
const HFuncTestCallPanicFullEP = coretypes.Hname(0x4c878834)
const HFuncTestCallPanicViewEPFromFull = coretypes.Hname(0xfd7e8c1d)
const HFuncTestChainOwnerIDFull = coretypes.Hname(0x2aff1167)
const HFuncTestContractIDFull = coretypes.Hname(0x95934282)
const HFuncTestEventLogDeploy = coretypes.Hname(0x96ff760a)
const HFuncTestEventLogEventData = coretypes.Hname(0x0efcf939)
const HFuncTestEventLogGenericData = coretypes.Hname(0x6a16629d)
const HFuncTestPanicFullEP = coretypes.Hname(0x24fdef07)
const HFuncWithdrawToChain = coretypes.Hname(0x437bc026)
const HViewCheckContextFromViewEP = coretypes.Hname(0x88ff0167)
const HViewFibonacci = coretypes.Hname(0x7940873c)
const HViewGetCounter = coretypes.Hname(0xb423e607)
const HViewGetInt = coretypes.Hname(0x1887e5ef)
const HViewJustView = coretypes.Hname(0x33b8972e)
const HViewPassTypesView = coretypes.Hname(0x1a5b87ea)
const HViewTestCallPanicViewEPFromView = coretypes.Hname(0x91b10c99)
const HViewTestChainOwnerIDView = coretypes.Hname(0x26586c33)
const HViewTestContractIDView = coretypes.Hname(0x28a02913)
const HViewTestPanicViewEP = coretypes.Hname(0x22bc4d72)
const HViewTestSandboxCall = coretypes.Hname(0x42d72b63)
