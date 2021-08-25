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
	ParamChainID     = wasmlib.Key("ci")
	ParamDeployer    = wasmlib.Key("dp")
	ParamDescription = wasmlib.Key("ds")
	ParamHname       = wasmlib.Key("hn")
	ParamName        = wasmlib.Key("nm")
	ParamProgramHash = wasmlib.Key("ph")
)

const (
	ResultContractFound    = wasmlib.Key("cf")
	ResultContractRecData  = wasmlib.Key("dt")
	ResultContractRegistry = wasmlib.Key("r")
)

const (
	FuncDeployContract         = "deployContract"
	FuncGrantDeployPermission  = "grantDeployPermission"
	FuncInit                   = "init"
	FuncRevokeDeployPermission = "revokeDeployPermission"
	ViewFindContract           = "findContract"
	ViewGetContractRecords     = "getContractRecords"
)

const (
	HFuncDeployContract         = wasmlib.ScHname(0x28232c27)
	HFuncGrantDeployPermission  = wasmlib.ScHname(0xf440263a)
	HFuncInit                   = wasmlib.ScHname(0x1f44d644)
	HFuncRevokeDeployPermission = wasmlib.ScHname(0x850744f1)
	HViewFindContract           = wasmlib.ScHname(0xc145ca00)
	HViewGetContractRecords     = wasmlib.ScHname(0x078b3ef3)
)
