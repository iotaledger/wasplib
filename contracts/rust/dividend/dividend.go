// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package dividend

import (
	"github.com/iotaledger/wasplib/packages/vm/wasmlib"
)

func funcDivide(ctx wasmlib.ScFuncContext, params *FuncDivideParams) {
	amount := ctx.Balances().Balance(wasmlib.IOTA)
	if amount == 0 {
		ctx.Panic("nothing to divide")
	}
	state := ctx.State()
	total := state.GetInt64(VarTotalFactor).Value()
	members := state.GetMap(VarMembers)
	memberList := state.GetAddressArray(VarMemberList)
	size := memberList.Length()
	parts := int64(0)
	for i := int32(0); i < size; i++ {
		address := memberList.GetAddress(i).Value()
		factor := members.GetInt64(address).Value()
		share := amount * factor / total
		if share != 0 {
			parts += share
			transfers := wasmlib.NewScTransfer(wasmlib.IOTA, share)
			ctx.TransferToAddress(address, transfers)
		}
	}
	if parts != amount {
		// note we truncated the calculations down to the nearest integer
		// there could be some small remainder left in the contract, but
		// that will be picked up in the next round as part of the balance
		remainder := ctx.Utility().String(amount - parts)
		ctx.Log("remainder in contract: " + remainder)
	}
}

func funcMember(ctx wasmlib.ScFuncContext, params *FuncMemberParams) {
	state := ctx.State()
	members := state.GetMap(VarMembers)
	address := params.Address.Value()
	currentFactor := members.GetInt64(address)
	if !currentFactor.Exists() {
		// add new address to member list
		memberList := state.GetAddressArray(VarMemberList)
		memberList.GetAddress(memberList.Length()).SetValue(address)
	}
	factor := params.Factor.Value()
	totalFactor := state.GetInt64(VarTotalFactor)
	newTotalFactor := totalFactor.Value() - currentFactor.Value() + factor
	totalFactor.SetValue(newTotalFactor)
	currentFactor.SetValue(factor)
}

func viewGetFactor(ctx wasmlib.ScViewContext, params *ViewGetFactorParams) {
	address := params.Address.Value()
	members := ctx.State().GetMap(VarMembers)
	factor := members.GetInt64(address).Value()
	ctx.Results().GetInt64(VarFactor).SetValue(factor)
}
