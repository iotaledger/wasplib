// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package tokenregistry

import "github.com/iotaledger/wasplib/packages/vm/wasmlib"

const (
	ScName  = "tokenregistry"
	HScName = wasmlib.ScHname(0xe1ba0c78)
)

const (
	ParamColor       = wasmlib.Key("color")
	ParamDescription = wasmlib.Key("description")
	ParamUserDefined = wasmlib.Key("userDefined")
)

const (
	StateColorList = wasmlib.Key("colorList")
	StateRegistry  = wasmlib.Key("registry")
)

const (
	FuncMintSupply        = "mintSupply"
	FuncTransferOwnership = "transferOwnership"
	FuncUpdateMetadata    = "updateMetadata"
	ViewGetInfo           = "getInfo"
)

const (
	HFuncMintSupply        = wasmlib.ScHname(0x564349a7)
	HFuncTransferOwnership = wasmlib.ScHname(0xbb9eb5af)
	HFuncUpdateMetadata    = wasmlib.ScHname(0xa26b23b6)
	HViewGetInfo           = wasmlib.ScHname(0xcfedba5f)
)
