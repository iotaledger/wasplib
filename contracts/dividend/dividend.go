// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package dividend

import (
	"github.com/iotaledger/wasp/packages/vm/wasmlib"
)

func funcDivide(ctx *wasmlib.ScFuncContext, params *FuncDivideParams) {
	amount := ctx.Balances().Balance(wasmlib.IOTA)
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
		m := NewMemberFromBytes(members.GetBytes(i).Value())
		part := amount * m.Factor / total
		if part != 0 {
			parts += part
			ctx.TransferToAddress(m.Address, wasmlib.NewScTransfer(wasmlib.IOTA, part))
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

func funcMember(ctx *wasmlib.ScFuncContext, params *FuncMemberParams) {
	member := &Member{
		Address: params.Address.Value(),
		Factor:  params.Factor.Value(),
	}
	state := ctx.State()
	totalFactor := state.GetInt(VarTotalFactor)
	total := totalFactor.Value()
	members := state.GetBytesArray(VarMembers)
	size := members.Length()
	for i := int32(0); i < size; i++ {
		m := NewMemberFromBytes(members.GetBytes(i).Value())
		if m.Address.Equals(member.Address) {
			total -= m.Factor
			total += member.Factor
			totalFactor.SetValue(total)
			members.GetBytes(i).SetValue(member.Bytes())
			ctx.Log("Updated: " + member.Address.String())
			return
		}
	}
	total += member.Factor
	totalFactor.SetValue(total)
	members.GetBytes(size).SetValue(member.Bytes())
	ctx.Log("Appended: " + member.Address.String())
}
