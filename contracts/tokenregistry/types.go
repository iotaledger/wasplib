// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package tokenregistry

import "github.com/iotaledger/wasplib/client"

type TokenInfo struct {
	Created     int64
	Description string
	MintedBy    *client.ScAgent
	Owner       *client.ScAgent
	Supply      int64
	Updated     int64
	UserDefined string
}

func EncodeTokenInfo(o *TokenInfo) []byte {
	return client.NewBytesEncoder().
		Int(o.Created).
		String(o.Description).
		Agent(o.MintedBy).
		Agent(o.Owner).
		Int(o.Supply).
		Int(o.Updated).
		String(o.UserDefined).
		Data()
}

func DecodeTokenInfo(bytes []byte) *TokenInfo {
	decode := client.NewBytesDecoder(bytes)
	data := &TokenInfo{}
	data.Created = decode.Int()
	data.Description = decode.String()
	data.MintedBy = decode.Agent()
	data.Owner = decode.Agent()
	data.Supply = decode.Int()
	data.Updated = decode.Int()
	data.UserDefined = decode.String()
	return data
}
