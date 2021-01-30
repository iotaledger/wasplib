package wasptest

import (
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/balance"
	"github.com/iotaledger/wasp/packages/coretypes"
	"github.com/iotaledger/wasp/packages/kv/codec"
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasp/packages/vm/wasmhost"
	"github.com/stretchr/testify/require"
	"testing"
)

const incName = "inccounter"
const varCounter = "counter"
const varNumRepeats = "num_repeats"

var incFile = wasmhost.WasmPath("inccounter_bg.wasm")
var incHname = coretypes.Hn(incName)

func TestIncSoloInc(t *testing.T) {
	al := solo.New(t, false, true)
	chain := al.NewChain(nil, "chain1")
	err := chain.DeployWasmContract(nil, incName, incFile)
	require.NoError(t, err)
	req := solo.NewCallParams(incName, "increment").
		WithTransfer(balance.ColorIOTA, 1)
	_, err = chain.PostRequest(req, nil)
	require.NoError(t, err)
	ret, err := chain.CallView(incName, "increment_view_counter")
	require.NoError(t, err)
	counter, _, err := codec.DecodeInt64(ret.MustGet(varCounter))
	require.NoError(t, err)
	require.EqualValues(t, 1, counter)
}

func TestIncSoloRepeatMany(t *testing.T) {
	al := solo.New(t, true, true)
	chain := al.NewChain(nil, "chain1")
	err := chain.DeployWasmContract(nil, incName, incFile)
	require.NoError(t, err)
	req := solo.NewCallParams(incName, "increment_repeat_many", varNumRepeats, 2).
		WithTransfer(balance.ColorIOTA, 1)
	_, err = chain.PostRequest(req, nil)
	require.NoError(t, err)
	chain.WaitForEmptyBacklog()
	ret, err := chain.CallView(incName, "increment_view_counter")
	require.NoError(t, err)
	counter, _, err := codec.DecodeInt64(ret.MustGet(varCounter))
	require.NoError(t, err)
	require.EqualValues(t, 3, counter)
}

func TestIncSoloResultsTest(t *testing.T) {
	al := solo.New(t, false, true)
	chain := al.NewChain(nil, "chain1")
	err := chain.DeployWasmContract(nil, incName, incFile)
	require.NoError(t, err)
	req := solo.NewCallParams(incName, "results_test").
		WithTransfer(balance.ColorIOTA, 1)
	ret, err := chain.PostRequest(req, nil)
	require.NoError(t, err)
	//ret, err = chain.CallView(incName, "results_check")
	//require.NoError(t, err)
	require.EqualValues(t, 8, len(ret))
}

func TestIncSoloStateTest(t *testing.T) {
	al := solo.New(t, false, true)
	chain := al.NewChain(nil, "chain1")
	err := chain.DeployWasmContract(nil, incName, incFile)
	require.NoError(t, err)
	req := solo.NewCallParams(incName, "state_test").
		WithTransfer(balance.ColorIOTA, 1)
	ret, err := chain.PostRequest(req, nil)
	require.NoError(t, err)
	ret, err = chain.CallView(incName, "state_check")
	require.NoError(t, err)
	require.EqualValues(t, 0, len(ret))
}
