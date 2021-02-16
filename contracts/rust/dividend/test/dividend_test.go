// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package test

import (
	"github.com/iotaledger/wasp/contracts/common"
	"github.com/iotaledger/wasp/packages/solo"
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

func TestAddMemberOk(t *testing.T) {
	chain := setupTest(t)

	member1 := chain.Env.NewSignatureSchemeWithFunds()
	req := solo.NewCallParams(ScName, FuncMember,
		ParamAddress, member1.Address(),
		ParamFactor, 100,
	)
	_, err := chain.PostRequest(req, nil)
	require.NoError(t, err)
}

func TestAddMemberFailMissingAddress(t *testing.T) {
	chain := setupTest(t)

	req := solo.NewCallParams(ScName, FuncMember,
		ParamFactor, 100,
	)
	_, err := chain.PostRequest(req, nil)
	require.Error(t, err)
}

func TestAddMemberFailMissingFactor(t *testing.T) {
	chain := setupTest(t)

	member1 := chain.Env.NewSignatureSchemeWithFunds()
	req := solo.NewCallParams(ScName, FuncMember,
		ParamAddress, member1.Address(),
	)
	_, err := chain.PostRequest(req, nil)
	require.Error(t, err)
}
