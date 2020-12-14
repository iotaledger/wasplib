package test

import (
	"github.com/iotaledger/wasp/packages/vm/alone"
	"github.com/stretchr/testify/require"
	"testing"
)

const (
	contractName = "test_env"
	fileName     = "../pkg/test_dummy_bg.wasm"
)

func TestSuccess(t *testing.T) {
	glb := alone.New(t, false, false)
	chain := glb.NewChain(nil, "chain1")
	defer chain.WaitEmptyBacklog()

	err := chain.DeployWasmContract(nil, contractName, fileName)
	require.NoError(t, err)
	_, _, rec := chain.GetInfo()
	require.EqualValues(t, 4, len(rec))
}

func TestFail(t *testing.T) {
	glb := alone.New(t, false, false)
	chain := glb.NewChain(nil, "chain1")
	err := chain.DeployWasmContract(nil, contractName, fileName,
		"failParam", 1,
	)
	require.Error(t, err)
	_, _, rec := chain.GetInfo()
	require.EqualValues(t, 3, len(rec))
}

func TestFailRepeat(t *testing.T) {
	glb := alone.New(t, false, false)
	chain := glb.NewChain(nil, "chain1")
	err := chain.DeployWasmContract(nil, contractName, fileName,
		"failParam", 1,
	)
	require.Error(t, err)

	err = chain.DeployWasmContract(nil, contractName, fileName)
	require.NoError(t, err)
	_, _, rec := chain.GetInfo()
	require.EqualValues(t, 4, len(rec))
}
