// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package tokenregistry

import "github.com/iotaledger/wasplib/client"

const ScName = "tokenregistry"
const ScHname = client.ScHname(0xe1ba0c78)

const ParamColor = client.Key("color")
const ParamDescription = client.Key("description")
const ParamUserDefined = client.Key("userDefined")

const VarColorList = client.Key("colorList")
const VarRegistry = client.Key("registry")

const FuncMintSupply = "mintSupply"
const FuncTransferOwnership = "transferOwnership"
const FuncUpdateMetadata = "updateMetadata"
const ViewGetInfo = "getInfo"

const HFuncMintSupply = client.ScHname(0x564349a7)
const HFuncTransferOwnership = client.ScHname(0xbb9eb5af)
const HFuncUpdateMetadata = client.ScHname(0xa26b23b6)
const HViewGetInfo = client.ScHname(0xcfedba5f)

func OnLoad() {
    exports := client.NewScExports()
    exports.AddCall(FuncMintSupply, funcMintSupply)
    exports.AddCall(FuncTransferOwnership, funcTransferOwnership)
    exports.AddCall(FuncUpdateMetadata, funcUpdateMetadata)
    exports.AddView(ViewGetInfo, viewGetInfo)
}
