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
		amount:    sc.Incoming().Balance(client.IOTA),
		donator:   sc.Caller(),
		error:     "",
		feedback:  sc.Params().GetString(keyFeedback).Value(),
		timestamp: sc.Timestamp(),
	}
	if donation.amount == 0 || len(donation.feedback) == 0 {
		donation.error = "error: empty feedback or donated amount = 0. The donated amount has been returned (if any)"
		if donation.amount > 0 {
			sc.Transfer(donation.donator, client.IOTA, donation.amount)
			donation.amount = 0
		}
	}
	state := sc.State()
	log := state.GetBytesArray(keyLog)
	log.GetBytes(log.Length()).SetValue(encodeDonationInfo(donation))

	largestDonation := state.GetInt(keyMaxDonation)
	totalDonated := state.GetInt(keyTotalDonation)
	if donation.amount > largestDonation.Value() {
		largestDonation.SetValue(donation.amount)
	}
	totalDonated.SetValue(totalDonated.Value() + donation.amount)
}

func withdraw(sc *client.ScCallContext) {
	scOwner := sc.Contract().Creator()
	if !sc.From(scOwner) {
		sc.Log("Cancel spoofed request")
		return
	}

	amount := sc.Balances().Balance(client.IOTA)
	withdrawAmount := sc.Params().GetInt(keyWithdrawAmount).Value()
	if withdrawAmount == 0 || withdrawAmount > amount {
		withdrawAmount = amount
	}
	if withdrawAmount == 0 {
		sc.Log("DonateWithFeedback: withdraw. nothing to withdraw")
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
		di := decodeDonationInfo(log.GetBytes(i).Value())
		donation := donations.GetMap(i)
		donation.GetInt(keyAmount).SetValue(di.amount)
		donation.GetString(keyDonator).SetValue(di.donator.String())
		donation.GetString(keyError).SetValue(di.error)
		donation.GetString(keyFeedback).SetValue(di.feedback)
		donation.GetInt(keyTimestamp).SetValue(di.timestamp)
	}
}
