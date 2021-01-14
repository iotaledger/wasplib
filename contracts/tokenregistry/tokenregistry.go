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

func mintSupply(sc *client.ScCallContext) {
	minted := sc.Incoming().Minted()
	if minted.Equals(client.MINT) {
		sc.Panic("TokenRegistry: No newly minted tokens found")
	}
	state := sc.State()
	registry := state.GetMap(KeyRegistry).GetBytes(minted)
	if registry.Exists() {
		sc.Panic("TokenRegistry: Color already exists")
	}
	params := sc.Params()
	token := &TokenInfo{
		Supply:      sc.Incoming().Balance(minted),
		MintedBy:    sc.Caller(),
		Owner:       sc.Caller(),
		Created:     sc.Timestamp(),
		Updated:     sc.Timestamp(),
		Description: params.GetString(KeyDescription).Value(),
		UserDefined: params.GetString(KeyUserDefined).Value(),
	}
	if token.Supply <= 0 {
		sc.Panic("TokenRegistry: Insufficient supply")
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
