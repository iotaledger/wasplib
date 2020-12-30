// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package dividend

import "github.com/iotaledger/wasplib/client"

const keyAddress = client.Key("address")
const keyFactor = client.Key("factor")
const keyMembers = client.Key("members")
const keyTotalFactor = client.Key("total_factor")

func OnLoad() {
	exports := client.NewScExports()
	exports.AddCall("member", member)
	exports.AddCall("dividend", dividend)
}

func member(sc *client.ScCallContext) {
	if !sc.From(sc.Contract().Creator()) {
		sc.Panic("Cancel spoofed request")
	}
	params := sc.Params()
	address := params.GetAddress(keyAddress)
	if !address.Exists() {
		sc.Panic("Missing address")
	}
	factor := params.GetInt(keyFactor)
	if !factor.Exists() {
		sc.Panic("Missing factor")
	}
	member := &Member{
		address: address.Value(),
		factor:  factor.Value(),
	}
	state := sc.State()
	totalFactor := state.GetInt(keyTotalFactor)
	total := totalFactor.Value()
	members := state.GetBytesArray(keyMembers)
	size := members.Length()
	for i := int32(0); i < size; i++ {
		m := decodeMember(members.GetBytes(i).Value())
		if m.address.Equals(member.address) {
			total -= m.factor
			total += member.factor
			totalFactor.SetValue(total)
			members.GetBytes(i).SetValue(encodeMember(member))
			sc.Log("Updated: " + member.address.String())
			return
		}
	}
	total += member.factor
	totalFactor.SetValue(total)
	members.GetBytes(size).SetValue(encodeMember(member))
	sc.Log("Appended: " + member.address.String())
}

func dividend(sc *client.ScCallContext) {
	amount := sc.Balances().Balance(client.IOTA)
	if amount == 0 {
		sc.Panic("Nothing to divide")
	}
	state := sc.State()
	totalFactor := state.GetInt(keyTotalFactor)
	total := totalFactor.Value()
	members := state.GetBytesArray(keyMembers)
	parts := int64(0)
	size := members.Length()
	for i := int32(0); i < size; i++ {
		m := decodeMember(members.GetBytes(i).Value())
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
