package main

import (
	"github.com/iotaledger/wasplib/client"
)

type DonationInfo struct {
	seq      int64
	id       string
	amount   int64
	sender   string
	feedback string
	error    string
}

func main() {
}

//export onLoad
func onLoadDonateWithFeedback() {
	exports := client.NewScExports()
	exports.Add("donate")
	exports.AddProtected("withdraw")
}

//export donate
func donate() {
	sc := client.NewScContext()
	tlog := sc.TimestampedLog("l")
	request := sc.Request()
	donation := &DonationInfo{
		seq:      int64(tlog.Length()),
		id:       request.Hash(),
		amount:   request.Balance("iota"),
		sender:   request.Address(),
		feedback: request.Params().GetString("f").Value(),
		error:    "",
	}
	if donation.amount == 0 || len(donation.feedback) == 0 {
		donation.error = "error: empty feedback or donated amount = 0. The donated amount has been returned (if any)"
		if donation.amount > 0 {
			sc.Transfer(donation.sender, "iota", donation.amount)
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

//export withdraw
func withdraw() {
	sc := client.NewScContext()
	owner := sc.Contract().Owner()
	request := sc.Request()
	if request.Address() != owner {
		sc.Log("Cancel spoofed request")
		return
	}

	account := sc.Account()
	amount := account.Balance("iota")
	withdrawAmount := request.Params().GetInt("s").Value()
	if withdrawAmount == 0 || withdrawAmount > amount {
		withdrawAmount = amount
	}
	if withdrawAmount == 0 {
		sc.Log("DonateWithFeedback: withdraw. nothing to withdraw")
		return
	}

	sc.Transfer(owner, "iota", withdrawAmount)
}

func decodeDonationInfo(bytes []byte) *DonationInfo {
	decoder := client.NewBytesDecoder(bytes)
	data := &DonationInfo{}
	data.seq = decoder.Int()
	data.id = decoder.String()
	data.amount = decoder.Int()
	data.sender = decoder.String()
	data.error = decoder.String()
	data.feedback = decoder.String()
	return data
}

func encodeDonationInfo(donation *DonationInfo) []byte {
	return client.NewBytesEncoder().
		Int(donation.seq).
		String(donation.id).
		Int(donation.amount).
		String(donation.sender).
		String(donation.error).
		String(donation.feedback).
		Data()
}
