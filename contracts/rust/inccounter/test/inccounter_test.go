// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package test

import (
	"fmt"
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/balance"
	"github.com/iotaledger/wasp/packages/kv"
	"github.com/iotaledger/wasp/packages/kv/codec"
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasp/packages/vm/wasmproc"
	"github.com/iotaledger/wasplib/contracts/common"
	"github.com/stretchr/testify/require"
	"sort"
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
	_, err := chain.PostRequestSync(req, nil)
	require.NoError(t, err)

	checkStateCounter(t, chain, 1)
}

func TestIncrementTwice(t *testing.T) {
	chain := setupTest(t)

	req := solo.NewCallParams(ScName, FuncIncrement)
	_, err := chain.PostRequestSync(req, nil)
	require.NoError(t, err)

	req = solo.NewCallParams(ScName, FuncIncrement)
	_, err = chain.PostRequestSync(req, nil)
	require.NoError(t, err)

	checkStateCounter(t, chain, 2)
}

func TestIncrementRepeatThrice(t *testing.T) {
	chain := setupTest(t)

	req := solo.NewCallParams(ScName, FuncRepeatMany,
		ParamNumRepeats, 3,
	).WithTransfer(balance.ColorIOTA, 1) // !!! posts to self
	_, err := chain.PostRequestSync(req, nil)
	require.NoError(t, err)

	chain.WaitForEmptyBacklog()

	checkStateCounter(t, chain, 4)
}

func TestIncrementCallIncrement(t *testing.T) {
	chain := setupTest(t)

	req := solo.NewCallParams(ScName, FuncCallIncrement)
	_, err := chain.PostRequestSync(req, nil)
	require.NoError(t, err)

	checkStateCounter(t, chain, 2)
}

func TestIncrementCallIncrementRecurse5x(t *testing.T) {
	chain := setupTest(t)

	req := solo.NewCallParams(ScName, FuncCallIncrementRecurse5x)
	_, err := chain.PostRequestSync(req, nil)
	require.NoError(t, err)

	checkStateCounter(t, chain, 6)
}

func TestIncrementPostIncrement(t *testing.T) {
	chain := setupTest(t)

	req := solo.NewCallParams(ScName, FuncPostIncrement).WithTransfer(balance.ColorIOTA, 1) // !!! posts to self
	_, err := chain.PostRequestSync(req, nil)
	require.NoError(t, err)

	chain.WaitForEmptyBacklog()

	checkStateCounter(t, chain, 2)
}

func TestIncrementLocalStateInternalCall(t *testing.T) {
	chain := setupTest(t)

	req := solo.NewCallParams(ScName, FuncLocalStateInternalCall)
	_, err := chain.PostRequestSync(req, nil)
	require.NoError(t, err)

	checkStateCounter(t, chain, 2)
}

func TestIncrementLocalStateSandboxCall(t *testing.T) {
	chain := setupTest(t)

	req := solo.NewCallParams(ScName, FuncLocalStateSandboxCall)
	_, err := chain.PostRequestSync(req, nil)
	require.NoError(t, err)

	if common.WasmRunner == 0 {
		// global var in wasm execution has no effect
		checkStateCounter(t, chain, nil)
		return
	}

	checkStateCounter(t, chain, 2)
}

func TestIncrementLocalStatePost(t *testing.T) {
	chain := setupTest(t)

	req := solo.NewCallParams(ScName, FuncLocalStatePost).WithTransfer(balance.ColorIOTA, 1) // !!! posts to self
	_, err := chain.PostRequestSync(req, nil)
	require.NoError(t, err)

	chain.WaitForEmptyBacklog()

	if common.WasmRunner == 0 {
		// global var in wasm execution has no effect
		checkStateCounter(t, chain, nil)
		return
	}

	checkStateCounter(t, chain, 1)
}

func TestLeb128(t *testing.T) {
	chain := setupTest(t)

	req := solo.NewCallParams(ScName, FuncTestLeb128)
	_, err := chain.PostRequestSync(req, nil)
	require.NoError(t, err)
	res, err := chain.CallView(
		ScName, wasmproc.ViewCopyAllState,
	)
	require.NoError(t, err)
	keys := make([]string, 0)
	for key := range res {
		keys = append(keys, string(key))
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Printf("%s: %v\n", key, res[kv.Key(key)])
	}
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
