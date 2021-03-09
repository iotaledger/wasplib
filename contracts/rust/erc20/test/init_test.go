package test

import (
	"github.com/iotaledger/wasp/packages/coretypes"
	"github.com/iotaledger/wasplib/contracts/common"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDeployErc20(t *testing.T) {
	chain := common.StartChain(t, ScName)
	creatorAgentID = coretypes.NewAgentIDFromAddress(common.CreatorWallet.Address())
	err := common.DeployWasmContractByName(chain, ScName,
		ParamSupply, 1000000,
		ParamCreator, creatorAgentID,
	)
	require.NoError(t, err)
	_, rec := chain.GetInfo()
	require.EqualValues(t, 5, len(rec))

	_, err = chain.FindContract(ScName)
	require.NoError(t, err)

	// deploy second time
	err = common.DeployWasmContractByName(chain, ScName,
		ParamSupply, 1000000,
		ParamCreator, creatorAgentID,
	)
	require.Error(t, err)
	_, rec = chain.GetInfo()
	require.EqualValues(t, 5, len(rec))
}

func TestDeployErc20Fail1(t *testing.T) {
	chain := common.StartChain(t, ScName)
	err := common.DeployWasmContractByName(chain, ScName)
	require.Error(t, err)
	_, rec := chain.GetInfo()
	require.EqualValues(t, 4, len(rec))
}

func TestDeployErc20Fail2(t *testing.T) {
	chain := common.StartChain(t, ScName)
	err := common.DeployWasmContractByName(chain, ScName,
		ParamSupply, 1000000,
	)
	require.Error(t, err)
	_, rec := chain.GetInfo()
	require.EqualValues(t, 4, len(rec))
}

func TestDeployErc20Fail3(t *testing.T) {
	chain := common.StartChain(t, ScName)
	creatorAgentID = coretypes.NewAgentIDFromAddress(common.CreatorWallet.Address())
	err := common.DeployWasmContractByName(chain, ScName,
		ParamCreator, creatorAgentID,
	)
	require.Error(t, err)
	_, rec := chain.GetInfo()
	require.EqualValues(t, 4, len(rec))
}

func TestDeployErc20Fail3Repeat(t *testing.T) {
	chain := common.StartChain(t, ScName)
	creatorAgentID = coretypes.NewAgentIDFromAddress(common.CreatorWallet.Address())
	err := common.DeployWasmContractByName(chain, ScName,
		ParamCreator, creatorAgentID,
	)
	require.Error(t, err)
	_, rec := chain.GetInfo()
	require.EqualValues(t, 4, len(rec))

	// repeat after failure
	err = common.DeployWasmContractByName(chain, ScName,
		ParamSupply, 1000000,
		ParamCreator, creatorAgentID,
	)
	require.NoError(t, err)
	_, rec = chain.GetInfo()
	require.EqualValues(t, 5, len(rec))

	_, err = chain.FindContract(ScName)
	require.NoError(t, err)
}
