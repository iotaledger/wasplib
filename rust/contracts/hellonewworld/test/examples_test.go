package test

import (
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/balance"
	"github.com/iotaledger/wasp/packages/kv/codec"
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEx1(t *testing.T) {
	glb := solo.New(t, false, false)
	chain := glb.NewChain(nil, "ch1")

	chainInfo, coreContracts := chain.GetInfo() // calls view root::GetInfo

	require.EqualValues(t, 4, len(coreContracts)) // 4 core contracts deployed by default

	t.Logf("chainID: %s", chainInfo.ChainID)
	t.Logf("chain owner ID: %s", chainInfo.ChainOwnerID)
	for hname, rec := range coreContracts {
		t.Logf("    Core contract #%d: %s", hname, rec.Name)
	}
}

func TestEx2(t *testing.T) {
	glb := solo.New(t, false, false)
	userWallet := glb.NewSignatureSchemeWithFunds()
	userAddress := userWallet.Address()
	t.Logf("Address of the userWallet is: %s", userAddress)
	numIotas := glb.GetUtxodbBalance(userAddress, balance.ColorIOTA)
	t.Logf("balance of the userWallet is: %d iotas", numIotas)
	glb.AssertUtxodbBalance(userAddress, balance.ColorIOTA, 1337)
}

func TestEx3(t *testing.T) {
	glb := solo.New(t, false, false)
	chain := glb.NewChain(nil, "ch1")
	err := chain.DeployWasmContract(nil, "hnw1", "../pkg/hellonewworld_bg.wasm")
	require.NoError(t, err)
	_, coreContracts := chain.GetInfo()
	require.EqualValues(t, 5, len(coreContracts)) // 4 core contracts deployed by default
}

func TestEx4(t *testing.T) {
	glb := solo.New(t, false, false)
	chain := glb.NewChain(nil, "ch1")
	err := chain.DeployWasmContract(nil, "hnw1", "../pkg/hellonewworld_bg.wasm")
	require.NoError(t, err)

	req := solo.NewCall("hnw1", "hello")
	for i := 0; i < 3; i++ {
		_, err = chain.PostRequest(req, nil)
		require.NoError(t, err)
	}

	res, err := chain.CallView("hnw1", "getCounter")
	require.NoError(t, err)
	counter, exists, err := codec.DecodeInt64(res.MustGet("counter"))
	require.NoError(t, err)
	require.True(t, exists)
	require.EqualValues(t, 3, counter)
}
