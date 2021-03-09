// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package dividend

import "github.com/iotaledger/wasplib/packages/vm/wasmlib"

type Member struct {
	Address wasmlib.ScAddress // address of dividend recipient
	Factor  int64             // relative division factor
}

func NewMemberFromBytes(bytes []byte) *Member {
	decode := wasmlib.NewBytesDecoder(bytes)
	data := &Member{}
	data.Address = decode.Address()
	data.Factor = decode.Int64()
	return data
}

func (o *Member) Bytes() []byte {
	return wasmlib.NewBytesEncoder().
		Address(o.Address).
		Int64(o.Factor).
		Data()
}