// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package tokenregistry

import "github.com/iotaledger/wasplib/client"

type TokenInfo struct {
	created     int64
	description string
	mintedBy    *client.ScAgent
	owner       *client.ScAgent
	supply      int64
	updated     int64
	userDefined string
}

func encodeTokenInfo(o *TokenInfo) []byte {
	return client.NewBytesEncoder().
		Int(o.created).
		String(o.description).
		Agent(o.mintedBy).
		Agent(o.owner).
		Int(o.supply).
		Int(o.updated).
		String(o.userDefined).
		Data()
}

func decodeTokenInfo(bytes []byte) *TokenInfo {
	decode := client.NewBytesDecoder(bytes)
	data := &TokenInfo{}
	data.created = decode.Int()
	data.description = decode.String()
	data.mintedBy = decode.Agent()
	data.owner = decode.Agent()
	data.supply = decode.Int()
	data.updated = decode.Int()
	data.userDefined = decode.String()
	return data
}
