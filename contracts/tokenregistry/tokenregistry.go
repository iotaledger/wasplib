// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package tokenregistry

import "github.com/iotaledger/wasplib/client"

const keyColorList = client.Key("color_list")
const keyDescription = client.Key("description")
const keyRegistry = client.Key("registry")
const keyUserDefined = client.Key("user_defined")

func OnLoad() {
	exports := client.NewScExports()
	exports.AddCall("mint_supply", mintSupply)
	exports.AddCall("update_metadata", updateMetadata)
	exports.AddCall("transfer_ownership", transferOwnership)
}

func mintSupply(sc *client.ScCallContext) {
	minted := sc.Incoming().Minted()
	if minted.Equals(client.MINT) {
		sc.Log("TokenRegistry: No newly minted tokens found")
		return
	}
	state := sc.State()
	registry := state.GetMap(keyRegistry).GetBytes(minted)
	if registry.Exists() {
		sc.Log("TokenRegistry: Color already exists")
		return
	}
	params := sc.Params()
	token := &TokenInfo{
		supply:      sc.Incoming().Balance(minted),
		mintedBy:    sc.Caller(),
		owner:       sc.Caller(),
		created:     sc.Timestamp(),
		updated:     sc.Timestamp(),
		description: params.GetString(keyDescription).Value(),
		userDefined: params.GetString(keyUserDefined).Value(),
	}
	if token.supply <= 0 {
		sc.Log("TokenRegistry: Insufficient supply")
		return
	}
	if len(token.description) == 0 {
		token.description += "no dscr"
	}
	registry.SetValue(encodeTokenInfo(token))
	colors := state.GetColorArray(keyColorList)
	colors.GetColor(colors.Length()).SetValue(minted)
}

func updateMetadata(_sc *client.ScCallContext) {
	//TODO
}

func transferOwnership(_sc *client.ScCallContext) {
	//TODO
}
