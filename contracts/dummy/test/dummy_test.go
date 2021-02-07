package test_dummy_wasm

import (
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/stretchr/testify/require"
	"testing"
)

const (
	contractName = "dummy"
	fileName     = "../pkg/dummy_bg.wasm"
)

func TestSuccess(t *testing.T) {
	env := solo.New(t, false, false)
	chain := env.NewChain(nil, "chain1")
	defer chain.WaitForEmptyBacklog()

	err := chain.DeployWasmContract(nil, contractName, fileName)
	require.NoError(t, err)
	_, rec := chain.GetInfo()
	require.EqualValues(t, 5, len(rec))
}

func TestFail(t *testing.T) {
	env := solo.New(t, false, false)
	chain := env.NewChain(nil, "chain1")
	err := chain.DeployWasmContract(nil, contractName, fileName,
		"failInitParam", 1,
	)
	require.Error(t, err)
	_, rec := chain.GetInfo()
	require.EqualValues(t, 4, len(rec))
}

func TestFailRepeat(t *testing.T) {
	env := solo.New(t, false, false)
	chain := env.NewChain(nil, "chain1")
	err := chain.DeployWasmContract(nil, contractName, fileName,
		"failInitParam", 1,
	)
	require.Error(t, err)

	err = chain.DeployWasmContract(nil, contractName, fileName)
	require.NoError(t, err)
	_, rec := chain.GetInfo()
	require.EqualValues(t, 5, len(rec))
}
