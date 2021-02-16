// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package test

import (
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/balance"
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

func TestStateAfterDeploy(t *testing.T) {
	chain := common.StartChainAndDeployWasmContractByName(t, ScName)

	checkStateCounter(t, chain, nil)
}

func TestIncrementOnce(t *testing.T) {
	chain := setupTest(t)

	req := solo.NewCallParams(ScName, FuncIncrement)
	_, err := chain.PostRequest(req, nil)
	require.NoError(t, err)

	checkStateCounter(t, chain, 1)
}

func TestIncrementTwice(t *testing.T) {
	chain := setupTest(t)

	req := solo.NewCallParams(ScName, FuncIncrement)
	_, err := chain.PostRequest(req, nil)
	require.NoError(t, err)

	req = solo.NewCallParams(ScName, FuncIncrement)
	_, err = chain.PostRequest(req, nil)
	require.NoError(t, err)

	checkStateCounter(t, chain, 2)
}

func TestIncrementRepeatThrice(t *testing.T) {
	chain := setupTest(t)

	req := solo.NewCallParams(ScName, FuncRepeatMany,
		ParamNumRepeats, 3,
	).WithTransfer(balance.ColorIOTA, 1) // !!! posts to self
	_, err := chain.PostRequest(req, nil)
	require.NoError(t, err)

	chain.WaitForEmptyBacklog()

	checkStateCounter(t, chain, 4)
}

func TestIncrementCallIncrement(t *testing.T) {
	chain := setupTest(t)

	req := solo.NewCallParams(ScName, FuncCallIncrement)
	_, err := chain.PostRequest(req, nil)
	require.NoError(t, err)

	checkStateCounter(t, chain, 2)
}

func TestIncrementCallIncrementRecurse5x(t *testing.T) {
	chain := setupTest(t)

	req := solo.NewCallParams(ScName, FuncCallIncrementRecurse5x)
	_, err := chain.PostRequest(req, nil)
	require.NoError(t, err)

	checkStateCounter(t, chain, 6)
}

func TestIncrementPostIncrement(t *testing.T) {
	chain := setupTest(t)

	req := solo.NewCallParams(ScName, FuncPostIncrement).WithTransfer(balance.ColorIOTA, 1) // !!! posts to self
	_, err := chain.PostRequest(req, nil)
	require.NoError(t, err)

	chain.WaitForEmptyBacklog()

	checkStateCounter(t, chain, 2)
}

func TestIncrementLocalStateInternalCall(t *testing.T) {
	chain := setupTest(t)

	req := solo.NewCallParams(ScName, FuncLocalStateInternalCall)
	_, err := chain.PostRequest(req, nil)
	require.NoError(t, err)

	checkStateCounter(t, chain, 2)
}

func TestIncrementLocalStateSandboxCall(t *testing.T) {
	chain := setupTest(t)

	req := solo.NewCallParams(ScName, FuncLocalStateSandboxCall)
	_, err := chain.PostRequest(req, nil)
	require.NoError(t, err)

	// global var in wasm execution has no effect
	checkStateCounter(t, chain, nil)
}

func TestIncrementLocalStatePost(t *testing.T) {
	chain := setupTest(t)

	req := solo.NewCallParams(ScName, FuncLocalStatePost).WithTransfer(balance.ColorIOTA, 1) // !!! posts to self
	_, err := chain.PostRequest(req, nil)
	require.NoError(t, err)

	chain.WaitForEmptyBacklog()

	// global var in wasm execution has no effect
	checkStateCounter(t, chain, nil)
}

func checkStateCounter(t *testing.T, chain *solo.Chain, expected interface{}) {
	res, err := chain.CallView(
		ScName, ViewGetCounter,
	)
	require.NoError(t, err)
	counter, exists, err := codec.DecodeInt64(res[VarCounter])
	require.NoError(t, err)
	if expected == nil {
		require.False(t, exists)
		return
	}
	require.True(t, exists)
	require.EqualValues(t, expected, counter)
}
