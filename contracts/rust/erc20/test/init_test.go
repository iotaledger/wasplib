package test

import (
	"testing"

	"github.com/iotaledger/wasp/packages/coretypes"
	"github.com/iotaledger/wasp/packages/vm/core"
	"github.com/iotaledger/wasplib/contracts/common"
	"github.com/stretchr/testify/require"
)

func TestDeployErc20(t *testing.T) {
	chain := common.StartChain(t, ScName)
	creator, creatorAddr = chain.Env.NewKeyPairWithFunds()
	creatorAgentID = coretypes.NewAgentID(creatorAddr, 0)
	err := common.DeployWasmContractByName(chain, ScName,
		ParamSupply, 1000000,
		ParamCreator, creatorAgentID,
	)
	require.NoError(t, err)
	_, _, rec := chain.GetInfo()
	require.EqualValues(t, len(core.AllCoreContractsByHash)+1, len(rec))

	_, err = chain.FindContract(ScName)
	require.NoError(t, err)

	// deploy second time
	err = common.DeployWasmContractByName(chain, ScName,
		ParamSupply, 1000000,
		ParamCreator, creatorAgentID,
	)
	require.Error(t, err)
	_, _, rec = chain.GetInfo()
	require.EqualValues(t, len(core.AllCoreContractsByHash)+1, len(rec))
}

func TestDeployErc20Fail1(t *testing.T) {
	chain := common.StartChain(t, ScName)
	err := common.DeployWasmContractByName(chain, ScName)
	require.Error(t, err)
	_, _, rec := chain.GetInfo()
	require.EqualValues(t, len(core.AllCoreContractsByHash), len(rec))
}

func TestDeployErc20Fail2(t *testing.T) {
	chain := common.StartChain(t, ScName)
	err := common.DeployWasmContractByName(chain, ScName,
		ParamSupply, 1000000,
	)
	require.Error(t, err)
	_, _, rec := chain.GetInfo()
	require.EqualValues(t, len(core.AllCoreContractsByHash), len(rec))
}

func TestDeployErc20Fail3(t *testing.T) {
	chain := common.StartChain(t, ScName)
	creator, creatorAddr = chain.Env.NewKeyPairWithFunds()
	creatorAgentID = coretypes.NewAgentID(creatorAddr, 0)
	err := common.DeployWasmContractByName(chain, ScName,
		ParamCreator, creatorAgentID,
	)
	require.Error(t, err)
	_, _, rec := chain.GetInfo()
	require.EqualValues(t, len(core.AllCoreContractsByHash), len(rec))
}

func TestDeployErc20Fail3Repeat(t *testing.T) {
	chain := common.StartChain(t, ScName)
	creator, creatorAddr = chain.Env.NewKeyPairWithFunds()
	creatorAgentID = coretypes.NewAgentID(creatorAddr, 0)
	err := common.DeployWasmContractByName(chain, ScName,
		ParamCreator, creatorAgentID,
	)
	require.Error(t, err)
	_, _, rec := chain.GetInfo()
	require.EqualValues(t, len(core.AllCoreContractsByHash), len(rec))

	// repeat after failure
	err = common.DeployWasmContractByName(chain, ScName,
		ParamSupply, 1000000,
		ParamCreator, creatorAgentID,
	)
	require.NoError(t, err)
	_, _, rec = chain.GetInfo()
	require.EqualValues(t, len(core.AllCoreContractsByHash)+1, len(rec))

	_, err = chain.FindContract(ScName)
	require.NoError(t, err)
}
