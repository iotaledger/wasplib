// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package donatewithfeedback

import "github.com/iotaledger/wasplib/client"

type DonationInfo struct {
	Amount    int64
	Donator   *client.ScAgent
	Error     string
	Feedback  string
	Timestamp int64
}

func EncodeDonationInfo(o *DonationInfo) []byte {
	return client.NewBytesEncoder().
		Int(o.Amount).
		Agent(o.Donator).
		String(o.Error).
		String(o.Feedback).
		Int(o.Timestamp).
		Data()
}

func DecodeDonationInfo(bytes []byte) *DonationInfo {
	decode := client.NewBytesDecoder(bytes)
	data := &DonationInfo{}
	data.Amount = decode.Int()
	data.Donator = decode.Agent()
	data.Error = decode.String()
	data.Feedback = decode.String()
	data.Timestamp = decode.Int()
	return data
}
