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

//export donate
func donate() {
	ctx := client.NewScContext()
	tlog := ctx.TimestampedLog("l")
	request := ctx.Request()
	di := &DonationInfo{
		seq:      int64(tlog.Length()),
		id:       request.Hash(),
		amount:   request.Balance("iota"),
		sender:   request.Address(),
		feedback: request.Params().GetString("f").Value(),
		error:    "",
	}
	if di.amount == 0 || len(di.feedback) == 0 {
		di.error = "error: empty feedback or donated amount = 0. The donated amount has been returned (if any)"
		if di.amount > 0 {
			ctx.Transfer(di.sender, "iota", di.amount)
			di.amount = 0
		}
	}
	data := encodeDonationInfo(di)
	tlog.Append(request.Timestamp(), data)

	state := ctx.State()
	maxd := state.GetInt("maxd")
	total := state.GetInt("total")
	if di.amount > maxd.Value() {
		maxd.SetValue(di.amount)
	}
	total.SetValue(total.Value() + di.amount)
}

//export withdraw
func withdraw() {
	ctx := client.NewScContext()
	owner := ctx.Contract().Owner()
	request := ctx.Request()
	if request.Address() != owner {
		ctx.Log("Cancel spoofed request")
		return
	}

	account := ctx.Account()
	bal := account.Balance("iota")
	withdrawSum := request.Params().GetInt("s").Value()
	if withdrawSum == 0 || withdrawSum > bal {
		withdrawSum = bal
	}
	if withdrawSum == 0 {
		ctx.Log("DonateWithFeedback: withdraw. nothing to withdraw")
		return
	}

	ctx.Transfer(owner, "iota", withdrawSum)
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

func encodeDonationInfo(data *DonationInfo) []byte {
	return client.NewBytesEncoder().
		Int(data.seq).
		String(data.id).
		Int(data.amount).
		String(data.sender).
		String(data.error).
		String(data.feedback).
		Data()
}
