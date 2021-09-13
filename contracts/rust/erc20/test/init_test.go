package test

import (
	"testing"

	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasp/packages/vm/core"
	"github.com/iotaledger/wasplib/contracts/common"
	"github.com/iotaledger/wasplib/contracts/rust/erc20"
	"github.com/stretchr/testify/require"
)

func TestDeployErc20(t *testing.T) {
	setupTest(t)
	ctx := common.NewSoloContext(t, chain, erc20.ScName, erc20.OnLoad,
		ParamSupply, solo.Saldo,
		ParamCreator, creator.ScAgentID().Bytes(),
	)
	require.NoError(t, ctx.Err)
	_, _, rec := chain.GetInfo()
	require.EqualValues(t, len(core.AllCoreContractsByHash)+1, len(rec))

	_, err := chain.FindContract(erc20.ScName)
	require.NoError(t, err)

	// deploy second time
	ctx = common.NewSoloContext(t, chain, erc20.ScName, erc20.OnLoad,
		ParamSupply, solo.Saldo,
		ParamCreator, creator.ScAgentID().Bytes(),
	)
	require.Error(t, ctx.Err)
	_, _, rec = chain.GetInfo()
	require.EqualValues(t, len(core.AllCoreContractsByHash)+1, len(rec))
}

func TestDeployErc20Fail1(t *testing.T) {
	setupTest(t)
	ctx := common.NewSoloContext(t, chain, erc20.ScName, erc20.OnLoad)
	require.Error(t, ctx.Err)
	_, _, rec := chain.GetInfo()
	require.EqualValues(t, len(core.AllCoreContractsByHash), len(rec))
}

func TestDeployErc20Fail2(t *testing.T) {
	setupTest(t)
	ctx := common.NewSoloContext(t, chain, erc20.ScName, erc20.OnLoad,
		ParamSupply, solo.Saldo,
	)
	require.Error(t, ctx.Err)
	_, _, rec := chain.GetInfo()
	require.EqualValues(t, len(core.AllCoreContractsByHash), len(rec))
}

func TestDeployErc20Fail3(t *testing.T) {
	setupTest(t)
	ctx := common.NewSoloContext(t, chain, erc20.ScName, erc20.OnLoad,
		ParamCreator, creator.ScAgentID().Bytes(),
	)
	require.Error(t, ctx.Err)
	_, _, rec := chain.GetInfo()
	require.EqualValues(t, len(core.AllCoreContractsByHash), len(rec))
}

func TestDeployErc20Fail3Repeat(t *testing.T) {
	setupTest(t)
	ctx := common.NewSoloContext(t, chain, erc20.ScName, erc20.OnLoad,
		ParamCreator, creator.ScAgentID().Bytes(),
	)
	require.Error(t, ctx.Err)
	_, _, rec := chain.GetInfo()
	require.EqualValues(t, len(core.AllCoreContractsByHash), len(rec))

	_, err := chain.FindContract(erc20.ScName)
	require.Error(t, err)

	// repeat after failure
	ctx = common.NewSoloContext(t, chain, erc20.ScName, erc20.OnLoad,
		ParamSupply, solo.Saldo,
		ParamCreator, creator.ScAgentID().Bytes(),
	)
	require.NoError(t, ctx.Err)
	_, _, rec = chain.GetInfo()
	require.EqualValues(t, len(core.AllCoreContractsByHash)+1, len(rec))

	_, err = chain.FindContract(erc20.ScName)
	require.NoError(t, err)
}
