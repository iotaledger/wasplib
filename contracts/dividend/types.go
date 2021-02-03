// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package dividend

import "github.com/iotaledger/wasplib/client"

type Member struct {
	Address *client.ScAddress // address of dividend recipient
	Factor  int64             // relative division factor
}

func EncodeMember(o *Member) []byte {
	return client.NewBytesEncoder().
		Address(o.Address).
		Int(o.Factor).
		Data()
}

func DecodeMember(bytes []byte) *Member {
	decode := client.NewBytesDecoder(bytes)
	data := &Member{}
	data.Address = decode.Address()
	data.Factor = decode.Int()
	return data
}
