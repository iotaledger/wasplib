// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package donatewithfeedback

import "github.com/iotaledger/wasplib/client"

const keyAmount = client.Key("amount")
const keyDonations = client.Key("donations")
const keyDonator = client.Key("donator")
const keyError = client.Key("error")
const keyFeedback = client.Key("feedback")
const keyLog = client.Key("log")
const keyMaxDonation = client.Key("max_donation")
const keyTimestamp = client.Key("timestamp")
const keyTotalDonation = client.Key("total_donation")
const keyWithdrawAmount = client.Key("withdraw")

func OnLoad() {
	exports := client.NewScExports()
	exports.AddCall("donate", donate)
	exports.AddCall("withdraw", withdraw)
	exports.AddView("view_donations", viewDonations)
}

func donate(sc *client.ScCallContext) {
	donation := &DonationInfo{
		Amount:    sc.Incoming().Balance(client.IOTA),
		Donator:   sc.Caller(),
		Error:     "",
		Feedback:  sc.Params().GetString(keyFeedback).Value(),
		Timestamp: sc.Timestamp(),
	}
	if donation.Amount == 0 || len(donation.Feedback) == 0 {
		donation.Error = "error: empty feedback or donated amount = 0. The donated amount has been returned (if any)"
		if donation.Amount > 0 {
			sc.Transfer(donation.Donator, client.IOTA, donation.Amount)
			donation.Amount = 0
		}
	}
	state := sc.State()
	log := state.GetBytesArray(keyLog)
	log.GetBytes(log.Length()).SetValue(EncodeDonationInfo(donation))

	largestDonation := state.GetInt(keyMaxDonation)
	totalDonated := state.GetInt(keyTotalDonation)
	if donation.Amount > largestDonation.Value() {
		largestDonation.SetValue(donation.Amount)
	}
	totalDonated.SetValue(totalDonated.Value() + donation.Amount)
}

func withdraw(sc *client.ScCallContext) {
	scOwner := sc.Contract().Creator()
	if !sc.From(scOwner) {
		sc.Panic("Cancel spoofed request")
	}

	amount := sc.Balances().Balance(client.IOTA)
	withdrawAmount := sc.Params().GetInt(keyWithdrawAmount).Value()
	if withdrawAmount == 0 || withdrawAmount > amount {
		withdrawAmount = amount
	}
	if withdrawAmount == 0 {
		sc.Log("DonateWithFeedback: nothing to withdraw")
		return
	}

	sc.Transfer(scOwner, client.IOTA, withdrawAmount)
}

func viewDonations(sc *client.ScViewContext) {
	state := sc.State()
	largestDonation := state.GetInt(keyMaxDonation)
	totalDonated := state.GetInt(keyTotalDonation)
	log := state.GetBytesArray(keyLog)
	results := sc.Results()
	results.GetInt(keyMaxDonation).SetValue(largestDonation.Value())
	results.GetInt(keyTotalDonation).SetValue(totalDonated.Value())
	donations := results.GetMapArray(keyDonations)
	size := log.Length()
	for i := int32(0); i < size; i++ {
		di := DecodeDonationInfo(log.GetBytes(i).Value())
		donation := donations.GetMap(i)
		donation.GetInt(keyAmount).SetValue(di.Amount)
		donation.GetString(keyDonator).SetValue(di.Donator.String())
		donation.GetString(keyError).SetValue(di.Error)
		donation.GetString(keyFeedback).SetValue(di.Feedback)
		donation.GetInt(keyTimestamp).SetValue(di.Timestamp)
	}
}
