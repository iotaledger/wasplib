// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package test

import (
	"testing"
	"time"

	"github.com/iotaledger/wasp/contracts/rust/fairroulette"
	"github.com/iotaledger/wasp/packages/vm/wasmsolo"
	"github.com/stretchr/testify/require"
)

func setupTest(t *testing.T) *wasmsolo.SoloContext {
	return wasmsolo.NewSoloContext(t, fairroulette.ScName, fairroulette.OnLoad)
}

func TestDeploy(t *testing.T) {
	ctx := setupTest(t)
	require.NoError(t, ctx.ContractExists(fairroulette.ScName))
}

func TestBets(t *testing.T) {
	ctx := setupTest(t)
	var better [10]*wasmsolo.SoloAgent
	for i := 0; i < 10; i++ {
		better[i] = ctx.NewSoloAgent()
		placeBet := fairroulette.ScFuncs.PlaceBet(ctx)
		placeBet.Params.Number().SetValue(3)
		placeBet.Func.TransferIotas(25).Post()
		require.NoError(t, ctx.Err)
	}
	ctx.AdvanceClockBy(1201 * time.Second)
	require.True(t, ctx.WaitForPendingRequests(1))
}
