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

func donate(sc *client.ScCallContext) {
	donation := &DonationInfo{
		Amount:    sc.Incoming().Balance(client.IOTA),
		Donator:   sc.Caller(),
		Error:     "",
		Feedback:  sc.Params().GetString(KeyFeedback).Value(),
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
	log := state.GetBytesArray(KeyLog)
	log.GetBytes(log.Length()).SetValue(EncodeDonationInfo(donation))

	largestDonation := state.GetInt(KeyMaxDonation)
	totalDonated := state.GetInt(KeyTotalDonation)
	if donation.Amount > largestDonation.Value() {
		largestDonation.SetValue(donation.Amount)
	}
	totalDonated.SetValue(totalDonated.Value() + donation.Amount)
}

func withdraw(sc *client.ScCallContext) {
	scOwner := sc.ContractCreator()
	if !sc.From(scOwner) {
		sc.Panic("Cancel spoofed request")
	}

	amount := sc.Balances().Balance(client.IOTA)
	withdrawAmount := sc.Params().GetInt(KeyWithdrawAmount).Value()
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
	largestDonation := state.GetInt(KeyMaxDonation)
	totalDonated := state.GetInt(KeyTotalDonation)
	log := state.GetBytesArray(KeyLog)
	results := sc.Results()
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
