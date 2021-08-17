// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package test

import (
	"testing"
	"time"

	"github.com/iotaledger/hive.go/crypto/ed25519"
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasplib/contracts/common"
	"github.com/iotaledger/wasplib/contracts/rust/fairroulette"
	"github.com/stretchr/testify/require"
)

func setupTest(t *testing.T) *solo.Chain {
	return common.StartChainAndDeployWasmContractByName(t, fairroulette.ScName)
}

func TestDeploy(t *testing.T) {
	chain := setupTest(t)
	_, err := chain.FindContract(fairroulette.ScName)
	require.NoError(t, err)
}

func TestBets(t *testing.T) {
	chain := setupTest(t)
	var better [10]*ed25519.KeyPair
	for i := 0; i < 10; i++ {
		better[i], _ = chain.Env.NewKeyPairWithFunds()
		req := solo.NewCallParams(ScName, FuncPlaceBet,
			ParamNumber, 3,
		).WithIotas(25)
		_, err := chain.PostRequestSync(req, better[i])
		require.NoError(t, err)
	}
	require.True(t, chain.WaitForRequestsThrough(15))
	chain.Env.AdvanceClockBy(1201 * time.Second)
	// require.True(t, chain.WaitForRequestsThrough(1))
}
