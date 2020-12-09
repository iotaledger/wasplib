package erc20

import (
	"github.com/iotaledger/wasp/packages/coretypes"
	"github.com/iotaledger/wasp/packages/vm/alone"
	"github.com/stretchr/testify/require"
	"testing"
)

const (
	erc20name = "erc20test"
	erc20file = "../pkg/erc20_bg.wasm"
)

func TestDeployErc20(t *testing.T) {
	e := alone.New(t, false, false)
	defer e.WaitEmptyBacklog()

	creator := e.NewSigScheme()
	creatorAgentID := coretypes.NewAgentIDFromAddress(creator.Address())
	err := e.DeployWasmContract(nil, erc20name, erc20file,
		PARAM_SUPPLY, 1000000,
		PARAM_CREATOR, creatorAgentID,
	)
	require.NoError(t, err)
	_, _, rec := e.GetInfo()
	require.EqualValues(t, 4, len(rec))

	_, err = e.FindContract(erc20name)
	require.NoError(t, err)

	// deploy second time
	err = e.DeployWasmContract(nil, erc20name, erc20file,
		PARAM_SUPPLY, 1000000,
		PARAM_CREATOR, creatorAgentID,
	)
	require.Error(t, err)
	_, _, rec = e.GetInfo()
	require.EqualValues(t, 4, len(rec))
}

func TestDeployErc20Fail1(t *testing.T) {
	e := alone.New(t, false, false)
	err := e.DeployWasmContract(nil, erc20name, erc20file)
	require.Error(t, err)
	_, _, rec := e.GetInfo()
	require.EqualValues(t, 3, len(rec))
}

func TestDeployErc20Fail2(t *testing.T) {
	e := alone.New(t, false, false)
	err := e.DeployWasmContract(nil, erc20name, erc20file,
		PARAM_SUPPLY, 1000000,
	)
	require.Error(t, err)
	_, _, rec := e.GetInfo()
	require.EqualValues(t, 3, len(rec))
}

func TestDeployErc20Fail3(t *testing.T) {
	e := alone.New(t, false, false)
	creator := e.NewSigScheme()
	creatorAgentID := coretypes.NewAgentIDFromAddress(creator.Address())
	err := e.DeployWasmContract(nil, erc20name, erc20file,
		PARAM_CREATOR, creatorAgentID,
	)
	require.Error(t, err)
	_, _, rec := e.GetInfo()
	require.EqualValues(t, 3, len(rec))
}

func TestDeployErc20Fail3Repeat(t *testing.T) {
	e := alone.New(t, false, false)
	creator := e.NewSigScheme()
	creatorAgentID := coretypes.NewAgentIDFromAddress(creator.Address())
	err := e.DeployWasmContract(nil, erc20name, erc20file,
		PARAM_CREATOR, creatorAgentID,
	)
	require.Error(t, err)
	_, _, rec := e.GetInfo()
	require.EqualValues(t, 3, len(rec))

	// repeat after failure
	err = e.DeployWasmContract(nil, erc20name, erc20file,
		PARAM_SUPPLY, 1000000,
		PARAM_CREATOR, creatorAgentID,
	)
	require.NoError(t, err)
	_, _, rec = e.GetInfo()
	require.EqualValues(t, 4, len(rec))

	_, err = e.FindContract(erc20name)
	require.NoError(t, err)
}
