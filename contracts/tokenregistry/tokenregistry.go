// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package tokenregistry

import "github.com/iotaledger/wasp/packages/vm/wasmlib"

func funcMintSupply(ctx *wasmlib.ScFuncContext, params *FuncMintSupplyParams) {
	minted := ctx.Incoming().Minted()
	if minted.Equals(wasmlib.MINT) {
		ctx.Panic("TokenRegistry: No newly minted tokens found")
	}
	state := ctx.State()
	registry := state.GetMap(VarRegistry).GetBytes(minted)
	if registry.Exists() {
		ctx.Panic("TokenRegistry: Color already exists")
	}
	token := &Token{
		Supply:      ctx.Incoming().Balance(minted),
		MintedBy:    ctx.Caller(),
		Owner:       ctx.Caller(),
		Created:     ctx.Timestamp(),
		Updated:     ctx.Timestamp(),
		Description: params.Description.Value(),
		UserDefined: params.UserDefined.Value(),
	}
	if token.Supply <= 0 {
		ctx.Panic("TokenRegistry: Insufficient supply")
	}
	if len(token.Description) == 0 {
		token.Description += "no dscr"
	}
	registry.SetValue(token.Bytes())
	colors := state.GetColorArray(VarColorList)
	colors.GetColor(colors.Length()).SetValue(minted)
}

func funcTransferOwnership(_sc *wasmlib.ScFuncContext, params *FuncTransferOwnershipParams) {
	//TODO
}

func funcUpdateMetadata(_sc *wasmlib.ScFuncContext, params *FuncUpdateMetadataParams) {
	//TODO
}

func viewGetInfo(_sc *wasmlib.ScViewContext, params *ViewGetInfoParams) {
	//TODO
}
