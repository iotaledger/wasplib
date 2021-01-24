package sandbox_tests

import (
	"github.com/iotaledger/wasp/packages/coretypes"
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasp/packages/vm/core/testcore/sandbox_tests/test_sandbox_sc"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestChainOwnerIDView(t *testing.T) {
	_, chain := setupChain(t, nil)
	setupTestSandboxSC(t, chain, nil)

	ret, err := chain.CallView(SandboxSCName, test_sandbox_sc.FuncChainOwnerIDView)
	require.NoError(t, err)

	c := ret.MustGet(test_sandbox_sc.VarChainOwner)

	require.EqualValues(t, chain.OriginatorAgentID.Bytes(), c)
}

func TestChainOwnerIDFull(t *testing.T) {
	_, chain := setupChain(t, nil)
	setupTestSandboxSC(t, chain, nil)

	req := solo.NewCall(SandboxSCName, test_sandbox_sc.FuncChainOwnerIDFull)
	ret, err := chain.PostRequest(req, nil)
	require.NoError(t, err)

	c := ret.MustGet(test_sandbox_sc.VarChainOwner)
	require.EqualValues(t, chain.OriginatorAgentID.Bytes(), c)
}

func TestContractIDView(t *testing.T) {
	_, chain := setupChain(t, nil)
	setupTestSandboxSC(t, chain, nil)

	ret, err := chain.CallView(SandboxSCName, test_sandbox_sc.FuncContractIDView)
	require.NoError(t, err)
	cID := coretypes.NewContractID(chain.ChainID, coretypes.Hn(SandboxSCName))
	require.EqualValues(t, cID[:], ret.MustGet(test_sandbox_sc.VarContractID))
}

func TestContractIDFull(t *testing.T) {
	_, chain := setupChain(t, nil)
	setupTestSandboxSC(t, chain, nil)

	req := solo.NewCall(SandboxSCName, test_sandbox_sc.FuncContractIDFull)
	ret, err := chain.PostRequest(req, nil)
	require.NoError(t, err)
	cID := coretypes.NewContractID(chain.ChainID, coretypes.Hn(SandboxSCName))
	require.EqualValues(t, cID[:], ret.MustGet(test_sandbox_sc.VarContractID))
}

func TestSandboxCall(t *testing.T) {
	_, chain := setupChain(t, nil)
	setupTestSandboxSC(t, chain, nil)

	ret, err := chain.CallView(SandboxSCName, test_sandbox_sc.FuncSandboxCall)
	require.NoError(t, err)

	d := ret.MustGet(test_sandbox_sc.VarSandboxCall)
	require.EqualValues(t, "'solo' testing chain", string(d))
}
