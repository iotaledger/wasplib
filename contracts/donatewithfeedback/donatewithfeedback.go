// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package donatewithfeedback

import "github.com/iotaledger/wasplib/client"

const KeyAmount = client.Key("amount")
const KeyDonations = client.Key("donations")
const KeyDonator = client.Key("donator")
const KeyError = client.Key("error")
const KeyFeedback = client.Key("feedback")
const KeyLog = client.Key("log")
const KeyMaxDonation = client.Key("max_donation")
const KeyTimestamp = client.Key("timestamp")
const KeyTotalDonation = client.Key("total_donation")
const KeyWithdrawAmount = client.Key("withdraw")

func OnLoad() {
	exports := client.NewScExports()
	exports.AddCall("donate", donate)
	exports.AddCall("withdraw", withdraw)
	exports.AddView("view_donations", viewDonations)
}

func donate(ctx *client.ScCallContext) {
	donation := &DonationInfo{
		Amount:    ctx.Incoming().Balance(client.IOTA),
		Donator:   ctx.Caller(),
		Error:     "",
		Feedback:  ctx.Params().GetString(KeyFeedback).Value(),
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
	log := state.GetBytesArray(KeyLog)
	log.GetBytes(log.Length()).SetValue(EncodeDonationInfo(donation))

	largestDonation := state.GetInt(KeyMaxDonation)
	totalDonated := state.GetInt(KeyTotalDonation)
	if donation.Amount > largestDonation.Value() {
		largestDonation.SetValue(donation.Amount)
	}
	totalDonated.SetValue(totalDonated.Value() + donation.Amount)
}

func withdraw(ctx *client.ScCallContext) {
	scOwner := ctx.ContractCreator()
	if !ctx.From(scOwner) {
		ctx.Panic("Cancel spoofed request")
	}

	amount := ctx.Balances().Balance(client.IOTA)
	withdrawAmount := ctx.Params().GetInt(KeyWithdrawAmount).Value()
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
	largestDonation := state.GetInt(KeyMaxDonation)
	totalDonated := state.GetInt(KeyTotalDonation)
	log := state.GetBytesArray(KeyLog)
	results := ctx.Results()
	results.GetInt(KeyMaxDonation).SetValue(largestDonation.Value())
	results.GetInt(KeyTotalDonation).SetValue(totalDonated.Value())
	donations := results.GetMapArray(KeyDonations)
	size := log.Length()
	for i := int32(0); i < size; i++ {
		di := DecodeDonationInfo(log.GetBytes(i).Value())
		donation := donations.GetMap(i)
		donation.GetInt(KeyAmount).SetValue(di.Amount)
		donation.GetString(KeyDonator).SetValue(di.Donator.String())
		donation.GetString(KeyError).SetValue(di.Error)
		donation.GetString(KeyFeedback).SetValue(di.Feedback)
		donation.GetInt(KeyTimestamp).SetValue(di.Timestamp)
	}
}
