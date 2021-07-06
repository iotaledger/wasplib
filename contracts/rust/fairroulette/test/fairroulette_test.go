// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package test

import (
	"testing"

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
