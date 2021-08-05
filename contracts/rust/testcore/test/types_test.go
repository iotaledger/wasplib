package test

import (
	"testing"

	"github.com/iotaledger/wasp/packages/hashing"
	"github.com/iotaledger/wasp/packages/iscp"
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasp/packages/vm/core/testcore/sbtests/sbtestsc"
	"github.com/stretchr/testify/require"
)

func TestTypesFull(t *testing.T) { run2(t, testTypesFull) }
func testTypesFull(t *testing.T, w bool) {
	_, chain := setupChain(t, nil)
	cID, _ := setupTestSandboxSC(t, chain, nil, w)

	req := solo.NewCallParams(ScName, sbtestsc.FuncPassTypesFull.Name,
		"string", "string",
		"string-0", "",
		"int64", 42,
		"int64-0", 0,
		"Hash", hashing.HashStrings("Hash"),
		"Hname", iscp.Hn("Hname"),
		"Hname-0", iscp.Hname(0),
		"ContractID", cID,
		"ChainID", chain.ChainID,
		"Address", chain.ChainID.AsAddress(),
		"AgentID", chain.OriginatorAgentID,
	).WithIotas(1)
	_, err := chain.PostRequestSync(req, nil)
	require.NoError(t, err)
}

func TestTypesView(t *testing.T) { run2(t, testTypesView) }
func testTypesView(t *testing.T, w bool) {
	_, chain := setupChain(t, nil)
	cID, _ := setupTestSandboxSC(t, chain, nil, w)

	_, err := chain.CallView(ScName, sbtestsc.FuncPassTypesView.Name,
		"string", "string",
		"string-0", "",
		"int64", 42,
		"int64-0", 0,
		"Hash", hashing.HashStrings("Hash"),
		"Hname", iscp.Hn("Hname"),
		"Hname-0", iscp.Hname(0),
		"ContractID", cID,
		"ChainID", chain.ChainID,
		"Address", chain.ChainID.AsAddress(),
		"AgentID", chain.OriginatorAgentID,
	)
	require.NoError(t, err)
}
