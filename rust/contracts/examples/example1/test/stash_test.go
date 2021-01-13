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

const WASM_FILE = "../pkg/example1_bg.wasm"

func TestExample1Test1(t *testing.T) {
	env := solo.New(t, false, false)
	chain := env.NewChain(nil, "ch1")

	err := chain.DeployWasmContract(nil, "example1", WASM_FILE)
	require.NoError(t, err)

	theString := "Hello, world!"
	req := solo.NewCall("example1", "storeString",
		"paramString", theString)
	_, err = chain.PostRequest(req, nil)
	require.NoError(t, err)

	// call the contract to extract value of the 'returnedString'. Must be equal 1
	res, err := chain.CallView("example1", "getString")
	require.NoError(t, err)
	returnedString, exists, err := codec.DecodeString(res.MustGet("paramString"))
	require.NoError(t, err)
	require.True(t, exists)
	require.EqualValues(t, theString, returnedString)
}

func TestExample1Test2(t *testing.T) {
	env := solo.New(t, false, false)
	chain := env.NewChain(nil, "ch1")

	err := chain.DeployWasmContract(nil, "example1", WASM_FILE)
	require.NoError(t, err)

	req := solo.NewCall("example1", "storeString")
	_, err = chain.PostRequest(req, nil)
	require.Error(t, err)

	// call the contract to extract value of the 'returnedString'. Must be equal 1
	res, err := chain.CallView("example1", "getString")
	require.NoError(t, err)
	returnedString, exists, err := codec.DecodeString(res.MustGet("paramString"))
	require.NoError(t, err)
	require.True(t, exists)
	require.EqualValues(t, "", returnedString)
}

func TestExample1Test3(t *testing.T) {
	env := solo.New(t, false, false)
	chain := env.NewChain(nil, "ch1")

	err := chain.DeployWasmContract(nil, "example1", WASM_FILE)
	require.NoError(t, err)

	userWallet := env.NewSignatureSchemeWithFunds()
	userAgentID := coretypes.NewAgentIDFromAddress(userWallet.Address())

	env.AssertAddressBalance(userWallet.Address(), balance.ColorIOTA, 1337)

	theString := "Hello, world!"
	req := solo.NewCall("example1", "storeString",
		"paramString", theString).WithTransfer(balance.ColorIOTA, 42)
	_, err = chain.PostRequest(req, userWallet)
	require.NoError(t, err)

	contractID := coretypes.NewContractID(chain.ChainID, coretypes.Hn("example1"))
	contractAgentID := coretypes.NewAgentIDFromContractID(contractID)
	chain.AssertAccountBalance(contractAgentID, balance.ColorIOTA, 42)
	chain.AssertAccountBalance(userAgentID, balance.ColorIOTA, 1)

	env.AssertAddressBalance(userWallet.Address(), balance.ColorIOTA, 1337-43)
}

func TestExample3(t *testing.T) {
	glb := solo.New(t, false, false)
	chain := glb.NewChain(nil, "ex3")

	userWallet := glb.NewSignatureSchemeWithFunds()
	userAddress := userWallet.Address()
	t.Logf("Address of the userWallet is: %s", userAddress)
	numIotas := glb.GetAddressBalance(userAddress, balance.ColorIOTA)
	t.Logf("balance of the userWallet is: %d iota", numIotas)
	glb.AssertAddressBalance(userAddress, balance.ColorIOTA, 1337)

	// send 42 iotas to the own account on-chain
	req := solo.NewCall(accounts.Name, accounts.FuncDeposit).
		WithTransfer(balance.ColorIOTA, 42)
	_, err := chain.PostRequest(req, userWallet)
	require.NoError(t, err)

	userAgentID := coretypes.NewAgentIDFromAddress(userAddress)
	chain.AssertAccountBalance(userAgentID, balance.ColorIOTA, 43) // 43!!

	// withdraw back all iotas
	req = solo.NewCall(accounts.Name, accounts.FuncWithdrawToAddress)
	_, err = chain.PostRequest(req, userWallet)
	require.NoError(t, err)

	chain.AssertAccountBalance(userAgentID, balance.ColorIOTA, 0) // empty
	glb.AssertAddressBalance(userAddress, balance.ColorIOTA, 1337)
}
