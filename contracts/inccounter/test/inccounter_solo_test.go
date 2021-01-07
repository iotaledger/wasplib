package wasptest

import (
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/balance"
	"github.com/iotaledger/wasp/packages/kv/codec"
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasp/packages/vm/wasmproc"
	"github.com/iotaledger/wasplib/contracts/inccounter"
	"github.com/iotaledger/wasplib/wasmlocalhost"
	"github.com/stretchr/testify/require"
	"testing"
)

// testing direct Go SC code execution

const incName = "inccounter"

const varCounter = "counter"
const varNumRepeats = "num_repeats"

var contracts = map[string]func(){
	incName: inccounter.OnLoad,
}

func DeployGoContract(t *testing.T, chain *solo.Chain, contractName string) error {
	wasmproc.GoWasmVM = wasmlocalhost.NewGoVM(contracts)
	hprog, err := chain.UploadWasm(nil, []byte("go:"+contractName))
	require.NoError(t, err)
	err = chain.DeployContract(nil, contractName, hprog)
	return err
}

func TestIncSoloInc(t *testing.T) {
	al := solo.New(t, false, true)
	chain := al.NewChain(nil, "chain1")
	err := DeployGoContract(t, chain, incName)
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
	err := DeployGoContract(t, chain, incName)
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
	err := DeployGoContract(t, chain, incName)
	require.NoError(t, err)
	req := solo.NewCall(incName, "results_test").
		WithTransfer(balance.ColorIOTA, 1)
	ret, err := chain.PostRequest(req, nil)
	require.NoError(t, err)
	//ret, err = chain.CallView(incName, "results_check")
	//require.NoError(t, err)
	require.EqualValues(t, 6, len(ret))
}

func TestIncSoloStateTest(t *testing.T) {
	al := solo.New(t, false, true)
	chain := al.NewChain(nil, "chain1")
	err := DeployGoContract(t, chain, incName)
	require.NoError(t, err)
	req := solo.NewCall(incName, "state_test").
		WithTransfer(balance.ColorIOTA, 1)
	ret, err := chain.PostRequest(req, nil)
	require.NoError(t, err)
	ret, err = chain.CallView(incName, "state_check")
	require.NoError(t, err)
	require.EqualValues(t, 0, len(ret))
}
