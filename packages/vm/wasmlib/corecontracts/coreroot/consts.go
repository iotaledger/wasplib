// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package coreroot

import "github.com/iotaledger/wasplib/packages/vm/wasmlib"

const (
	ScName        = "root"
	ScDescription = "Core root contract"
	HScName       = wasmlib.ScHname(0xcebf5908)
)

const (
	ParamChainOwner   = wasmlib.Key("$$owner$$")
	ParamDeployer     = wasmlib.Key("$$deployer$$")
	ParamDescription  = wasmlib.Key("$$description$$")
	ParamHname        = wasmlib.Key("$$hname$$")
	ParamName         = wasmlib.Key("$$name$$")
	ParamOwnerFee     = wasmlib.Key("$$ownerfee$$")
	ParamProgramHash  = wasmlib.Key("$$proghash$$")
	ParamValidatorFee = wasmlib.Key("$$validatorfee$$")
)

const (
	ResultChainID             = wasmlib.Key("c")
	ResultChainOwnerID        = wasmlib.Key("o")
	ResultContractRegistry    = wasmlib.Key("r")
	ResultData                = wasmlib.Key("dt")
	ResultDefaultOwnerFee     = wasmlib.Key("do")
	ResultDefaultValidatorFee = wasmlib.Key("dv")
	ResultDescription         = wasmlib.Key("d")
	ResultFeeColor            = wasmlib.Key("f")
	ResultOwnerFee            = wasmlib.Key("of")
	ResultValidatorFee        = wasmlib.Key("vf")
)

const (
	FuncClaimChainOwnership    = "claimChainOwnership"
	FuncDelegateChainOwnership = "delegateChainOwnership"
	FuncDeployContract         = "deployContract"
	FuncGrantDeployPermission  = "grantDeployPermission"
	FuncRevokeDeployPermission = "revokeDeployPermission"
	FuncSetContractFee         = "setContractFee"
	FuncSetDefaultFee          = "setDefaultFee"
	ViewFindContract           = "findContract"
	ViewGetChainInfo           = "getChainInfo"
	ViewGetFeeInfo             = "getFeeInfo"
)

const (
	HFuncClaimChainOwnership    = wasmlib.ScHname(0x03ff0fc0)
	HFuncDelegateChainOwnership = wasmlib.ScHname(0x93ecb6ad)
	HFuncDeployContract         = wasmlib.ScHname(0x28232c27)
	HFuncGrantDeployPermission  = wasmlib.ScHname(0xf440263a)
	HFuncRevokeDeployPermission = wasmlib.ScHname(0x850744f1)
	HFuncSetContractFee         = wasmlib.ScHname(0x8421a42b)
	HFuncSetDefaultFee          = wasmlib.ScHname(0x3310ecd0)
	HViewFindContract           = wasmlib.ScHname(0xc145ca00)
	HViewGetChainInfo           = wasmlib.ScHname(0x434477e2)
	HViewGetFeeInfo             = wasmlib.ScHname(0x9fe54b48)
)
