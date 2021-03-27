// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package tokenregistry

import (
	"github.com/iotaledger/wasplib/packages/vm/wasmlib"
)

func funcMintSupply(ctx wasmlib.ScFuncContext, params *FuncMintSupplyParams) {
    minted := ctx.Minted()
    mintedColors := minted.Colors()
    ctx.Require(mintedColors.Length() == 1, "need single minted color")
    mintedColor := mintedColors.GetColor(0).Value()
    state := ctx.State()
    registry := state.GetMap(VarRegistry).GetBytes(mintedColor)
    if registry.Exists() {
        // should never happen, because transaction id is unique
        ctx.Panic("TokenRegistry: registry for color already exists")
    }
    token := &Token {
        Supply: minted.Balance(mintedColor),
        MintedBy: ctx.Caller(),
        Owner: ctx.Caller(),
        Created: ctx.Timestamp(),
        Updated: ctx.Timestamp(),
        Description: params.Description.Value(),
        UserDefined: params.UserDefined.Value(),
    }
    if len(token.Description) == 0 {
        token.Description += "no dscr"
    }
    registry.SetValue(token.Bytes())
    colors := state.GetColorArray(VarColorList)
    colors.GetColor(colors.Length()).SetValue(mintedColor)
}

func funcTransferOwnership(ctx wasmlib.ScFuncContext, params *FuncTransferOwnershipParams) {
    //TODO
}

func funcUpdateMetadata(ctx wasmlib.ScFuncContext, params *FuncUpdateMetadataParams) {
    //TODO
}

func viewGetInfo(ctx wasmlib.ScViewContext, params *ViewGetInfoParams) {
    //TODO
}
