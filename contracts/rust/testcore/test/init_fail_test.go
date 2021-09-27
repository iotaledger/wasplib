package test

import (
	"testing"

	"github.com/iotaledger/wasp/packages/vm/core"
	"github.com/iotaledger/wasp/packages/vm/core/testcore/sbtests/sbtestsc"
	"github.com/stretchr/testify/require"
)

func TestInitSuccess(t *testing.T) { run2(t, testInitSuccess) }
func testInitSuccess(t *testing.T, w bool) {
	ctx := setupTest(t, w)
	require.NoError(t, ctx.Err)
}

//func TestInitFail(t *testing.T) { run2(t, testInitFail) }
//func testInitFail(t *testing.T, w bool) {
//	init := testcore.ScFuncs.Init(nil)
//	init.Params.Fail().SetValue(1)
//	ctx := setupTestInit(t, w, init.Func)
//	require.Error(t, ctx.Err)
//}

func TestInitFail(t *testing.T) {
	_, chain := setupChain(t, nil)
	err := chain.DeployContract(nil, ScName, sbtestsc.Contract.ProgramHash,
		sbtestsc.ParamFail, 1)
	require.Error(t, err)
}

func TestInitFailRepeat(t *testing.T) {
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
