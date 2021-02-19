package test

import (
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasp/packages/vm/core/testcore/sbtests/sbtestsc"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSuccess(t *testing.T) {
	env := solo.New(t, false, false)
	chain := env.NewChain(nil, "chain1")
	err := chain.DeployContract(nil, sbtestsc.Name, sbtestsc.Interface.ProgramHash)
	require.NoError(t, err)
}

func TestFail(t *testing.T) {
	env := solo.New(t, false, false)
	chain := env.NewChain(nil, "chain1")
	err := chain.DeployContract(nil, sbtestsc.Name, sbtestsc.Interface.ProgramHash,
		sbtestsc.ParamFail, 1)
	require.Error(t, err)
}

func TestFailRepeat(t *testing.T) {
	env := solo.New(t, false, false)
	chain := env.NewChain(nil, "chain1")
	err := chain.DeployContract(nil, sbtestsc.Name, sbtestsc.Interface.ProgramHash,
		sbtestsc.ParamFail, 1)
	require.Error(t, err)
	_, rec := chain.GetInfo()
	require.EqualValues(t, 4, len(rec))

	// repeat must succeed
	err = chain.DeployContract(nil, sbtestsc.Name, sbtestsc.Interface.ProgramHash)
	require.NoError(t, err)
	_, rec = chain.GetInfo()
	require.EqualValues(t, 5, len(rec))
}
