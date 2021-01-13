package test

import (
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/balance"
	"github.com/iotaledger/wasp/packages/coretypes"
	"github.com/iotaledger/wasp/packages/kv/codec"
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasp/packages/vm/core/accounts"
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
	chain := env.NewChain(nil, "ex3")
	// deploy the contract on chain
	err := chain.DeployWasmContract(nil, "example1", "../pkg/example1_bg.wasm")
	require.NoError(t, err)

	// call contract to store string
	req := solo.NewCall("example1", "storeString", "paramString", "Hello, world!")
	_, err = chain.PostRequest(req, nil)
	require.NoError(t, err)

	// call the contract to extract value of the 'paramString' and check
	res, err := chain.CallView("example1", "getString")
	require.NoError(t, err)
	returnedString, exists, err := codec.DecodeString(res.MustGet("paramString"))
	require.NoError(t, err)
	require.True(t, exists)
	require.EqualValues(t, "Hello, world!", returnedString)
}

func TestSolo4(t *testing.T) {
	env := solo.New(t, false, false)
	chain := env.NewChain(nil, "ex4")
	// deploy the contract on chain
	err := chain.DeployWasmContract(nil, "example1", "../pkg/example1_bg.wasm")
	require.NoError(t, err)

	// call contract incorrectly
	req := solo.NewCall("example1", "storeString")
	_, err = chain.PostRequest(req, nil)
	require.Error(t, err)
}

func TestSolo5(t *testing.T) {
	env := solo.New(t, false, false)
	chain := env.NewChain(nil, "ex5")

	// create wallet with 1337 iotas.
	// wallet has address and it is globally identified through
	// universal identifier: the agent ID
	userWallet := env.NewSignatureSchemeWithFunds()
	userAddress := userWallet.Address()
	userAgentID := coretypes.NewAgentIDFromAddress(userAddress)

	env.AssertAddressBalance(userAddress, balance.ColorIOTA, 1337) // 1337 on address
	chain.AssertAccountBalance(userAgentID, balance.ColorIOTA, 0)  // empty on-chain

	t.Logf("Address of the userWallet is: %s", userAddress)
	numIotas := env.GetAddressBalance(userAddress, balance.ColorIOTA)
	t.Logf("balance of the userWallet is: %d iota", numIotas)
	env.AssertAddressBalance(userAddress, balance.ColorIOTA, 1337)

	// send 42 iotas from wallet to own account on-chain, controlled by the same wallet
	req := solo.NewCall(accounts.Name, accounts.FuncDeposit).
		WithTransfer(balance.ColorIOTA, 42)
	_, err := chain.PostRequest(req, userWallet)
	require.NoError(t, err)

	// check address balance: must be 43 (!) iotas less
	env.AssertAddressBalance(userAddress, balance.ColorIOTA, 1337-43)
	// check the on-chain account. Must contain 43 (!) iotas
	chain.AssertAccountBalance(userAgentID, balance.ColorIOTA, 43)

	// withdraw back all iotas
	req = solo.NewCall(accounts.Name, accounts.FuncWithdrawToAddress)
	_, err = chain.PostRequest(req, userWallet)
	require.NoError(t, err)

	// we are back to initial situation: IOTA is fee-less!
	env.AssertAddressBalance(userAddress, balance.ColorIOTA, 1337)
	chain.AssertAccountBalance(userAgentID, balance.ColorIOTA, 0) // empty
}

func TestSolo6(t *testing.T) {
	env := solo.New(t, false, false)
	chain := env.NewChain(nil, "ex6")

	err := chain.DeployWasmContract(nil, "example1", "../pkg/example1_bg.wasm")
	require.NoError(t, err)

	// global ID of the deployed contract
	contractID := coretypes.NewContractID(chain.ChainID, coretypes.Hn("example1"))
	// contract id in the form of the agent ID
	contractAgentID := coretypes.NewAgentIDFromContractID(contractID)

	userWallet := env.NewSignatureSchemeWithFunds()
	userAddress := userWallet.Address()
	userAgentID := coretypes.NewAgentIDFromAddress(userWallet.Address())

	env.AssertAddressBalance(userAddress, balance.ColorIOTA, 1337)
	chain.AssertAccountBalance(contractAgentID, balance.ColorIOTA, 0) // empty on-chain
	chain.AssertAccountBalance(userAgentID, balance.ColorIOTA, 0)     // empty on-chain

	req := solo.NewCall("example1", "storeString", "paramString", "Hello, world!").
		WithTransfer(balance.ColorIOTA, 42)
	_, err = chain.PostRequest(req, userWallet)
	require.NoError(t, err)

	chain.AssertAccountBalance(contractAgentID, balance.ColorIOTA, 42)
	chain.AssertAccountBalance(userAgentID, balance.ColorIOTA, 1)
	env.AssertAddressBalance(userAddress, balance.ColorIOTA, 1337-43)
}

func TestSolo7(t *testing.T) {
	env := solo.New(t, false, false)
	chain := env.NewChain(nil, "ex7")

	err := chain.DeployWasmContract(nil, "example1", "../pkg/example1_bg.wasm")
	require.NoError(t, err)

	// global ID of the deployed contract
	contractID := coretypes.NewContractID(chain.ChainID, coretypes.Hn("example1"))
	// contract id in the form of the agent ID
	contractAgentID := coretypes.NewAgentIDFromContractID(contractID)

	userWallet := env.NewSignatureSchemeWithFunds()
	userAddress := userWallet.Address()
	userAgentID := coretypes.NewAgentIDFromAddress(userWallet.Address())

	env.AssertAddressBalance(userAddress, balance.ColorIOTA, 1337)
	chain.AssertAccountBalance(contractAgentID, balance.ColorIOTA, 0) // empty on-chain
	chain.AssertAccountBalance(userAgentID, balance.ColorIOTA, 0)     // empty on-chain

	req := solo.NewCall("example1", "storeString").
		WithTransfer(balance.ColorIOTA, 42)
	_, err = chain.PostRequest(req, userWallet)
	require.Error(t, err)

	chain.AssertAccountBalance(contractAgentID, balance.ColorIOTA, 0)
	chain.AssertAccountBalance(userAgentID, balance.ColorIOTA, 1)
	env.AssertAddressBalance(userAddress, balance.ColorIOTA, 1337-1)
}
