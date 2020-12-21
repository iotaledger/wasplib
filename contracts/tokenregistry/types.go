// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package tokenregistry

import "github.com/iotaledger/wasplib/client"

type TokenInfo struct {
    created int64
    description string
    mintedBy *client.ScAgent
    owner *client.ScAgent
    supply int64
    updated int64
    userDefined string
}

func encodeTokenInfo(o *TokenInfo) []byte {
    return client.NewBytesEncoder().
        Int(o.created).
        String(o.description).
        Agent(o.mintedBy).
        Agent(o.owner).
        Int(o.supply).
        Int(o.updated).
        String(o.userDefined).
        Data()
}

func decodeTokenInfo(bytes []byte) *TokenInfo {
    d := client.NewBytesDecoder(bytes)
    data := &TokenInfo{}
    data.created = d.Int()
    data.description = d.String()
    data.mintedBy = d.Agent()
    data.owner = d.Agent()
    data.supply = d.Int()
    data.updated = d.Int()
    data.userDefined = d.String()
    return data
}
