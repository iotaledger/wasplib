// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package tokenregistry

import (
	"github.com/iotaledger/wasp/packages/vm/wasmlib"
)

func funcMintSupply(ctx wasmlib.ScFuncContext, f *MintSupplyContext) {
	minted := ctx.Minted()
	mintedColors := minted.Colors()
	ctx.Require(mintedColors.Length() == 1, "need single minted color")
	mintedColor := mintedColors.GetColor(0).Value()
	currentToken := f.State.Registry().GetToken(mintedColor)
	if currentToken.Exists() {
		// should never happen, because transaction id is unique
		ctx.Panic("TokenRegistry: registry for color already exists")
	}
	token := &Token{
		Supply:      minted.Balance(mintedColor),
		MintedBy:    ctx.Caller(),
		Owner:       ctx.Caller(),
		Created:     ctx.Timestamp(),
		Updated:     ctx.Timestamp(),
		Description: f.Params.Description().Value(),
		UserDefined: f.Params.UserDefined().Value(),
	}
	if token.Description == "" {
		token.Description += "no dscr"
	}
	currentToken.SetValue(token)
	colorList := f.State.ColorList()
	colorList.GetColor(colorList.Length()).SetValue(mintedColor)
}

func funcTransferOwnership(ctx wasmlib.ScFuncContext, f *TransferOwnershipContext) {
	// TODO
}

func funcUpdateMetadata(ctx wasmlib.ScFuncContext, f *UpdateMetadataContext) {
	// TODO
}

func viewGetInfo(ctx wasmlib.ScViewContext, f *GetInfoContext) {
	// TODO
}
