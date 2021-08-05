package test

import (
	"testing"

	"github.com/iotaledger/wasp/packages/vm/core"
	"github.com/iotaledger/wasp/packages/vm/core/testcore/sbtests/sbtestsc"
	"github.com/stretchr/testify/require"
)

func TestSuccess(t *testing.T) {
	_, chain := setupChain(t, nil)
	err := chain.DeployContract(nil, ScName, sbtestsc.Contract.ProgramHash)
	require.NoError(t, err)
}

func TestFail(t *testing.T) {
	_, chain := setupChain(t, nil)
	err := chain.DeployContract(nil, ScName, sbtestsc.Contract.ProgramHash,
		sbtestsc.ParamFail, 1)
	require.Error(t, err)
}

func TestFailRepeat(t *testing.T) {
	_, chain := setupChain(t, nil)
	err := chain.DeployContract(nil, ScName, sbtestsc.Contract.ProgramHash,
		sbtestsc.ParamFail, 1)
	require.Error(t, err)
	_, _, rec := chain.GetInfo()
	require.EqualValues(t, len(core.AllCoreContractsByHash), len(rec))

	// repeat must succeed
	err = chain.DeployContract(nil, ScName, sbtestsc.Contract.ProgramHash)
	require.NoError(t, err)
	_, _, rec = chain.GetInfo()
	require.EqualValues(t, len(core.AllCoreContractsByHash)+1, len(rec))
}
