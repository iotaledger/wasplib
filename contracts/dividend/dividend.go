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

func member(ctx *client.ScCallContext) {
	if !ctx.From(ctx.ContractCreator()) {
		ctx.Panic("Cancel spoofed request")
	}
	params := ctx.Params()
	address := params.GetAddress(ParamAddress)
	if !address.Exists() {
		ctx.Panic("Missing address")
	}
	factor := params.GetInt(ParamFactor)
	if !factor.Exists() {
		ctx.Panic("Missing factor")
	}
	member := &Member{
		Address: address.Value(),
		Factor:  factor.Value(),
	}
	state := ctx.State()
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
			ctx.Log("Updated: " + member.Address.String())
			return
		}
	}
	total += member.Factor
	totalFactor.SetValue(total)
	members.GetBytes(size).SetValue(EncodeMember(member))
	ctx.Log("Appended: " + member.Address.String())
}

func divide(ctx *client.ScCallContext) {
	amount := ctx.Balances().Balance(client.IOTA)
	if amount == 0 {
		ctx.Panic("Nothing to divide")
	}
	state := ctx.State()
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
			ctx.TransferToAddress(m.Address, client.NewScTransfer(client.IOTA, part))
		}
	}
	if parts != amount {
		// note we truncated the calculations down to the nearest integer
		// there could be some small remainder left in the contract, but
		// that will be picked up in the next round as part of the balance
		remainder := amount - parts
		ctx.Log("Remainder in contract: " + ctx.Utility().String(remainder))
	}
}
