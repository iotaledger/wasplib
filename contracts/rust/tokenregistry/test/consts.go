// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package test

import (
	"github.com/iotaledger/wasp/packages/coretypes"
)

const ScName = "tokenregistry"
const ScHname = coretypes.Hname(0xe1ba0c78)

const ParamColor = "color"
const ParamDescription = "description"
const ParamUserDefined = "userDefined"

const VarColorList = "colorList"
const VarRegistry = "registry"

const FuncMintSupply = "mintSupply"
const FuncTransferOwnership = "transferOwnership"
const FuncUpdateMetadata = "updateMetadata"
const ViewGetInfo = "getInfo"

const HFuncMintSupply = coretypes.Hname(0x564349a7)
const HFuncTransferOwnership = coretypes.Hname(0xbb9eb5af)
const HFuncUpdateMetadata = coretypes.Hname(0xa26b23b6)
const HViewGetInfo = coretypes.Hname(0xcfedba5f)
