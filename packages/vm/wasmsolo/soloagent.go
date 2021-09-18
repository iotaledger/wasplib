// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package wasmsolo

import (
	"github.com/iotaledger/goshimmer/packages/ledgerstate"
	"github.com/iotaledger/hive.go/crypto/ed25519"
	"github.com/iotaledger/wasp/packages/iscp/colored"
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasplib/packages/vm/wasmlib"
	"github.com/stretchr/testify/require"
)

type SoloAgent struct {
	env     *solo.Solo
	pair    *ed25519.KeyPair
	address ledgerstate.Address
}

func NewSoloAgent(env *solo.Solo) *SoloAgent {
	agent := &SoloAgent{env: env}
	agent.pair, agent.address = agent.env.NewKeyPairWithFunds()
	return agent
}

func (a *SoloAgent) ScAddress() wasmlib.ScAddress {
	return wasmlib.NewScAddressFromBytes(a.address.Bytes())
}

func (a *SoloAgent) ScAgentID() wasmlib.ScAgentID {
	return a.ScAddress().AsAgentID()
}

func (a *SoloAgent) Balance(color ...wasmlib.ScColor) int64 {
	switch len(color) {
	case 0:
		return int64(a.env.GetAddressBalance(a.address, colored.IOTA))
	case 1:
		col, err := colored.ColorFromBytes(color[0].Bytes())
		require.NoError(a.env.T, err)
		return int64(a.env.GetAddressBalance(a.address, col))
	default:
		require.Fail(a.env.T, "too many color arguments")
		return 0
	}
}

func (a *SoloAgent) Mint(amount int64) (wasmlib.ScColor, error) {
	color, err := a.env.MintTokens(a.pair, uint64(amount))
	return wasmlib.NewScColorFromBytes(color.Bytes()), err
}
