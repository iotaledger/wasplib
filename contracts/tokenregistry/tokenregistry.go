// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package tokenregistry

import "github.com/iotaledger/wasplib/client"

const KeyColorList = client.Key("color_list")
const KeyDescription = client.Key("description")
const KeyRegistry = client.Key("registry")
const KeyUserDefined = client.Key("user_defined")

func OnLoad() {
	exports := client.NewScExports()
	exports.AddCall("mint_supply", mintSupply)
	exports.AddCall("update_metadata", updateMetadata)
	exports.AddCall("transfer_ownership", transferOwnership)
}

func mintSupply(ctx *client.ScCallContext) {
	minted := ctx.Incoming().Minted()
	if minted.Equals(client.MINT) {
		ctx.Panic("TokenRegistry: No newly minted tokens found")
	}
	state := ctx.State()
	registry := state.GetMap(KeyRegistry).GetBytes(minted)
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
		Description: params.GetString(KeyDescription).Value(),
		UserDefined: params.GetString(KeyUserDefined).Value(),
	}
	if token.Supply <= 0 {
		ctx.Panic("TokenRegistry: Insufficient supply")
	}
	if len(token.Description) == 0 {
		token.Description += "no dscr"
	}
	registry.SetValue(EncodeTokenInfo(token))
	colors := state.GetColorArray(KeyColorList)
	colors.GetColor(colors.Length()).SetValue(minted)
}

func updateMetadata(_sc *client.ScCallContext) {
	//TODO
}

func transferOwnership(_sc *client.ScCallContext) {
	//TODO
}
