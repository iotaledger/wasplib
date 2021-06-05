package test

import (
	"testing"

	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasp/packages/vm/core/testcore/sbtests/sbtestsc"
	"github.com/stretchr/testify/require"
)

func TestBasicBlockContext1(t *testing.T) {
	_, chain := setupChain(t, nil)
	_, _ = setupTestSandboxSC(t, chain, nil, false)

	req := solo.NewCallParams(SandboxSCName, sbtestsc.FuncTestBlockContext1).WithIotas(1)
	_, err := chain.PostRequestSync(req, nil)
	require.NoError(t, err)
}

func TestBasicBlockContext2(t *testing.T) {
	_, chain := setupChain(t, nil)
	_, _ = setupTestSandboxSC(t, chain, nil, false)

	req := solo.NewCallParams(SandboxSCName, sbtestsc.FuncTestBlockContext2).WithIotas(1)
	_, err := chain.PostRequestSync(req, nil)
	require.NoError(t, err)

	res, err := chain.CallView(sbtestsc.Interface.Name, sbtestsc.FuncGetStringValue, sbtestsc.ParamVarName, "atTheEndKey")
	require.NoError(t, err)
	b, err := res.Get("atTheEndKey")
	require.NoError(t, err)
	require.EqualValues(t, "atTheEndValue", string(b))
}
