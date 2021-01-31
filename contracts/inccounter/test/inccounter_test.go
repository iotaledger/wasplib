package wasptest

import (
	"github.com/iotaledger/wasp/packages/vm/wasmhost"
	"github.com/iotaledger/wasplib/contracts/inccounter"
	"github.com/iotaledger/wasplib/govm"
	"github.com/stretchr/testify/require"
	"testing"
)

const incName = "inccounter"

const varCounter = "counter"
const varNumRepeats = "num_repeats"

func TestIncrementDeploy(t *testing.T) {
	te := govm.NewTestEnv(t, incName)
	checkCounter(te, nil)
}

func TestIncrementOnce(t *testing.T) {
	te := govm.NewTestEnv(t, incName)
	_ = te.NewCallParams("increment").Post(0)
	checkCounter(te, 1)
}

func TestIncrementTwice(t *testing.T) {
	te := govm.NewTestEnv(t, incName)
	_ = te.NewCallParams("increment").Post(0)
	_ = te.NewCallParams("increment").Post(0)
	checkCounter(te, 2)
}

func TestIncrementRepeatThrice(t *testing.T) {
	te := govm.NewTestEnv(t, incName)
	_ = te.NewCallParams("increment_repeat_many",
		varNumRepeats, 3).
		Post(1) // !!! posts to self
	te.WaitForEmptyBacklog()
	checkCounter(te, 4)
}

func TestIncrementCallIncrement(t *testing.T) {
	te := govm.NewTestEnv(t, incName)
	_ = te.NewCallParams("increment_call_increment").Post(0)
	checkCounter(te, 2)
}

func TestIncrementCallIncrementRecurse5x(t *testing.T) {
	te := govm.NewTestEnv(t, incName)
	_ = te.NewCallParams("increment_call_increment_recurse5x").Post(0)
	checkCounter(te, 6)
}

func TestIncrementPostIncrement(t *testing.T) {
	te := govm.NewTestEnv(t, incName)
	_ = te.NewCallParams("increment_post_increment").
		Post(1) // !!! posts to self
	te.WaitForEmptyBacklog()
	checkCounter(te, 2)
}

func TestIncrementLocalStateInternalCall(t *testing.T) {
	te := govm.NewTestEnv(t, incName)
	_ = te.NewCallParams("increment_local_state_internal_call").Post(0)
	checkCounter(te, 2)
}

func TestIncrementLocalStateSandboxCall(t *testing.T) {
	te := govm.NewTestEnv(t, incName)
	_ = te.NewCallParams("increment_local_state_sandbox_call").Post(0)
	if govm.WasmRunner == govm.WasmRunnerGoDirect {
		// global var in direct go execution has effect
		checkCounter(te, 2)
		return
	}
	// global var in wasm execution has no effect
	checkCounter(te, nil)
}

func TestIncrementLocalStatePost(t *testing.T) {
	te := govm.NewTestEnv(t, incName)
	_ = te.NewCallParams("increment_local_state_post").
		Post(1)
	te.WaitForEmptyBacklog()
	if govm.WasmRunner == govm.WasmRunnerGoDirect {
		// global var in direct go execution has effect
		checkCounter(te, 1)
		return
	}
	// global var in wasm execution has no effect
	checkCounter(te, nil)
}

func TestIncrementViewCounter(t *testing.T) {
	te := govm.NewTestEnv(t, incName)
	_ = te.NewCallParams("increment").Post(0)
	//TODO FIXME checkCounter(te, 1)
	ret := te.CallView("increment_view_counter")
	results := govm.GetClientMap(t, wasmhost.KeyResults, ret)
	counter := results.GetInt(inccounter.KeyCounter)
	require.True(te.T, counter.Exists())
	require.EqualValues(t, 1, counter.Value())
}

func TestIncResultsTest(t *testing.T) {
	te := govm.NewTestEnv(t, incName)
	ret := te.NewCallParams("results_test").Post(0)
	//ret = te.CallView( "results_check")
	require.EqualValues(t, 8, len(ret))
}

func TestIncStateTest(t *testing.T) {
	te := govm.NewTestEnv(t, incName)
	ret := te.NewCallParams("state_test").Post(0)
	ret = te.CallView("state_check")
	require.EqualValues(t, 0, len(ret))
}

func checkCounter(te *govm.TestEnv, expected interface{}) {
	state := te.State()
	counter := state.GetInt(inccounter.KeyCounter)
	if expected == nil {
		require.False(te.T, counter.Exists())
		return
	}
	require.True(te.T, counter.Exists())
	require.EqualValues(te.T, expected, counter.Value())
}
