// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// +build wasmtest

package wasptest

import (
	"github.com/iotaledger/wasp/contracts/rust/inccounter"
	"github.com/iotaledger/wasp/contracts/testenv"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIncrementDeploy(t *testing.T) {
	te := testenv.NewTestEnv(t, inccounter.ScName)
	checkStateCounter(te, nil)
}

func TestIncrementOnce(t *testing.T) {
	te := testenv.NewTestEnv(t, inccounter.ScName)
	_ = te.NewCallParams(inccounter.FuncIncrement).Post(0)
	checkStateCounter(te, 1)
}

func TestIncrementTwice(t *testing.T) {
	te := testenv.NewTestEnv(t, inccounter.ScName)
	_ = te.NewCallParams(inccounter.FuncIncrement).Post(0)
	_ = te.NewCallParams(inccounter.FuncIncrement).Post(0)
	checkStateCounter(te, 2)
}

func TestIncrementRepeatThrice(t *testing.T) {
	te := testenv.NewTestEnv(t, inccounter.ScName)
	_ = te.NewCallParams(inccounter.FuncRepeatMany,
		inccounter.ParamNumRepeats, 3).
		Post(1) // !!! posts to self
	te.WaitForEmptyBacklog()
	checkStateCounter(te, 4)
}

func TestIncrementCallIncrement(t *testing.T) {
	te := testenv.NewTestEnv(t, inccounter.ScName)
	_ = te.NewCallParams(inccounter.FuncCallIncrement).Post(0)
	checkStateCounter(te, 2)
}

func TestIncrementCallIncrementRecurse5x(t *testing.T) {
	te := testenv.NewTestEnv(t, inccounter.ScName)
	_ = te.NewCallParams(inccounter.FuncCallIncrementRecurse5x).Post(0)
	checkStateCounter(te, 6)
}

func TestIncrementPostIncrement(t *testing.T) {
	te := testenv.NewTestEnv(t, inccounter.ScName)
	_ = te.NewCallParams(inccounter.FuncPostIncrement).
		Post(1) // !!! posts to self
	te.WaitForEmptyBacklog()
	checkStateCounter(te, 2)
}

func TestIncrementLocalStateInternalCall(t *testing.T) {
	te := testenv.NewTestEnv(t, inccounter.ScName)
	_ = te.NewCallParams(inccounter.FuncLocalStateInternalCall).Post(0)
	checkStateCounter(te, 2)
}

func TestIncrementLocalStateSandboxCall(t *testing.T) {
	te := testenv.NewTestEnv(t, inccounter.ScName)
	_ = te.NewCallParams(inccounter.FuncLocalStateSandboxCall).Post(0)
	if testenv.WasmRunner == testenv.WasmRunnerGoDirect {
		// global var in direct go execution has effect
		checkStateCounter(te, 2)
		return
	}
	// global var in wasm execution has no effect
	checkStateCounter(te, nil)
}

func TestIncrementLocalStatePost(t *testing.T) {
	te := testenv.NewTestEnv(t, inccounter.ScName)
	_ = te.NewCallParams(inccounter.FuncLocalStatePost).
		Post(1)
	te.WaitForEmptyBacklog()
	if testenv.WasmRunner == testenv.WasmRunnerGoDirect {
		// global var in direct go execution has effect
		checkStateCounter(te, 1)
		return
	}
	// global var in wasm execution has no effect
	checkStateCounter(te, nil)
}

func TestIncrementViewCounter(t *testing.T) {
	te := testenv.NewTestEnv(t, inccounter.ScName)
	_ = te.NewCallParams(inccounter.FuncIncrement).Post(0)
	checkStateCounter(te, 1)

	ret := te.CallView(inccounter.ViewGetCounter)
	results := te.Results(ret)
	counter := results.GetInt(inccounter.VarCounter)
	require.True(te.T, counter.Exists())
	require.EqualValues(t, 1, counter.Value())
}

func TestIncResultsTest(t *testing.T) {
	te := testenv.NewTestEnv(t, inccounter.ScName)
	ret := te.NewCallParams(inccounter.FuncResultsTest).Post(0)
	//ret = te.CallView( inccounter.ViewResultsCheck)
	require.EqualValues(t, 8, len(ret))
}

func TestIncStateTest(t *testing.T) {
	te := testenv.NewTestEnv(t, inccounter.ScName)
	ret := te.NewCallParams(inccounter.FuncStateTest).Post(0)
	ret = te.CallView(inccounter.ViewStateCheck)
	require.EqualValues(t, 0, len(ret))
}

func checkStateCounter(te *testenv.TestEnv, expected interface{}) {
	state := te.State()
	counter := state.GetInt(inccounter.VarCounter)
	if expected == nil {
		require.False(te.T, counter.Exists())
		return
	}
	require.True(te.T, counter.Exists())
	require.EqualValues(te.T, expected, counter.Value())
}
