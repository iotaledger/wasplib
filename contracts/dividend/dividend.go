// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package dividend

import (
	"github.com/iotaledger/wasplib/client"
)

type Member struct {
	address *client.ScAddress
	factor  int64
}

func OnLoad() {
	exports := client.NewScExports()
	exports.AddCall("member", member)
	exports.AddCall("divide", divide)
}

func member(sc *client.ScCallContext) {
	if !sc.From(sc.Contract().Owner()) {
		sc.Log("Cancel spoofed request")
		return
	}
	params := sc.Params()
	address := params.GetAddress("address")
	if !address.Exists() {
		sc.Log("Missing address")
		return
	}
	factor := params.GetInt("factor")
	if !factor.Exists() {
		sc.Log("Missing factor")
		return
	}
	member := &Member{
		address: address.Value(),
		factor:  factor.Value(),
	}
	state := sc.State()
	totalFactor := state.GetInt("totalFactor")
	total := totalFactor.Value()
	members := state.GetBytesArray("members")
	size := members.Length()
	for i := int32(0); i < size; i++ {
		bytes := members.GetBytes(i).Value()
		m := decodeMember(bytes)
		if m.address.Equals(member.address) {
			total -= m.factor
			total += member.factor
			totalFactor.SetValue(total)
			bytes := encodeMember(member)
			members.GetBytes(i).SetValue(bytes)
			sc.Log("Updated: " + member.address.String())
			return
		}
	}
	total += member.factor
	totalFactor.SetValue(total)
	bytes := encodeMember(member)
	members.GetBytes(size).SetValue(bytes)
	sc.Log("Appended: " + member.address.String())
}

func divide(sc *client.ScCallContext) {
	amount := sc.Balances().Balance(client.IOTA)
	if amount == 0 {
		sc.Log("Nothing to divide")
		return
	}
	state := sc.State()
	totalFactor := state.GetInt("totalFactor")
	total := totalFactor.Value()
	members := state.GetBytesArray("members")
	size := members.Length()
	parts := int64(0)
	for i := int32(0); i < size; i++ {
		bytes := members.GetBytes(i).Value()
		m := decodeMember(bytes)
		part := amount * m.factor / total
		if part != 0 {
			parts += part
			sc.Transfer(m.address.AsAgent(), client.IOTA, part)
		}
	}
	if parts != amount {
		// note we truncated the calculations down to the nearest integer
		// there could be some small remainder left in the contract, but
		// that will be picked up in the next round as part of the balance
		remainder := amount - parts
		sc.Log("Remainder in contract: " + sc.Utility().String(remainder))
	}
}

func encodeMember(member *Member) []byte {
	return client.NewBytesEncoder().
		Address(member.address).
		Int(member.factor).
		Data()
}

func decodeMember(bytes []byte) *Member {
	decoder := client.NewBytesDecoder(bytes)
	member := &Member{}
	member.address = decoder.Address()
	member.factor = decoder.Int()
	return member
}
