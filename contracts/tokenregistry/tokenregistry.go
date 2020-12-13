// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package tokenregistry

import (
	"github.com/iotaledger/wasplib/client"
)

const (
	keyColorList   = client.Key("color_list")
	keyDescription = client.Key("description")
	keyRegistry    = client.Key("registry")
	keyUserDefined = client.Key("user_defined")
)

type TokenInfo struct {
	supply      int64
	mintedBy    *client.ScAgent
	owner       *client.ScAgent
	created     int64
	updated     int64
	description string
	userDefined string
}

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
	bytes := encodeTokenInfo(token)
	registry.SetValue(bytes)
	colors := state.GetColorArray(keyColorList)
	colors.GetColor(colors.Length()).SetValue(minted)
}

func updateMetadata(sc *client.ScCallContext) {
	//TODO
}

func transferOwnership(sc *client.ScCallContext) {
	//TODO
}

func decodeTokenInfo(bytes []byte) *TokenInfo {
	decoder := client.NewBytesDecoder(bytes)
	data := &TokenInfo{}
	data.supply = decoder.Int()
	data.mintedBy = decoder.Agent()
	data.owner = decoder.Agent()
	data.created = decoder.Int()
	data.updated = decoder.Int()
	data.description = decoder.String()
	data.userDefined = decoder.String()
	return data
}

func encodeTokenInfo(token *TokenInfo) []byte {
	return client.NewBytesEncoder().
		Int(token.supply).
		Agent(token.mintedBy).
		Agent(token.owner).
		Int(token.created).
		Int(token.updated).
		String(token.description).
		String(token.userDefined).
		Data()
}
