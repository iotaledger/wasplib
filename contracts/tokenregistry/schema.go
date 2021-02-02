// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package tokenregistry

import "github.com/iotaledger/wasplib/client"

const ScName = "tokenregistry"
const ScHname = client.Hname(0xe1ba0c78)

const ParamDescription = client.Key("description")
const ParamUserDefined = client.Key("user_defined")

const VarColorList = client.Key("color_list")
const VarRegistry = client.Key("registry")

const FuncMintSupply = "mint_supply"
const FuncTransferOwnership = "transfer_ownership"
const FuncUpdateMetadata = "update_metadata"

const HFuncMintSupply = client.Hname(0x5b0b99b9)
const HFuncTransferOwnership = client.Hname(0xea337e10)
const HFuncUpdateMetadata = client.Hname(0xaee46d94)

func OnLoad() {
	exports := client.NewScExports()
	exports.AddCall(FuncMintSupply, funcMintSupply)
	exports.AddCall(FuncTransferOwnership, funcTransferOwnership)
	exports.AddCall(FuncUpdateMetadata, funcUpdateMetadata)
}
