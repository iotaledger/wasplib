// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package test

import (
	"github.com/iotaledger/wasp/packages/kv/codec"
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasplib/contracts/common"
	"github.com/stretchr/testify/require"
	"testing"
)

func setupTest(t *testing.T) *solo.Chain {
	return common.StartChainAndDeployWasmContractByName(t, ScName)
}

func TestDeploy(t *testing.T) {
	chain := common.StartChainAndDeployWasmContractByName(t, ScName)
	_, err := chain.FindContract(ScName)
	require.NoError(t, err)
}

func TestFuncHelloWorld(t *testing.T) {
	chain := setupTest(t)

	req := solo.NewCallParams(ScName, FuncHelloWorld)
	_, err := chain.PostRequestSync(req, nil)
	require.NoError(t, err)
}

func TestViewGetHelloWorld(t *testing.T) {
	chain := setupTest(t)

	res, err := chain.CallView(
		ScName, ViewGetHelloWorld,
	)
	require.NoError(t, err)
	hw, _, err := codec.DecodeString(res[VarHelloWorld])
	require.NoError(t, err)
	require.EqualValues(t, "Hello, world!", hw)
}
