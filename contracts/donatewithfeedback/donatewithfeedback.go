// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package donatewithfeedback

import (
	"github.com/iotaledger/wasplib/client"
)

const (
	keyAmount         = client.Key("amount")
	keyData           = client.Key("data")
	keyDonations      = client.Key("donations")
	keyDonator        = client.Key("donator")
	keyError          = client.Key("error")
	keyFeedback       = client.Key("feedback")
	keyLog            = client.Key("log")
	keyMaxDonation    = client.Key("maxDonation")
	keyTimestamp      = client.Key("timestamp")
	keyTotalDonation  = client.Key("totalDonation")
	keyWithdrawAmount = client.Key("withdraw")
)

type DonationInfo struct {
	seq      int64
	donator  *client.ScAgent
	amount   int64
	feedback string
	error    string
}

func OnLoad() {
	exports := client.NewScExports()
	exports.AddCall("donate", donate)
	exports.AddCall("withdraw", withdraw)
	exports.AddView("viewDonations", viewDonations)
}

func donate(sc *client.ScCallContext) {
	tlog := sc.TimestampedLog(keyLog)
	donation := &DonationInfo{
		seq:      int64(tlog.Length()),
		amount:   sc.Incoming().Balance(client.IOTA),
		donator:  sc.Caller(),
		error:    "",
		feedback: sc.Params().GetString(keyFeedback).Value(),
	}
	if donation.amount == 0 || len(donation.feedback) == 0 {
		donation.error = "error: empty feedback or donated amount = 0. The donated amount has been returned (if any)"
		if donation.amount > 0 {
			sc.Transfer(donation.donator, client.IOTA, donation.amount)
			donation.amount = 0
		}
	}
	bytes := encodeDonationInfo(donation)
	tlog.Append(sc.Timestamp(), bytes)

	state := sc.State()
	largestDonation := state.GetInt(keyMaxDonation)
	totalDonated := state.GetInt(keyTotalDonation)
	if donation.amount > largestDonation.Value() {
		largestDonation.SetValue(donation.amount)
	}
	totalDonated.SetValue(totalDonated.Value() + donation.amount)
}

func withdraw(sc *client.ScCallContext) {
	scOwner := sc.Contract().Owner()
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
	tlog := sc.TimestampedLog(keyLog)
	results := sc.Results()
	results.GetInt(keyMaxDonation).SetValue(largestDonation.Value())
	results.GetInt(keyTotalDonation).SetValue(totalDonated.Value())
	donations := results.GetMapArray(keyDonations)
	size := tlog.Length()
	for i := int32(0); i < size; i++ {
		log := tlog.GetMap(i)
		donation := donations.GetMap(i)
		donation.GetInt(keyTimestamp).SetValue(log.GetInt(keyTimestamp).Value())
		bytes := log.GetBytes(keyData).Value()
		di := decodeDonationInfo(bytes)
		donation.GetInt(keyAmount).SetValue(di.amount)
		donation.GetString(keyFeedback).SetValue(di.feedback)
		donation.GetString(keyDonator).SetValue(di.donator.String())
		donation.GetString(keyError).SetValue(di.error)
	}
}

func decodeDonationInfo(bytes []byte) *DonationInfo {
	decoder := client.NewBytesDecoder(bytes)
	data := &DonationInfo{}
	data.seq = decoder.Int()
	data.donator = decoder.Agent()
	data.amount = decoder.Int()
	data.feedback = decoder.String()
	data.error = decoder.String()
	return data
}

func encodeDonationInfo(donation *DonationInfo) []byte {
	return client.NewBytesEncoder().
		Int(donation.seq).
		Agent(donation.donator).
		Int(donation.amount).
		String(donation.feedback).
		String(donation.error).
		Data()
}
