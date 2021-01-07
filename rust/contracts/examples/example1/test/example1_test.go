package test

import (
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/balance"
	"github.com/iotaledger/wasp/packages/coretypes"
	"github.com/iotaledger/wasp/packages/kv/codec"
	"github.com/iotaledger/wasp/packages/solo"
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
		"paramString", theString).WithTransfer(map[balance.Color]int64{
		balance.ColorIOTA: 42,
	})
	_, err = chain.PostRequest(req, userWallet)
	require.NoError(t, err)

	contractID := coretypes.NewContractID(chain.ChainID, coretypes.Hn("example1"))
	contractAgentID := coretypes.NewAgentIDFromContractID(contractID)
	chain.AssertAccountBalance(contractAgentID, balance.ColorIOTA, 42)
	chain.AssertAccountBalance(userAgentID, balance.ColorIOTA, 1)

	env.AssertAddressBalance(userWallet.Address(), balance.ColorIOTA, 1337-43)
}