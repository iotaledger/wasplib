package test

import (
	"testing"

	"github.com/iotaledger/wasp/packages/kv/kvdecoder"
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasp/packages/vm/core"
	"github.com/iotaledger/wasp/packages/vm/core/testcore/sbtests/sbtestsc"
	"github.com/stretchr/testify/require"
)

func TestSpawn(t *testing.T) {
	_, chain := setupChain(t, nil)
	_, _ = setupTestSandboxSC(t, chain, nil, false)

	req := solo.NewCallParams(ScName, sbtestsc.FuncSpawn.Name).WithIotas(1)
	_, err := chain.PostRequestSync(req, nil)
	require.NoError(t, err)

	ret, err := chain.CallView(ScName+"_spawned", sbtestsc.FuncGetCounter.Name)
	require.NoError(t, err)
	res := kvdecoder.New(ret, chain.Log)
	counter := res.MustGetUint64(sbtestsc.VarCounter)
	require.EqualValues(t, 5, counter)

	_, _, recs := chain.GetInfo()
	require.EqualValues(t, len(core.AllCoreContractsByHash)+2, len(recs))
}
