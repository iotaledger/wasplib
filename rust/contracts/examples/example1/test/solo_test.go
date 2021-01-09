package test

import (
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/balance"
	"github.com/iotaledger/wasp/packages/coretypes"
	"github.com/iotaledger/wasp/packages/kv/codec"
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSolo1(t *testing.T) {
	env := solo.New(t, false, false)
	chain := env.NewChain(nil, "ex1")

	chainInfo, coreContracts := chain.GetInfo()   // calls view root::GetInfo
	require.EqualValues(t, 4, len(coreContracts)) // 4 core contracts deployed by default

	t.Logf("chainID: %s", chainInfo.ChainID)
	t.Logf("chain owner ID: %s", chainInfo.ChainOwnerID)
	for hname, rec := range coreContracts {
		t.Logf("    Core contract '%s': %s", rec.Name, coretypes.NewContractID(chain.ChainID, hname))
	}
}

func TestSolo2(t *testing.T) {
	env := solo.New(t, false, false)
	userWallet := env.NewSignatureSchemeWithFunds() // create new wallet with 1337 iotas
	userAddress := userWallet.Address()
	t.Logf("Address of the userWallet is: %s", userAddress)
	numIotas := env.GetAddressBalance(userAddress, balance.ColorIOTA) // how many iotas contains the address
	t.Logf("balance of the userWallet is: %d iota", numIotas)
	env.AssertAddressBalance(userAddress, balance.ColorIOTA, 1337) // assert the address has 1337 iotas
}

func TestSolo3(t *testing.T) {
	env := solo.New(t, false, false)
	chain := env.NewChain(nil, "ex1")
	// deploy the contract on chain
	err := chain.DeployWasmContract(nil, "example1", "../pkg/example1_bg.wasm")
	require.NoError(t, err)

	// call contract to store string
	theString := "Hello, world!"
	req := solo.NewCall("example1", "storeString",
		"paramString", theString)
	_, err = chain.PostRequest(req, nil)
	require.NoError(t, err)

	// call the contract to extract value of the 'paramString' and check
	res, err := chain.CallView("example1", "getString")
	require.NoError(t, err)
	returnedString, exists, err := codec.DecodeString(res.MustGet("paramString"))
	require.NoError(t, err)
	require.True(t, exists)
	require.EqualValues(t, theString, returnedString)
}

func TestSolo4(t *testing.T) {
	env := solo.New(t, false, false)
	chain := env.NewChain(nil, "ex1")
	// deploy the contract on chain
	err := chain.DeployWasmContract(nil, "example1", "../pkg/example1_bg.wasm")
	require.NoError(t, err)

	// call contract incorrectly
	req := solo.NewCall("example1", "storeString")
	_, err = chain.PostRequest(req, nil)
	require.Error(t, err)

}
