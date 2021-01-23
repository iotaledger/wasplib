package wasptest

import (
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/balance"
	"github.com/iotaledger/wasp/packages/kv/codec"
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasplib/govm"
	"github.com/stretchr/testify/require"
	"testing"
)

// testing direct Go SC code execution

const incName = "inccounter"

const varCounter = "counter"
const varNumRepeats = "num_repeats"

func TestIncSoloInc(t *testing.T) {
	al := solo.New(t, false, true)
	chain := al.NewChain(nil, "chain1")
	err := govm.DeployGoContract(chain, nil, incName, incName)
	require.NoError(t, err)

	req := solo.NewCall(incName, "increment").
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
	al := solo.New(t, false, true)
	chain := al.NewChain(nil, "chain1")
	err := govm.DeployGoContract(chain, nil, incName, incName)
	require.NoError(t, err)
	req := solo.NewCall(incName, "increment_repeat_many", varNumRepeats, 2).
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
	err := govm.DeployGoContract(chain, nil, incName, incName)
	require.NoError(t, err)
	req := solo.NewCall(incName, "results_test").
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
	err := govm.DeployGoContract(chain, nil, incName, incName)
	require.NoError(t, err)
	req := solo.NewCall(incName, "state_test").
		WithTransfer(balance.ColorIOTA, 1)
	ret, err := chain.PostRequest(req, nil)
	require.NoError(t, err)
	ret, err = chain.CallView(incName, "state_check")
	require.NoError(t, err)
	require.EqualValues(t, 0, len(ret))
}
