// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package fairroulette

import "github.com/iotaledger/wasplib/client"

type BetInfo struct {
	Amount int64
	Better *client.ScAgentId
	Color  int64
}

func EncodeBetInfo(o *BetInfo) []byte {
	return client.NewBytesEncoder().
		Int(o.Amount).
		Agent(o.Better).
		Int(o.Color).
		Data()
}

func DecodeBetInfo(bytes []byte) *BetInfo {
	decode := client.NewBytesDecoder(bytes)
	data := &BetInfo{}
	data.Amount = decode.Int()
	data.Better = decode.Agent()
	data.Color = decode.Int()
	return data
}
