// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package donatewithfeedback

import (
	"github.com/iotaledger/wasplib/client"
)

type DonationInfo struct {
	seq      int64
	id       *client.ScRequestId
	amount   int64
	sender   *client.ScAgent
	error    string
	feedback string
}

func OnLoad() {
	exports := client.NewScExports()
	exports.AddCall("donate", donate)
	exports.AddCall("withdraw", withdraw)
}

func donate(sc *client.ScCallContext) {
	tlog := sc.TimestampedLog("l")
	request := sc.Request()
	donation := &DonationInfo{
		seq:      int64(tlog.Length()),
		id:       request.Id(),
		amount:   request.Balance(client.IOTA),
		sender:   request.Sender(),
		error:    "",
		feedback: request.Params().GetString("f").Value(),
	}
	if donation.amount == 0 || len(donation.feedback) == 0 {
		donation.error = "error: empty feedback or donated amount = 0. The donated amount has been returned (if any)"
		if donation.amount > 0 {
			sc.Transfer(donation.sender, client.IOTA, donation.amount)
			donation.amount = 0
		}
	}
	bytes := encodeDonationInfo(donation)
	tlog.Append(request.Timestamp(), bytes)

	state := sc.State()
	largestDonation := state.GetInt("maxd")
	totalDonated := state.GetInt("total")
	if donation.amount > largestDonation.Value() {
		largestDonation.SetValue(donation.amount)
	}
	totalDonated.SetValue(totalDonated.Value() + donation.amount)
}

func withdraw(sc *client.ScCallContext) {
	scOwner := sc.Contract().Owner()
	request := sc.Request()
	if !request.From(scOwner) {
		sc.Log("Cancel spoofed request")
		return
	}

	account := sc.Account()
	amount := account.Balance(client.IOTA)
	withdrawAmount := request.Params().GetInt("s").Value()
	if withdrawAmount == 0 || withdrawAmount > amount {
		withdrawAmount = amount
	}
	if withdrawAmount == 0 {
		sc.Log("DonateWithFeedback: withdraw. nothing to withdraw")
		return
	}

	sc.Transfer(scOwner, client.IOTA, withdrawAmount)
}

func decodeDonationInfo(bytes []byte) *DonationInfo {
	decoder := client.NewBytesDecoder(bytes)
	data := &DonationInfo{}
	data.seq = decoder.Int()
	data.id = decoder.RequestId()
	data.amount = decoder.Int()
	data.sender = decoder.Agent()
	data.error = decoder.String()
	data.feedback = decoder.String()
	return data
}

func encodeDonationInfo(donation *DonationInfo) []byte {
	return client.NewBytesEncoder().
		Int(donation.seq).
		RequestId(donation.id).
		Int(donation.amount).
		Agent(donation.sender).
		String(donation.error).
		String(donation.feedback).
		Data()
}
