package test

import (
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/address"
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/balance"
	"github.com/iotaledger/wasp/packages/coretypes"
	"github.com/iotaledger/wasp/packages/kv/codec"
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasp/packages/vm/core/testcore/sbtests/sbtestsc"
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

	contractID := coretypes.NewContractID(chain.ChainID, coretypes.Hn(sbtestsc.Interface.Name))
	agentID := coretypes.NewAgentIDFromContractID(contractID)

	req := solo.NewCallParams(sbtestsc.Interface.Name, sbtestsc.FuncCheckContextFromFullEP,
		sbtestsc.ParamChainID, chain.ChainID,
		sbtestsc.ParamAgentID, agentID,
		sbtestsc.ParamCaller, userAgentID,
		sbtestsc.ParamChainOwnerID, chain.OriginatorAgentID,
		sbtestsc.ParamContractID, contractID,
		sbtestsc.ParamContractCreator, userAgentID,
	)
	_, err := chain.PostRequestSync(req, user)
	require.NoError(t, err)
}

func TestMainCallsFromViewEP(t *testing.T) { run2(t, testMainCallsFromViewEP) }
func testMainCallsFromViewEP(t *testing.T, w bool) {
	_, chain := setupChain(t, nil)

	user := setupDeployer(t, chain)
	userAddress := user.Address()
	userAgentID := coretypes.NewAgentIDFromAddress(userAddress)

	setupTestSandboxSC(t, chain, user, w)

	contractID := coretypes.NewContractID(chain.ChainID, coretypes.Hn(sbtestsc.Interface.Name))
	agentID := coretypes.NewAgentIDFromContractID(contractID)

	_, err := chain.CallView(sbtestsc.Interface.Name, sbtestsc.FuncCheckContextFromViewEP,
		sbtestsc.ParamChainID, chain.ChainID,
		sbtestsc.ParamAgentID, agentID,
		sbtestsc.ParamChainOwnerID, chain.OriginatorAgentID,
		sbtestsc.ParamContractID, contractID,
		sbtestsc.ParamContractCreator, userAgentID,
	)
	require.NoError(t, err)
}

func TestMintedSupplyOk(t *testing.T) { run2(t, testMintedSupplyOk) }
func testMintedSupplyOk(t *testing.T, w bool) {
	_, chain := setupChain(t, nil)

	user := setupDeployer(t, chain)
	setupTestSandboxSC(t, chain, user, w)

	supply := int64(42)
	req := solo.NewCallParams(sbtestsc.Interface.Name, sbtestsc.FuncGetMintedSupply).WithMinting(
		map[address.Address]int64{
			user.Address(): supply,
		},
	)
	tx, ret, err := chain.PostRequestSyncTx(req, user)
	require.NoError(t, err)

	extraIota := int64(0)
	if w {
		extraIota = 1
	}
	chain.Env.AssertAddressBalance(user.Address(), balance.ColorIOTA, solo.Saldo-3-extraIota-supply)
	chain.Env.AssertAddressBalance(user.Address(), balance.Color(tx.ID()), supply)

	supplyBack, ok, err := codec.DecodeInt64(ret.MustGet(sbtestsc.VarMintedSupply))
	require.NoError(t, err)
	require.True(t, ok)
	require.EqualValues(t, supply, supplyBack)
}
