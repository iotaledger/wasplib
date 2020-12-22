// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package donatewithfeedback

import "github.com/iotaledger/wasplib/client"

type DonationInfo struct {
	amount    int64
	donator   *client.ScAgent
	error     string
	feedback  string
	timestamp int64
}

func encodeDonationInfo(o *DonationInfo) []byte {
	return client.NewBytesEncoder().
		Int(o.amount).
		Agent(o.donator).
		String(o.error).
		String(o.feedback).
		Int(o.timestamp).
		Data()
}

func decodeDonationInfo(bytes []byte) *DonationInfo {
	d := client.NewBytesDecoder(bytes)
	data := &DonationInfo{}
	data.amount = d.Int()
	data.donator = d.Agent()
	data.error = d.String()
	data.feedback = d.String()
	data.timestamp = d.Int()
	return data
}
