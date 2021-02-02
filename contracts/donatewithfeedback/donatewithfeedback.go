// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package donatewithfeedback

import "github.com/iotaledger/wasplib/client"

func funcDonate(ctx *client.ScCallContext) {
	donation := &DonationInfo{
		Amount:    ctx.Incoming().Balance(client.IOTA),
		Donator:   ctx.Caller(),
		Error:     "",
		Feedback:  ctx.Params().GetString(ParamFeedback).Value(),
		Timestamp: ctx.Timestamp(),
	}
	if donation.Amount == 0 || len(donation.Feedback) == 0 {
		donation.Error = "error: empty feedback or donated amount = 0. The donated amount has been returned (if any)"
		if donation.Amount > 0 {
			ctx.TransferToAddress(donation.Donator.Address(), client.NewScTransfer(client.IOTA, donation.Amount))
			donation.Amount = 0
		}
	}
	state := ctx.State()
	log := state.GetBytesArray(VarLog)
	log.GetBytes(log.Length()).SetValue(EncodeDonationInfo(donation))

	largestDonation := state.GetInt(VarMaxDonation)
	totalDonated := state.GetInt(VarTotalDonation)
	if donation.Amount > largestDonation.Value() {
		largestDonation.SetValue(donation.Amount)
	}
	totalDonated.SetValue(totalDonated.Value() + donation.Amount)
}

func funcWithdraw(ctx *client.ScCallContext) {
	scOwner := ctx.ContractCreator()
	if !ctx.From(scOwner) {
		ctx.Panic("Cancel spoofed request")
	}

	amount := ctx.Balances().Balance(client.IOTA)
	withdrawAmount := ctx.Params().GetInt(ParamWithdrawAmount).Value()
	if withdrawAmount == 0 || withdrawAmount > amount {
		withdrawAmount = amount
	}
	if withdrawAmount == 0 {
		ctx.Log("DonateWithFeedback: nothing to withdraw")
		return
	}

	ctx.TransferToAddress(scOwner.Address(), client.NewScTransfer(client.IOTA, withdrawAmount))
}

func viewDonations(ctx *client.ScViewContext) {
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
		di := DecodeDonationInfo(log.GetBytes(i).Value())
		donation := donations.GetMap(i)
		donation.GetInt(VarAmount).SetValue(di.Amount)
		donation.GetString(VarDonator).SetValue(di.Donator.String())
		donation.GetString(VarError).SetValue(di.Error)
		donation.GetString(VarFeedback).SetValue(di.Feedback)
		donation.GetInt(VarTimestamp).SetValue(di.Timestamp)
	}
}
