package test

import (
	"github.com/iotaledger/goshimmer/packages/ledgerstate"
	"github.com/iotaledger/goshimmer/packages/ledgerstate/utxoutil"
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

	user, _, userAgentID := setupDeployer(t, chain)

	setupTestSandboxSC(t, chain, user, w)

	req := solo.NewCallParams(ScName, sbtestsc.FuncCheckContextFromFullEP,
		sbtestsc.ParamChainID, chain.ChainID,
		sbtestsc.ParamAgentID, coretypes.NewAgentID(chain.ChainID.AsAddress(), HScName),
		sbtestsc.ParamCaller, userAgentID,
		sbtestsc.ParamChainOwnerID, chain.OriginatorAgentID,
		sbtestsc.ParamContractCreator, userAgentID,
	).WithIotas(1)
	_, err := chain.PostRequestSync(req, user)
	require.NoError(t, err)
}

func TestMainCallsFromViewEP(t *testing.T) { run2(t, testMainCallsFromViewEP) }
func testMainCallsFromViewEP(t *testing.T, w bool) {
	_, chain := setupChain(t, nil)

	user, _, userAgentID := setupDeployer(t, chain)

	setupTestSandboxSC(t, chain, user, w)

	_, err := chain.CallView(ScName, sbtestsc.FuncCheckContextFromViewEP,
		sbtestsc.ParamChainID, chain.ChainID,
		sbtestsc.ParamAgentID, coretypes.NewAgentID(chain.ChainID.AsAddress(), HScName),
		sbtestsc.ParamChainOwnerID, chain.OriginatorAgentID,
		sbtestsc.ParamContractCreator, userAgentID,
	)
	require.NoError(t, err)
}

func TestMintedSupplyOk(t *testing.T) { run2(t, testMintedSupplyOk) }
func testMintedSupplyOk(t *testing.T, w bool) {
	_, chain := setupChain(t, nil)

	user, userAddress, _ := setupDeployer(t, chain)
	setupTestSandboxSC(t, chain, user, w)

	newSupply := uint64(42)
	req := solo.NewCallParams(ScName, sbtestsc.FuncGetMintedSupply).
		WithIotas(1).
		WithMint(userAddress, newSupply)
	tx, ret, err := chain.PostRequestSyncTx(req, user)
	require.NoError(t, err)

	mintedAmounts := utxoutil.GetMintedAmounts(tx)
	t.Logf("minting request tx: %s", tx.ID().Base58())

	require.Len(t, mintedAmounts, 1)
	var color ledgerstate.Color
	for col := range mintedAmounts {
		color = col
		break
	}
	t.Logf("Minted: amount = %d color = %s", newSupply, color.Base58())

	extraIota := uint64(0)
	if w {
		extraIota = 1
	}
	chain.Env.AssertAddressIotas(userAddress, solo.Saldo-3-extraIota-newSupply)
	chain.Env.AssertAddressBalance(userAddress, color, newSupply)

	colorBack, ok, err := codec.DecodeColor(ret.MustGet(sbtestsc.VarMintedColor))
	require.NoError(t, err)
	require.True(t, ok)
	t.Logf("color back: %s", colorBack.Base58())
	require.EqualValues(t, color, colorBack)
	supplyBack, ok, err := codec.DecodeUint64(ret.MustGet(sbtestsc.VarMintedSupply))
	require.NoError(t, err)
	require.True(t, ok)
	t.Logf("supply back: %d", supplyBack)
	require.EqualValues(t, int(newSupply), int(supplyBack))
}
