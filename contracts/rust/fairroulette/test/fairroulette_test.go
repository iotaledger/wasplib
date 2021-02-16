// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package test

import (
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasplib/contracts/common"
	"github.com/stretchr/testify/require"
	"testing"
)

func setupTest(t *testing.T) *solo.Chain {
	return common.DeployContract(t, ScName)
}

func TestDeploy(t *testing.T) {
	chain := common.DeployContract(t, ScName)
	_, err := chain.FindContract(ScName)
	require.NoError(t, err)
}
