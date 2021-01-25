package sandbox_tests

import (
	"github.com/iotaledger/wasp/packages/coretypes"
	"github.com/iotaledger/wasp/packages/hashing"
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasp/packages/vm/core/testcore/sandbox_tests/test_sandbox_sc"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTypesFull(t *testing.T) { run2(t, testTypesFull) }
func testTypesFull(t *testing.T, w bool) {
	_, chain := setupChain(t, nil)
	cID, _ := setupTestSandboxSC(t, chain, nil, w)

	req := solo.NewCall(SandboxSCName, test_sandbox_sc.FuncPassTypesFull,
		"string", "string",
		"string-0", "",
		"int64", 42,
		"int64-0", 0,
		"Hash", hashing.HashStrings("Hash"),
		"Hname", coretypes.Hn("Hname"),
		"Hname-0", coretypes.Hname(0),
		"ContractID", cID,
		"ChainID", chain.ChainID,
		"Address", chain.ChainAddress,
		"AgentID", chain.OriginatorAgentID,
	)
	_, err := chain.PostRequest(req, nil)
	require.NoError(t, err)
}

func TestTypesView(t *testing.T) { run2(t, testTypesView) }
func testTypesView(t *testing.T, w bool) {
	_, chain := setupChain(t, nil)
	cID, _ := setupTestSandboxSC(t, chain, nil, w)

	_, err := chain.CallView(SandboxSCName, test_sandbox_sc.FuncPassTypesView,
		"string", "string",
		"string-0", "",
		"int64", 42,
		"int64-0", 0,
		"Hash", hashing.HashStrings("Hash"),
		"Hname", coretypes.Hn("Hname"),
		"Hname-0", coretypes.Hname(0),
		"ContractID", cID,
		"ChainID", chain.ChainID,
		"Address", chain.ChainAddress,
		"AgentID", chain.OriginatorAgentID,
	)
	require.NoError(t, err)
}
