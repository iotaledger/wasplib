// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package dividend

import "github.com/iotaledger/wasplib/client"

type Member struct {
	address *client.ScAddress
	factor  int64
}

func encodeMember(o *Member) []byte {
	return client.NewBytesEncoder().
		Address(o.address).
		Int(o.factor).
		Data()
}

func decodeMember(bytes []byte) *Member {
	decode := client.NewBytesDecoder(bytes)
	data := &Member{}
	data.address = decode.Address()
	data.factor = decode.Int()
	return data
}
