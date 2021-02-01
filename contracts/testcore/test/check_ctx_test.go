package sandbox_tests

import (
	"github.com/iotaledger/wasp/packages/coretypes"
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasp/packages/vm/core/testcore/sandbox_tests/test_sandbox_sc"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMainCallsFromFullEP(t *testing.T) { run2(t, testMainCallsFromFullEP) }
func testMainCallsFromFullEP(t *testing.T, w bool) {
	_, chain := setupChain(t, nil)

	user := setupDeployer(t, chain)
	userAddress := user.Address()
	userAgentID := coretypes.NewAgentIDFromAddress(userAddress)

	setupTestSandboxSC(t, chain, user, w)

	contractID := coretypes.NewContractID(chain.ChainID, coretypes.Hn(test_sandbox_sc.Interface.Name))
	agentID := coretypes.NewAgentIDFromContractID(contractID)

	req := solo.NewCallParams(test_sandbox_sc.Interface.Name, test_sandbox_sc.FuncCheckContextFromFullEP,
		test_sandbox_sc.ParamChainID, chain.ChainID,
		test_sandbox_sc.ParamAgentID, agentID,
		test_sandbox_sc.ParamCaller, userAgentID,
		test_sandbox_sc.ParamChainOwnerID, chain.OriginatorAgentID,
		test_sandbox_sc.ParamContractID, contractID,
		test_sandbox_sc.ParamContractCreator, userAgentID,
	)
	_, err := chain.PostRequest(req, user)
	require.NoError(t, err)
}

func TestMainCallsFromViewEP(t *testing.T) { run2(t, testMainCallsFromViewEP) }
func testMainCallsFromViewEP(t *testing.T, w bool) {
	_, chain := setupChain(t, nil)

	user := setupDeployer(t, chain)
	userAddress := user.Address()
	userAgentID := coretypes.NewAgentIDFromAddress(userAddress)

	setupTestSandboxSC(t, chain, user, w)

	contractID := coretypes.NewContractID(chain.ChainID, coretypes.Hn(test_sandbox_sc.Interface.Name))
	agentID := coretypes.NewAgentIDFromContractID(contractID)

	_, err := chain.CallView(test_sandbox_sc.Interface.Name, test_sandbox_sc.FuncCheckContextFromViewEP,
		test_sandbox_sc.ParamChainID, chain.ChainID,
		test_sandbox_sc.ParamAgentID, agentID,
		test_sandbox_sc.ParamChainOwnerID, chain.OriginatorAgentID,
		test_sandbox_sc.ParamContractID, contractID,
		test_sandbox_sc.ParamContractCreator, userAgentID,
	)
	require.NoError(t, err)
}
