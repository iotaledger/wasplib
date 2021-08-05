package test

import (
	"testing"

	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasp/packages/vm/core/testcore/sbtests/sbtestsc"
	"github.com/stretchr/testify/require"
)

func TestChainOwnerIDView(t *testing.T) { run2(t, testChainOwnerIDView) }
func testChainOwnerIDView(t *testing.T, w bool) {
	_, chain := setupChain(t, nil)
	setupTestSandboxSC(t, chain, nil, w)

	ret, err := chain.CallView(ScName, sbtestsc.FuncChainOwnerIDView.Name)
	require.NoError(t, err)

	c := ret.MustGet(sbtestsc.ParamChainOwnerID)

	require.EqualValues(t, chain.OriginatorAgentID.Bytes(), c)
}

func TestChainOwnerIDFull(t *testing.T) { run2(t, testChainOwnerIDFull) }
func testChainOwnerIDFull(t *testing.T, w bool) {
	_, chain := setupChain(t, nil)
	setupTestSandboxSC(t, chain, nil, w)

	req := solo.NewCallParams(ScName, sbtestsc.FuncChainOwnerIDFull.Name).WithIotas(1)
	ret, err := chain.PostRequestSync(req, nil)
	require.NoError(t, err)

	c := ret.MustGet(sbtestsc.ParamChainOwnerID)
	require.EqualValues(t, chain.OriginatorAgentID.Bytes(), c)
}

func TestSandboxCall(t *testing.T) { run2(t, testSandboxCall) }
func testSandboxCall(t *testing.T, w bool) {
	_, chain := setupChain(t, nil)
	setupTestSandboxSC(t, chain, nil, w)

	ret, err := chain.CallView(ScName, sbtestsc.FuncSandboxCall.Name)
	require.NoError(t, err)

	d := ret.MustGet(sbtestsc.VarSandboxCall)
	require.EqualValues(t, "'solo' testing chain", string(d))
}
