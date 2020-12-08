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
	e := alone.New(t, false, false)
	defer e.WaitEmptyBacklog()

	err := e.DeployWasmContract(nil, contractName, fileName)
	require.NoError(t, err)
	_, _, rec := e.GetInfo()
	require.EqualValues(t, 4, len(rec))
}

func TestFail(t *testing.T) {
	e := alone.New(t, false, false)
	err := e.DeployWasmContract(nil, contractName, fileName,
		"failParam", 1,
	)
	require.Error(t, err)
	_, _, rec := e.GetInfo()
	require.EqualValues(t, 3, len(rec))
}

func TestFailRepeat(t *testing.T) {
	e := alone.New(t, false, false)
	err := e.DeployWasmContract(nil, contractName, fileName,
		"failParam", 1,
	)
	require.Error(t, err)

	err = e.DeployWasmContract(nil, contractName, fileName)
	require.NoError(t, err)
	_, _, rec := e.GetInfo()
	require.EqualValues(t, 4, len(rec))
}
