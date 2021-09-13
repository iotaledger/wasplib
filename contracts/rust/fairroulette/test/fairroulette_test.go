// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package test

import (
	"testing"
	"time"

	"github.com/iotaledger/wasplib/contracts/common"
	"github.com/iotaledger/wasplib/contracts/rust/fairroulette"
	"github.com/stretchr/testify/require"
)

func setupTest(t *testing.T) *common.SoloContext {
	return common.NewSoloContract(t, fairroulette.ScName, fairroulette.OnLoad)
}

func TestDeploy(t *testing.T) {
	ctx := setupTest(t)
	_, err := ctx.Chain.FindContract(fairroulette.ScName)
	require.NoError(t, err)
}

func TestBets(t *testing.T) {
	ctx := setupTest(t)
	var better [10]*common.SoloAgent
	for i := 0; i < 10; i++ {
		better[i] = common.NewSoloAgent(ctx)
		placeBet := fairroulette.ScFuncs.PlaceBet(ctx)
		placeBet.Params.Number().SetValue(3)
		placeBet.Func.TransferIotas(25).Post()
		require.NoError(t, ctx.Err)
	}
	ctx.Chain.Env.AdvanceClockBy(1201 * time.Second)
	require.True(t, ctx.WaitForRequestsThrough(15))
}
