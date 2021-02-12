// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package donatewithfeedback

import (
	"github.com/iotaledger/wasp/packages/vm/wasmlib"
)

func funcDonate(ctx *wasmlib.ScFuncContext, params *FuncDonateParams) {
	donation := &Donation{
		Amount:    ctx.Incoming().Balance(wasmlib.IOTA),
		Donator:   ctx.Caller(),
		Error:     "",
		Feedback:  params.Feedback.Value(),
		Timestamp: ctx.Timestamp(),
	}
	if donation.Amount == 0 || len(donation.Feedback) == 0 {
		donation.Error = "error: empty feedback or donated amount = 0. The donated amount has been returned (if any)"
		if donation.Amount > 0 {
			ctx.TransferToAddress(donation.Donator.Address(), wasmlib.NewScTransfer(wasmlib.IOTA, donation.Amount))
			donation.Amount = 0
		}
	}
	state := ctx.State()
	log := state.GetBytesArray(VarLog)
	log.GetBytes(log.Length()).SetValue(donation.Bytes())

	largestDonation := state.GetInt(VarMaxDonation)
	totalDonated := state.GetInt(VarTotalDonation)
	if donation.Amount > largestDonation.Value() {
		largestDonation.SetValue(donation.Amount)
	}
	totalDonated.SetValue(totalDonated.Value() + donation.Amount)
}

func funcWithdraw(ctx *wasmlib.ScFuncContext, params *FuncWithdrawParams) {
	balance := ctx.Balances().Balance(wasmlib.IOTA)
	amount := params.Amount.Value()
	if amount == 0 || amount > balance {
		amount = balance
	}
	if amount == 0 {
		ctx.Log("DonateWithFeedback: nothing to withdraw")
		return
	}

	scCreator := ctx.ContractCreator().Address()
	ctx.TransferToAddress(scCreator, wasmlib.NewScTransfer(wasmlib.IOTA, amount))
}

func viewDonations(ctx *wasmlib.ScViewContext, params *ViewDonationsParams) {
	state := ctx.State()
	largestDonation := state.GetInt(VarMaxDonation)
	totalDonated := state.GetInt(VarTotalDonation)
	log := state.GetBytesArray(VarLog)
	results := ctx.Results()
	results.GetInt(VarMaxDonation).SetValue(largestDonation.Value())
	results.GetInt(VarTotalDonation).SetValue(totalDonated.Value())
	donations := results.GetMapArray(VarDonations)
	size := log.Length()
	for i := int32(0); i < size; i++ {
		di := NewDonationFromBytes(log.GetBytes(i).Value())
		donation := donations.GetMap(i)
		donation.GetInt(VarAmount).SetValue(di.Amount)
		donation.GetString(VarDonator).SetValue(di.Donator.String())
		donation.GetString(VarError).SetValue(di.Error)
		donation.GetString(VarFeedback).SetValue(di.Feedback)
		donation.GetInt(VarTimestamp).SetValue(di.Timestamp)
	}
}
