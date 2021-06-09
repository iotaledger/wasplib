// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package test

import (
	"strings"
	"testing"

	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasplib/contracts/common"
	"github.com/stretchr/testify/require"
)

func setupTest(t *testing.T) *solo.Chain {
	return common.StartChainAndDeployWasmContractByName(t, ScName)
}

func TestDeploy(t *testing.T) {
	chain := common.StartChainAndDeployWasmContractByName(t, ScName)
	_, err := chain.FindContract(ScName)
	require.NoError(t, err)
}

func TestAddMemberOk(t *testing.T) {
	chain := setupTest(t)

	_, member1Addr := chain.Env.NewKeyPairWithFunds()
	req := solo.NewCallParams(ScName, FuncMember,
		ParamAddress, member1Addr,
		ParamFactor, 100,
	)
	req.WithIotas(1)
	_, err := chain.PostRequestSync(req, nil)
	require.NoError(t, err)
}

func TestAddMemberFailMissingAddress(t *testing.T) {
	chain := setupTest(t)

	req := solo.NewCallParams(ScName, FuncMember,
		ParamFactor, 100,
	)
	req.WithIotas(1)
	_, err := chain.PostRequestSync(req, nil)
	require.Error(t, err)
	require.True(t, strings.HasSuffix(err.Error(), "missing mandatory address"))
}

func TestAddMemberFailMissingFactor(t *testing.T) {
	chain := setupTest(t)

	_, member1Addr := chain.Env.NewKeyPairWithFunds()
	req := solo.NewCallParams(ScName, FuncMember,
		ParamAddress, member1Addr,
	)
	req.WithIotas(1)
	_, err := chain.PostRequestSync(req, nil)
	require.Error(t, err)
	require.True(t, strings.HasSuffix(err.Error(), "missing mandatory factor"))
}
