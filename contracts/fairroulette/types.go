// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package fairroulette

import "github.com/iotaledger/wasplib/client"

type BetInfo struct {
    amount int64
    better *client.ScAgent
    color int64
}

func encodeBetInfo(o *BetInfo) []byte {
    return client.NewBytesEncoder().
        Int(o.amount).
        Agent(o.better).
        Int(o.color).
        Data()
}

func decodeBetInfo(bytes []byte) *BetInfo {
    d := client.NewBytesDecoder(bytes)
    data := &BetInfo{}
    data.amount = d.Int()
    data.better = d.Agent()
    data.color = d.Int()
    return data
}
