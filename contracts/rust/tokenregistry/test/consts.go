// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package test

import "github.com/iotaledger/wasp/packages/coretypes"

const (
	ScName  = "tokenregistry"
	HScName = coretypes.Hname(0xe1ba0c78)
)

const (
	ParamColor       = "color"
	ParamDescription = "description"
	ParamUserDefined = "userDefined"
)

const (
	StateColorList = "colorList"
	StateRegistry  = "registry"
)

const (
	FuncMintSupply        = "mintSupply"
	FuncTransferOwnership = "transferOwnership"
	FuncUpdateMetadata    = "updateMetadata"
	ViewGetInfo           = "getInfo"
)

const (
	HFuncMintSupply        = coretypes.Hname(0x564349a7)
	HFuncTransferOwnership = coretypes.Hname(0xbb9eb5af)
	HFuncUpdateMetadata    = coretypes.Hname(0xa26b23b6)
	HViewGetInfo           = coretypes.Hname(0xcfedba5f)
)
