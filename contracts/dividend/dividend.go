// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package dividend

import "github.com/iotaledger/wasplib/client"

const ParamAddress = client.Key("address")
const ParamFactor = client.Key("factor")

const VarMembers = client.Key("members")
const VarTotalFactor = client.Key("total_factor")

func OnLoad() {
	exports := client.NewScExports()
	exports.AddCall("member", member)
	exports.AddCall("divide", divide)
}

func member(sc *client.ScCallContext) {
	if !sc.From(sc.ContractCreator()) {
		sc.Panic("Cancel spoofed request")
	}
	params := sc.Params()
	address := params.GetAddress(ParamAddress)
	if !address.Exists() {
		sc.Panic("Missing address")
	}
	factor := params.GetInt(ParamFactor)
	if !factor.Exists() {
		sc.Panic("Missing factor")
	}
	member := &Member{
		Address: address.Value(),
		Factor:  factor.Value(),
	}
	state := sc.State()
	totalFactor := state.GetInt(VarTotalFactor)
	total := totalFactor.Value()
	members := state.GetBytesArray(VarMembers)
	size := members.Length()
	for i := int32(0); i < size; i++ {
		m := DecodeMember(members.GetBytes(i).Value())
		if m.Address.Equals(member.Address) {
			total -= m.Factor
			total += member.Factor
			totalFactor.SetValue(total)
			members.GetBytes(i).SetValue(EncodeMember(member))
			sc.Log("Updated: " + member.Address.String())
			return
		}
	}
	total += member.Factor
	totalFactor.SetValue(total)
	members.GetBytes(size).SetValue(EncodeMember(member))
	sc.Log("Appended: " + member.Address.String())
}

func divide(sc *client.ScCallContext) {
	amount := sc.Balances().Balance(client.IOTA)
	if amount == 0 {
		sc.Panic("Nothing to divide")
	}
	state := sc.State()
	totalFactor := state.GetInt(VarTotalFactor)
	total := totalFactor.Value()
	members := state.GetBytesArray(VarMembers)
	parts := int64(0)
	size := members.Length()
	for i := int32(0); i < size; i++ {
		m := DecodeMember(members.GetBytes(i).Value())
		part := amount * m.Factor / total
		if part != 0 {
			parts += part
			sc.TransferToAddress(m.Address, client.IOTA, part)
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
