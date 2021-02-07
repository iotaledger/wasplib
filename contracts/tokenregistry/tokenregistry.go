// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package tokenregistry

import "github.com/iotaledger/wasp/packages/vm/wasmlib"

func funcMintSupply(ctx *wasmlib.ScCallContext) {
	minted := ctx.Incoming().Minted()
	if minted.Equals(wasmlib.MINT) {
		ctx.Panic("TokenRegistry: No newly minted tokens found")
	}
	state := ctx.State()
	registry := state.GetMap(VarRegistry).GetBytes(minted)
	if registry.Exists() {
		ctx.Panic("TokenRegistry: Color already exists")
	}
	params := ctx.Params()
	token := &TokenInfo{
		Supply:      ctx.Incoming().Balance(minted),
		MintedBy:    ctx.Caller(),
		Owner:       ctx.Caller(),
		Created:     ctx.Timestamp(),
		Updated:     ctx.Timestamp(),
		Description: params.GetString(ParamDescription).Value(),
		UserDefined: params.GetString(ParamUserDefined).Value(),
	}
	if token.Supply <= 0 {
		ctx.Panic("TokenRegistry: Insufficient supply")
	}
	if len(token.Description) == 0 {
		token.Description += "no dscr"
	}
	registry.SetValue(EncodeTokenInfo(token))
	colors := state.GetColorArray(VarColorList)
	colors.GetColor(colors.Length()).SetValue(minted)
}

func funcTransferOwnership(_sc *wasmlib.ScCallContext) {
	//TODO
}

func funcUpdateMetadata(_sc *wasmlib.ScCallContext) {
	//TODO
}

func viewGetInfo(_sc *wasmlib.ScViewContext) {
	//TODO
}
